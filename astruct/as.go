package astruct

import (
	"gogen/bstruct"
	"gogen/cstruct"
)

type (
	A struct {
		Bb           []bstruct.B    `fixType:"1"`
		Cc           []cstruct.C    `fixType:"2"`
		Ss           string         `fixType:"10"`
		SmcContext5g []SmcContext5g `fixType:"optional"`
	}

	SmcContext5g struct {
		s int
	}
)

func (*A) Af1() {

}
