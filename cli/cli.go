package cli

import (
	"errors"
	"fmt"
	"os"
)

type Command struct {
	Name string
	Subcommands []*Command
	Aliases []string
}

type PersistentFlag struct {
}

type Flag struct {
}

type commandNotFoundError struct {
	err  string
	name *string
}

type flagNotFoundError struct {
	err  string
	flag *Flag
}

type noSubcommandsProvided struct {
	err string
}

type helpError struct {
	err string
}

func NewCommand(name string, aliases []string) *Command {
	return &Command{
		Name:        name,
		Subcommands: nil,
		Aliases:     aliases,
	}
}

func (command *Command) AddCommand(name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, NewCommand(name, aliases))
}

func NewCommandNotFoundError(name *string) *commandNotFoundError {
	return &commandNotFoundError{
		err:  fmt.Sprintf("%s: invalid command", *name),
		name: name,
	}
}

func (c *commandNotFoundError) Error() error {
	return errors.New(c.err)
}

func (c *flagNotFoundError) Error() error {
	return errors.New(c.err)
}

func getActiveSubcommands() ([]string, error) {
	if len(os.Args) > 1 {
		return os.Args[1:], nil
	}
	return os.Args[1:], nil
}

func (parentCommand *Command) hasSubcommand(childCommand *Command) (*commandNotFoundError, bool) {
	for _, subcommand := range parentCommand.Subcommands {
		if subcommand.Name == childCommand.Name {
			return nil, true
		}
	}
	return NewCommandNotFoundError(&childCommand.Name), false
}

func (parentCommand *Command) validate() *commandNotFoundError {
	for _, subcommand := range parentCommand.Subcommands {
		if err, ok := parentCommand.hasSubcommand(subcommand); !ok {
			return err
		}
	}
	return nil
}

func validateCommands(parent *Command, child *Command) *commandNotFoundError {
	parent.hasSubcommand(child)
	for _, subcommand := range parent.Subcommands {
		if subcommand.Name == child.Name {
			return nil
		}
	}
	return NewCommandNotFoundError(&child.Name)
}

func Execute() error {

	return nil
}
