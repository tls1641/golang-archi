package main

import (
	"fmt"

	architecture "github.com/tls1641/golang-archi"
	"github.com/tls1641/golang-archi/storage/mongo"
	"github.com/tls1641/golang-archi/storage/postgres"
)

func main() {

	dbm := mongo.DB{}
	dbp := postgres.DB{}

	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "Jenny",
	}

	ps := architecture.NewPersonService(dbm)

	ps.Get(1)
	// or store in some other data storage
	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))

}
