package save

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"urlshortener/internal/lib/api/response"
	"urlshortener/internal/lib/logger/sl"
	"urlshortener/internal/lib/random"
)

const aliasLength = 8

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op))
		slog.String("request_id", middleware.GetReqID(r.Context()))
		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))
			render.JSON(w, r, response.Error("failed to decode "))
			return
		}

		log.Info("reques body decoded", slog.Any("req", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("failed to validate request", sl.Err(err))
			render.JSON(w, r, response.Error("failed to validate "))
			return
		}
		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

		id, err := urlSaver.SaveURL(req.URL, alias)
		if err != nil {
			log.Error("failed to save url", sl.Err(err))
			render.JSON(w, r, response.Error("failed to save url"))
			return
		}
		log.Info("url saved", slog.Int("id", id))

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    alias,
		})
	}
}
