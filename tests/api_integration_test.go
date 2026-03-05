package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/api"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/registry"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	r, err := registry.New("../examples/skills")
	if err != nil {
		t.Fatal(err)
	}
	return httptest.NewServer(api.NewServer(r, time.Second).Router())
}

func TestHealthAndReadyDoNotRequireAuth(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	healthResp, err := http.Get(ts.URL + "/healthz")
	if err != nil {
		t.Fatal(err)
	}
	if healthResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 for /healthz got %d", healthResp.StatusCode)
	}

	readyResp, err := http.Get(ts.URL + "/readyz")
	if err != nil {
		t.Fatal(err)
	}
	if readyResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 for /readyz got %d", readyResp.StatusCode)
	}
}

func TestCatalogRequiresAuth(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/v1/catalog")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 got %d", resp.StatusCode)
	}
}

func TestInstallAndTestFlow(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()
	client := &http.Client{}

	readyBefore, err := http.Get(ts.URL + "/readyz")
	if err != nil {
		t.Fatal(err)
	}
	if readyBefore.StatusCode != http.StatusOK {
		t.Fatalf("ready before status: %d", readyBefore.StatusCode)
	}
	var readyBeforePayload map[string]any
	if err := json.NewDecoder(readyBefore.Body).Decode(&readyBeforePayload); err != nil {
		t.Fatal(err)
	}
	if got := int(readyBeforePayload["installed"].(float64)); got != 0 {
		t.Fatalf("expected installed=0 before install, got %d", got)
	}

	installBody := bytes.NewBufferString(`{"name":"echo"}`)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/install", installBody)
	req.Header.Set("Authorization", "Bearer dev-token")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("install status: %d", res.StatusCode)
	}

	readyAfter, err := http.Get(ts.URL + "/readyz")
	if err != nil {
		t.Fatal(err)
	}
	if readyAfter.StatusCode != http.StatusOK {
		t.Fatalf("ready after status: %d", readyAfter.StatusCode)
	}
	var readyAfterPayload map[string]any
	if err := json.NewDecoder(readyAfter.Body).Decode(&readyAfterPayload); err != nil {
		t.Fatal(err)
	}
	if got := int(readyAfterPayload["installed"].(float64)); got != 1 {
		t.Fatalf("expected installed=1 after install, got %d", got)
	}

	testBody := bytes.NewBufferString(`{"input":"hello"}`)
	req2, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/v1/test/echo", testBody)
	req2.Header.Set("Authorization", "Bearer dev-token")
	req2.Header.Set("Content-Type", "application/json")
	res2, err := client.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	if res2.StatusCode != http.StatusOK {
		t.Fatalf("test status: %d", res2.StatusCode)
	}
	var payload map[string]any
	if err := json.NewDecoder(res2.Body).Decode(&payload); err != nil {
		t.Fatal(err)
	}
	if payload["status"] != "ok" {
		t.Fatalf("expected status ok, got %+v", payload)
	}
}
