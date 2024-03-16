package main

import (
	"testing"
	"reflect"
)

func TestConstantVal(t *testing.T){
	expect := 55
	constantVal := newConstantVal(expect)
	
	if !reflect.DeepEqual(expect, constantVal.get()){
		t.Fatal("Test Failed!")
	}
}
