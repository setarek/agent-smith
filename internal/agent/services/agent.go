package services

import (
	"math"
)

var ActiveAgents map[int64]Agent

func init() {
	ActiveAgents = make(map[int64]Agent)
}

type Agent struct {
	ID              int64
	Coordinate      Coordinate
	Done            bool
	CurrentDistance float64
}

func (a *Agent) IncID() int64 {
	a.ID = a.ID + 1
	return a.ID
}

func (a *Agent) Distance(c Coordinate) float64 {
	return math.Sqrt(math.Pow((c.X-a.Coordinate.X), 2) - math.Pow((c.Y-a.Coordinate.Y), 2))
}

func (a *Agent) GetActiveAgents() map[int64]Agent {
	return ActiveAgents
}

func (a *Agent) SetAgent() {
	a.ID = a.IncID()
	ActiveAgents[a.ID] = *a
}

func (a *Agent) TrigerAgent(x float64, y float64) {
	a.Coordinate.X = x
	a.Coordinate.Y = y
	a.ID = a.IncID()

	activeAgents := a.GetActiveAgents()

	if len(activeAgents) == 0 {
		distance := a.Coordinate.Calculate(0, 0)
		a.Coordinate.Walk(a.ID, distance, 0)
		return
	}

	var temp float64

	selectedAgent := Agent{}
	for _, agent := range activeAgents {
		distance := agent.Coordinate.Calculate(agent.Coordinate.X, agent.Coordinate.Y)
		if temp == 0 {
			temp = distance
		} else if distance < temp {
			temp = distance
			selectedAgent = agent
		}
	}

	a.Coordinate.Walk(a.ID, temp, selectedAgent.CurrentDistance)
	if selectedAgent.CurrentDistance == 0 {
		a.CurrentDistance = temp
	} else {
		a.CurrentDistance = selectedAgent.CurrentDistance + temp
	}
	a.SetAgent()

	delete(activeAgents, selectedAgent.ID)
}
