package chart

var baseHtml = `<html lang="en">
<head>
    <meta charset="utf-8">
    <title>{{.Title}}</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/echarts/4.7.0/echarts-en.min.js"></script>
</head>
<body>
<div id="main" style="width:{{.Width}}px;height:{{.Height}}px;margin:auto;"></div>
<script type="text/javascript">
    let myChart = echarts.init(document.getElementById("main"));
    let Option = {{.Option}};
    myChart.setOption(Option);
</script>
</body>
</html>
`
