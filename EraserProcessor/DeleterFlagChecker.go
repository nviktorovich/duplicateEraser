package EraserProcessor

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func RemoveOperator(m map[string][]string) error {
	errChan := make(chan error)
	workInfoChan := make(chan string)
	wg := sync.WaitGroup{}

	go func() {
		select {
		case workInfo := <-workInfoChan:
			fmt.Println(workInfo)
		case err := <-errChan:
			log.Print(err)
			os.Exit(1)
		}
	}()

	for k := range m {
		for i := range m[k] {
			wg.Add(1)
			go DeleteDuplicate(m[k][i], &wg, workInfoChan, errChan)
		}
	}
	wg.Wait()
	close(workInfoChan)

	return nil

}

func DeleteDuplicate(path string, wg *sync.WaitGroup, infoCh chan<- string, errCh chan<- error) {
	defer wg.Done()
	if err := os.Remove(path); err != nil {
		errCh <- err
		return
	}
	mess := "file: " + path + " was successfully deleted"
	infoCh <- mess

}
