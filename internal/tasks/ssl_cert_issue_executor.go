package tasks

import (
	"github.com/TeaOSLab/EdgeAPI/internal/db/models/acme"
	"github.com/TeaOSLab/EdgeAPI/internal/goman"
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/logs"
	"sync"
	"time"
)

func init() {
	dbs.OnReadyDone(func() {
		goman.New(func() {
			NewSSLCertIssueExecutor(5 * time.Second).Start()
		})
	})
}

// SSLCertIssueExecutor 证书签发任务
type SSLCertIssueExecutor struct {
	BaseTask
	ticker *time.Ticker
}

func NewSSLCertIssueExecutor(duration time.Duration) *SSLCertIssueExecutor {
	return &SSLCertIssueExecutor{
		ticker: time.NewTicker(duration),
	}
}

// Start 启动任务
func (this *SSLCertIssueExecutor) Start() {
	for range this.ticker.C {
		err := this.Loop()
		if err != nil {
			this.logErr("SSLCertExpireCheckExecutor", err.Error())
		}
	}
}

var (
	concurrent     = 10
	mutex          sync.Mutex
	queue          = make(chan *acme.ACMETask, 10)
	runningTaskIds = make(map[int64]bool)
)

func taskConsumer(wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range queue {
		mutex.Lock()
		runningTaskIds[int64(task.Id)] = true
		mutex.Unlock()

		go func(t *acme.ACMETask) {
			defer func() {
				mutex.Lock()
				delete(runningTaskIds, int64(t.Id))
				mutex.Unlock()
			}()

			ok, errMsg := acme.SharedACMETaskDAO.RunTaskAndAutoBindServer(nil, int64(t.Id), t.DecodeDomains())
			if !ok {
				logs.Printf(errMsg)
			}
		}(task)
	}
}

func taskProducer(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		time.Sleep(2 * time.Second)
		if len(runningTaskIds) < concurrent {
			excludeTasks := make([]int64, 0, len(runningTaskIds))
			for k := range runningTaskIds {
				excludeTasks = append(excludeTasks, k)
			}
			tasks, err := acme.SharedACMETaskDAO.FindIssueACMETask(nil, 2, int64(concurrent-len(runningTaskIds)), excludeTasks)
			if err != nil {
				logs.Printf("Failed to find tasks from database, %v", err)
				continue
			}
			if len(tasks) == 0 {
				continue
			}

			for _, task := range tasks {
				queue <- task
			}
		}
	}
}

// Loop 单次执行
func (this *SSLCertIssueExecutor) Loop() error {
	// 检查是否为主节点
	if !this.IsPrimaryNode() {
		return nil
	}
	var queue = make(chan *acme.ACMETask, concurrent)
	tasks, err := acme.SharedACMETaskDAO.FindIssueACMETask(nil, 2, int64(concurrent), nil)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}
	for _, task := range tasks {
		queue <- task
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go taskConsumer(&wg)
	go taskProducer(&wg)
	wg.Wait()
	return nil
}
