# golang-from-ground-up
Complete golang guide

### Things covered are the following:
* Array, Slices , Maps 
* Concurrency 
* Functions 
* Structs 
* Interfaces 
* Packages 
* Pointers 

# Explanations and Details 

## Modules and Packages
* Run using `go run .` or `go run <filename>.go`
* Built-in Commands are Provided by Packages
    * For Example : `fmt` is a `built-in package` which provides I/O functions
* In go you cannot use `single quotes('')` , only the `double quotes ("")` and `back-ticks (``)` are allowed
* Function is a command executed by ourselves but the main function is executed by go itself
* The function must be named main in the main file of the application
* Each go file must have a package , a single package can be split across multiple files 
* We can use a package from one file into another file 
* Some Packages are included in go and some need to be installed 
* `main` is a special Package name which tells go that this package will be the `main entry point` of the application 
* `go build` makes an executible file that can also be executed on systems on which go is not installed but u cannot run it unless u have a go.mod file in the parent folder or the current folder.
* On windows run the executible file by simply running it but on MacOS u have to run it by `./<filename>`
* One module consist of multiple Packages
* run the `go mod init <location>` to make module. The location should be the path where u wish to host your module
* The `go.mod` file tells that the folder in which the file is stored and the sub folder belong to the module of this name 
* GO does not run chronologically from top to bottom like JS
* You also should have a single main function
* If you are building a third party utility you don't need a main function

## Conventions
* It is better to use `_` instead of `-` when handling `multi-word` file names
* CamelCase naming is used in GO

## Variables and Types
* Variables can be re-assigned
* You can add comma if 2 var have the same type or different type
* Use `:=` as often as possible
* GO is a statically typed language so types are important in go 
* You cannot mismatch types in calculations
* Adding Types Explicitly help us get rid off explicit type conversions
* We need to `pass pointer` to a variable to scan so scan populates it with user input 

## FMT Package
* `fmt.Scan()`You can't (easily) fetch multi-word input values. Fetching text that consists of more than a single word is tricky with this function.
* If u dont want to print output on screen but format string and use it internally use `fmt.Sprintf`

## Function 
* It is a code on demand block
* Must add return type for the function 
* If u declare return variable name in the function , u dont need to redefine it inside the function body
* If u are returning 2 values by defining the names in the function you can simply use `return` instead of `return a,b`

## Loops and Switch 
* If u want to get out of the loop and do further processing dont use switch as u cannot break in default.

## File Permissions 
* 755 means you can do anything with the file or directory, and other users can read and execute it but not alter it.Suitable for programs and directories you want to make publicly available.
* 644 means you can read and write the file or directory and other users can only read it. Suitable for public text files.
*711 means you can do anything with the file or directory and other users can only execute it. Suitable for directories where you don't want other people browsing through the contents but do want to give them access to selected files. This is the default for your home directory and the minimum access required for your public_html directory if you have a personal website.
* 700 means you can do anything with the file or directory and other users have no access to it at all. Suitable for private directories and programs.
* 600 means you can read and write the file or directory and other users have no access to it. Suitable for private text files.

## Pacakges 
* Two packages declaration is now allowed in the same folder
* Function name which needs to be exported should always be capital in a package 
* You can search for third-party packages here [go-search](https://pkg.go.dev/)
* All third party dependecies are listed in the `go.mod` file
* When `go get` is run without any path , it downloads all the dependencies in the `go.mod` file

## Pointers 
* Pointers are variables which stores variables addresses instead of values 
* `&` retrieves address of stored value
* Pointers enable direct mutation of values 
* Pointers prevent unnecessary value copies (when passing pointers as argument to function , copy is not created. Function uses the address to look up the value)
* With pointers only one value is stored in memory and the address is passed around
* All values in Go have a so-called `Null Value` - i.e., the value that's set as a default if no value is assigned to a variable.
For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".
For a `pointer`, it's `nil` - a special value built-into Go.
nil represents the `absence` of an `address value `- i.e., a pointer pointing at no address / no value in memory.
* Scan function takes pointer , internally dereference the pointer and overwrites the value stored in the address with value entered by the user 

## Structs
* Related data needs to be passed individually but struct solve this problem by grouping the related data 
* Always accept the struct as pointer and never as a value when making a method which edits the struct ( otherwise it makes a copy and only edits a copy)
* Struct Embedding is building a new struct that builds up on an existing struct
* Always `Prefer anonymous` embedding over `Explicit Embedding`(where u add a fieldname for a struct)
* Struct Tag Assign MetaDeta

## Types 
* type keyword can also be used to assign alias to built-in types 
* To read a single character use `reader.ReadString('\n')` use single quotes instead of double

## Interfaces 
* An interface is a contract that guarantees that a certain value (struct) has a certain method 
* If an interface defines only one method than the interface name is the name of the methid with er at the end
* Use embedding if you want to create composable, reusable interfaces
* Use `interface{}` for any kind of value 

## Generics 
* Generics allow u to write more resuable and more generic functions where go can infer the type of values u are working with 