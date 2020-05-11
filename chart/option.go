package chart

// Option - チャートのオプション
type option struct {
	Title   *title    `json:"title,omitempty"`   // タイトル
	Legend  *legend   `json:"legend,omitempty"`  // 凡例
	XAxis   []*axis   `json:"xAxis"`             // X軸
	YAxis   []*axis   `json:"yAxis"`             // Y軸
	Tooltip *tooltip  `json:"tooltip,omitempty"` // ツールチップ
	Toolbox *toolbox  `json:"toolbox,omitempty"` // ツールボックス
	Series  []*series `json:"series"`            // シリーズ
}

// title - タイトル
type title struct {
	Id   string `json:"id,omitempty"`   // ID
	Text string `json:"text,omitempty"` // タイトル
}

// legend - 凡例
type legend struct {
	Type string   `json:"type,omitempty"` // 種別
	Id   string   `json:"id,omitempty"`   // ID
	Data []string `json:"data,omitempty"` // 凡例
}

// tooltip - ツールチップ
type tooltip struct {
	Trigger     TooltipTrigger `json:"trigger,omitempty"`     // トリガー
	AxisPointer *axisPointer   `json:"axisPointer,omitempty"` // 軸ポインタ
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
	Show  bool   `json:"show,omitempty"`  // 表示するか
	Title string `json:"title,omitempty"` // タイトル
}

// dataViewFeature - データの閲覧機能
type dataViewFeature struct {
	Show     bool      `json:"show,omitempty"`     // 表示するか
	Title    string    `json:"title,omitempty"`    // タイトル
	Lang     [3]string `json:"lang,omitempty"`     // ボタン
	ReadOnly bool      `json:"readOnly,omitempty"` // 読み込み専用か
}

// magicTypeFeature - シリーズ種別の変更機能
type magicTypeFeature struct {
	Show  bool         `json:"show,omitempty"`  // 表示するか
	Title string       `json:"title,omitempty"` // タイトル
	Type  []SeriesType `json:"type,omitempty"`  // 選べるシリーズ種別
}

// restoreFeature - 元に戻す機能
type restoreFeature struct {
	Show  bool   `json:"show,omitempty"`  // 表示するか
	Title string `json:"title,omitempty"` // タイトル
}

// axis - 軸
type axis struct {
	Id          string        `json:"id,omitempty"`          // ID
	Type        AxisType      `json:"type,omitempty"`        // 種別
	Name        string        `json:"name,omitempty"`        // 名前
	Min         string        `json:"min,omitempty"`         // 最小
	Max         string        `json:"max,omitempty"`         // 最大
	Interval    int           `json:"interval,omitempty"`    // 間隔
	Data        []interface{} `json:"data,omitempty"`        // データ
	AxisLabel   *axisLabel    `json:"axisLabel,omitempty"`   // 軸のラベル
	AxisPointer *axisPointer  `json:"axisPointer,omitempty"` // 軸ポインタ
}

// axisLabel - 軸のラベル
type axisLabel struct {
	Formatter string `json:"formatter,omitempty"` // フォーマッター
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

// series - シリーズ
type series struct {
	Type       SeriesType    `json:"type,omitempty"`       // 種別
	Id         string        `json:"id,omitempty"`         // ID
	Name       string        `json:"name,omitempty"`       // 名前
	YAxisIndex int           `json:"yAxisIndex,omitempty"` // Y軸のインデックス
	Smooth     bool          `json:"smooth,omitempty"`     // スムーズにするか
	Data       []interface{} `json:"data,omitempty"`       // データ
}
