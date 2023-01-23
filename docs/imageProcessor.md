[apyhub](../README.md) / imageProcessor

# Namespace: imageProcessor

## Table of contents

### Functions

- [compress](imageProcessor.md#compress)
- [crop](imageProcessor.md#crop)
- [resize](imageProcessor.md#resize)
- [watermark](imageProcessor.md#watermark)

## Functions

### compress

Compresses an image file or URL.

**`Example 1 : compress a local image`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// input is a file
File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.CompressImageAsFile(File, 100)
```
**`Example 2 : compress a image from url`**

```go

// input is a url
imageUrl:="image url"

byt,err :=apyhub.CompressImageAsFile(imageUrl, 100)
```
#### Defined in

[imageprocessor/compress.go:14](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/compress.go#L14)

**`Example 3 : compress a local image and return AWS presigned url`**
```go
// input is a file
File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

url,err:=apyhub.CompressImageAsURL(File, 100)
```
**`Example 4 : compress a image from url and return AWS presigned url`**
```go
// input is a url
imageUrl:="image url"
url,err:=apyhub.CompressImageAsURL(imageUrl, 100)
```

#### Defined in

[imageprocessor/compress.go:42](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/compress.go#L42)

**`Link`**

https://apyhub.com/utility/image-processor-compress

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` \| `Buffer` | The image file or URL to compress. |
| `output` | `[]byte`\| `string` | The format for the response.   |
| `quality` | `int` | The desired quality of the compressed image. |


#### Returns

 - The data for the output file as a slice of byte or Url as a string.
___

### crop

Crops an image file or URL.

**`Example 1 : crop a local image`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.CropImageAsFile(File, 350, 250)
```

**`Example 2 : crop a image from url`**

```go

imageUrl:="image url"

byt,err :=apyhub.CropImageAsFile(imageUrl, 350, 250)
```

#### Defined in

[imageprocessor/crop.go:14](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/crop.go#L14)

**`Example 3 : crop a local image and return AWS presigned url`**

```go
File,err :=os.Open("any image")
if err!=nil {
    log.Fatal(err)
}

url,err:=apyhub.CropImageAsURL(File, 350, 250)
```
**`Example 4 : crop a image from url and return AWS presigned url`**

```go
imageUrl:="image url"

url,err:=apyhub.CropImageAsURL(imageUrl, 350, 250)
```

#### Defined in

[imageprocessor/crop.go:42](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/crop.go#L42)

**`Link`**

https://apyhub.com/utility/image-processor-crop

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` \| `Buffer` | The image file or URL to crop. |
| `height` | `int` | The desired height of the cropped image. |
| `width` | `int` | The desired width of the cropped image. |
| `output` | `[]byte`\| `string` | The format for the response.   |


#### Returns

-   The data for the output file as a slice of byte or Url as a string.

___

### resize

Resizes an image file or URL.

**`Example 1 : resize a local image`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// input is a file
File,err :=os.Open("any file")
if err!=nil {
    log.Fatal(err)
}

byt,err :=apyhub.ResizeImageAsFile(File, 350, 250)
```
**`Example 2 : resize a image from url`**

```go
imageUrl:="image url"

byt,err :=apyhub.ResizeImageAsFile(imageUrl, 350, 250)
if err!=nil {
    log.Fatal(err)
}
```

#### Defined in

[imageprocessor/resize.go:14](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/resize.go#L14)


**`Example 3 : resize a local image and return AWS presigned url`**

```go
// input is a file
File,err :=os.Open("any file")
if err!=nil {
    log.Fatal(err)
}

url,err:=apyhub.ResizeImageAsURL(File, 350, 250)
```
**`Example 4 : resize a image from url and return AWS presigned url`**

```go
imageUrl:="image url"

url,err:=apyhub.ResizeImageAsURL(imageUrl, 350, 250)
```

#### Defined in

[imageprocessor/resize.go:42](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/resize.go#L42)


**`Link`**

https://apyhub.com/utility/image-processor-resize

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` \| `Buffer` | The image file or URL to resize. |
| `width` | `int` | The desired width of the resized image. |
| `height` | `int` | The desired height of the resized image. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

-   The data for the output file as a slice of byte or Url as a string.

___

### watermark

Adds a watermark to an image.

**`Example 1 : watermark to a local image`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

image,err :=os.Open("source image")
if err!=nil {
    log.Fatal(err)
}

watermarkimage,err :=os.Open("watermark image")
if err!=nil {
    log.Fatal(err)
}

conv := h.Watermark{
    Image:     image,
    Watermark: watermarkimage,
}

byt,err :=apyhub.WatermarkImageAsFile(conv, 350, 250)
```

**`Example 2 : watermark to a image from url`**

```go

imageUrl:="image url"
watermarkimageUrl:="image url"

conv = h.Watermark{
    Image:     imageUrl,
    Watermark: watermarkimageUrl,
}
byt,err :=apyhub.WatermarkImageAsFile(conv, 350, 250)
```

#### Defined in

[imageprocessor/watermark.go:15](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/watermark.go#L15)


**`Example 3 : watermark to a local image and return AWS presigned url`**

```go

import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

image,err :=os.Open("source image")
if err!=nil {
    log.Fatal(err)
}

watermarkimage,err :=os.Open("watermark image")
if err!=nil {
    log.Fatal(err)
}

conv := h.Watermark{
    Image:     image,
    Watermark: watermarkimage,
}

url,err:=apyhub.WatermarkImageAsURL(conv, 350, 250)
```
**`Example 4 : watermark to a image from url and return AWS presigned url`**

```go
imageUrl:="image url"
watermarkimageUrl:="image url"

conv = h.Watermark{
    Image:     imageUrl,
    Watermark: watermarkimageUrl,
}

url,err:=apyhub.WatermarkImageAsURL(conv, 350, 250)
```

#### Defined in

[imageprocessor/watermark.go:54](https://github.com/apyhub/apyhub.go/blob/main/imageprocessor/watermark.go#L54)

**`Link`**

https://apyhub.com/utility/image-processor-watermark

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `input` | `string` \| `Buffer` | The input image as a file path or URL, or as a Buffer if it is a file. |
| `image` | `string` \| `Buffer` | The image as a file path or URL, or as a Buffer if it is a file. |
| `watermarkimage` | `string` \| `Buffer` | The watermark image as a file path or URL, or as a Buffer if it is a file. |
| `width` | `int` | The desired width of the resized image. |
| `height` | `int` | The desired height of the resized image. |
| `output` | `[]byte`\| `string` | The format for the response.   |

#### Returns

-   The data for the output file as a slice of byte or Url as a string.