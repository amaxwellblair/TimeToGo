package main

import (
	"os"
	"testing"
)

func TestStackOverflow_ParsePosts(t *testing.T) {
	s := NewTestStackOverflow()
	page, err := s.GetXMLTest()
	if err != nil {
		t.Fatal(err)
	}
	s.ParseXML(page)
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
