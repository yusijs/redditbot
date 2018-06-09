package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"redditbot/common"
	"redditbot/watcher"
)

func main() {
	filename := flag.String("conf", "./config.json", "Path to your config file (defaults to ./config.json)")
	conf, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal("Failed to read config file", err)
	}

	var config common.Config
	err = json.Unmarshal(conf, &config)
	if err != nil {
		log.Fatal("Failed to parse config", err)
	}
	watcher.WatchAndComment(&config)
}
