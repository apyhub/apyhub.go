package helper

const (
	// file convertion
	CsvUrltoExcelFile  = "https://api.apyhub.com/convert/csv-url/excel-file"
	CsvUrltoExcelUrl   = "https://api.apyhub.com/convert/csv-url/excel-url"
	CsvFiletoExcelFile = "https://api.apyhub.com/convert/csv-file/excel-file"
	CsvFiletoExcelUrl  = "https://api.apyhub.com/convert/csv-file/excel-url"

	HtmlUrltoPdfFile  = "https://api.apyhub.com/convert/html-url/pdf-file"
	HtmlUrltoPdfUrl   = "https://api.apyhub.com/convert/html-url/pdf-url"
	HtmlFiletoPdfFile = "https://api.apyhub.com/convert/html-file/pdf-file"
	HtmlFiletoPdfUrl  = "https://api.apyhub.com/convert/html-file/pdf-url"

	ImageUrltoPdfFile = "https://api.apyhub.com/convert/image-url/pdf-file"
	ImageUrltoPdfUrl  = "https://api.apyhub.com/convert/image-url/pdf-url"
	ImagFiletoPdfFile = "https://api.apyhub.com/convert/image-file/pdf-file"
	ImageFiletoPdfUrl = "https://api.apyhub.com/convert/image-file/pdf-url"

	MdUrltoHtmlFile  = "https://api.apyhub.com/convert/md-url/html-file"
	MdUrltoHtmlUrl   = "https://api.apyhub.com/convert/md-url/html-url"
	MdFiletoHtmlFile = "https://api.apyhub.com/convert/md-file/html-file"
	MdFiletoHtmlUrl  = "https://api.apyhub.com/convert/md-file/html-url"

	PresentationUrltoPdfFile  = "https://api.apyhub.com/convert/presentation-url/pdf-file"
	PresentationUrltoPdfUrl   = "https://api.apyhub.com/convert/presentation-url/pdf-url"
	PresentationFiletoPdfFile = "https://api.apyhub.com/convert/presentation-file/pdf-file"
	PresentationFiletoPdfUrl  = "https://api.apyhub.com/convert/presentation-file/pdf-url"

	SpreadsheetUrltoPdfFile  = "https://api.apyhub.com/convert/spreadsheet-url/pdf-file"
	SpreadsheetUrltoPdfUrl   = "https://api.apyhub.com/convert/spreadsheet-url/pdf-url"
	SpreadsheetFiletoPdfFile = "https://api.apyhub.com/convert/spreadsheet-file/pdf-file"
	SpreadsheetFiletoPdfUrl  = "https://api.apyhub.com/convert/spreadsheet-file/pdf-url"

	WordUrltoPdfFile  = "https://api.apyhub.com/convert/word-url/pdf-file"
	WordUrltoPdfUrl   = "https://api.apyhub.com/convert/word-url/pdf-url"
	WordFiletoPdfFile = "https://api.apyhub.com/convert/word-file/pdf-file"
	WordFiletoPdfUrl  = "https://api.apyhub.com/convert/word-file/pdf-url"

	// image processing
	CompressImageUrlToFile  = "https://api.apyhub.com/processor/image/compress/url/file"
	CompressImageUrlToUrl   = "https://api.apyhub.com/processor/image/compress/file-urls"
	CompressImageFileToFile = "https://api.apyhub.com/processor/image/compress/file"
	CompressImageFileToUrl  = "https://api.apyhub.com/processor/image/compress/file/url"

	WaterMarkImageUrlToFile  = "https://api.apyhub.com/processor/image/watermark/url/file"
	WaterMarkImageUrlToUrl   = "https://api.apyhub.com/processor/image/watermark/file-urls"
	WaterMarkImageFileToFile = "https://api.apyhub.com/processor/image/watermark/file"
	WaterMarkImageFileToUrl  = "https://api.apyhub.com/processor/image/watermark/file/url"

	CropImageUrlToFile  = "https://api.apyhub.com/processor/image/crop/url/file"
	CropImageUrlToUrl   = "https://api.apyhub.com/processor/image/crop/file-urls"
	CropImageFileToFile = "https://api.apyhub.com/processor/image/crop/file"
	CropImageFileToUrl  = "https://api.apyhub.com/processor/image/crop/file/url"

	ResizeImageUrlToFile  = "https://api.apyhub.com/processor/image/resize/url/file"
	ResizeImageUrlToUrl   = "https://api.apyhub.com/processor/image/resize/file-urls"
	ResizeImageFileToFile = "https://api.apyhub.com/processor/image/resize/file"
	ResizeImageFileToUrl  = "https://api.apyhub.com/processor/image/resize/file/urls"

	ThumnailsImageUrlToFile  = "https://api.apyhub.com/generate/image/thumbnail/url/file"
	ThumnailsImageUrlToUrl   = "https://api.apyhub.com/generate/image/thumbnail/file-urls"
	ThumnailsImageFileToFile = "https://api.apyhub.com/generate/image/thumbnail/file"
	ThumnailsImageFileToUrl  = "https://api.apyhub.com/generate/image/thumbnail/file/url"

	// validation
	VatValidate        = "https://api.apyhub.com/validate/vat"
	INPostcodeValidate = "https://api.apyhub.com/validate/postcodes/in"
	UKPostcodeValidate = "https://api.apyhub.com/validate/postcodes/uk"
	EmailDNS           = "https://api.apyhub.com/validate/email/dns"
	EmailAcadamic      = "https://api.apyhub.com/validate/email/academic"

	// file-extraction
	TextFromWebpage       = "https://api.apyhub.com/extract/text/webpage"
	TextFromWordUrl       = "https://api.apyhub.com/extract/text/word-url"
	TextFromWordFile      = "https://api.apyhub.com/extract/text/word-file"
	TextFromPdfUrl        = "https://api.apyhub.com/extract/text/pdf-url"
	TextFromPdfFile       = "https://api.apyhub.com/extract/text/pdf-file"
	ImageMetadataFromFile = "https://api.apyhub.com/processor/image/metadata/file"
	ImageMetadataFromUrl  = "https://api.apyhub.com/processor/image/metadata/file-urls"
	FuzzySearch           = "https://api.apyhub.com/search/text/fuzzy"

	// data
	CountryInfo    = "https://api.apyhub.com/data/info/country"
	CountryList    = "https://api.apyhub.com/data/dictionary/country"
	CurrenciesList = "https://api.apyhub.com/data/dictionary/currency"
	CurrencyConv   = "https://api.apyhub.com/data/convert/currency"
	Timezones      = "https://api.apyhub.com/data/dictionary/timezone"

	// generate
	BarChartFile = "https://api.apyhub.com/generate/charts/bar/file"
	BarChartUrl  = "https://api.apyhub.com/generate/charts/bar/url"

	PieChartFile = "https://api.apyhub.com/generate/charts/pie/file"
	PieChartUrl  = "https://api.apyhub.com/generate/charts/pie/url"

	StackedBarChartFile = "https://api.apyhub.com/generate/charts/stacked/file"
	StackedBarChartUrl  = "https://api.apyhub.com/generate/charts/stacked/url"

	BarcodeAsFile = "https://api.apyhub.com/generate/barcode/file"
	BarcodeAsURL  = "https://api.apyhub.com/generate/barcode/url"

	QrcodeAsFile = "https://api.apyhub.com/generate/qr-code/file"
	QrcodeAsURL  = "https://api.apyhub.com/generate/qr-code/url"

	HtmlContentToPdfFile = "https://api.apyhub.com/generate/html-content/pdf-file"
	HtmlContentToPdfURL  = "https://api.apyhub.com/generate/html-content/pdf-url"

	WebpageToPdfFile = "https://api.apyhub.com/generate/webpage/pdf-file"
	WebpageToPdfURL  = "https://api.apyhub.com/generate/webpage/pdf-url"

	ScreenshotWebpageAsFile = "https://api.apyhub.com/generate/screenshot/webpage/image-file"
	ScreenshotWebpageAsURL  = "https://api.apyhub.com/generate/screenshot/webpage/image-url"

	ArchiveAsFile     = "https://api.apyhub.com/generate/archive/file-urls/archive-file"
	ArchiveAsURL      = "https://api.apyhub.com/generate/archive/file-urls/archive-url"
	SecureArchiveFile = "https://api.apyhub.com/generate/secure-archive/file-urls/archive-file"
	SecureArchiveURL  = "https://api.apyhub.com/generate/secure-archive/file-urls/archive-url"

	IcalAsFile = "https://api.apyhub.com/generate/ical/file"
	IcalAsURL  = "https://api.apyhub.com/generate/ical/url"
)
