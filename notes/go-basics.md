## Basic structure

package main

import "fmt"

func main() {
    fmt.Println("hello world!!")  //one-line comments
}

---

## Variables

- Long way to declare variabls
- wallrus declaration :=
- declaration and initialization (just as their name suggests)
- walrus operator cannot be used to initialize global variables
- non initialised variables have a base value
- type is infered when the walrus operator is used

---

## Constants

- declaration and initialisation

<pre> ``` const NAME = "bwanaChairman" ``` <pre>

- numeric const has no type untill it is used in a situation that can help it infer the type
- performs arithmethic manipulation with infinite/arbitrary precision
- this means that it has no limits
- can be used anywhere