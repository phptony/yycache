package yycache

import (
	"fmt"
	"github.com/phptony/yycache/util"
)
import "testing"

func TestT(t *testing.T)  {
	q := util.NewQueue()
	q.Push("1")
	q.Push("2")
	q.Push("3")
	fmt.Println(q.Poll())
	fmt.Println(q.Size())
	fmt.Println(q.Poll())

}