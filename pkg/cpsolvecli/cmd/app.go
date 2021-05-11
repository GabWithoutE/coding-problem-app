package cmd

import "github.com/spf13/cobra"

type CPSolveCLIApp struct {
	Config  *CPSolveCLIConfig
	command *cobra.Command
}

func NewCPSolveCLIApp() (*CPSolveCLIApp, error) {
	c, err := LoadCommands()
	if err != nil {
		return nil, err
	}

	return &CPSolveCLIApp{
		Config: c,
	}, nil
}

func (a *CPSolveCLIApp) BuildCommands() {
	buildCobraCommands(a.Config)
}
