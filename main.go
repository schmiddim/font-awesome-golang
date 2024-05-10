package main

import (
	"fmt"
	fa "github.com/schmiddim/font-awesome-golang/download"
	"github.com/schmiddim/font-awesome-golang/generated"
	"log"
)

func main() {
	f := fa.FontAwesome{}
	err := f.FetchMedataJson("5.x")
	if err != nil {
		fmt.Println(err)
	}

	f.ParseMetaData()
	err = f.GenerateGoFileFromIcons()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)
	list := generated.Icons
	for item := range list {
		fmt.Println(list[item])
	}
}
