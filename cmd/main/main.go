package main

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-6/pkg/transactions"
	"math/rand"
	"time"
)

func main() {

	const users = 5
	const transitionsPerUser = 5

	transitionsSort := make([]int64, users*transitionsPerUser)
	for idx := range transitionsSort {
		transitionsSort[idx] = int64(rand.Intn(99))
	}

	task1 := transactions.Sort(transitionsSort)
	fmt.Println(task1)


	transitions := make([]transactions.Transaction, users*transitionsPerUser)
	for idx := range transitions {
		transitions[idx].Amount = int64(rand.Intn(99))
		transitions[idx].Date = time.Date(2020, time.Month(rand.Intn(12)), 1, 0, 0, 0, 0, time.Local)
	}

	total := transactions.SumConcurrently(transitions)
	fmt.Println(total)
}

