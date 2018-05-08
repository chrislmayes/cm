package main

import (
	"flag"
	"fmt"

	"github.com/ogier/pflag"
)

func init()

//Global Vars
// flags
var (
	user string
)

func main() {
	fmt.Println("Hello, World")

	pflag.SomeFunction(fmt.Println("flag"))
	//a := "a"

	flag.Parse()

	flag.StringVarP(&user, "user", "u", "", "Search Users")

}
