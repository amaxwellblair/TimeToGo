package main_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/amaxwellblair/TimeToGo/scrapers"
)

func TestStackOverflow_ParsePage(t *testing.T) {
	s := NewTestStackOverflow()
	xml, err := s.GetXMLTest()
	if err != nil {
		t.Fatal(err)
	} else if err := s.ParsePage(xml); err != nil {
		t.Fatal(err)
	}
	title := "Hackers wanted - Leidenschaftliche Programmierer / Entwickler (m/w) gesucht. at freiheit.com technologies gmbh (Hamburg, Deutschland)"
	categories := []string{"java", "scala", "javascript", "clojure", "go"}
	item := s.Pages[0].Items[0]

	if item.Title != title {
		t.Fatalf("unexpected title: %s", title)
	} else if reflect.DeepEqual(categories, item.Categories) != true {
		t.Fatalf("unexpected categories: %s", item.Categories)
	}
}

type TestStackOverflow struct {
	url string
	tag string
	main.StackOverflow
}

func NewTestStackOverflow() *TestStackOverflow {
	return &TestStackOverflow{
		url: "./fixtures/stackoverflow_golang.xml",
		tag: "",
	}
}

func (ts *TestStackOverflow) GetXMLTest() ([]byte, error) {
	// Open fixture of sample StackOverflow XML page
	f, err := os.Open(ts.url)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	// Read and return as bytes
	buf := make([]byte, stat.Size())
	if _, err := f.Read(buf); err != nil {
		return nil, err
	}

	return buf, nil
}
