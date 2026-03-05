package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/executor"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/registry"
)

type Server struct {
	registry      *registry.Registry
	defaultTimout time.Duration
	installedMu   sync.RWMutex
	installed     map[string]time.Time
	runsCounter   *prometheus.CounterVec
}

type InstallRequest struct {
	Name string `json:"name"`
}

type TestRequest struct {
	Input string `json:"input"`
}

func NewServer(r *registry.Registry, timeout time.Duration) *Server {
	runs := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "skills_runs_total",
		Help: "Total skill test runs",
	}, []string{"skill", "status"})
	if err := prometheus.Register(runs); err != nil {
		if ar, ok := err.(prometheus.AlreadyRegisteredError); ok {
			runs = ar.ExistingCollector.(*prometheus.CounterVec)
		}
	}

	return &Server{
		registry:      r,
		defaultTimout: timeout,
		installed:     map[string]time.Time{},
		runsCounter:   runs,
	}
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	r.Get("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		s.installedMu.RLock()
		installed := len(s.installed)
		s.installedMu.RUnlock()
		respondJSON(w, http.StatusOK, map[string]any{"status": "ready", "installed": installed})
	})

	r.Group(func(r chi.Router) {
		r.Use(JSONLoggingMiddleware)
		r.Use(AuthMiddleware)
		r.Use(RateLimitMiddleware(60, time.Minute))

		r.Handle("/metrics", promhttp.Handler())
		r.Get("/api/v1/catalog", s.handleCatalog)
		r.Get("/api/v1/installed", s.handleInstalled)
		r.Post("/api/v1/install", s.handleInstall)
		r.Post("/api/v1/test/{name}", s.handleTest)
	})

	return otelhttp.NewHandler(r, "skills-api")
}

func (s *Server) handleCatalog(w http.ResponseWriter, _ *http.Request) {
	respondJSON(w, http.StatusOK, s.registry.List())
}

func (s *Server) handleInstalled(w http.ResponseWriter, _ *http.Request) {
	s.installedMu.RLock()
	defer s.installedMu.RUnlock()
	respondJSON(w, http.StatusOK, s.installed)
}

func (s *Server) handleInstall(w http.ResponseWriter, req *http.Request) {
	var payload InstallRequest
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		respondErr(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if _, ok := s.registry.Get(payload.Name); !ok {
		respondErr(w, http.StatusNotFound, "skill not found")
		return
	}

	s.installedMu.Lock()
	s.installed[payload.Name] = time.Now().UTC()
	s.installedMu.Unlock()

	respondJSON(w, http.StatusOK, map[string]string{"status": "installed", "name": payload.Name})
}

func (s *Server) handleTest(w http.ResponseWriter, req *http.Request) {
	name := chi.URLParam(req, "name")
	sk, ok := s.registry.Get(name)
	if !ok {
		respondErr(w, http.StatusNotFound, "skill not found")
		return
	}

	var payload TestRequest
	if req.Body != nil {
		_ = json.NewDecoder(req.Body).Decode(&payload)
	}

	result := executor.Run(context.Background(), sk, payload.Input, s.defaultTimout)
	s.runsCounter.WithLabelValues(name, result.Status).Inc()
	respondJSON(w, http.StatusOK, result)
}

func respondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func respondErr(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, map[string]string{"error": msg})
}
