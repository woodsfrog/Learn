package main

import(
    "fmt"
)

func main() {
    scores := make([]int, 0, 5)
    c := cap(scores)
    fmt.Println(c)

    for i := 0; i < 25; i++ {
        scores = append(scores, i)
        // if our capacity has changed,
        // Go had to grow our array to accommodate the new data
        if cap(scores) != c {
            c = cap(scores)
            fmt.Println(c)
        }
    }
}


