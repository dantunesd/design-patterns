package main

import (
	"design-patterns/proxy/repository"
)

func main() {
	articleRepository := repository.NewArticles()
	articleRepository.Save("content one")

	println(articleRepository.Get(1))
	println(articleRepository.Get(2))

	articleRepositoryProxy := repository.NewArticlesProxy(repository.NewArticles())
	articleRepositoryProxy.Save("content one")

	println(articleRepositoryProxy.Get(1))
	println(articleRepositoryProxy.Get(2))
}
