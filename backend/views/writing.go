package views

import (
	"fmt"
	"log"
	"my-blog/backend/common"
	"my-blog/backend/service"
	"net/http"
	"strconv"
)

/*
@author: shg
@since: 2022/10/19
@desc: //TODO
*/

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	r.ParseForm()
	println(r.Form.Get("id"))
	writing := common.Template.Writing
	writingRes, err := service.Write()
	if err != nil {
		log.Println("Writing:", err)
	}
	writing.WriteData(w, writingRes)
}

func (*HTMLApi) GetPost(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	r.ParseForm()
	pid, err := strconv.Atoi(r.Form.Get("id"))
	fmt.Println("pid  = ", pid)
	writingRes, err := service.GetPostDetail(pid)
	if err != nil {
		log.Println("Writing:", err)
	}
	writing.WriteData(w, writingRes)
}
