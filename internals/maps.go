package internals

import (
	"fmt"
	"hash/maphash"
)

func hashString() {
	var h1 maphash.Hash
	h1.WriteString("abc")
	fmt.Printf("%d\n", h1.Sum64())

	var h2 maphash.Hash
	h2.SetSeed(h1.Seed())
	h2.WriteString("abc")
	fmt.Printf("%d\n", h2.Sum64())
	fmt.Println(h1)
}

func StartMaps() {
	hashString()
}
