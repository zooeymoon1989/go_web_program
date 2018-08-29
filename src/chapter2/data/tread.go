package data

import "time"

type Thread struct {
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}

func Threads() (threads []Thread , err error)  {
	//查询数据
	rows , err := Db.Query("SELECT id ,uuid , topic , user_id , created_at FROM threads ORDER BY created_at DESC")
	if err != nil{
		return
	}

	//迭代查找记录
	for rows.Next() {
		//创建结构体实例
		th := Thread{}

		//把记录复制到th中
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId,&th.CreatedAt);err != nil{
			return
		}

		threads = append(threads , th)
	}

	rows.Close()

	return
}

/**
  通过id查询总数
 */
func (t *Thread) NumReplies() (count int) {

	rows , err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1" , t.Id)

	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count) ; err != nil {
			return
		}
	}

	rows.Close()

	return

}