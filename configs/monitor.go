package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Monitor() monitor.Config {
	return monitor.Config{
		Title:      "Stock Server Monitor",
		Refresh:    time.Second * 5,
		APIOnly:    false,
		Next:       nil,
		CustomHead: "",
		FontURL:    "https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap",
		ChartJsURL: "https://cdn.jsdelivr.net/npm/chart.js@2.9/dist/Chart.bundle.min.js",
	}
}
