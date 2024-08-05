# Yog SDK

```ASCII
  ,--.   ,--..-'),-----.   ,----.     
   \  `.'  /( OO'  .-.  ' '  .-./-')  
 .-')     / /   |  | |  | |  |_( O- ) 
(OO  \   /  \_) |  |\|  | |  | .--, \ 
 |   /  /\_   \ |  | |  |(|  | '. (_/ 
 `-./  /.__)   `'  '-'  ' |  '--'  |  
   `--'          `-----'   `------'
```

## Examples 
### customer use
```go
package main

import (
	"fmt"
	"src.nextbillion.io/shared/yog"
)

func main() {
	y := yog.New("xurui-test-task-id", "test", "./tmp")
	err := y.Load()
	if err != nil {
		return
	}

	for true {
		durations, distances, err := y.ReadChunk(30)
		if err != nil {
			break
		}
		fmt.Println(durations)
		fmt.Println(len(durations))
		fmt.Println(distances)
		fmt.Println(len(distances))
	}
}
```

### internal use
```go
package main

import (
	"fmt"
	"src.nextbillion.io/shared/yog"
)

func main() {
	y := yog.NewInternal("xurui-test-task-id", "test", "./tmp", "/Users/xurui/Workspace/env/dcsa.json", "nb-data")
	err := y.Load()
	if err != nil {
		return
	}

	for true {
		durations, distances, err := y.ReadChunk(30)
		if err != nil {
			break
		}
		fmt.Println(durations)
		fmt.Println(len(durations))
		fmt.Println(distances)
		fmt.Println(len(distances))
	}
}

```

## lint

```bash
  # install lint tool
  sh lint.sh install
  # run lint tool use config in .golangci.yaml 
  sh lint.sh run
```
