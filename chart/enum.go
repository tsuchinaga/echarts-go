package chart

// Bool - 真偽値
type Bool string

const (
	BoolEmpty Bool = ""
	BoolNull  Bool = "null"
	True      Bool = "true"
	False     Bool = "false"
)

func (e Bool) String() string {
	return string(e)
}

func (e Bool) MarshalJSON() ([]byte, error) {
	switch e {
	case True:
		return []byte("true"), nil
	case False:
		return []byte("false"), nil
	default:
		return []byte("null"), nil
	}
}

// AxisType - 軸の種別
type AxisType string

const (
	AxisTypeEmpty    AxisType = ""
	AxisTypeNull     AxisType = "null"
	AxisTypeValue    AxisType = "value"
	AxisTypeCategory AxisType = "category"
	AxisTypeTime     AxisType = "time"
	AxisTypeLog      AxisType = "log"
)

func (e AxisType) String() string {
	return string(e)
}

// SeriesType - シリーズの種別
type SeriesType string

const (
	SeriesTypeEmpty         SeriesType = ""
	SeriesTypeNull          SeriesType = "null"
	SeriesTypeLine          SeriesType = "line"
	SeriesTypeBar           SeriesType = "bar"
	SeriesTypePie           SeriesType = "pie"
	SeriesTypeScatter       SeriesType = "scatter"
	SeriesTypeEffectScatter SeriesType = "effectScatter"
	SeriesTypeRadar         SeriesType = "radar"
	SeriesTypeTree          SeriesType = "tree"
	SeriesTypeTreemap       SeriesType = "treemap"
	SeriesTypeSunburst      SeriesType = "sunburst"
	SeriesTypeBoxplot       SeriesType = "boxplot"
	SeriesTypeCandlestick   SeriesType = "candlestick"
	SeriesTypeHeadmap       SeriesType = "heatmap"
	SeriesTypeMap           SeriesType = "map"
	SeriesTypeParallel      SeriesType = "parallel"
	SeriesTypeLines         SeriesType = "lines"
	SeriesTypeGraph         SeriesType = "graph"
	SeriesTypeSnakey        SeriesType = "snakey"
	SeriesTypeFunnel        SeriesType = "funnel"
	SeriesTypeGauge         SeriesType = "gauge"
	SeriesTypePictorialBar  SeriesType = "pictorialBar"
	SeriesTypeThemeRiver    SeriesType = "themeRiver"
	SeriesTypeCustom        SeriesType = "custom"
)

func (e SeriesType) String() string {
	return string(e)
}

// AxisPointerType - 軸ポインタの種別
type AxisPointerType string

const (
	AxisPointerTypeEmpty  AxisPointerType = ""
	AxisPointerTypeNull   AxisPointerType = "null"
	AxisPointerTypeNone   AxisPointerType = "none"
	AxisPointerTypeLine   AxisPointerType = "line"
	AxisPointerTypeShadow AxisPointerType = "shadow"
	AxisPointerTypeCross  AxisPointerType = "cross"
)

func (e AxisPointerType) String() string {
	return string(e)
}

// TooltipTrigger - ツールチップのトリガー
type TooltipTrigger string

const (
	TooltipTriggerEmpty TooltipTrigger = ""
	TooltipTriggerNull  TooltipTrigger = "null"
	TooltipTriggerNone  TooltipTrigger = "none"
	TooltipTriggerItem  TooltipTrigger = "item"
	TooltipTriggerAxis  TooltipTrigger = "axis"
)

func (e TooltipTrigger) String() string {
	return string(e)
}

// SymbolType - シンボルの形
type SymbolType string

const (
	SymbolTypeEmpty     SymbolType = ""
	SymbolTypeNull      SymbolType = "null"
	SymbolTypeNone      SymbolType = "none"
	SymbolTypeCircle    SymbolType = "circle"
	SymbolTypeRect      SymbolType = "rect"
	SymbolTypeRoundRect            = "roundRect"
	SymbolTypeTriangle  SymbolType = "triangle"
	SymbolTypeDiamond   SymbolType = "diamond"
	SymbolTypePin       SymbolType = "pin"
	SymbolTypeArrow     SymbolType = "arrow"
)

func (e SymbolType) String() string {
	return string(e)
}
