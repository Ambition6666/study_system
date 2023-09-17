package pool

import "studysystem/logs"

//-------------------------------------------创建简易的协程池--------------------------------------------------------
type Task struct {
	F      func(...interface{}) []interface{}
	agrs   []interface{}
	Result chan interface{}
}
type Pool struct {
	EmptyChan  chan *Task
	JobsChan   chan *Task
	Worker_num int
}

//定义一个全局协程池
var P *Pool

// 创建新任务
func NewTask(F func(...interface{}) []interface{}, val ...interface{}) *Task {
	t := new(Task)
	t.F = F
	t.agrs = val
	t.Result = make(chan any, 10)
	return t
}
func (t *Task) Close() {
	close(t.Result)
}

// 执行任务
func (t *Task) exec() {
	rval := t.F(t.agrs)
	err := rval[len(rval)-1].(error)
	if err != nil {
		t.Close()
		logs.SugarLogger.Debugf("协程中的错误:%v", err)
		return
	}
	t.Result <- rval[0]
	t.Close()
}

// 创建新协程池
func NewPool() *Pool {
	p := new(Pool)
	p.EmptyChan = make(chan *Task, 10)
	p.JobsChan = make(chan *Task, 10)
	p.Worker_num = 3
	return p
}

//关闭协程池
func (p *Pool) Close() {
	close(p.EmptyChan)
	close(p.JobsChan)
}

// 协程工作者
func (p *Pool) Worker(id int) {
	for v := range p.JobsChan {
		v.exec()
	}
}

//运行协程池
func (p *Pool) Run() {
	for i := 1; i <= p.Worker_num; i++ {
		go p.Worker(i)
	}
	for v := range p.EmptyChan {
		p.JobsChan <- v
	}
}
