package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 5; i++ {
		items = append(items, opts.LineData{Name: strconv.Itoa(i + 1), Value: rand.Intn(20), Symbol: "diamond", SymbolSize: 10})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWonderland}),
		charts.WithTitleOpts(opts.Title{
			Title:    "www.google.com",
			Subtitle: "1.1.1.1",
		}))

	// Put data into instance
	line.SetXAxis([]int32{1, 2, 3, 4, 5}).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
