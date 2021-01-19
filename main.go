package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-ping/ping"
)

func getPingData(url string) ([]opts.LineData, []int32, string) {
	pinger, _ := ping.NewPinger(url)
	pinger.Count = 10
	pinger.Run()
	var rttsMS []int64
	var XAxis []int32
	for i := range pinger.Statistics().Rtts {
		rttsMS = append(rttsMS, int64(pinger.Statistics().Rtts[i]/time.Millisecond))
	}
	var items []opts.LineData
	for i := 0; i < len(rttsMS); i++ {
		XAxis = append(XAxis, int32(i+1))
		items = append(items, opts.LineData{Value: rttsMS[i], Symbol: "diamond", SymbolSize: 10})
	}
	return items, XAxis, pinger.Statistics().IPAddr.String()
}

func httpserver(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()["target"][0]
	lineData, XAxis, ipInfo := getPingData(url)

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    url,
			Subtitle: ipInfo,
		}))
	line.SetXAxis(XAxis).
		AddSeries("Category A", lineData).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	log.Println("start to listen on :8081")
	http.ListenAndServe(":8081", nil)
}
