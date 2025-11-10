package cli

import "fmt"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command doesn't exist")
	}

	err := handler(s, cmd)
	if err != nil {
		return fmt.Errorf("command handler failed: %w", err)
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}
