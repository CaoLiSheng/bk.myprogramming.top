package blogTag

type BlogTag struct {
	ID   int64  `db:"id" json:"id" form:"id" binding:"required"`
	Name string `db:"name" json:"name" form:"name" binding:"required"`
}

type BlogTags []BlogTag

type AddBlogTag struct {
	ID   int64  `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}

type BlogTagID struct {
	ID   int64  `form:"id" binding:"required"`
}
