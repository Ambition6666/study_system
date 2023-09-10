package pool

//-------------------------------------------创建简易的协程池--------------------------------------------------------
type Task struct {
	F    func(...any) any
	agrs []interface{}
}
type Pool struct {
	EmptyChan  chan *Task
	JobsChan   chan *Task
	Worker_num int
}

//定义一个全局协程池
var P *Pool

// 创建新任务
func NewTask(F func(...any) any, val ...any) *Task {
	t := new(Task)
	t.F = F
	t.agrs = val
	return t
}

// 执行任务
func (t *Task) exec() {
	t.F(t.agrs)
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

//运行协程池
func (p *Pool) Run() {
	for i := 1; i <= p.Worker_num; i++ {
		go p.Worker(i)
	}
	for v := range p.EmptyChan {
		p.JobsChan <- v
	}
}
