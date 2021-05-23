package main

import (
	"math/rand"
	"sort"
	"time"
)

const (
	RedColor = "\033[1;31m"
	EndColor = "\u001B[0m"
)

type account struct {
	cash float64
}

type accounts []account

func (a accounts) Len() int           { return len(a) }
func (a accounts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a accounts) Less(i, j int) bool { return a[i].cash < a[j].cash }

const (
	N    = 50
	P    = .10
	GMax = 10000
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func task(accounts accounts) float64 {
	l := len(accounts)
	for i := 0; i < N-1; i++ {
		sort.Sort(sort.Reverse(accounts))
		accounts[l-i-2] = join(accounts[l-i-2], accounts[l-i-1])
		accounts = accounts[:len(accounts)-1]
	}
	return accounts[0].cash
}

func join(a, b account) account {
	sum := a.cash + b.cash
	return account{sum - (sum)*P}
}
