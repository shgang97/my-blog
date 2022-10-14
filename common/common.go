package common

import (
	"log"
	"my-blog/config"
	"my-blog/models"
	"sync"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

var Template models.HTMLTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	// 启动携程加载模版
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Config.System.CurrentDir + "/template/")
		if err != nil {
			log.Println(err)
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
