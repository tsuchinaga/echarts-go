package chart

type seriesOpt interface {
	getSeriesOptType() string
	setOpt(series *series)
}

// SeriesBaseOpt - シリーズの基本的な設定
type SeriesBaseOpt struct {
	seriesOpt
	Name           string     // 名前
	XAxisIndex     int        // X軸のインデックス
	YAxisIndex     int        // Y軸のインデックス
	Smooth         Bool       // スムーズにするか
	HoverAnimation Bool       // hover時のアニメーション
	Symbol         SymbolType // シンボルの形
	SymbolSize     int        // シンボルのサイズ
	ShowSymbol     Bool       // シンボルの表示非表示
	Stack          string     // 積み上げ
	Large          Bool       //
}

func (o *SeriesBaseOpt) getSeriesOptType() string {
	return "base"
}
func (o *SeriesBaseOpt) setOpt(series *series) {
	series.Name = o.Name
	series.XAxisIndex = o.XAxisIndex
	series.YAxisIndex = o.YAxisIndex
	series.Smooth = o.Smooth
	series.HoverAnimation = o.HoverAnimation
	series.Symbol = o.Symbol
	series.SymbolSize = o.SymbolSize
	series.ShowSymbol = o.ShowSymbol
	series.Stack = o.Stack
	series.Large = o.Large
}

// ItemStyleOpt - データのスタイル
type ItemStyleOpt struct {
	seriesOpt
	Color        string // 色
	Color0       string // 色0
	BorderColor  string // 境界線の色
	BorderColor0 string // 境界線の色0
}

func (o *ItemStyleOpt) getSeriesOptType() string {
	return "itemStyle"
}
func (o *ItemStyleOpt) setOpt(series *series) {
	series.ItemStyle = &itemStyle{Color: o.Color, Color0: o.Color0, BorderColor: o.BorderColor, BorderColor0: o.BorderColor0}
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

// EncodeOpt - データをデータセットから取り出す設定
type EncodeOpt struct {
	seriesOpt
	X []int
	Y []int
}

func (o *EncodeOpt) getSeriesOptType() string {
	return "encode"
}
func (o *EncodeOpt) setOpt(series *series) {
	series.Encode = &encode{X: o.X, Y: o.Y}
}
