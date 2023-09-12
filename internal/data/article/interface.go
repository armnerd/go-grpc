package article

type IArticleDao interface {
	GetAll(start int, category int, timeline int, search string) ([]Article, error)
	GetTotal(category int, timeline int, search string) int64
	GetOne(id int) (Article, error)
}
