<div align="center">
  <img src="https://apyhub.com/logo.svg" alt="Alt text" width="256px" height="256px">

# ApyHub for Golang
---
[![Reference](https://img.shields.io/badge/api-reference-blue.svg)](https://apyhub.com/)
[![Chat](https://camo.githubusercontent.com/75ee4c902e099cde950b194d37afca7ba236cc0892cf7db0804b85a46e9c4bc6/68747470733a2f2f696d672e736869656c64732e696f2f646973636f72642f393938393639313139363537343237303234)](https://discord.com/invite/JZagJJcw6F)


</div>

## About 

Welcome to the ApyHub SDK for Golang! This library provides easy access to the ApyHub APIs through a set of simple and powerful functions.

Our underlying philosophy is to not bloat applications, which is why we have made sure to keep the package as minimal as possible. With this SDK, you can easily perform tasks such as file conversion, text and metadata extraction, document and image generation, image processing, and access to data lists and currency conversion. As our platform grows and develops new APIs, we will be adding more functions to this library.

To learn more about ApyHub, visit our website at **[apyhub.com](https://www.apyhub.com/)** or join our Discord server at **[discord.com/invite/apyhub](https://discord.gg/JZagJJcw6F)**.

## **API Documentation**

Our SDK provides access to the following API categories:

- [Convert API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/convert.md)
- [Data API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/data.md)
- [Extract API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/extract.md)
- [Generate API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/generate.md)
- [Image Processing API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/imageProcessor.md)
- [Validate API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/validate.md)
- [Search API](https://github.com/apyhub/apyhub.go/tree/main/docs/modules/search.md)

Each of these categories contains multiple functions for interacting with the corresponding API. For more examples and detailed API documentation, see the links above.

## **Installation**

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

```bash
go get github.com/apyhub/apyhub.go
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```bash
go get -u github.com/apyhub/apyhub.go
```

## **Usage**

To use the library, you will need to initialize it with your ApyHub API token.

#### Token

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

apyhub.InitApyHub("YOUR_APY_TOKEN")
```

## **Example**

Once the library is initialized, you can import and use the various functions and methods provided by the library. For example, to access the data list of countries, you can do the following:

```go
import (
   apyhub "github.com/apyhub/apyhub.go/data"
)

 country,err:=data.Country("countrycode")
```

For more examples, see the **[API documentation](https://github.com/apyhub/apyhub.go/tree/main/src#README.md)**.

## **Contributing**

We welcome contributions to this project. If you have an idea for a new feature or improvement, please open an issue to discuss it. If you'd like to contribute code, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and commit them.
4. Push your branch to your fork on GitHub.
5. Open a pull request from your branch to the **`main`** branch of this repository.