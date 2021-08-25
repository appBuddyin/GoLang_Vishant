package pkg

type Operation interface {
	op() float32
}
type Variables struct {
	var1     float32
	var2     float32
	operator string
}

func (v Variables) op() float32 {
	var result float32
	switch v.operator {
	case "mul":
		result = v.var1 * v.var2
	case "add":
		result = v.var1 + v.var2
	case "sub":
		result = v.var1 - v.var2
	case "dev":
		result = v.var1 / v.var2
	}
	return result
}
