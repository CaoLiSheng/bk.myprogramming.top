package blogCfg

type BlogCfgRow struct {
	ID  string `db:"id" json:"id" form:"id" binding:"required"`
	Cfg string `db:"cfg" json:"cfg" form:"cfg" binding:"required"`
}

type BlogCfgDictIDs []string

type BlogCfgDict map[string]string

type BlogCfgRowId struct {
	ID string `form:"id" binding:"required"`
}
