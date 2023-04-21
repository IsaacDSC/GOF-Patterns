package main

import "fmt"

// type Body struct {
// 	Sku string
// }

// type AbsTesting[T any] struct {
// 	InferValue T
// }

// type TEEEE[T any] struct {
// 	ACTION string
// 	Body   T
// }

// func main() {
// 	body := Body{Sku: "123"}
// 	v := AbsTesting[TEEEE[Body]]{}
// 	v.InferValue.ACTION = "CHECKOUT"
// 	v.InferValue.Body = body
// 	fmt.Println(v.InferValue.ACTION)
// 	fmt.Println(v.InferValue.Body.Sku)
// }

type Body struct {
	Sku string
}

type TEEEE[T any] struct {
	ACTION string
	Body   T
}

func main() {
	body := Body{Sku: "123"}
	v := TEEEE[Body]{ACTION: "CHECKOUT"}
	v.Body = body
	fmt.Println(v.Body.Sku)
}
