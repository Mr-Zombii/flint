package main

// #include <string.h>
import "C"
import "unsafe"

//export printString
func printString(str *C.char) {
	ptr := unsafe.Pointer(str)
	length := int(C.strlen(str))

	for i := range length {
		unsafeCharacterPtr := unsafe.Add(ptr, i)
		characterPtr := (*C.char)(unsafeCharacterPtr)
		character := *characterPtr

		print(string(character))
	}
}

func main() {}
