package main

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-6/pkg/transactions"
	"math/rand"
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

}
