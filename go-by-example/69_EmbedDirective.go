// go:embed is a compiler directive that allows programs to include arbitrary
// files and folders in the Go binary at build time.
// Read more about the embed directive here.

package main

// Import the embed package; if you don’t use any exported identifiers from
// this package, you can do a blank import with _ "embed".

import (
	"embed"
)

// embed directives accept paths relative to the directory containing the Go
// source file. This directive embeds the contents of the file into
// the string variable immediately following it.

//go:embed dir/single_file.txt
var filestring string

// Or embed the contents of the file into a []byte.
//go:embed dir/single_file.txt
var fileByte []byte

// We can also embed multiple files or even folders with wildcards.
// This uses a variable of the embed.FS type,
// which implements a simple virtual file system.
//go:embed dir/single_file.txt
//go:embed dir/*.hash
var folder embed.FS

func main() {
	// Print out the contents of single_file.txt.
	print(filestring)
	print(string(fileByte))

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("tmp/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("tmp/file2.hash")
	print(string(content2))
}
