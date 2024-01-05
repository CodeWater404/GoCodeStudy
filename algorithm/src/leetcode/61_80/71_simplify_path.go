package _1_80

/**
  @author: CodeWater
  @since: 2024/1/5
  @desc: 71. 简化路径
**/

func simplifyPath(path string) string {
	res, name := "", "" //name保存正常的一个路径
	//不是/结尾，加上一个/表示读到末尾了
	if path[len(path)-1] != '/' {
		path += "/"
	}
	for _, c := range path {
		if c != '/' {
			//正常目录名
			name += string(c)
		} else {
			//目录名读取完成后有三种情况：..
			if name == ".." {
				//回到上一级
				for len(res) > 0 && res[len(res)-1] != '/' {
					res = res[:len(res)-1]
				}
				if len(res) > 0 { //去掉“/”
					res = res[:len(res)-1]
				}
			} else if name != "." && name != "" { //name为空就是“///”这种情况
				res += "/" + name
			} //第三种情况就是name是“.”和“”，不需要做处理
			name = ""
		}
	}

	if len(res) == 0 {
		res = "/"
	}
	return res
}
