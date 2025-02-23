# Introduction to Golang

- Golang is a compiled language.

- Golang compiles faster than C, C++, RUST, etc.

- Go Tool can run file directly without VM.

- Executables are different for different OS.

- Golang can be used to build from Web Applications to System apps - Cloud.

- It is already in Production.

- Golang is kind of Object Oriented Programming Language. We don't have classes in Golang but we have structs as an alternative.

- We don't have Method Overloading and Overrinding.

## What is Missing?

- No Try & Catch.

- Lexer does not work.(No use of Semi Colon)

## LEXER

- A "lexer" in programming, short for "lexical analyzer," is a component of a compiler that reads raw source code and breaks it down into smaller, meaningful units called "tokens," essentially acting like a first step to understand the structure of the code by identifying keywords, operators, identifiers, and other basic elements within the input stream of characters.

```
Note: In short You can think of lexer to be sort of like a Grammer Checker for Programming Languages
```
```
Example:

Input: "int x = 5;"
Output (tokens): [ "int", "identifier:x", "=", "integer:5", ";"] 
```


## Types

- Almost Case Insensitive

- Variable type should be known before hand

- Everything is a type:
    - Primary
        - String
        - Bool
        - Integer [uint8, uint64, int8, int64, uintptr] ; aliases are involved too
        - Floating [float32, float64]
        - Complex
    
    - Advance Types
        - Array
        - Slices
        - Maps
        - Structs
        - Pointers
        - Functions
        - Channels

```
Note: In golang, a variable with a first capital letter is a public variable.

For eg: 
    const LoginToken string = "qwertyuiop"
    
    This 'LoginToken' is now accessible in any other file in the folder or in the file itself.

```
- The above is the reason that you will notice that in print statement fmt.Println(), the 'P' is capital.

## Go Build

- use `go env` in the terminal to find more about environment variables.
- We can create different builds for different OS using the GOOS="windows" or GOOS="darwin" or GOOS="linux".
```
Eg:
    GOOS="windows" go build filename
```

## Memory Management in Golang

- Garbage Collection happens automatically
- Out of Scope or nil

### new()
- Allocate memory but no INIT
- you will get a Memory address
- zeroed storage

### make()
- Allocate memory and INIT
- you will get a Memory address
- non-zeroed storage
