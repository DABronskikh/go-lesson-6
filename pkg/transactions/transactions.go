package transactions

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Transaction struct {
	Amount int64
	Date   time.Time
}

func Sum(transitions []Transaction) int64 {
	sum := int64(0)
	for _, v := range transitions {
		sum += v.Amount
	}

	return sum
}

func SumConcurrently(transactions []Transaction) int64 {
	transactionsDateGroup := map[string][]Transaction{}

	for _, v := range transactions {
		period := v.Date.Format("2006-01")
		_, ok := transactionsDateGroup[period]
		if !ok {
			transactionsDateGroup[period] = []Transaction{v}
		} else {
			transactionsDateGroup[period] = append(transactionsDateGroup[period], v)
		}
	}

	goroutines := len(transactionsDateGroup)
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)

	for k, v := range transactionsDateGroup {
		go func(k string, v []Transaction) {
			sum := Sum(v)
			total += sum
			fmt.Printf("date: %v sum: %v\n", k, sum)
			wg.Done()
		}(k, v)
	}

	wg.Wait()
	return total
}

func Sort(transitions []int64) []int64 {
	sort.SliceStable(transitions, func(i, j int) bool {
		return transitions[i] > transitions[j]
	})
	return transitions
}
