# Unique ID generator
This is a very simple package that generates URL safe unique IDs.

> Note: This isn't safe for production use yet

## Example Usage 

### Install 

```shell
$ go get github.com/codehakase/uid
```

### Use in any part of existing code

```go
package main

import (
  "fmt"
  "github.com/codehakase/uid"
)

func main() {
  // ...

  requestID := uid.New(32) // Generates a 32 bit long random string

  fmt.Println(requestID) // you'd get something like:  niwVnfcOcmtWksDCNavCXeTdjWwBWZsC
```

}
