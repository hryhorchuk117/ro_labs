package main

import (
	"fmt"
	"sync"
	"time"
)

const NORTH = 1
const SOUTH = -1

func train(
	index int,
	comingFromNorth bool,
	startOnA bool,
	movingInA *int,
	movingInB *int,
	aCheckLock *sync.Mutex,
	bCheckLock *sync.Mutex,
	wg *sync.WaitGroup) {

	waitingOnA := startOnA

	waitOnTunnel:
	for true {
		fmt.Printf("%d: waiting from A? %v, from north? %v\n", index, waitingOnA, comingFromNorth)
		if waitingOnA && comingFromNorth {
			startWaitingTime := time.Now()

			for true {
				success := false

				aCheckLock.Lock()
				if *movingInA >= 0 {
					*movingInA = *movingInA + NORTH
					success = true
				}
				aCheckLock.Unlock()

				if !success {
					time.Sleep(50 * time.Millisecond)
					//runtime.Gosched()
					if time.Now().Unix()-startWaitingTime.Unix() > 500 {
						fmt.Printf("%d: waiting too long\n", index)
						waitingOnA = !waitingOnA
						continue waitOnTunnel
					}
				} else {
					fmt.Printf("%d: moving from north in A!\n", index)
					time.Sleep(1000 * time.Millisecond)
					fmt.Printf("%d: arrived from north in A!\n", index)

					aCheckLock.Lock()
					*movingInA = *movingInA - NORTH
					aCheckLock.Unlock()

					break waitOnTunnel
				}
			}
		}
		if waitingOnA && !comingFromNorth {
			startWaitingTime := time.Now()

			for true {
				success := false

				aCheckLock.Lock()
				if *movingInA <= 0 {
					*movingInA = *movingInA + SOUTH
					success = true
				}
				aCheckLock.Unlock()

				if !success {
					time.Sleep(50 * time.Millisecond)
					if time.Now().Unix()-startWaitingTime.Unix() > 500 {
						fmt.Printf("%d: waiting too long\n", index)
						waitingOnA = !waitingOnA
						continue waitOnTunnel
					}
				} else {
					fmt.Printf("%d: moving from south in A!\n", index)
					time.Sleep(1000 * time.Millisecond)
					fmt.Printf("%d: arrived from south in A!\n", index)

					aCheckLock.Lock()
					*movingInA = *movingInA - SOUTH
					aCheckLock.Unlock()

					break waitOnTunnel
				}
			}
		}
		if !waitingOnA && comingFromNorth {
			startWaitingTime := time.Now()

			for true {
				success := false

				bCheckLock.Lock()
				if *movingInB >= 0 {
					*movingInB = *movingInB + NORTH
					success = true
				}
				bCheckLock.Unlock()

				if !success {
					time.Sleep(50 * time.Millisecond)
					//runtime.Gosched()
					if time.Now().Unix()-startWaitingTime.Unix() > 500 {
						fmt.Printf("%d: waiting too long\n", index)
						waitingOnA = !waitingOnA
						continue waitOnTunnel
					}
				} else {
					fmt.Printf("%d: moving from north in B!\n", index)
					time.Sleep(1000 * time.Millisecond)
					fmt.Printf("%d: arrived from north in B!\n", index)

					bCheckLock.Lock()
					*movingInB = *movingInB - NORTH
					bCheckLock.Unlock()

					break waitOnTunnel
				}
			}
		}
		if !waitingOnA && !comingFromNorth {
			startWaitingTime := time.Now()

			for true {
				success := false

				bCheckLock.Lock()
				if *movingInB <= 0 {
					*movingInB = *movingInB + SOUTH
					success = true
				}
				bCheckLock.Unlock()

				if !success {
					time.Sleep(50 * time.Millisecond)
					//runtime.Gosched()
					if time.Now().Unix()-startWaitingTime.Unix() > 500 {
						fmt.Printf("%d: waiting too long\n", index)
						waitingOnA = !waitingOnA
						continue waitOnTunnel
					}
				} else {
					fmt.Printf("%d: moving from south in B!\n", index)
					time.Sleep(1000 * time.Millisecond)
					fmt.Printf("%d: arrived from south in B!\n", index)

					bCheckLock.Lock()
					*movingInB = *movingInB - SOUTH
					bCheckLock.Unlock()

					break waitOnTunnel
				}
			}
		}
	}
	wg.Done()
}

func main() {
	movingInA, movingInB := 0, 0
	aCheckLock := sync.Mutex{}
	bCheckLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(6)
	go train(0, true, true, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	go train(1, true, true, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	go train(2, true, true, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	go train(3, false, true, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	go train(4, true, false, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	go train(5, false, false, &movingInA, &movingInB, &aCheckLock, &bCheckLock, &wg)
	wg.Wait()
}
