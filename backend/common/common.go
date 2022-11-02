package common

import (
	"encoding/json"
	"io"
	"log"
	"my-blog/backend/config"
	models2 "my-blog/backend/models"
	"net/http"
	"sync"
)

/*
@author: shg
@since: 2022/10/14
@desc: //TODO
*/

var Template models2.HTMLTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	// 启动携程加载模版
	go func() {
		var err error
		Template, err = models2.InitTemplate(config.Config.System.CurrentDir + "/template/")
		if err != nil {
			log.Println(err)
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models2.Result
	result.Code = 200
	result.Msg = ""
	result.Data = data
	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Println("failed to convert result to json:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println("common success:", err)
	}
}

func Fail(w http.ResponseWriter, err error) {
	var result models2.Result
	result.Code = -999
	result.Msg = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
