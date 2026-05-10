package httptransport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

type rpcHandler struct {
	client      client.Client
	serviceName string
	addresses   []string
}

func newRPCHandler(client client.Client, serviceName string, addresses ...string) rpcHandler {
	return rpcHandler{
		client:      client,
		serviceName: serviceName,
		addresses:   addresses,
	}
}

func (h *rpcHandler) call(c *gin.Context, method string, request interface{}, response interface{}) {
	req := h.client.NewRequest(
		h.serviceName,
		method,
		request,
		client.WithContentType("application/json"),
	)

	callOptions := make([]client.CallOption, 0, 1)
	if len(h.addresses) > 0 {
		callOptions = append(callOptions, client.WithAddress(h.addresses...))
	}

	if err := h.client.Call(c.Request.Context(), req, response, callOptions...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
