package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/skitta/gotrans/api/baidu"
)

var to = flag.String("t", "zh", "target language")
var h = flag.String("h", "usage", "help messages")

func main() {
	flag.Parse()
	args := os.Args

	if len(args) == 2 {
		translate(args[1])
	} else if len(args) == 3 {
		if args[1] == "-h" {
			help(args[2])
		} else {
			usage()
		}
	} else if len(args) == 4 {
		if args[1] == "-t" {
			translate(args[3])
		} else {
			usage()
		}
	} else {
		usage()
	}
}

func usage() {
	text, _ := ioutil.ReadFile("./doc/usage")
	fmt.Println(string(text))
}

func translate(qurey string) {
	result, err := baidu.Translator(qurey, "auto", *to)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.Result[0].Dst)
}

func help(qurey string) {
	switch qurey {
	case "langcode":
		text, _ := ioutil.ReadFile("./doc/langcode")
		fmt.Println(string(text))
	}
}
