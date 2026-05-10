package httptransport

import (
	"net/http"

	appuser "goTrellisDemo/internal/app/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
)

type UserHandler struct {
	rpc rpcHandler
}

func NewUserHandler(client client.Client, serviceName string, addresses ...string) *UserHandler {
	return &UserHandler{
		rpc: newRPCHandler(client, serviceName, addresses...),
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var loginReq appuser.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login request"})
		return
	}

	rsp := new(appuser.LoginResponse)
	h.rpc.call(c, "User.Login", &loginReq, rsp)
}
