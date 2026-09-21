package yog

import "testing"

func TestRemoveSchemePrefix(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "gcs address keeps bucket and prefix",
			addr: "gs://oregon-cdn.nextbillion.io/mdm/api-auto-oregon/task-1",
			want: "oregon-cdn.nextbillion.io/mdm/api-auto-oregon/task-1",
		},
		{
			name: "s3 address keeps bucket and prefix",
			addr: "s3://nb-mdm-oregon/mdm/api-auto-oregon/task-1",
			want: "nb-mdm-oregon/mdm/api-auto-oregon/task-1",
		},
		{
			name: "oci address drops the region qualifier with the scheme",
			addr: "oci://nb-mdm-oregon@us-phoenix-1/mdm/api-auto-oregon/task-1",
			want: "nb-mdm-oregon/mdm/api-auto-oregon/task-1",
		},
		{
			name: "oci address drops an explicit namespace too",
			addr: "oci://nb-mdm-oregon@axabc123.us-phoenix-1/mdm/task-1",
			want: "nb-mdm-oregon/mdm/task-1",
		},
		{
			name: "oci bucket with no prefix",
			addr: "oci://nb-mdm-oregon@us-phoenix-1",
			want: "nb-mdm-oregon",
		},
		{
			name: "an @ below the bucket belongs to the object name",
			addr: "oci://nb-mdm-oregon@us-phoenix-1/mdm/task@1/0",
			want: "nb-mdm-oregon/mdm/task@1/0",
		},
		{
			name: "a scheme is only stripped at the start",
			addr: "gs://bucket/gs://not-a-scheme",
			want: "bucket/gs://not-a-scheme",
		},
		{
			name: "an address with no scheme is left alone",
			addr: "bucket/mdm/task-1",
			want: "bucket/mdm/task-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSchemePrefix(tt.addr); got != tt.want {
				t.Errorf("removeSchemePrefix(%q) = %q, want %q", tt.addr, got, tt.want)
			}
		})
	}
}

// TestNew_BuildsReachableHTTPSURL is the property that matters to the caller: the
// merge path reads chunks over plain HTTPS, so whatever scheme the task address was
// written in, what comes out has to be a URL a GET can reach.
func TestNew_BuildsReachableHTTPSURL(t *testing.T) {
	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "gcs",
			addr: "gs://oregon-cdn.nextbillion.io/mdm/api-auto-oregon/task-1",
			want: "https://oregon-cdn.nextbillion.io/mdm/api-auto-oregon/task-1",
		},
		{
			name: "oci",
			addr: "oci://oregon-cdn.nextbillion.io@us-phoenix-1/mdm/api-auto-oregon/task-1",
			want: "https://oregon-cdn.nextbillion.io/mdm/api-auto-oregon/task-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := New("task-1", t.TempDir(), tt.addr)
			c, ok := y.storage.(*HTTPGetClient)
			if !ok {
				t.Fatalf("New built a %T, want *HTTPGetClient", y.storage)
			}
			if c.Host != tt.want {
				t.Errorf("host = %q, want %q", c.Host, tt.want)
			}
		})
	}
}
