package server

import "errors"

type Args struct {
	A, B int
}

type Quotinet struct {
	Quo, Rem int
}
type Arith int

func (t *Arith) Multiply(args *Args, replay *int) error {
	*replay = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotinet) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
