# GoLang Basics

<img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" width="200" />

## Functions

Functions can be defined as follows:

```golang
func User(name string) *User {
    u := User{Name: "Martin", Mail: "bla@github.com"}
    u.Age = 42
    return &u
}
```

## Data structures

### Structs
Structs are a combination of multiple variables and can be defined as follows:

```golang
type User struct {
	Name string
	Age int16
	Mail string
}
```

Remarks:

- Attributes starting with upper case are `public`
- Attributes with lower case are `private`

### Arrays
Arrays can be created by using 
```golang
array := [6]bool{true, false, true, true, false, true}
```

- Have the attribute `len(array)` and cannot be extended

### Slices
Slices can be created by using
 
```golang
slice := []bool{true, false, true, true, false, true}
```

- Have the attributes `len(slice)` and `cap(slice)`
- Can be extended with `append(slice, values...)` which is returning the new slice

### Pointers
