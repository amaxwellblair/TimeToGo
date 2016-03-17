package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/moovweb/gokogiri"
)

func main() {
	// FIXME This is test code to figure out the stackoverflow API
	// fetch and read a web page
	resp, err := http.Get("http://stackoverflow.com/jobs/feed?tags=golang")
	if err != nil {
		fmt.Printf("unexpected error: %s \n", err)
		return
	}
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("unexpected error: %s \n", err)
		return
	}
	defer resp.Body.Close()

	// parse the web page
	doc, err := gokogiri.ParseXml(page)
	if err != nil {
		fmt.Printf("unexpected error: %s \n", err)
		return
	}
	defer doc.Free()

	node := doc.Root().FirstChild().FirstChild()
	for i := 0; i < 5; i, node = i+1, node.NextSibling() {
		fmt.Println(node.Path())
	}

	node = doc.Root()
	fmt.Println(node.Search("/rss/channel/item[1]/category"))

	// perform operations on the parsed page -- consult the tests for examples

	// important -- don't forget to free the resources when you're done!

}
