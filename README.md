GoLang Notes
==========
Pointers
--------
A pointer holds the memory address of a value. 
The type `*T` is a pointer to a T value (e.g. type `*int` is a pointer to an integer). 
The zero value of a point is `nil`

To generate a pointer:
```
var p *int
```
or 
```
i := 42
p := &i
```

The `*` operator denotes the pointer's underlying value
```
*p = 21 // set i through the pointer p
```
This is known as "dereferencing" or "indirecting"

Important: Go has no pointer arithmetic

Structs
------
A `struct` is a collection of fields. For example:
```
type Vertex struct {
	X int
	Y int
}
```
The fields can be accessed with a dot.
```
v := Vertex{1,2}
v.X = 4 // sets X int as 4
```
To set a pointer to a struct:
```
p := &v
```
Fields can then be accessed through either `(*p).X` or `p.X`

Arrays
------

Slices
------

Range
------

Maps
-----
Maps can be created via a map literal or using the `make` function.

Map literal
```
var m = map[key type]value type{
	key: value,
	key: value,
}
```

Make function

```
m := make(map[key type]value type)
m[key] = value
```

Insert or update elements
```
m[key] = value
```
Retrieve an element
```
elem = m[key]
```
Delete an element
```
delete(m, key)
```
Test that a key is present
```
elem, ok := m[key]
```
If `key` is not present, then `elem` is zero and `ok` is `false`

Function values
-------
Functions are values that can be passed around like other values, e.g. as function arguments and return values.

