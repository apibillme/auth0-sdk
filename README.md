# Auth0-SDK for Go (golang)

[![Go Report](https://goreportcard.com/badge/github.com/apibillme/auth0-sdk)](https://goreportcard.com/report/github.com/apibillme/auth0-sdk) [![GolangCI](https://golangci.com/badges/github.com/apibillme/auth0-sdk.svg)](https://golangci.com/r/github.com/apibillme/auth0-sdk) [![Travis](https://travis-ci.org/apibillme/auth0-sdk.svg?branch=master)](https://travis-ci.org/apibillme/auth0-sdk#) [![codecov](https://codecov.io/gh/apibillme/auth0-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/apibillme/auth0-sdk) ![License](https://img.shields.io/github/license/mashape/apistatus.svg) ![Maintenance](https://img.shields.io/maintenance/yes/2018.svg) [![GoDoc](https://godoc.org/github.com/apibillme/auth0-sdk?status.svg)](https://godoc.org/github.com/apibillme/auth0-sdk)


## Features:
* Full Auth0-SDK for Go (golang)
* Uses [restly](https://github.com/apibillme/restly) for API calls
* Simple to understand, fix issues (if any), and add-on
* Only about 700 LOC

```bash
go get github.com/apibillme/auth0-sdk
```

## Example

```go
func main() {
    auth0Domain := `example.auth0.com`
    clientID := `X1V34JKDJEOMCNENR`
    clientSecret := `secret`

    sdk, err := auth0sdk.New(auth0Domain, clientID, clientSecret)
    if err != nil {
        log.Panic("error!")
    }
    configs, err := sdk.GetRulesConfigs()
    if err != nil {
        log.Panic(err.Error())
    }
    for _, config := range configs {
        name := config.Get("name").String()
        log.Println(name)
    }
}
```
