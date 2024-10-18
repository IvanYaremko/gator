package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	request, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error getting feed: %w", err)
	}
	request.Header.Add("User-Agent", "gator")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error with response: %w", err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error io readall: %w", err)
	}

	rssFeed := RSSFeed{}
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshalling xml: %w", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	fmt.Println("Blog: ", rssFeed.Channel.Title)
	fmt.Println("Description: ", rssFeed.Channel.Title)

	for _, item := range rssFeed.Channel.Item {
		fmt.Println("Title: ", item.Title)
		fmt.Println("Publication date: ", item.PubDate)
		fmt.Println("Link: ", item.Link)
		fmt.Println("Description: ", item.Description)
	}
	return &RSSFeed{}, nil
}
