package rpcdemo

import "errors"

// 以Service.Method方式来调用

type DemoService struct {
}

// 除法所需的2个运算数
type Args struct {
	A, B int
}

// go中对rpc的要求：2个参数args和result，并返回error
// 因为结果要写入result中，所以result必须是指针类型
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}