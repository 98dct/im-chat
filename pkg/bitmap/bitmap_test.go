package bitmap

import (
	"fmt"
	"testing"
)

//	1 0 0 0 0 0 0
//
// [[0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0]]
func TestBitmap(t *testing.T) {

	b := NewBitmap(10)

	b.Set("pppp")
	b.Set("222")
	b.Set("pppp")
	b.Set("ccc")
	b.Set("eee")
	b.Set("fff")
	for _, bit := range b.bits {
		t.Logf("%b, %v", bit, bit)
	}

}

func Test111(t *testing.T) {

	// 0000 0001
	res := 1 << 5
	fmt.Printf("%b", res)
}
