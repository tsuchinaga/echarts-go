package chart

type chartOpt interface {
	getChartOptType() string
	setOpt(option *option)
}

// TitleOpt - タイトルに関する設定
type TitleOpt struct {
	chartOpt
	Id      string // ID
	Text    string // タイトル
	Subtext string // サブタイトル
	Left    string // 左からの位置
}

func (o *TitleOpt) getChartOptType() string {
	return "title"
}
func (o *TitleOpt) setOpt(option *option) {
	option.Title = &title{Id: o.Id, Text: o.Text, Subtext: o.Subtext, Left: o.Left}
}

// LegendOpt - 凡例に関する設定
type LegendOpt struct {
	chartOpt
	Type string   // 種別
	Id   string   // ID
	Data []string // 凡例
}

func (o *LegendOpt) getChartOptType() string {
	return "legend"
}
func (o *LegendOpt) setOpt(option *option) {
	option.Legend = &legend{Type: o.Type, Id: o.Id, Data: o.Data}
}

// TooltipBaseOpt - ツールチップに関する設定
type TooltipBaseOpt struct {
	chartOpt
	Trigger   TooltipTrigger // トリガー
	Formatter string         // フォーマッター
}

func (o *TooltipBaseOpt) getChartOptType() string {
	return "tooltipBase"
}
func (o *TooltipBaseOpt) setOpt(option *option) {
	option.Tooltip = &tooltip{Trigger: o.Trigger, Formatter: o.Formatter}
}

// TooltipLinePointerOpt - ツールチップのラインポインタに関する設定
type TooltipLinePointerOpt struct {
	chartOpt
}

func (o *TooltipLinePointerOpt) getChartOptType() string {
	return "tooltipLinePointer"
}
func (o *TooltipLinePointerOpt) setOpt(option *option) {
	if option.Tooltip == nil {
		option.Tooltip = new(tooltip)
	}
	option.Tooltip.AxisPointer = &axisPointer{
		Type: AxisPointerTypeLine,
	}
}

// TooltipCrossPointerOpt - ツールチップのクロスポインタに関する設定
type TooltipCrossPointerOpt struct {
	chartOpt
	Color string // 色
}

func (o *TooltipCrossPointerOpt) getChartOptType() string {
	return "tooltipCrossPointer"
}
func (o *TooltipCrossPointerOpt) setOpt(option *option) {
	if option.Tooltip == nil {
		option.Tooltip = new(tooltip)
	}
	option.Tooltip.AxisPointer = &axisPointer{
		Type:       AxisPointerTypeCross,
		CrossStyle: &axisStyle{Color: o.Color},
	}
}

// GridOpt - グリッドに関するオプション
type GridOpt struct {
	chartOpt
	Id           string // ID
	Left         string // 左
	Top          string // 上
	Right        string // 右
	Bottom       string // 下
	Height       string // 高さ
	ContainLabel Bool   // グリッド領域に軸目盛を含めるか
}

func (o *GridOpt) getChartOptType() string {
	return "grid"
}
func (o *GridOpt) setOpt(option *option) {
	if option.Grid == nil {
		option.Grid = make([]*grid, 0)
	}
	g := &grid{
		Id:           o.Id,
		Left:         o.Left,
		Top:          o.Top,
		Right:        o.Right,
		Bottom:       o.Bottom,
		Height:       o.Height,
		ContainLabel: o.ContainLabel,
	}

	for i, grid := range option.Grid {
		if grid.Id == g.Id {
			option.Grid[i] = g
			return
		}
	}

	option.Grid = append(option.Grid, g)
}

// DatasetOpt - データセットに関する設定
type DatasetOpt struct {
	chartOpt
	Source []interface{} // データセット
}

func (o *DatasetOpt) getChartOptType() string {
	return "dataset"
}
func (o *DatasetOpt) setOpt(option *option) {
	option.Dataset = &dataset{Source: o.Source}
}

// DataZoomInsideOpt - チャート内のデータズームに関する設定
type DataZoomInsideOpt struct {
	chartOpt
	XAxisIndex []int  //
	Start      string //
	End        string //
}

func (o *DataZoomInsideOpt) getChartOptType() string {
	return "dataZoomInside"
}
func (o *DataZoomInsideOpt) setOpt(option *option) {
	if option.DataZoom == nil {
		option.DataZoom = make([]*dataZoom, 0)
	}

	d := &dataZoom{Type: "inside", XAxisIndex: o.XAxisIndex, Start: o.Start, End: o.End}
	for i, dataZoom := range option.DataZoom {
		if dataZoom.Type == d.Type {
			option.DataZoom[i] = d
			return
		}
	}

	option.DataZoom = append(option.DataZoom, d)
}

// DataZoomSliderOpt - スライダーのデータズームに関する設定
type DataZoomSliderOpt struct {
	chartOpt
	Show       Bool   // 表示するか
	XAxisIndex []int  //
	Bottom     string //
	Start      string //
	End        string //
	HandleIcon string //
	HandleSize string //
}

func (o *DataZoomSliderOpt) getChartOptType() string {
	return "dataZoomSlider"
}
func (o *DataZoomSliderOpt) setOpt(option *option) {
	if option.DataZoom == nil {
		option.DataZoom = make([]*dataZoom, 0)
	}

	d := &dataZoom{Type: "slider", Show: o.Show, XAxisIndex: o.XAxisIndex, Bottom: o.Bottom, Start: o.Start, End: o.End, HandleIcon: o.HandleIcon, HandleSize: o.HandleSize}
	for i, dataZoom := range option.DataZoom {
		if dataZoom.Type == d.Type {
			option.DataZoom[i] = d
			return
		}
	}

	option.DataZoom = append(option.DataZoom, d)
}
