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

## Loops
Arrays, slices and maps can be iterated as follows:

```golang
x := []string{"Hello", "Go"}
for index, value := range x {
	fmt.Printf("%i -> %s", index, value)
}
```

Otherwise, loops can be created as follows:
```golang
for i := 1; i <= 5; i++ {
	
}
```

## Data structures

### Variables


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

### Maps
Maps can be used for assigning a key to a value and can be created as follows:
```golang
m := map[string]string
```

### Pointers
Basic usage of pointers as follows:

```golang
type User struct {
	Name string
}

pointer := &User {Name: "Martin"}
```

- Pointers can be created by calling using the `&`-symbol in front of an object
- Passing or returning values as pointers can be signaled by using the asterisk symbol `*User`
