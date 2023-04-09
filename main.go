package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	helpers "github.com/readme-update-actions/pkg/utils"

	"github.com/mmcdole/gofeed"
)

func main() {
	// get the rss list from the actions env
	rss_url, _ := helpers.GetEnvString("RSS_LIST")

	// get the number of posts or stories to commit
	max_post, _ := helpers.GetEnvInteger("MAX_POST")

	// if max_post not in env var set default to 3
	if max_post == 0 {
		max_post = 3
	}

	// get readme path from the actions env
	readme_path, _ := helpers.GetEnvString("README_PATH")

	// if path not provided default to root readme
	if readme_path == "" {
		readme_path = "./README.md"
	}

	// get username
	commit_user, _ := helpers.GetEnvString("COMMIT_USER")
	if commit_user == "" {
		commit_user = "readme-update-bot"
	}

	// git user email
	commit_email, _ := helpers.GetEnvString("COMMIT_EMAIL")
	if commit_email == "" {
		commit_email = "readme-update-actions@example.com"
	}

	// git commit message
	commit_message, _ := helpers.GetEnvString("COMMIT_MESSAGE")
	if commit_message == "" {
		commit_message = "Update readme with latest blogs"
	}

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(rss_url)

	// store the posts
	var items []string

	// get the posts
	// format it according to readme links format
	for i := 0; i < max_post; i++ {
		if i < len(feed.Items) {
			item := fmt.Sprintf("- [%s](%s)\n", feed.Items[i].Title, feed.Items[i].Link)
			items = append(items, item)
		}
	}

	// find readme and replace with our result
	err := helpers.ReplaceFile(readme_path, items)
	if err != nil {
		log.Fatalf("Error updating readme %s", err)
	}

	// set git user name
	nameCmd := exec.Command("git", "config", "user.name", commit_user)
	err = nameCmd.Run()
	if err != nil {
		log.Fatalf("Error setting git user %s", err)
	}

	// set git user email
	emailCmd := exec.Command("git", "config", "user.email", commit_email)
	err = emailCmd.Run()
	if err != nil {
		log.Fatalf("Error setting git email %s", err)
	}

	// check git status
	statusCmd, err := exec.Command("git", "status").Output()
	if err != nil {
		log.Fatal(err)
	}

	statusOutput := string(statusCmd)
	if !strings.Contains(statusOutput, "nothing to commit") {
		// add to staging area
		addCmd := exec.Command("git", "add", readme_path)
		err = addCmd.Run()
		if err != nil {
			log.Fatalf("Error adding to staging area %s", err)
			return
		}

		// do git commit
		commitCmd := exec.Command("git", "commit", "-m", commit_message)
		err = commitCmd.Run()
		if err != nil {
			log.Fatalf("Error commiting to repo %s", err)
			return
		}

		// do git push
		pushCmd := exec.Command("git", "push")
		err = pushCmd.Run()
		if err != nil {
			log.Fatalf("Error pushing to repo %s", err)
			return
		}
	}
}
