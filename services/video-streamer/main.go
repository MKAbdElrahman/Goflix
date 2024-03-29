package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"

	"github.com/ardanlabs/conf/v3"
)

var build = "develop"

func main() {
	// CONFIG
	cfg := struct {
		conf.Version
		Web struct {
			APIHost string `conf:"default:0.0.0.0:3000"`
		}
	}{
		Version: conf.Version{
			Desc:  "Video Streamer Service",
			Build: build,
		},
	}

	const prefix = "STREAMER"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return
		}
		log.Fatal(err.Error())
	}

	// ROUTER
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./videos/"))
	mux.Handle("GET /videos/", http.StripPrefix("/videos", fileServer))

	// SERVER
	log.Info("Streaming server listening", "APIHost", cfg.Web.APIHost)
	http.ListenAndServe(cfg.Web.APIHost, mux)
}
