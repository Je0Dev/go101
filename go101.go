package main

import (
	"fmt"
	"time"
	"math/rand"
	"math/cmplx"
	"runtime"
)

var c, python, java bool//variables declaration
var i, j int = 1, 2//global init

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.pi)

	fmt.Println(add(42, 13))//call add()

	a, b := swap("hello", "world")
	fmt.Println(a, b)//world,hello
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)//1 int,3 bools
	var c, python, java = true, false, "no!"//init
	fmt.Println(i, j, c, python, java)

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}


func add(x int, y int) int {
	return x + y
}
//or 
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Basic Types:
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
//      // represents a Unicode code point
// float32 float64
// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1//shift
	z      complex128 = cmplx.Sqrt(-5 + 12i)//complex
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//zero val->0:num,false:bool,"":strings

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
//or simpler form
i := 42 // int
f := 3.142 // float64
g := 0.867 + 0.5i // complex128
f := float64(i)
u := uint(f)

var i int
j := i // j is int

v := 42 // change me!
fmt.Printf("v is of type %T\n", v)

}


const Pi = 3.14//char,str,bool,num
//cannot be declared with :=

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}


const (
	// huge num by shifting a 1 bit left 100 places.
	// the binary num that is 1 followed by 100 0's.
	Big = 1 << 100
	// Shift right again 99 places, so we end up with 1<<1= 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

//the init statement: executed before the first iteration
//the condition expression: evaluated before every iteration
//the post statement: executed at the end of every iteration
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//or 
	for ; sum < 1000; {
		sum += sum
	}
	//drop ; for is go's while!
	for sum < 1000 {
		sum += sum
	}
	//infinity & behond
	for {
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	//// can't use v here, though
	return lim
}


func main() {
	//calls
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func main() {
	//only runs the selected case
	//automatic break
	// Go's switch cases need not be constants, 
	// and the values involved need not be integers.
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

switch i {
case 0:
case f():
}//if(i==0) doesnt call

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defers the execution of a function until the surrounding function returns
	//The deferred call's arguments are evaluated immediately, but the function call is
	//  not executed until the surrounding function returns
	
	defer fmt.Println("world")

	fmt.Println("hello")


	//Deferred function calls are pushed onto a stack. 
	// When a function returns, 
	// its deferred calls are executed in last-in-first-out order
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//pointer holds the memory address of a value. 

var p *int
i := 42
p = &i
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p(dereference-indirect)



import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

//collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	p := &v
	p.X = 1e9//(*p).X
	fmt.Println(v)
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
var a [10]int//array declare-fixed

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
//slice:a[low : high] like a[1:4]-1 through 3

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

//slice doenst store-describers section of underlying arr
//Other slices that share the same underlying array will see those changes(if changing elem)
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	
}


//[]bool{true, true, false},[3]bool{true, true, false}

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

//var a [10]int->
// a[0:10]
// a[:10]
// a[0:]
// a[:] are all equivalent

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	//A slice has both a length and a capacity. 

	//The length of a slice is the number of elements it contains. 

	//The capacity of a slice is the number of elements in the
	//  underlying array, counting from the first element in the slice. 

	//You can extend a slice's length by re-slicing it,
	//  provided it has sufficient capacity.

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//zero val of slice is nil->length and capacity of 0,no underly arr
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

a := make([]int, 5)  // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4


func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

//When ranging over a slice, two values are returned 
// for each iteration. The first is the index, 
// and the second is a copy of the element at that index. 
func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//skip index or value
for i, _ := range pow
for _, value := range pow

//keep only index
for i := range pow

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}


//A map maps keys to values
//The zero value of a map is nil. 
// A nil map has no keys, nor can keys be added


type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//The make function returns a map of the given type, 
	// initialized and ready for use. 
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)//keys required
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
//or
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

m[key] = elem//insert/update elem
elem = m[key]//retrieve elem
delete(m, key)//delete elem
elem, ok = m[key]//test if key present->it is?ok=true:false
elem, ok := m[key]//declare


func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

//Functions are values-used as func args or ret vals
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}


//A closure is a function value 
// that references variables from outside its body.

//The function may access and assign to the referenced variables; 
// in this sense the function is "bound" to the variables.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}//each closer bound to its own sum val

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

type Vertex struct {
	X, Y float64
}
//method->function with special receiver arg

//Receiver appears in its own argument list between the func keyword 
// and the method name

//Abs method has a receiver of type Vertex named v. 
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//regular function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64
//You cannot declare a method with a receiver whose type is defined 
// in another package (which includes the built-in types such as int)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
//Methods with pointer receivers 
// can modify the value to which the receiver points 
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v := Vertex{3, 4}
	v.Scale(10)//change here
	fmt.Println(v.Abs())
}

//again as regular functions
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)//hm it changes here
	fmt.Println(Abs(v))
}

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

//method with ptr receivers
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK


func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)//(&v).Scale(2) 
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!

//method with val receivers
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK-(*p).Abs(). 


func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}


//use ptr receiver->
//method can modify the value that its receiver points to. 
// avoid copying the value on each method call(large struct eg)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}


//An interface type is defined as a set of method signatures. 
//A value of interface type can hold any value that implements those methods. 


type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser->error
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


//A type implements an interface by implementing its methods.
type I interface {
	M()
}

type T struct {
	S string
}

// method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

//interface->(value, type)
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
//an interface value that holds a nil concrete value is itself non-null(huh?)
func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


//A nil interface value holds neither value nor concrete type. 
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()//run-time error because there is no type inside the interface 
	// tuple to indicate which concrete method to call


	var i interface{}//empty->vals of any type
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe(i interface{}) {//empty interface mentioned
	fmt.Printf("(%v, %T)\n", i, i)
}


//t, ok := i.(T) if i holds T=>t underlying val & ok=true,else ok=F & t=0


var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)//no panic
	fmt.Println(s, ok)

	f, ok := i.(float64)//no panic
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}

func do(i interface{}) {
	// interface value i holds a value of type T or S. 
	// In each of the T and S cases, the variable v will be of type T or S 
	// respectively and hold the value held by i. 
	// In the default case (where there is no match), 
	// the variable v is of the same interface type and value as i
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}


//type that can describe itself as a string
type Stringer interface {
    String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

//build-in
type error interface {
    Error() string
}
//A nil error denotes success; a non-nil error denotes failure. 
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

import (
	"fmt"
	"io"
	"strings"
)

//io.Reader interface
func (T) Read(b []byte) (n int, err error)
//Read populates the given byte slice with data and returns 
// the number of bytes populated and an error value. 
// It returns an io.EOF error when the stream ends. 
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}


package image

//from package image
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}


import (
	"fmt"
	"image"
)

func main() {
	//drawing stuff
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}


package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}


// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}


// goroutine is a lightweight thread by go runtime
//Goroutines run in the same address space, so access to shared memory must be synchronized



import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

//Channels are a typed conduit through which you can send and receive values  with <-
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	//must be declared before use
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

//buffered with length
//sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. 
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//no more values will be sent
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
//v, ok := <-ch test whether a channel has been closed


//select statement lets a goroutine wait on multiple communication operations. 
//select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready. 
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}


func main() {
	start := time.Now()
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
// // select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }


import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
