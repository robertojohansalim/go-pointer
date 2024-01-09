# Go Pointer

Very simple package to pointerize variable using Go Generic

Inspired by https://github.com/nicklaros/gopointer

Implemented using generics for go 1.18 or higher

## Requirement
> Require Go version 1.18 or higher
### Installation
```bash
go get github.com/robertojohansalim/go-pointer
```

## Problem
Imagine a scenario with a struct that need to be created
```go
type NewStruct struct{
    Value *string
}
```
Sometimes when a sturct require a pointer field this is what usually happened
```go
tempValue := "Lorem Ipsum"

newStruct := NewStruct{
    Value: &tempValue
}
```
Or some sorcerry that ends up looks like this
```go
newStruct := NewStruct{
    Value: &[]string{"Lorem Ipsum"}[0]
}
```

## Usage
The usage of this package simplify the needs of creating a temporary variable just to create a pointer of it
```go
import gopointer "github.com/robertojohansalim/go-pointer"
...
newStruct := NewStruct{
    Value: gopointer.Pointerize("Lorem Ipsum")
}
```


Utilizing Go 1.18 Generics, this function can be used for any type of variable
```go
ptrString := gopointer.Pointerize("Lorem Ipsum")
ptrInt := gopointer.Pointerize(20)
```