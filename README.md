# Go Language

A collection of focused Go examples for learning the language, from basic syntax to concurrency and generics. Each topic lives in its own directory so you can read, run, and adapt examples as you progress.

## Requirements

- [Go 1.18 or later](https://go.dev/dl/) (required for the generics examples)

Check your installation with:

```sh
go version
```

## Getting started

Clone the repository, move into its directory, and run an example by providing the path to its Go file:

```sh
git clone https://github.com/Vineeth-28/Go-lang.git
cd Go-lang
go run ./1_hello_world/main.go
```

Replace the example path with a file from another topic directory. Some directories contain multiple independent programs; run the file you want to explore. For example:

```sh
go run ./16_Struct/construtorstruct.go
go run ./16_Struct/StructComposition.go
```

## Topics

| Directory | Topics |
| --- | --- |
| `1_hello_world` | Your first Go program |
| `2_Simple_values` | Basic values and types |
| `3_Variables` | Variable declarations |
| `4_Constants` | Constants |
| `5_For_Loop` | `for` loops |
| `6_if_Else` | Conditional statements |
| `07_Switch` | `switch` statements |
| `8_Arrays` | Arrays |
| `9_Slices` | Slices |
| `10_Maps` | Maps |
| `11_Range` | Iteration with `range` |
| `12_Functions` | Functions |
| `13_Variadic_Function` | Variadic functions |
| `14_Closures` | Closures |
| `15_Pointers` | Pointers |
| `16_Struct` | Structs, initialization, and composition |
| `17_Interfaces` | Interfaces |
| `18_Enums` | Enum-style patterns |
| `19_Generics` | Type parameters and constraints |
| `20_GoRoutines` | Goroutines |
| `21_WaitGroup` | Coordinating goroutines with a wait group |
| `22_Channels` | Channels |
| `23_Mutex` | Mutexes and synchronization |
| `24_FileHandling` | Reading and writing files |
| `25_Packages` | Packages and imports |

## How to learn with the examples

1. Choose a topic directory.
2. Read the Go source file and follow the code in `main`.
3. Run it with `go run ./<directory>/<file>.go`.
4. Change the example and run it again to see how the behavior changes.

These are learning examples rather than a single application or Go module. In `24_FileHandling`, the read example expects an `example.txt` file in the current working directory; the write example creates `output.txt` there.

## Further reading

- [The Go Programming Language](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/)
- [Go documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
