package commenter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"redditbot/common"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// Comment inserts a new comment, based on a developer comment
func Comment(auth *common.Auth, c *common.Comment) {
	token := getAuthToken(auth)
	existingID, body := existingComment(c.PostLink, auth, c.ID)
	if existingID != "" {
		editComment(token, existingID, body, c)
	} else {
		writeComment(token, c)
	}
	c = nil
}

func existingComment(postLink string, auth *common.Auth, commentID string) (string, string) {
	hq := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s.json", postLink), nil)
	if err != nil {
		log.Warnln("Failed to create request", err)
	}
	req.Header.Set("User-Agent", "swbf2bot")

	res, err := hq.Do(req)
	if err != nil {
		log.Warnln("Failed to retrieve post", err)
	}

	response, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Warnln("Failed to retrieve post", err)
	}
	var parsed common.PostResponse
	json.Unmarshal(response, &parsed)
	comments := parsed[1].Data.Children
	var thingID string
	var body string

	for i := range comments {
		if comments[i].Data.Author == auth.Username && comments[i].Data.ID != commentID {
			asd := comments[i].Data
			thingID = fmt.Sprintf("t1_%s", asd.ID)
			body = asd.Body
			break
		}
	}
	return thingID, body
}

func editComment(token string, thingID string, body string, c *common.Comment) {
	href := "https://oauth.reddit.com/api/editusertext"
	form := url.Values{}
	form.Add("api_type", "json")
	form.Add("return_rtjson", "false")
	form.Add("richtext_json", "")
	text := createCommentText(c)
	fullText := fmt.Sprintf("%s\n\n%s", body, text)
	form.Add("text", fullText)
	form.Add("thing_id", thingID)
	executeRequest(form, href, token)
}

func executeRequest(form url.Values, href string, token string) {
	hq := http.Client{}
	req, err := http.NewRequest("POST", href, strings.NewReader(form.Encode()))
	if err != nil {
		log.Warnln("Failed to create request", err)
		return
	}
	req.Header.Set("User-Agent", "swbf2bot")
	req.Header.Set("Authorization", "bearer "+token)
	res, err := hq.Do(req)
	if err != nil {
		log.Warnln("Failed to perform request", err)
		return
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warnln("Failed to parse response from reddit", err)
		return
	}
	if res.StatusCode > 299 {
		log.Warnln("Error from reddit", string(response))
		return
	}
	var body common.CommentResponse
	json.Unmarshal(response, &body)
	if len(body.JSON.Errors) > 0 {
		log.Warnln("Failed to insert comment", body.JSON.Errors)
	} else {
		log.Println("Inserted / Updated comment")
	}
}

func getAuthToken(auth *common.Auth) string {
	authURL := "https://www.reddit.com/api/v1/access_token"
	hq := http.Client{}
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", auth.Username)
	form.Add("password", auth.Password)

	req, err := http.NewRequest("POST", authURL, strings.NewReader(form.Encode()))
	if err != nil {
		log.Warnln("Failed to get auth token")
		log.Warnln(err)
	}
	req.SetBasicAuth(auth.ClientID, auth.ClientSecret)
	req.Header.Set("User-Agent", "swbf2bot")
	res, e := hq.Do(req)
	if e != nil {
		log.Warnln(e)
	}
	var tokenAuth common.Redditauth
	bodyText, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Warnln(e)
	}
	e = json.Unmarshal(bodyText, &tokenAuth)
	if e != nil {
		log.Warnln(e)
	}
	token := tokenAuth.AccessToken
	return token
}

func writeComment(token string, c *common.Comment) {
	href := "https://oauth.reddit.com/api/comment"
	form := url.Values{}
	form.Add("api_type", "json")
	form.Add("return_rtjson", "false")
	form.Add("richtext_json", "")
	text := createCommentText(c)
	form.Add("text", text)
	form.Add("thing_id", c.ParentID)
	executeRequest(form, href, token)
}

func createCommentText(c *common.Comment) string {
	var snippet string
	if len(c.Body) > 30 {
		runes := []rune(c.Body)
		snippet = string(runes[0:30])
	} else {
		snippet = c.Body
	}
	//     p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	date := c.Timestamp.Format("2006-01-02 15:04")
	comment := fmt.Sprintf("%s: Developer response (%s): %s", date, c.Author, snippet)
	href := c.Link
	body := fmt.Sprintf("[%s](%s)", comment, href)
	return body
}
