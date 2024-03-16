package main

import (
	"fmt"
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

func TestBinaryOp(t *testing.T) {
	cases := map[string]struct {
		in   *Expr
		want int
	}{
		"13+(5+1)": {
			in:   tAdd(t, tInt(t, 13), tAdd(t, tInt(t, 5), tInt(t, 1))),
			want: 13 + (5 + 1),
		},
		"(2*3)-(4/2)": {
			in:   tSub(t, tMul(t, tInt(t, 2), tInt(t, 3)), tDiv(t, tInt(t, 4), tInt(t, 2))),
			want: (2 * 3) - (4 / 2),
		},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.eval()

			if !reflect.DeepEqual(tt.want, got) {
				fmt.Println("want: ", tt.want)
				fmt.Println("got: ", got)
				t.Fatal("Unexpected data")
			}
		})
	}
}

func tAdd(t *testing.T, a *Expr, b *Expr) *Expr {
	t.Helper()
	return newExpr(Binary, &Exprs{BinaryOp: newBinaryOp(Add, a, b)})
}

func tSub(t *testing.T, a *Expr, b *Expr) *Expr {
	t.Helper()
	return newExpr(Binary, &Exprs{BinaryOp: newBinaryOp(Sub, a, b)})
}
func tMul(t *testing.T, a *Expr, b *Expr) *Expr {
	t.Helper()
	return newExpr(Binary, &Exprs{BinaryOp: newBinaryOp(Mul, a, b)})
}
func tDiv(t *testing.T, a *Expr, b *Expr) *Expr {
	t.Helper()
	return newExpr(Binary, &Exprs{BinaryOp: newBinaryOp(Div, a, b)})
}
func tInt(t *testing.T, a int) *Expr {
	t.Helper()
	return newExpr(Constant, &Exprs{ConstantVal: newConstantVal(a)})
}
