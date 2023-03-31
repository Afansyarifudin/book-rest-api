package route

import (
	"book-rest-api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/book") //prefix
	{
		api.POST("", server.CreateBook)
		api.GET("", server.GetAllBook)
		api.GET("/:id", server.GetBookById)
		api.PUT("/:id", server.UpdateBook)
		api.DELETE("/:id", server.DeleteBook)
	}
}