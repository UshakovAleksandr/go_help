package main

import "testing"

//func BenchmarkToCompare(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		//OldFillSlice()
//		NewFillSlice()
//	}
//}
//
func BenchmarkOldFillSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OldFillSlice()
	}
}

//func BenchmarkNewFillSlice(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		NewFillSlice()
//	}
//}


