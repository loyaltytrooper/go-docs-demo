// This is a demonstration of godocs
//
// only the part of overview section
package src

import "fmt"

// bytes per kb represents the number of bytes in a kilobyte
const BytesPerKB = 1024

// PrintKB prints the size of constant
func PrintKb(size int) {
	fmt.Printf("The Bytes Per KB is %d", BytesPerKB)
}
