# Go Language

**This README will only contain stuff which are either non-obvious or different from C++, along with a few basic syntaxes and definitions, since I personally am well-versed with C++**

## Introduction

Unlike Java, Packages are not hierarchical
> Eg. `import "math"` does not include `import "math/cmplx"`
```go
package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("The time is", time.Now())
    // fmt is mainly for scan and print
}
```

### Defining variables:
```go
var xyz int
var abc, pqr int = 4, 6
var def = "hi" //type can be omitted, variable def will take the type of "hi" which is a string
ghi := 7 // you can use the := operator to directly ignore the 'var' keyword, this is not allowed outside functions
```

Types in go: int, int8/16/32/64, uint, uint8/16/32/64, uintptr, float32/64, complex64/128, string, byte, rune

Typecasting is same as that in C, but it is compulsory in Go
Eg.
```go
var f float64 = float64(i)
```

### Basic flow statements syntax:
```go
for i := 0; i < 10; i++ {
	sum += i
}

abc := 1
for abc < 1000{  //This is like while() of C/C++
    abc += abc
} 

if x < 0 {
}
```

Some weird behavior:
Variables can be declared inside an ```if``` short statement
Calls to functions return their results before call to fmt.Println()
What do I mean by the above statement? Here's an example:

The output for the following piece of code:
```go
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("Parth\n")
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
func main() {
	fmt.Println(
		"hi2\n",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
would be this:
```go
Parth
27 >= 20
hi2
 9 20  //ya the initial single space character before 9 was also intentional
```

Switch statements are kinda like C/C++, except that Go runs only the selected cases (and not all that follow), hence removing the need of the break statement
Also, Go's switch statements need not be constants or integers

A defer statement puts the execution of a function on hold, until the surrounding function returns. Deferred function calls are pushed on a stack, and are executed at the end in LIFO order

Pointers are like C/C++, but there's no pointer arithmetic in Go

Struct fields are like C/C++, and are accessed using dots. They can also be accessed via pointers

Arrays are initialized like : `var arr [10]int`, and are of fixed length like C
But 'Slices' are dynamic sized arrays :-
```go
primes := [6]int{2, 3, 5, 7, 11, 13}
var s []int = primes[1:4] //This means index 1,2,3; which is {3,5,7}, default values are assumed if any of the limits are omitted
```
Slices are not a separate datatype, they just represent a section of an underlying array
So changes in slices actually change the array, which obc changes other slices involved

We can also directly initialize slices: ```sli := []int{2,3,4,5}```, this essentially just creates an array and builds a slice on it

We can extend a slice from right, unless it lies within the underlying array's capacity
but we cannot access the leftward elements of the slice using that slice
len() is length of slice, cap() is capacity of underlying array
Eg. 
```go
s := []int{0,1,2,3,4,5,6,7} //cap = 8, len = 8
s = s[2:5] // s = {2,3,4} :- cap = 6, len = 3
s = s[:5] // s = {2,3,4,5,6}, this is possible since it lies within capacity :- cap = 6, len = 5
//cannot access 0,1 whatsoever using this s
//it could've been possible if initial array was different from a slice s
orig := []int{0,1,2,3,4,5,6,7}
s = orig[2:5]
//now ofc orig still has the 0,1 with it
```

`make()` function:
Allocates a zeroes array and returns a slice referring to that array
```go
a := make([]int, x, y) //x is length of the slice a, y is its capacity; y is optional
```

`s = append(s, 67, 69)` appends the elements to the slice s, replacing the underlying array elements after end of slice s, if any

`range` function
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
	for i, v := range pow { // use _ instead, if u dont wanna use either index or value
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```
This outputs : 2\*\*0 = 1 \n 2\*\*1 = 2 \n 2\*\*2 = 4 \n ....

Maps:
How to initialize? :-
`m := make(map[string]int)` or `var m map[string]int; m = make(map[string]int)`
To insert or update elemts just `m[key] = elem`; To delete: `delete(m, key)`; To retrieve: `abc := m[key]`; and to check presence: `abc, flag := m[key]`


### Functions
Functions are values too, can be passed around like other values
Similar to C/C++ kinda

Closures are functions that reference variables from outside their body - those variables are bound to the closure
```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x // sum is captured
		return sum
	}
}
```

### Methods
A method is a function with a special receiver argument
Eg.
```go
type Vertex struct {
    X, Y float64
}
func (v Vertex) Abs() float64 { // Syntax: func receiver_type func_name return_type{}
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
Receivers can be pointers too, if u wanna use the method to modify values using pass by reference instead of pass by value
Receivers are called using a dot, like `a.Abs()`
If a method has a pointer receiver, it interprets the statement v.Scale(5) and (&v).Scale(5) automatically for convenience, unlike functions

### Interfaces
An interface type is a set of method signatures
A value of interface type can hold any value that implements those methods
Implementation is implicit - no "implements" keyword needed, just define the methods

```go
type Abser interface {
    Abs() float64
}

func (f MyFloat) Abs() float64 { //MyFloat implements Abser automatically
    return float64(f)
}
```

Empty interface `interface{}` can hold values of any type since every type implements at least zero methods
Useful for handling values of unknown type

Type assertions provide access to an interface value's underlying concrete value
```go
t := i.(T) // panics if i doesn't hold a T
t, ok := i.(T) // ok is false if i doesn't hold a T, doesn't panic
```

Type switches are like regular switches but the cases specify types not values
```go
switch v := i.(type) {
case int:
    // v is an int
case string:
    // v is a string
}
```

### Concurrency

Goroutines are lightweight threads managed by the Go runtime
Just put `go` before a function call to run it in a new goroutine
```go
go f(x, y, z) // starts new goroutine running f
```

Channels are typed conduits through which you can send and receive values with the channel operator `<-`
```go
ch := make(chan int) // creates channel
ch <- v    // Send v to channel ch
v := <-ch  // Receive from ch, assign to v
```

By default sends and receives block until the other side is ready - this allows goroutines to sync without explicit locks

Channels can be buffered - provide buffer length as second arg to make
```go
ch := make(chan int, 100) // buffered channel with capacity 100
```
Sends block only when buffer is full, receives block when buffer is empty

A sender can `close(ch)` a channel to indicate no more values will be sent
Receivers can test if channel closed: `v, ok := <-ch` where ok is false if closed
Only sender should close, never receiver

`range` can loop over channels, receiving values until channel is closed
```go
for i := range ch {
    // receives values until ch is closed
}
```

`select` lets a goroutine wait on multiple communication operations
It blocks until one of its cases can run, then executes that case
```go
select {
case msg := <-ch1:
    // received from ch1
case msg := <-ch2:
    // received from ch2
default:
    // runs if no other case is ready (makes select non-blocking)
}
```

`sync.Mutex` provides mutual exclusion - use Lock() and Unlock() to protect shared data
`sync.WaitGroup` waits for a collection of goroutines to finish - use Add(), Done(), and Wait()