package apyhub

import (
	"github.com/apyhub/apyhub.go/convert"
	"github.com/apyhub/apyhub.go/data"
	"github.com/apyhub/apyhub.go/extract"
	"github.com/apyhub/apyhub.go/generate"
	h "github.com/apyhub/apyhub.go/helper"
	"github.com/apyhub/apyhub.go/imageprocessor"
	"github.com/apyhub/apyhub.go/search"
	"github.com/apyhub/apyhub.go/validation"
)

// Set the token
func InitApyHub(auth interface{}) error {
	switch auth := auth.(type) {
	case string:
		h.Token = auth
	case h.BasicAuth:
		h.Auth = h.BasicAuth(auth)
	default:
		return h.ErrTypeNotMatch
	}
	return nil
}

/*
##Functions

### csvToExcel

Converts a CSV file or URL to an Excel file.

**`Example 1 : convert csv file to Excel file`**

```go

	import (
		apyhub "github.com/apyhub/apyhub.go"
	)

	File,err :=os.Open("any csv file")
	if err!=nil {
		log.Fatal(err)
	}

	// Input is a file
	byt,err :=apyhub.CsvToExcelAsFile(File)

```

**`Example 2 : convert csv file url to Excel file`**

```go

	fileUrl :="csv file url"

	// Input is a url
	byt,err :=apyhub.CsvToExcelAsFile(fileUrl)

```

#### Defined in

[convert/csvToExcel.go:63](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L63)
*/
func CsvToExcelAsFile(urlorFile interface{}) ([]byte, error) {
	return convert.CsvToExcelAsFile(urlorFile)
}

/*
**`Example 3 : convert csv file to AWS presigned url`**

```go

	// Input is a file
	File,err :=os.Open("any csv file")
	if err!=nil {
		log.Fatal(err)
	}

	url,err :=apyhub.CsvToExcelAsUrl(File)

```

**`Example 4 : convert csv file url to AWS presigned url`**

```go

	fileUrl :="csv file url"

	// Input is a url
	url,err :=apyhub.CsvToExcelAsUrl(fileUrl)

```

#### Defined in

[convert/csvToExcel.go:38](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L38)

**`Link`**

https://apyhub.com/utility/converter-csv-excel

#### Parameters

<table>
<tr><th>Name</th><th>Type</th><th>Description</th></tr>
<tr><td>input</td><td> <code>"string"</code> | <code>Buffer</code></td><td>The html  file or URL.</td></tr>
<tr><td>output</td><td><code>[]byte</code> | <code>"string"</code> </td><td>The format for the response.</td></tr>
</table>

#### Returns

The data for the output file as slice of byte or Url as a string.
*/
func CsvToExcelAsUrl(urlorFile interface{}) (string, error) {
	return convert.CsvToExcelAsURL(urlorFile)
}

func HtmlToPdfAsFile(urlorFile interface{}, landscape bool) ([]byte, error) {
	return convert.HtmlToPdfAsFile(urlorFile, landscape)
}

func HtmlToPdfAsUrl(urlorFile interface{}, landscape bool) (string, error) {
	return convert.HtmlToPdfAsUrl(urlorFile, landscape)
}

/*
### imageToPdf

Converts a image file or URL to an pdf file.

**`Example 1 : convert image to pdf file`**

```go

	import (
	apyhub "github.com/apyhub/apyhub.go"
	)

	// Input is a file
	File,err :=os.Open("any image")
	if err!=nil {
		log.Fatal(err)
	}

	byt,err :=apyhub.ImageToPdfAsFile(File)

```

**`Example 2 : convert image url to pdf file`**

```go

	// Input is a url
	fileUrl :="image url"

	byt,err :=apyhub.ImageToPdfAsFile(fileUrl)

```

#### Defined in

[convert/imageToPdf.go:12](https://github.com/apyhub/apyhub.go/blob/main/convert/imageToPdf.go#L12)
*/
func ImageToPdfAsFile(urlorFile interface{}, landscape bool) ([]byte, error) {
	return convert.ImageToPdfAsFile(urlorFile, landscape)
}

func ImageToPdfAsUrl(urlorFile interface{}, landscape bool) (string, error) {
	return convert.ImageToPdfAsUrl(urlorFile, landscape)
}

func MarkdownToHtmlAsFile(urlorFile interface{}) ([]byte, error) {
	return convert.MarkdownToHtmlAsFile(urlorFile)
}

func MarkdownToHtmlAsUrl(urlorFile interface{}) (string, error) {
	return convert.MarkdownToHtmlAsUrl(urlorFile)
}

func PresentationToPdfAsFile(urlorFile interface{}, landscape bool) ([]byte, error) {
	return convert.PresentationToPdfAsFile(urlorFile, landscape)
}

func PresentationToPdfAsUrl(urlorFile interface{}, landscape bool) (string, error) {
	return convert.PresentationToPdfAsUrl(urlorFile, landscape)
}

func SpreadsheetToPdfAsFile(urlorFile interface{}, landscape bool) ([]byte, error) {
	return convert.SpreadsheetToPdfAsFile(urlorFile, landscape)
}

func SpreadsheetToPdfAsUrl(urlorFile interface{}, landscape bool) (string, error) {
	return convert.SpreadsheetToPdfAsUrl(urlorFile, landscape)
}

func WordToPdfAsFile(urlorFile interface{}, landscape bool) ([]byte, error) {
	return convert.WordToPdfAsFile(urlorFile, landscape)
}

func WordToPdfAsUrl(urlorFile interface{}, landscape bool) (string, error) {
	return convert.WordToPdfAsUrl(urlorFile, landscape)
}

// Validation
func VatValidation(vat string) (bool, error) {
	return validation.VatValidation(vat)
}

// IN and UK postcode
func PostCodeValidation(postcode string, checkType validation.CheckType) (bool, error) {
	return validation.PostCodeValidation(postcode, checkType)
}

// Email DNS and Acadamic
func EmailValidation(email string, checkType validation.CheckType) (bool, error) {
	return validation.EmailValidation(email, checkType)
}

// Extraction
func TextFromWebpage(url string) (string, error) {
	return extract.TextFromWebpage(url)
}

func TextFromWord(urlorFile interface{}) (string, error) {
	return extract.TextFromWord(urlorFile)
}

func TextFromPdf(urlorFile interface{}) (string, error) {
	return extract.TextFromPdf(urlorFile)
}

func ImageMetadata(urlorFile interface{}) (map[string]interface{}, error) {
	return extract.ImageMetadata(urlorFile)
}

func FuzzySearch(source string, target string, unicodeNormalize bool) ([]string, error) {
	return search.FuzzySearch(source, target, unicodeNormalize)
}

// unarchiver
func UnArchiveAsURL(conv interface{}) ([]string, error) {
	return extract.UnArchiveAsURL(conv)
}

func SecureUnArchiveAsURL(password string, conv interface{}) ([]string, error) {
	return extract.SecureUnArchiveAsURL(password, conv)
}

// Data getting Information
func Country(countrycode string) (h.Country, error) {
	return data.Country(countrycode)
}

func CountryList() ([]h.Country, error) {
	return data.CountryList()
}

func CurrencyConv(currency h.CurrencyConverter) (float64, error) {
	return data.CurrencyConv(currency)
}

func CurrencyList() ([]h.Currency, error) {
	return data.CurrencyList()
}

func Timezones() ([]h.Timezone, error) {
	return data.Timezones()
}

// Generate Charts, bar code,qr code and pdf
func BarChartAsFile(charts h.Chart) ([]byte, error) {
	return generate.BarChartAsFile(charts)
}

func BarChartAsURL(charts h.Chart) (string, error) {
	return generate.BarChartAsURL(charts)
}

func PieChartAsFile(charts h.Chart) ([]byte, error) {
	return generate.PieChartAsFile(charts)
}

func PieChartAsURL(charts h.Chart) (string, error) {
	return generate.PieChartAsURL(charts)
}

func StackedBarChartAsFile(charts h.StackedBarChart) ([]byte, error) {
	return generate.StackedBarChartAsFile(charts)
}

func StackedBarChartAsURL(charts h.StackedBarChart) (string, error) {
	return generate.StackedBarChartAsURL(charts)
}

func BarcodeAsFile(content string) ([]byte, error) {
	return generate.BarcodeAsFile(content)
}
func BarcodeAsURL(content string) (string, error) {
	return generate.BarcodeAsURL(content)
}

func QrcodeAsFile(content string) ([]byte, error) {
	return generate.QrcodeAsFile(content)
}

func QrcodeAsURL(content string) (string, error) {
	return generate.QrcodeAsURL(content)
}

func HtmlContentToPdfFile(content string) ([]byte, error) {
	return generate.HtmlContentToPdfFile(content)
}

func HtmlContentToPdfURL(content string) (string, error) {
	return generate.HtmlContentToPdfAsUrl(content)
}

func WebpageToPdfFile(url string, margin float64) ([]byte, error) {
	return generate.WebpageToPdfFile(url, margin)
}

func WebpageToPdfURL(url string, margin float64) (string, error) {
	return generate.WebpageToPdfURL(url, margin)
}

func ScreenshotWebpageAsFile(url string, quality int, delay int) ([]byte, error) {
	return generate.ScreenshotWebpageAsFile(url, quality, delay)
}
func ScreenshotWebpageAsURL(url string, quality int, delay int) (string, error) {
	return generate.ScreenshotWebpageAsURL(url, quality, delay)
}

func IcalGeneratorAsFile(calendar h.CalenderEvent) ([]byte, error) {
	return generate.IcalGeneratorAsFile(calendar)
}

func IcalGeneratorAsURL(calendar h.CalenderEvent) (string, error) {
	return generate.IcalGeneratorAsURL(calendar)
}

// Archiver and secure archiver
func ArchiveAsFile(conv ...interface{}) ([]byte, error) {
	return generate.ArchiveAsFile(conv)
}

func ArchiveAsURL(conv ...interface{}) (string, error) {
	return generate.ArchiveAsURL(conv)
}

func SecureArchiveFile(password string, conv ...interface{}) ([]byte, error) {
	return generate.SecureArchiveFile(password, conv)
}

func SecureArchiveAsURL(password string, conv ...interface{}) (string, error) {
	return generate.SecureArchiveAsURL(password, conv)
}

// image processing
func CompressImageAsFile(conv interface{}, quality int) ([]byte, error) {
	return imageprocessor.CompressImageAsFile(conv, quality)
}

func CompressImageAsURL(conv interface{}, quality int) (string, error) {
	return imageprocessor.CompressImageAsURL(conv, quality)
}

func CropImageAsFile(conv interface{}, width int, height int) ([]byte, error) {
	return imageprocessor.CropImageAsFile(conv, width, height)
}

func CropImageAsURL(conv interface{}, width int, height int) (string, error) {
	return imageprocessor.CropImageAsURL(conv, width, height)
}

func ResizeImageAsFile(conv interface{}, width int, height int) ([]byte, error) {
	return imageprocessor.ResizeImageAsFile(conv, width, height)
}

func ResizeImageAsURL(conv interface{}, width int, height int) (string, error) {
	return imageprocessor.ResizeImageAsURL(conv, width, height)
}

func WatermarkImageAsFile(conv h.Watermark, width int, height int) ([]byte, error) {
	return imageprocessor.WatermarkImageAsFile(conv, width, height)
}

func WatermarkImageAsURL(conv h.Watermark, width int, height int) (string, error) {
	return imageprocessor.WatermarkImageAsURL(conv, width, height)
}

func ThumnailsImageAsFile(conv interface{}, width int, height int) ([]byte, error) {
	return generate.ThumnailsImageAsFile(conv, width, height)
}

func ThumnailsImageAsURL(conv interface{}, width int, height int) (string, error) {
	return generate.ThumnailsImageAsURL(conv, width, height)
}
