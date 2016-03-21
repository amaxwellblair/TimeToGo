package main

import "github.com/amaxwellblair/TimeToGo/models"

func main() {

}

type itemExtractor interface {
	ItemExtract(models.Item) error
}

// Scrape data from APIs
func Scrape(a API, s itemExtractor) error {
	if err := a.ExtractData(); err != nil {
		return err
	}
	pages := a.GetPages()

	for _, page := range pages {
		items := page.Items
		for _, item := range items {
			if err := s.ItemExtract(item); err != nil {
				return err
			}
		}
	}
	return nil
}

// API interface
type API interface {
	ExtractData() error
	GetPages() []*Page
}

// Page represents a single page
type Page struct {
	Items []*Item
}

// Item represents a single item
type Item struct {
	Title       string
	Categories  []string
	Description string
}

// GetTitle serves as a getter for Title
func (i *Item) GetTitle() string { return i.Title }

// GetCategories serves as a getter for Categories
func (i *Item) GetCategories() []string { return i.Categories }

// GetDescription serves as a getter for Description
func (i *Item) GetDescription() string { return i.Description }
