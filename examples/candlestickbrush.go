package main

import (
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"net/http"
)

func candlestickBrush(w http.ResponseWriter, _ *http.Request) {
	getSide := func(o, c float64) float64 {
		if o > c {
			return 1
		} else {
			return -1
		}
	}

	x := make([]interface{}, len(set2))
	price := make([]interface{}, len(set2))
	volume := make([]interface{}, len(set2))
	maPeriods := []int{5, 10, 20, 30}
	maSum := make([]float64, len(maPeriods))
	ma := make([][]interface{}, len(maPeriods))
	for i, d := range set2 {
		x[i] = d.Date
		price[i] = []float64{d.Open, d.Close, d.High, d.Low}                  // 始値、終値、高値、安値
		volume[i] = []float64{float64(i), d.Volume, getSide(d.Open, d.Close)} // インデックス、値、向き

		for j, period := range maPeriods {
			maSum[j] += d.Close
			if i < period-1 {
				continue
			}
			if ma[j] == nil {
				ma[j] = make([]interface{}, len(set2))
			}
			if i >= period {
				maSum[j] -= set2[i-period].Close
			}
			ma[j][i] = maSum[j] / float64(period)
		}
	}

	err := chart.New("Candlestick Brush", 800, 600).
		SetOption(
			&chart.BaseOpt{BackgroundColor: "#fff", Animation: chart.False},
			&chart.LegendOpt{Bottom: "10", Left: "center", Data: []string{"Dow-Jones index", "MA5", "MA10", "MA20", "MA30"}},
			&chart.TooltipBaseOpt{
				Trigger:         chart.TooltipTriggerAxis,
				BackgroundColor: "rgba(245, 245, 245, 0.8)",
				BorderWidth:     "1",
				BorderColor:     "#cccccc",
				Padding:         "10",
				TextColor:       "#000",
			},
			&chart.GlobalAxisPointer{LinkXAxisIndex: []int{0, 1}},
			&chart.GridOpt{Id: "grid1", Left: "10%", Right: "8%", Height: "50%"},
			&chart.GridOpt{Id: "grid2", Left: "10%", Right: "8%", Top: "63%", Height: "16%"},
			&chart.DataZoomInsideOpt{XAxisIndex: []int{0, 1}, Start: 98, End: 100},
			&chart.DataZoomSliderOpt{Show: chart.True, XAxisIndex: []int{0, 1}, Top: "85%", Start: 98, End: 100}).
		SetToolboxFeature(
			&chart.DataZoomFeatureOpt{YAxisIndex: chart.False},
			&chart.BrushFeatureOpt{Type: []chart.BrushType{chart.BrushTypeLineX, chart.BrushTypeClear}}).
		SetXAxis("x1", chart.AxisTypeCategory, x,
			&chart.AxisBaseOpt{Scale: chart.True, BoundaryGap: chart.False, SplitNumber: 20, Min: "dataMin", Max: "dataMax"},
			&chart.AxisLineOpt{OnZero: chart.False}, &chart.SplitLineOpt{Show: chart.False}, &chart.AxisPointerOpt{Z: 100}).
		SetXAxis("x2", chart.AxisTypeCategory, x,
			&chart.AxisBaseOpt{GridIndex: 1, Scale: chart.True, BoundaryGap: chart.False, SplitNumber: 20, Min: "dataMin", Max: "dataMax"},
			&chart.AxisLineOpt{OnZero: chart.False}, &chart.AxisTickOpt{Show: chart.False}, &chart.SplitLineOpt{Show: chart.False}, &chart.AxisLabelOpt{Show: chart.False}).
		SetYAxis("y1", chart.AxisTypeEmpty, nil, &chart.AxisBaseOpt{Scale: chart.True}, &chart.SplitAreaOpt{Show: chart.True}).
		SetYAxis("y2", chart.AxisTypeEmpty, nil, &chart.AxisBaseOpt{Scale: chart.True, GridIndex: 1, SplitNumber: 2},
			&chart.AxisLabelOpt{Show: chart.False}, &chart.AxisLineOpt{Show: chart.False}, &chart.AxisTickOpt{Show: chart.False}, &chart.SplitLineOpt{Show: chart.False}).
		SetSeries("price", chart.SeriesTypeCandlestick, price, &chart.SeriesBaseOpt{Name: "Dow-Jones index"}, &chart.ItemStyleOpt{Color: "#00da3c", Color0: "#ec0000", BorderColor: "#00da3c", BorderColor0: "#ec0000"}).
		SetSeries("ma5", chart.SeriesTypeLine, ma[0], &chart.SeriesBaseOpt{Name: "MA5", Smooth: chart.True}, &chart.LineStyleOpt{Opacity: 0.5}).
		SetSeries("ma10", chart.SeriesTypeLine, ma[1], &chart.SeriesBaseOpt{Name: "MA10", Smooth: chart.True}, &chart.LineStyleOpt{Opacity: 0.5}).
		SetSeries("ma20", chart.SeriesTypeLine, ma[2], &chart.SeriesBaseOpt{Name: "MA20", Smooth: chart.True}, &chart.LineStyleOpt{Opacity: 0.5}).
		SetSeries("ma30", chart.SeriesTypeLine, ma[3], &chart.SeriesBaseOpt{Name: "MA30", Smooth: chart.True}, &chart.LineStyleOpt{Opacity: 0.5}).
		SetSeries("volume", chart.SeriesTypeBar, volume, &chart.SeriesBaseOpt{Name: "Volume", XAxisIndex: 1, YAxisIndex: 1}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}
