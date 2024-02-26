package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Relatorio struct {
	URL             string
	Racer           int64
	Duration        *time.Duration
	RequestCount    int64
	StatusCodeCount map[int]int64
	ErrorCount      int64
	mutex           sync.Mutex
}

func NewRelatorio(url string, racer int64) *Relatorio {
	Relatorio := &Relatorio{}
	Relatorio.URL = url
	Relatorio.Racer = racer
	Relatorio.StatusCodeCount = map[int]int64{}
	return Relatorio
}

func (r *Relatorio) AddRequest(statusCode int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.RequestCount++
	currentCount, ok := r.StatusCodeCount[statusCode]
	if !ok {
		r.StatusCodeCount[statusCode] = 1
	} else {
		r.StatusCodeCount[statusCode] = currentCount + 1
	}
}

func (r *Relatorio) AddError() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.ErrorCount++
}

func (r *Relatorio) PrintRelatorio(finished bool) {
	duration := "-"
	if r.Duration != nil {
		if r.Duration.Seconds() >= 1 {
			duration = fmt.Sprintf("%.1fs", r.Duration.Seconds())
		} else {
			duration = fmt.Sprintf("%dms", r.Duration.Milliseconds())
		}
	}

	total2XXRequests := int64(0)
	for key, value := range r.StatusCodeCount {
		if key >= 200 && key <= 299 {
			total2XXRequests = total2XXRequests + value
		}
	}

	twInner := table.NewWriter()
	twInner.SetStyle(table.StyleLight)
	twInner.Style().Options.SeparateRows = true
	twInner.Style().Options.SeparateColumns = true
	twInner.Style().Options.DrawBorder = true

	if len(r.StatusCodeCount) == 0 {
		twInner.AppendRows([]table.Row{{"-", "-"}})
	}

	for key, value := range r.StatusCodeCount {
		twInner.AppendRows([]table.Row{{key, value}})
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleColoredDark)
	tw.Style().Options.SeparateRows = true
	tw.Style().Options.SeparateColumns = true
	tw.Style().Options.DrawBorder = true

	if finished {
		tw.SetTitle(fmt.Sprintf("Relatório de testes do %s", APP_NAME))
	} else {
		tw.SetTitle(fmt.Sprintf("Relatório de testes %s foi interrompido", APP_NAME))
	}

	tw.AppendRows([]table.Row{
		{"URL", r.URL},
		{"Racer", r.Racer},
		{"Duração", duration},
		{"Requests", r.RequestCount},
		{"2XX Requests", total2XXRequests},
		{"Requests por Status", twInner.Render()},
		{"Erros", r.ErrorCount},
	})

	fmt.Println(tw.Render())
}
