package main

import (
    "errors"
    "fmt"
)

// Divide returns the result of dividing a by b, or an error if b is zero.
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a/b, nil
}

func main() {
    a, b := 10.0, 2.0
    result, err := Divide(a, b)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
