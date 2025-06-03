package internal

import "slices"

func CompareStrings(args1, args2 []string) bool {
	if len(args1) != len(args2) {
		return false
	}
	slices.Sort(args1)
	slices.Sort(args2)
	return slices.Equal(args1, args2)
}
