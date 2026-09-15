package test

import (
	"fmt"
	"library/util"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := util.Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	} else {
		fmt.Println("test pass")
	}
}
