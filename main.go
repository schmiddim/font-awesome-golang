package main

import (
	"fmt"
	fa "github.com/schmiddim/font-awesome-golang/download"
)

func main() {
	f := fa.FontAwesome{}
	err := f.FetchMedataJson("5.x")
	if err != nil {
		fmt.Println(err)
	}

	f.ParseMetaData()

	fmt.Println(err)

}
