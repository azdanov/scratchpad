package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_secureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %q; got %q", "deny", frameOptions)
	}
	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
	}
	xContentTypeOptions := rs.Header.Get("X-Content-Type-Options")
	if xContentTypeOptions != "nosniff" {
		t.Errorf("want %q; got %q", "nosniff", xContentTypeOptions)
	}
	referrerPolicy := rs.Header.Get("Referrer-Policy")
	if referrerPolicy != "no-referrer-when-downgrade" {
		t.Errorf("want %q; got %q", "no-referrer-when-downgrade", referrerPolicy)
	}
	contentSecurityPolicy := rs.Header.Get("Content-Security-Policy")
	if contentSecurityPolicy != "default-src 'self' http: https: data: blob: 'unsafe-inline'; frame-ancestors 'self';" {
		t.Errorf("want %q; got %q", "default-src 'self' http: https: data: blob: 'unsafe-inline'; frame-ancestors 'self';", contentSecurityPolicy)
	}
	permissionsPolicy := rs.Header.Get("Permissions-Policy")
	if permissionsPolicy != "interest-cohort=()" {
		t.Errorf("want %q; got %q", "interest-cohort=()", permissionsPolicy)
	}
	strictTransportSecurity := rs.Header.Get("Strict-Transport-Security")
	if strictTransportSecurity != "max-age=31536000; includeSubDomains" {
		t.Errorf("want %q; got %q", "max-age=31536000; includeSubDomains", strictTransportSecurity)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
