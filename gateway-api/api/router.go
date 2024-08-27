package api

import (
	"api/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car moyka
// @version 1.0
// @description API for Book Shop
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	// enforcer, err := casbin.NewEnforcer("./api/model.conf", "./api/policy.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Use(middlerware.NewAuth(enforcer))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	authorGroup := router.Group("/authors")
	{
		authorGroup.POST("/", h.CreateAuthor)
		authorGroup.GET("/:id", h.GetByIdAuthor)
		authorGroup.PUT("/", h.UpdateAuthor)
		authorGroup.GET("/", h.GetAllAuthors)
		authorGroup.DELETE("/:id", h.DeleteAuthor)
	}

	bookGroup := router.Group("/books")
	{
		bookGroup.POST("/", h.CreateBooks)
		bookGroup.GET("/:id", h.GetByIdBook)
		bookGroup.PUT("/", h.UpdateBook)
		bookGroup.GET("/", h.GetAllBook)
		bookGroup.DELETE("/:id", h.DeleteBooks)
	}

	publisherGroup := router.Group("/publishers")
	{
		publisherGroup.POST("/", h.CreatePublisher)
		publisherGroup.GET("/:id", h.GetByIdPublisher)
		publisherGroup.PUT("/", h.UpdatePublisher)
		publisherGroup.GET("/", h.GetAllPublishers)
		publisherGroup.DELETE("/:id", h.DeletePublisher)
	}

	translatorGroup := router.Group("/translators")
	{
		translatorGroup.POST("/", h.CreateTranslator)
		translatorGroup.GET("/:id", h.GetByIdTranslator)
		translatorGroup.PUT("/", h.UpdateTranslator)
		translatorGroup.GET("/", h.GetAllTranslators)
		translatorGroup.DELETE("/:id", h.DeleteTranslator)
	}

	return router
}
