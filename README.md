# golang-from-ground-up
Complete golang guide

## Things covered are the following:
* Array, Slices , Maps 
* Concurrency 
* Functions 
* Structs 
* Interfaces 
* Packages 
* Pointers 

## Explanations and Details 

### Modules and Packages
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

### Conventions
* It is better to use `_` instead of `-` when handling `multi-word` file names
* CamelCase naming is used in GO
* GO is a statically typed language so types are important in go 
* You cannot mismatch types in calculations
* Adding Types Explicitly help us get rid off explicit type conversions
* Use `:=` as often as possible
* You can add comma if 2 var have the same type or different type
* 