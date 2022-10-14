package models

import "my-blog/config"

/*
@author: shg
@since: 2022/10/13
@desc: //TODO
*/

type HomeResponse struct {
	config.Viewer
	Category []Category
	Posts    []PostMore
	Total    int
	Page     int
	Pages    []int
	PageEnd  bool
}
