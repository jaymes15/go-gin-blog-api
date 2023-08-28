package services

import (
	"blog/internal/requests"
	"blog/internal/responses"

	"github.com/gin-gonic/gin"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetStoriesArticles() responses.Articles
	CreateArticle(c *gin.Context, request requests.CreateArticle) (responses.Article, error)
	Find(id int) (responses.Article, error)
}
