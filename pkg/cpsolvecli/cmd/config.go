package cmd

import (
	"github.com/spf13/viper"
)

type CPSolveCLIConfig struct{
	Commands *[]Command
}

type Command struct {
	Usage string `mapstructure:"usage"`
	Method string `mapstructure:"method"`
	Class string `mapstructure:"class"`
	Short string `mapstructure:"short"`
	Long string `mapstructure:"long"`
	Inputs []Input `mapstructure:"inputs"`
}

type Input struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
	Usage string `mapstructure:"usage"`
}

func LoadCommands() (*CPSolveCLIConfig, error) {
	coms, err := ReadInCommandsConfig()
	if err != nil {
		return nil, err
	}

	return &CPSolveCLIConfig{
		Commands: coms,
	}, nil
}



func ReadInCommandsConfig() (*[]Command, error) {
	v := viper.New()
	v.AddConfigPath("pkg/cpsolvecli/cmd")
	v.SetConfigName("cmds")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var commands []Command
	if err := v.UnmarshalKey("commands", &commands); err != nil {
		return nil, err
	}

	return &commands, nil
}