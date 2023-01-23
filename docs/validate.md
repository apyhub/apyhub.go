[apyhub](../README.md) / validate

# Namespace: validate

## Table of contents

### Functions

- [email](validate.md#email)
- [postcode](validate.md#postcode)
- [vat](validate.md#vat)

## Functions

### email

Validates an email address.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

// Validate DNS
dns, err := apyhub.EmailValidation("hello@apyhub.com", validation.CheckType(validation.DNS))

// Validate Acadamic Emails
acadamic, err := apyhub.EmailValidation("google.edu.in", validation.CheckType(validation.Acadamic))
```

**`Link`**

https://apyhub.com/utility/validator-dns-email

**`Link`**

https://apyhub.com/utility/validator-academic-email

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `checkType` | ``"dns"`` \| ``"academic"`` | The type of validation to perform. Can be either "dns" or "academic". |
| `email` | `string` | The email address to validate. |
| `output` | `bool` | It's return the email is valid dns or acadamic. |


#### Returns
It's return boolean of the email dns or acadamic.

#### Defined in

[validate/emailValidate.go:14](https://github.com/apyhub/apyhub.go/blob/main/validation/emailValidate.go#L14)

___

### postcode

Validates a postcode.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

INPostCode, err := apyhub.PostCodeValidation("520010", validation.CheckType(validation.IN))


UKPostCode, err := apyhub.PostCodeValidation("AB11 5QN", validation.CheckType(validation.UK))
```

**`Link`**

https://apyhub.com/utility/data-postcodes-india

**`Link`**

https://apyhub.com/utility/data-postcodes-uk

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `countryCode` | ``"IN"`` \| ``"UK"`` | The country code of the postcode. Can be either "IN" or "UK". |
| `postcode` | `string` | The postcode to validate. |
| `output` | `bool` | Bool of postcode validation |

#### Returns
It's return valid postcode or not in boolean type.

#### Defined in

[validate/postCode.go:12](https://github.com/apyhub/apyhub.go/blob/main/validation/postCode.go#L12)

___

### vat

Validates a VAT number.

**`Example`**

```go
import (
   apyhub "github.com/apyhub/apyhub.go"
)

vat, err := apyhub.VatValidation("RO19")
```

**`Link`**

https://apyhub.com/utility/vat-number-validator

#### Parameters

| Name | Type | Description |
| :------ | :----- | :----- |
| `vat` | `string` | It's taking a vat number. |
| `output` | `bool` | Bool of postcode validation |

#### Returns
It's return bool value of vat is valid.

#### Defined in

[validate/vat.go:7](https://github.com/apyhub/apyhub.go/blob/main/validation/vat.go#L7)