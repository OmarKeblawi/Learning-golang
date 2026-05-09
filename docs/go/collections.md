# Go Collections

Go has three common built-in collection types that show up often in everyday code: arrays, slices, and maps. They solve different problems, and in most real programs you will use slices and maps far more often than arrays.

## Arrays

An array has a fixed length that is part of its type. That means `[3]int` and `[5]int` are different types.

Use arrays when the size is known ahead of time and will not change, or when you want a value with a fixed number of elements.

```go
package main

import "fmt"

func main() {
    numbers := [3]int{1, 2, 3}
    fmt.Println(numbers)

    var names [2]string
    names[0] = "Go"
    names[1] = "Tutorial"
    fmt.Println(names)
}
```

Common uses:

- Fixed-size data where the number of elements never changes.
- Low-level code that needs exact size and value semantics.
- Small, known-length groups of values.

## Slices

A slice is a flexible view over an underlying array. It has a length and a capacity, and it can grow with `append`.

Use slices when you want a dynamic list of values. This is the default choice for most sequence data in Go.

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4)

    fmt.Println(numbers)
    fmt.Println("len:", len(numbers))
    fmt.Println("cap:", cap(numbers))

    names := make([]string, 0, 5)
    names = append(names, "Alice", "Bob")
    fmt.Println(names)
}
```

Ways to initialize slices:

- `var s []int` creates a nil slice.
- `s := []int{1, 2, 3}` creates a populated slice.
- `s := make([]int, 3)` creates a slice with length 3.
- `s := make([]int, 0, 10)` creates an empty slice with room to grow.

Common uses:

- Lists of items that can grow or shrink.
- Function inputs and outputs for ordered data.
- Building batches of values before processing them.

Important note: slicing an existing slice does not copy the data. The new slice still points to the same backing array.

```go
numbers := []int{10, 20, 30, 40}
part := numbers[1:3]
fmt.Println(part) // [20 30]
```

## Maps

A map stores key-value pairs. You look up a value by key instead of by index.

Use maps when you need fast lookup by a meaningful key, such as a name, ID, or code.

```go
package main

import "fmt"

func main() {
    ages := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }

    ages["Carol"] = 28
    fmt.Println(ages["Alice"])

    age, ok := ages["Dave"]
    fmt.Println(age, ok)
}
```

Ways to initialize maps:

- `var m map[string]int` creates a nil map that cannot be written to yet.
- `m := map[string]int{}` creates an empty, ready-to-use map.
- `m := make(map[string]int)` also creates an empty map.

Common uses:

- Lookup tables.
- Counting values, such as word frequencies.
- Caching computed results.
- Grouping data by key.

Example: counting words.

```go
counts := make(map[string]int)

words := []string{"go", "go", "tutorial"}
for _, word := range words {
    counts[word]++
}

fmt.Println(counts)
```

## Quick comparison

- Array: fixed length, value with size baked into the type.
- Slice: dynamic, resizable sequence backed by an array.
- Map: key-value lookup structure.

## Practical rule of thumb

- Use arrays when the size is fixed and part of the design.
- Use slices for most ordered collections.
- Use maps when you need fast lookup by key.

## Loops in Go (`for` and while-style)

Go has one loop keyword: `for`. What other languages call `while` is written using `for` with only a condition.

### `range` loop (arrays, slices, maps)

Use `range` when you want to iterate over elements directly.

```go
intArr := [...]int32{1, 2, 3}
for i, v := range intArr {
    fmt.Printf("Index: %d Value: %d\n", i, v)
}

myMap := map[string]uint8{"adam": 1, "eve": 2}
for name, age := range myMap {
    fmt.Println(name, age)
}
```

Notes:

- For arrays/slices, `range` gives index and value.
- For maps, `range` gives key and value.
- If you only need one value, use `_` for the unused one.

### While-style loop in Go

This is equivalent to a classic `while (condition)` loop in other languages.

```go
i := 0
for i < len(intArr) {
    fmt.Printf("Index: %d Value: %d\n", i, intArr[i])
    i++
}
```

Use this form when:

- You need explicit control of loop state.
- You do not want `range` semantics.

### Classic C-style `for`

Use initializer, condition, and post statement.

```go
for i := 0; i < len(intArr); i++ {
    fmt.Printf("Index: %d Value: %d\n", i, intArr[i])
}
```

Use this form when:

- You need precise index control.
- You need non-default step logic.

### Infinite loop form

```go
for {
    // do work
    break
}
```

Useful for event loops or workers that run until `break`, `return`, or `panic`.

### Practical tip for your `timeLoop` style example

If you are appending many items, pre-allocate capacity to reduce reallocations:

```go
slice := make([]int, 0, n)
for i := 0; i < n; i++ {
    slice = append(slice, i)
}
```

## Related example in this repo

The file [cmd/tut2/main.go](../../cmd/tut2/main.go) shows arrays, slices, maps, and multiple loop styles in a runnable example.