package main

import (
	"bootcamp/aula_04/ex_4/sort"
	"bootcamp/aula_04/ex_4/types"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	small := rand.Perm(100)
	medium := rand.Perm(1000)
	large := rand.Perm(10000)

	fmt.Printf("Insertion: %+v\n", insertion(small, medium, large))
	fmt.Printf("Selection: %+v\n", selection(small, medium, large))
	fmt.Printf("Bubble: %+v\n", bubble(small, medium, large))

}

func insertion(a, b, c types.Values) time.Duration {
	first := make(chan struct{})
	second := make(chan struct{})
	third := make(chan struct{})

	t1 := time.Now()
	go func() {
		sort.Insertion(a)
		first <- struct{}{}
	}()
	go func() {
		<-first
		sort.Insertion(b)
		second <- struct{}{}
	}()
	go func() {
		<-second
		sort.Insertion(c)
		third <- struct{}{}
	}()
	<-third
	t2 := time.Now()

	return t2.Sub(t1)
}

func selection(a, b, c types.Values) time.Duration {
	first := make(chan struct{})
	second := make(chan struct{})
	third := make(chan struct{})

	t1 := time.Now()
	go func() {
		sort.Selection(a)
		first <- struct{}{}
	}()
	go func() {
		<-first
		sort.Selection(b)
		second <- struct{}{}
	}()
	go func() {
		<-second
		sort.Selection(c)
		third <- struct{}{}
	}()
	<-third
	t2 := time.Now()

	return t2.Sub(t1)
}

func bubble(a, b, c types.Values) time.Duration {
	first := make(chan struct{})
	second := make(chan struct{})
	third := make(chan struct{})

	t1 := time.Now()
	go func() {
		sort.Bubble(a)
		first <- struct{}{}
	}()
	go func() {
		<-first
		sort.Bubble(b)
		second <- struct{}{}
	}()
	go func() {
		<-second
		sort.Bubble(c)
		third <- struct{}{}
	}()
	<-third
	t2 := time.Now()

	return t2.Sub(t1)
}
