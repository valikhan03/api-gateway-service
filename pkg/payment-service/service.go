package payment_service

import (
	"net/http"
	"net/http/httputil"
	"api-gateway-service/pkg/models"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Host        string `json:"host"`
}

func InitService() *Service{
	return &Service{
		Host: models.ConfigGlobal.PaymentService.URL,
	}
}


func (s *Service) ServeService(req_path string) gin.HandlerFunc {
	return func(c *gin.Context){
		c.Request.URL.Scheme = "http"
		c.Request.URL.Host = s.Host
		c.Request.URL.Path = req_path
		director := func(req *http.Request){
			req = c.Request
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}




