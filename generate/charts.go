package generate

import (
	h "github.com/apyhub/apyhub/helper"
)

func BarChartAsFile(chart h.Chart) ([]byte, error) {
	return prepareAPI(chart, "POST", h.BarChartFile, false)
}

func BarChartAsURL(chart h.Chart) (string, error) {
	byt, err := prepareAPI(chart, "POST", h.BarChartUrl, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

func PieChartAsFile(chart h.Chart) ([]byte, error) {
	return prepareAPI(chart, "POST", h.PieChartFile, false)
}

func PieChartAsURL(chart h.Chart) (string, error) {
	byt, err := prepareAPI(chart, "POST", h.PieChartUrl, true)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

func StackedBarChartAsFile(chart h.StackedBarChart) (byt []byte, err error) {
	return prepareAPI(chart, "POST", h.StackedBarChartFile, false)
}

func StackedBarChartAsURL(chart h.StackedBarChart) (url string, err error) {
	byt, err := prepareAPI(chart, "POST", h.StackedBarChartUrl, true)
	if err != nil {
		return
	}
	return string(byt), nil
}
