package main

type ExprType int

const (
	Constant ExprType = iota
	Plus     ExprType = iota
)

type Exprs struct {
	PlusOp      *PlusOp
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
	case Plus:
		return e.Operator.PlusOp.eval()
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

type PlusOp struct {
	LeftExpr  *Expr
	RightExpr *Expr
}

func newPlusOp(leftExpr *Expr, rightExpr *Expr) *PlusOp {
	po := new(PlusOp)
	po.LeftExpr = leftExpr
	po.RightExpr = rightExpr
	return po
}

func (po *PlusOp) eval() int {
	return po.LeftExpr.eval() + po.RightExpr.eval()
}
