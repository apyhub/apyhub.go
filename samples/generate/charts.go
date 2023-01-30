package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/apyhub/apyhub.go"
	h "github.com/apyhub/apyhub.go/helper"
)

func init() {
	apyhub.InitApyHub("Enter your apyhub token")
}

func Barchart() {
	barChart := h.Chart{
		Title: "BarChart",
		Theme: "Dark",
		Data: []h.Values{
			{
				Label: "bar 1",
				Value: 10.5,
			},
			{
				Label: "bar 2",
				Value: 15,
			},
			{
				Label: "bar 3",
				Value: 5,
			},
		},
	}
	byt, err := apyhub.BarChartAsFile(barChart)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/bar.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.BarChartAsURL(barChart)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func Piechart() {
	pieChart := h.Chart{
		Title: "PieChart",
		Theme: "Dark",
		Data: []h.Values{
			{
				Label: "pie 1",
				Value: 10.5,
			},
			{
				Label: "pie 2",
				Value: 15,
			},
			{
				Label: "pie 3",
				Value: 5,
			},
		},
	}
	byt, err := apyhub.PieChartAsFile(pieChart)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/pie.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.PieChartAsURL(pieChart)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}

func Stackedchart() {
	stackedChart := h.StackedBarChart{
		Title: "BarChart",
		Theme: "Light",
		Data: []h.StackedValue{
			{
				Name: "",
				Values: []h.Values{
					{
						Label: "stacked 1",
						Value: 10.5,
					},
					{
						Label: "stacked 2",
						Value: 15,
					},
					{
						Label: "stacked 3",
						Value: 5,
					},
				},
			},
		},
	}
	byt, err := apyhub.StackedBarChartAsFile(stackedChart)
	if err != nil {
		log.Fatal(err)
	}

	// write bytes in specific file
	if err := os.WriteFile("./test_file/stacked.png", byt, 0666); err != nil {
		log.Fatal(err)
	}

	url, err := apyhub.StackedBarChartAsURL(stackedChart)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
