package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/4/16
  @desc: 统一的业务逻辑处理
**/

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, "/list/") != 0 {
		return userError("path must start: " + "/list/")
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//1.普通的不做处理
		//panic(err)

		//2.一般错误处理
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//return

		//3.错误类型处理,有错就返回，外面一层会有统一的处理
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
