# CLI Design
## Command
The main command is cpsolve.

## Child Commands
Each coding problem is a child command.

**Design Considerations based on cool points brought up in [Carolyn Van's lecture on CLIs](https://www.youtube.com/watch?v=eMz0vni6PAw)**
- Separating command implementation from Cobra, and Viper
    - Why? Being able to separate Cobra wiring from command implementation is beneficial, because it allows commands to be implemented without understanding of Cobra (and related wiring), as well as makes testing commands easier.
    - Difficulties? Each Coding Problem command requires a different type of resource. Realistically, this application does not make much sense as a CLI, and without significant workarounds, each command will be hard coded as child commands, instead of positional parameters
      - ^ currently considering ways to solve this