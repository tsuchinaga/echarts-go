package chart

import (
	"encoding/json"
	"html/template"
	"io"
)

// Charter - チャートのインターフェース
type Charter interface {
	Render(w io.Writer) error
	Json() ([]byte, error)
	SetOption(options ...chartOpt) Charter
	SetXAxis(id string, t AxisType, data []interface{}, options ...axisOpt) Charter
	SetYAxis(id string, t AxisType, data []interface{}, options ...axisOpt) Charter
	SetSeries(id string, t SeriesType, data []interface{}, options ...seriesOpt) Charter
	SetToolboxFeature(options ...toolboxOpt) Charter
}

// New - 新しいチャートの生成
func New(title string, width, height int) Charter {
	return &Chart{
		Title:  title,
		Width:  width,
		Height: height,
		Option: &option{
			XAxis:  []*axis{},
			YAxis:  []*axis{},
			Series: []*series{},
		},
		x:      map[string]int{},
		y:      map[string]int{},
		series: map[string]int{},
	}
}

// Chart - チャート
type Chart struct {
	Charter
	Title  string         // HTML headのタイトル
	Width  int            // グラフエリアの幅
	Height int            // グラフエリアの高さ
	Option *option        // チャートのオプション
	x      map[string]int // x軸
	y      map[string]int // y軸
	series map[string]int // シリーズ
}

// Render - 引数のWriterにhtmlを出力する
func (c *Chart) Render(w io.Writer) error {
	tpl, err := template.New("").Parse(baseHtml)
	if err != nil {
		return err
	}

	// 必須項目を埋めておく
	if len(c.x) == 0 {
		c.Option.XAxis = []*axis{{}}
	}
	if len(c.y) == 0 {
		c.Option.YAxis = []*axis{{}}
	}

	err = tpl.Execute(w, c)
	return nil
}

// Json - optionの中身をjsonで出力
func (c *Chart) Json() ([]byte, error) {
	return json.Marshal(c.Option)
}

// SetOption - オプションの追加
func (c *Chart) SetOption(options ...chartOpt) Charter {
	for _, o := range options {
		o.setOpt(c.Option)
	}
	return c
}

// SetXAxis - x軸の追加
func (c *Chart) SetXAxis(id string, t AxisType, data []interface{}, options ...axisOpt) Charter {
	a := &axis{Id: id, Type: t, Data: data}
	for _, o := range options {
		o.setOpt(a)
	}

	// 同一のidが存在すれば上書き
	if i, ok := c.x[id]; ok {
		c.Option.XAxis[i] = a
		return c
	}

	c.Option.XAxis = append(c.Option.XAxis, a)
	c.x[id] = len(c.Option.XAxis) - 1
	return c
}

// SetYAxis - y軸の追加
func (c *Chart) SetYAxis(id string, t AxisType, data []interface{}, options ...axisOpt) Charter {
	a := &axis{Id: id, Type: t, Data: data}
	for _, o := range options {
		o.setOpt(a)
	}

	// 同一のidが存在すれば上書き
	if i, ok := c.y[id]; ok {
		c.Option.YAxis[i] = a
		return c
	}

	c.Option.YAxis = append(c.Option.YAxis, a)
	c.y[id] = len(c.Option.YAxis) - 1
	return c
}

// SetSeries - シリーズの追加
func (c *Chart) SetSeries(id string, t SeriesType, data []interface{}, options ...seriesOpt) Charter {
	s := &series{Id: id, Type: t, Data: data}
	for _, o := range options {
		o.setOpt(s)
	}

	// 同一のidが存在すれば上書き
	if i, ok := c.series[id]; ok {
		c.Option.Series[i] = s
		return c
	}

	c.Option.Series = append(c.Option.Series, s)
	c.series[id] = len(c.Option.Series) - 1
	return c
}

// SetToolboxFeature - ツールボックスの機能の追加
func (c *Chart) SetToolboxFeature(options ...toolboxOpt) Charter {
	if c.Option.Toolbox == nil {
		c.Option.Toolbox = new(toolbox)
	}

	for _, o := range options {
		o.setOpt(c.Option.Toolbox)
	}

	return c
}
