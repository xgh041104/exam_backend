package lib

import (
	model "StudyExamPlatformAPI/Model"
	"errors"
	"fmt"
	"sync"

	"gopkg.in/ini.v1"
)

var config *Database
var once sync.Once

var dsn string

type Database struct {
	Type              string
	User              string
	Password          string
	Host              string
	Name              string
	Path              string
	YzAccount         string
	YzOnlyMark        string
	YzYLH             string
	YzMAC             string
	ENDPOINT          string
	ACCESS_KEY_ID     string
	ACCESS_KEY_SECRET string
	BACKET_NAME       string
	OSShttp           string
}

// 通过单例模式初始化全局配置
func LoadConfig() *Database {
	once.Do(func() {
		var DatabaseSetting = &Database{}
		var cfg *ini.File
		cfg, _ = ini.Load("conf.ini")
		cfg.Section("database").MapTo(DatabaseSetting)
		config = DatabaseSetting
	})
	return config
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

type MQueue struct {
	Front      int
	Rear       int
	QueueArray []*model.FileInfo
	Mutex      sync.Mutex
}

var queue *MQueue
var onceQueue sync.Once

func LoadQueue() *MQueue {

	onceQueue.Do(func() {
		queue = &MQueue{
			Front:      0,
			Rear:       -1,
			QueueArray: []*model.FileInfo{},
		}
	})
	return queue
}

// 队列-入队
func (q *MQueue) AddQueue(v *model.FileInfo) error {
	var mutex sync.Mutex
	if mutex.TryLock() {
		q.Rear += 1                            //队尾下标+1
		q.QueueArray = append(q.QueueArray, v) //数据插入队尾
		mutex.Unlock()

	}

	return nil
}

// 队列-出队
func (q *MQueue) DeleteQueue() (*model.FileInfo, error) {

	var mutex sync.Mutex
	if mutex.TryLock() {

		//判断队列否为空
		file := new(model.FileInfo)
		file = nil
		if len(q.QueueArray) == 0 {
			return file, errors.New("队列为空")
		}

		v := q.QueueArray[0] //获取队列头部元素值
		q.QueueArray = append(q.QueueArray[:q.Front], q.QueueArray[q.Front+1:]...)
		q.Front += 1 //队头下标+1
		mutex.Unlock()
		return v, nil
	}
	return nil, nil
}

var onceChan sync.Once
var fileinfochanmutex *MutexFileChan

type MutexFileChan struct {
	Mutex        sync.Mutex
	Fileinfochan chan *model.FileInfo
}

func LoadChan() *MutexFileChan {
	onceChan.Do(func() {
		fileinfochanmutex = new(MutexFileChan)
		fileinfochanmutex.Fileinfochan = make(chan *model.FileInfo, 1000)

	})
	return fileinfochanmutex
}

type MutexOfficeChan struct {
	Mutex     sync.Mutex
	OfficePdf chan *OfficePdf
}

type OfficePdf struct {
	InputFile  string
	OutputFile string
}

var onceofficeChan sync.Once
var officechanmutex *MutexOfficeChan

func LoadOfficeChan() *MutexOfficeChan {
	onceofficeChan.Do(func() {
		officechanmutex = new(MutexOfficeChan)
		officechanmutex.OfficePdf = make(chan *OfficePdf, 1000)

	})
	return officechanmutex

}
