[apyhub](../README.md) / data

# Namespace: data

## Table of contents

### Functions

- [countries](data.md#countryList)
- [country](data.md#country)
- [currencyConverter](data.md#currencyconverter)
- [currencyList](data.md#currencylist)
- [timezones](data.md#timezones)

## Functions

### countryList

Retrieves a list of countries and their associated data.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

country,err :=apyhub.CountryList()
```

**`Link`**

https://apyhub.com/utility/data-lists-country

#### Returns

```[{ "calling_codes": []number ; "cca3": "string" ; "emoji": "string" ; "key": "string" ; "value": "string"  }]```

- A output returns an array of country objects, each containing calling codes, cca3 code, emoji, key, and value properties.

#### Defined in

[data/countryList.go:7](https://github.com/apyhub/apyhub.go/blob/main/data/countryList.go#L7)

___

### country

Retrieves information about a specific country.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

country,err :=apyhub.Country("IN")
```

**`Link`**

https://apyhub.com/utility/data-info-country

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `country` | `string` | The name or cca3 code of the country. |

#### Returns

 ```{ "calling_code": []number ; "cca3": "string" ; "emoji": "string" ; "key": "string" ; "subdivision": [{ "code": "string" ; "name": "string"  }] ; "value": "string" }```

- A output returns an object containing the country's value, key, cca3 code, emoji, calling codes, and an array of subdivisions, each with a code and name.

#### Defined in

[data/country.go:9](https://github.com/apyhub/apyhub.go/blob/main/data/country.go#L9)

___

### currencyConverter

Converts an amount in one currency to another.

**`Example`**

```go

import (
   apyhub "github.com/apyhub/apyhub.go"
   h      "github.com/apyhub/helper"
)

currency:=h.CurrencyConverter{
      Source: "inr",
      Target: "usd",
      Date: "2023-01-20",
}

currencyrate,err :=apyhub.CurrencyConv(currency)
```

**`Link`**

https://apyhub.com/utility/currency-conversion

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `date` | `string` | The date for the conversion (in the format "YYYY-MM-DD"). If not provided, the current date is used. |
| `source` | `string` | The source currency. |
| `target` | `string` | The target currency. |

#### Returns

```{ data: number }```

- A output returns an object containing the converted amount.

#### Defined in

[data/currencyConverter.go:10](https://github.com/apyhub/apyhub.go/blob/main/data/currencyConverter.go#L10)

___

### currencyList

Retrieves a list of currencies and their associated data.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

currencylist,err:=apyhub.CurrencyList()
```

**`Link`**

https://apyhub.com/utility/data-lists-currency

#### Returns

    [{ "emoji": "string" ; "key": "string" ; "value": "string"  }]

- A output returns an array of currency objects, each containing emoji, key, and value properties.

#### Defined in

[data/currencyList.go:7](https://github.com/apyhub/apyhub.go/blob/main/data/currencyList.go#L7)

___

### timezones

Retrieves a list of timezones and their associated data.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

timezone,err :=apyhub.Timezones()
```

**`Link`**

https://apyhub.com/utility/data-lists-timezone

#### Returns

```{ "data": [{ "abbreviation": []string ; "key": "string" ; "utc_time": "string" ; "value": "string"  }]  }```

- A output returns an object containing an array of timezone objects, each with a key, value, abbreviation, and utc_time property.

#### Defined in

[data/timezone.go:7](https://github.com/apyhub/apyhub.go/blob/main/data/timezone.go#L7)