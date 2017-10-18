GoLang Notes
==========
Pointers
--------

Structs
------

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

