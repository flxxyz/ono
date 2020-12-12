package crontab

import (
	"github.com/flxxyz/ono/queue"
	"github.com/roylee0704/gron"
	"log"
	"sync"
	"time"
)

var cron *gron.Cron

func init() {
	log.Println("初始化任务作业...")
	cron = gron.New()
	cron.Start()
	cron.AddFunc(gron.Every(time.Microsecond), func() {
		var wg sync.WaitGroup

		for queue.Head() != nil {
			item := queue.Dequeue()

			wg.Add(1)
			go func() {
				item.(gron.Job).Run()
				wg.Done()
			}()

			wg.Wait()
		}
	})
}

// Once 添加一次性任务
func Once(job gron.Job) {
	queue.Enqueue(job)
}

// Add 添加需要执行的任务
func Add(schedule gron.Schedule, job gron.Job) {
	cron.Add(schedule, job)
}

// AddFunc 添加需要执行的匿名函数
func AddFunc(schedule gron.Schedule, job func()) {
	cron.AddFunc(schedule, job)
}
