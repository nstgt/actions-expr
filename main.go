package main

import (
	"fmt"

	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version(version.Print("actions-expr"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	fmt.Println(gopherSay("Hello World!"))
}

func gopherSay(s string) string {
	return fmt.Sprintf("ʕ◔ϖ◔ʔ > %s", s)
}
