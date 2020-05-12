package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"math"
	"net/http"
)

func confidenceBand(w http.ResponseWriter, _ *http.Request) {
	base := 0.0
	for _, d := range set1 {
		if d.Low < base {
			base = math.Floor(d.Low)
		}
	}
	base *= -1

	x := make([]interface{}, len(set1))
	s1 := make([]interface{}, len(set1))
	s2 := make([]interface{}, len(set1))
	s3 := make([]interface{}, len(set1))
	for i, d := range set1 {
		x[i] = d.Date
		s1[i] = d.Value + base
		s2[i] = d.Low + base
		s3[i] = d.High - d.Low
	}

	err := chart.New("Confidence Band", 1000, 600).
		SetOption(
			&chart.TitleOpt{Text: "Confidence Band", Subtext: "Example in MetricsGraphics.js", Left: "center"},
			&chart.TooltipBaseOpt{Trigger: chart.TooltipTriggerAxis},
			&chart.TooltipCrossPointerOpt{},
			&chart.GridOpt{Left: "3%", Right: "4%", Bottom: "3%", ContainLabel: chart.True}).
		SetXAxis("x", chart.AxisTypeCategory, x, &chart.AxisBaseOpt{BoundaryGap: chart.False}, &chart.SplitLineOpt{Show: chart.False}).
		SetYAxis("y", chart.AxisTypeEmpty, nil, &chart.AxisBaseOpt{SplitNumber: 3}, &chart.SplitLineOpt{Show: chart.False}).
		SetSeries("s1", chart.SeriesTypeLine, s1, &chart.SeriesBaseOpt{HoverAnimation: chart.False, SymbolSize: 6, ShowSymbol: chart.False}, &chart.ItemStyleOpt{Color: "#c23531"}).
		SetSeries("s2", chart.SeriesTypeLine, s2, &chart.SeriesBaseOpt{Name: "L", Symbol: chart.SymbolTypeNone, Stack: "stack1"}, &chart.LineStyleOpt{Opacity: 0}).
		SetSeries("s3", chart.SeriesTypeLine, s3, &chart.SeriesBaseOpt{Name: "H", Symbol: chart.SymbolTypeNone, Stack: "stack1"}, &chart.LineStyleOpt{Opacity: 0}, &chart.AreaStyleOpt{Color: "#ccc"}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
