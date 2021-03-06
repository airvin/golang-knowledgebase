GoLang Notes
==========

These notes are taken from the [A Tour of Go](https://tour.golang.org/list) tutorial series. 

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

Function closures
-------
A closer is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function below returns a closure and each closure is bound to its own `sum` variable.

```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	
	return func(x int) int {
		fmt.Printf("sum is: %d\n",sum)
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```


Methods
-------
Methods can be defined on types (either struct or non-struct). A method is a function with a special receiver argument. the receiver appears in its own argument list between the `func` keyword and the method name. 

```
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

v.Abs()
```

Methods can only be declared with a receiver whose type is defined in the same package as the method. Methods cannot be declared with a receiver whose type is define in another package (including in-built types such as `int`)

Pointer receivers
-------
Methods can be declared with pointer receivers. This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)
For example, the `Scale` method here is defined on `*Vertex`.

```
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

v.Scale(10)
```

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Methods and pointer indirection
-------
Methods with pointer receivers take either a value or a pointer as the receiver when they are called:
```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
This is because Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

The equivalent happens in the reverse direction in that methods with value receivers take either a value or a pointer as the receiver when they are called:

```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

Value vs pointer receivers
-------
There are two reasons to use a pointer receiver.
1. So that the method can modify the value that its receiver points to.
2. To avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example. In this case, the method needn't modify its receiver.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.

Interfaces
-------
An interface type is defined as a set of method signatures. A value of interface type can hold any value that implements those methods. 
Interfaces don't allow pointer indirection. For example:

```
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	
	v := Vertex{3, 4}

	a = &v // a *Vertex implements Abser
	p := v

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	// these will print the same thing as p.Abs() is interpreted as (&p).Abs(). This is not allowed with a
	fmt.Println(a.Abs())
	fmt.Println(p.Abs())
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Interface values
-------
Interface values can be thought of as a tuple of a value and a concrete type:
```
(value,type)
```

Interface values will nil underlying values
-------
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver. 

Note that an interface value that holds a nil concrete value is itself non-nil.

Nil interface values
-------
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

```
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

The empty interface
-------
The interface type that specifies zero methods is known as the empty interface:
```
interface{}
```
An empty interface may hold values of any type (as every type implements at least zero methods).

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

Type assertions
------

A type assertion provides access to an interface value's underlying concrete value.
```
t := i.(T)
```
This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.
If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
```
t, ok := i.(T)
```
If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Type switches
-------
A type switch is a construct that permits several type assertions in series. 
A type switch is like a regular switch statment, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value. 
```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.
This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

Stringer
-------
One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.
```
type Stringer interface {
    String() string
}
```
A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

Errors
------
Go programs express error state with `error` values.

The `error` type is a built-in interface similar to `fmt.Stringer`:
```
type error interface {
    Error() string
}
```
As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.
```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```
A nil `error` denotes success; a non-nil `error` denotes failure.

Goroutines
-------
A goroutine is a lightweight thread managed by the Go runtime.
`go f(x, y, z)` starts a new goroutine running `f(x, y, z)`
The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.
Goroutines run in the same address space, so access to shared memory must be synchronised. The `sync` package provides useful primitives. 

Channels
------
Channels are a typed conduit through which you can send and receive values with the channel operator `<-`
```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```

Further resources
-------
[How to Write Go Code](https://golang.org/doc/code.html)<br/>
[Writing, building, installing, and testing Go code screencast](https://www.youtube.com/watch?v=XCsL89YtqCs)<br/>
[Command go](https://golang.org/cmd/go/)<br/>
[Effective Go](https://golang.org/doc/effective_go.html)<br/>
[Error handling and Go](https://blog.golang.org/error-handling-and-go)<br/>
[Google I/O 2012 - Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)<br/>
[Google I/O 2013 - Advanced Go Concurrency Patterns](https://www.youtube.com/watch?<br/>v=QDDwwePbDtw)<br/>
[GO Frequently Asked Questions (FAQ)](https://golang.org/doc/faq)<br/>
[The Go Memory Model](https://golang.org/ref/mem)<br/>
[Codewalk: Generating arbitrary text: a Markov chain algorithm](https://golang.org/doc/codewalk/markov/)<br/>
[Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)<br/>
[JSON and Go](https://blog.golang.org/json-and-go)<br/>
[Diagnostics](https://golang.org/doc/diagnostics.html)<br/>
[The Go Programming Language Specification](https://golang.org/ref/spec)<br/>
[Codewalk: Share Memory By Communicating](https://golang.org/doc/codewalk/sharemem/)<br/>
[Go: a simple programming environment](https://vimeo.com/53221558)<br/>
[Writing Web Applications](https://golang.org/doc/articles/wiki/)<br/>
[Codewalk: First-Class Functions in Go](https://golang.org/doc/codewalk/functions/)<br/>