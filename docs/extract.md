[apyhub](../README.md) / extract

# Namespace: extract

## Table of contents

### Functions

- [imageMetadata](extract.md#imagemetadata)
- [textFromPdf](extract.md#textfrompdf)
- [textFromWebpage](extract.md#textfromwebpage)
- [textFromWord](extract.md#textfromword)
- [unarchive](generate.md#unarchive)


## Functions

### imageMetadata

Extracts metadata from an image.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// input is a file
File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

data,err :=apyhub.ImageMetadata(File)

// input is a url
imageUrl :="https://assets.apyhub.com/samples/sample.jpg"

data,err :=apyhub.ImageMetadata(imageUrl)
```

**`Link`**

https://apyhub.com/utility/image-processor-extract-metadata

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `image` | `string` \| `Buffer` | The image file or URL. |
| `output` | `map[string]interface{}` | It's returns a metadata for the image |
#### Returns

```{ data: map[string]interface{} }```

A output returns an object to the metadata for the image.

#### Defined in

[extract/ImageMetadata.go:11](https://github.com/apyhub/apyhub.go/blob/main/extract/ImageMetadata.go#L11)

___

### textFromPdf

Extracts text from a PDF file.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// input is a file
File,err :=os.Open("any pdf file")
if err!=nil {
    log.Fatal(err)
}

data,err:=apyhub.TextFromPdf(File)

// input is a url 
fileUrl :="https://assets.apyhub.com/samples/sample.pdf"

data,err:=apyhub.TextFromPdf(fileUrl)
```

**`Link`**

https://apyhub.com/utility/extractor-pdf-text

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `pdf` | `string` \| `Buffer` | The PDF file or URL. |
| `output` | `string` | output return as a string. |

#### Returns

```{ data: string }```

A output returns to the text for the PDF.

#### Defined in

[extract/textFromPdf.go:11](https://github.com/apyhub/apyhub.go/blob/main/extract/textFromPdf.go#L11)

___

### textFromWebpage

Extracts text from a webpage.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

data,err :=apyhub.TextFromWebpage("any url")
```

**`Link`**

https://apyhub.com/utility/extractor-webpage-text

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `url` | `string` | The URL of the webpage. |
| `output` | `string` | output return as a string. |

#### Returns

```{ data: string }```

A output returns to the text for the webpage.

#### Defined in

[extract/textFromWebpage.go:9](https://github.com/apyhub/apyhub.go/blob/main/extract/textFromWebpage.go#L9)

___

### textFromWord

Extracts text from a Word document.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// input is a file
File,err :=os.Open("any word file")
if err!=nil {
    log.Fatal(err)
}

data,err := apyhub.TextFromWord(File)

// input is a url
fileUrl :="https://assets.apyhub.com/samples/sample.docx"

data,err := apyhub.TextFromWord(fileUrl)
```

**`Link`**

https://apyhub.com/utility/extractor-word-text

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `word` | `string` \| `Buffer` | The Word document file or URL. |

#### Returns

```{ data: string }```

A output returns to the text for the Word document.

#### Defined in

[extract/textFromWord.go:11](https://github.com/apyhub/apyhub.go/blob/main/extract/textFromWord.go#L11)


### unarchive

Unarchive a unsecure zip file or fileurl and return as list of url 

**`Example 1: unsecure file-unarchiver a zip file and return as list of AWS presigned url`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

File, err := os.Open("unsecure zip file")
if err != nil {
    log.Fatal(err)
}

sliceStr, err := apyhub.UnArchiveAsURL(File)
```
**`Example 2: unsecure file-unarchiver a zip file url and return as list of AWS presigned url`**

```go

zipUrl:="unsecure zip file url"

sliceStr, err := apyhub.UnArchiveAsURL(zipUrl)

```
#### Defined in
[extract/unarchive.go:11](https://github.com/apyhub/apyhub.go/blob/main/generate/archive.go#L11)


**`Link`**

https://apyhub.com/utility/extract-file-unarchive


Unarchive a secure zip file or fileurl and return as list of url 

**`Example 3: secure file-unarchiver a zip file and return as list of AWS presigned url`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

File, err := os.Open("secure zip file")
if err != nil {
    log.Fatal(err)
}

sliceStr, err := apyhub.SecureUnArchiveAsURL("password",File)
```
**`Example 4: secure file-unarchiver a zip file url and return as list of AWS presigned url`**

```go

zipUrl:="secure zip file url"

sliceStr, err := apyhub.SecureUnArchiveAsURL("password",zipUrl)

```
#### Defined in
[generate/archive.go:36](https://github.com/apyhub/apyhub.go/blob/main/generate/archive.go#L36)

**`Link`**

https://apyhub.com/utility/extract-file-secure-unarchive

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `file | url` | The list of files or URLs to archive. |
| `output` | `[]string` | Format for the response |
| `password` | `string` | The password to use for secure unarchiving. |

#### Returns

- The data for the output as a slice of string as presigned Url.
