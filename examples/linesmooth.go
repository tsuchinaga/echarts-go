package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func lineSmooth(w http.ResponseWriter, _ *http.Request) {
	err := chart.New("Basic Line Chart", 800, 600).
		SetXAxis("x", chart.AxisTypeCategory, []interface{}{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		SetYAxis("y", chart.AxisTypeValue, nil).
		SetSeries("series", chart.SeriesTypeLine, []interface{}{820, 932, 901, 934, 1290, 1330, 1320}, &chart.SeriesBaseOpt{Smooth: chart.True}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
