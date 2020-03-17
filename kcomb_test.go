package kcomb

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestCombineGenerator(t *testing.T) {
	cases := []struct {
		columns  []Set
		expected []Set
	}{
		{
			columns: []Set{
				Set{Datum{"apple"}, Datum{"orange"}},
				Set{Datum{"celery"}, Datum{"broccoli"}},
			},
			expected: []Set{
				Set{Datum{"apple"}, Datum{"celery"}},
				Set{Datum{"apple"}, Datum{"broccoli"}},
				Set{Datum{"orange"}, Datum{"celery"}},
				Set{Datum{"orange"}, Datum{"broccoli"}},
			},
		},
		{
			columns: []Set{
				Set{Datum{1}, Datum{2}},
				Set{Datum{3}, Datum{4}},
			},
			expected: []Set{
				Set{Datum{1}, Datum{3}},
				Set{Datum{1}, Datum{4}},
				Set{Datum{2}, Datum{3}},
				Set{Datum{2}, Datum{4}},
			},
		},
		{
			columns: []Set{
				Set{Datum{1}, Datum{2}, Datum{3}},
				Set{Datum{4}, Datum{5}, Datum{6}, Datum{7}},
				Set{Datum{8}},
			},
			expected: []Set{
				Set{Datum{1}, Datum{4}, Datum{8}},
				Set{Datum{1}, Datum{5}, Datum{8}},
				Set{Datum{1}, Datum{6}, Datum{8}},
				Set{Datum{1}, Datum{7}, Datum{8}},
				Set{Datum{2}, Datum{4}, Datum{8}},
				Set{Datum{2}, Datum{5}, Datum{8}},
				Set{Datum{2}, Datum{6}, Datum{8}},
				Set{Datum{2}, Datum{7}, Datum{8}},
				Set{Datum{3}, Datum{4}, Datum{8}},
				Set{Datum{3}, Datum{5}, Datum{8}},
				Set{Datum{3}, Datum{6}, Datum{8}},
				Set{Datum{3}, Datum{7}, Datum{8}},
			},
		},
	}

	for idx, tc := range cases {
		done := make(chan struct{})
		stream := CombineGenerator(done, tc.columns)

		streamIdx := 0
		for v := range stream {
			if reflect.DeepEqual(v, tc.expected[streamIdx]) == false {
				t.Fatalf("Case Index: %d, Stream Index: %d - Expected '%v' to equal '%v'", idx, streamIdx, v, tc.expected[streamIdx])
			}
			time.Sleep(time.Nanosecond)
			streamIdx++
		}
	}
}

func TestCombine(t *testing.T) {
	cases := []struct {
		columns  []Set
		expected []Set
	}{
		{
			columns: []Set{
				Set{Datum{"apple"}, Datum{"orange"}},
				Set{Datum{"celery"}, Datum{"broccoli"}},
			},
			expected: []Set{
				Set{Datum{"apple"}, Datum{"celery"}},
				Set{Datum{"apple"}, Datum{"broccoli"}},
				Set{Datum{"orange"}, Datum{"celery"}},
				Set{Datum{"orange"}, Datum{"broccoli"}},
			},
		},
		{
			columns: []Set{
				Set{Datum{1}, Datum{2}},
				Set{Datum{3}, Datum{4}},
			},
			expected: []Set{
				Set{Datum{1}, Datum{3}},
				Set{Datum{1}, Datum{4}},
				Set{Datum{2}, Datum{3}},
				Set{Datum{2}, Datum{4}},
			},
		},
		{
			columns: []Set{
				Set{Datum{1}, Datum{2}, Datum{3}},
				Set{Datum{4}, Datum{5}, Datum{6}, Datum{7}},
				Set{Datum{8}},
			},
			expected: []Set{
				Set{Datum{1}, Datum{4}, Datum{8}},
				Set{Datum{1}, Datum{5}, Datum{8}},
				Set{Datum{1}, Datum{6}, Datum{8}},
				Set{Datum{1}, Datum{7}, Datum{8}},
				Set{Datum{2}, Datum{4}, Datum{8}},
				Set{Datum{2}, Datum{5}, Datum{8}},
				Set{Datum{2}, Datum{6}, Datum{8}},
				Set{Datum{2}, Datum{7}, Datum{8}},
				Set{Datum{3}, Datum{4}, Datum{8}},
				Set{Datum{3}, Datum{5}, Datum{8}},
				Set{Datum{3}, Datum{6}, Datum{8}},
				Set{Datum{3}, Datum{7}, Datum{8}},
			},
		},
	}

	for idx, tc := range cases {
		r := Combine(tc.columns)
		t.Logf("Index: %d - Created %d permutations", idx, len(r))
		if reflect.DeepEqual(r, tc.expected) == false {
			t.Fatalf("Index: %d - Expected '%v' to equal '%v'", idx, r, tc.expected)
		}
	}
}

func BenchmarkCombine2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkCombine(2, b)
	}
}

func BenchmarkCombine10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkCombine(10, b)
	}
}

func BenchmarkCombine50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkCombine(50, b)
	}
}
func BenchmarkCombine1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkCombine(1000, b)
	}
}

func benchmarkCombine(n int, b *testing.B) {
	col1 := make(Set, n)
	col2 := make(Set, n)
	col3 := make(Set, n)

	for i := 1; i < n; i++ {
		col1[i] = Datum{i}
		col2[i] = Datum{i}
		col3[i] = Datum{i}
	}

	log.Printf("created %d permutations", len(Combine([]Set{col1, col2, col3})))
}
