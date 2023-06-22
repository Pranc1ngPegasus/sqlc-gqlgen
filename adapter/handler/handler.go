package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	gqlgenHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/wire"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	schema graphql.ExecutableSchema
	router http.Handler
}

func NewHandler(
	schema graphql.ExecutableSchema,
) *Handler {
	mux := http.NewServeMux()

	h := &Handler{
		schema: schema,
		router: mux,
	}

	mux.HandleFunc("/healthcheck", h.healthcheck)
	mux.HandleFunc("/graphql", h.graphql)

	return h
}

func (h *Handler) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) graphql(w http.ResponseWriter, r *http.Request) {
	gqlgenHandler.NewDefaultServer(h.schema).ServeHTTP(w, r)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
