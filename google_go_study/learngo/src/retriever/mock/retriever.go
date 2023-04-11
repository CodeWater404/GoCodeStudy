package mock

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
