Golang proposal regarding reduce of memory alloc call count

Assume me have such struct:
```go 
type User struct { // 32 bytes (rather small)
    ID     uint64  // 8 bytes
    Name   string  // 16 bytes
    Age    int     // 4 bytes + 4 bytes alignment
}
```
or just
```go
type SomeSmallStruct struct{...}
```


Let's say I write something like this:
```go
func NewStruct() *SomeSmallStruct {
    return &SomeSmallStruct{...}
}

func foo(...) ... {
    v := NewStruct()
    vOut := NewStruct()
    ... // v does not go out of foo
        // vOut goes out of foo
}
```


What I should get generated in fact by the compiler, to prevent unnessasary heap allocs:
```go
// _variableName           : means implicit hidden variable
// word_SomeIdentifierName : means that there was just SomeIdentifierName in the source code
//                           and word_ is a prefix added by the compiler*
// compiler* means hypothetical Go compiler which adds heap_ and stack_ functions

func heap_NewStruct() *SomeSmallStruct {
    return &SomeSmallStruct{...}
}

func stack_NewStruct(_buffer *SomeSmallStruct) *SomeSmallStruct {
    // here ~smth like (in c++):
    //     *_buffer = std::move(SomeSmallStruct{...})
    // or
    //     new (_buffer) SomeSmallStruct({...})

    // or in go:
    _buffer.field1 = ... // corresponding value
    _buffer.field2 = ...
    ...
    return _buffer
}

func foo(...) ... {
    var _buf SomeSmallStruct
    v := stack_NewStruct(&_buf) // `stack_NewStruct` choosen by the compiler instead of `NewStruct`
    ...                         //      as hs pointer does not leave `foo` function
    vOut := heap_NewStruct()    // `NewStruct_Heap` choosen by the compiler
    ...
}

```

Results (go1.22.1):
```
goos: linux
goarch: amd64
pkg: allocs
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkOldFoo-8       14548387               132.9 ns/op            64 B/op          2 allocs/op
BenchmarkNewFoo-8       15364522                79.40 ns/op           32 B/op          1 allocs/op
PASS
ok      allocs  3.323s
```
