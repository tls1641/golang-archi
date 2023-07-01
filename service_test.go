package architecture

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Put(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)

	p := Person{
		First: "James",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)

	ctl.Finish()
}

// func ExamplePut() {
// 	mdb := DB{}
// 	p := Person{
// 		First: "James",
// 	}

// 	Put(mdb, 1, p)
// 	got := mdb.Retrieve(1)
// 	fmt.Println(got)
// 	//Output: {James}

// }
