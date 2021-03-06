package blogTag

type BlogTagRow struct {
	ID    int64  `db:"id" json:"id" form:"id" binding:"required"`
	Name  string `db:"name" json:"name" form:"name" binding:"required"`
	Alias string `db:"alias" json:"alias" form:"alias" binding:"required"`
}

type BlogTagRows []BlogTagRow

type AddBlogTagRow struct {
	ID    int64  `json:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Alias string `json:"alias" form:"alias" binding:"required"`
}

type BlogTagRowID struct {
	ID   int64  `form:"id" binding:"required"`
}
