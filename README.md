# ipdata - Go library for ipdata.co

[![Build Status](https://travis-ci.org/cuonglm/ipdata.svg?branch=master)](https://travis-ci.org/cuonglm/ipdata)
[![Go Report Card](https://goreportcard.com/badge/github.com/cuonglm/ipdata)](https://goreportcard.com/report/github.com/cuonglm/ipdata)
[![GoDoc](https://godoc.org/github.com/cuonglm/ipdata?status.svg)](https://godoc.org/github.com/cuonglm/ipdata)

# Installation

```sh
go get -u github.com/cuonglm/ipdata/cmd/goipdata
```

# Usage

```sh
$ goipdata -ip 8.8.8.8
{
    "ip": "8.8.8.8",
    "city": "",
    "region": "",
    "country_name": "United States",
    "country_code": "US",
    "continent_name": "North America",
    "continent_code": "NA",
    "latitude": 37.751,
    "longitude": -97.822,
    "asn": "AS15169",
    "organisation": "Google LLC",
    "postal": "",
    "currency": "USD",
    "currency_symbol": "$",
    "calling_code": "1",
    "flag": "https://ipdata.co/flags/us.png",
    "time_zone": ""
}
```

# Author

Cuong Manh Le <cuong.manhle.vn@gmail.com>

# License

See [LICENSE](https://github.com/cuonglm/ipdata/blob/master/LICENSE)
