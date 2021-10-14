package model

import (
	"go-blog/model/entity"
	"go-blog/utils/mylog"
)

type Markdown = entity.Markdown

type MarkdownModel struct {

}

// InsertMarkdown 插入markdown
func (m *MarkdownModel) InsertMarkdown(markdown *Markdown) (articleId int)  {
	res := DB.Create(markdown)

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}

	if res.RowsAffected == 0 {
		articleId = -1
		return
	}
	articleId = markdown.ArticleId
	return
}

// FindMarkdownByArticleId 根据文章id获取文章的markdown内容
func (m *MarkdownModel) FindMarkdownByArticleId(articleId int) (markdown *Markdown) {
	var findMarkdown Markdown

	res := DB.Raw("select * from markdown where article_id = ? ;", articleId).Scan(&findMarkdown)

	if res.RowsAffected == 0 {
		return
	}

	if res.Error != nil {
		mylog.Error.Printf("execute sql error, message: %v", res.Error.Error())
		return
	}
	markdown = &findMarkdown
	return
}