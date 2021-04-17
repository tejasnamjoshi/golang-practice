package main

import (
	"errors"
	"flag"
	"fmt"
	"golang_practice/Gophercizes/HackerNews/hn"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

// var (
// 	cache      []item
// 	cacheExp   time.Time
// 	cacheMutex sync.Mutex
// )

func main() {
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "Port number to start the application on.")
	flag.IntVar(&numStories, "numStories", 30, "Number of stories to dispaly")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))
	http.HandleFunc("/", handler(numStories, tpl))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {

	sc := storyCache{
		numStories: numStories,
		duration:   6 * time.Second,
	}

	go func() {
		ticket := time.NewTicker(3 * time.Second)
		for {
			// wait till ticker sends a tick
			<-ticket.C
			temp := storyCache{
				numStories: numStories,
				duration:   6 * time.Second,
			}

			temp.stories()
			var cache []item
			cache = temp.cache

			sc.mutex.Lock()
			sc.cache = cache
			sc.exp = temp.exp
			sc.mutex.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := templateData{Stories: stories, Time: time.Now().Sub(start)}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

type storyCache struct {
	numStories int
	cache      []item
	exp        time.Time
	duration   time.Duration
	mutex      sync.Mutex
}

// Pointer because we want to maintain the same cache in the same ref.
func (sc *storyCache) stories() ([]item, error) {

	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if time.Now().Sub(sc.exp) < 0 {
		return sc.cache, nil
	}

	stories, err := getTopStories(sc.numStories)
	sc.exp = time.Now().Add(sc.duration)
	if err != nil {
		return nil, err
	}
	sc.cache = stories
	return sc.cache, nil

}

// func getCachedStories(numStories int) ([]item, error) {
// 	cacheMutex.Lock()
// 	defer cacheMutex.Unlock()
// 	if time.Now().Sub(cacheExp) < 0 {
// 		return cache, nil
// 	}
// 	stories, err := getTopStories(numStories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cache = stories
// 	cacheExp = time.Now().Add(5 * time.Minute)

// 	return cache, nil
// }

func getTopStories(numStories int) ([]item, error) {
	var c hn.Client
	ids, err := c.GetTopStories()
	if err != nil {
		return nil, errors.New("Failed to load stories.")
	}
	var stories []item
	at := 0
	for len(stories) < numStories {
		need := numStories - len(stories)
		stories = append(stories, getStories(ids[at:at+need])...)
		at += need
	}
	return stories, nil
}

func getStories(ids []int) []item {
	var c hn.Client
	type result struct {
		idx  int
		item item
		err  error
	}
	resultChan := make(chan result)

	for i := 0; i < len(ids); i++ {
		go func(i int, id int) {
			hnItem, err := c.GetItemById(id)
			if err != nil {
				resultChan <- result{idx: i, err: err}
			}

			resultChan <- result{idx: i, item: parseHNItem(hnItem)}
		}(i, ids[i])
	}

	var results []result
	for i := 0; i < len(ids); i++ {
		results = append(results, <-resultChan)
	}
	sort.Slice(results, func(i, j int) bool { return results[i].idx < results[j].idx })

	var stories []item
	for _, res := range results {
		if res.err != nil {
			continue
		}
		if isStoryLink(res.item) {
			stories = append(stories, res.item)
		}
	}

	return stories
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Host, "www.")
	}

	return ret
}

type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
