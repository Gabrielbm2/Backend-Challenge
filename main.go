package main

import (
	"Backend-Challenge/models"
	"Backend-Challenge/routes"
	"Backend-Challenge/script"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.CreateArticleTable()

	// Chamar a função SeedArticles para armazenar os dados dos artigos
	script.SeedArticles()

	// Configurar o CRON para executar diariamente às 9h
	cron := cron.New()
	cron.AddFunc("0 9 * * *", func() {
		// Chamar a função SeedArticles para atualizar os dados dos artigos diariamente
		script.SeedArticles()
	})
	cron.Start()
	defer cron.Stop()

	r := chi.NewRouter()
	routes.LoadRoutes(r)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	address := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}

	fmt.Printf("Servidor rodando em %v", address)
	log.Fatal(srv.ListenAndServe())
}
