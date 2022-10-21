package models

import "my-blog/config"

/*
@author: shg
@since: 2022/10/21
@desc: //TODO
*/

type PigeonholeRes struct {
	config.Viewer
	config.System
	Categories []Category
	Lines      map[string][]Post
}
