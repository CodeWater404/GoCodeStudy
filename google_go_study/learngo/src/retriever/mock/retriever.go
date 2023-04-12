package mock

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/10
  @desc: $
**/

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
func (r *Retriever) GetPointer(url string) string {
	return r.Contents
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s", r.Contents)
}
