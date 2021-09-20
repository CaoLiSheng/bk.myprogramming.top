package blogCfg

type blogCfgRow struct {
	ID  string `db:"id" json:"id"`
	Cfg string `db:"cfg" json:"cfg"`
}
