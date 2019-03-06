package shortest_lenth

import (
	"testing"
	"fmt"
)

func TestSl(t *testing.T) {
	s := []int{1, 2, 2, 3, 4, 4, 3, 2, 3, 3, 2, 2, 1,1,2,3,4}
	M := 4
	result := shortestLength(s, M)
	fmt.Println("result:", result)

}
