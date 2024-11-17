package cmd

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/grpcreflect"
	"github.com/ei-sugimoto/microtodo/server/member/gen/health/v1/healthv1connect"
	"github.com/ei-sugimoto/microtodo/server/member/gen/member/v1/memberv1connect"
	"github.com/ei-sugimoto/microtodo/server/member/handler"
	"github.com/ei-sugimoto/microtodo/server/member/handler/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Serve() {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"health.v1.HealthService",
		"member.v1.MemberService",
	)

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	healthHandler := handler.NewHealthHandler()
	path, healthHand := healthv1connect.NewHealthServiceHandler(healthHandler)
	mux.Handle(path, middleware.LoggingMiddleware(healthHand))

	memberHandler := handler.NewMemberHandler()
	path, memberHand := memberv1connect.NewMemberServiceHandler(memberHandler)
	mux.Handle(path, middleware.LoggingMiddleware(memberHand))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	srv := &http.Server{
		Addr:    ":5556",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	go func() {
		slog.Info("server listening", "Addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatalf("shutdown: %s\n", err)
	}

}
