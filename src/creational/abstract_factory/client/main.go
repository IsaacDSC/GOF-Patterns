package main

import (
	"errors"
	"fmt"
	"log"

	abstract_factory "github.com/IsaacDSC/GOF-Patterns/src/creational/abstract_factory"
)

// EXAMPLE LAYER INFRA -> CONTROLLER

func main() {
	// EXEMPLO CHECKOUT DA ENTRADA DOS DADOS NA CAMADA DE EXTERNA DA COMuNICACAO (INFRA)
	commerce := abstract_factory.AbsctractCommerce{}
	commerce.InputCheckout = []struct {
		Sku      string
		Quantity int
	}{{Sku: "123", Quantity: 2}, {Sku: "456", Quantity: 2}}
	response := commerce.Implements("enterpriseId_123", "checkout")a
	if response != "SUCCESS" {
		log.Fatalf(errors.New("NOT-SUCCESS").Error())
	}
	fmt.Println(commerce.ReturnCheckout)

	// EXEMPLO CHECKSTOCK DA ENTRADA DOS DADOS NA CAMADA DE EXTERNA DA COMuNICACAO (INFRA)
	commerce.InputCheckStock = struct{ Sku string }{Sku: "133"}
	checkStockResponse := commerce.Implements("enterpriseId_123", "checkstock")
	if checkStockResponse != "SUCCESS" {
		log.Fatalf(errors.New("NOT-SUCCESS").Error())
	}
	fmt.Println(commerce.ReturnCheckStock)
}
