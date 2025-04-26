package dbUtils

func RemoveCompleted() {
	db := OpenDb()

	db.Delete(&StoredTodo{}, "is_completed = true")
}
