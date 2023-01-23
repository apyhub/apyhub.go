[apyhub](../README.md) / search

# Namespace: search

## Table of contents

### Functions

- [fuzzy](search.md#fuzzy)

## Functions

### fuzzy

Fuzzy search. 

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

matchedText, err := apyhub.FuzzySearch("taching", "teaching is the best qualifying compare another profession", true)

fmt.Println(matchedText) //[teaching]
```

**`Link`**

https://apyhub.com/utility/search-fuzzy-text

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `source` | `string` | The source text. |
| `target` | `string` | The target text. |
| `unicodeNormalized` | `boolean` | Whether the text is unicode normalized. Default is false. |
| `output` | `[]string` | Matched text |

#### Returns

It's return the slice of string(matched text).

#### Defined in

[search/fuzzySearch.go:11](https://github.com/apyhub/apyhub.go/blob/main/search/fuzzySearch.go#L11)