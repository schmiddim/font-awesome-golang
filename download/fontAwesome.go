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

}

func (fa *FontAwesome) GenerateGoFileFromIcons() error {
	// Set the filename to the desired path
	filename := "generated/iconList.go"

	// Open file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write package declaration
	_, err = file.WriteString("package generated\n\n")
	if err != nil {
		return err
	}

	// Write import declaration
	_, err = file.WriteString("import \"github.com/schmiddim/font-awesome-golang/download\"\n\n")
	if err != nil {
		return err
	}

	// Write variable declaration for the array
	_, err = file.WriteString("var Icons = []download.Icon{\n")
	if err != nil {
		return err
	}

	// Write each icon as an element of the array
	for _, icon := range fa.Icons {
		_, err = file.WriteString(fmt.Sprintf("\t{Changes: %#v, Ligatures: %#v, Search: %#v, Styles: %#v, Unicode: %#v, Label: %#v, Voted: %#v, Free: %#v},\n",
			icon.Changes, icon.Ligatures, icon.Search, icon.Styles, icon.Unicode, icon.Label, icon.Voted, icon.Free))
		if err != nil {
			return err
		}
	}

	// Write closing bracket
	_, err = file.WriteString("}\n")
	if err != nil {
		return err
	}

	return nil
}
