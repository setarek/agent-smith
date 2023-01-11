package services

import (
	zlog "agent-smith/logger"

	"fmt"
	"math"
	"time"
)

type Coordinate struct {
	X           float64
	Y           float64
	Distanation float64
}

func (p *Coordinate) Calculate(x float64, y float64) float64 {
	return float64(math.Pow(math.Pow(p.X-x, 2)+math.Pow(p.Y-y, 2), 0.5))

}

func (p Coordinate) Walk(id int64, distance float64, oldDistance float64) {

	if distance-oldDistance < 1 {
		zlog.Logger.Log().Int64("ID", id).Str("Coordinate", fmt.Sprintf("X:%f, Y:%f", p.X, p.Y)).Float64("current step:", distance).Msg("Agent just arrived at disentation")
	} else {
		diff := int(math.Ceil(distance-oldDistance)) - 1
		for i := 0; i <= diff; i++ {

			if i == diff {
				zlog.Logger.Log().Int64("ID", id).Str("Coordinate", fmt.Sprintf("X:%f, Y:%f", p.X, p.Y)).Float64("current step:", distance).Msg("Agent just arrived at disentation")
			} else {
				zlog.Logger.Log().Int64("ID", id).Float64("current step:", distance+float64(i)).Msg("Keep on...")
			}
			time.Sleep(1 * time.Second)
		}
	}

}
