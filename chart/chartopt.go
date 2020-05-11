package chart

type chartOpt interface {
	getChartOptType() string
	setOpt(option *option)
}

// TitleOpt - タイトルに関する設定
type TitleOpt struct {
	chartOpt
	Id   string // ID
	Text string // テキスト
}

func (o *TitleOpt) getChartOptType() string {
	return "title"
}
func (o *TitleOpt) setOpt(option *option) {
	option.Title = &title{Id: o.Id, Text: o.Text}
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
	Trigger TooltipTrigger // トリガー
}

func (o *TooltipBaseOpt) getChartOptType() string {
	return "tooltipBase"
}
func (o *TooltipBaseOpt) setOpt(option *option) {
	option.Tooltip = &tooltip{Trigger: o.Trigger}
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
