package main

import (
	"io/ioutil"
	"net/http"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

// StackOverflow represents the StackOverflow API
type StackOverflow struct {
	url   string
	tag   string
	Pages []*StackOverflowPage
}

// StackOverflowPage represents a single page
type StackOverflowPage struct {
	Items []*StackOverflowItem
}

// StackOverflowItem represents a single item
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

	s.ParsePage(page)

	return nil
}

// ParsePage parses scraped RSS data from stackoverflow
func (s *StackOverflow) ParsePage(data []byte) (*StackOverflowPage, error) {
	// Parse the web page
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		return nil, err
	}
	defer doc.Free()

	var page StackOverflowPage

	// Find raw XML items
	items, err := doc.Root().Search("/rss/channel/item")
	if err != nil {
		return nil, err
	}

	// Parse each of the XML items
	for _, item := range items {
		i, err := s.ParseItem(item)
		if err != nil {
			return nil, err
		}
		page.Items = append(page.Items, i)
	}

	return &page, nil
}

// ParseItem parses XML data into items
func (s *StackOverflow) ParseItem(node xml.Node) (*StackOverflowItem, error) {
	// Parse different data points based on path
	var item *StackOverflowItem
	paths := []string{"category", "title", "description"}

	data := make(map[string][]string)
	var err error
	for _, path := range paths {
		if data[path], err = s.ParseData(path, node); err != nil {
			return nil, err
		}
	}
	item = &StackOverflowItem{
		Title:       data["title"][0],
		Categories:  data["category"],
		Description: data["description"][0],
	}

	return item, nil
}

// ParseData parses XML data into skills
func (s *StackOverflow) ParseData(path string, node xml.Node) ([]string, error) {
	// Find data in each node
	d, err := node.Search(path)
	if err != nil {
		return nil, err
	}
	var data []string
	for _, v := range d {
		content, err := s.ParseHTML([]byte(v.Content()))
		if err != nil {
			return nil, err
		}
		data = append(data, content)
	}
	return data, nil
}

// ParseHTML will take in HTML and return a string
func (s *StackOverflow) ParseHTML(data []byte) (string, error) {
	doc, err := gokogiri.ParseHtml(data)
	if err != nil {
		return "", err
	}
	defer doc.Free()

	return doc.Root().Content(), nil
}
