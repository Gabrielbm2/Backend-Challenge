package routes

import (
	"Backend-Challenge/controllers"
	"github.com/go-chi/chi"
	"net/http"
)

func LoadRoutes(r *chi.Mux) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	r.Get("/", controllers.HomeHandler)
	r.Get("/articles", controllers.ListArticlesHandler)
	r.Get("/articles/{id}", controllers.GetArticleHandler)
	r.Post("/articles", controllers.AddArticleHandler)
	r.Put("/articles/{id}", controllers.UpdateArticleHandler)
	r.Delete("/articles/{id}", controllers.DeleteArticleHandler)
}
