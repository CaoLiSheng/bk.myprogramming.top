package blogDir

type BlogDirRow struct {
	ID    int64  `db:"id" json:"id" form:"id" binding:"required"`
	PID   int64  `db:"pid" json:"pid" form:"pid"`
	Name  string `db:"name" json:"name" form:"name" binding:"required"`
	Alias string `db:"alias" json:"alias" form:"alias" binding:"required"`
}

type BlogDirRows []BlogDirRow

type AddBlogDirRow struct {
	ID    int64  `json:"id"`
	PID   int64  `json:"pid" form:"pid"`
	Name  string `json:"name" form:"name" binding:"required"`
	Alias string `json:"alias" form:"alias" binding:"required"`
}

type BlogDirRowId struct {
	ID    int64  `form:"id" binding:"required"`
}
