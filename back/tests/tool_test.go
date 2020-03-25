package test

import (
	"fmt"
	"sharemusic/models/tool"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(tool.GetTime(true))
}

func TestHashCode(t *testing.T) {
	fmt.Println(tool.HashCode("123"))
}

func TestConvert2(t *testing.T) {
	a := map[string]string{}
	a["id"] = "123"
	fmt.Println(a)
	fmt.Println(tool.Convert2(a))
}
