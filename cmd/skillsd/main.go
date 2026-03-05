package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dsantoreis/ai-agent-skills-demo/internal/api"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/executor"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/registry"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/telemetry"
	"github.com/dsantoreis/ai-agent-skills-demo/internal/watcher"
)

func main() {
	if len(os.Args) < 2 {
		fatal("usage: skillsd <list|run|watch|serve> [flags]")
	}

	switch os.Args[1] {
	case "list":
		runList(os.Args[2:])
	case "run":
		runSkill(os.Args[2:])
	case "watch":
		runWatch(os.Args[2:])
	case "serve":
		runServe(os.Args[2:])
	default:
		fatal("unknown command")
	}
}

func runList(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	skillsDir := fs.String("skills-dir", "./examples/skills", "skills directory")
	_ = fs.Parse(args)

	r, err := registry.New(*skillsDir)
	if err != nil {
		fatal(err.Error())
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(r.List())
}

func runSkill(args []string) {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	skillsDir := fs.String("skills-dir", "./examples/skills", "skills directory")
	name := fs.String("name", "", "skill name")
	input := fs.String("input", "", "input payload sent to stdin")
	timeoutMs := fs.Int("timeout-ms", 3000, "default timeout in milliseconds")
	_ = fs.Parse(args)

	if *name == "" {
		fatal("run requires --name")
	}

	r, err := registry.New(*skillsDir)
	if err != nil {
		fatal(err.Error())
	}
	s, ok := r.Get(*name)
	if !ok {
		fatal("skill not found")
	}

	res := executor.Run(context.Background(), s, *input, time.Duration(*timeoutMs)*time.Millisecond)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(res)
}

func runWatch(args []string) {
	fs := flag.NewFlagSet("watch", flag.ExitOnError)
	skillsDir := fs.String("skills-dir", "./examples/skills", "skills directory")
	_ = fs.Parse(args)

	r, err := registry.New(*skillsDir)
	if err != nil {
		fatal(err.Error())
	}
	fmt.Printf("watching %s (loaded=%d)\n", *skillsDir, len(r.List()))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := watcher.Start(ctx, *skillsDir, r); err != nil {
		fatal(err.Error())
	}
}

func runServe(args []string) {
	fs := flag.NewFlagSet("serve", flag.ExitOnError)
	skillsDir := fs.String("skills-dir", "./examples/skills", "skills directory")
	addr := fs.String("addr", ":8080", "HTTP listen address")
	timeoutMs := fs.Int("timeout-ms", 3000, "default timeout in milliseconds")
	_ = fs.Parse(args)

	r, err := registry.New(*skillsDir)
	if err != nil {
		fatal(err.Error())
	}
	shutdownTelemetry, err := telemetry.Init(context.Background())
	if err != nil {
		fatal(err.Error())
	}
	defer func() { _ = shutdownTelemetry(context.Background()) }()

	srv := &http.Server{
		Addr:              *addr,
		Handler:           api.NewServer(r, time.Duration(*timeoutMs)*time.Millisecond).Router(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fatal(err.Error())
		}
	}()
	fmt.Printf("skills registry listening on %s\n", *addr)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
