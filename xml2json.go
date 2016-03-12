package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const VERSION = "20160312"

func main() {
	flag.Parse()
	if *port > 0 {
		if *port < 65535 {
			startHTTPServer(*port)
		} else {
			log.Fatalln("wrong http port:", *port)
		}
	} else {
		command()
	}
}

func init() {
	d := flag.Usage
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Convert XML to JSON")
		fmt.Fprintln(os.Stderr, "version:", VERSION)
		eg := `
Example:
xml2json -xml b.xml
xml2json -xml b.xml -jsonschema b-schema.json  #fix json with JSONSchema
xml2json -xml b.xml -out b.json
cat b.xml|xml2json
xml2json -port 8080
`
		fmt.Fprintln(os.Stderr, eg)
		fmt.Fprintln(os.Stderr, "site:", "https://github.com/hidu/xml2json")
		fmt.Fprintln(os.Stderr, "\n")
		d()
	}
}
