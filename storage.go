package yog

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/nextbillion-ai/gsg/lib/object"
	"github.com/sirupsen/logrus"
)

type ICloudStorage interface {
	DownloadSingleFile(ctx context.Context, objectName string, path string) error
}

type GCSClient struct {
	Client     *storage.Client
	Bucket     string
	Prefix     string
	APIKeyPath string
}

// GSGClient can deal with aws and gcs at the same time
type GSGClient struct {
	FullPrefix string
}

// DownloadSingleFile gs://nb-data/mdm/test/d111ae24-0f3e-3617-b64d-ca9ae0243c7c/0
func (g *GSGClient) DownloadSingleFile(ctx context.Context, objectName string, path string) (err error) {
	if path != "" && path[len(path)-1] != '/' {
		path = path + "/"
	}

	fullPath := g.FullPrefix + "/" + objectName
	gsgObject, err := object.New(fullPath)
	if err != nil {
		return fmt.Errorf("object.New: %w", err)
	}

	f, err := os.Create(path + objectName)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	err = gsgObject.Read(f)
	if errors.Is(err, object.ErrObjectNotFound) {
		return errors.New("file not exist")
	}
	if err != nil {
		return fmt.Errorf("object.Read: %w", err)
	}

	s, err := f.Stat()
	if err != nil {
		return fmt.Errorf("file stat failed: %w", err)
	}
	if s.Size() == 0 {
		return errors.New("file not exist")
	}

	return nil
}

//export GOOGLE_APPLICATION_CREDENTIALS="/Users/xurui/Workspace/env/dcsa.json"
func NewGCSClient(bucket, prefix, APIKey, fullpath string) ICloudStorage {
	useGSG := os.Getenv("USE_GSG") == "true"
	if useGSG {
		logrus.Info("Using GSG for cloud storage")
		// this call is just to test access, we do not need to retain the object
		_, err := object.New(fullpath)
		if err != nil {
			panic(fmt.Sprintf("failed to init gsg object.New: %s", err.Error()))
		}
		return &GSGClient{
			FullPrefix: strings.TrimSuffix(fullpath, "/"),
		}
	}

	if APIKey != "" {
		err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", APIKey)
		if err != nil {
			panic(fmt.Sprintf("failed to init storage.NewClient: %s", err.Error()))
		}
	}

	client, err := storage.NewClient(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to init storage.NewClient: %s", err.Error()))
	}

	return &GCSClient{
		Client: client,
		Bucket: bucket,
		Prefix: prefix,
	}
}

// DownloadSingleFile gs://nb-data/mdm/test/d111ae24-0f3e-3617-b64d-ca9ae0243c7c/0
func (g *GCSClient) DownloadSingleFile(ctx context.Context, objectName string, path string) error {

	if path != "" && path[len(path)-1] != '/' {
		path = path + "/"
	}

	_, err := g.Client.Bucket(g.Bucket).Object(g.Prefix + "/" + objectName).Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return errors.New("file not exist")
		}
		return fmt.Errorf("check remote storage failed bucket:%v object:%v err:%v", g.Bucket, g.Prefix+"/"+objectName, err)
	}

	//ctx, cancel := context.WithTimeout(ctx, time.Minute*50)
	//defer cancel()

	f, err := os.Create(path + objectName)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	rc, err := g.Client.Bucket(g.Bucket).Object(g.Prefix + "/" + objectName).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("Object(%q).NewReader: %w", g.Prefix, err)
	}
	defer func(rc *storage.Reader) {
		_ = rc.Close()
	}(rc)

	if _, err := io.Copy(f, rc); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("f.Close: %w", err)
	}
	return nil
}

type HTTPGetClient struct {
	Host string
}

func NewHTTPGetClient(Host string) *HTTPGetClient {
	return &HTTPGetClient{
		Host: Host,
	}
}

func (g *HTTPGetClient) DownloadSingleFile(ctx context.Context, objectName string, path string) error {
	if path != "" && path[len(path)-1] != '/' {
		path = path + "/"
	}

	file, err := os.OpenFile(path+objectName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	c := &http.Client{
		Timeout: 0,
	}

	rsp, err := c.Get(g.Host + "/" + objectName)
	if err != nil {
		return err
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	n, err := io.Copy(file, rsp.Body)
	if n < 200 {
		b := make([]byte, 127)
		_, err := file.ReadAt(b, 0)
		if err != nil && err != io.EOF {
			return err
		}
		if string(b) == "<?xml version='1.0' encoding='UTF-8'?><Error><Code>NoSuchKey</Code><Message>The specified key does not exist.</Message></Error>" {
			_ = os.Remove(path + objectName)
			return errors.New("file not exist")
		}
	}

	return err
}
