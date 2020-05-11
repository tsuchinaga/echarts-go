package chart

type seriesOpt interface {
	getSeriesOptType() string
	setOpt(series *series)
}

// SeriesBaseOpt - シリーズの基本的な設定
type SeriesBaseOpt struct {
	seriesOpt
	Name       string // 名前
	YAxisIndex int    // Y軸のインデックス
	Smooth     bool   // スムーズにするか
}

func (o *SeriesBaseOpt) getSeriesOptType() string {
	return "base"
}
func (o *SeriesBaseOpt) setOpt(series *series) {
	series.Name = o.Name
	series.YAxisIndex = o.YAxisIndex
	series.Smooth = o.Smooth
}
