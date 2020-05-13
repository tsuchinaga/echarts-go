package chart

type axisOpt interface {
	getAxisOptType() string
	setOpt(axis *axis)
}

// AxisBaseOpt - 軸の基本的な設定
type AxisBaseOpt struct {
	axisOpt
	Name        string // 名前
	Min         string // 最小
	Max         string // 最大
	Interval    int    // 間隔
	GridIndex   int    // グリッド設定のインデックス
	SplitNumber int    // 分割
	BoundaryGap Bool   // 境界部分のギャップ
	Scale       Bool   // スケールするか
}

func (o *AxisBaseOpt) getAxisOptType() string {
	return "base"
}

func (o *AxisBaseOpt) setOpt(axis *axis) {
	axis.Name = o.Name
	axis.Min = o.Min
	axis.Max = o.Max
	axis.Interval = o.Interval
	axis.GridIndex = o.GridIndex
	axis.SplitNumber = o.SplitNumber
	axis.BoundaryGap = o.BoundaryGap
	axis.Scale = o.Scale
}

// AxisPointerOpt - 軸ポインタの設定
type AxisPointerOpt struct {
	axisOpt
	Type AxisPointerType // 軸ポインタの種別
}

func (o *AxisPointerOpt) getAxisOptType() string {
	return "axisPointer"
}

func (o *AxisPointerOpt) setOpt(axis *axis) {
	axis.AxisPointer = &axisPointer{Type: o.Type}
}

// AxisLabelOpt - 軸ラベルの設定
type AxisLabelOpt struct {
	axisOpt
	Show      Bool   // 表示するか
	Formatter string // フォーマッター
}

func (o *AxisLabelOpt) getAxisOptType() string {
	return "axisLabel"
}

func (o *AxisLabelOpt) setOpt(axis *axis) {
	axis.AxisLabel = &axisLabel{Show: o.Show, Formatter: o.Formatter}
}

// SplitLineOpt - 分割線の設定
type SplitLineOpt struct {
	axisOpt
	Show Bool // 表示するか
}

func (o *SplitLineOpt) getAxisOptType() string {
	return "splitLine"
}

func (o *SplitLineOpt) setOpt(axis *axis) {
	axis.SplitLine = &splitLine{Show: o.Show}
}

// AxisLineOpt - 軸線に関する設定
type AxisLineOpt struct {
	axisOpt
	Show   Bool // 表示するか
	OnZero Bool // 原点表示について
}

func (o *AxisLineOpt) getAxisOptType() string {
	return "axisLine"
}

func (o *AxisLineOpt) setOpt(axis *axis) {
	axis.AxisLine = &axisLine{Show: o.Show, OnZero: o.OnZero}
}

// AxisTickOpt - 軸目盛に関する設定
type AxisTickOpt struct {
	axisOpt
	Show Bool // 表示するか
}

func (o *AxisTickOpt) getAxisOptType() string {
	return "axisTick"
}

func (o *AxisTickOpt) setOpt(axis *axis) {
	axis.AxisTick = &axisTick{Show: o.Show}
}

// SplitAreaOpt - 軸のエリア区切りに関する設定
type SplitAreaOpt struct {
	axisOpt
	Show Bool // 表示するか
}

func (o *SplitAreaOpt) getAxisOptType() string {
	return "axisArea"
}

func (o *SplitAreaOpt) setOpt(axis *axis) {
	axis.SplitAreaOpt = &splitAreaOpt{Show: o.Show}
}
