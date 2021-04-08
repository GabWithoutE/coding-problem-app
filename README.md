# Coding Problem Application
This application implements coding challenge solutions in a go application, where the user can either interact with it using the REST API, or as a CLI. This application uses the `go-build-template` from [here](https://github.com/thockin/go-build-template) as a starting point.

## An Exercise In...
**Code Structure:**
using `DDD (Domain Driven Development)` or `Group by Context` and `Hexagonal Architecture` as described by ["How do you structure your go apps" Kat Zien](https://www.youtube.com/watch?v=oL6JBUk6tj0)

**Go HTTP Web Service Patterns:**
using patterns described by [Mat Ryer - How I Write HTTP Web Services after Eight Years](https://www.youtube.com/watch?v=rWBSMsLG8po&list=PLbs4C-FV19dTDAtq1pfvcqx-h8utbiVLu&index=1&t=1553s)

**Testing:**
using patterns described by [Mat Ryer - How I Write HTTP Web Services after Eight Years](https://www.youtube.com/watch?v=rWBSMsLG8po&list=PLbs4C-FV19dTDAtq1pfvcqx-h8utbiVLu&index=1&t=1553s)

**Profiling:**
using techniques shown by [Dave Cheney - Two Go Programs, Three Different Profiling Techniques](https://www.youtube.com/watch?v=nok0aYiGiYA&list=PLbs4C-FV19dTDAtq1pfvcqx-h8utbiVLu&index=2)

## Group By Context
Context: Coding Interview Problems Solving

### Language
| Language  | Definition |
| --------  | ---------- |
| Problem   | code interview question with clear parameters and desired solutions |
| Solver    | code implementation of  | 
| Solution  | the return value of the solver | 
| Helper    | a computation that is general, and useful in solutions to problems |
| Handler   | handles the incoming API call, and outgoing response |
| ServeHTTP | handles the http request, and the outgoing http response |
| ServeCLI  | handles the CLI usage, and the output using standard io |

### Entities
| Entity | Definition |
| ------ | ---------- |
| `Problem`Problem  | Has properties for the parameters `Problem` |
| `Problem`Solution | Has properties for the return values of `Solution` |

### Service
| Service    | Description |
| ---------- | ----------- | 
| `Problem`Solving | Given parameters for the `Problem`, return the solution | 
| `Problem``Solver`Implementation | Given the type of `Solver` for the `Problem` return the `Solver`'s implementation | 

### Events
| Event | Description |
| ----- | ----------- |
| `Problem`Solved | Successfully solved `problem` |
| `Problem`InvalidInputs | Invalid inputs to `problem` |
| `Problem`NotFound | Invalid problem was requested |
| `Problem`Unsolvable | The solution was run, and determined that the problem could not be solved with the given valid parameters | 
| `Solver`NotImplemented | The solver for the problem has not been implemented |
| `Solution`InternalFailure | `Solution` implementation has a bug, and failed to run |

## API
| Problem    | Input | Output |
| ---------- | ----- | ------ |
| CoinChange | Coin `denominations`, `total` amount of money | Returns: Fewest number of coins of type `denominations` to make up `total` |
