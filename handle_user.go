package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("User successfully switched to '%s'!\n", name)
	return nil
}
