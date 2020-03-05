package kcomb

// Datum represents a single value within a set.
type Datum struct {
	Value interface{}
}

// Set represents a column (or set) of data of a similar kind.
type Set []Datum

// Combine will generate every distinct permutation of values within a
// collection of sets.
//
// Example:
//   argument: [ [ apple, orange ], [ celery, broccoli ] ]
//   output: => [
//                [ apple, celery ],
//                [ apple, broccoli ],
//                [ orange, celery ],
//                [ orange, broccoli ]
//              ]
func Combine(columns []Set) []Set {
	n := len(columns)
	indices := make([]int, n)
	combset := make([]Set, 0)

	for {
		var comb Set

		for i := 0; i < n; i++ {
			comb = append(comb, columns[i][indices[i]])
		}

		combset = append(combset, comb)
		next := n - 1

		for next >= 0 && (indices[next]+1 >= len(columns[next])) {
			next--
		}

		if next < 0 {
			return combset
		}

		indices[next]++

		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}
