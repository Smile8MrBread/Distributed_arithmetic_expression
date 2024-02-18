// реализация замысла со счетом выражения за определенное время
package logic

import (
	"Distributed_arithmetic_expression_evaluator/database/operationTime"
	"strings"
	"sync"
	"time"
)

var execTime = map[string]time.Time{} // массив с требуемым временем на исполнение
var mu = sync.Mutex{}

func ExecutionTime(id, form string) {
	pluses := strings.Count(form, "+")
	minuses := strings.Count(form, "-")
	multes := strings.Count(form, "*")
	divisions := strings.Count(form, "/")

	// счет требуемого времени
	mu.Lock()
	endTiime := time.Now().Add(time.Duration(pluses) * operationTime.GetTimePlus() + time.Duration(minuses) * operationTime.GetTimeMinus() + time.Duration(multes) * operationTime.GetTimeMult() + time.Duration(divisions) * operationTime.GetTimeDivision())
	execTime[id] = endTiime
	mu.Unlock()
}

// возвращает время, которое осталось
func GetExecTime(id string) time.Duration {
	defer mu.Unlock()

	mu.Lock()
	return time.Until(execTime[id])
}