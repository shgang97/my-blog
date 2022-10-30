package models

import (
	"my-blog/backend/config"
)

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/

type HomeResponse struct {
	config.Viewer
	Categories []Category
	Posts      []PostMore
	Total      int
	Page       int
	Pages      []int
	PageEnd    bool
}
