[apyhub](../README.md) / convert

# Namespace: convert

## Table of contents

### Functions

- [csvToExcel](convert.md#csvtoexcel)
- [htmlToPdf](convert.md#htmltopdf)
- [imageToPdf](convert.md#imagetopdf)
- [markdownToHtml](convert.md#markdowntohtml)
- [presentationToPdf](convert.md#presentationtopdf)
- [spreadsheetToPdf](convert.md#spreadsheettopdf)
- [wordToPdf](convert.md#wordtopdf)

## Functions

### csvToExcel

Converts a CSV file or URL to an Excel file.

**`Example 1 : convert csv file to Excel file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any csv file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.CsvToExcelAsFile(File)
```

**`Example 2 : convert csv file url to Excel file`**
```go

// Input is a url
fileUrl :="csv file url"

byt,err :=apyhub.CsvToExcelAsFile(fileUrl)
```
#### Defined in

[convert/csvToExcel.go:13](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L13)

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
// Input is a url
fileUrl :="csv file url"

url,err :=apyhub.CsvToExcelAsUrl(fileUrl)
```
#### Defined in

[convert/csvToExcel.go:38](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L38)

**`Link`**

https://apyhub.com/utility/converter-csv-excel

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The csv  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.

___
### htmlToPdf


Converts an HTML file or URL to a PDF file.

**`Example 1 : convert html file to pdf file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any html file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.HtmlToPdfAsFile(File)
```

**`Example 2 : convert html file url to pdf file`**
```go

// Input is a url
fileUrl :="html file url"

byt,err :=apyhub.HtmlToPdfAsFile(fileUrl)
```
#### Defined in

[convert/htmlToPdf.go:14](https://github.com/apyhub/apyhub.go/blob/main/convert/htmlToPdf.go#L14)

**`Example 3 : convert html file to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any html file")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.HtmlToPdfAsUrl(File)
```
**`Example 4 : convert html file url to AWS presigned url`**
```go
// Input is a url
fileUrl :="html file url"

url,err :=apyhub.HtmlToPdfAsUrl(File)
```
#### Defined in

[convert/htmlToPdf.go:39](https://github.com/apyhub/apyhub.go/blob/main/convert/htmlToPdf.go#L39)

**`Link`**

https://apyhub.com/utility/converter-html-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The html  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.
  
___
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

**`Example 3 : convert image to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.ImageToPdfAsUrl(File)
```
**`Example 4 : convert image url to AWS presigned url`**
```go
// Input is a url
fileUrl :="image url"

url,err :=apyhub.ImageToPdfAsUrl(fileUrl)
```
#### Defined in

[convert/imageToPdf.go:38](https://github.com/apyhub/apyhub.go/blob/main/convert/imageToPdf.go#L38)

**`Link`**

https://apyhub.com/utility/converter-image-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The image  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.
  
___

### markdownToHtml


Convert markdown to HTML.

This function converts the given markdown input to HTML and returns the
result.

**`Example 1 : convert markdown to html file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any markdown file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.MarkdownToHtmlAsFile(File)
```

**`Example 2 : convert markdown file url to html file`**
```go

// Input is a url
fileUrl :="markdown file url"

byt,err :=apyhub.MarkdownToHtmlAsFile(fileUrl)
```
#### Defined in

[convert/markdownToHtml.go:13](https://github.com/apyhub/apyhub.go/blob/main/convert/markdownToHtml.go#L13)

**`Example 3 : convert markdown file to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any markdown file")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.MarkdownToHtmlAsUrl(File)
```
**`Example 4 : convert markdown file url to AWS presigned url`**
```go
// Input is a url
fileUrl :="markdown file url"

url,err :=apyhub.MarkdownToHtmlAsUrl(fileUrl)
```
#### Defined in

[convert/markdownToHtml.go:38](https://github.com/apyhub/apyhub.go/blob/main/convert/markdownToHtml.go#L38)

**`Link`**

https://apyhub.com/utility/converter-md-html

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The markdown  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.

___

### presentationToPdf


Convert a presentation to PDF.

This function converts the given presentation input to PDF and returns the
result.

**`Example 1 : convert presentation file to pdf file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any presentation file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.PresentationToPdfAsFile(File)
```

**`Example 2 : convert presentation file url to pdf file`**
```go

// Input is a url
fileUrl :="presentation file url"

byt,err :=apyhub.PresentationToPdfAsFile(fileUrl)
```
#### Defined in

[convert/presentationToPdf.go:14](https://github.com/apyhub/apyhub.go/blob/main/convert/presentationToPdf.go#L14)

**`Example 3 : convert presentation file to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any presentation file")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.PresentationToPdfAsUrl(File)
```
**`Example 4 : convert presentation file url to AWS presigned url`**
```go
// Input is a url
fileUrl :="presentation file url"

url,err :=apyhub.PresentationToPdfAsUrl(fileUrl)
```
#### Defined in

[convert/presentationToPdf.go:40](https://github.com/apyhub/apyhub.go/blob/main/convert/presentationToPdf.go#L40)

**`Link`**

https://apyhub.com/utility/converter-presentation-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The presentation  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.

___

### spreadsheetToPdf


Converts a spreadsheet file or URL to an pdf file.

**`Example 1 : convert spreadsheet file to pdf file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any spreadsheet file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.SpreadsheetToPdfAsFile(File)
```

**`Example 2 : convert spreadsheet file url to pdf file`**
```go

// Input is a url
fileUrl :="spreadsheet file url"

byt,err :=apyhub.SpreadsheetToPdfAsFile(fileUrl)
```
#### Defined in

[convert/spreadsheetToPdf.go:14](https://github.com/apyhub/apyhub.go/blob/main/convert/spreadsheetToPdf.go#L14)

**`Example 3 : convert spreadsheet file to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any spreadsheet file")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.SpreadsheetToPdfAsUrl(File)
```
**`Example 4 : convert spreadsheet file url to AWS presigned url`**
```go
// Input is a url
fileUrl :="spreadsheet file url"

url,err :=apyhub.SpreadsheetToPdfAsUrl(fileUrl)
```
#### Defined in

[convert/spreadsheetToPdf.go:42](https://github.com/apyhub/apyhub.go/blob/main/convert/spreadsheetToPdf.go#L42)

**`Link`**

https://apyhub.com/utility/converter-spreadsheet-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The spreadsheet  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.

___

### csvToExcel


Converts a CSV file or URL to an Excel file.

**`Example 1 : convert csv file to Excel file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any csv file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.CsvToExcelAsFile(File)
```

**`Example 2 : convert csv file url to Excel file`**
```go

// Input is a url
fileUrl :="csv file url"

byt,err :=apyhub.CsvToExcelAsFile(fileUrl)
```
#### Defined in

[convert/csvToExcel.go:13](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L13)

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
// Input is a url
fileUrl :="csv file url"

url,err :=apyhub.CsvToExcelAsUrl(fileUrl)
```
#### Defined in

[convert/csvToExcel.go:38](https://github.com/apyhub/apyhub.go/blob/main/convert/csvToExcel.go#L38)

**`Link`**

https://apyhub.com/utility/converter-csv-excel

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The csv  file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.

___
### wordToPdf


Converts a Word file or URL to a PDF file.

**`Example 1 : convert word file to pdf file`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Input is a file
File,err :=os.Open("any word file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.WordToPdfAsFile(File)
```

**`Example 2 : convert word file url to pdf file`**
```go

// Input is a url
fileUrl :="word file url"

byt,err :=apyhub.WordToPdfAsFile(fileUrl)
```
#### Defined in

[convert/wordToPdf.go:14](https://github.com/apyhub/apyhub.go/blob/main/convert/wordToPdf.go#L14)

**`Example 3 : convert word file to AWS presigned url`**
```go

// Input is a file
File,err :=os.Open("any word file")
if err!=nil {
    log.Fatal(err)
}

url,err :=apyhub.WordToPdfAsUrl(File)
```
**`Example 4 : convert word file url to AWS presigned url`**
```go
// Input is a url
fileUrl :="word file url"

url,err :=apyhub.WordToPdfAsUrl(File)
```
#### Defined in

[convert/wordToPdf.go:42](https://github.com/apyhub/apyhub.go/blob/main/convert/wordToPdf.go#L42)

**`Link`**

https://apyhub.com/utility/converter-doc-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `"string"` \| `Buffer` | The word file or URL. |
| `output` | `[]byte`\| `"string"` | The format for the response.  |

#### Returns

  The data for the output file as slice of byte or Url as a string.
  
___