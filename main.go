package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

const netIF = "eth0"

type StatsResponse struct {
	TXBytes   string
	RXBytes   string
	TXPackets string
	RXPackets string
}

func main() {
	e := echo.New()
	e.GET("/stats", func(c echo.Context) error {
		res := &StatsResponse{}

		txBytes, err := get("tx_bytes")
		if err != nil {
			c.Logger().Errorf("failed to get tx bytes: %+v", err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.TXBytes = txBytes

		rxBytes, err := get("rx_bytes")
		if err != nil {
			c.Logger().Errorf("failed to get rx bytes: %+v", err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.RXBytes = rxBytes

		txPackets, err := get("tx_packets")
		if err != nil {
			c.Logger().Errorf("failed to get tx packets: %+v", err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.TXPackets = txPackets

		rxPackets, err := get("rx_packets")
		if err != nil {
			c.Logger().Errorf("failed to get rx packets: %+v", err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.RXPackets = rxPackets

		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func get(target string) (string, error) {
	file, err := os.Open(fmt.Sprintf("/sys/class/net/%s/statistics/%s", netIF, target))
	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}
