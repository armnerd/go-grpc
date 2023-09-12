package article

import (
	"go-grpc/internal/base"

	"gorm.io/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db base.BlogDB) IArticleDao {
	return &ArticleDao{db: db}
}

func (a *ArticleDao) GetAll(start int, category int, timeline int, search string) ([]Article, error) {
	var data = make([]Article, 0)

	// 参数绑定
	var params = &Article{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := a.db.Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	err := tx.Offset(start).Limit(6).Order("id desc").Find(&data).Error
	return data, err
}

// GetTotal 获取文章总数
func (a *ArticleDao) GetTotal(category int, timeline int, search string) int64 {
	var count int64

	// 参数绑定
	var params = &Article{}
	if category != 0 {
		params.Type = category
	}
	if timeline != 0 {
		params.Timeline = timeline
	}
	tx := a.db.Model(&Article{}).Where(params)

	// 搜索参数
	if search != "" {
		tx = tx.Where("title LIKE ?", "%"+search+"%")
	}
	tx.Count(&count)

	return count
}

// GetOne 获取文章详情
func (a *ArticleDao) GetOne(id int) (Article, error) {
	var data = Article{}
	err := a.db.First(&data, id).Error
	return data, err
}
