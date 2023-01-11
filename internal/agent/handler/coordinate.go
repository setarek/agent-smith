package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"agent-smith/internal/agent/services"
	rsErr "agent-smith/internal/error"
	"agent-smith/internal/queue"
)

type Handler struct {
	agent *services.Agent
}

type CoordinateRequest struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CoordinateResponse struct {
	Message string `json: "message`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) NewCoordinate(ctx gin.Context) {
	var request CoordinateRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, ErrorResponse{
			Message: rsErr.EmptyBodyRequest.Error(),
		})
	}

	agentChannel := queue.AgentChan{
		Agent: *h.agent,
		X:     request.X,
		Y:     request.Y,
	}

	queue.AgentChannel <- agentChannel

	ctx.JSON(http.StatusCreated, CoordinateResponse{
		Message: "new coordinate added",
	})
}
