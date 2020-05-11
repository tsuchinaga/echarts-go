package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func getStarted(w http.ResponseWriter, _ *http.Request) {
	c := chart.New("ECharts", 800, 600).
		SetOption(
			&chart.TitleOpt{Text: "ECharts entry example"},
			&chart.LegendOpt{Data: []string{"shirt", "cardign", "chiffon shirt", "pants", "heels", "socks"}},
			&chart.TooltipBaseOpt{}).
		SetXAxis("x", chart.AxisTypeEmpty, []interface{}{"shirt", "cardign", "chiffon shirt", "pants", "heels", "socks"}).
		SetSeries("sales", chart.SeriesTypeBar, []interface{}{5, 20, 36, 10, 10, 20}, &chart.SeriesBaseOpt{Name: "Sales"})
	if err := c.Render(w); err != nil {
		log.Println(err)
	}
}
