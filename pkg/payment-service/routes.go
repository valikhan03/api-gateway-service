package payment_service

import(
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	service := InitService()

	payment := r.Group("payment")
	payment.POST("/pay", service.ServeService("/pay"))
}