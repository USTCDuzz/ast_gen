package gogen

import (
	"fmt"
	"gogen/astruct"
	"gogen/bstruct"
	"gogen/cstruct"
)

func Main() {
	a := &astruct.A{
		Bb: []bstruct.B{
			{
				Ee: []byte{0, 1, 2, 3},
			},
		},
		Cc: []cstruct.C{
			{
				StrC: []string{"a", "bv", "c", "d", "e"},
				Dd:   []uint32{1, 2, 3, 4, 5, 6},
			},
			{
				StrC: []string{"a", "bv", "c", "d", "e"},
				Dd:   []uint32{1, 2, 3, 4, 5, 6},
			},
		},
		Ss: "knickknack",
	}
	fmt.Printf("%#v\n", a)
}
