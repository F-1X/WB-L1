package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.
func main() {
	set1 := []int{4, 5, 1, 2, 3, 1, 2, 4}
	set2 := []int{1, 2, 1, 5, 6, 1, 2, 3}

	intersectSlice := sliceSet(set1, set2)
	fmt.Println("intersectSlice:", intersectSlice)

	intersectMap := mapSet(set1, set2)
	fmt.Println("intersectMap:", Keys(intersectMap))

	intersectSet := Intersection(NewSet(set1...), NewSet(set2...))
	fmt.Println("intersectSet:", Keys(intersectSet))
}

func sliceSet(set1, set2 []int) []int {
	intersection := []int{}
	for i := 0; i < len(set1); i++ {
		for j := 0; j < len(set2); j++ {
			if set1[i] == set2[j] {

				alreadyExist := false
				for k := 0; k < len(intersection); k++ {
					if intersection[k] == set1[i] {
						alreadyExist = true
						break
					}
				}

				if !alreadyExist {
					intersection = append(intersection, set1[i])
				}

				if len(intersection) == 0 {
					intersection = append(intersection, set1[i])
				}
			}
		}
	}
	return intersection
}

func mapSet(set1, set2 []int) map[int]struct{} {
	intersection := make(map[int]struct{})
	for i := 0; i < len(set1); i++ {
		for j := 0; j < len(set2); j++ {
			if set1[i] == set2[j] {
				intersection[set1[i]] = struct{}{}
			}
		}
	}

	return intersection
}

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements ...T) Set[T] {
	set := make(Set[T])
	for _, e := range elements {
		set[e] = struct{}{}
	}
	return set
}

func Intersection[T comparable](set1, set2 Set[T]) Set[T] {
	intersection := make(Set[T])
	for e := range set1 {
		if _, found := set2[e]; found {
			intersection[e] = struct{}{}
		}
	}
	return intersection
}

func Keys[T comparable](m map[T]struct{}) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
