package main

// A short form for import multiple packages.
import (
    t "time"  // import package with an alias name
    "os"
)
// It is the same as
//   import t "time"
//   import "os"

// For direct use without using fmt, not recommended!
import . "fmt"

func main() {
     // We can use Println from "fmt" without fmt.
     Println("Hello World from process", os.Getpid(), "at", t.Now());
}
