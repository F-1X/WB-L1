package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать
// набор из N воркеров, которые читают произвольные данные из канала и
// выводят в stdout. Необходима возможность выбора количества воркеров при
// старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
// способ завершения работы всех воркеров.
func main() {

	// вариант с использованием context.CancelFunc
	// variantCancel()

	// вариант с использованием вспомогательного канала done
	// variantDone()

	// вариант где горутина читает в цикле for из канала
	variantRange()

}

func variantCancel() {
	N := 34

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// канал из которого будут читать воркеры и куда будут записываться данные (может быть небуфиризированный или буфиризированный)
	ch := make(chan any, 100)

	ctx, cancel := context.WithCancel(context.Background())

	// для отмены контекста
	go func() {
		<-sig
		cancel()
	}()

	// wg нужен для грамотного ожидания завершения работы всех воркеров, чтобы дождаться их завершения перед закрытием программы
	var wg sync.WaitGroup
	wg.Add(N)

	// запуск worker pool на N воркеров
	for i := 0; i < N; i++ {
		go spawnWorker(ctx, i, ch, &wg)
	}

	// данные для отправки
	data := []interface{}{1, 2.3, "строка", true, struct{ name string }{name: "test"}}

	for {
		for _, v := range data {
			select {
			case <-ctx.Done(): // перед отправкой следует проверить а не закрыт ли контекст, иначе будет блокировка так как некому будет читать
				fmt.Println("main: context closed")
				close(ch)
				wg.Wait()
				return
			default:
				ch <- v // вынес в дефолт чтобы сперва проверился контекст, если контекст не закрыт то отправить в канал
			}
		}
	}

}

func spawnWorker(ctx context.Context, worker int, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case x, ok := <-ch:
			if ok { // если канал закрыт то можно завершить воркер
				fmt.Println("worker id", worker, "got:", x)
			} else {
				fmt.Println("worker id", worker, ":ch not ok")
				return // можем выйти если канал закрыт
			}
		case <-ctx.Done(): //  case выбирается случайно, но благодаря проверке ok лишний раз не будет проверяться этот case
			fmt.Println("worker id", worker, ":context closed")
			return
		}
	}
}

func variantDone() {
	// количество воркеров
	N := 14

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// канал из которого будут читать воркеры и куда будут записываться данные (может быть небуфиризированный или буфиризированный)
	ch := make(chan any)

	ctx, cancel := context.WithCancel(context.Background())

	// для оповещения воркеров о закрытии
	done := make(chan struct{}, N)

	go func() {
		<-sig
		for i := 0; i < N; i++ { // здесь необходимо отправить столько оповещений, сколько и воркеров, так как канал буфиризированный происходить чтение/запись должно быстрее
			done <- struct{}{}
		}
		cancel()
	}()

	// wg нужен для грамотного ожидания завершения работы всех воркеров, чтобы дождаться их завершения перед закрытием программы
	var wg sync.WaitGroup
	wg.Add(N)

	// запуск worker pool на N воркеров
	for i := 0; i < N; i++ {
		go workerWithDone(done, i, ch, &wg)
	}

	// данные для отправки
	data := []interface{}{1, 2.3, "строка", true, struct{ name string }{name: "test"}}

	// цикл отправителя, отправка данных в канал ch и проверка на закрытие контекста
	for {

		for _, v := range data {
			select {
			case <-ctx.Done(): // перед отправкой следует проверить а не закрыт ли контекст, иначе будет блокировка так как некому будет читать
				fmt.Println("main: context closed")
				close(ch)
				wg.Wait()
				return
			default:
				ch <- v // вынес в дефолт чтобы сперва проверился контекст, если контекст не закрыт то отправить в канал
			}
		}
	}
}

func workerWithDone(done chan struct{}, worker int, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case x, ok := <-ch:
			if ok { // если канал закрыт то можно завершить воркер
				fmt.Println("worker id", worker, "got:", x)
			} else {
				fmt.Println("worker id", worker, ":ch not ok")
				return // можем выйти если канал закрыт
			}
		case <-done: //  case выбирается случайно, но благодаря проверке ok лишний раз не будет проверяться этот case
			fmt.Println("worker id", worker, ":done recivied")
			return
		}
	}
}

func variantRange() {
	// количество воркеров
	N := 14

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// канал из которого будут читать воркеры и куда будут записываться данные (может быть небуфиризированный или буфиризированный)
	ch := make(chan any)

	ctx, cancel := context.WithCancel(context.Background())

	// для оповещения воркеров о закрытии
	done := make(chan struct{}, N)

	go func() {
		<-sig
		for i := 0; i < N; i++ { // здесь необходимо отправить столько оповещений, сколько и воркеров, так как канал буфиризированный происходить чтение/запись должно быстрее
			done <- struct{}{}
		}
		cancel()
	}()

	// wg нужен для грамотного ожидания завершения работы всех воркеров, чтобы дождаться их завершения перед закрытием программы
	var wg sync.WaitGroup
	wg.Add(N)

	// запуск worker pool на N воркеров
	for i := 0; i < N; i++ {
		go readerRange(i, ch, &wg)
	}

	// данные для отправки
	data := []interface{}{1, 2.3, "строка", true, struct{ name string }{name: "test"}}

	// цикл отправителя, отправка данных в канал ch и проверка на закрытие контекста
	for {
		for _, v := range data {
			select {
			case <-ctx.Done(): // перед отправкой следует проверить а не закрыт ли контекст, иначе будет блокировка так как некому будет читать
				fmt.Println("main: context closed")
				close(ch)
				wg.Wait()
				return
			default:
				ch <- v // вынес в дефолт чтобы сперва проверился контекст, если контекст не закрыт то отправить в канал
			}
		}
	}

}

func readerRange(worker int, ch chan any, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range ch {
		fmt.Println("worker id", worker, "got:", i)
	}

	fmt.Println("worker id", worker, "finished")
}
