package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type dept interface {
	orm.Inserter[models.Department]
}

type deptImpl app.AppContext

// Insert implements dept.
func (d *deptImpl) Insert(dept *models.Department) error {
	return orm.NamedInsert(d.DB, sqls.Dept.Insert(), dept)
}

// InsertWithTx implements dept.
func (d *deptImpl) InsertTx(tx *sqlx.Tx, dept *models.Department) error {
	return orm.NamedInsert(tx, sqls.Dept.Insert(), dept)
}

var _ dept = (*deptImpl)(nil)

func Dept(appCtx *app.AppContext) dept {
	return (*deptImpl)(appCtx)
}
