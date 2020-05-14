package chart

// Option - チャートのオプション
type option struct {
	Title           *title       `json:"title,omitempty"`           // タイトル
	Legend          *legend      `json:"legend,omitempty"`          // 凡例
	Grid            []*grid      `json:"grid,omitempty"`            // グリッド
	XAxis           []*axis      `json:"xAxis"`                     // X軸
	YAxis           []*axis      `json:"yAxis"`                     // Y軸
	Tooltip         *tooltip     `json:"tooltip,omitempty"`         // ツールチップ
	Toolbox         *toolbox     `json:"toolbox,omitempty"`         // ツールボックス
	AxisPointer     *axisPointer `json:"axisPointer,omitempty"`     // 軸ポインタの設定
	DataZoom        []*dataZoom  `json:"dataZoom,omitempty"`        // ツールボックス
	Dataset         *dataset     `json:"dataset,omitempty"`         // データセット
	Series          []*series    `json:"series"`                    // シリーズ
	BackgroundColor string       `json:"backgroundColor,omitempty"` // 背景色
	Animation       Bool         `json:"animation,omitempty"`       // アニメーションするか
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
	Type   string   `json:"type,omitempty"`   // 種別
	Id     string   `json:"id,omitempty"`     // ID
	Left   string   `json:"left,omitempty"`   // 左からの位置
	Top    string   `json:"top,omitempty"`    // 上からの位置
	Right  string   `json:"right,omitempty"`  // 右からの位置
	Bottom string   `json:"bottom,omitempty"` // 下からの位置
	Data   []string `json:"data,omitempty"`   // 凡例
}

type grid struct {
	Id           string `json:"id,omitempty"`           // ID
	Left         string `json:"left,omitempty"`         // 左
	Top          string `json:"top,omitempty"`          // 上
	Right        string `json:"right,omitempty"`        // 右
	Bottom       string `json:"bottom,omitempty"`       // 下
	Height       string `json:"height,omitempty"`       // 高さ
	ContainLabel Bool   `json:"containLabel,omitempty"` // グリッド領域に軸目盛を含めるか
}

// tooltip - ツールチップ
type tooltip struct {
	Trigger         TooltipTrigger `json:"trigger,omitempty"`         // トリガー
	BackgroundColor string         `json:"backgroundColor,omitempty"` // 背景色
	BorderWidth     string         `json:"borderWidth,omitempty"`     // 境界線の幅
	BorderColor     string         `json:"borderColor,omitempty"`     // 境界線の色
	Padding         string         `json:"padding,omitempty"`         // 余白
	AxisPointer     *axisPointer   `json:"axisPointer,omitempty"`     // 軸ポインタ
	TextStyle       *textStyle     `json:"textStyle,omitempty"`       // テキストのスタイル
	Formatter       string         `json:"formatter,omitempty"`       // フォーマッター
}

// textStyle - テキストのスタイル
type textStyle struct {
	Color string `json:"color,omitempty"`
}

// toolbox - ツールボックス
type toolbox struct {
	Feature *toolboxFeature `json:"feature,omitempty"` // 機能
}

// toolboxFeature - ツールボックスの機能
type toolboxFeature struct {
	SaveAsImage *saveAsImageFeature `json:"saveAsImage,omitempty"`
	DataView    *dataViewFeature    `json:"dataView,omitempty"`
	DataZoom    *dataZoomFeature    `json:"dataZoom,omitempty"`
	MagicType   *magicTypeFeature   `json:"magicType,omitempty"`
	Restore     *restoreFeature     `json:"restore,omitempty"`
	Brush       *brushFeature       `json:"brush,omitempty"`
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

// dataZoomFeature - データのズーム機能
type dataZoomFeature struct {
	YAxisIndex Bool `json:"yAxisIndex,omitempty"` // Y軸
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

// brushFeature - ブラシ機能
type brushFeature struct {
	Type []BrushType `json:"type"`
}

// dataZoom - データの拡大縮小
type dataZoom struct {
	Type       string `json:"type,omitempty"`       // 種別
	Show       Bool   `json:"show,omitempty"`       // 表示するか
	XAxisIndex []int  `json:"xAxisIndex,omitempty"` //
	Top        string `json:"top,omitempty"`        //
	Bottom     string `json:"bottom,omitempty"`     //
	Start      int    `json:"start,omitempty"`      //
	End        int    `json:"end,omitempty"`        //
	HandleIcon string `json:"handleIcon,omitempty"` //
	HandleSize string `json:"handleSize,omitempty"` //
}

// axis - 軸
type axis struct {
	Id           string        `json:"id,omitempty"`          // ID
	Type         AxisType      `json:"type,omitempty"`        // 種別
	Name         string        `json:"name,omitempty"`        // 名前
	Min          string        `json:"min,omitempty"`         // 最小
	Max          string        `json:"max,omitempty"`         // 最大
	Interval     int           `json:"interval,omitempty"`    // 間隔
	GridIndex    int           `json:"gridIndex,omitempty"`   // グリッド設定のインデックス
	SplitNumber  int           `json:"splitNumber,omitempty"` // 分割
	BoundaryGap  Bool          `json:"boundaryGap,omitempty"` // 境界部分のギャップ
	Scale        Bool          `json:"scale,omitempty"`       // スケールするか
	Data         []interface{} `json:"data,omitempty"`        // データ
	AxisLine     *axisLine     `json:"axisLine,omitempty"`    // 軸線
	AxisTick     *axisTick     `json:"axisTick,omitempty"`    // 軸線
	AxisLabel    *axisLabel    `json:"axisLabel,omitempty"`   // 軸のラベル
	AxisPointer  *axisPointer  `json:"axisPointer,omitempty"` // 軸ポインタ
	SplitLine    *splitLine    `json:"splitLine,omitempty"`   // 分割線
	SplitAreaOpt *splitAreaOpt `json:"splitArea,omitempty"`   // 分割線
}

// axisLine - 軸線
type axisLine struct {
	Show   Bool `json:"show,omitempty"`   // 表示するか
	OnZero Bool `json:"onZero,omitempty"` // 原点表示について
}

// axisTick - 軸目盛
type axisTick struct {
	Show Bool `json:"show,omitempty"` // 表示するか
}

// axisLabel - 軸のラベル
type axisLabel struct {
	Show      Bool   `json:"show,omitempty"`      // 表示するか
	Formatter string `json:"formatter,omitempty"` // フォーマッター
}

// axisPointer - 軸ポインタ
type axisPointer struct {
	Type       AxisPointerType `json:"type,omitempty"`       // 種別
	Z          int             `json:"z,omitempty"`          // 奥行
	CrossStyle *axisStyle      `json:"crossStyle,omitempty"` // クロスポインタのスタイル
	Link       *link           `json:"link,omitempty"`       // リンク
}

// link - リンクさせる軸線
type link struct {
	XAxisIndex []int `json:"xAxisIndex,omitempty"` // リンクさせるX軸
}

// axisStyle - 軸のスタイル
type axisStyle struct {
	Color string `json:"color,omitempty"` // 色
}

// splitLine - 分割線
type splitLine struct {
	Show Bool `json:"show"` // 表示するか
}

// splitAreaOpt - エリア分割
type splitAreaOpt struct {
	Show Bool `json:"show"` // 表示するか
}

// dataset - データセット
type dataset struct {
	Source []interface{} `json:"source,omitempty"` // データソース
}

// series - シリーズ
type series struct {
	Type           SeriesType    `json:"type,omitempty"`           // 種別
	Id             string        `json:"id,omitempty"`             // ID
	Name           string        `json:"name,omitempty"`           // 名前
	XAxisIndex     int           `json:"xAxisIndex,omitempty"`     // X軸のインデックス
	YAxisIndex     int           `json:"yAxisIndex,omitempty"`     // Y軸のインデックス
	Smooth         Bool          `json:"smooth,omitempty"`         // スムーズにするか
	HoverAnimation Bool          `json:"hoverAnimation,omitempty"` // hover時のアニメーション
	Symbol         SymbolType    `json:"symbol,omitempty"`         // シンボルの形
	SymbolSize     int           `json:"symbolSize,omitempty"`     // シンボルのサイズ
	ShowSymbol     Bool          `json:"showSymbol,omitempty"`     // シンボルの表示非表示
	Stack          string        `json:"stack,omitempty"`          // 積み上げ
	Large          Bool          `json:"large,omitempty"`          //
	Data           []interface{} `json:"data,omitempty"`           // データ
	ItemStyle      *itemStyle    `json:"itemStyle,omitempty"`      // データのスタイル
	LineStyle      *lineStyle    `json:"lineStyle,omitempty"`      // 線のスタイル
	AreaStyle      *areaStyle    `json:"areaStyle,omitempty"`      // エリアのスタイル
	Encode         *encode       `json:"encode,omitempty"`         // エンコード
}

// itemStyle - データのスタイル
type itemStyle struct {
	Color        string `json:"color,omitempty"`        // 色
	Color0       string `json:"color0,omitempty"`       // 色0
	BorderColor  string `json:"borderColor,omitempty"`  // 境界線の色
	BorderColor0 string `json:"borderColor0,omitempty"` // 境界線の色0
}

// lineStyle - 線のスタイル
type lineStyle struct {
	Opacity float64 `json:"opacity"` // 不透明度
}

// areaStyle - エリアのスタイル
type areaStyle struct {
	Color string `json:"color,omitempty"` // 色
}

// encode - データセットからどのデータを取るか
type encode struct {
	X []int `json:"x"`           // X
	Y []int `json:"y,omitempty"` // Y
}
