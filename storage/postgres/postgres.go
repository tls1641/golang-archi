package postgres

import architecture "github.com/tls1641/golang-archi"

type DB map[int]architecture.Person

func (pg DB) Save(n int, p architecture.Person) {
	pg[n] = p
}

func (pg DB) Retrieve(n int) architecture.Person {
	return pg[n]
}
