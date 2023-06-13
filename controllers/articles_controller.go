package controllers

import (
	"Backend-Challenge/models"
	"Backend-Challenge/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ArticleController struct {
	ArticleService *services.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		ArticleService: services.NewArticleService(),
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Back-end Challenge 2021 🏅 - Space Flight News"))
}

func ListArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := services.GetAllArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, articles, http.StatusOK)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id") // Extrair o ID do artigo da URL

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := services.GetArticleByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, article, http.StatusOK)
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	err = services.AddArticle(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, article, http.StatusCreated)
}

func UpdateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Failed to decode JSON body", http.StatusBadRequest)
		return
	}

	err = services.UpdateArticle(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponse(w, article, http.StatusOK)
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id") // Extrair o ID do artigo da URL

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	err = services.DeleteArticle(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
