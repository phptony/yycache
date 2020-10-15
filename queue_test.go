package yycache

import (
	"fmt"
	"github.com/phptony/yycache/util"
	"testing"
)
​
func TestQueue(t *testing.T) {
	queue := util.NewQueue()
	queue.Push("a")
	queue.Push("b")
	fmt.Println(queue.Poll())
}