package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type group interface {
	inkstone.ORM[model.Group]
	CountWithTx(gtype, appId int, tx *sqlx.Tx) (int, error)
	PaginationWithTx(gtype, appId, offset, limit int, tx *sqlx.Tx) ([]model.GroupWithApp, error)
}

type groupImpl inkstone.AppContext

// CountWithTx implements group.
func (*groupImpl) CountWithTx(gtype, appId int, tx *sqlx.Tx) (c int, err error) {
	err = tx.Get(&c, sql.Group.Count(), gtype, appId)
	return
}

// PaginationWithTx implements group.
func (*groupImpl) PaginationWithTx(gtype, appId, offset, limit int, tx *sqlx.Tx) (groups []model.GroupWithApp, err error) {
	err = tx.Select(
		&groups,
		sql.Group.Pagination(),
		gtype,
		appId,
		limit,
		offset,
	)
	return
}

// Delete implements group.
func (*groupImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements group.
func (*groupImpl) Find() ([]model.Group, error) {
	panic("unimplemented")
}

// Get implements group.
func (*groupImpl) Get(int) (*model.Group, error) {
	panic("unimplemented")
}

// Save implements group.
func (*groupImpl) Save(*model.Group) error {
	panic("unimplemented")
}

// SaveWithTx implements group.
func (*groupImpl) SaveWithTx(*model.Group, *sqlx.Tx) error {
	panic("unimplemented")
}

var _ group = (*groupImpl)(nil)

func Group(appCtx *inkstone.AppContext) group {
	return (*groupImpl)(appCtx)
}
