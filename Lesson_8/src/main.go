package main	

import (
	"log"
	"time"
	"errors"
	"sync"
)

func main() {
	RunParallel([]func()error {  
		makeTestTask(0, 1, false),
		makeTestTask(1, 1, false),
		makeTestTask(2, 2, true),
		makeTestTask(3, 3, true),
		makeTestTask(4, 2, false),
		makeTestTask(5, 3, false),
		makeTestTask(6, 1, false),
	}, 2, 2)
}

// RunParallel функция запускает задания паралелльно
func RunParallel(tasks []func() error, workersCount, maxErrorsCount int) {
	var wg sync.WaitGroup
	workerChannel := make(chan struct{}, workersCount)
	var errorsCount int
	log.Printf("%+v\n", "Запуск заданий...")
	for _, task := range tasks {		
			wg.Add(1)
			workerChannel <- struct{}{}
			go func(t func() error) {
				defer wg.Done()
				if t() != nil {					
					if errorsCount++; errorsCount>= maxErrorsCount {		
						log.Fatalf("Достигнуто максимальное количество ошибок!!")					
					}
				}			
				<-workerChannel
			}(task)		
	}
	wg.Wait()
	log.Println("Задания завершены!")
}

func makeTestTask(id int, sleepSeconds time.Duration, isError bool) func() error {
	return func() error {
		time.Sleep((time.Second * sleepSeconds))
		if isError {
			log.Printf("Задание %v завершено с ОШИБКОЙ\n", id)
			return errors.New("Тестовая ошибка")
		}
		log.Printf("Задание %v завершено успешно\n", id)
		return nil
	}
}




