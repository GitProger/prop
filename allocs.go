package allocs

type User struct {
    ID     uint64
    Name   string
    Age    int
    bytes  [4000]byte
}

func Adult(u *User) bool {
    return u.Age >= 18
}


func NewUser(c Counter, name string, age int) *User {
    c.Succ()
    return &User{ID: c.Get(), Name: name, Age: age}
}

func OldFoo(c Counter) *User {
    u1 := NewUser(c, "Alice", 19)
    _ = Adult(u1)

    u2 := NewUser(c, "Bob", 17)
    return u2
}


func heap_NewUser(c Counter, name string, age int) *User {
    return NewUser(c, name, age)
}

func stack_NewUser(_buf *User, c Counter, name string, age int) *User {
    c.Succ()
    _buf.ID = c.Get()
    _buf.Name = name 
    _buf.Age = age
    return _buf
}

func NewFoo(c Counter) *User {
    var _u User
    u1 := stack_NewUser(&_u, c, "Alice", 19)
    _ = Adult(u1)

    u2 := heap_NewUser(c, "Bob", 17)
    return u2
}
