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

type GridOpt struct {
	chartOpt
	Left         string // 左
	Top          string // 上
	Right        string // 右
	Bottom       string // 下
	ContainLabel Bool   // グリッド領域に軸目盛を含めるか
}

func (o *GridOpt) getChartOptType() string {
	return "grid"
}
func (o *GridOpt) setOpt(option *option) {
	option.Grid = &grid{
		Left:         o.Left,
		Top:          o.Top,
		Right:        o.Right,
		Bottom:       o.Bottom,
		ContainLabel: o.ContainLabel,
	}
}
