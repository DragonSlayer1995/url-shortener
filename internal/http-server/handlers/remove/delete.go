package remove

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
)

type Response struct {
	resp.Response
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type URLDeleter interface {
	DeleteURL(alias string) (int64, error)
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.remove.New"

		log = log.With(slog.String("op", op))

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("alias is empty"))

			return
		}

		count, err := urlDeleter.DeleteURL(alias)

		if err != nil {
			log.Error("failed to remove url by alias", sl.Err(err))

			render.JSON(w, r, resp.Error("internal error"))

			return
		}

		if count == 0 {
			log.Info("nothing to delete", slog.Int64("count", count))

			responseOK(w, r, false, fmt.Sprintf("Can't find alias %v", alias))

			return
		}

		log.Info("url deleted", slog.String("alias", alias))

		responseOK(w, r, true, "")
	}

}

func responseOK(w http.ResponseWriter, r *http.Request, success bool, message string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Success:  success,
		Message:  message,
	})
}
