package query_service

import(
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	service := InitService()

	query_group := r.Group("")

	query_group.GET("/auctions/:id", service.ServeService("/auctions/:id"))
	query_group.GET("/products/:id", service.ServeService("/products/:id")) 
	query_group.GET("/search", service.ServeService("/search"))
	query_group.GET("/suggestions", service.ServeService("/suggestions"))
	query_group.GET("/health", service.ServeService("/health"))
}


