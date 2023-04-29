package query_service

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Host string
	Port string 
	Protocol string
	RequestPath string
	Method string
}

func InitService() *Service{
	return &Service{

	}
}


func (s *Service) ServeService() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Request.URL.Scheme = s.Protocol
		c.Request.URL.Host = s.Host
		c.Request.URL.Path = s.RequestPath
		director := func(req *http.Request){
			req = c.Request
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterServiceEndpoint(router *gin.Engine, service *Service, path string) {
	switch service.Method{
	case "POST":
		router.POST(path, service.ServeService())
	case "GET":
		router.GET(path, service.ServeService())
	case "PUT":
		router.PUT(path, service.ServeService())
	case "DELETE":
		router.DELETE(path, service.ServeService())
	case "ANY":
		router.Any(path, service.ServeService())
	}
}


func RegisterServiceEndpointWithMiddleware(router *gin.RouterGroup, service *Service, path string) {
	switch service.Method{
	case "POST":
		router.POST(path, service.ServeService())
	case "GET":
		router.GET(path, service.ServeService())
	case "PUT":
		router.PUT(path, service.ServeService())
	case "DELETE":
		router.DELETE(path, service.ServeService())
	case "ANY":
		router.Any(path, service.ServeService())
	}
}