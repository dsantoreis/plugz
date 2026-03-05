package tests

import (
	"net/http"
	"testing"
)

func TestE2ERateLimitEventually429(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	client := &http.Client{}

	seen429 := false
	for i := 0; i < 120; i++ {
		req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/catalog", nil)
		req.Header.Set("Authorization", "Bearer dev-token")
		res, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode == http.StatusTooManyRequests {
			seen429 = true
			break
		}
	}
	if !seen429 {
		t.Fatalf("expected at least one 429 response")
	}
}
