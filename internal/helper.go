package internal

import (
	"errors"
	"fmt"
	"time"
)

var wStFormat = "2006-01-02 15:04:05"

func MessageLog(format string, a ...any) (n int, err error) {
	timeString := time.Now().UTC().Format(wStFormat)
	args := []any{timeString, APP_NAME}
	args = append(args, a...)
	return fmt.Printf("%s [%s] "+format+"\n", args...)
}

func MessageWithPrefixLog(prefix string, format string, a ...any) (n int, err error) {
	timeString := time.Now().UTC().Format(wStFormat)
	args := []any{timeString, APP_NAME, prefix}
	args = append(args, a...)
	return fmt.Printf("%s [%s]%s "+format+"\n", args...)
}

func GetLoadRacer(requests int64, racer int64) (*[]int64, error) {
	if requests < 0 {
		return nil, errors.New("valor invalido para os requests")
	}

	if racer < 0 {
		return nil, errors.New("valor invalido para os racers")
	}

	baseLoad := requests / racer
	remainder := requests % racer

	loadPerRacer := make([]int64, racer)
	for i := int64(0); i < racer; i++ {
		loadPerRacer[i] = baseLoad
		if i < remainder {
			loadPerRacer[i]++
		}
	}

	return &loadPerRacer, nil
}
