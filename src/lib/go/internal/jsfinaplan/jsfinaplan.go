package jsfinaplan

import (
	"strconv"
	"syscall/js"

	"github.com/sprkweb/finaplan-web/src/lib/go/internal/jsresult"
	"github.com/sprkweb/finaplan/pkg/finaplan"
)

// jsFinaplan is a FinancialPlan wrapper for JS objects
type jsFinaplan struct {
	finaplan *finaplan.FinancialPlan
}

func InitPlan(this js.Value, args []js.Value) any {
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
	intervalAmount := uint32(args[2].Int())

	jsPlan := &jsFinaplan{
		finaplan: finaplan.Init(&finaplan.PlanConfig{
			IntervalType:   intervalType,
			IntervalLength: intervalLength,
		}, intervalAmount),
	}
	return js.ValueOf(map[string]any{
		"add":       js.FuncOf(jsPlan.Add),
		"invest":    js.FuncOf(jsPlan.Invest),
		"inflation": js.FuncOf(jsPlan.Inflation),
		"print":     js.FuncOf(jsPlan.Print),
	})
}

func (p *jsFinaplan) Add(this js.Value, args []js.Value) any {
	amount := args[0].String()
	each := uint32(args[1].Int())
	start := uint32(args[2].Int())
	err := p.finaplan.Add(amount, each, start)
	return jsresult.Result(nil, err)
}

func (p *jsFinaplan) Invest(this js.Value, args []js.Value) any {
	interest := args[0].String()
	interval := uint32(args[1].Int())
	start := uint32(args[2].Int())
	compound := args[3].Bool()
	err := p.finaplan.Invest(interest, interval, start, compound)
	return jsresult.Result(nil, err)
}

func (p *jsFinaplan) Inflation(this js.Value, args []js.Value) any {
	inflation := args[0].String()
	intervals := uint32(args[1].Int())
	err := p.finaplan.Inflation(inflation, intervals)
	return jsresult.Result(nil, err)
}

func (p *jsFinaplan) Print(this js.Value, args []js.Value) any {
	projection := p.finaplan.PrettyPrint()
	out := make([]any, 0, len(projection))
	for _, v := range projection {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return jsresult.Result(nil, err)
		}
		out = append(out, num)
	}
	return jsresult.Result(out, nil)
}
