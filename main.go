package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/mmcdole/gofeed"
	"log"
	"os"
	"time"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type article struct {
	title string
	body string
	url string
	timestamp *time.Time
}

type rssFeed struct {
	title string
	articles []article
}

var feedMap map[string]rssFeed

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("error initilizing termui: %v", err)
	}
	defer ui.Close()

	feedMap = make(map[string]rssFeed)
	fmt.Println(os.Args)

	err := mainMenu()
	if err != nil {
		log.Fatalf("error rendering main menu %v", err)
	}

}
func mainMenu() error {
	menu := promptui.Select{
		Label: "Main Menu",
		Items: []string{"read", "add"},
	}
	_, c, err := menu.Run()

	if err != nil {
		return err
	}

	switch c {
	case "read":
		err := parseFeed(os.Args[1], feedMap)
		if err != nil {
			log.Fatalf("unable to parse rss feed %v", err)
		}
	case "add":
		fmt.Println("Add")
	}
	return nil
}

//func addFeed(feed string, feedMap map[string]rssFeed) error {
//
//}

func parseFeed(feed string, feedMap map[string]rssFeed) error {
	fp := gofeed.NewParser()
	f, err := fp.ParseURL(feed)
	if err != nil {
		return fmt.Errorf("error parsing feed %v", err)
	}
	fm := feedMap[feed]
	fm.title = f.Title
	var articles []article
	//fmt.Println(f)

	for _, i := range f.Items {
		a := article{}
		a.title, a.body, a.timestamp, a.url = i.Title, i.Description, i.PublishedParsed, i.Link
		articles = append(articles, a)
	}
	fm.articles = articles
    for _, a := range fm.articles {
    	p := widgets.NewParagraph()
    	p.Title**

	}
	return nil
}
