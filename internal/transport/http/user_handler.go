package httptransport

import (
	"net/http"

	appuser "goTrellisDemo/internal/app/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

type UserHandler struct {
	client      client.Client
	serviceName string
	addresses   []string
}

func NewUserHandler(client client.Client, serviceName string, addresses ...string) *UserHandler {
	return &UserHandler{
		client:      client,
		serviceName: serviceName,
		addresses:   addresses,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginReq appuser.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login request"})
		return
	}

	req := h.client.NewRequest(
		h.serviceName,
		"User.Login",
		&loginReq,
		client.WithContentType("application/json"),
	)

	rsp := new(appuser.LoginResponse)
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
