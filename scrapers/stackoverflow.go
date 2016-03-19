package main

import (
	"io/ioutil"
	"net/http"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

// StackOverflow represents the StackOverflow API
type StackOverflow struct {
	url string
	tag string
}

// StackOverflowItem represents a job posting
type StackOverflowItem struct {
	Title       string
	Categories  []string
	Description string
}

// NewStackOverflow returns a new instance of the StackOverflow API
func NewStackOverflow(tag string) *StackOverflow {
	return &StackOverflow{
		url: "http://stackoverflow.com/jobs/feed?tags=",
		tag: tag,
	}
}

// GetXML scrapes RSS data from stackoverflow
func (s *StackOverflow) GetXML() (*http.Response, error) {
	return http.Get(s.url + s.tag)
}

// ParseXML parses scraped RSS data from stackoverflow
func (s *StackOverflow) ParseXML(page []byte) error {
	// Parse the web page
	doc, err := gokogiri.ParseXml(page)
	if err != nil {
		return err
	}
	defer doc.Free()

	s.ParseItems(doc.Root())

	return nil
}

// ExtractJobs pulls data from StackOverflow and places them in a database
func (s *StackOverflow) ExtractJobs() error {
	resp, err := s.GetXML()
	if err != nil {
		return nil
	}
	// Extract the response body
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s.ParseXML(page)

	return nil
}

// ParseItems parses XML data into items
func (s *StackOverflow) ParseItems(root xml.Node) ([]*StackOverflowItem, error) {
	i, err := root.Search("/rss/channel/item")
	if err != nil {
		return nil, err
	}
	var items []*StackOverflowItem
	for _, v := range i {
		categories, err := s.ParseCategories(v)
		if err != nil {
			return nil, err
		}
		items = append(items, &StackOverflowItem{
			Categories: categories,
		})
	}

	return items, nil
}

// ParseCategories parses XML data into skills
func (s *StackOverflow) ParseCategories(node xml.Node) ([]string, error) {
	// Find categories in each item
	c, err := node.Search("category")
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, v := range c {
		categories = append(categories, v.Content())
	}
	return nil, nil
}

// func main() {
// 	page, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("unexpected error: %s \n", err)
// 		return
// 	}
// 	defer resp.Body.Close()
//
// 	// parse the web page
// 	doc, err := gokogiri.ParseXml(page)
// 	if err != nil {
// 		fmt.Printf("unexpected error: %s \n", err)
// 		return
// 	}
// 	defer doc.Free()
//
// 	node := doc.Root().FirstChild().FirstChild()
// 	for i := 0; i < 5; i, node = i+1, node.NextSibling() {
// 		fmt.Println(node.Path())
// 	}
//
// 	node = doc.Root()
// 	fmt.Println(node.Search("/rss/channel/item[1]/category"))
//
// 	// perform operations on the parsed page -- consult the tests for examples
//
// 	// important -- don't forget to free the resources when you're done!
//
// }
