package rpc

import (
    "calculator-app/pkg/calculator"
    "net"
    "net/rpc"
    "log"
)

type CalculatorService struct {
    Calc calculator.Calculator
}

type Args struct {
    A, B float64
}

type Result struct {
    Value float64
    Err   string
}

func (cs *CalculatorService) Add(args *Args, result *Result) error {
    result.Value = cs.Calc.Add(args.A, args.B)
    return nil
}

func (cs *CalculatorService) Subtract(args *Args, result *Result) error {
    result.Value = cs.Calc.Subtract(args.A, args.B)
    return nil
}

func (cs *CalculatorService) Multiply(args *Args, result *Result) error {
    result.Value = cs.Calc.Multiply(args.A, args.B)
    return nil
}

func (cs *CalculatorService) Divide(args *Args, result *Result) error {
    value, err := cs.Calc.Divide(args.A, args.B)
    if err != nil {
        result.Err = err.Error()
    } else {
        result.Value = value
    }
    return nil
}

func StartRPCServer() {
    calcService := &CalculatorService{Calc: &calculator.CalculatorImpl{}}
    rpc.Register(calcService)

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("Listen error:", err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Connection error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
