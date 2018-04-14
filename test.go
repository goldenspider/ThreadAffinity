package main

import (
	"fmt"
	"runtime"

	"time"

	"golang.org/x/sys/windows"
)

func main() {

	go func() {
		runtime.LockOSThread()
		fmt.Println("2", windows.GetCurrentThreadId())
		handle, err := windows.GetCurrentThread()
		if err != nil {
			fmt.Printf("err=%v", err)
		}
		//bind to core 2
		err = windows.SetThreadAffinityMask(handle, uint64(2)) 
		if err != nil {
			fmt.Printf("err=%v", err)
		}
		////cpu busy test start
		var b float64 = 1
		var c float64 = 1
		for i := 0; i < 500000000; i++ {
			b = b/c + b/10.0 + c*12.2
			c = c/b + c*10.0 + b*12.2
			b = b/c + b*10.0 + c/12.2
			c = c/b + c*10.0 + b*12.2
			b = b/c + b/10.0 + c*12.2
			c = c/b + c*10.0 + b*12.2
			b = b/c + b*10.0 + c/12.2
			c = c/b + c*10.0 + b*12.2
		}
		////cpubusy test end
		fmt.Printf("b=%f c=%f\n", b, c)
		time.Sleep(10 * time.Second)
	}()

	fmt.Println("1", windows.GetCurrentThreadId())

	time.Sleep(30 * time.Second)
}
