package st

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

type Creature struct {
	Species string
}

func (c Creature) String() string {
	return c.Species
}

func (c *Creature) Reset() {
	c.Species = ""
}

func ChangeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}

type Number interface {
	constraints.Integer | constraints.Float
}

func Max[T constraints.Ordered](x, y T) T {
	return Ternary(x > y, x, y)
}

func Ternary[T any](cond bool, x, y T) T {
	if cond {
		return x
	}
	return y
}

func isContains(n int, list []int) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}

func Sum[T Number](list []T) T {
	var res T
	for _, n := range list {
		res += n
	}
	return res
}

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accumulator(init, el)
	}
	return init
}

func Bianary_search(list []int, item int) int {
	var low, mid, guess int = 0, 0, 0
	high := len(list) - 1

	for low <= high {
		mid = low + high
		guess = list[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of range uint8 range", x)
}
