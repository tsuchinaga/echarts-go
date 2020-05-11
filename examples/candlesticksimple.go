package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func candlestickSimple(w http.ResponseWriter, _ *http.Request) {
	err := chart.New("Basic Candlestick", 800, 600).
		SetXAxis("x", chart.AxisTypeCategory, []interface{}{"2017-10-24", "2017-10-25", "2017-10-26", "2017-10-27"}).
		SetYAxis("y", chart.AxisTypeValue, nil).
		SetSeries("series", chart.SeriesTypeCandlestick, []interface{}{
			[]int{20, 30, 10, 35},
			[]int{40, 35, 30, 55},
			[]int{33, 38, 33, 40},
			[]int{40, 40, 32, 42},
		}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
