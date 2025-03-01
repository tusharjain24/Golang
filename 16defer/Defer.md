# Defer Statement in Golang

In Go, the `defer` statement is used to delay the execution of a function until the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function completes.

## Syntax

```go
defer functionName(parameters)
```

## Key Points

- Deferred functions are executed in LIFO (Last In, First Out) order.
- Commonly used for resource cleanup, such as closing files or network connections.
- Helps in writing cleaner code by keeping cleanup code close to the resource allocation.

## Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred")
    fmt.Println("End")
}
```

### Output

```
Start
End
Deferred
```

In this example, the `defer` statement ensures that "Deferred" is printed after "End", even though the `defer` statement appears before it in the code.

## Use Cases

- Closing files
- Unlocking mutexes
- Logging
- Releasing resources

Using `defer` can make your code more readable and maintainable by handling cleanup tasks in a structured manner.