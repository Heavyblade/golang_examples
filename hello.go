package main

import (
    "C"
    "fmt"
)

func main() {
    fmt.Println("from main")
}

//export Test
func Test() int {
    return 30
}

//export TestString
func TestString() string {
    return "s"
}

//export TestBool
func TestBool() bool {
    return true
}

//export StrFxn
func StrFxn() *C.char {
    // C data needs to be manually managed in memory.
    // But we will do it from Haskell.
    return C.CString("Hola mundo")
}

//export StrFxnEcho
func StrFxnEcho(x *C.char) *C.char {
    input := C.GoString(x)
    return C.CString(input)
}

// https://sakshamsharma.com/2018/02/haskell-golang-ffi/
// env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o main.dll hello.go
//https://medium.com/@master.rta/golang-create-a-web-view-app-for-any-platform-54917dea397