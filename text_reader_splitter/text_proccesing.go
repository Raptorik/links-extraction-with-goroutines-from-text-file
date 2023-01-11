package text_reader_splitter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Collect struct {
	_arrayOfLinks []string
}

func NewLinksCollect() *Collect {
	res := &Collect{}
	return res
}

func (c *Collect) FindLinks(fileContent string) string {
	re, err := regexp.Compile(`(?:https?://)?(?:[^/.]+\.)*google\.com(?:/[^/\s]+)*/?`) // regex for http links in text
	if err != nil {
		fmt.Println(err)
	}
	c._arrayOfLinks = re.FindAllString(string(fileContent), -1) // creation of slice of links
	data := strings.Join(c._arrayOfLinks, " ")                  // each link as a string from a slice separated with ""
	return data
}

func CreateFileWriteDataToFile(filename, data string) {
	file, err := os.Create(filename) // creation of new file
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(data) // writing strings with links to the new file
}

func OpenPrintFileData(s string) {
	contentLinks, err := os.ReadFile(s) // new file opening
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(contentLinks))
}
