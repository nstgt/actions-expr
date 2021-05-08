package main

import "fmt"

func main() {
	fmt.Println(gopherSay("Hello World!"))
}

func gopherSay(s string) string {
	return fmt.Sprintf("ʕ◔ϖ◔ʔ > %s", s)
}
