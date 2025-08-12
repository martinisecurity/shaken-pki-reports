package linter_test

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/linter"
)

func TestLintUrl_HttpStatus404(t *testing.T) {
	// Create a test server that returns a 404 status
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Parse the test server URL
	u, _ := url.Parse(ts.URL)

	// Call the function under test
	res := linter.LintUrl(u)

	// Define the expected result
	want := &lint.LintResult{
		Status:  lint.Error,
		Details: "HTTP response shall have StatusCode 200, but it is 404 Not Found",
	}

	// Compare the actual result with the expected result
	if test := res.Results["e_http_status_200"]; !reflect.DeepEqual(test, want) {
		t.Errorf("lint.LintUrl() = %v, want %v", test, want)
	}
}

func TestLintUrl_AtisContentType(t *testing.T) {
	// Mock server without the expected Content-Type header
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// No Content-Type or a different one, to trigger the warning
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("mock body"))
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	res := linter.LintUrl(u)

	want := &lint.LintResult{
		Status:  lint.Warn,
		Details: "HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain",
	}
	if test := res.Results["w_atis_content_type"]; !reflect.DeepEqual(test, want) {
		t.Errorf("lint.LintUrl() = %v, want %v", test, want)
	}
}

func TestLintUrl_AtisProtocol(t *testing.T) {
	code := "w_atis_protocol"

	// Helper: start a TLS server bound to a specific addr (e.g., 127.0.0.1:8443)
	startTLSServerOn := func(addr string) (*httptest.Server, func()) {
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		})
		ts := httptest.NewUnstartedServer(handler)
		l, err := net.Listen("tcp", addr)
		if err != nil {
			// If the port is taken or unavailable, skip this test case.
			t.Skipf("unable to listen on %s: %v", addr, err)
		}
		ts.Listener = l
		ts.StartTLS()
		return ts, ts.Close
	}

	t.Run("https on 8443 (pass)", func(t *testing.T) {
		// Bind to 127.0.0.1:8443 to satisfy allowed port rule
		ts, cleanup := startTLSServerOn("127.0.0.1:8443")
		defer cleanup()
		u, _ := url.Parse(ts.URL)
		res := linter.LintUrl(u)
		want := &lint.LintResult{Status: lint.Pass}
		if got := res.Results[code]; !reflect.DeepEqual(got, want) {
			t.Errorf("lint.LintUrl() = %v, want %v", got, want)
		}
	})

	t.Run("https on random port (warn)", func(t *testing.T) {
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		}))
		defer ts.Close()
		u, _ := url.Parse(ts.URL)
		res := linter.LintUrl(u)
		want := &lint.LintResult{
			Status:  lint.Warn,
			Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		}
		if got := res.Results[code]; !reflect.DeepEqual(got, want) {
			t.Errorf("lint.LintUrl() = %v, want %v", got, want)
		}
	})

	t.Run("http url (warn)", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		}))
		defer ts.Close()
		u, _ := url.Parse(ts.URL)
		res := linter.LintUrl(u)
		want := &lint.LintResult{
			Status:  lint.Warn,
			Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		}
		if got := res.Results[code]; !reflect.DeepEqual(got, want) {
			t.Errorf("lint.LintUrl() = %v, want %v", got, want)
		}
	})
}

func TestLintUrl_RequestTimeout(t *testing.T) {
	// Server that sleeps longer than client timeout (3s)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sleep to trigger client timeout
		<-time.After(5 * time.Second)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("late"))
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	res := linter.LintUrl(u)

	got := res.Results["e_request_timeout"]
	if got == nil || got.Status != lint.Error {
		t.Fatalf("expected e_request_timeout error, got %v", got)
	}
}

func TestLintUrl_TLSInsecureFallback(t *testing.T) {
	// Start a TLS server with self-signed cert. First request with verification should fail.
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	res := linter.LintUrl(u)

	got := res.Results["e_tls_transport"]
	if got == nil || got.Status != lint.Error {
		t.Fatalf("expected e_tls_transport error, got %v", got)
	}
	if got.Details == "" {
		t.Errorf("expected non-empty details for e_tls_transport")
	}
}

func TestLintUrl_BadURLScheme(t *testing.T) {
	// Unsupported scheme should produce e_bad_url
	u, _ := url.Parse("bad://host/path")
	res := linter.LintUrl(u)

	got := res.Results["e_bad_url"]
	if got == nil || got.Status != lint.Error {
		t.Fatalf("expected e_bad_url error, got %v", got)
	}
	if got.Details == "" {
		t.Errorf("expected non-empty details for e_bad_url")
	}
}

func TestLintUrl_RedirectDetection(t *testing.T) {
	// First request returns 200, second (used for redirect check) returns a redirect
	var count int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if count == 0 {
			count++
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("ok"))
			return
		}
		// On subsequent call, redirect
		http.Redirect(w, r, "/next", http.StatusFound)
	}))
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	res := linter.LintUrl(u)

	got := res.Results["e_atis_redirect"]
	want := &lint.LintResult{Status: lint.Error, Details: "The STI-VS shall not follow HTTP redirections"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected redirect error %v, got %v", want, got)
	}
}
