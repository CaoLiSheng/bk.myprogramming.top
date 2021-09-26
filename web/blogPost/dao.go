package blogPost

import (
	"time"

	"bk.myprogramming.top/db"
	srv "bk.myprogramming.top/server"
	"bk.myprogramming.top/utils"
	"bk.myprogramming.top/web/blogDir"
	"bk.myprogramming.top/web/blogTag"
	"github.com/huandu/go-sqlbuilder"
)

func (data *BlogPostRows) ListBlogPost(c *db.Core) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("a.id", "a.type", "a.name", "a.path", "a.date", "a.dir_id", "c.id", "c.name", "c.alias")
	sb.From(sb.As("blog_post", "a"))
	sb.JoinWithOption(sqlbuilder.LeftJoin, sb.As("blog_post_tags", "b"), "a.id=b.post_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, sb.As("blog_tag", "c"), "b.tag_id=c.id")
	sql, args := sb.Build()
	rows, err := c.DB.QueryxContext(*c.Ctx, sql, args...)

	srv.IsPanic(err)

	dict := make(map[int64]BlogPostRow)
	for rows.Next() {
		var r BlogPostRow
		srv.IsPanic(rows.Scan(&r.ID, &r.Type, &r.Name, &r.Path, &r.Date, &r.DirID, &r.Tag.TagId, &r.Tag.TagName, &r.Tag.TagAlias))
		if row, ok := dict[r.ID]; ok {
			if r.Tag.TagId.Valid {
				row.Tags = append(row.Tags, blogTag.BlogTagRow{
					ID: r.Tag.TagId.Int64,
					Name: r.Tag.TagName.String,
					Alias: r.Tag.TagAlias.String,
				})
			}
		} else {
			r.DirDict = (&blogDir.BlogDirRowId{ID: r.DirID}).GetBlogDirDict(c)
			if r.Tag.TagId.Valid {
				r.Tags = []blogTag.BlogTagRow{{
					ID: r.Tag.TagId.Int64,
					Name: r.Tag.TagName.String,
					Alias: r.Tag.TagAlias.String,
				}}
			}
			dict[r.ID] = r
		}
	}

	for _, row := range dict {
		*data = append(*data, row)
	}
}

func (row *AddBlogPostRow) AddBlogPost(c *db.Core) {
	row.Path = utils.Pinyin(row.Name) + ".md"
	row.Date = time.Now()
	
	sb := sqlbuilder.NewInsertBuilder().InsertInto("blog_post")
	sb.Cols("type", "name", "path", "date", "dir_id")
	sb.Values(row.Type, row.Name, row.Path, row.Date, row.DirID)
	sql, args := sb.Build()
	result := c.MustExec(sql, args...)
	id, err := result.LastInsertId()

	srv.IsPanic(err)

	row.ID = id
	row.DirDict = (&blogDir.BlogDirRowId{ID: row.DirID}).GetBlogDirDict(c)
}

func (row *ModifyBlogPostRow) ModifyBlogPost(c *db.Core) {
	sb := sqlbuilder.NewUpdateBuilder().Update("blog_post")
	sb.Set(sb.Assign("type", row.Type), sb.Assign("name", row.Name), sb.Assign("dir_id", row.DirID), sb.Assign("date", time.Now()))
	sql, args := sb.Where(sb.E("id", row.ID)).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()
	
	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}

	row.DirDict = (&blogDir.BlogDirRowId{ID: row.DirID}).GetBlogDirDict(c)
}

func (row *BlogPostTagRow) BindPostTag(c *db.Core) {
	sql, args := sqlbuilder.NewInsertBuilder().InsertInto("blog_post_tags").Cols("post_id", "tag_id").Values(row.PostID, row.TagID).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()
	srv.IsPanic(err)
	if ra <= 0 {panic(db.InsertErr)}
}

func (row *BlogPostTagRow) UnbindPostTag(c *db.Core) {
	sb := sqlbuilder.NewDeleteBuilder().DeleteFrom("blog_post_tags")
	sql, args := sb.Where(sb.And(sb.E("post_id", row.PostID), sb.E("tag_id", row.TagID))).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()
	srv.IsPanic(err)
	if ra <= 0 {panic(db.RemoveErr)}
}

func (form *BlogPostSaveForm) RefreshAndGetPath(c *db.Core) string {
	updateSB := sqlbuilder.NewUpdateBuilder().Update("blog_post")
	updateSB.Set(updateSB.Assign("date", time.Now()))
	sql, args := updateSB.Where(updateSB.E("id", form.ID)).Build()
	result := c.MustExec(sql, args...)
	ra, err := result.RowsAffected()
	
	srv.IsPanic(err)
	if ra <= 0 {panic(db.UpdateErr)}

	var res string
	selectSB := sqlbuilder.NewSelectBuilder().Select("path").From("blog_post")
	sql, args = selectSB.Where(selectSB.E("id", form.ID)).Build()
	srv.IsPanic(c.DB.QueryRowxContext(*c.Ctx, sql, args...).Scan(&res))
	return res
}
