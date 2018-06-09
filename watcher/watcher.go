package watcher

import (
	"fmt"
	"log"
	"redditbot/commenter"
	"redditbot/common"
	"regexp"
	"strings"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type announcer struct{}

var auth *common.Auth
var subreddit string

// WatchAndComment initializes the bot, and monitors the given users
func WatchAndComment(appConfig *common.Config) {
	users := &appConfig.Accounts
	subreddit = appConfig.SubReddit
	auth = appConfig.Auth
	cfg := initConfig(users)
	bot, _ := reddit.NewScript("graw:swbf2bot:0.1.0 by /u/_yusi_", 5*time.Second)

	_, wait, err := graw.Scan(&announcer{}, bot, cfg)
	if err != nil {
		log.Print("Failed to init reddit scan")
		log.Fatal(err)
	}
	log.Println("Scanning reddit...")
	if err := wait(); err != nil {
		log.Fatal("Scanner bot crashed: ", err)
	}
}

func (a *announcer) UserComment(orig *reddit.Comment) error {
	p := regexp.MustCompile(fmt.Sprintf("(?i)%s", subreddit))
	if p.MatchString(orig.Subreddit) {
		converted := convertComment(orig)
		go commenter.Comment(auth, converted)
	}
	return nil
}
func (a *announcer) UserPost(orig *reddit.Post) error {
	return nil
}

func convertComment(orig *reddit.Comment) *common.Comment {
	strct := common.Comment{}
	strct.ID = orig.ID
	strct.Author = orig.Author
	strct.Timestamp = time.Unix(int64(orig.CreatedUTC), 0)
	strct.AuthorFlair = orig.AuthorFlairText
	strct.Name = orig.Name
	strct.Link = getCommentURL(orig.Permalink)
	strct.PostLink = getCommentPost(orig.Permalink, orig.ID)
	strct.PostTitle = orig.LinkTitle
	strct.Subreddit = orig.Subreddit
	strct.SubredditID = orig.SubredditID
	strct.Body = orig.Body
	strct.ParentID = orig.ParentID
	return &strct
}

func initConfig(users *[]string) graw.Config {
	cfg := graw.Config{Users: *users}
	return cfg
}

func getCommentPost(commentLink string, id string) string {
	index := strings.Index(commentLink, id)
	if index < 0 {
		return ""
	}
	permaLink := common.Substring(commentLink, 0, index)
	redditURL := "https://reddit.com"
	return fmt.Sprintf("%v%v", redditURL, permaLink)
}

func getCommentURL(permalink string) string {
	redditURL := "https://reddit.com"
	context := "?context=3"
	return fmt.Sprintf("%v%v%v", redditURL, permalink, context)
}
