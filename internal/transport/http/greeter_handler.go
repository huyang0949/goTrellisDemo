package httptransport

import (
	"net/http"

	appgreeter "goTrellisDemo/internal/app/greeter"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

type GreeterHandler struct {
	rpc rpcHandler
}

func NewGreeterHandler(client client.Client, serviceName string, addresses ...string) *GreeterHandler {
	return &GreeterHandler{
		rpc: newRPCHandler(client, serviceName, addresses...),
	}
}

func (h *GreeterHandler) Hello(c *gin.Context) {
	name := c.Query("name")
	if name == "" && c.Request.Body != nil {
		var req appgreeter.HelloRequest
		if err := c.ShouldBindJSON(&req); err == nil {
			name = req.Name
		}
	}

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	rsp := new(appgreeter.HelloResponse)
	h.rpc.call(c, "Greeter.Hello", &appgreeter.HelloRequest{Name: name}, rsp)
}
