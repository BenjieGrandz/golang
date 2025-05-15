package main

import "fmt"

func main() {
	// empty array/default -- holds 5 int 
    var a [6]int
    fmt.Println("emp:", a)

    a[4] = 100 //seting an index specifically
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

	// declare and initialize
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// automatically counts elements
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	// anything in between will be zeroed
    b = [...]int{100, 3: 700, 500}
    fmt.Println("idx:", b)

	// 2d arrays
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	// create and declare multiarreys
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}