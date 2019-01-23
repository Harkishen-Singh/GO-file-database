package main

import(
    "fmt"
)

func try(){
    var p *[]int
    x := []int{4,5}
    fmt.Println(x)
    p = &x
    i := 0
    fmt.Println(*p)
    for i<10 {
        *p = append(*p, i)
        fmt.Println(*p)
        i++
    }
}