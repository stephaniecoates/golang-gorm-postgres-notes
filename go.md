# foundations of Go

- it's built around speed and concurrency
- imperative programming language, not obj-oriented, or logical, functional
- combination of C++ and Java, with syntax and style like Python

- double quotes are necessary for string data type
- single quotes are used for characters (strings are collections of characters)

- go is a compiled language
- before we can execute it, we need to translate it into a language the computer can understand (assembly? or some other lower level lang)

- go run automatically compiles and runs the file in one step
  - if you want to compile and run separately, type `go build file.go`, which will generate compiled file `file.exe` 
  - that executable file can then directly be run from the command line (literally just typing `./file.exe`) since it's in a lower level language the computer can directly understand


  # variables and data types
  - variable is a way to store and access information
  - once a variable is declared as a specific type, it CANT change
    - golang is statically typed
  - variables are declared as `var name type`, i.e. `var last_name string`
  - can be assigned in the same line, or later, i.e. `var storage int = 50`

  ## variable types

  - Integers
    - unsigned (not negative)
      - uint8 (8 bits/1 byte, numbers 0 - 255) 
      - uint16 (16 bits, numbers 0 - 65535) (~2^16)
      - uint 32 (32 bits, nums 0 - 4294967295) (~2^32)
      - uint 64 (64 bits, 0 - 18446744073709551615) (~2^64)
    - signed integers (both positive and negative)
      - (now, the first bit will represent whether the following int is + or -, and then it fans out equal directions from 0 space-wise)
      - int8(numbers -128 to 127)
      - int16 (numbers -32768 to 32767)
      - int 32 (nums -2147483648 to 2147483647)
      - int 64 (nums -9223372036854775808 to 9223372036854775807)
      - * a lot of people end up using int64 as a base since computers have so much memory nowadays that the extra allocation doesn't really make a difference. But if you want to write a very memory-efficient program, you can minimize how much memory you allocate to ints based on their bit size
    - 3 machine dependent types
      - if you just type `uint`, depending on the operating system you use, it will either allocate 32 or 64 bits based on the OS
      - `int` allocates same memory size as `uint`, which'll either be 32 or 64 bits based on OS
      - `uintptr` provides an unsigned integer to store the uninterpreted bits of a pointer value
  
  - Floating Point Numbers - `float32 / float64`
    - signifies numbers with a floating point (decimal)
    - float32 (32 bits to represent the float)
    - float64 (64 bits to represent the float, will be longer and more precise)

  - Complex (Imaginary Parts)
    - only used for advanced computational math or physics
    - complex32 (complex nums with float32 real and imaginary parts)
    - complex64 (complex nums with float64 real and imaginary parts)

- Strings (double quotes)
`"Hello World"`

- Characters (single quotes)
`'a'`

- Booleans
  - true
  - false


# explicit vs implicit, assignment expression
- with explicit, we ourselves tie a data type to our variable, `var number uint8 = 5`
- with implicit, we write `var number = 5` and we tell golang to "guess" and how to size/apply a type to our variable, based on the value we give it
- if you use implicit, you have less flexibility to change the variable down the road, i.e. if you first give a variable `num` the value 5, and go implicitly applies uint8 datatype, however if you try later to do `num + 500` to the original variable, it'll break. To avoid this, use explicit assignments.
  - this above scenario wouldn't happen though, because the computer I think would apply a machine dependent uint/int type, i.e. uint32 or uint64 (not uint8)
- In 99% of cases, it's fine to not define the type. But it's a good tool to have.
- *tip* `fmt.Printf("%T", num)` would print the data type of `num`

## expression assignment operator (walrus operator) :=
- essentially saying, do what we did above with `var number = 5`, but omit `var`
- only works with implicit assignment (can't define type)

- when you declare a variable but don't initialize/assign a value to it, it has a _default value_.
	- `var unsignedNumber uint` (0)
	- `var signedNumber int` (0)
	- `var floatingNum float32` (0)
	- `var bl bool` (false)
	- `var word string` ("")

# fmt module
- `fmt`
  - `.Printf()`
    - format a separate variable within a string
    - `%T` prints var type

    - `%v` prints value in default format
    - `%+v` prints value in default format with variable name

    - `%%` prints literal % sign (escaped value)

    - `%t` prints boolean value (true or false)

    ints
    - `%b` prints base2 (binary) number
    - `%o` prints base8 (octal) number
    - `%d` prints base10 (digit) number
    - `%x` prints base16 (hexadecimal) number
      - use `%X` for capital hexadecimal nums

    floating points
    - `%e` prints scientific notation (exponential form, i.e. 1.556413e+20)
    - `%f / %F` prints decimal with no exponent (full decimal number up to memory capabilities)
    - `%g` prints large exponents (preserves decimal precision and adds exponent, i.e. 2356766363.80001 --> 2.35676636380001e+09) use when we don't want decimal cut off like in `%f`

    strings
    - `%s` prints string value without doubles quotes
    - `%q` prints string value with doubles quotes

    Width (for any data type)
    - when we want to display formatted printed lines a certain way, we can use width
    - rather than typing out "Start:     steph" (5 spaces from Start:, then 5 chars for "steph"), if I want the whole thing to be 10 chars wide, I can do `printF("Start: %10s", "Steph")`
    - %10 will make it 10 chars left-justified
    - %-10 will make it 10 chars right-justified
    - `s` stands for the data type, can be s for string, q for quoted string, d for base10 int, etc

    Precision 
    - (designed for floating point nums, but can be used to chop off characters/bytes from any line)
    - round floating point numbers to a precise number of decimals (it mathematically rounds)
    - `%9f` sets left width, but default precision
    - `%.2f` sets no width, but precision 2 (would chop "Steph" to "St")
    - `%9.2f` sets 9 left width, precision 2
    - `%9.f` sets 9 left width, precision 0 (round to whole num)
      - for strings this chops it to empty string, .1 leaves it at one char

    Padding
    - only works with `0` int
    - `%09d` pads digit to length nine with preceeding 0's (doesn't work appending them after)
    - `%-4d` pads with spaces (width 4, left justified) -- this is just width in action
      - but if you try to do `%-04d`, it doesn't append 0's after

  - `.Sprintf(..., ...)`
    - _formats_ and stores a variable at the same time
    - var x string = `fmt.Sprintf("Hello %v", 24)`
    - the above line doesn't print anything to the console, instead takes the int (24) and formats it into the string and then stores it into var x

    ## escaped characters
    - `\n` newline
    - `\t` tab


### Getting user input and converting it into numeric types

** when doing logic and a function returns more than one value and you only want to use one, you can set the 2nd value as _ and go won't throw an error if it's unused **

if an err is 2nd arg and you don't account for it (just using _), the logic using the 1st arg (the result) will just be empty

```js
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// under the hood, Go always looks for the main function as the entrypoint, that's what runs when the file is executed
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type in the year you were born: ")
	// scanner.Scan() is synchronously waiting for us to type something
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d years old at the end of 2020.", 2020-input)
}

```


# Arithmetic operators and math
- both values happening in an operation MUST be the same data type. Even if you're adding 4 + 4, but one is stored in a float64 variable and the other in an int variable, it won't work

- can do basic conversions using data types as functions

```js
func main() {
	var num1 float64 = 8
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Printf("%f", answer)
}
```

- warning: when doing a math equations in ints, the final value will stay an int and not deduce to a floating point num. So 9 / 2 = 4, not 4.5. If you want precision those original two nums need to be stored as floating point values.
- however, when using the mod `%` operator, two ints can be divided with a remainder, and as long as that remainder is also an int, ints can be used here rather than floating points
  - you can't mod or divide by 0, the code will compile OK but you'll get a runtime error
- `math` package gives you access to many mathematical formulas -- absolute value, exponents, etc.

# Conditions & Boolean expressions
- `<`
- `>`
- `<=`
- `>=`
- `==`
- `!=`
- putting two values between a conditional will return a `true` or `false` boolean
- just like with arithmetic operators, we can't compare two values from two different data types, go will throw an error
  - if you're trying to check
  ```js
  func main() {
    x := 5
    y := 6.5
    val := float64(x) == y
    fmt.Printf("%t", val)
    // false
  }
  ```

  - note: arithmetic operators are all evaluated before conditional operators

  ```go
  // comparing strings
  func main() {
    x := "Hello"
    y := "hello"
    val := x == y
    fmt.Printf("%t", val)
    // false
  }

  func main() {
    x := "hello"
    y := "hello"
    val := x == y
    fmt.Printf("%t", val)
    // true
  }
  // I think this works differently than C, because in C this would return false because it doesn't reference the same location in memory? Not sure though.

  // when comparing strings with < and >, it will compare the numeric location of the ASCII characters from left to right

  // you can compare single quote characters as well
  ```

# Chained Conditionals
- combine multiple boolean expressions to come to one answer
- logical operators
  - && (if both are true, return true, otherwise return false)
  - || (if one or both is true, return true, if both are false, return false)
  - ! (return inverse)

# If, Else If, Else

```go
func main() {
	name := "Sue"
	fmt.Println("before if")
	if name == "Steph" {
		fmt.Println("the name is steph")
	} else if name == "Sue" {
		fmt.Println("the name is sue")
  } else {
		fmt.Println("the name is neither steph nor sue")
  }
	fmt.Println("after if")

}
// only diff between js and golang is the condition isn't housed in ()
```

# For Loops
- in Go, For and While loops are the same, but there are several different implementations of For loops
- the purpose of a For loop is to execute the same block of code a certain # of times

```go
// while-like For Loop
func main() {
	x := 1
	for x <= 5 {
		// while x is less than 5, do this thing repeatedly
		fmt.Printf("x (%d) is still less than or equal to 5\n", x)
		x++
	}
}
```

```go
	// more traditional for loop -- initializing, setting contraints, and incrementing all here
func main() {
	for x := 0; x <= 5; x += 2 {
		fmt.Println(x)
	}
}
```

- break: similar to return, just exits for loop no matter what
- continue: skips everything below, but then goes back to the beginning of the loop


```go
func main() {
	// print all nums that are divisible by 3, 7, and 9
	// if number isn't divisible by those nums, print N
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue
		}
		fmt.Println("N")
	}
}
```

```go
func main() { // could also use else block, but using continue is considered cleaner
	// print all nums that are divisible by 3, 7, and 9
	// if number isn't divisible by those nums, print N
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
		} else {
			fmt.Println("N")
		}
	}
}
```

```go
func main() {
	// print all nums that are divisible by 3, 7, and 9
	// don't print any non-divisible nums
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
		}
	}
}
```

```go
func main() {
	// print only the first number that's divisible by 3, 7, and 9 between 0 - 1000
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			break
		}
	}
}
```

# Switch statement

- cases need to match same data type as passed in variable

```go
func main() {
	// similar to if statement, but cleaner if you're looking for a specific case
	// swtich only runs once, so if you change the var at play in any of the cases, it won't run thru the switch again w/ updated value
	ans := 1

	switch ans {
	case 1:
		fmt.Println("ans is 1")
		ans++
		fmt.Println(ans)
	case 2:
		fmt.Println("ans is 2")
	case 3:
		fmt.Println("ans is 3")
	default:
		fmt.Println("ans is not 1, 2, or 3")
	}
}
```

```go
  func main() {
    // similar to if statement, but cleaner if you're looking for a specific case
    // swtich only runs once, so if you change the var at play in any of the cases, it won't run thru the switch again w/ updated value
    ans := 3

    switch ans {
    // each individual case can have multiple "catches"
    case 1, 2, 3, 4:
      fmt.Println("ans is 1, 2, 3, 4")
    case 5, 6, 7, 8:
      fmt.Println("ans is 5, 6, 7, 8")
    case 9:
      fmt.Println("ans is 9")
    default:
      fmt.Println("ans is none of these")
    }
  }
```

```go

func main() {
	// you can run a switch statement with no initial variable, instead reference the var in each case. But this isn't much more beneficial than an if/else at this point
	ans := 0

	switch {
	case ans > 0:
		fmt.Println("ans is greater than 0")
	case ans < 0:
		fmt.Println("ans is less than 0")
	default:
		fmt.Println("ans is 0")
	}
}
```

# Arrays
- Slices are the main thing used in Golang, but Arrays can complement them

- Collection of values/elements that are ordered
- How arrays work in Golang
  - Arrays are represented with 3 things:
  - Pointer: the location of the first element in the array
  - Length: how many elements can be in the array
  - Capacity: always the same as the length, the max amount of elems that can be in the array
- This works because as long as we know the starting point of the array and how many subsequent elements come after it, we know all the contents of our array
- To access arr[2], all we need to do is add 2 to the original location in memory of the start of our array to find the value that's 2 down from it

```go
// to initialize and populate an array separately

func main() {
	// all elements in the array need to be the same type
	// need to say exactly how many elements we want in the array in order to specify max size
	var arr [5]int
	// how to access elements in an array
	arr[0] = 100
	arr[4] = 500
	fmt.Println(arr)
	fmt.Println(arr[4])
}
```

```go

func main() {
	// all elements in the array need to be the same type
	// need to say exactly how many elements we want in the array in order to specify max size
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// to get length of array
	fmt.Println(len(arr))
}
```
```go
// using arrays with for loops
func main() {
	arr := [3]int{4, 5, 6}
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
```
```go
func main() {
	// two dimensional array
	// first bracket signifies how many internal arrays will be inside our main array
	// second bracket signifies how many elements each internal array will hold (must be same?)
	arr2D := [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	for i := 0; i < len(arr2D); i++ {
    // just print the 2nd element of each array
		fmt.Println(arr2D[i][1])
	}
}
```

# Slices
- slices add on functionality and fix some of the shortcomings we may have with Arrays
- with standard arrays, we need to pick the size off the bat
- Slices let us take portions of an array, and we can change the size of it as we add/remove elements

- we represent slices the same way we represent arrays, except:
  - we have a pointer, which tells us the memory location of the first element in our slice. This will be the same memory location as whichever element # it is in the original array
  - length: num of elements in the slice itself
  - capacity: how many elements we _could_ have. How many elements are left in the original array we're building on top of that we can continue to expand toward

  ** length and capacity are different in slices
  - if orig array is [1, 2, 3, 4, 5, 6, 7, 8] and our slice grabs [3, 4, 5], the pointer is the same location as origArr[2], the length is 3, and the capacity is 6 (because we could extend our slice all the way to 8, like the orig array)

  ```go
  func main() {
    var x [5]int = [5]int{1, 2, 3, 4, 5}
    // when I don't put the # of elements inside the square brackets when declaring a var type, Go knows its a slice
    // on right assignment side, colon inside [] denotes what indexes to grab from underlying array and put in slice
    // just [:] means copy entire thing, go all the way from beginning to end
    // no value on left? start at beginning
    // no value on right? go to very end
    // if you specify a # on the right, it means _go up to_ this var, but don't include it
    var s []int = x[:4]
    fmt.Printf("slice: %v\n", s)
    fmt.Printf("length of slice: %d\n", len(s))
    // capacity, remember this extends to end of orig array
    fmt.Printf("capacity of slice: %d\n", cap(s))

    // slice our orig slice, still references topmost array for capacity
    secondSlice := s[2:3]
    fmt.Printf("second minislice: %v\n", secondSlice)
    fmt.Printf("second minislice length: %v\n", len(secondSlice))
    fmt.Printf("second minislice capacity: %v\n", cap(secondSlice))
  }

  // create slice w/o array
  func main() {
    // if you do this and create a slice without referencing an original array,
    // golang will automatically create a base array for you and then copy/slice the entire thing
    var a []string = []string{"a", "b", "c", "d"}
    // slice of our slice still retains capacity of orig created array
    fmt.Println(cap(a[:3]))

    // adding elems to a slice
    // doesn't extend orig array, instead creates a new slice with more space
    // but you can overwrite orig variable for consistency
    // extension := append(a, "e")
    a = append(a, "e")

    fmt.Println(a)

    // make fn, take in a slice declaration and how long you want that slice capacity to be
    b := make([]int, 5)
    fmt.Printf("make slice: %v\n", b)
    // comes back as type slice of ints!
    fmt.Printf("make slice type: %T\n", b)
  }
```

# Range

```go
func main() {
	// slice of ints
	var a []int = []int{1, 3, 5, 67, 90, 332, 3, 55}

	// traditional for loop to iterate thru slice
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// using range
	// i represents the index of the element in the slice we're looking at
	// element represents the value at the index we're at

	for i, element := range a {
		// if you don't use either i or element, replace with _
		fmt.Printf("%d: %d\n", i, element)

		// 0: 1
		// 1: 3
		// 2: 5
		// 3: 67
		// 4: 90
		// 5: 332
		// 6: 3
		// 7: 55
	}
}
```
```go
func main() {
	a := []int{6, 1, 3, 4, 56, 7, 12, 4, 6}

	// iterate thru and print out each num in slice, but remove duplicates

	// novice way
	for i, element := range a {
		for j, element2 := range a {
			// print out duplicates, but make sure not to double print them by checking if second loop index is ahead of first loop index
			if element == element2 && j > i {
				fmt.Println(element)
			}
		}
	}

	// advanced way
	for i, element := range a {
		// house more intricate j for loop inside range for i
		for j := i + 1; j < len(a); j++ {
			//in the case where j is greater than the length of "a" (8 elems)
			// because j is always i + 1, j will be 9 on final loop and it wont run bc 9 !< 8
			element2 := a[j]
			if element == element2 {
				fmt.Printf("element: %d\n index: %d\n element2: %d\n index2: %d\n\n", element, i, element2, j)
			}
		}
	}
}
```

# Maps
- maps are Golang's version of JS objects
- Maps are collections of key:value pairs, positions don't matter
- lower level func: access happens almost instantly to grab a value by its key in a map...access happens faster in maps than in slices/arrays

```go
func main() {
	// make a map with string keys that map to int values
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	// this makes an empty map
	mp2 := make(map[string]int)

	fmt.Println(mp)
	fmt.Println(mp2)

	// to change values:
	fmt.Println(mp["apple"])

	mp["apple"] = 900

	// add a value:
	mp["steph"] = 1001

	fmt.Println(mp)

	// to delete a value (if the specified key arg doesn't exist, nothing happens)
	delete(mp, "steph")
	fmt.Println(mp)
}
```
```go
func main() {
	// make a map with string keys that map to int values
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	// check if a key exists
	// if it does, val will be value and ok will be true
	// if it doesn't, ok will be false and val will be the default value for that type
	// NOTE: checking doesn't add not-there keys and default values to the map
	val, ok := mp["no"]

	fmt.Printf("value: %d\nexists? %t\n", val, ok)
	fmt.Println(mp)

	// value: 0
	// exists? false
	// map[apple:5 orange:9 pear:6]

	// to get length (# of key:value pairs in a map)
	fmt.Printf("length of map: %d", len(mp))

}
```

# Functions
- a function is a block of reusable code
- you can call a function, pass in a value, and get a value out

```go
  func add(x, y int) { // can also write x, y int (if they're both int types)
    fmt.Println(x + y)
  }

  func main() {
    add(16, 5)
    add(6, 7)
  }
```
```go
func add(x, y int) (int, int) { // can also write x, y int (if they're both int types)
	// go can return multiple values
  return x + y, x - y
}

func main() {
	ans1, ans2 := add(16, 5)
	fmt.Println(ans1, ans2)
}
```
- named return values
```go
func add(x, y int) (z1 int, z2 int) {
	// if you name return variables like above, you just need to declare them somewhere in the fn body and then just write return
	z1 = x + y // notice I'm just setting values, not re-declaring
	z2 = x - y
	return
}

func main() {
	ans1, ans2 := add(16, 5)
	fmt.Println(ans1, ans2)
}
```
- defer
```go
func add(x, y int) (z1 int, z2 int) {
	// defer runs once the function hits its return, "cleanup" call of sorts
	defer fmt.Println("Hello! I'm done")
	z1 = x + y // notice I'm just setting values, not re-declaring
	fmt.Println("doing stuff before return...")
	z2 = x - y
	return
}

func main() {
	ans1, ans2 := add(16, 5)
	fmt.Println(ans1, ans2)
}
```

- functions inside functions
```go
func main() {

	// test is a var that is storing an anonymous function that takes no args, doesn't return anything
	test := func(x int) int {
		return x * -1
	}

	fmt.Println(test(5))

	// functions can be invoked inline
	immediate := func(s string) string {
		return fmt.Sprintf("Hey %s!\n", s)
	}("Steph")

	fmt.Printf("immediate value: %s", immediate)

}
```
```go
func main() {
	test := func(x int) int {
		return x * -1
	}
	// outer expects another function to execute inside, but we create and pass it outside of that function's context
	outer(test)

	test2 := func(x int) int {
		return x + 5
	}
	// calling outer with another inner function
	outer(test2)
}
```
- closures
```go
func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	// closures are when we use a value outside of the function, inside of it

	// first invocation returns the inner function
	// second invocation actually invokes that inner fn to println
	returnFunc("hey")()

	// closure concept: even though param "yes" is only passed to first arg of returnFunc, second func has access to it inside
	container := returnFunc("yes")
	container()
}
```

# Mutable & Immutable data types
- ints, floats, strings are immutable data types
```go
// if z = 6, z + 1 does not change z to 7. You'd have to reassign the new val back to z in order to overwrite it
z = 6
z + 1
// 6
z = z + 1 // 7
```
- slices, maps, and arrays are mutable data types
```go
func main() {
	var x int = 5
	// y is equal to the value stored in x
	// this line of code does NOT bind y and x together
	// changes in x or y will not affect the other
	// it essentially just copies whatever the value was at the moment
	y := x
	// change the value of y
	y = 7

	fmt.Println(x, y)
}
```
- slices and maps are mutable data types
```go
func main() {
	var x []int = []int{3, 4, 5}
	// slices are mutable
	// y := x says y is equal to the slice x is pointing to
	// so they end up pointing to the same underlying slice, so changes to either will be changes to both
	y := x
	// slices are mutable, so this change below will also affect x
	// maps are also mutable
	// new assignments are references, not copies
	y[0] = 100
	fmt.Println(x, y)
}
```
```go
func main() {
	// maps are mutable
	mp := map[string]string{
		"apples": "green",
		"lemons": "yellow",
		"grapes": "purple",
	}
	// assigning another pointer to the same underlying map, these now reference eachother
	mp2 := mp

	mp2["lemons"] = "neon"

	delete(mp, "grapes")

	fmt.Println(mp, mp2)

	//map[apples:green lemons:neon] map[apples:green lemons:neon]
}
```
- arrays are mutable (they can change), however it doesn't have the same reference behavior as maps and slices
```go
func main() {
	var arr [5]int = [5]int{2, 4, 6, 8, 10}
	// not the same behavior as maps and slices, these two vars do not point to the same place in memory
	// instead it just makes a copy, so they're detached
	arr2 := arr
	arr2[0] = 20
	fmt.Println(arr, arr2)
	// [2 4 6 8 10] [20 4 6 8 10]
}
```

### why some variables are mutable and others aren't
- with maps and slices, the variables assigned to them hold a **value that points** to a place in memory where the data structure is stored.
  - you can create functions that modify arrays (changeFirstVal, deleteLastVal) and it will work, since the param passed in will point to the same place in memory as the orig value
- with single value values, the value itself is stored in directly in the variable's location in memory, so when you do x := y, it just copies that value stored in the y and puts it in x. 
- With arrays, values are also copied, not referenced

# Pointers & Dereference operator (& and *)
- although examples here show pointers and dereferencers being used in functions and not, the main use case for them would be while passing in params into fns and wanting to manipulate the original argument value if it's a single-value data type
```go
func main() {
	// pointers allow us to do more complex things in go
	// &: get the pointer
	// *: dereference
	x := 7
	// value is 7, reference (x) is the location of where this piece of data is stored
	// to look at the reference (x) and not the value, use &
	// tells computer, tell me the reference of x (where is the value 7 actually stored?)
	fmt.Printf("pointer x: %d\n", &x) // 0x14000130008, hexadecimal code of where in our computer's memory x is stored
	// a reference == a pointer
	fmt.Printf("value x: %d\n", x)
	// by accessing the pointer (location of the reference x), this allows us to:

	// y is equal to the pointer of x
	// doesn't mean y = 7, instead means y = location in memory of where the value 7 is stored
	// you could change the value 7, but it's location in memory wouldn't change
	y := &x

	x = 9
	fmt.Println("changed value, x = 9")
	// even tho x value is different, address is same
	fmt.Printf("value of y, which = &x: %d\n", y)

	fmt.Printf("value x: %d\nvalue y: %d\n", x, y)

	// *y tells y: hey, I know that your value is the address of memory of variable x,
	// but rather than changing the memory address I prefixed you with * to have you change
	// the value at that address directly, not the address itself

	*y = 8 // this goes to mem address (y MUST be pointer) and opens value and changes it

	fmt.Printf("value x: %d\nvalue y dereferenced: %d\n", x, y)

	// pointer x: 1374389649520
	// value x: 7
	// changed value, x = 9
	// value of y, which = &x: 1374389649520
	// value x: 9
	// value y: 1374389649520
	// value x: 8
	// value y dereferenced: 1374389649520
}
```
```go
// when you pass in a asterisk in front of a var data type,
// you're asking for the pointer to the value, not the value itself
// *string = param must be a pointer to a type: string
// the reason we do this is because we can change the value using dereferencing
// if we just passed in the value itself, it doesn't change
// this is because single-value data types are immutable.
func changeValue(str *string) {
	// since str here is a pointer, I have to dereference it (*) in order to manipulate the actual value, rather than the address itself
	*str = "changed!"

}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	toChange := "hello"
	// because toChange requires a pointer param, we pass in &toChange, the pointer to the address in memory where "hello" is stored
	changeValue(&toChange)
	fmt.Println(toChange)

	toChange = "hello"
	// when passing an arg into a fn,
	// I'm essentially doing str := toChange --> str = "changed", but because str and toChange aren't references to eachother (just copies),
	// a change to str doesn't effect toChange, the original value
	// (this is why pointer are necessary for immutable data types)
	changeValue2(toChange)
	fmt.Println(toChange)

}
```

```go
func main() {
	toChange := "hello"
	// how to declare a pointer
	var pointer *string = &toChange
	// before dereferencing pointer to change value at address
	fmt.Printf("pointer: %d\noriginal val: %q\n", pointer, toChange)
	fmt.Printf("pointer value (*pointer): %q\n", *pointer)

	*pointer = "hello Steph!"
	// after dereferencing pointer to change value at address
	fmt.Printf("new toChange after dereferencing pointer: %q\n", toChange)

	// pointer: 1374389600832
	// original val: "hello"
	// pointer value (*pointer): "hello"
	// new toChange after dereferencing pointer: "hello Steph!"
}
```

# Structs and Custom Types
- every variable type has a specific set of features and behavior
- struct defines a custom type we want to use, essentially is an object with specific types within that we define
- while a `map` is defined like this: 	
```go
var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
  ```
  A struct is defined as this and refers to whatever the struct definition contains defined above.'
  ```go
type Point struct {
	x int32
	y int32
}

func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}
	fmt.Println(p1, p2)
	// {1 2} {-5 7}
	// can also print specific values within the struct, since you know they exist in the type
	fmt.Println(p1.y)
	// 2

  // can also directly change the fields
  p1.y = 100
  fmt.Println(p1)
}
```
```go
func main() {
	var p1 Point = Point{1, 2, true}
	// simpler way to write:
	p2 := Point{-5, 7, false}
	// if you don't have all the necessary data to fill in all fields, can do the following:
	// unnamed fields will be assigned default type vals
	p3 := Point{isOnGrid: false}

	fmt.Println(p1, p2, p3)
	// {1 2 true} {-5 7 false} {0 0 false}
}
```
- structs and pointers
```go
type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {
	// why are we not doing (*p1).x here?
	// the reason here we don't directly dereference pt:
	// structs behave differently, we dont need to dereference pt and get the obj value and then access the x value,
	// instead we can access x value directly from the pointer with the dot (.) operator
	pt.x = 100
}

func main() {
	// declare as pointer so you can pass it into functions and still manipulate the actual value
	// otherwise values inside a struct would be immutable (just like single-value Go types), it would just make a copy inside the function and not change the original custom value
	// so by pointing/dereferencing we get around that when passing structs into fns
	p1 := &Point{y: 3}
	fmt.Println(p1)
	// &{0 3}
	changeX(p1)
	fmt.Println(p1)
	// &{100 3}

	// NOTE: both instances of using pointers and just passing structs into functions are valid
	// in some cases, you might just want the function to make a copy
	// in other cases, you'll need to mutate the original struct
}
```
- embedded structs
```go
type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center *Point
	// you can also define embedded pointers by themselves without a key
	// *Point
	// their subfields (x and y), as long as they don't conflict with other base fields on Circle
	// will be available on Circle now (similar to obj spreading in JS)
}

func main() {
	// accessing embedded structs

	// when you pass a struct into another struct, make sure you pass a pointer to the embedded struct, rather than the struct itself
	c1 := Circle{4.56, &Point{4, 5}}
	// if you log c1 out as a whole, it'll print the pointer address for the center key
	fmt.Println(c1)
	// in order to access the value (struct) itself, use dot notation
	fmt.Println(c1.center)
	// can also access sub values x and y
	fmt.Println(c1.center.x)

	// {4.56 0x140000ae000}
	// &{4 5}
	// 4
}
```

# Struct Methods
- methods are functions we perform on objects
- methods will be specific to the fields of an object they're supposed to read/modify
- for almost all methods, you're going to want to specify the pointer when passing in the connected object parameter

```go
type Student struct {
	name   string
	grades []int
	age    int
}

// a method acts on a specific object, in this case, Student
// arg after func says: this is the param that represents the object we want to act on, as well as it's struct
func (s *Student) getAge() int {
	// getAge is a method that acts on a Student object and returns an int
	return s.age
}

func (s *Student) setAge(age int) {
	fmt.Printf("s1 inside setAge %v\n", s)
	s.age = age
}

// don't need a pointer here because we're just reading and returning data, not trying to manipulate the data
func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	// notice that Student doesn't have to be &Student for setAge (method that accepts s1 pointer) to work
	s1 := &Student{"Steph", []int{90, 95, 98}, 25}

	// getAge works on s1 because s1 is a Student obj
	stephAge := s1.getAge()
	fmt.Println(stephAge)
	// 25

	// because setAge specifically grabs the s1 pointer, s1.age (inside setAge) is able to manipulate actual value in memory (without dereferencing because of dot notation)
	// rather than just changing a copy of it, which would happen if we just grabbed the value itself
	s1.setAge(22)

	fmt.Println(s1)

	average := s1.getAverageGrade()
	fmt.Println(average)
}
```
- using methods w/o pointers (just reading data from obj, not changing it)
```go
type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getMaxGrade() int {
	currentMax := 0
	for _, grade := range s.grades {
		if currentMax < grade {
			currentMax = grade
		}
	}
	return currentMax
}
func main() {
	// notice that Student doesn't have to be &Student for setAge (method that accepts s1 pointer) to work
	s1 := &Student{"Steph", []int{90, 95, 98}, 25}

	maxGrade := s1.getMaxGrade()
	fmt.Println(maxGrade)
}
```

# Interfaces
- an interface is a way to look at a set of related objects or types
- if a group of objects all have the same named method (with the same return), even if the implementations of the method are different, they are said to to type _____ (whatever the interface name is) and that they implement that interface
	- the objects that implement the interface MUST have methods for all defined method fields on the interface, not just one of two of them

- you can use in interface as a type when defining a variable, but all the other struct properties on the passed in objs to that variable aren't accessible from within the interface

- we achieve polymorphism (OOP concept where different classes can be used with the same interface) in Go by implementing an interface in more than one struct

- using interfaces as a slice:
```go
import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	width float64
}

type RightTriangle struct {
	aLength float64
	bLength float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (s Square) area() float64 {
	return math.Pow(2, s.width)
}

func (t RightTriangle) area() float64 {
	// obviously this isn't how to get triangle area, fudged for example purposes
	return t.aLength * t.bLength
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (s Square) perimeter() float64 {
	return s.width * 4
}

func (t RightTriangle) perimeter() float64 {
	cLength := t.getCLength()
	return t.aLength + t.bLength + cLength
}

func (t RightTriangle) getCLength() float64 {
	cSqrd := math.Pow(t.aLength, 2) + math.Pow(t.bLength, 2)

	return math.Sqrt(cSqrd)
}

func main() {
	sqr := Square{4}
	rect := Rectangle{2, 5}
	tri := RightTriangle{3, 3}

	fmt.Println(sqr, rect, tri)

	shapes := []Shape{sqr, rect, tri}
	fmt.Println(shapes)

	for i, shape := range shapes {
		fmt.Printf("printing area for shape %d: %f\n", i, shape.area())
		// can't call shape.width or shape.bLength, those properties are struct-specific and aren't implemented in the interface
		fmt.Printf("printing perimeter for shape %d: %f\n", i, shape.perimeter())
	}
}
```
```go
// can pass interfaces in as function parameters too! and access their methods there.
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	sqr := Square{4}
	rect := Rectangle{2, 5}
	tri := RightTriangle{3, 3}

	fmt.Println(sqr, rect, tri)

	shapes := []Shape{sqr, rect, tri}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
```
- with pointers:
```go
import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	width float64
}

type RightTriangle struct {
	aLength float64
	bLength float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (s *Square) area() float64 {
	return math.Pow(2, s.width)
}

func (t *RightTriangle) area() float64 {
	// obviously this isn't how to get triangle area, fudged for example purposes
	return t.aLength * t.bLength
}

func (r *Rectangle) perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func (s *Square) perimeter() float64 {
	return s.width * 4
}

func (t *RightTriangle) perimeter() float64 {
	cLength := t.getCLength()
	return t.aLength + t.bLength + cLength
}

func (t *RightTriangle) getCLength() float64 {
	cSqrd := math.Pow(t.aLength, 2) + math.Pow(t.bLength, 2)

	return math.Sqrt(cSqrd)
}

// can pass interfaces in as function parameters too! and access their methods there.
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	sqr := Square{4}
	rect := Rectangle{2, 5}
	tri := RightTriangle{3, 3}

	fmt.Println(sqr, rect, tri)

	// because interface methods now accept pointers, must change these args to pointers, rather than values themselves
	// this would be necessary if we were modifying data on the objs, rather than just reading it
	// that said, it's always best practice to pass the pointer so we have access to the original value if we need it.
	// never hurts to pass the pointer
	shapes := []Shape{&sqr, &rect, &tri}
	fmt.Println(shapes)

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
```