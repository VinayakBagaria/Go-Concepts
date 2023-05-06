package bloomfilters

import (
	"fmt"
	"hash/fnv"
	"math"
)

type hashFunction func(element string) int

func hash32Bit(element string) int {
	h := fnv.New32a()
	h.Write([]byte(element))
	floatedHash := float64(h.Sum32())
	return int(math.Abs(floatedHash))
}

func hash64Bit(element string) int {
	h := fnv.New64a()
	h.Write([]byte(element))
	floatedHash := float64(h.Sum64())
	return int(math.Abs(floatedHash))
}

type BloomFilters struct {
	size           int
	data           []int
	hashFunctions  []hashFunction
	insertionCount int
}

func NewBloomFilters() *BloomFilters {
	size := 10
	return &BloomFilters{size: size, data: make([]int, size), hashFunctions: []hashFunction{hash32Bit, hash64Bit}}
}

func (b *BloomFilters) insert(element string) {
	for _, hashFunction := range b.hashFunctions {
		index := hashFunction(element) % b.size
		b.data[index] = 1
	}
	b.insertionCount += 1
}

func (b *BloomFilters) search(element string) string {
	for _, hashFunction := range b.hashFunctions {
		index := hashFunction(element) % b.size
		if b.data[index] == 0 {
			return "Not in bloom filter"
		}
	}

	// (1 - ((1 - 1/m) ** (n * k))) ** k
	// m -> array size; n -> elements inserted; k -> no of hash functions

	m := b.size
	n := b.insertionCount
	k := len(b.hashFunctions)

	inner := math.Pow((1 - 1/float64(m)), float64(n*k))
	probability := math.Pow(1-inner, float64(k))
	return fmt.Sprintf("Found with a probability of: %f", probability)
}

func DoWork() {
	bloom := NewBloomFilters()
	bloom.insert("hello")
	bloom.insert("bloom")
	bloom.insert("filters")

	fmt.Printf("Bloom filters inserted with %d filters\n", bloom.insertionCount)
	fmt.Println("------")

	termList := []string{"hello", "ammsmsm", "filters", "filtersss"}
	for _, term := range termList {
		fmt.Printf("%s: %s\n", term, bloom.search(term))
	}
}
