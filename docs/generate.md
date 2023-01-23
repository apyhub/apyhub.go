[apyhub](../README.md) / generate

# Namespace: generate

## Table of contents

### Functions

- [archive](generate.md#archive)
- [barcode](generate.md#barcode)
- [charts](generate.md#charts)
- [ical](generate.md#ical)
- [pdf](generate.md#pdf)
- [qr](generate.md#qr)
- [screenshot](generate.md#screenshot)
- [thumbnail](generate.md#thumbnail)

## Functions

### archive

Creates an archive file (zip) from a list of URLs.

**`Example of Unsecure File-archiver`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

zip:= h.Zip{
        Urls: []string{
          "any file url",
          "any file url",
          "any file url"
        }
      }

// It return the file bytes
byt, err := a.ArchiveAsFile(zip)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

zip:= h.Zip{
        Urls: []string{
          "any file url",
          "any file url",
          "any file url"
        }
      }
      
// It's return file presigned url. It's valid for 15 mins.
url, err := a.ArchiveAsURL(zip)
```

**`Example of Secure File-archiver`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

securedzip:= h.ZipSecure{
    Urls: []string{
            "any file url",
            "any file url",
            "any file url"
      },
    Password: "password",
  }

byt, err := a.SecureArchiveFile(securedzip)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

securedzip:= h.ZipSecure{
    Urls: []string{
            "any file url",
            "any file url",
            "any file url"
      },
    Password: "password",
  }

url, err := a.SecureArchiveAsURL(securedzip)
```

**`Link`**

https://apyhub.com/utility/generate-file-archive

**`Link`**

https://apyhub.com/utility/generate-secure-file-archive

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | []`string` | The list of files or URLs to archive. |
| `output` | `[]byte | string` | The desired file name for the output archive. |
| `password` | `string` | The password to use for secure archiving. |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/archive.go:11](https://github.com/apyhub/apyhub.go/blob/main/generate/archive.go#L11)

___

### barcode

Generates a barcode image from a string of up to 80 characters.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// It will return image bytes
byt,err :=a.BarcodeAsFile("Text you want to embed in barcode")
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// It will return presigned url
url,err :=a.BarcodeAsURL("Text you want to embed in barcode")
```

**`Link`**

https://apyhub.com/utility/generate-bar-code

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` & { `maxLength`: ``80``  } | The string to generate the barcode from. |
| `output` | `[]byte`\| `string`| The format for the response.    |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/barcode.go:7](https://github.com/apyhub/apyhub.go/blob/main/generate/barcode.go#L7)

___

### charts

Generates a bar chart image from data options.

▸ **barchart**

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

charts := h.Chart{
        Title: "bar chart",
        Theme: "light",
        Data: []h.Values{
            {
                Value: 10,
                Label: "label1",
            },
            {
                Value: 20,
                Label: "label2",
            },
        },
    }
// it will return image as bytes
byt,err:=a.BarChartAsFile(charts)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

charts := h.Chart{
        Title: "bar chart",
        Theme: "light",
        Data: []h.Values{
            {
                Value: 10,
                Label: "label1",
            },
            {
                Value: 20,
                Label: "label2",
            },
        },
    }

// it will return presigned url
url,err :=a.BarChartAsURL(charts)
```
**`Link`**

https://apyhub.com/utility/bar-graph

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `title` | `string` | The title of the chart. |
| `theme` | ``"light"`` \| ``"dark"`` | The theme of the chart. Valid values are "light" or "dark". |
| `data` | [{ `label`: `string` ; `value`: `float`  }] | The data to be used to generate the chart. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/charts.go:7](https://github.com/apyhub/apyhub.go/blob/main/generate/charts.go#L7)

___

▸ **piechart**

Generates a pie chart image from data options.

**`Example`**
```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

    charts := h.Chart{
        Title: "pie chart",
        Theme: "light",
        Data: []h.Values{
            {
                Value: 10,
                Label: "label1",
            },
            {
                Value: 20,
                Label: "label2",
            },
        },
    }
// it will return image as bytes
byt,err :=a.PieChartAsFile(charts)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

charts := h.Chart{
        Title: "pie chart",
        Theme: "light",
        Data: []h.Values{
            {
                Value: 10,
                Label: "label1",
            },
            {
                Value: 20,
                Label: "label2",
            },
        },
    }
// it will return presigned url
url,err :=a.PieChartAsURL(charts)
```

**`Link`**

https://apyhub.com/utility/pie-chart

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `title` | `string` | The title of the chart. |
| `theme` | ``"light"`` \| ``"dark"`` | The theme of the chart. Valid values are "light" or "dark". |
| `data` | [{ `label`: `string` ; `value`: `float`  }] | The data to be used to generate the chart. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/charts.go:19](https://github.com/apyhub/apyhub.go/blob/main/generate/charts.go#L19)

___

▸ **Stackedbar**

Generates a Stackedbar chart image from data options.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

charts := h.StackedBarChart{
    Title: "StackedBarChart",
    Theme: "dark",
    Data: []h.StackedValue{
        {
            Name: "bar1",
            Values: []h.Values{
                {
                        Label: "label1",
                        Value: 10,
                },
                {
                        Label: "label1",
                        Value: 10,
                },
            },
        },
    },
}
// it will return image as bytes
byt,err :=a.StackedBarChartAsFile(charts)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

charts := h.StackedBarChart{
    Title: "StackedBarChart",
    Theme: "dark",
    Data: []h.StackedValue{
        {
            Name: "bar1",
            Values: []h.Values{
                {
                    Label: "label1",
                    Value: 10,
                },
                {
                    Label: "label1",
                    Value: 10,
                },
            },
        },
    },
}
// it will return presigned url
url,err :=a.StackedBarChartAsURL(charts)
```

**`Link`**

https://apyhub.com/utility/stacked-graph


#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `title` | `string` | The title of the chart. |
| `theme` | ``"light"`` \| ``"dark"`` | The theme of the chart. Valid values are "light" or "dark". |
| `data` | [{ `label`: `string` ; `value`: `float`  }] | The data to be used to generate the chart. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.


#### Defined in

[generate/charts.go:31](https://github.com/apyhub/apyhub.go/blob/main/generate/charts.go#L31)

___

### ical

Generates an iCalendar file from a list of parameters.

**`Example`**

```go

import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

calender := h.CalenderEvent{
    Summary:         "John",
    Description:     "Google Meeting",
    OrganizerEmail:  "John.j@outlook.com",
    AttendeesEmails: []string{"Micheal@outlook.com", "Richard@outlook.com"},
    Location:        "india",
    TimeZone:        "asia/kolkata",
    StartTime:       "05:30",
    EndTime:         "06:30",
    MeetingDate:     "05-10-2022",
    Reccuring:       true,
    Recurrence: h.Reccurent{
        Frequency: "Monthly",
        Count:     3,
    },
}
// it will return .ics file in bytes
byt,err :=a.IcalGeneratorAsFile(calender)
```
```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

calender := h.CalenderEvent{
    Summary:         "John",
    Description:     "Google Meeting",
    OrganizerEmail:  "John.j@outlook.com",
    AttendeesEmails: []string{"Micheal@outlook.com", "Richard@outlook.com"},
    Location:        "india",
    TimeZone:        "asia/kolkata",
    StartTime:       "05:30",
    EndTime:         "06:30",
    MeetingDate:     "05-10-2022",
    Reccuring:       true,
    Recurrence: h.Reccurent{
        Frequency: "Monthly",
        Count:     3,
    },
}
// it will return presigned url of the .ics file
url,err :=a.IcalGeneratorAsURL(calender)
```

**`Link`**

https://apyhub.com/utility/generator-ical

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `summary` | `string` | The meeting summary. |
| `description` | `string` | The meeting description. |
| `organizerEmail` | `string` | The meeting organizer's email address. |
| `attendeesEmails` | `string`[] | The meeting attendees' email addresses. |
| `location` | `string` | The meeting location. |
| `timeZone` | `string` | The meeting time zone. |
| `startTime` | `string` | The meeting start time. |
| `endTime` | `string` | The meeting end time. |
| `meetingDate` | `string` | The meeting date. |
| `recurring` | `boolean` | Whether the meeting is recurring. |
| `recurrence.frequency` | `string` | - |
| `recurrence.count` | `int` | - |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/ical.go:7](https://github.com/apyhub/apyhub.go/blob/main/generate/ical.go#L7)

___

### pdf

▸ **html**

Generates a PDF file from an HTML content.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

content :="<b>If you click on me, I will disappear.</b><p> this is an sample data</p><p> added new line </p>"
// it will return pdf file as bytes
byt,err :=a.HtmlContentToPdfFile(content)
```
```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

content :="<b>If you click on me, I will disappear.</b><p> this is an sample data</p><p> added new line </p>"
// it will return presigned url
url,err :=a.HtmlContentToPdfURL(content)
```
**`Link`**

https://apyhub.com/utility/generate-html-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `content-string` | The HTML content to generate the PDF form. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

___

▸ **Webpage**

Generates a PDF file from an Webpage url.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// it will return pdf as bytes
byt,err :=apyhub.WebpageToPdfFile("any url", 0.1)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// it will presigned url
url,err :=apyhub.WebpageToPdfURL("any url", 0.1)
```

**`Link`**

https://apyhub.com/utility/generate-url-pdf

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `content-string` | The HTML content to generate the PDF form. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/webpageToPdf.go:9](https://github.com/apyhub/apyhub.go/blob/main/generate/webpageToPdf.go#L9)

___

### qr

Generates a QR code image from a string.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// It will return qr image as bytes
byt,err :=apyhub.QrcodeAsFile("your message")
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// it will return presigned url
url,err :=a.QrcodeAsURL("https://apyhub.com")
```

**`Link`**

https://apyhub.com/utility/generate-qr-code

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` | The string to generate the QR code from. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/qr.go:7](https://github.com/apyhub/apyhub.go/blob/main/generate/qrcode.go#L7)

___

### screenshot

Takes a screenshot of a webpage.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// it will return image as bytes
byt,err :=apyhub.ScreenshotWebpageAsFile("any url", 3,5)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)
// it will return presigned url
url,err :=apyhub.ScreenshotWebpageAsURL("any url", 3, 5)
```

**`Link`**

https://apyhub.com/utility/generate-webpage-screenshot

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` | The URL of the webpage to take a screenshot of. |
| `delay` | `int` | The delay, in seconds, before taking the screenshot. |
| `quality` | ``1`` \| ``2`` \| ``3`` \| ``4`` \| ``5`` | The quality of the screenshot image. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

- The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/screenshot.go:9](https://github.com/apyhub/apyhub.go/blob/main/generate/screenshot.go#L9)

___

### thumbnail

Creates a thumbnail image from a file or URL.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

File,err :=os.Open("local image")
if err!=nil {
    log.Fatal(err)
}
// it will return image as bytes
byt,err :=apyhub.ThumnailsImageAsFile(File, 350, 250)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

imageUrl:="image url"

byt,err :=apyhub.ThumnailsImageAsFile(imageUrl, 350, 250)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

File,err :=os.Open("local image")
if err!=nil {
    log.Fatal(err)
}

url,err:=apyhub.ThumnailsImageAsURL(File, 350, 250)
```

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

imageUrl:="image url"

url,err:=apyhub.ThumnailsImageAsURL(imageUrl, 350, 250)
```

**`Link`**

https://apyhub.com/utility/image-processor-thumbnail

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `params` | `Object` | The parameters for the thumbnail creation. |
| `height` | `int` | The desired height of the thumbnail. |
| `input` | `string` \| `Buffer` | The file or URL to create a thumbnail from. |
| `output?` | `string` | The desired file name for the output thumbnail. |
| `responseFormat` | ``"file"`` \| ``"url"`` | The desired response format. |
| `width` | `int` | The desired width of the thumbnail. |

#### Returns

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` \| `Buffer` | The image file or URL to resize. |
| `width` | `int` | The desired width of the resized image. |
| `height` | `int` | The desired height of the resized image. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

-   The data for the output file as a slice of byte or Url as a string.

#### Defined in

[generate/thumnails.go:14](https://github.com/apyhub/apyhub.go/blob/main/generate/thumnails.go#L14)