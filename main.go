package main

import (
	"fmt" //Default format import
	"os"

	"io/ioutil" // Used to read file system
	"strconv"   //

	flag "github.com/ogier/pflag" // Used to pass flags from cli to runtime
)

//Global Variables can be used in any function
var (
	user      string
	directory string
	reverse   string
)

// entry point
func main() {
	//Begin with setting reverse to false, this is for the -r flags usage to trigger listing the ls backwards
	var rev bool = false

	//Parses flags for user
	flag.Parse()

	// if user does not supply flags, use the current directory hence "."
	if flag.NFlag() == 0 {
		directory = "."

	} else {
		if reverse != "" {
			directory = reverse
			rev = true
		}
	}

	files, err := ioutil.ReadDir(directory)
	if err == nil {
		if rev == true {
			var tempfiles []os.FileInfo = make([]os.FileInfo, len(files))
			for i := len(files) - 1; i > -1; i-- {
				tempfiles[len(files)-1-i] = files[i]

			}
			files = tempfiles
		}
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
	flag.StringVarP(&reverse, "reverse", "r", "", "Reverse List")

}
