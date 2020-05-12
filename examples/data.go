package main

type Data struct {
	Date  string
	Value float64
	Low   float64
	High  float64
}

var set1 = []Data{
	{Date: "2012-08-28", Value: -1.1618426259, Low: -2.6017329022, High: 0.2949717757},
	{Date: "2012-08-29", Value: -0.5828247293, Low: -1.3166963635, High: 0.1324086347},
	{Date: "2012-08-30", Value: -0.3790770636, Low: -0.8712221305, High: 0.0956413566},
	{Date: "2012-08-31", Value: -0.2792926002, Low: -0.6541832008, High: 0.0717120241},
	{Date: "2012-09-01", Value: -0.2461165469, Low: -0.5222677907, High: 0.0594188803},
	{Date: "2012-09-02", Value: -0.2017354137, Low: -0.4434280535, High: 0.0419213465},
	{Date: "2012-09-03", Value: -0.1457476871, Low: -0.3543957712, High: 0.0623761171},
	{Date: "2012-09-04", Value: -0.002610973, Low: -0.3339911495, High: 0.031286929},
	{Date: "2012-09-05", Value: -0.0080692734, Low: -0.2951839941, High: 0.0301762553},
	{Date: "2012-09-06", Value: -0.0296490933, Low: -0.2964395801, High: -0.0029821004},
	{Date: "2012-09-07", Value: 0.001317397, Low: -0.2295443759, High: 0.037903312},
	{Date: "2012-09-08", Value: -0.0117649838, Low: -0.2226376418, High: 0.0239720183},
	{Date: "2012-09-09", Value: 0.0059394263, Low: -0.2020479849, High: 0.0259489347},
	{Date: "2012-09-10", Value: -0.0115565898, Low: -0.2042048037, High: 0.0077863806},
	{Date: "2012-09-11", Value: 0.0041183019, Low: -0.1837263172, High: 0.0137898406},
	{Date: "2012-09-12", Value: 0.0353559544, Low: -0.136610008, High: 0.051403828},
	{Date: "2012-09-13", Value: 0.0070046011, Low: -0.1569988647, High: 0.0202266411},
	{Date: "2012-09-14", Value: -0.0004251807, Low: -0.1410340292, High: 0.0273410185},
	{Date: "2012-09-15", Value: -0.0035461023, Low: -0.1438653689, High: 0.0165445684},
	{Date: "2012-09-16", Value: 0.007797889, Low: -0.1291975355, High: 0.0232461153},
	{Date: "2012-09-17", Value: 0.0025402723, Low: -0.133972479, High: 0.0116753921},
	{Date: "2012-09-18", Value: -0.005317381, Low: -0.1269266586, High: 0.0129723291},
	{Date: "2012-09-19", Value: -0.0075841521, Low: -0.1283478383, High: 0.0056371616},
	{Date: "2012-09-20", Value: -0.0391388721, Low: -0.1571172198, High: -0.0311678828},
	{Date: "2012-09-21", Value: 0.0075430252, Low: -0.1097354417, High: 0.0141132062},
	{Date: "2012-09-22", Value: 0.1850284663, Low: 0.0333682152, High: 0.2140709422},
	{Date: "2012-09-23", Value: 0.076629596, Low: -0.0068472967, High: 0.1101280569},
	{Date: "2012-09-24", Value: -0.0314292271, Low: -0.1074281762, High: 0.0032669363},
	{Date: "2012-09-25", Value: -0.0232608674, Low: -0.0905197842, High: 0.0164250295},
	{Date: "2012-09-26", Value: -0.01968615, Low: -0.084319856, High: 0.0193319465},
	{Date: "2012-09-27", Value: -0.0310196816, Low: -0.0914356781, High: 0.0094436256},
	{Date: "2012-09-28", Value: -0.0758746967, Low: -0.1169814745, High: -0.019659551},
	{Date: "2012-09-29", Value: 0.0233974572, Low: -0.0356839258, High: 0.0610712506},
	{Date: "2012-09-30", Value: 0.011073579, Low: -0.0558712863, High: 0.0346160081},
	{Date: "2012-10-01", Value: -0.002094822, Low: -0.0707143388, High: 0.0152899266},
	{Date: "2012-10-02", Value: -0.1083707096, Low: -0.1718101335, High: -0.0886271057},
	{Date: "2012-10-03", Value: -0.1098258972, Low: -0.1881274065, High: -0.1072157972},
	{Date: "2012-10-04", Value: -0.0872970297, Low: -0.1731903321, High: -0.064381434},
	{Date: "2012-10-05", Value: -0.0761992047, Low: -0.1770373817, High: 0.100085727},
	{Date: "2012-10-06", Value: -0.0416654249, Low: -0.1502479611, High: 0.0751148102},
	{Date: "2012-10-07", Value: -0.0410128962, Low: -0.1618694445, High: 0.0881453482},
	{Date: "2012-10-08", Value: -0.0214289042, Low: -0.1590852977, High: 0.0871880288},
	{Date: "2012-10-09", Value: 0.2430880604, Low: 0.063624221, High: 0.2455101587},
	{Date: "2012-10-10", Value: 0.3472823479, Low: 0.1553854927, High: 0.3583991097},
	{Date: "2012-10-11", Value: 0.3360734074, Low: 0.2055952772, High: 0.3812162823},
	{Date: "2012-10-12", Value: -0.0463648355, Low: -0.0626466998, High: 0.0037342957},
	{Date: "2012-10-13", Value: -0.0867009379, Low: -0.0867594055, High: -0.0223791074},
	{Date: "2012-10-14", Value: -0.1288672826, Low: -0.1161709129, High: -0.0534789124},
	{Date: "2012-10-15", Value: -0.1474426821, Low: -0.1559759048, High: -0.0646995092},
	{Date: "2012-10-16", Value: -0.1502405066, Low: -0.1604364638, High: -0.0602562376},
	{Date: "2012-10-17", Value: -0.1203765529, Low: -0.1569023195, High: -0.0578129637},
	{Date: "2012-10-18", Value: -0.0649122919, Low: -0.0782987564, High: -0.0501999174},
	{Date: "2012-10-19", Value: -0.015525562, Low: -0.1103873808, High: -0.0132131311},
	{Date: "2012-10-20", Value: -0.006051357, Low: -0.1089644497, High: 0.0230384197},
	{Date: "2012-10-21", Value: 0.0003154213, Low: -0.1073849227, High: 0.0017290437},
	{Date: "2012-10-22", Value: -0.0063018298, Low: -0.1120298155, High: 0.0173284555},
	{Date: "2012-10-23", Value: -0.004294834, Low: -0.1076841119, High: 0.0547933965},
	{Date: "2012-10-24", Value: -0.0053400832, Low: -0.1096991408, High: 0.0560555803},
	{Date: "2012-10-25", Value: 0.0070057212, Low: -0.0940613813, High: 0.0425517607},
	{Date: "2012-10-26", Value: 0.0082121656, Low: -0.0906810455, High: 0.0396884383},
	{Date: "2012-10-27", Value: 0.0141422884, Low: -0.0841305678, High: 0.0340050012},
	{Date: "2012-10-28", Value: 0.0041613553, Low: -0.0886723749, High: 0.039426727},
	{Date: "2012-10-29", Value: -0.0013614287, Low: -0.0923481608, High: 0.0438725574},
	{Date: "2012-10-30", Value: -0.0052144933, Low: -0.0937763043, High: 0.0459998555},
	{Date: "2012-10-31", Value: 0.0078904741, Low: -0.0807028001, High: 0.0334824169},
	{Date: "2012-11-01", Value: 0.0099598702, Low: -0.0740001323, High: 0.0280264274},
	{Date: "2012-11-02", Value: 0.0001146029, Low: -0.0820430294, High: 0.0326771125},
	{Date: "2012-11-03", Value: 0.0047572651, Low: -0.0754113825, High: 0.0294912577},
	{Date: "2012-11-04", Value: 0.006204557, Low: -0.0750627059, High: 0.029693607},
	{Date: "2012-11-05", Value: 0.0115231406, Low: -0.0663484142, High: 0.0214084056},
	{Date: "2012-11-06", Value: -0.0032634994, Low: -0.0793170451, High: 0.0355159827},
	{Date: "2012-11-07", Value: -0.0108985452, Low: -0.0846123893, High: 0.0409797057},
	{Date: "2012-11-08", Value: -0.0092766813, Low: -0.0802668328, High: 0.0373886301},
	{Date: "2012-11-09", Value: 0.0095972086, Low: -0.0623739694, High: 0.0194918693},
	{Date: "2012-11-10", Value: -0.0111809358, Low: -0.0819555908, High: 0.038335749},
	{Date: "2012-11-11", Value: -0.0023572296, Low: -0.0745443377, High: 0.0306093592},
	{Date: "2012-11-12", Value: 0.0084213775, Low: -0.0657707155, High: 0.0227270619},
	{Date: "2012-11-13", Value: 0.0107446453, Low: -0.0617995017, High: 0.0196547867},
	{Date: "2012-11-14", Value: 0.009457792, Low: -0.0597697849, High: 0.0191832343},
	{Date: "2012-11-15", Value: 0.0031194779, Low: -0.0589126783, High: 0.0186409442},
	{Date: "2012-11-16", Value: -0.0115128213, Low: -0.0767105447, High: 0.0370292452},
	{Date: "2012-11-17", Value: 0.0058347339, Low: -0.0592236472, High: 0.0198181452},
	{Date: "2012-11-18", Value: -0.0235630436, Low: -0.083529944, High: 0.046280909},
	{Date: "2012-11-19", Value: -0.0479795964, Low: -0.1086422529, High: 0.0113044645},
	{Date: "2012-11-21", Value: -0.0218184359, Low: -0.0881634878, High: 0.0448568265},
	{Date: "2012-11-28", Value: -0.0071361172, Low: -0.0807350229, High: 0.0453599734},
	{Date: "2012-12-05", Value: -0.0151966912, Low: -0.089995793, High: 0.0558329569},
	{Date: "2012-12-12", Value: -0.0097784855, Low: -0.089466481, High: 0.0550191387},
	{Date: "2012-12-19", Value: -0.0095681495, Low: -0.090513354, High: 0.057073314},
	{Date: "2012-12-27", Value: -0.0034165915, Low: -0.0907151292, High: 0.0561479112},
	{Date: "2012-12-31", Value: 0.3297981389, Low: 0.1537781522, High: 0.3499473316},
}