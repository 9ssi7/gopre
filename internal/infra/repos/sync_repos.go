package repos

import "sync"

type syncRepo struct {
	mu sync.Mutex
}

func (r *syncRepo) Lock() {
	r.mu.Lock()
}

func (r *syncRepo) Unlock() {
	r.mu.Unlock()
}
