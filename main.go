package main

import (
	"fmt"

	"io/ioutil"
	"strconv"

	flag "github.com/ogier/pflag"
)

//Global Vars
var (
	user      string
	directory string
)

// entry point
func main() {
	//flag.SomeFunction(user.string)
	//a := "a"
	//Parses flags for user
	flag.Parse()

	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		directory = "."

		//fmt.Printf("Usage: %s [options]\n", os.Args[0])
		//fmt.Println("Options:")
		//flag.PrintDefaults()
		//os.Exit(1)
	}

	// if multiple users are passed separated by commas, store them in a "users" array

	//users := strings.Split(user, ",")
	//fmt.Printf("Searching user(s): %s\n", users)
	files, err := ioutil.ReadDir(directory)
	if err == nil {
		for _, fi := range files {
			fmt.Printf("%s         %s        %s\n", fi.Name(), strconv.FormatInt(fi.Size(), 10), fi.ModTime())
		}
	} else {
		fmt.Println(err)

	}

}

func init() {

	//flag.StringVarP(&user, "user", "u", "", "Search Users")
	flag.StringVarP(&directory, "directory", "d", "", "Directory to list files")

}
