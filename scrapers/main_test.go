package main_test

import (
	"testing"

	"github.com/amaxwellblair/TimeToGo/scrapers"
	"github.com/amaxwellblair/TimeToGo/testHelpers"
)

func Test_Scrape_StackOverflow(t *testing.T) {
	ts := testhelper.NewTestStore()
	ts.MustOpen()
	defer ts.Close()
	defer ts.CleanDatabase()
	s := main.NewStackOverflow([]string{"golang", "ruby"})

	if err := main.Scrape(s, ts); err != nil {
		t.Fatal(err)
	}

	pages := s.GetPages()
	items := pages[0].Items
	if len(pages) != 2 {
		t.Fatalf("unexpected number of pages: %#v", len(pages))
	} else if len(items) < 1 {
		t.Fatalf("less than expected number of items: %#v", len(items))
	}
}

func Test_ItemExtract(t *testing.T) {
	ts := testhelper.NewTestStore()
	ts.MustOpen()
	defer ts.Close()
	defer ts.CleanDatabase()

	item := &main.Item{
		Title:       "Bacon taster",
		Categories:  []string{"java", "clojure", "love"},
		Description: "Bacon eating ninja who wants to turn their passion into cash",
	}
	if err := ts.ItemExtract(item); err != nil {
		t.Fatal(err)
	}

	if job, err := ts.Job(item.Title); err != nil {
		t.Fatal(err)
	} else if job.Name != item.Title {
		t.Fatalf("unexpected name: %s", item.Title)
	}

	skills, err := ts.SkillsByJob(item.Title)
	if err != nil {
		t.Fatal(err)
	}

	for _, actual := range skills {
		comp := false
		for _, expected := range item.Categories {
			if expected == actual.Name {
				comp = true
			}
		}
		if comp != true {
			t.Fatalf("unexpected skill %s", actual.Name)
		}
	}

	if skills, err = ts.Skills(); err != nil {
		t.Fatal(err)
	} else if len(skills) != 3 {
		t.Fatalf("unexpected length: %#v", len(skills))
	}
}
