package repository

type ArticlesProxy struct {
	articles       *Articles
	lastInsertedID int
}

func NewArticlesProxy(article *Articles) *ArticlesProxy {
	return &ArticlesProxy{
		articles: article,
	}
}

func (a *ArticlesProxy) Save(content string) int {
	lastID := a.articles.Save(content)
	a.lastInsertedID = lastID
	return lastID
}

func (a *ArticlesProxy) Get(id int) string {
	if id > a.lastInsertedID {
		return "unknown content, try another"
	}

	return a.articles.Get(id)
}
