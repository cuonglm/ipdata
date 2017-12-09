# ipdata - Go library for ipdata.co

[![Build Status](https://travis-ci.org/Gnouc/ipdata.svg?branch=master)](https://travis-ci.org/Gnouc/ipdata)
[![Go Report Card](https://goreportcard.com/badge/github.com/Gnouc/ipdata)](https://goreportcard.com/report/github.com/Gnouc/ipdata)
[![GoDoc](https://godoc.org/github.com/Gnouc/ipdata?status.svg)](https://godoc.org/github.com/Gnouc/ipdata)

# Installation

```sh
go get -u github.com/Gnouc/ipdata/cmd/goipdata
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

See [LICENSE](https://github.com/Gnouc/ipdata/blob/master/LICENSE)
