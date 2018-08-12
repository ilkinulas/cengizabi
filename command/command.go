package command

import (
	"fmt"
	"strings"
	"github.com/ilkinulas/cengizabi/config"
	"log"
)

type Command interface {
	Execute(input Input) (*Output, error)
}

type Output struct {
	Text string
}

type Input struct {
	cmd  string
	args [] string
}

type Registry struct {
	commands map[string]Command
	cfg      config.Config
	logger   *log.Logger
}

func (r *Registry) addCommand(key string, command Command) {
	r.commands[key] = command
}

func NewRegistry(cfg config.Config, logger *log.Logger) *Registry {
	registry := Registry{
		cfg:      cfg,
		commands: make(map[string]Command),
		logger:   logger,
	}
	registry.addCommand("/echo", &Echo{})
	registry.addCommand("/podcast", NewPodcast(cfg.Podcast, logger))
	return &registry
}

func (r *Registry) HandleMessage(rawMessage string) (*Output, error) {
	parts := strings.Fields(rawMessage)
	cmdStr := parts[0]
	args := parts[1:]

	input := Input{cmd: cmdStr, args: args}

	if cmd, ok := r.commands[input.cmd]; ok {
		out, err := cmd.Execute(input)
		if err != nil {
			return nil, err
		}
		return out, nil
	}

	return nil, fmt.Errorf("command not found. message = %q", rawMessage)
}
