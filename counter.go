package allocs

type Counter interface {
    Succ()
    Get() uint64
}

type SimpleCnt struct {
    v uint64
}

func (s SimpleCnt) Get() uint64 {
    return s.v
}

func (s *SimpleCnt) Succ() {
    s.v++
}

func NewSimpleCnt() *SimpleCnt {
    return &SimpleCnt{v: 0}
}
