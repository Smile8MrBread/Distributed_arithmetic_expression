// отдельный пакет для установки времени операций
package operationTime

import (
	"time"
)

// структура для динамического html
type Operation struct {
	Operations []OperationTime
}

type OperationTime struct {
	Operation	string
	Time		time.Duration // время выполнения операции
}

var operators = Operation{
	Operations: []OperationTime{
		{Operation: "+", Time: time.Second},
		{Operation: "-", Time: time.Second},
		{Operation: "*", Time: time.Second},
		{Operation: "/", Time: time.Second},
	},
}

// блок функций для установки нового времени операциям
// решил оставить возможность установить отрацательное время :-)
func SetTimePlus(timeForPlus time.Duration) {
	operators.Operations[0].Time = timeForPlus
}

func SetTimeMinus(timeForMinus time.Duration) {
	operators.Operations[1].Time = timeForMinus
}

func SetTimeMult(timeForMult time.Duration) {
	operators.Operations[2].Time = timeForMult
}

func SetTimeDivision(timeForDivision time.Duration) {
	operators.Operations[3].Time = timeForDivision
}

// возврат всех операций для странички html
func GetOperators() Operation {
	return operators
}

// поотдельное получение операций для дальнейшего счета
func GetTimePlus() time.Duration {
	return operators.Operations[0].Time
}

func GetTimeMinus() time.Duration {
	return operators.Operations[1].Time
}

func GetTimeMult() time.Duration {
	return operators.Operations[2].Time
}

func GetTimeDivision() time.Duration {
	return operators.Operations[3].Time
}