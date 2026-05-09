package httptransport

import (
	"net/http"

	appgreeter "goTrellisDemo/internal/app/greeter"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

type GreeterHandler struct {
	client      client.Client
	serviceName string
	addresses   []string
}

func NewGreeterHandler(client client.Client, serviceName string, addresses ...string) *GreeterHandler {
	return &GreeterHandler{
		client:      client,
		serviceName: serviceName,
		addresses:   addresses,
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

	req := h.client.NewRequest(
		h.serviceName,
		"Greeter.Hello",
		&appgreeter.HelloRequest{Name: name},
		client.WithContentType("application/json"),
	)

	rsp := new(appgreeter.HelloResponse)
	callOptions := make([]client.CallOption, 0, 1)
	if len(h.addresses) > 0 {
		callOptions = append(callOptions, client.WithAddress(h.addresses...))
	}

	if err := h.client.Call(c.Request.Context(), req, rsp, callOptions...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rsp)
}
