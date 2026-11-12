package main

import (
	// "fmt"
	"math/rand"
	"testing"
)

func genRandSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}

	return slice
}

func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenRandSlice(t *testing.T) {
	size := 100
	slice := genRandSlice(size)

	if len(slice) != size {
		t.Errorf("expected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for b.Loop() {
		genRandSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := genRandSlice(1000)
	
	for b.Loop() {
		Sum(slice)
	}
}

// func Add(a, b int) int {
// 	return a + b
// }

// func BenchmarkAdd(b *testing.B) {
// 	for range b.N {
// 		Add(2, 3)
// 	}
// }

// func TestAddSubTests(t *testing.T) {
// 	tests := []struct {a, b, expected int}{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}

// 	for _,test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			res := Add(test.a, test.b) 
// 				if res != test.expected {
// 					t.Errorf("result = %d; want %d", res, test.expected)
// 				}
// 		})
// 	}
// }

// func TestAddTable(t * testing.T) {
// 	tests := []struct{a, b, expected int}{
// 		{2, 3, 5},
// 		{0, 0, 0},
// 		{-1, 1, 0},
// 	}	

// 	for _,test := range tests {
// 		res := Add(test.a, test.b)

// 		if res != test.expected {
// 			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, res, test.expected)
// 		}
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := Add(2, 3)
// 	expected := 5

// 	if result != expected {
// 		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
// 	}
// }
