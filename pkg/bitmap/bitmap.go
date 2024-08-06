package bitmap

import "fmt"

type Bitmap struct {
	bits []byte
	size int
}

func NewBitmap(size int) *Bitmap {

	if size == 0 {
		size = 250
	}

	return &Bitmap{
		bits: make([]byte, size),
		size: size * 8,
	}
}

// [[0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0], [0,0,0,0,0,0,0,0]]
func (b *Bitmap) Set(id string) {
	idx := hash(id) % b.size // 22  32
	fmt.Println("id:" + id + " ,idx:" + fmt.Sprintf("%d", idx))
	byteIdx := idx / 8 // 2  4
	bitIdx := idx % 8  // 6  0
	b.bits[byteIdx] |= 1 << bitIdx
}

func (b *Bitmap) IsSet(id string) bool {
	idx := hash(id) % b.size // 22  32
	fmt.Println("id:" + id + "idx:" + fmt.Sprintf("%d", idx))
	byteIdx := idx / 8 // 2  4
	bitIdx := idx % 8  // 6  0
	return (b.bits[byteIdx] & 1 << bitIdx) != 0
}

func (b *Bitmap) Export() []byte {
	return b.bits
}

func Load(bits []byte) *Bitmap {
	if len(bits) == 0 {
		return NewBitmap(0)
	}

	return &Bitmap{
		bits: bits,
		size: len(bits) * 8,
	}
}

func hash(id string) int {
	// 使用BKDR哈希算法
	seed := 131313 // 31 131 1313 13131 131313, etc
	hash := 0
	for _, c := range id {
		hash = hash*seed + int(c)
	}
	return hash & 0x7FFFFFFF
}
