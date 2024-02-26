package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

var httpClient = &http.Client{}

type stressTestRacer struct {
	name         string
	url          string
	requestsToDo int64
	requestsDone int64
	relatorio    *Relatorio
}

func NewRacer(name string, url string, requests int64, relatorio *Relatorio) *stressTestRacer {
	Racer := &stressTestRacer{}
	Racer.name = name
	Racer.url = url
	Racer.requestsToDo = requests
	Racer.relatorio = relatorio
	return Racer
}

func (r *stressTestRacer) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	logPrefix := fmt.Sprintf("[Testando Racer %s]", strings.ToUpper(r.name))
	log := func(format string, a ...any) {
		MessageWithPrefixLog(logPrefix, format, a...)
	}

	log("Iniciando com %d requests", r.requestsToDo)

	for i := int64(0); i < r.requestsToDo; i++ {
		statusCode, err := r.doRequest(ctx)
		if err != nil {
			log("erro: %s", err.Error())
			r.relatorio.AddError()
			continue
		}

		r.relatorio.AddRequest(statusCode)
		r.requestsDone++
	}

	log("Finalizado (%d/%d)", r.requestsDone, r.requestsToDo)
}

func (r *stressTestRacer) doRequest(ctx context.Context) (int, error) {
	request, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return 0, err
	}
	request.Close = true
	request = request.WithContext(ctx)

	response, err := httpClient.Do(request)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, response.Body)
	defer response.Body.Close()

	return response.StatusCode, nil
}
