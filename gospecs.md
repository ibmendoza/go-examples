Basic Go
============

<h3>About</h3>

This book is a verbatim subset version of [Go language specification] (https://golang.org/ref/spec). Installation procedure is specific for Windows users only but the concepts apply for non-Windows operating systems.

Go is brought to you by the Go authors courtesy of Google, Inc.

This PDF is generated from GitHub Flavored Markdown by Isagani Mendoza [**http://itjumpstart.wordpress.com**] (http://itjumpstart.wordpress.com) 

<h3>Share</h3>

This PDF can be shared subject to terms of Creative Commons Attribution 3.0 License and copyright by Google, Inc.


<h2>Part 1 - Install Go</h2>

<h3>Download Go</h3>

- Go to https://golang.org/dl to download the Go installer
- For Windows 32-bit - download the file ending with windows-386.msi
- For Windows 64-bit - download the file ending with windows-amd64.msi


<h3>Go Installation Folder</h3>

- By default, Go for Windows will install it at C:\Go. The Go installer will automatically set the Path system environment variable 
- For other operating systems, please follow instructions at https://golang.org/doc/install

<h3>Setup Go workspace</h3>

From http://golang.org/doc/code.html

> Go code must be kept inside a workspace. A workspace is a directory hierarchy with three directories at its root: 

- src contains Go source files organized into packages (one package per directory),
- pkg contains package objects, and
- bin contains executable commands

Under Windows, it is a convention to create the Go workspace under C:\mygo

Under C:\mygo, create three more sub-folders named bin, pkg and src. 

- bin
- pkg
- src

Your folder structure should be:

```
C:\mygo\bin
C:\mygo\pkg
C:\mygo\src
```

<h3>The GOPATH environment variable</h3>

Once you have setup the Go workspace (C:\mygo), create a new system environment variable named GOPATH and put C:\mygo as its value.

```
GOPATH=C:\mygo
```

Under Windows 7, access Control Panel\System and Security\System\Advanced system settings\Environment variables\System variables, click New.

```
Variable name = GOPATH
Variable value = C:\mygo
```

Reboot your computer to reflect the changes.

<h3>Install GitHub client</h3>

Go was designed to access open source third-party packages from repositories like GitHub so you need a GitHub client depending on your operating system.  

- For Windows, download at https://windows.github.com
- For Mac, download at https://mac.github.com
- For RPM-based Linux, ```sudo yum install git```
- For Debian-based Linux, ```sudo apt-get install git```

**Resource:** https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

For Windows, you need to ensure that Git has been added to your system variable named Path.

For example,

```
Variable name = Path
Variable value = C:\Go\bin;C:\Program Files\Git\bin
```

**Reboot** your computer to reflect the changes.

<h3>Download LiteIDE</h3>

A minimal IDE (integrated development environment) has been built exclusively for Go. Ironically, the latest version is getting bloated so I suggest you download an earlier version (23.2) although your mileage may vary.

```
http://sourceforge.net/projects/liteide/files
```

<h3>Your First Program</h3>

Remember that under Windows, the Go workspace by convention is located at C:\mygo (or at any folder of your choosing). The same goes with non-Windows operating systems.

The src folder under the Go workspace is the location of your Go programs such that there is one program per folder.

For example, this is our Go workspace under Windows. 

```
c:\mygo\src
```

Creating our first program is as easy as creating a text file using your favorite editor (in our case, we will use LiteIDE).

So the customary Hello world program will be like this:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

To compile, click Build under LiteIDE. Under Windows, open Windows explorer and you will the executable named src.exe under ```C:\mygo\src```.

To run, double-click the executable. If you noticed, you have not seen the output because the program exited just before you were able to see it!

Output:

```
Hello World
```

<h3>Organize Your Programs</h3>

Our first program was located at ```C:\mygo\src\main.go``` and the executable was named src.exe.

That's because Go compiles your program and name it based on location folder.


> To organize our programs, create a folder per program. Each program can be nested to a sub-folder.

For example, create main.go under hello folder of C:\mygo\src,

```
c:\mygo\src\hello\main.go
```

When you compile the above main.go program, it will output ```hello.exe```


<h3>Get Third-Party Packages</h3>

For now, we will get our first package named gosl (Go Script Language) from David Deng's repository under GitHub.

At this stage, I assume you have already setup Git on your system. For Windows and Mac users, you need to install the native GitHub client.

At the command line (aka terminal), type the following:

```
go get github.com/daviddengcn/gosl
```

Go has built-in tools to manage your programs. Read https://golang.org/cmd/go.

The commands are:

```
build       compile packages and dependencies
clean       remove object files
env         print Go environment information
fix         run go tool fix on packages
fmt         run gofmt on package sources
generate    generate Go files by processing source
get         download and install packages and dependencies
install     compile and install packages and dependencies
list        list packages
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         run go tool vet on packages
```

Under the hood, **LiteIDE** builds, formats and runs your program so you need not learn the command-line for now. However, once you get past the advanced stage of learning Go, you will find these commands indispensable especially for scripting and other tasks.

Under Windows, you will notice that Go installed gosl at

```
C:\mygo\src\github.com\daviddengcn\gosl
``` 

Storing packages under github.com folder of your Go workspace is Go's way of organizing programs such that you can reference it on your local computer.

Using LiteIDE, build gosl.go so it outputs gosl.exe.


**Why gosl?**

Since Go is a compiled language, gosl serves as a [REPL] (http://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) or language shell so that you can play around Go language constructs without the compile-then-run cycle.

As David Deng has said:

> Gosl is an application that can make you write script with the Go language. It is NOT an interpreter but the pure Go. The preprocessor tranforms the script into a Go program, instantly compiles and runs. So it is almost the same as standard Go with the same efficiency.

In other words, we need not learn another scripting language just to play around with Go.

<h2>Part 2 - The Basics</h2>

<h3>Hello World</h3>

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

There are two kinds of programming languages.

- compiled - translates a program (source code) into an executable (Examples: Go, C, etc)
- interpreted - translates a program (source code) into a format that is executed by a runtime software (Examples: JavaScript, Python, PHP, etc)

When programming, you have to bear in mind the syntax and semantics of each programming language construct.

In our example above, Go has particular syntax when it comes to compiling a program.

```go
package main
```

This indicates to Go compiler that this is an executable program.

```go
import "fmt"
```

import is a reserved Go keyword which means we have to import or use a built-in Go package named "fmt" that prints output to the screen.

```go
func main() {

}
```

This is the body of the program. Like the body of a letter, you write here the instructions that you want to be executed.

```go
	fmt.Println("Hello World")
```

This is the code that prints "Hello World" to the screen.

<h3>Syntax</h3>

From https://golang.org/ref/spec

- The syntax is specified using Extended Backus-Naur Form (EBNF)
- Source code is Unicode text encoded in [UTF-8] (http://en.wikipedia.org/wiki/UTF-8)

<h3>Comments</h3>

A comment is ignored by the compiler (meaning it does not get translated into executable code). As such, comments are useful to document your program either inline or within a block.

```go
//this is a comment
```

```go
/*
	This is a block comment
	Anything you put within this block
	are all comments
*/
```

**Tokens**

Tokens form the vocabulary of the Go language. There are four classes: 

- identifiers
- keywords
- operators and delimiters
- literals. 

White space, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token.

<h3>Identifiers</h3>

Identifiers name program entities such as variables and types.

An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

For example, in physics:

```go
f = ma
```

where

f - force
m - mass
a - acceleration

In physics, scientists use notation to denote symbols that represent a certain variable. In programming, it's the same thing.

However, [programming is not physics nor algebra] (https://andy.wordpress.com/2012/05/30/programming-is-not-algebra/).

To illustrate, I borrowed Andy Skelton's example.

```go
a = 10
b = 20
a = b

// What are the values of a and b?
```

In the first line, a equals 10.
In the second line, b equals 20.
In the third line, how can a be equal to b?

So if you're coming from a math or science class, you have to abandon that line of thinking in order to do programming.

The third line is an illustration of **variable assignment**.

In this case, we assign to variable **a** whatever is the value of variable **b**!

<h3>Keywords</h3>

The following keywords are reserved and may not be used as identifiers. 

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

<h3>Operators and Delimiters</h3>

The following character sequences represent operators, delimiters, and other special tokens:

```go
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

<h3>Integer Literals</h3>

An integer literal is a sequence of digits representing an integer constant. An optional prefix sets a non-decimal base: 0 for octal, 0x or 0X for hexadecimal. In hexadecimal literals, letters a-f and A-F represent values 10 through 15.

```go
42
0600
0xBadFace
170141183460469231731687303715884105727
```

<h3>Floating-point Literals</h3>

A floating-point literal is a decimal representation of a floating-point constant. It has an integer part, a decimal point, a fractional part, and an exponent part. The integer and fractional part comprise decimal digits; the exponent part is an e or E followed by an optionally signed decimal exponent. One of the integer part or the fractional part may be elided; one of the decimal point or the exponent may be elided.

```go
0.
72.40
072.40  // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
```

<h3>Rune Literals</h3>

A rune literal represents a rune constant, an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes. Within the quotes, any character may appear except single quote and newline. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats. 

```go
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'aa'         // illegal: too many characters
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

<h3>String Literals</h3>

A string literal represents a string constant obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

Raw string literals are character sequences between back quotes ``. Within the quotes, any character is legal except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.

```go
`abc`  // same as "abc"
`\n
\n`    // same as "\\n\n\\n"
"\n"
""
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"       // illegal: surrogate half
"\U00110000"   // illegal: invalid Unicode code point
``` 
 
<h3>Constants</h3>

```go
package main

import "fmt"

const (
	pi = 3.14
	isBoolean = true
	thousand = 1000
)

func main() {

	const s string = "This is a string"
	
	fmt.Println(s)
	fmt.Println(pi)
	fmt.Println(isBoolean)
	fmt.Println(thousand)
}
```
<h3>Variables</h3>

A variable is a storage location for holding a value. The set of permissible values is determined by the variable's type.

```go
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```

<h2>Part 3 - Variable Types</h2>

<h3>Boolean Types</h3>

```go
var isAvailable = false
var isOk = true 
```

<h3>Numeric Types</h3>

A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are: 

```go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

<h3>String Types</h3>

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. 

**Strings are immutable**: once created, it is impossible to change the contents of a string. The predeclared string type is string.

The length of a string s (its size in bytes) can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer indices 0 through len(s)-1. It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid.

```go
var str = "The quick brown fox jumps over the lazy dog"
```

<h3>Array Types</h3>

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative. 

```go
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

<h3>Slice Types</h3>

A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is nil.
 Like arrays, slices are indexable and have a length. The length of a slice s can be discovered by the built-in function len; unlike with arrays it may change during execution. The elements can be addressed by integer indices 0 through len(s)-1. The slice index of a given element may be less than the index of the same element in the underlying array.

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function cap(a).

A new, initialized slice value for a given element type T is made using the built-in function make, which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing 

```go
make([]T, length, capacity)
```

produces the same slice as allocating an array and slicing it, so these two expressions are equivalent:

```go
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

<h3>Struct Types</h3>

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField). Within a struct, non-blank field names must be unique. 

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

<h3>Pointer Types</h3>

A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. The value of an uninitialized pointer is nil.

```go
*Point
*[4]int
```

<h3>Function Types</h3>

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil.

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

The final parameter in a function signature may have a type prefixed with .... A function with such a parameter is called **variadic** and may be invoked with zero or more arguments for that parameter. 

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

<h3>Interface Types</h3>

An interface type specifies a method set called its interface. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is nil. 

As with all method sets, in an interface type, each method must have a unique non-blank name.

```go
// A simple File interface
interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}
```

More than one type may implement an interface. For instance, if two types S1 and S2 have the method set.

```go
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
```

(where T stands for either S1 or S2) then the File interface is implemented by both S1 and S2, regardless of what other methods S1 and S2 may have or share.

A type implements any interface comprising any subset of its methods and may therefore implement several distinct interfaces. For instance, all types implement the empty interface: 

```go
interface{}
```

Similarly, consider this interface specification, which appears within a type declaration to define an interface called Locker:

```go
type Locker interface {
	Lock()
	Unlock()
}
```

If S1 and S2 also implement

```go
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

they implement the Locker interface as well as the File interface.

An interface T may use a (possibly qualified) interface type name E in place of a method specification. This is called embedding interface E in T; it adds all (exported and non-exported) methods of E to the interface T.

```go
type ReadWriter interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type File interface {
	ReadWriter  // same as adding the methods of ReadWriter
	Locker      // same as adding the methods of Locker
	Close()
}

type LockedFile interface {
	Locker
	File        // illegal: Lock, Unlock not unique
	Lock()      // illegal: Lock not unique
}
```

An interface type T may not embed itself or any interface type that embeds T, recursively.

```go
// illegal: Bad cannot embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}
```

<h3>Map Types</h3>

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil.

The comparison operators == and != must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a run-time panic. 

```go
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

The number of map elements is called its length. For a map m, it can be discovered using the built-in function len and may change during execution. Elements may be added during execution using assignments and retrieved with index expressions; they may be removed with the delete built-in function.

A new, empty map value is made using the built-in function make, which takes the map type and an optional capacity hint as arguments:

```go
make(map[string]int)
make(map[string]int, 100)
```

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of nil maps. A nil map is equivalent to an empty map except that no elements may be added. 

<h3>Channel Types</h3>

A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil. 

The optional <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bidirectional. A channel may be constrained only to send or only to receive by conversion or assignment.

```go
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

The <- operator associates with the leftmost chan possible:

```go
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments: 

```go
make(chan int, 100)
```

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A nil channel is never ready for communication.

A channel may be closed with the built-in function close. The multi-valued assignment form of the receive operator reports whether a received value was sent before the channel was closed.

A single channel may be used in send statements, receive operations, and calls to the built-in functions cap and len by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent. 

<h2>Part 4 - Declarations and scope</h2>

A declaration binds a non-blank identifier to a constant, type, variable, function, label, or package. Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

The blank identifier may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier init may only be used for init function declarations, and like the blank identifier it does not introduce a new binding. 

 The scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

Go is lexically scoped using blocks:

- The scope of a predeclared identifier is the universe block.
- The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
- The scope of the package name of an imported package is the file block of the file containing the import declaration.
- The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
- The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
- The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

The package clause is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same package and to specify the default package name for import declarations.

<h3>Label scopes</h3>

Labels are declared by labeled statements and are used in the "break", "continue", and "goto" statements. It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.

<h3>Blank identifier</h3>

The blank identifier is represented by the underscore character _. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in declarations, as an operand, and in assignments.

<h3>Exported identifier</h3>

 An identifier may be exported to permit access to it from another package. An identifier is exported if both:

- the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
- the identifier is declared in the package block or it is a field name or method name.

All other identifiers are not exported. 

<h3>Uniqueness of identifier</h3>

Given a set of identifiers, an identifier is called unique if it is different from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different packages and are not exported. Otherwise, they are the same.

<h3>Constant declaration</h3>

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of constant expressions. The number of identifiers must be equal to the number of expressions, and the nth identifier on the left is bound to the value of the nth expression on the right.

If the type is present, all constants take the type specified, and the expressions must be assignable to that type. If the type is omitted, the constants take the individual types of the corresponding expressions. If the expression values are untyped constants, the declared constants remain untyped and the constant identifiers denote the constant values. For instance, if the expression is a floating-point literal, the constant identifier denotes a floating-point constant, even if the literal's fractional part is zero.

```go
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)

// a = 3, b = 4, c = "foo", untyped integer and string constants
const a, b, c = 3, 4, "foo"  
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
```

Within a parenthesized const declaration list the expression list may be omitted from any but the first declaration. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list. The number of identifiers must be equal to the number of expressions in the previous list. Together with the iota constant generator this mechanism permits light-weight declaration of sequential values:

```go
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
```

<h3>Iota</h3>

Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. It is reset to 0 whenever the reserved word const appears in the source and increments after each ConstSpec. It can be used to construct a set of related constants: 

```go
const (  // iota is reset to 0
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 << iota  // a == 1 (iota has been reset)
	b = 1 << iota  // b == 2
	c = 1 << iota  // c == 4
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0 (iota has been reset)
const y = iota  // y == 0 (iota has been reset)
```

Within an ExpressionList, the value of each iota is the same because it is only incremented after each ConstSpec: 

```go
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	bit1, mask1                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3, mask3                           // bit3 == 8, mask3 == 7
)
```

This last example exploits the implicit repetition of the last non-empty expression list. 



<h3>Type Declarations</h3>

A type declaration binds an identifier, the type name, to a new type that has the same underlying type as an existing type, and operations defined for the existing type are also defined for the new type. The new type is different from the existing type.

```go
type IntArray [16]int

type (
	Point struct{ x, y float64 }
	Polar Point
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

The declared type does not inherit any methods bound to the existing type, but the method set of an interface type or of elements of a composite type remains unchanged: 

```go
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of the base type of PtrMutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its anonymous field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
```

A type declaration may be used to define a different boolean, numeric, or string type and attach methods to it: 

```go
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT+%dh", tz)
}
```

<h3>Variable Declarations</h3>

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value. 

```go
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
```

 If a list of expressions is given, the variables are initialized with the expressions following the rules for assignments. Otherwise, each variable is initialized to its zero value.

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first converted to its default type; if it is an untyped boolean value, it is first converted to type bool. The predeclared value nil cannot be used to initialize a variable with no explicit type. 

```go
var d = math.Sin(0.5)  // d is int64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
```

Implementation restriction: A compiler may make it illegal to declare a variable inside a function body if the variable is never used. 

<h3>Short Variable Declarations</h3>

```go
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w := os.Pipe(fd)  // os.Pipe() returns two values

// coord() returns three values; only interested in y coordinate
_, y, _ := coord(p)  
```

Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block with the same type, and at least one of the non-blank variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original. 

```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2  
```

Short variable declarations may appear only inside functions. In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables. 

<h3>Function Declarations</h3>

If the function's signature declares result parameters, the function body's statement list must end in a terminating statement. 

```go
func findMarker(c <-chan int) int {
	for i := range c {
		if x := <-c; isMarker(x) {
			return x
		}
	}
	// invalid: missing return statement.
}
```

A function declaration may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine. 

```go
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func flushICache(begin, end uintptr)  // implemented externally
```

<h3>Method Declarations</h3>

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type. 

 The receiver is specified via an extra parameter section preceeding the method name. That parameter section must declare a single parameter, the receiver. Its type must be of the form T or *T (possibly using parentheses) where T is a type name. The type denoted by T is called the receiver base type; it must not be a pointer or interface type and it must be declared in the same package as the method. The method is said to be bound to the base type and the method name is visible only within selectors for that type.

A non-blank receiver identifier must be unique in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.

For a base type, the non-blank names of methods bound to it must be unique. If the base type is a struct type, the non-blank method and field names must be distinct.

Given type Point, the declarations 

```go
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

bind the methods Length and Scale, with receiver type *Point, to the base type Point.

The type of a method is the type of a function with the receiver as first argument. For instance, the method Scale has type 

```go
func(p *Point, factor float64)
```

However, a function declared this way is not a method.

<h2>Part 5 - Expressions</h2>

An expression specifies the computation of a value by applying operators and functions to operands. 

<h3>Operands</h3>

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly qualified) non-blank identifier denoting a constant, variable, or function, a method expression yielding a function, or a parenthesized expression.

The blank identifier may appear as an operand only on the left-hand side of an assignment.

<h3>Qualified Identifier</h3>

A qualified identifier is an identifier qualified with a package name prefix. Both the package name and the identifier must not be blank. 

A qualified identifier accesses an identifier in a different package, which must be imported. The identifier must be exported and declared in the package block of that package. 

```go
math.Sin	// denotes the Sin function in package math
```

<h3>Composite Literals</h3>

Composite literals construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated. They consist of the type of the value followed by a brace-bound list of composite elements. An element may be a single expression or a key-value pair. 

 The LiteralType must be a struct, array, slice, or map type (the grammar enforces this constraint except when the type is given as a TypeName). The types of the expressions must be assignable to the respective field, element, and key types of the LiteralType; there is no additional conversion. The key is interpreted as a field name for struct literals, an index for array and slice literals, and a key for map literals. For map literals, all elements must have a key. It is an error to specify multiple elements with the same field name or constant key value.

For struct literals the following rules apply:

- A key must be a field name declared in the LiteralType.
- An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
- If any element has a key, every element must have a key.
- An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
- A literal may omit the element list; such a literal evaluates to the zero value for its type.
- It is an error to specify an element for a non-exported field of a struct belonging to a different package.

Given the declarations 

```go
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
```

one may write 

```go
origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
```

 For array and slice literals the following rules apply:

- Each element has an associated integer index marking its position in the array.
- An element with a key uses the key as its index; the key must be a constant integer expression.
- An element without a key uses the previous element's index plus one. If the first element has no key, its index is zero.

Taking the address of a composite literal generates a pointer to a unique variable initialized with the literal's value.

```go
var pointer *Point3D = &Point3D{y: 1000}
```

The length of an array literal is the length specified in the LiteralType. If fewer elements than the length are provided in the literal, the missing elements are set to the zero value for the array element type. It is an error to provide elements with index values outside the index range of the array. The notation ... specifies an array length equal to the maximum element index plus one. 

```go
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
```

A slice literal describes the entire underlying array literal. Thus, the length and capacity of a slice literal are the maximum element index plus one. A slice literal has the form 

[]T{x1, x2, … xn}

and is shorthand for a slice operation applied to an array: 

```go
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
```

Within a composite literal of array, slice, or map type T, elements that are themselves composite literals may elide the respective literal type if it is identical to the element type of T. Similarly, elements that are addresses of composite literals may elide the &T when the element type is *T. 

```go
// same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[...]Point{{1.5, -3.5}, {0, 0}}   

// same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]int{{1, 2, 3}, {4, 5}}        

// same as [...]*Point{&Point{1.5, -3.5}, &Point{0, 0}}
[...]*Point{{1.5, -3.5}, {0, 0}}  
```

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the keyword and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses. 

```go
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
```

Examples of valid array, slice, and map literals: 

```go
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```

<h3>Function Literal</h3>

A function literal represents an anonymous function. 

```go
func(a, b int, z float64) bool { return a*b < int(z) }
```

A function literal can be assigned to a variable or invoked directly. 

```go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals are closures: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible. 

<h3>Primary Expression</h3>

Primary expressions are the operands for unary and binary expressions. 

```go
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
```

<h3>Method expression</h3>

If M is in the method set of type T, T.M is a function that is callable as a regular function with the same arguments as M prefixed by an additional argument that is the receiver of the method. 

Consider a struct type T with two methods, Mv, whose receiver is of type T, and Mp, whose receiver is of type *T. 

```go
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
```

The expression 

```go
T.Mv
```

yields a function equivalent to Mv but with an explicit receiver as its first argument; it has signature 

```go
func(tv T, a int) int
```
That function may be called normally with an explicit receiver, so these five invocations are equivalent: 

```go
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
```

Similarly, the expression 

```go
(*T).Mp
```

yields a function value representing Mp with signature 

```go
func(tp *T, f float32) float32
```

For a method with a value receiver, one can derive a function with an explicit pointer receiver, so 

```go
(*T).Mv
```

yields a function value representing Mv with signature 

```go
func(tv *T, a int) int
```

 Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.

Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given f := T.Mv, f is invoked as f(t, 7) not t.f(7). To construct a function that binds the receiver, use a function literal or method value.

It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type. 

<h3>Slice expression</h3>

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity. 

**Simple slice expressions**

For a string, array, pointer to array, or slice a, the primary expression 

```go
a[low : high]
```

constructs a substring or slice. The indices low and high select which elements of operand a appear in the result. The result has indices starting at 0 and length equal to high - low. After slicing the array a

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

the slice s has type []int, length 3, capacity 4, and elements 

```go
s[0] == 2
s[1] == 3
s[2] == 4
```

For convenience, any of the indices may be omitted. A missing low index defaults to zero; a missing high index defaults to the length of the sliced operand: 

```go
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

 If a is a pointer to an array, a[low : high] is shorthand for (*a)[low : high].

For arrays or strings, the indices are in range if 0 <= low <= high <= len(a), otherwise they are out of range. For slices, the upper index bound is the slice capacity cap(a) rather than the length. A constant index must be non-negative and representable by a value of type int; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy low <= high. If the indices are out of range at run time, a run-time panic occurs.

Except for untyped strings, if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type string. If the sliced operand is an array, it must be addressable and the result of the slice operation is a slice with the same element type as the array.

If the sliced operand of a valid slice expression is a nil slice, the result is a nil slice. Otherwise, the result shares its underlying array with the operand

**Full slice expressions**

For an array, pointer to array, or slice a (but not a string), the primary expression 

```go
a[low : high : max]
```

constructs a slice of the same type, and with the same length and elements as the simple slice expression a[low : high]. Additionally, it controls the resulting slice's capacity by setting it to max - low. Only the first index may be omitted; it defaults to 0. After slicing the array a

```go
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

the slice t has type []int, length 2, capacity 4, and elements 

```go
t[0] == 2
t[1] == 3
```

 As for simple slice expressions, if a is a pointer to an array, a[low : high : max] is shorthand for (*a)[low : high : max]. If the sliced operand is an array, it must be addressable.

The indices are in range if 0 <= low <= high <= max <= cap(a), otherwise they are out of range. A constant index must be non-negative and representable by a value of type int; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a run-time panic occurs. 


<h3>Type Assertions</h3>

For an expression x of interface type and a type T, the primary expression 

```go
x.(T)
```

 asserts that x is not nil and that the value stored in x is of type T. The notation x.(T) is called a type assertion.

More precisely, if T is not an interface type, x.(T) asserts that the dynamic type of x is identical to the type T. In this case, T must implement the (interface) type of x; otherwise the type assertion is invalid since it is not possible for x to store a value of type T. If T is an interface type, x.(T) asserts that the dynamic type of x implements the interface T.

If the type assertion holds, the value of the expression is the value stored in x and its type is T. If the type assertion is false, a run-time panic occurs. In other words, even though the dynamic type of x is known only at run time, the type of x.(T) is known to be T in a correct program. 

```go
var x interface{} = 7  // x has dynamic type int and value 7
i := x.(int)           // i has type int and value 7

type I interface { m() }
var y I

// illegal: string does not implement I (missing method m)
s := y.(string)        

// r has type io.Reader and y must implement both I and io.Reader
r := y.(io.Reader)     
```

A type assertion used in an assignment or initialization of the special form 

```go
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
```

yields an additional untyped boolean value. The value of ok is true if the assertion holds. Otherwise it is false and the value of v is the zero value for type T. No run-time panic occurs in this case. 

<h3>Calls</h3>

Given an expression f of function type F, 

```go
f(a1, a2, … an)
```

calls f with arguments a1, a2, … an. Except for one special case, arguments must be single-valued expressions assignable to the parameter types of F and are evaluated before the function is called. The type of the expression is the result type of F. A method invocation is similar but the method itself is specified as a selector upon a value of the receiver type for the method. 

```go
math.Atan2(x, y)  // function call
var pt *Point
pt.Scale(3.5)     // method call with receiver pt
```

 In a function call, the function value and arguments are evaluated in the usual order. After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the calling function when the function returns.

Calling a nil function value causes a run-time panic.

As a special case, if the return values of a function or method g are equal in number and individually assignable to the parameters of another function or method f, then the call f(g(parameters_of_g)) will invoke f after binding the return values of g to the parameters of f in order. The call of f must contain no parameters other than the call of g, and g must have at least one return value. If f has a final ... parameter, it is assigned the return values of g that remain after assignment of regular parameters. 

```go
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
```

A method call x.m() is valid if the method set of (the type of) x contains m and the argument list can be assigned to the parameter list of m. If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m(): 

```go
var p Point
p.Scale(3.5)
```

There is no distinct method type and there are no method literals. 

<h3>Passing arguments to ... parameters</h3>

If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent to type []T. If f is invoked with no actual arguments for p, the value passed to p is nil. Otherwise, the value passed is a new slice of type []T with a new underlying array whose successive elements are the actual arguments, which all must be assignable to T. The length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site.

Given the function and calls 

```go
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
```

within Greeting, who will have the value nil in the first call, and []string{"Joe", "Anna", "Eileen"} in the second.

If the final argument is assignable to a slice type []T, it may be passed unchanged as the value for a ...T parameter if the argument is followed by .... In this case no new slice is created.

Given the slice s and call 

```go
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
```

within Greeting, who will have the same value as s with the same underlying array.

<h2>Part 6 - Statements</h2>

<h3>Statements</h3>

A labeled statement may be the target of a goto, break or continue statement. 

```go
Error: log.Panic("error encountered")
```

<h3>Expression statements</h3>

With the exception of specific built-in functions, function and method calls and receive operations can appear in statement context. Such statements may be parenthesized. 

The following built-in functions are not permitted in statement context: 

```go
append cap complex imag len make new real
unsafe.Alignof unsafe.Offsetof unsafe.Sizeof
```

```go
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

<h3>Send statement</h3>

A send statement sends a value on a channel. The channel expression must be of channel type, the channel direction must permit send operations, and the type of the value to be sent must be assignable to the channel's element type. 

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a run-time panic. A send on a nil channel blocks forever. 

```go
ch <- 3  // send value 3 to channel ch
```

<h3>IncDec statement</h3>

The "++" and "--" statements increment or decrement their operands by the untyped constant 1. As with an assignment, the operand must be addressable or a map index expression. 

The following assignment statements are semantically equivalent: 

```go
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

<h3>Assignment</h3>

Each left-hand side operand must be addressable, a map index expression, or (for = assignments only) the blank identifier. Operands may be parenthesized. 

```go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

An assignment operation x op= y where op is a binary arithmetic operation is equivalent to x = x op y but evaluates x only once. The op= construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier. 

```go
a[i] <<= 2
i &^= 1<<n
```

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a channel or map operation, or a type assertion. The number of operands on the left hand side must match the number of values. For instance, if f is a function returning two values, 

```go
x, y = f()
```

assigns the first value to x and the second to y. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the nth expression on the right is assigned to the nth operand on the left: 

```go
one, two, three = '一', '二', '三'
```

The blank identifier provides a way to ignore right-hand side values in an assignment: 

```go
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

The assignment proceeds in two phases. First, the operands of index expressions and pointer indirections (including implicit pointer indirections in selectors) on the left and the expressions on the right are all evaluated in the usual order. Second, the assignments are carried out in left-to-right order. 

```go
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
```

In assignments, each value must be assignable to the type of the operand to which it is assigned, with the following special cases:

1. Any typed value may be assigned to the blank identifier.
2. If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first converted to its default type.
3. If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first converted to type bool.

<h3>If statement</h3>

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed. 

```go
if x > max {
	x = max
}
```

The expression may be preceded by a simple statement, which executes before the expression is evaluated. 

```go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

<h3>Switch statement</h3>

"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute. 

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. 


<h3>Expression switches</h3>

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value true.

In a case or default clause, the last non-empty statement may be a (possibly labeled) "fallthrough" statement to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

The expression may be preceded by a simple statement, which executes before the expression is evaluated. 

```go
	switch tag {
	default:
		s3()
	case 0, 1, 2, 3:
		s1()
	case 4, 5, 6, 7:
		s2()
	}

	switch x := f(); { // missing switch expression means "true"
	case x < 0:
		return -x
	default:
		return x
	}

	switch {
	case x < y:
		f1()
	case x < z:
		f2()
	case x == 4:
		f3()
	}
```

<h3>Type switches</h3>

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a type assertion using the reserved word type rather than an actual type: 

```go
switch x.(type) {
// cases
}
```

Cases then match actual types T against the dynamic type of the expression x. As with type assertions, x must be of interface type, and each non-interface type T listed in a case must implement the type of x. 

```go
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
TypeList        = Type { "," Type } .
```

The TypeSwitchGuard may include a short variable declaration. When that form is used, the variable is declared at the beginning of the implicit block in each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

The type in a case may be nil; that case is used when the expression in the TypeSwitchGuard is a nil interface value.

Given an expression x of type interface{}, the following type switch: 

```go
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

could be rewritten: 

```go
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
```

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

The "fallthrough" statement is not permitted in a type switch. 

<h3>For statement</h3>

A "for" statement specifies repeated execution of a block. The iteration is controlled by a condition, a "for" clause, or a "range" clause. 

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value true. 

```go
for a < b {
	a *= 2
}
```

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement. The init statement may be a short variable declaration, but the post statement must not. Variables declared by the init statement are re-used in each iteration. 

```go
for i := 0; i < 10; i++ {
	f(i)
}
```

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the semicolons are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value true. 

```go
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns iteration values to corresponding iteration variables if present and then executes the block. 

he expression on the right in the "range" clause is called the range expression, which may be an array, pointer to an array, slice, string, map, or channel permitting receive operations. As with an assignment, if present the operands on the left must be addressable or map index expressions; they denote the iteration variables. If the range expression is a channel, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the blank identifier, the range clause is equivalent to the same clause without that identifier.

The range expression is evaluated once before beginning the loop, with one exception: if the range expression is an array or a pointer to an array and at most one iteration variable is present, only the range expression's length is evaluated; if that length is constant, by definition the range expression itself will not be evaluated.

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present: 

```go
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```



1. For an array, pointer to array, or slice value a, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to len(a)-1 and does not index into the array or slice itself. For a nil slice, the number of iterations is 0.
2. For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type rune, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be 0xFFFD, the Unicode replacement character, and the next iteration will advance a single byte in the string.
3. The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If map entries that have not yet been reached are removed during iteration, the corresponding iteration values will not be produced. If map entries are created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is nil, the number of iterations is 0.
4. For channels, the iteration values produced are the successive values sent on the channel until the channel is closed. If the channel is nil, the range expression blocks forever.

The iteration values are assigned to the respective iteration variables as in an assignment statement.

The iteration variables may be declared by the "range" clause using a form of short variable declaration (:=). In this case their types are set to the types of the respective iteration values and their scope is the block of the "for" statement; they are re-used in each iteration. If the iteration variables are declared outside the "for" statement, after execution their values will be those of the last iteration. 

```go
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface {}  // value type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}
```

<h3>Go statement</h3>

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space. 

 The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for expression statements.

The function value and parameters are evaluated as usual in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes. 

```go
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true; }} (c)
```

<h3>Select statement</h3>

A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations. 

A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a short variable declaration. The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

Execution of a "select" statement proceeds in several steps: 

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
3. Unless the selected case is the default case, the respective communication operation is executed.
4. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
5. The statement list of the selected case is executed.

Since communication on nil channels can never proceed, a select with only nil channels and no default case blocks forever. 

```go
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
	}
}

select {}  // block forever
```

<h3>Return statement</h3>

A "return" statement in a function F terminates the execution of F, and optionally provides one or more result values. Any functions deferred by F are executed before F returns to its caller. 

In a function without a result type, a "return" statement must not specify any result values. 

```go
func noResult() {
	return
}
```

There are three ways to return values from a function with a result type: 

1. The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and assignable to the corresponding element of the function's result type. 

```go
func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}
```

2. The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply. 

```go
func complexF2() (re float64, im float64) {
	return complexF1()
}
```

3. The expression list may be empty if the function's result type specifies names for its result parameters. The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables. 

```go
func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0
	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)
	return
}
```

Regardless of how they are declared, all the result values are initialized to the zero values for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in scope at the place of the return. 

```go
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
```

<h3>Break statement</h3>

A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function. 

If there is a label, it must be that of an enclosing "for", "switch", or "select" statement, and that is the one whose execution terminates. 

```go
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

<h3>Continue statement</h3>

A "continue" statement begins the next iteration of the innermost "for" loop at its post statement. The "for" loop must be within the same function. 

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances. 

```go
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

<h3>Defer statement</h3>

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking. 

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for expression statements.

Each time a "defer" statement executes, the function value and parameters to the call are evaluated as usual and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. If a deferred function value evaluates to nil, execution panics when the function is invoked, not when the "defer" statement is executed.

For instance, if the deferred function is a function literal and the surrounding function has named result parameters that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on handling panics.) 

```go
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 1
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
```

<h2>Part 7 - Built-in functions</h2>

Built-in functions are predeclared. They are called like any other function but some of them accept a type instead of an expression as the first argument.

The built-in functions do not have standard Go types, so they can only appear in call expressions; they cannot be used as function values. 

<h3>Close</h3>

For a channel c, the built-in function close(c) records that no more values will be sent on the channel. It is an error if c is a receive-only channel. Sending to or closing a closed channel causes a run-time panic. Closing the nil channel also causes a run-time panic. After calling close, and after any previously sent values have been received, receive operations will return the zero value for the channel's type without blocking. The multi-valued receive operation returns a received value along with an indication of whether the channel is closed. 

<h3>Length and capacity</h3>

The built-in functions len and cap take arguments of various types and return a result of type int. The implementation guarantees that the result always fits into an int. 

```go
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
```

The capacity of a slice is the number of elements for which there is space allocated in the underlying array. At any time the following relationship holds: 

```go
0 <= len(s) <= cap(s)
```

 The length of a nil slice, map or channel is 0. The capacity of a nil slice or channel is 0.

The expression len(s) is constant if s is a string constant. The expressions len(s) and cap(s) are constants if the type of s is an array or pointer to an array and the expression s does not contain channel receives or (non-constant) function calls; in this case s is not evaluated. Otherwise, invocations of len and cap are not constant and s is evaluated. 

```go
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(x) is a (non-constant) function call
)
var z complex128
```

<h3>Allocation</h3>

The built-in function new takes a type T, allocates storage for a variable of that type at run time, and returns a value of type *T pointing to it. The variable is initialized as described in the section on initial values. 

```go
new(T)
```

For instance 

```go
type S struct { a int; b float64 }
new(S)
```

allocates storage for a variable of type S, initializes it (a=0, b=0.0), and returns a value of type *S containing the address of the location. 


<h3>Making slices, maps and channels</h3>

The built-in function make takes a type T, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions. It returns a value of type T (not *T). The memory is initialized as described in the section on initial values. 

```go
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```

The size arguments n and m must be of integer type or untyped. A constant size argument must be non-negative and representable by a value of type int. If both n and m are provided and are constant, then n must be no larger than m. If n is negative or larger than m at run time, a run-time panic occurs. 

```go
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for 100 elements
```

<h3>Appending to and copying slices</h3>

 The built-in functions append and copy assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps.

The variadic function append appends zero or more values x to s of type S, which must be a slice type, and returns the resulting slice, also of type S. The values x are passed to a parameter of type ...T where T is the element type of S and the respective parameter passing rules apply. As a special case, append also accepts a first argument assignable to type []byte with a second argument of string type followed by .... This form appends the bytes of the string. 

```go
append(s S, x ...T) S  // T is the element type of S
```

If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, append re-uses the underlying array. 

```go
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")                                  t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }
```

The function copy copies slice elements from a source src to a destination dst and returns the number of elements copied. Both arguments must have identical element type T and must be assignable to a slice of type []T. The number of elements copied is the minimum of len(src) and len(dst). As a special case, copy also accepts a destination argument assignable to type []byte with a source argument of a string type. This form copies the bytes from the string into the byte slice. 

```go
copy(dst, src []T) int
copy(dst []byte, src string) int
```

Examples:

```go
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```

<h3>Deletion of map elements</h3>

The built-in function delete removes the element with key k from a map m. The type of k must be assignable to the key type of m. 

```go
delete(m, k)  // remove element m[k] from map m
```

If the map m is nil or the element m[k] does not exist, delete is a no-op. 

<h3>Handling panics</h3>

Two built-in functions, panic and recover, assist in reporting and handling run-time panics and program-defined error conditions. 

```go
func panic(interface{})
func recover() interface{}
```

While executing a function F, an explicit call to panic or a run-time panic terminates the execution of F. Any functions deferred by F are then executed as usual. Next, any deferred functions run by F's caller are run, and so on up to any deferred by the top-level function in the executing goroutine. At that point, the program is terminated and the error condition is reported, including the value of the argument to panic. This termination sequence is called panicking. 

```go
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
```

The recover function allows a program to manage behavior of a panicking goroutine. Suppose a function G defers a function D that calls recover and a panic occurs in a function on the same goroutine in which G is executing. When the running of deferred functions reaches D, the return value of D's call to recover will be the value passed to the call of panic. If D returns normally, without starting a new panic, the panicking sequence stops. In that case, the state of functions called between G and the call to panic is discarded, and normal execution resumes. Any functions deferred by G before D are then run and G's execution terminates by returning to its caller.

The return value of recover is nil if any of the following conditions holds: 

- panic's argument was nil;
- the goroutine is not panicking;
- recover was not called directly by a deferred function.

The protect function in the example below invokes the function argument g and protects callers from run-time panics raised by g. 

```go
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
```
