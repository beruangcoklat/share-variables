package repository

import "sync"

type Repository struct {
	sync.Mutex
	allVariables map[string]string
}

var (
	once sync.Once
	repo *Repository
)

func GetRepository() *Repository {
	once.Do(func() {
		repo = &Repository{
			allVariables: make(map[string]string),
		}
	})
	return repo
}

func (r *Repository) UpdateVariable(key, value string) {
	r.Lock()
	defer r.Unlock()
	r.allVariables[key] = value
}

func (r *Repository) GetVariable(key string) string {
	return r.allVariables[key]
}
