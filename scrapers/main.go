package main

import (
	"fmt"

	"github.com/amaxwellblair/TimeToGo/keys"
	"github.com/amaxwellblair/TimeToGo/models"
)

func main() {
	// Create new API structs
	a := NewStackOverflow([]string{"golang", "ruby", "python", "clojure", "c", "c++"})

	// Open the database
	s := models.NewStore()
	if err := s.Open(keys.DATASOURCEDEV); err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}

	// Scrape the APIs for jobs data
	if err := Scrape(a, s); err != nil {
		fmt.Printf("unexpected error: %s\n", err)
		return
	}
	fmt.Println("The scraper has successfully consumed the API(s)")
}

type itemExtractor interface {
	ItemExtract(models.Item) error
}

// Scrape data from APIs
func Scrape(a API, s itemExtractor) error {
	// Extract data from API and place in API struct
	if err := a.ExtractData(); err != nil {
		return err
	}
	pages := a.GetPages()

	// Iterate through each page and place the items into the database
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
