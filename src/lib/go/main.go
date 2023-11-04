//go:build wasm
// +build wasm

package main

import (
	"syscall/js"

	"github.com/sprkweb/finaplan-cli/finaplan/pkg/finaplan"
)

func main() {
	js.Global().Set("initFinaplan", js.FuncOf(initPlan))

	<-make(chan struct{})
}

// jsFinaplan is a FinancialPlan wrapper for JS objects
type jsFinaplan struct {
	finaplan *finaplan.FinancialPlan
}

func initPlan(this js.Value, args []js.Value) any {
	var intervalType finaplan.IntervalType
	switch args[0].String() {
	case "days":
		intervalType = finaplan.Days
	case "weeks":
		intervalType = finaplan.Weeks
	case "months":
		intervalType = finaplan.Months
	case "years":
		intervalType = finaplan.Years
	default:
		panic("wrong interval type")
	}

	intervalLength := uint32(args[1].Int())
	intervalAmount := uint64(args[2].Int())

	jsPlan := &jsFinaplan{
		finaplan: finaplan.Init(&finaplan.PlanConfig{
			IntervalType:   intervalType,
			IntervalLength: intervalLength,
		}, intervalAmount),
	}
	return js.ValueOf(map[string]any{
		"add":    js.FuncOf(jsPlan.Add),
		"invest": js.FuncOf(jsPlan.Invest),
		"print":  js.FuncOf(jsPlan.Print),
	})
}

func (p *jsFinaplan) Add(this js.Value, args []js.Value) any {
	amount := finaplan.ProjectionUnit(args[0].Float())
	each := uint64(args[1].Int())
	start := uint64(args[2].Int())
	p.finaplan.Add(amount, each, start)
	return nil
}

func (p *jsFinaplan) Invest(this js.Value, args []js.Value) any {
	interest := args[0].Float()
	interval := uint64(args[1].Int())
	start := uint64(args[2].Int())
	compound := args[3].Bool()
	err := p.finaplan.Invest(interest, interval, start, compound)
	if err != nil {
		panic(err)
	}
	return nil
}

func (p *jsFinaplan) Print(this js.Value, args []js.Value) any {
	out := make([]any, 0, len(p.finaplan.Projection))
	for _, v := range p.finaplan.Projection {
		out = append(out, float64(v))
	}
	return out
}
