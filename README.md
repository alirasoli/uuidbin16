# UUIDBIN16
This package adds support for binary(16) type to sql in golang to store UUID as binary(16).

## How to use
Add the package to your project:
```bash
go get github.com/alirasoli/uuidbin16
```

Convert uuid strings to uuidbin16 type:
```go
var uuidbin uuidbin16.StringToBin16(uuid)
```
And then you can define your database models with this type and it handles sql part for you.