package yog

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// Return a recognizable OAuth error instead of a token so authentication can be
// checked through the public storage API without sending a request to GCS.
func credentialFiles(t *testing.T) func(string) string {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error(err)
			http.Error(w, "invalid token request", http.StatusBadRequest)
			return
		}
		parts := strings.Split(r.Form.Get("assertion"), ".")
		if len(parts) != 3 {
			t.Error("missing service-account assertion")
			http.Error(w, "missing assertion", http.StatusBadRequest)
			return
		}
		payload, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			t.Error(err)
			http.Error(w, "invalid assertion", http.StatusBadRequest)
			return
		}
		var claims struct {
			Issuer string `json:"iss"`
		}
		if err := json.Unmarshal(payload, &claims); err != nil {
			t.Error(err)
			http.Error(w, "invalid claims", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{
			"error":             "invalid_grant",
			"error_description": "synthetic credential: " + claims.Issuer,
		}); err != nil {
			t.Error(err)
		}
	}))
	t.Cleanup(server.Close)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	privateKey := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	dir := t.TempDir()
	return func(identity string) string {
		t.Helper()
		data, err := json.Marshal(map[string]string{
			"type":         "service_account",
			"project_id":   "synthetic-project",
			"client_email": identity,
			"private_key":  string(privateKey),
			"token_uri":    server.URL,
		})
		if err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(dir, identity+".json")
		if err := os.WriteFile(path, data, 0600); err != nil {
			t.Fatal(err)
		}
		return path
	}
}

func configureGCSTest(t *testing.T) {
	t.Helper()
	t.Setenv("USE_GSG", "false")
	t.Setenv("STORAGE_EMULATOR_HOST", "")
	t.Setenv("STORAGE_USE_GRPC", "")
	t.Setenv("GOOGLE_API_USE_MTLS_ENDPOINT", "never")
	t.Setenv("GOOGLE_API_USE_CLIENT_CERTIFICATE", "false")
}

func checkClientCredential(t *testing.T, client *GCSClient, identity string) {
	t.Helper()
	defer func() {
		if err := client.Client.Close(); err != nil {
			t.Error(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Client.Bucket("synthetic-bucket").Attrs(ctx)
	if err == nil || !strings.Contains(err.Error(), "synthetic credential: "+identity) {
		t.Errorf("client for %s used unexpected credentials: %v", identity, err)
	}
}

func TestNewGCSClientConcurrentCredentials(t *testing.T) {
	for _, adcConfigured := range []bool{false, true} {
		t.Run(fmt.Sprintf("adc-configured=%t", adcConfigured), func(t *testing.T) {
			configureGCSTest(t)
			writeCredential := credentialFiles(t)
			adcIdentity := "default@example.test"
			adcPath := writeCredential(adcIdentity)
			t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", adcPath)
			if !adcConfigured {
				if err := os.Unsetenv("GOOGLE_APPLICATION_CREDENTIALS"); err != nil {
					t.Fatal(err)
				}
			}
			before, wasSet := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
			const jobs = 32
			paths := make([]string, jobs)
			identities := make([]string, jobs)
			clients := make([]*GCSClient, jobs)
			for i := range paths {
				identities[i] = fmt.Sprintf("task-%02d@example.test", i)
				paths[i] = writeCredential(identities[i])
			}
			var wg sync.WaitGroup
			start := make(chan struct{})
			for i := range clients {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					<-start
					clients[i] = NewGCSClient("synthetic-bucket", "task-prefix", paths[i], "gs://synthetic-bucket/task-prefix").(*GCSClient)
				}(i)
			}
			close(start)
			wg.Wait()
			after, isSet := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
			if after != before || isSet != wasSet {
				t.Error("explicit client credentials changed process-wide ADC configuration")
			}
			for i, client := range clients {
				if client.Bucket != "synthetic-bucket" || client.Prefix != "task-prefix" {
					t.Error("client lost bucket or prefix configuration")
				}
				checkClientCredential(t, client, identities[i])
			}
			if adcConfigured {
				// Explicit clients must not change the credentials of later clients
				// that rely on application-default credential discovery.
				client := NewGCSClient("synthetic-bucket", "", "", "gs://synthetic-bucket").(*GCSClient)
				checkClientCredential(t, client, adcIdentity)
			}
		})
	}
}

func TestNewGCSClientInvalidCredentials(t *testing.T) {
	configureGCSTest(t)
	writeCredential := credentialFiles(t)
	adcPath := writeCredential("default@example.test")
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", adcPath)
	for _, path := range []string{filepath.Join(t.TempDir(), "missing.json"), filepath.Join(t.TempDir(), "invalid.json")} {
		t.Run(filepath.Base(path), func(t *testing.T) {
			if filepath.Base(path) == "invalid.json" {
				if err := os.WriteFile(path, []byte("not JSON"), 0600); err != nil {
					t.Fatal(err)
				}
			}
			defer func() {
				if recovered := recover(); recovered == nil {
					t.Error("invalid explicit credentials must still panic")
				}
				if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != adcPath {
					t.Error("failed constructor changed process-wide ADC configuration")
				}
			}()
			NewGCSClient("synthetic-bucket", "", path, "gs://synthetic-bucket")
		})
	}
}
