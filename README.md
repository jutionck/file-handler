## File Handler

### Install

```bash
go get github.com/jutionck/file-handler
```

### Use

```go
package main

import (
	"fmt"

	filehandler "github.com/jutionck/file-handler"
)

func main() {

	filePath := "/Users/jutioncandrakirana/Desktop/belajar-file-golang/data.txt"

	file, err := filehandler.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	filehandler.CloseFile(file)
}
```
