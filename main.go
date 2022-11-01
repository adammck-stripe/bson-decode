package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	w := flag.CommandLine.Output() // stderr
	noNewline := flag.Bool("n", false, "don't print newline")

	flag.Usage = func() {
		fmt.Fprintf(w, "Usage: %s [<args>] <b64-bson>\n", os.Args[0])
		fmt.Fprintf(w, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	v1, err := base64.StdEncoding.DecodeString(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(w, "Error decoding b64: %v\n", err)
		os.Exit(1)
	}

	var v2 map[string]interface{}
	err = bson.Unmarshal(v1, &v2)
	if err != nil {
		fmt.Fprintf(w, "Error unmarshalling BSON: %v\n", err)
		os.Exit(1)
	}

	out, err := json.Marshal(&v2)
	if err != nil {
		fmt.Fprintf(w, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(out))
	if !*noNewline {
		fmt.Print("\n")
	}
}
