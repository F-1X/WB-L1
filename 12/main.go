package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
// собственное множество.
func main() {

	testCases := []struct {
		input []string
		want  []string
	}{
		{[]string{"cat", "cat", "dog", "cat", "tree"}, []string{"cat", "dog", "tree"}},
	}

	for i := range testCases {
		// на дженериках
		got := genericsSet(testCases[i].input)
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", testCases[i].want) {
			fmt.Printf("testCase(%d) failed: genericsSet(%v) = %v (expected: %v)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: genericsSet(%v) == %v (result: %v)\n", i, testCases[i].input, got, testCases[i].want)
		}

		// на слайсах
		got2 := sliceSet(testCases[i].input)
		if !EqualsAnyOrder(got2, testCases[i].want) {
			fmt.Printf("testCase(%d) failed: sliceSet(%v) = %v (expected: %v)\n", i, testCases[i].input, got2, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: sliceSet(%v) == %v (result: %v)\n", i, testCases[i].input, got2, testCases[i].want)
		}

		// на типе мап
		got3 := mapSet(testCases[i].input)
		if !EqualsAnyOrder(Keys(got3), testCases[i].want) {
			fmt.Printf("testCase(%d) failed: mapSet(%v) = %v (expected: %v)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: mapSet(%v) == %v (result: %v)\n", i, testCases[i].input, Keys(got3), testCases[i].want)
		}
	}
}

func sliceSet(words []string) []string {
	set := []string{}

	for i := 0; i < len(words); i++ {
		contains := false
		for j := 0; j < len(set); j++ {
			if set[j] == words[i] {
				contains = true
				break
			}
		}

		if !contains {
			set = append(set, words[i])
		}

	}

	return set

}

func mapSet(words []string) map[string]struct{} {

	set := make(map[string]struct{})
	for i := 0; i < len(words); i++ {
		set[words[i]] = struct{}{}
	}
	return set
}

// определяем constraint
type MyType interface {
	string
}

type Set[T MyType] map[T]struct{}

// аналогично с решением map реализуем множетсов
func genericsSet[T MyType](elements []T) Set[T] {
	set := make(Set[T])
	for _, e := range elements {
		set[e] = struct{}{}
	}
	return set
}

func Keys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func EqualsAnyOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := map[string]struct{}{}
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			return false
		}
	}
	return true
}
