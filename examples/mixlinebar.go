package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func mixLineBar(w http.ResponseWriter, _ *http.Request) {
	err := chart.New("Mixed Line and Bar", 800, 600).
		SetOption(
			&chart.LegendOpt{Data: []string{"蒸発量", "降水量", "平均温度"}},
			&chart.TooltipBaseOpt{Trigger: chart.TooltipTriggerAxis},
			&chart.TooltipCrossPointerOpt{Color: "#999"}).
		SetXAxis("x", chart.AxisTypeCategory, []interface{}{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"}, &chart.AxisPointerOpt{Type: chart.AxisPointerTypeShadow}).
		SetYAxis("y1", chart.AxisTypeValue, nil, &chart.AxisBaseOpt{Name: "水量", Min: "0", Max: "250", Interval: 50}, &chart.AxisLabelOpt{Formatter: "{value} ml"}).
		SetYAxis("y2", chart.AxisTypeValue, nil, &chart.AxisBaseOpt{Name: "温度", Min: "0", Max: "25", Interval: 5}, &chart.AxisLabelOpt{Formatter: "{value} ℃"}).
		SetSeries("s1", chart.SeriesTypeBar, []interface{}{2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0, 6.4, 3.3}, &chart.SeriesBaseOpt{Name: "蒸発量"}).
		SetSeries("s2", chart.SeriesTypeBar, []interface{}{2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3}, &chart.SeriesBaseOpt{Name: "降水量"}).
		SetSeries("s3", chart.SeriesTypeLine, []interface{}{2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2}, &chart.SeriesBaseOpt{Name: "平均温度", YAxisIndex: 1}).
		SetToolboxFeature(
			&chart.DataViewFeatureOpt{Show: true, ReadOnly: false},
			&chart.MagicTypeFeatureOpt{Show: true, Type: []chart.SeriesType{chart.SeriesTypeLine, chart.SeriesTypeBar}},
			&chart.RestoreFeatureOpt{Show: true},
			&chart.SaveAsImageFeatureOpt{Show: true}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
