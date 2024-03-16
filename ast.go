package main

type ExprType int
type OpKind int

const (
	Constant ExprType = iota
	Binary
)

const (
	Add OpKind = iota
	Sub
	Mul
	Div
)

type Exprs struct {
	BinaryOp    *BinaryOp
	ConstantVal *ConstantVal
}

type Expr struct {
	Type     ExprType
	Operator *Exprs
}

func newExpr(t ExprType, e *Exprs) *Expr {
	ex := new(Expr)
	ex.Type = t
	ex.Operator = e
	return ex

}

func (e *Expr) eval() int {
	switch e.Type {
	case Constant:
		return e.Operator.ConstantVal.eval()
	case Binary:
		return e.Operator.BinaryOp.eval()
	default:
		panic("Invalid operator")
	}
}

type ConstantVal int

func newConstantVal(val int) *ConstantVal {
	cv := ConstantVal(val)
	return &cv
}

func (cv *ConstantVal) eval() int {
	return int(*cv)
}

type BinaryOp struct {
	Kind      OpKind
	LeftExpr  *Expr
	RightExpr *Expr
}

func newBinaryOp(kind OpKind, leftExpr *Expr, rightExpr *Expr) *BinaryOp {
	po := new(BinaryOp)
	po.Kind = kind
	po.LeftExpr = leftExpr
	po.RightExpr = rightExpr
	return po
}

func (po *BinaryOp) eval() int {
	right := po.RightExpr.eval()
	left := po.LeftExpr.eval()
	switch po.Kind {
	case Add:
		return left + right
	case Sub:
		return left - right
	case Mul:
		return left * right
	case Div:
		return left / right
	default:
		panic("Invalid Operator kind")
	}
}
