package utils

import (
	"fmt"
	"github.com/kr/pretty"
)

func Dump(vs ...interface{}) {
	for _, v := range vs {
		switch v.(type) {
		case string:
			fmt.Printf("%v\n", pretty.Formatter(v))
		default:
			fmt.Printf("%# v\n", pretty.Formatter(v))
		}
	}
}
