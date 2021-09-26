package blogPost

import (
	"database/sql"
	"time"

	"bk.myprogramming.top/web/blogDir"
	"bk.myprogramming.top/web/blogTag"
)

type BlogPostRow struct {
	ID      int64               		 `json:"id"`
	Type    string						 `json:"type"`
	Name    string						 `json:"name"`
	Path    string						 `json:"path"`
	Date    time.Time					 `json:"date"`
	DirID   int64						 `json:"dirId"`
	DirDict map[int64]blogDir.BlogDirRow `json:"dirDict"`
	Tags    blogTag.BlogTagRows			 `json:"tags,omitempty"`
	Tag struct {
		TagId    sql.NullInt64
		TagName  sql.NullString
		TagAlias sql.NullString
	} `json:"-"`
}

type BlogPostRows []BlogPostRow

type AddBlogPostRow struct {
	ID      int64                        `json:"id"`
	Type    string                       `json:"type" form:"type" binding:"required"`
	Name    string                       `json:"name" form:"name" binding:"required"`
	DirID   int64                        `json:"dirId" form:"dirId" binding:"required"`
	DirDict map[int64]blogDir.BlogDirRow `json:"dirDict"`
	Path    string                       `json:"path"`
	Date    time.Time                    `json:"date"`
}

type ModifyBlogPostRow struct {
	ID      int64                        `json:"id" form:"id" binding:"required"`
	Type    string                       `json:"type" form:"type" binding:"required"`
	Name    string                       `json:"name" form:"name" binding:"required"`
	DirID   int64                        `json:"dirId" form:"dirId" binding:"required"`
	DirDict map[int64]blogDir.BlogDirRow `json:"dirDict"`
}

type BlogPostSaveForm struct {
	ID      int64  `form:"id" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type BlogPostTagRow struct {
	PostID int64 `form:"postId" binding:"required"`
	TagID  int64 `form:"tagId" binding:"required"`
}
