package articles

import (
	"fmt"
	"go-blog/models"
)

type articlesModel = models.Articles

func GetHomeArticles() (articles *articlesModel)  {
	var articlesModel articlesModel
	row, _ := articlesModel.GetLatestArticles(0, 2)

	fmt.Printf("\n%v\n", row)
	fmt.Printf("%v", articlesModel)

	if row == 0 {
		articles = nil
		return
	}
	return
}
