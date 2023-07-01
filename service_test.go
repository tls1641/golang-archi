package architecture

import (
	"fmt"
	"testing"
)

type DB map[int]Person

func (m DB) Save(n int, p Person) {
	m[n] = p
}

func (m DB) Retrieve(n int) Person {
	return m[n]
}

func Test_Put(t *testing.T) {
	mdb := DB{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if got != p {
		t.Fatalf("want %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := DB{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	//Output: {James}

}
