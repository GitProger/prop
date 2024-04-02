package allocs

import "testing"

func BenchmarkOldFoo(b *testing.B) {
    c := NewSimpleCnt()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        u := OldFoo(c)
        if (u.Age >= 18) != Adult(u) {
            b.FailNow()
        }
    }
}

func BenchmarkNewFoo(b *testing.B) {
    c := NewSimpleCnt()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        u := NewFoo(c)
        if (u.Age >= 18) != Adult(u) {
            b.FailNow()
        }
    }
}
