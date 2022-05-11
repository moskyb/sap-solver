package solver

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Side string

const (
	Left  Side = "left"
	Right Side = "right"
)

type Team struct {
	Side    Side
	Friends []*Pet
}

func (t *Team) CollectFainted() {
	faintedIdxes := make([]int, 0, len(t.Friends))
	for i, friend := range t.Friends {
		if friend.hasFainted {
			faintedIdxes = append(faintedIdxes, i)
			fmt.Printf("%v faints!\n", friend)
		}
	}

	for _, idx := range faintedIdxes {
		t.Friends = slices.Delete(t.Friends, idx, idx+1)
	}
}

func (t *Team) String() string {
	s := "["
	friends := t.Friends
	if t.Side == Left {
		friends = reverse(friends)
	}

	friendStrs := make([]string, len(t.Friends))
	for _, friend := range friends {
		friendStrs = append(friendStrs, friend.String())
	}

	s += strings.Join(friendStrs, " ")

	s += "]"
	return s
}

func reverse[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
