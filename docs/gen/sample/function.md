# Function

## Imports

```go
import (
	"fmt"
	"time"
	"github.com/thuongtruong1009/gouse/function"
)
```
## Functions


### funcDelay

```go
func funcDelay() {
	println("Delay start:")

	result := function.DelayF(func() string {
		return "Delayed return after 2s"
	}, 2)

	if result.HasReturn {
		println(result.Value)
	} else {
		println("No result")
	}

	function.Delay(func() {
		println("Delayed not return")
	}, 3)
}
```

### funcRetry

```go
func funcRetry() {
	function.Retry(func() error {
		println("Retry")
		return nil
	}, 3, 1)
}
```

### funcTimes

```go
func funcTimes() {
	function.Times(func() {
		println("Times")
	}, 3)
}
```

### funcInterval

```go
func funcInterval() {
	function.Interval(func() {
		println("Interval")
	}, 1)
}
```

### funcLock

```go
func funcLock() {
	oneInOneOutFuc := function.LockFunc(func(i interface{}) interface{} {
		return i
	}).(func(interface{}) interface{})("one input - one output")
	fmt.Println(oneInOneOutFuc)

	twoInTwoOutFunc1, twoInTwoOutFunc2 := function.LockFunc(func(i1, i2 interface{}) (interface{}, interface{}) {
		return i1, i2
	}).(func(interface{}, interface{}) (interface{}, interface{}))("two input - two output (a)", "two input - two output (b)")
	fmt.Println(twoInTwoOutFunc1, twoInTwoOutFunc2)

	function.LockFunc(func() {
		println("no input - no output")
	}).(func())()
}
```

### funcRunTime

```go
func funcRunTime() {
	exampleFunc := func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task completed.")
	}

	duration := function.RunTime(time.Now(), exampleFunc)
	fmt.Printf("Function run in: %v\n", duration)
}
```
