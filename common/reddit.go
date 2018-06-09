package common

type PostResponse []struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				Subreddit     string      `json:"subreddit"`
				ID            string      `json:"id"`
				Body          string      `json:"body"`
				Gilded        int         `json:"gilded"`
				Clicked       bool        `json:"clicked"`
				Name          string      `json:"name"`
				Distinguished interface{} `json:"distinguished"`
				Author        string      `json:"author"`
			} `json:"data"`
		} `json:"children"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}
