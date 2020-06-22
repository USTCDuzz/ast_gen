package astruct

import (
	"gogen/bstruct"
	"gogen/cstruct"
)

type (
	A struct {
		Bb    []bstruct.B `fixType:"optional"`
		Cc    []cstruct.C `fixType:"var 2"` // 可能需要循环指针
		Dd    []D         `fixType:"optional"`
		Ss    string      `fixType:"var 10"`
		BtFix []byte      `fixType:"fix 11"`
		BtVar []byte      `fixType:"var 12"`
	}
	D struct {

	}
)

func (*A) Af1() {

}
