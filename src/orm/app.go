package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type app interface {
	inkstone.ORM[model.App]
	GetWithTx(int, *sqlx.Tx) (*model.App, error)
}

type appImpl inkstone.AppContext

// Update implements app.
func (a *appImpl) Update(app *model.App) error {
	return namedExec(a.DB, sql.App.Update(), app, nil)
}

// UpdateWithTx implements app.
func (*appImpl) UpdateWithTx(app *model.App, tx *sqlx.Tx) error {
	return namedExec(tx, sql.App.Update(), app, nil)
}

// Insert implements app.
func (a *appImpl) Insert(app *model.App) error {
	return namedExec(a.DB, sql.App.Insert(), app, handleInsertResult)
}

// InsertWithTx implements app.
func (*appImpl) InsertWithTx(app *model.App, tx *sqlx.Tx) error {
	return namedExec(tx, sql.App.Insert(), app, handleInsertResult)
}

// GetWithTx implements app.
func (*appImpl) GetWithTx(id int, tx *sqlx.Tx) (app *model.App, err error) {
	app = new(model.App)
	err = tx.Get(
		app,
		sql.App.GetForUpdate(),
		id,
	)
	return
}

// Delete implements app.
func (*appImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements app.
func (a *appImpl) Find() (apps []model.App, err error) {
	err = a.DB.Select(
		&apps,
		sql.App.Find(),
	)
	return
}

// Get implements app.
func (a *appImpl) Get(id int) (app *model.App, err error) {
	app = new(model.App)
	err = a.DB.Get(
		app,
		sql.App.Get(),
		id,
	)
	return
}

// Save implements app.
func (a *appImpl) Save(app *model.App) error {
	return namedExec(a.DB, sql.App.Save(), app, handleSaveResult)
}

// SaveWithTx implements app.
func (*appImpl) SaveWithTx(app *model.App, tx *sqlx.Tx) error {
	return namedExec(tx, sql.App.Save(), app, handleSaveResult)
}

var _ app = (*appImpl)(nil)

func App(appCtx *inkstone.AppContext) app {
	return (*appImpl)(appCtx)
}
