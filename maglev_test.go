package maglev

import (
	"fmt"
	"log"
	"testing"
)

// func TestPopulate(t *testing.T) {

// 	var tests = []struct {
// 		dead []int
// 		want []int
// 	}{
// 		{nil, []int{1, 0, 1, 0, 2, 2, 0}},
// 		{[]int{1}, []int{0, 0, 0, 0, 2, 2, 2}},
// 	}

// 	permutations := [][]uint64{
// 		{3, 0, 4, 1, 5, 2, 6},
// 		{0, 2, 4, 6, 1, 3, 5},
// 		{3, 4, 5, 6, 0, 1, 2},
// 	}

// 	for _, tt := range tests {
// 		if got := populate(permutations, tt.dead); !reflect.DeepEqual(got, tt.want) {
// 			t.Errorf("populate(...,%v)=%v, want %v", tt.dead, got, tt.want)
// 		}
// 	}
// }

const sizeN = 5
const lookupSizeM = 13 //need prime and

func TestDistribution(t *testing.T) {
	// const size = 10

	var names []string
	for i := 0; i < sizeN; i++ {
		names = append(names, fmt.Sprintf("backend-%d", i))
	}

	mm := NewMaglev(names, lookupSizeM)
	log.Println("kk1:", mm.permutation[0])
	mm.populate()
	log.Println("pp1:", mm.lookup)
	mm.populate()
	log.Println("pp2:", mm.lookup)
	log.Println("node1:", mm.Get("IP1"))
	log.Println("node2:", mm.Get("IP2"))
	log.Println("node3:", mm.Get("IPasdasdwni2"))
	if err := mm.Remove("backend-0"); err != nil {
		t.Error("Remove failed", err)
	}
	log.Println("node3-D:", mm.Get("IPasdasdwni2"))

	if err := mm.Remove("backend-1"); err != nil {
		t.Error("Remove failed", err)
	}
	log.Println("node2-D:", mm.Get("IP2"))
}
