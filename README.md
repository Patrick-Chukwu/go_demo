# Summary

## Package
a package is a way to group functions, and it's made up of all the files in the same directory

### fmt package
The `fmt` package contains functions for formatting text, including printing to the console. It is one of the standard library packages you get upon installing go
## Binary
range of bit: (256)
Max:  2^8 - 1
Min: 0


### signed binary numbers (For negative numbers)
The first digit is signed bit and it is here below:
1 - negative
0 - positive


## Signed Integers: int
```
Max : 2^(8-1)-1
Min: - 2^(8-1)
```
### Know the difference between signed integers and unsigned integer

uint is always positive
Int can either be positive or negative

### byte
A byte is 8 bits.
- it is uint8 datatype
It can be used to store a number or character
type = int8
use single quotation

### bits
The number of values that can be represented. e.g int8

## Rune
type= int32

## bool
true or false


## string
use double quotation

## Implicit Vs Explicit

### 
```
Println - standard output
Printf()- Formatted
```

### %T
```
%T - Data type
%v - value
%b - binary
%d - decimal
%e - scientific representation
%f - float representation (e.g fmt.Printf(".2f"), 10.2f)
%s- string
%% - To add %
\" - for quotation 
```


## Sprintf
creates a string but not print it out


## Arithmetic operator

## The math package
- Min(a,b) (function)
- Max(a,b)
- Sqrt(a)
- Pow(a,b)
- Round(a)
- Ceil(a)
- Floor(a) : round down

### Convert from string to integer
- Use `strconv` package
    - .Atoi(): string to integer
    - .ParseInt(value, base, bit_size)
 
## Conditions
- A condition is any expression that evaluates to `true` or `false`
### Comparison Operators
- <
- >
- <=
- >=
- ==
- !=

### Logical Operators
- || - Or
- && - And
- ! - Not


### If 
```
x:= 2
if x < 3 {
    fmt.Println("run")
} else if x >5 {
    fmt.Println("jump")
} else {
    fmt.Println("sit")
}
```

### switch
```
a := 1
switch a {
    case 1: 
    fmt.Println("one")
    case 2:
    fmt.Println("two")
    default:a package is a way to group functions, and it's made up of all the files in the same directory
    fmt.Println("default")

}
```

to fall through every case, you should type `fallthrough`
    - You can also check for multiple cases
a package is a way to group functions, and it's made up of all the files in the same directory

### For loop

```
for inx := 0; idx <10; idx++ {
    fmt.Println(idx)

}a package is a way to group functions, and it's made up of all the files in the same directory
```
- In Go, we don't have access to a `while` loop
    - To run a while, use a for loop with a single condition

### Looping through strings

```
str := "hello world"
```

    - We have index acces, however it gives the values of the integer representation of the character
    - strings are slices of characters or bytes
    - you need to convert back to string


## ASCII vs UTF-8
### ASCII 
1 byte and 256

UTF-8 - 4 byte

## Array
Array are a fixed sized data structure that can store values of the same type.
- The type of array indicates its size
- Once we initialize an array the size cannot be changed
- set the value with {} not ()
- `array literal` is writing the values of the array upon declaration
- using [...] allows the compiler automatically count the value in our array literal
```
var arr [2]int
var arr [3]int{1 3 4}
arr1 := [2]int{4 6}
```

### Properties of array
- Index values can be accessed
- It can be mutated
- Length can be calculated
- You can't change the size
- You change the

## Slices