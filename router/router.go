package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavimoreno548/BookStore/app/controllers"
	"github.com/xavimoreno548/BookStore/app/middlewares"
)

func InitRouter(port string){

	var err error
	router := gin.Default()

	/////////////////// BEGIN PUBLIC ROUTES ///////////////////
	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})

	public := router.Group("/api")

	// Register Routes
	public.POST("/register", controllers.Register)

	// Login routes
	public.POST("/login", controllers.Login)
	/////////////////// END PUBLIC ROUTES ///////////////////


	/////////////////// BEGIN PRIVATE ROUTES ///////////////////
	protected := router.Group("/api/auth")
	protected.Use(middlewares.Authorize())

	protected.GET("/user", controllers.CurrentUser)

	// Book routes
	protected.GET("/books",controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/book/:id", controllers.FindBook)
	router.PATCH("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	// Editorials routes
	router.GET("/editorials", controllers.GetEditorials)
	router.POST("/editorial", controllers.CreateEditorial)

	// Author routes
	router.GET("/authors", controllers.GetAuthors)
	router.POST("/author", controllers.CreateAuthor)
	/////////////////// END PRIVATE ROUTES ///////////////////

	err = router.Run("0.0.0.0:"+port)

	if err != nil {
		fmt.Println(err)
	}
}
