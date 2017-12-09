package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/Gnouc/ipdata"
)

const (
	defaultServer = "https://api.ipdata.co"
	defaultLang   = "en"
)

var (
	serverFlag string
	langFlag   string
	ipFlag     string
	apiKeyFlag string
	helpFlag   bool
)

func init() {
	flag.StringVar(&ipFlag, "ip", "", "ip to gather information")
	flag.StringVar(&serverFlag, "server", defaultServer, fmt.Sprintf("the ip data server, default to %s", defaultServer))
	flag.StringVar(&langFlag, "lang", defaultLang, "the language to use")
	flag.StringVar(&apiKeyFlag, "key", "", "the api key to use")
	flag.BoolVar(&helpFlag, "help", false, "show usage message")
}

func usage() {
	usageStr :=
		`
Usage: %s -ip ip [-lang lang] [-server url] [-key api_key]
`
	fmt.Fprintf(os.Stderr, usageStr, os.Args[0])
}

func main() {
	flag.Parse()

	if helpFlag {
		usage()
		os.Exit(1)
	}
	c := ipdata.NewClient(
		ipdata.WithAPIKey(apiKeyFlag),
		ipdata.WithURL(serverFlag),
		ipdata.WithLanguage(langFlag),
	)

	r, err := c.Lookup(ipdata.WithIP(ipFlag))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

	b, _ := json.MarshalIndent(r, "", "    ")
	fmt.Fprint(os.Stdout, string(b))
}
