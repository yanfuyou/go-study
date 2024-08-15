package anonymous

import "log"

// Job 匿名组合了一个*log.Logger指针;以指针的方式从一个类型派生,
// 区别在于此类在实例化的时候需要外部提供一个对应指针类的实例
type Job struct {
	Command string
	*log.Logger
}

// Start 开始一个任务.可以在成员方法中方便的使用所有log.Logger提供的方法
func (job *Job) Start() {
	job.Logger.Println("job Starting")
	// doing
	job.Logger.Printf("doing:%s", job.Command)
	//
	job.Logger.Println("Job ended")
}
