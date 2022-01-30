package loader

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"messages_api/internal/messages/model"
	"messages_api/internal/messages/repository"
	"os"
	"runtime"
	"sync"
	"time"
)

var workers = runtime.NumCPU()

type messageLoader struct {
	messagesRepository repository.Repository
	filePath           string
}

func NewMessageLoader(messagesRepository repository.Repository, filePath string) *messageLoader {
	return &messageLoader{
		messagesRepository: messagesRepository,
		filePath:           filePath,
	}
}

func (l *messageLoader) Load(errChan chan error) {
	var (
		wg          sync.WaitGroup
		workerInput = make(chan *model.Message)
	)

	for i := 0; i < workers; i++ {
		go l.worker(&wg, workerInput, errChan)
	}

	csvFile, err := os.Open(l.filePath)
	if err != nil {
		errChan <- err
		return
	}

	reader := csv.NewReader(csvFile)
	_, _ = reader.Read() //skip header

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			errChan <- err
			return
		}

		if len(record) != 5 {
			log.Println("invalid row, skipping")
			continue
		}

		creationDate, err := time.Parse(time.RFC3339, record[4])
		if err != nil {
			log.Printf("failed to parse date %s with error: %s, skipping", record[4], err.Error())
			continue
		}

		workerInput <- &model.Message{
			ID:           record[0],
			Name:         record[1],
			Email:        record[2],
			Text:         record[3],
			CreationDate: creationDate,
		}
	}

	close(workerInput)
	wg.Wait()
}

func (l *messageLoader) worker(wg *sync.WaitGroup, input <-chan *model.Message, errChan chan error) {
	wg.Add(1)
	defer wg.Done()

	for msg := range input {
		err := l.messagesRepository.Store(context.Background(), msg)
		if err != nil {
			errChan <- err
		}
	}
}
