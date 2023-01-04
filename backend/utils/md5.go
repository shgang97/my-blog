package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

/*
@author: shg
@since: 2022/12/4
@desc: //TODO
*/

func Md5Crypt(str string, salt ...interface{}) (Crypto string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
