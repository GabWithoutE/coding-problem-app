package cmd

import (
	"fmt"
	"github.com/gabriellukechen/coding-problem-app/pkg/solving"
	"github.com/spf13/cobra"
	"reflect"
)

func buildCobraCommands(config *CPSolveCLIConfig) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "cpsolve",
		Short: "CLI for interacting with implemented coding problem solutions",
		Long: `
cpsolve is a CLI for the backend logic implemented in package:
	github.com/gabriellukechen/coding-problem-app/pkg/solving

Using cpsolve, a user can specify a particular coding problem as a command,
specify the inputs and get the solution from the output of the command.
This is really just an exercise in making a CLI application
`,
	}

	buildChildCommands(rootCmd, config)

	cobra.CheckErr(rootCmd.Execute())

	return rootCmd
}

func buildChildCommands(root *cobra.Command, config *CPSolveCLIConfig) {
	commands := *config.Commands

	for _, command := range commands {
		flagVars := make([]interface{}, 0)

		cmd := &cobra.Command{
			Use:   command.Usage,
			Short: command.Short,
			Long:  command.Long,
			Run: func(cmd *cobra.Command, args []string) {
				probs := solving.NewProblemsCatalogue()
				c := reflect.ValueOf(probs)

				vs := make([]reflect.Value, 0)
				for i, _ := range flagVars {
					switch reflect.TypeOf(flagVars[i]).String() {
					case "*string":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*string))))
					case "*[]string":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*[]string))))
					case "*int":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*int))))
					case "*[]int":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*[]int))))
					case "*bool":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*bool))))
					case "*[]bool":
						vs = append(vs, reflect.ValueOf(*(flagVars[i].(*[]bool))))
					}
				}

				problem := c.MethodByName(command.Method).Call(vs)
				i, _ := problem[0].Interface().(solving.Problem).Solve()
				fmt.Printf("%v", i)
			},
		}

		for _, in := range command.Inputs {
			n := in.Name

			switch in.Type {
			case "string":
				var tmp string
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().(string)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().StringVar(&flagVar, n, "", in.Usage)
			case "[]string":
				var tmp []string
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().([]string)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().StringSliceVar(&flagVar, n, nil, in.Usage)
			case "int":
				var tmp int
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().(int)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().IntVar(&flagVar, n, 0, in.Usage)
			case "[]int":
				var tmp []int
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().([]int)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().IntSliceVar(&flagVar, n, nil, in.Usage)
			case "bool":
				var tmp bool
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().(bool)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().BoolVar(&flagVar, n, false, in.Usage)
			case "[]bool":
				var tmp []bool
				t := reflect.TypeOf(tmp)

				flagVar := reflect.Zero(t).Interface().([]bool)
				flagVars = append(flagVars, &flagVar)
				cmd.Flags().BoolSliceVar(&flagVar, n, nil, in.Usage)
			}
		}

		root.AddCommand(cmd)
	}
}