package httpbymatryer

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func handleTest(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// use thing to handle request
			w.Write([]byte("Hello from Snippetbox"))
			logger.Info("handleTest")
		},
	)
}
func someMiddleware1(h http.Handler) http.Handler {
	fmt.Println("Set up happens once since not returned as part of next handler")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("In Middleware")
			h.ServeHTTP(w, r)
			fmt.Println("Should happen after request")
		},
	)
}
func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
) {
	mux.Handle("/test", handleTest(logger))
	mux.Handle("/", http.NotFoundHandler())
}

type config struct {
	addr string
}

func newServer(
	logger *slog.Logger,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
	)
	var handler http.Handler = mux
	handler = someMiddleware1(handler)
	return handler
}

func Run() {
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// flag.Parse()
	var config = config{
		addr: ":4000",
	}
	// addr := ":4000"
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	srv := newServer(
		logger,
	)
	httpServer := &http.Server{
		Addr:    config.addr,
		Handler: srv,
	}
	fmt.Printf("listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}
}
