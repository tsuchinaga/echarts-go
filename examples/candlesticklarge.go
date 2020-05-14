package main

import (
	"fmt"
	"gitlab.com/tsuchinaga/echarts-go/chart"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

func candlestickLarge(w http.ResponseWriter, _ *http.Request) {
	count := 200000
	data := generateOHLC(count)
	err := chart.New("Candle Stick Large", 800, 600).
		SetOption(
			&chart.DatasetOpt{Source: data},
			&chart.TitleOpt{Text: fmt.Sprintf("Data Amount: %d", count)},
			&chart.TooltipBaseOpt{Trigger: chart.TooltipTriggerAxis},
			&chart.TooltipLinePointerOpt{},
			&chart.GridOpt{Id: "grid1", Left: "10%", Right: "10%", Bottom: "200"},
			&chart.GridOpt{Id: "grid2", Left: "10%", Right: "10%", Height: "80", Bottom: "80"},
			&chart.DataZoomInsideOpt{XAxisIndex: []int{0, 1}, Start: 10, End: 100},
			&chart.DataZoomSliderOpt{Show: chart.True, XAxisIndex: []int{0, 1}, Bottom: "10", Start: 10, End: 100, HandleSize: "105%", HandleIcon: "M10.7,11.9H9.3c-4.9,0.3-8.8,4.4-8.8,9.4c0,5,3.9,9.1,8.8,9.4h1.3c4.9-0.3,8.8-4.4,8.8-9.4C19.5,16.3,15.6,12.2,10.7,11.9z M13.3,24.4H6.7V23h6.6V24.4z M13.3,19.6H6.7v-1.4h6.6V19.6z"}).
		SetToolboxFeature(&chart.DataZoomFeatureOpt{YAxisIndex: chart.False}).
		SetXAxis("x1", chart.AxisTypeCategory, nil,
			&chart.AxisBaseOpt{Scale: chart.True, BoundaryGap: chart.False, SplitNumber: 20, Min: "dataMin", Max: "dataMax"}, &chart.AxisLineOpt{OnZero: chart.False}, &chart.SplitLineOpt{Show: chart.False}).
		SetXAxis("x2", chart.AxisTypeCategory, nil,
			&chart.AxisBaseOpt{Scale: chart.True, BoundaryGap: chart.False, GridIndex: 1, SplitNumber: 20, Min: "dataMin", Max: "dataMax"}, &chart.AxisLineOpt{OnZero: chart.False}, &chart.AxisLabelOpt{Show: chart.False}, &chart.AxisTickOpt{Show: chart.False}, &chart.SplitLineOpt{Show: chart.False}).
		SetYAxis("y1", chart.AxisTypeEmpty, nil, &chart.AxisBaseOpt{Scale: chart.True}, &chart.SplitAreaOpt{Show: chart.True}).
		SetYAxis("y2", chart.AxisTypeEmpty, nil, &chart.AxisBaseOpt{Scale: chart.True, GridIndex: 1, SplitNumber: 2},
			&chart.AxisLabelOpt{Show: chart.False}, &chart.AxisLineOpt{Show: chart.False}, &chart.AxisTickOpt{Show: chart.False}, &chart.SplitLineOpt{Show: chart.False}).
		SetSeries("s1", chart.SeriesTypeCandlestick, nil, &chart.EncodeOpt{X: []int{0}, Y: []int{1, 4, 3, 2}}, &chart.ItemStyleOpt{Color: "#ec0000", Color0: "#00da3c", BorderColor: "#8A0000", BorderColor0: "#008F28"}).
		SetSeries("s2", chart.SeriesTypeBar, nil, &chart.SeriesBaseOpt{XAxisIndex: 1, YAxisIndex: 1, Large: chart.True}, &chart.EncodeOpt{X: []int{0}, Y: []int{5}}, &chart.ItemStyleOpt{Color: "#7fbe9e"}).
		Render(w)
	if err != nil {
		log.Println(err)
	}
}

// generateOHLC - 指定された本数の4本値をランダムに作る
func generateOHLC(count int) []interface{} {
	rand.Seed(time.Now().UnixNano())
	data := make([]interface{}, count)

	datetime := time.Date(2011, 1, 1, 0, 0, 0, 0, time.Local)
	baseValue := rand.Float64() * 12000
	dayRange := 12.0

	for i := 0; i < count; i++ {
		datetime = datetime.Add(time.Minute)
		baseValue += rand.Float64()*20 - 10

		fourPrice := make([]float64, 4)
		for j := 0; j < 4; j++ {
			fourPrice[j] = (rand.Float64()-0.5)*dayRange + baseValue
		}
		sort.Slice(fourPrice, func(i, j int) bool {
			return fourPrice[i] < fourPrice[j]
		})

		openIdx := int(math.Round(rand.Float64() * 3))
		closeIdx := int(math.Round(rand.Float64() * 2))
		sign := 1
		if fourPrice[openIdx] < fourPrice[closeIdx] {
			sign = -1
		}
		data[i] = []interface{}{
			datetime.Format("2006-01-02\n15:04:05"), // datetime
			fourPrice[openIdx],                      // open
			fourPrice[3],                            // high
			fourPrice[0],                            // low
			fourPrice[closeIdx],                     // close
			int(fourPrice[3] * (1000 + rand.Float64()*500)), // volume
			sign, // 向き
		}
	}

	return data
}
