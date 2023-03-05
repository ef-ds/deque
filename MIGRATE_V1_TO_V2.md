# How to Migrate from Deque V1 to V2

After have downloaded Deque v2 version into your project (`go get -u github.com/ef-ds/deque/v2`), below code changes are required in order to migrate to v2.

- Update the imports path from `github.com/ef-ds/deque` to `deque github.com/ef-ds/deque/v2`
- Update the Deque's declaration from `var d deque.Deque` to `var d deque.Deque[data_type]` and/or `d := deque.New()` to `d := deque.New[data_type]()`
- Remove any `interface{}` type casts (i.e. `var v, _ := d.Pop(); typeCast := v.(data_type)`)

Below are examples of migrated code:

Simple int as data type - V1 version:
```go
import "github.com/ef-ds/deque"

var d deque.Deque
for i := 1; i <= 5; i++ {
    d.PushBack(i)
}

for d.Len() > 0 {
    v, _ := d.PopFront()
    intV := v.(int)
    fmt.Print(intV)
}
```

Simple int as data type - V2 version:
```go
import deque "github.com/ef-ds/deque/v2"

var d deque.Deque[int]
for i := 1; i <= 5; i++ {
    d.PushBack(i)
}

for d.Len() > 0 {
    intV, _ := d.PopFront()
    fmt.Print(intV)
}
```

Custom struct as data type - V1 version:
```go
import "github.com/ef-ds/deque"

type myType struct {
    value int
}

d := deque.New()
for i := 1; i <= 5; i++ {
    d.PushBack(&myType{value: i})
}

for d.Len() > 0 {
    v, _ := d.PopFront()
    myTypeV := v.(*myType)
    fmt.Print(myTypeV.value)
}
```

Custom struct as data type - V2 version:
```go
import deque "github.com/ef-ds/deque/v2"

type myType struct {
    value int
}

d := deque.New[*myType]()
for i := 1; i <= 15; i++ {
    d.PushBack(&myType{value: i})
}

for d.Len() > 0 {
    v, _ := d.PopFront()
    fmt.Print(v.value)
}
```

Mixed data types used as data type - V1 version:
```go
import "github.com/ef-ds/deque"

d := deque.New()
for i := 1; i <= 5; i++ {
    if i%2 == 0 {
        // Push value as int
        d.PushBack(i)
    } else {
        // Push value as string
        d.PushBack(strconv.Itoa(i))
    }
}

for d.Len() > 0 {
    v, _ := d.PopFront()
    switch t := v.(type) {
    case int:
        fmt.Printf("Int value: %d", t)
    case string:
        fmt.Printf("String value: %s", t)
    default:
        fmt.Print("Unrecognized type")
    }
}
```

Mixed data types used as data type - V2 version:
```go
import deque "github.com/ef-ds/deque/v2"

d := deque.New[interface{}]()
for i := 1; i <= 5; i++ {
    if i%2 == 0 {
        // Push value as int
        d.PushBack(i)
    } else {
        // Push value as string
        d.PushBack(strconv.Itoa(i))
    }
}

for d.Len() > 0 {
    v, _ := d.PopFront()
    switch t := v.(type) {
    case int:
        fmt.Printf("Int value: %d", t)
    case string:
        fmt.Printf("String value: %s", t)
    default:
        fmt.Print("Unrecognized type")
    }
}
```
