package mongo

import "github.com/tls1641/architecture"

type DB map[int]architecture.Person

func (m DB) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m DB) Retrieve(n int) architecture.Person {
	return m[n]
}
