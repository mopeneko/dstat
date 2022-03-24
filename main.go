package main

import (
	"net/http"

	"github.com/akhenakh/statgo"
	"github.com/labstack/echo/v4"
)

type StatsResponse struct {
	TXBytes   int
	RXBytes   int
	TXPackets int
	RXPackets int
}

func main() {
	e := echo.New()
	e.GET("/stats", func(c echo.Context) error {
		s := statgo.NewStat()
		netIOStats := s.NetIOStats()[0]
		res := &StatsResponse{
			TXBytes:   netIOStats.TX,
			RXBytes:   netIOStats.RX,
			TXPackets: netIOStats.OPackets,
			RXPackets: netIOStats.IPackets,
		}
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
