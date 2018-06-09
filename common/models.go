package common

import "time"

// CommentResponse is the return value when posting a comment
type CommentResponse struct {
	JSON struct {
		Errors []interface{} `json:"errors"`
		Data   struct {
			Things []struct {
				Kind string    `json:"kind"`
				Data ThingData `json:"data"`
			} `json:"things"`
		} `json:"data"`
	} `json:"json"`
}

// ThingData is the actual thing returned from reddit comments
type ThingData struct {
	SubredditID                string        `json:"subreddit_id"`
	ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
	Distinguished              interface{}   `json:"distinguished"`
	ModReasonBy                interface{}   `json:"mod_reason_by"`
	BannedBy                   interface{}   `json:"banned_by"`
	AuthorFlairType            string        `json:"author_flair_type"`
	RemovalReason              interface{}   `json:"removal_reason"`
	LinkID                     string        `json:"link_id"`
	AuthorFlairTemplateID      interface{}   `json:"author_flair_template_id"`
	Likes                      bool          `json:"likes"`
	Replies                    string        `json:"replies"`
	UserReports                []interface{} `json:"user_reports"`
	Saved                      bool          `json:"saved"`
	ID                         string        `json:"id"`
	BannedAtUtc                interface{}   `json:"banned_at_utc"`
	ModReasonTitle             interface{}   `json:"mod_reason_title"`
	Gilded                     int           `json:"gilded"`
	Archived                   bool          `json:"archived"`
	ReportReasons              interface{}   `json:"report_reasons"`
	Author                     string        `json:"author"`
	CanModPost                 bool          `json:"can_mod_post"`
	SendReplies                bool          `json:"send_replies"`
	ParentID                   string        `json:"parent_id"`
	Score                      int           `json:"score"`
	ApprovedBy                 interface{}   `json:"approved_by"`
	ModNote                    interface{}   `json:"mod_note"`
	Collapsed                  bool          `json:"collapsed"`
	Body                       string        `json:"body"`
	Edited                     bool          `json:"edited"`
	AuthorFlairCSSClass        interface{}   `json:"author_flair_css_class"`
	Name                       string        `json:"name"`
	Downs                      int           `json:"downs"`
	AuthorFlairRichtext        []interface{} `json:"author_flair_richtext"`
	IsSubmitter                bool          `json:"is_submitter"`
	CollapsedReason            interface{}   `json:"collapsed_reason"`
	BodyHTML                   string        `json:"body_html"`
	Stickied                   bool          `json:"stickied"`
	CanGild                    bool          `json:"can_gild"`
	Subreddit                  string        `json:"subreddit"`
	AuthorFlairTextColor       interface{}   `json:"author_flair_text_color"`
	ScoreHidden                bool          `json:"score_hidden"`
	Permalink                  string        `json:"permalink"`
	NumReports                 interface{}   `json:"num_reports"`
	NoFollow                   bool          `json:"no_follow"`
	Created                    float64       `json:"created"`
	AuthorFlairText            interface{}   `json:"author_flair_text"`
	RteMode                    string        `json:"rte_mode"`
	CreatedUtc                 float64       `json:"created_utc"`
	SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
	Controversiality           int           `json:"controversiality"`
	AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
	ModReports                 []interface{} `json:"mod_reports"`
	SubredditType              string        `json:"subreddit_type"`
	Ups                        int           `json:"ups"`
}

// Config contains the user information & the accounts to monitor
type Config struct {
	Auth      *Auth
	Accounts  []string
	SubReddit string `json:"subReddit"`
}

// Auth contains the logon information to retrieve a Oauth token
type Auth struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// Comment represents a parsed comment from the reddit live api
type Comment struct {
	ID        string `json:"id" gorethink:"_id" mapstructure:"_id"`
	Name      string `json:"name" gorethink:"name" mapstructure:"name"`
	Link      string `json:"link" gorethink:"link" mapstructure:"link"`
	PostLink  string `json:"postLink" gorethink:"postLink" mapstructure:"postLink"`
	PostTitle string `json:"title" gorethink:"title" mapstructure:"title"`

	Timestamp   time.Time `json:"timestamp" gorethink:"timestamp" mapstructure:"timestamp"`
	Author      string    `json:"author" gorethink:"author" mapstructure:"author"`
	AuthorFlair string    `json:"authorFlair" gorethink:"authorFlair" mapstructure:"authorFlair"`

	Subreddit   string `json:"subreddit" gorethink:"subreddit" mapstructure:"subreddit"`
	SubredditID string `json:"subredditId" gorethink:"subredditId" mapstructure:"subredditId"`

	Body     string `json:"body" gorethink:"body" mapstructure:"body"`
	ParentID string `json:"parentId" gorethink:"parentId" mapstructure:"parentId"`
}

// Redditauth is the model used when posting comments
type Redditauth struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"bearer"`
}

// CommentBody is the data sent when creating a new comment
type CommentBody struct {
	APIType      string `json:"api_type"`
	ReturnRtJSON bool   `json:"return_rtjson"`
	RtJSON       string `json:"richtext_json"`
	Body         string `json:"text"`
	PostID       string `json:"thing_id"`
}
