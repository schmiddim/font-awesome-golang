package download

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// define constants for filename
const (
	fileName = "fontawesome.json"
)

type Icon struct {
	Changes   []string            `json:"changes"`
	Ligatures []string            `json:"ligatures"`
	Search    map[string][]string `json:"search"`
	Styles    []string            `json:"styles"`
	Unicode   string              `json:"unicode"`
	Label     string              `json:"label"`
	Voted     bool                `json:"voted"`
	Free      []string            `json:"free"`
}

type FontAwesome struct {
	Icons map[string]Icon `json:"icons"`
}

func (fa *FontAwesome) FetchMedataJson(version string) error {
	url := "https://raw.githubusercontent.com/FortAwesome/Font-Awesome/" + version + "/metadata/icons.json"
	path := fileName
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (fa *FontAwesome) ParseMetaData() {
	jsonFile, err := os.Open(fileName) // replace with the path to your icons.json file
	if err != nil {
		panic(err)

	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &fa.Icons)
	if err != nil {
		panic(err)
	}

	for iconName := range fa.Icons {
		fmt.Println(iconName)
	}

}

// Move to other lib
func (fa *FontAwesome) GetIcon(name string) Icon {
	return fa.Icons[name]
}
