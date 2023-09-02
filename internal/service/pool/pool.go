package pool

//-------------------------------------------创建简易的协程池--------------------------------------------------------
type Task struct {
	F func()
}
type Pool struct {
	EmptyChan  chan *Task
	JobsChan   chan *Task
	Worker_num int
}

var P *Pool

// 创建新任务
func NewTask(F func()) *Task {
	t := new(Task)
	t.F = F
	return t
}

// 执行
func (t *Task) exec() {
	t.F()
}

// 创建新协程池
func NewPool() *Pool {
	p := new(Pool)
	p.EmptyChan = make(chan *Task, 10)
	p.JobsChan = make(chan *Task, 10)
	p.Worker_num = 3
	return p
}

// 协程工作者
func (p *Pool) Worker(id int) {
	for v := range p.JobsChan {
		v.exec()
	}
}
func (p *Pool) Run() {
	for i := 1; i <= p.Worker_num; i++ {
		go p.Worker(i)
	}
	for v := range p.EmptyChan {
		p.JobsChan <- v
	}
}
