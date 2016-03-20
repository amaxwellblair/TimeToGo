package main

import (
	"os"
	"testing"
)

func TestStackOverflow_ParsePage(t *testing.T) {
	s := NewTestStackOverflow()
	xml, err := s.GetXMLTest()
	if err != nil {
		t.Fatal(err)
	}
	page, err := s.ParsePage(xml)
	title := "Hackers wanted - Leidenschaftliche Programmierer / Entwickler (m/w) gesucht. at freiheit.com technologies gmbh (Hamburg, Deutschland)"
	categories := []string{"java", "scala", "javascript", "clojure", "go"}
	item := page.Items[0]

	if item.Title != title {
		t.Fatalf("unexpected title: %s", title)
	}
	var comp bool
	for _, actual := range item.Categories {
		comp = false
		for _, expect := range categories {
			if actual == expect {
				comp = true
			}
		}
		if comp != true {
			t.Fatalf("unexpected category: %s", actual)
		}
	}
}

type TestStackOverflow struct {
	url string
	tag string
	StackOverflow
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
