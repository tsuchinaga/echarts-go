# examples

公式にあったチャートを描いてみる

## 起動
`go run examples/* -p 8080`

`-p`で指定したポート番号でWebサーバが立ちます(省略した場合は8080)

あとはlocalhostで各ページにアクセスすることで結果が見れます

## 一覧

* [Get Started with ECharts in 5 minutes](https://echarts.apache.org/en/tutorial.html#Get%20Started%20with%20ECharts%20in%205%20minutes)
    * [http://localhost:8080/get-started](http://localhost:8080/get-started)
    * [getstarted.go](getstarted.go)
* Line
    * [Basic Line Chart](https://echarts.apache.org/examples/en/editor.html?c=line-simple)
        * [http://localhost:8080/line-simple](http://localhost:8080/line-simple)
        * [linesimple.go](linesimple.go)
    * [Smoothed Line Chart](https://echarts.apache.org/examples/en/editor.html?c=line-smooth)
        * [http://localhost:8080/line-smooth](http://localhost:8080/line-smooth)
        * [linesmooth.go](linesmooth.go)
    * [Confidence Band](https://echarts.apache.org/examples/en/editor.html?c=confidence-band)
        * [http://localhost:8080/confidence-band](http://localhost:8080/confidence-band)
        * [confidenceband.go](confidenceband.go)
        * ※ラベルのフォーマッタなどで無名関数を使っている部分が実現できないためにY軸とかおかしい
* Bar
    * [Bar Simple](https://echarts.apache.org/examples/en/editor.html?c=bar-simple)
        * [http://localhost:8080/bar-simple](http://localhost:8080/bar-simple)
        * [barsimple.go](barsimple.go)
* Candlestick
    * [Basic Candlestick](https://echarts.apache.org/examples/en/editor.html?c=candlestick-simple)
        * [http://localhost:8080/candlestick-simple](http://localhost:8080/candlestick-simple)
        * [candlesticksimple.go](candlesticksimple.go)
    * [Large Scale Candlestick](https://echarts.apache.org/examples/en/editor.html?c=candlestick-large)
        * [http://localhost:8080/candlestick-large](http://localhost:8080/candlestick-large)
        * [candlesticklarge.go](candlesticklarge.go)
    * [Candlestick Brush](https://echarts.apache.org/examples/en/editor.html?c=candlestick-brush)
        * [http://localhost:8080/candlestick-brush](http://localhost:8080/candlestick-brush)
        * [candlestickbrush.go](candlestickbrush.go)
* Mixed
    * [Mixed Line and Bar](https://echarts.apache.org/examples/en/editor.html?c=mix-line-bar)
        * [http://localhost:8080/mix-line-bar](http://localhost:8080/mix-line-bar) 
        * [mixlinebar.go](mixlinebar.go)
