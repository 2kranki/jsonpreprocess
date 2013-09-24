package main

import (
	"io/ioutil"
	"os"

	"github.com/hnakamur/jsonutil"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = jsonutil.WriteUncommentedTo(os.Stdout, string(input))
}