package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(150 * time.Millisecond)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on other task")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task is done")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func asyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}
func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():

		fmt.Println(ret)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	}
}
