package levelgen

import (
	//"time"
	//"math/rand"
	"testing"
)

func TestSimpleLevelGen(t *testing.T) {

	var seed int64 = 1

	lvl := MakeLevel(500, 500, seed)

	lvl.GenSimpleRandomLVL(10, seed)
	
}