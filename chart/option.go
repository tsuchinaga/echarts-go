package chart

// Option - チャートのオプション
type option struct {
	Title   *title    `json:"title,omitempty"`   // タイトル
	Legend  *legend   `json:"legend,omitempty"`  // 凡例
	Grid    *grid     `json:"grid,omitempty"`    // グリッド
	XAxis   []*axis   `json:"xAxis"`             // X軸
	YAxis   []*axis   `json:"yAxis"`             // Y軸
	Tooltip *tooltip  `json:"tooltip,omitempty"` // ツールチップ
	Toolbox *toolbox  `json:"toolbox,omitempty"` // ツールボックス
	Series  []*series `json:"series"`            // シリーズ
}

// title - タイトル
type title struct {
	Id      string `json:"id,omitempty"`      // ID
	Text    string `json:"text,omitempty"`    // タイトル
	Subtext string `json:"subtext,omitempty"` // サブタイトル
	Left    string `json:"left,omitempty"`    // 左からの位置
}

// legend - 凡例
type legend struct {
	Type string   `json:"type,omitempty"` // 種別
	Id   string   `json:"id,omitempty"`   // ID
	Data []string `json:"data,omitempty"` // 凡例
}

type grid struct {
	Left         string `json:"left,omitempty"`         // 左
	Top          string `json:"top,omitempty"`          // 上
	Right        string `json:"right,omitempty"`        // 右
	Bottom       string `json:"bottom,omitempty"`       // 下
	ContainLabel Bool   `json:"containLabel,omitempty"` // グリッド領域に軸目盛を含めるか
}

// tooltip - ツールチップ
type tooltip struct {
	Trigger     TooltipTrigger `json:"trigger,omitempty"`     // トリガー
	AxisPointer *axisPointer   `json:"axisPointer,omitempty"` // 軸ポインタ
	Formatter   string         `json:"formatter,omitempty"`   // フォーマッター
}

// toolbox - ツールボックス
type toolbox struct {
	Feature *toolboxFeature `json:"feature,omitempty"` // 機能
}

// toolboxFeature - ツールボックスの機能
type toolboxFeature struct {
	SaveAsImage *saveAsImageFeature `json:"saveAsImage,omitempty"`
	DataView    *dataViewFeature    `json:"dataView,omitempty"`
	MagicType   *magicTypeFeature   `json:"magicType,omitempty"`
	Restore     *restoreFeature     `json:"restore,omitempty"`
}

// saveAsImageFeature - 画像保存機能
type saveAsImageFeature struct {
	Show Bool `json:"show,omitempty"` // 表示するか
}

// dataViewFeature - データの閲覧機能
type dataViewFeature struct {
	Show     Bool `json:"show,omitempty"`     // 表示するか
	ReadOnly Bool `json:"readOnly,omitempty"` // 読み込み専用か
}

// magicTypeFeature - シリーズ種別の変更機能
type magicTypeFeature struct {
	Show Bool         `json:"show,omitempty"` // 表示するか
	Type []SeriesType `json:"type,omitempty"` // 選べるシリーズ種別
}

// restoreFeature - 元に戻す機能
type restoreFeature struct {
	Show Bool `json:"show,omitempty"` // 表示するか
}

// axis - 軸
type axis struct {
	Id          string        `json:"id,omitempty"`          // ID
	Type        AxisType      `json:"type,omitempty"`        // 種別
	Name        string        `json:"name,omitempty"`        // 名前
	Min         string        `json:"min,omitempty"`         // 最小
	Max         string        `json:"max,omitempty"`         // 最大
	Interval    int           `json:"interval,omitempty"`    // 間隔
	SplitNumber int           `json:"splitNumber,omitempty"` // 分割
	BoundaryGap Bool          `json:"boundaryGap,omitempty"` // 境界部分のギャップ
	Data        []interface{} `json:"data,omitempty"`        // データ
	AxisLabel   *axisLabel    `json:"axisLabel,omitempty"`   // 軸のラベル
	AxisPointer *axisPointer  `json:"axisPointer,omitempty"` // 軸ポインタ
	SplitLine   *splitLine    `json:"splitLine,omitempty"`   // 分割線
}

// axisLabel - 軸のラベル
type axisLabel struct {
	Formatter interface{} `json:"formatter,omitempty"` // フォーマッター
}

// axisPointer - 軸ポインタ
type axisPointer struct {
	Type       AxisPointerType `json:"type,omitempty"`       // 種別
	CrossStyle *axisStyle      `json:"crossStyle,omitempty"` // クロスポインタのスタイル
}

// axisStyle - 軸のスタイル
type axisStyle struct {
	Color string `json:"color,omitempty"` // 色
}

// splitLine - 分割線
type splitLine struct {
	Show Bool `json:"show"` // 表示するか
}

// series - シリーズ
type series struct {
	Type           SeriesType    `json:"type,omitempty"`           // 種別
	Id             string        `json:"id,omitempty"`             // ID
	Name           string        `json:"name,omitempty"`           // 名前
	YAxisIndex     int           `json:"yAxisIndex,omitempty"`     // Y軸のインデックス
	Smooth         Bool          `json:"smooth,omitempty"`         // スムーズにするか
	HoverAnimation Bool          `json:"hoverAnimation,omitempty"` // hover時のアニメーション
	Symbol         SymbolType    `json:"symbol,omitempty"`         // シンボルの形
	SymbolSize     int           `json:"symbolSize,omitempty"`     // シンボルのサイズ
	ShowSymbol     Bool          `json:"showSymbol,omitempty"`     // シンボルの表示非表示
	Stack          string        `json:"stack,omitempty"`          // 積み上げ
	Data           []interface{} `json:"data,omitempty"`           // データ
	ItemStyle      *itemStyle    `json:"itemStyle,omitempty"`      // データのスタイル
	LineStyle      *lineStyle    `json:"lineStyle,omitempty"`      // 線のスタイル
	AreaStyle      *areaStyle    `json:"areaStyle,omitempty"`      // エリアのスタイル
}

// itemStyle - データのスタイル
type itemStyle struct {
	Color string `json:"color,omitempty"` // 色
}

// lineStyle - 線のスタイル
type lineStyle struct {
	Opacity float64 `json:"opacity"` // 不透明度
}

// areaStyle - エリアのスタイル
type areaStyle struct {
	Color string `json:"color,omitempty"` // 色
}
