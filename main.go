package main

import (
	"github.com/skitta/gotrans/api/baidu"
	"io/ioutil"
	"flag"
	"fmt"
	"os"
)

var to = flag.String("t", "zh", "target language")

func usage() {
	text, _ := ioutil.ReadFile("./doc/usage")
	fmt.Println(string(text))
}

func main() {
	flag.Parse()
	args := os.Args

	var query string

	if len(args) == 2 {
		query = args[1]
	} else if len(args) > 2 {
		query = args[3]
	} else {
		usage()
		return
	}

	result, err := baidu.Translator(query, "auto", *to)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.Result[0].Dst)
}
