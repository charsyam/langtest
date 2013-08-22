package main

import (
    "fmt"
    "./avl"
)

func main() {
    tree := tree.New()
    tree.Add(50)
    tree.Add(40)
    tree.Add(30)
    tree.Add(20)
    tree.Add(10)

    tree.PostOrder()
}
