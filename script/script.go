package script

import (
	"Backend-Challenge/models"
	"Backend-Challenge/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SeedArticles() {
	resp, err := http.Get("https://api.spaceflightnewsapi.net/v3/articles")
	if err != nil {
		log.Fatalf("Erro ao obter os dados dos artigos: %s", err)
	}
	defer resp.Body.Close()

	var articles []*models.Article
	err = json.NewDecoder(resp.Body).Decode(&articles)
	if err != nil {
		log.Fatalf("Erro ao decodificar os dados dos artigos: %s", err)
	}

	for _, article := range articles {
		err := services.AddArticle(article)
		if err != nil {
			log.Printf("Erro ao adicionar o artigo '%s': %s", article.Title, err)
		} else {
			log.Printf("Artigo '%s' adicionado com sucesso", article.Title)
		}
	}

	fmt.Println("Alimentação dos dados concluída")
}
