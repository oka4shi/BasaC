package main

import (
	"reflect"
	"testing"
)

func TestConstantVal(t *testing.T) {
	expect := 55
	constantVal := newConstantVal(expect)

	if !reflect.DeepEqual(expect, constantVal.eval()) {
		t.Fail()
	}
}

func TestPlusOp(t *testing.T) {
	// 13+(5+1)
	thirteen := newExpr(Constant, &Exprs{ConstantVal: newConstantVal(13)})
	five := newExpr(Constant, &Exprs{ConstantVal: newConstantVal(5)})
	one := newExpr(Constant, &Exprs{ConstantVal: newConstantVal(1)})

	fivePlusOne := newExpr(Plus, &Exprs{PlusOp: newPlusOp(five, one)})

	plusOp := newExpr(Plus, &Exprs{PlusOp: newPlusOp(thirteen, fivePlusOne)})

	expect := 13 + (5 + 1)

	if !reflect.DeepEqual(plusOp.eval(), expect) {
		t.Fail()
	}
}
