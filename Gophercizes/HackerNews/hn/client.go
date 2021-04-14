package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiBase = "https://hacker-news.firebaseio.com/v0"
)

type Client struct {
	baseUrl string
}

type Item struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Id          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	URL         string `json:"url"`
}

func (c *Client) setDefaults() {
	c.baseUrl = apiBase
}

func (c Client) GetTopStories() ([]int, error) {
	c.setDefaults()
	resp, err := http.Get(fmt.Sprintf("%s/newstories.json", c.baseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r := json.NewDecoder(resp.Body)
	var ids []int
	err = r.Decode(&ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (c Client) GetItemById(id int) (Item, error) {
	c.setDefaults()
	var item Item
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.baseUrl, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()
	r := json.NewDecoder(resp.Body)
	err = r.Decode(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (item Item) Print() {
	fmt.Printf("Id - %d \nBy - %s \nTitle - %s \nScore - %d\n", item.Id, item.By, item.Text, item.Score)
}
