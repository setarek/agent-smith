package queue

import (
	"agent-smith/internal/agent/services"
)

var AgentChannel chan AgentChan

type AgentChan struct {
	Agent services.Agent
	X     float64
	Y     float64
}

func init() {
	AgentChannel = make(chan AgentChan, 3)
}

func (a *AgentChan) AddNewCoordinate() {
	a.Agent.TrigerAgent(a.X, a.Y)
}

func Walker() {
	for w := range AgentChannel {
		w.AddNewCoordinate()
	}
}
