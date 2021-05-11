# CLI Design
## Command
The main command is cpsolve.

## Child Commands
Each coding problem is a child command.

### Configuring Child Commands/Extending the CLI
**GOALS:**
1. Extending the CLI--adding new commands to it--... should not require knowledge of underlying CLI framework (Cobra)
2. ... should not require knowledge of unexported interfaces in "business logic"
3. Each command should do a task--a `Problem Solve` task, e.g. solve Problem: WordBreakProblem

**Implementation (Addressing Goals):**
1. New commands are configured using toml, yaml, or json, any config format supported by viper
```toml
# Format Example in TOML
[[commands]] # list of command Tables
usage = "wordbreak" # the name of the child command 
method = "NewWordBreakProblem" # the name of the factory for generating the Problem interface
short = "Break up the string, so that all substrings are in the word dictionary." # short description
long = "Problem: Given a string, and a dictionary of words, add spaces to s to create a valid set of dictionary words. Consider and enumerate all valid such sets." # long description
[[commands.inputs]] # nested list of Tables for command input flags
name = "UnbrokenString" # flag name
type = "string" # type of input
usage = "Input string to be broken up into sets of valid WordDictionary substrings" # description of flag
[[commands.inputs]]
name = "WordDictionary"
type = "[]string"
usage = "WordDictionary as list of strings"
# ... more [[commands]] can be added with the above format
```

2. The business logic exports an interface that has the `Problem` factory 
   implementations necessary for creating and running a `Problem Solve Task`.
   The CLI implements a layer that calls these exported methods.
    - Considerations, tradeoffs, and subgoals:
    - The interface exported by the solving package is meant to enforce a 
      particular pattern for calling the business logic. In this case, create 
      a Problem, then call the Method `Solve` on the Problem. Enforcing this
      pattern means the struct for each problem shouldn't be exported. See
      root project README for why this is enforced [README.md](../../README.md).
      The CLI implementation also follows this pattern, wherein the config for
      a command, the factory method is specified, and the inputs in the exported
      function are specified as well.
    - TradeOffs: the input type configuration is coupled to the Golang implementation.
    - The input implementation is also quite inconvenient. Cobra requires flag
      variables to be declared in scope before the command is defined, and flags
      require the command to be defined to be "binded". In addition, commands can
      take any generic as input types. These create an awkward scenario where the vars
      need to be stored generically, but before they're turned into type reflect.Value
      for making Method call, they need to be dereferenced... Maybe there's a better
      workaround for this in a future refactor. The future implementation of generics
      in Golang would be very helpful with this.
    - The order of inputs in the commands configuration matters, as that's the order 
      they'll be passed into the factory as arguments.
    
### Commands
```shell
$ cpsolve wordbreak --UnbrokenString --WordDictionary
```