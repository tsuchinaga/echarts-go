package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func barSimple(w http.ResponseWriter, _ *http.Request) {
	err := chart.New("Bar Simple", 800, 600).
		SetXAxis("x", chart.AxisTypeCategory, []interface{}{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		SetYAxis("y", chart.AxisTypeValue, nil).
		SetSeries("series", chart.SeriesTypeBar, []interface{}{120, 200, 150, 80, 70, 110, 130}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
