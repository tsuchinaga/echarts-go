package chart

type seriesOpt interface {
	getSeriesOptType() string
	setOpt(series *series)
}

// SeriesBaseOpt - シリーズの基本的な設定
type SeriesBaseOpt struct {
	seriesOpt
	Name           string     // 名前
	YAxisIndex     int        // Y軸のインデックス
	Smooth         Bool       // スムーズにするか
	HoverAnimation Bool       // hover時のアニメーション
	Symbol         SymbolType // シンボルの形
	SymbolSize     int        // シンボルのサイズ
	ShowSymbol     Bool       // シンボルの表示非表示
	Stack          string     // 積み上げ
}

func (o *SeriesBaseOpt) getSeriesOptType() string {
	return "base"
}
func (o *SeriesBaseOpt) setOpt(series *series) {
	series.Name = o.Name
	series.YAxisIndex = o.YAxisIndex
	series.Smooth = o.Smooth
	series.HoverAnimation = o.HoverAnimation
	series.Symbol = o.Symbol
	series.SymbolSize = o.SymbolSize
	series.ShowSymbol = o.ShowSymbol
	series.Stack = o.Stack
}

// ItemStyleOpt - データのスタイル
type ItemStyleOpt struct {
	seriesOpt
	Color string // 色
}

func (o *ItemStyleOpt) getSeriesOptType() string {
	return "itemStyle"
}
func (o *ItemStyleOpt) setOpt(series *series) {
	series.ItemStyle = &itemStyle{Color: o.Color}
}

// LineStyleOpt - 線のスタイル
type LineStyleOpt struct {
	seriesOpt
	Opacity float64 // 不透明度
}

func (o *LineStyleOpt) getSeriesOptType() string {
	return "lineStyle"
}
func (o *LineStyleOpt) setOpt(series *series) {
	series.LineStyle = &lineStyle{Opacity: o.Opacity}
}

// AreaStyleOpt - エリアのスタイル
type AreaStyleOpt struct {
	seriesOpt
	Color string // 色
}

func (o *AreaStyleOpt) getSeriesOptType() string {
	return "areaStyle"
}
func (o *AreaStyleOpt) setOpt(series *series) {
	series.AreaStyle = &areaStyle{Color: o.Color}
}
