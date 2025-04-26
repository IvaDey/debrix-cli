package dbUtils

func FetchAll() *[]StoredTodo {
	db := OpenDb()

	var res []StoredTodo
	db.Find(&res)

	return &res
}
