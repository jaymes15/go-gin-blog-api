package controllers

import (
	"blog/internal/requests"
	"blog/internal/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Controller struct {
	articleService services.ArticleServiceInterface
}

func NewArticleController() *Controller {
	return &Controller{
		articleService: services.NewArticleService(),
	}
}

func (controllers *Controller) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"featuredArticles": controllers.articleService.GetFeaturedArticles(),
		"storiesArticles":  controllers.articleService.GetStoriesArticles(),
		"app name":         viper.Get("App.Name"),
	})
}

func (controllers *Controller) Details(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// Find the article from the database
	article, err := controllers.articleService.Find(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"title": "Entity not found", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article":  article,
		"app name": viper.Get("App.Name"),
	})
}

func (controllers *Controller) Create(c *gin.Context) {

	var request requests.CreateArticle

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Could not Bind:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := controllers.articleService.CreateArticle(c, request)

	if err != nil {
		log.Printf("Error:::::: %s", err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article":  article,
		"app name": viper.Get("App.Name"),
	})

}
