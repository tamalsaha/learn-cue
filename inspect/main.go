package main

import (
	"fmt"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

func main() {
	c := cuecontext.New()

	// read and compile value
	d, err := os.ReadFile("examples/values.cue")
	if err != nil {
		panic(err)
	}
	val := c.CompileBytes(d)

	paths := []string{
		//"a",
		//"d.f",
		//"l",
		"r",
	}

	for _, path := range paths {
		fmt.Printf("====  %s  ====\n", path)
		v := val.LookupPath(cue.ParsePath(path))
		p := v.Path()
		var x interface{}
		if err := v.Decode(&x); err != nil {
			panic(err)
		} else {
			fmt.Printf("%v\n", x)
		}
		fmt.Printf("%q\n%# v\n", p, v)
	}
}
