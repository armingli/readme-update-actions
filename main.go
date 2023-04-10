package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mmcdole/gofeed"
)

func main() {
	// get the rss list from the actions env
	rss_url := os.Getenv("INPUT_RSS_LIST")

	// get the number of posts or stories to commit
	max_post, err := strconv.Atoi(os.Getenv("INPUT_MAX_POST"))
	if err != nil || max_post == 0 {
		max_post = 5
	}

	// get readme path from the actions env
	readme_path := os.Getenv("INPUT_README_PATH")

	// if path not provided default to root readme
	if readme_path == "" {
		readme_path = "./README.md"
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss_url)

	// store the posts
	var items []string

	// get the posts
	// format it according to readme links format
	if feed != nil && len(feed.Items) > 0 {
		for i := 0; i < max_post; i++ {
			if i < len(feed.Items) {
				item := fmt.Sprintf("- [%s](%s)", feed.Items[i].Title, feed.Items[i].Link)
				items = append(items, item)
			} else {
				break
			}
		}
	}

	// find readme and replace with our result
	WriteToFile(readme_path, items)

}

func WriteToFile(path string, items []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f.Close()

	f, err = os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stop := false
	w := bufio.NewWriter(f)
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if !stop && line == "<!-- BLOG-POST-LIST:START -->" {
			fmt.Fprintln(w, line)
			stop = true
		}

		if line == "<!-- BLOG-POST-LIST:END -->" {
			stop = false
			for _, item := range items {
				fmt.Fprintln(w, item)
			}
		}

		if !stop {
			fmt.Fprintln(w, line)
		}

	}
	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}
