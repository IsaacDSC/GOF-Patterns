package main

import (
	"fmt"

	"github.com/IsaacDSC/GOF-Patterns/src/creational/factory"
)

func main() {
	fmt.Println("MAIN INITIALIZED")
	Factory := factory.User{}
	user := Factory.User("Isaac", "DSC")
	fmt.Println(user)

}
