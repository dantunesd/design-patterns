package repository

type Articles struct {
	db map[int]string
}

func NewArticles() *Articles {
	return &Articles{
		db: map[int]string{},
	}
}

func (e *Articles) Save(content string) int {
	nextID := len(e.db) + 1
	e.db[nextID] = content

	return nextID
}

func (e *Articles) Get(id int) string {
	return e.db[id]
}
