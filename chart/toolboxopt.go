package chart

type toolboxOpt interface {
	getToolboxOptType() string
	setOpt(toolbox *toolbox)
}

// SaveAsImageFeatureOpt - 画像保存機能に関する設定
type SaveAsImageFeatureOpt struct {
	toolboxOpt
	Show Bool // 表示するか
}

func (o *SaveAsImageFeatureOpt) getToolboxOptType() string {
	return "saveAsImage"
}
func (o *SaveAsImageFeatureOpt) setOpt(toolbox *toolbox) {
	if toolbox.Feature == nil {
		toolbox.Feature = new(toolboxFeature)
	}
	toolbox.Feature.SaveAsImage = &saveAsImageFeature{Show: o.Show}
}

// DataViewFeatureOpt - データ閲覧機能に関する設定
type DataViewFeatureOpt struct {
	toolboxOpt
	Show     Bool // 表示するか
	ReadOnly Bool // 読み込み専用か
}

func (o *DataViewFeatureOpt) getToolboxOptType() string {
	return "dataView"
}
func (o *DataViewFeatureOpt) setOpt(toolbox *toolbox) {
	if toolbox.Feature == nil {
		toolbox.Feature = new(toolboxFeature)
	}
	toolbox.Feature.DataView = &dataViewFeature{Show: o.Show, ReadOnly: o.ReadOnly}
}

// DataZoomFeatureOpt - データズーム機能に関する設定
type DataZoomFeatureOpt struct {
	toolboxOpt
	YAxisIndex Bool // Y軸
}

func (o *DataZoomFeatureOpt) getToolboxOptType() string {
	return "dataZoom"
}
func (o *DataZoomFeatureOpt) setOpt(toolbox *toolbox) {
	if toolbox.Feature == nil {
		toolbox.Feature = new(toolboxFeature)
	}
	toolbox.Feature.DataZoom = &dataZoomFeature{YAxisIndex: o.YAxisIndex}
}

// MagicTypeFeatureOpt - シリーズ種別変更に関する設定
type MagicTypeFeatureOpt struct {
	toolboxOpt
	Show Bool         // 表示するか
	Type []SeriesType // 選べるシリーズ種別
}

func (o *MagicTypeFeatureOpt) getToolboxOptType() string {
	return "magicType"
}
func (o *MagicTypeFeatureOpt) setOpt(toolbox *toolbox) {
	if toolbox.Feature == nil {
		toolbox.Feature = new(toolboxFeature)
	}
	toolbox.Feature.MagicType = &magicTypeFeature{Show: o.Show, Type: o.Type}
}

// RestoreFeatureOpt - 元のチャートに戻す機能に関する設定
type RestoreFeatureOpt struct {
	toolboxOpt
	Show Bool // 表示するか
}

func (o *RestoreFeatureOpt) getToolboxOptType() string {
	return "restore"
}
func (o *RestoreFeatureOpt) setOpt(toolbox *toolbox) {
	if toolbox.Feature == nil {
		toolbox.Feature = new(toolboxFeature)
	}
	toolbox.Feature.Restore = &restoreFeature{Show: o.Show}
}
