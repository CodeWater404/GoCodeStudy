package mysql

import "web_exercise_qimi/bluebell/models"

/**
  @author: CodeWater
  @since: 2023/11/14
  @desc: $
**/

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id , title , content , author_id , community_id) 
				values(? , ? , ? , ? , ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id , title , content , author_id , community_id , create_time 
			from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := "select post_id , title , content , author_id , community_id , create_time from  post limit ?,?"
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	/*
					1.posts = make([]*models.Post, 0, 2)虽然指定了容量2，但是size超过2的时候，会自动扩容，也就可以返回超过2条的数据
					2.在 db.Select(&posts, sqlStr, (page-1)*size, size) 中，&posts 传递的是 posts 切片的指针，因为 db.Select
				函数期望接收一个指向切片的指针，以便能够将查询结果填充到切片中。虽然 posts 已经是一个指向切片的指针，但在传递给函数时，
				你需要再次使用 & 操作符获取指向切片的指针，以确保 db.Select 函数可以将数据正确填充到切片中。这是因为 posts 是一个指向切片的指针，
				而 db.Select 期望接收一个指向指针的指针（**），所以需要传递 &posts。
					3.make([]int , 0 , 2) 和make([]int , 2)区别？
					make([]T, length, capacity) 初始化一个切片时，你可以显式指定切片的容量。这表示底层数组的大小将至少为指定的容量，
		即底层数组最多可以容纳 capacity 个元素。通常，capacity 的设置用于优化切片的性能，尤其是在你知道切片的最大元素数量时。
			make([]T, length) 初始化一个切片时，不显式指定容量。这意味着底层数组的大小将与切片的长度相等，即底层数组只能容纳 length 个元素。
		切片的容量将与长度相等，这可能会导致在添加元素时需要重新分配底层数组，以支持更多的元素，这可能会增加额外的性能开销。
			比较这两种方式：
			make([]T, length, capacity) 适用于你预先知道切片的最大元素数量，并希望在切片操作期间减少内部数组的重新分配次数。这可以提高性能，尤其是对于大型切片。
			make([]T, length) 适用于你不需要显式控制容量的情况，只关心切片的长度。这种方式会将容量设置为与长度相等，因此在切片长度达到容量时，需要重新分配内部数组以支持更多的元素。
	*/
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
