package command

import (
	"net/http"
	"github.com/ilkinulas/cengizabi/config"
	"fmt"
	"io/ioutil"
	"log"
)

type Podcast struct {
	cfg    config.Podcast
	logger *log.Logger
}

func NewPodcast(cfg config.Podcast, logger *log.Logger) *Podcast {
	return &Podcast{
		cfg:    cfg,
		logger: logger,
	}
}

func (c *Podcast) Execute(input Input) (*Output, error) {
	if len(input.args) < 1 {
		return nil, fmt.Errorf("missing command action")
	}
	action := input.args[0]
	switch action {
	case "save":
		return c.handleSaveUrl(input.args)
	default:
		return nil, fmt.Errorf("unrecognized action %v", action)
	}
}

func (c *Podcast) handleSaveUrl(args [] string) (*Output, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("missing url")
	}
	youtubeUrl := args[1]
	url := fmt.Sprintf("%v/save?url=%v", c.cfg.BaseUrl, youtubeUrl)
	c.logger.Printf("URL = %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call url %v. %v", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read http rsponse %v", err)
	}

	return &Output{Text: string(body[:])}, nil
}
