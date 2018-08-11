package command

import "strings"

type EchoCommand struct {
}

func (c *EchoCommand) Execute(input Input) (*Output, error) {
	resp := strings.Join(input.args, " ")
	if len(input.args) == 0 {
		resp = "You need to type a message, dude. '/echo you are great"
	}
	out := Output{Text: resp}
	return &out, nil
}
