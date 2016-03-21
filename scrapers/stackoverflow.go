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
	tags  []string
	Pages []*Page
}

// NewStackOverflow returns a new instance of the StackOverflow API
func NewStackOverflow(tags []string) *StackOverflow {
	return &StackOverflow{
		url:  "http://stackoverflow.com/jobs/feed?tags=",
		tags: tags,
	}
}

// GetXML scrapes RSS data from stackoverflow
func (s *StackOverflow) GetXML(index int) (*http.Response, error) {
	return http.Get(s.url + s.tags[index])
}

// GetPages returns the APIs pages
func (s *StackOverflow) GetPages() []*Page {
	return s.Pages
}

// ExtractData pulls data from StackOverflow and places them in a database
func (s *StackOverflow) ExtractData() error {
	for i := 0; i < len(s.tags); i++ {
		// Request data from StackOverflow
		resp, err := s.GetXML(i)
		if err != nil {
			return nil
		}

		// Extract the response body
		page, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := s.ParsePage(page); err != nil {
			return err
		}
	}

	return nil
}

// ParsePage parses scraped RSS data from stackoverflow
func (s *StackOverflow) ParsePage(data []byte) error {
	// Parse the web page
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		return err
	}
	defer doc.Free()

	var page Page

	// Find raw XML items
	items, err := doc.Root().Search("/rss/channel/item")
	if err != nil {
		return err
	}

	// Parse each of the XML items
	for _, item := range items {
		i, err := s.ParseItem(item)
		if err != nil {
			return err
		}
		page.Items = append(page.Items, i)
	}

	// Place the completed page into the API
	s.Pages = append(s.Pages, &page)

	return nil
}

// ParseItem parses XML data into items
func (s *StackOverflow) ParseItem(node xml.Node) (*Item, error) {
	// Parse different data points based on path
	var item *Item
	paths := []string{"category", "title", "description"}

	data := make(map[string][]string)
	var err error
	for _, path := range paths {
		if data[path], err = s.ParseData(path, node); err != nil {
			return nil, err
		}
	}
	item = &Item{
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
