package levelgen

import (
	"time"
	"testing"
)

func TestSimpleLevelGen(t *testing.T) {

	var seed int64 = time.Now().Unix()

	lvl := MakeLevel(500, 500, seed)

	lvl.GenSimpleRandomLVL(10, seed)
	
	lvl.Draw("test_simple_random_lvl.svg")

}