package api

// import (
// 	"sync"

// 	"github.com/google/uuid"
// 	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/ptr"
// )

// type TaskStore struct {
// 	mu    sync.RWMutex
// 	tasks map[string]Task
// }

// var taskStore = NewTaskStore()

// func NewTaskStore() *TaskStore {
// 	return &TaskStore{
// 		tasks: make(map[string]Task, 0),
// 		mu:    sync.RWMutex{},
// 	}
// }

// func (s *TaskStore) List() []Task {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	return s.getSilceFromStore()
// }

// func (s *TaskStore) Add(task Task) *string {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	task.Id = ptr.To(uuid.NewString())
// 	s.tasks[*task.Id] = task
// 	return task.Id
// }

// func (s *TaskStore) Delete(id string) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	delete(s.tasks, id)
// }

// func (s *TaskStore) Get(id string) (Task, bool) {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	task, ok := s.tasks[id]
// 	if !ok {
// 		return Task{}, false
// 	}
// 	return task, true
// }

// func (s *TaskStore) getSilceFromStore() []Task {
// 	s.mu.RLock()
// 	defer s.mu.RUnlock()

// 	var tasks = make([]Task, 0, len(s.tasks))
// 	for _, task := range s.tasks {
// 		tasks = append(tasks, task)
// 	}
// 	return tasks

// }
