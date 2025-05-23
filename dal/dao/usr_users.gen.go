// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/faiz/llm-code-review/dal/model"
)

func newUsrUser(db *gorm.DB, opts ...gen.DOOption) usrUser {
	_usrUser := usrUser{}

	_usrUser.usrUserDo.UseDB(db, opts...)
	_usrUser.usrUserDo.UseModel(&model.UsrUser{})

	tableName := _usrUser.usrUserDo.TableName()
	_usrUser.ALL = field.NewAsterisk(tableName)
	_usrUser.ID = field.NewInt64(tableName, "id")
	_usrUser.Username = field.NewString(tableName, "username")
	_usrUser.Token = field.NewString(tableName, "token")
	_usrUser.Email = field.NewString(tableName, "email")
	_usrUser.AesKey = field.NewString(tableName, "aes_key")
	_usrUser.GmtCreate = field.NewTime(tableName, "gmt_create")
	_usrUser.GmtUpdate = field.NewTime(tableName, "gmt_update")

	_usrUser.fillFieldMap()

	return _usrUser
}

type usrUser struct {
	usrUserDo usrUserDo

	ALL       field.Asterisk
	ID        field.Int64
	Username  field.String // 用户名
	Token     field.String // github token （加密存储）
	Email     field.String // 邮箱
	AesKey    field.String // 对称加密密钥
	GmtCreate field.Time
	GmtUpdate field.Time

	fieldMap map[string]field.Expr
}

func (u usrUser) Table(newTableName string) *usrUser {
	u.usrUserDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u usrUser) As(alias string) *usrUser {
	u.usrUserDo.DO = *(u.usrUserDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *usrUser) updateTableName(table string) *usrUser {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.Username = field.NewString(table, "username")
	u.Token = field.NewString(table, "token")
	u.Email = field.NewString(table, "email")
	u.AesKey = field.NewString(table, "aes_key")
	u.GmtCreate = field.NewTime(table, "gmt_create")
	u.GmtUpdate = field.NewTime(table, "gmt_update")

	u.fillFieldMap()

	return u
}

func (u *usrUser) WithContext(ctx context.Context) *usrUserDo { return u.usrUserDo.WithContext(ctx) }

func (u usrUser) TableName() string { return u.usrUserDo.TableName() }

func (u usrUser) Alias() string { return u.usrUserDo.Alias() }

func (u usrUser) Columns(cols ...field.Expr) gen.Columns { return u.usrUserDo.Columns(cols...) }

func (u *usrUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *usrUser) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["username"] = u.Username
	u.fieldMap["token"] = u.Token
	u.fieldMap["email"] = u.Email
	u.fieldMap["aes_key"] = u.AesKey
	u.fieldMap["gmt_create"] = u.GmtCreate
	u.fieldMap["gmt_update"] = u.GmtUpdate
}

func (u usrUser) clone(db *gorm.DB) usrUser {
	u.usrUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u usrUser) replaceDB(db *gorm.DB) usrUser {
	u.usrUserDo.ReplaceDB(db)
	return u
}

type usrUserDo struct{ gen.DO }

func (u usrUserDo) Debug() *usrUserDo {
	return u.withDO(u.DO.Debug())
}

func (u usrUserDo) WithContext(ctx context.Context) *usrUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u usrUserDo) ReadDB() *usrUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u usrUserDo) WriteDB() *usrUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u usrUserDo) Session(config *gorm.Session) *usrUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u usrUserDo) Clauses(conds ...clause.Expression) *usrUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u usrUserDo) Returning(value interface{}, columns ...string) *usrUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u usrUserDo) Not(conds ...gen.Condition) *usrUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u usrUserDo) Or(conds ...gen.Condition) *usrUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u usrUserDo) Select(conds ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u usrUserDo) Where(conds ...gen.Condition) *usrUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u usrUserDo) Order(conds ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u usrUserDo) Distinct(cols ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u usrUserDo) Omit(cols ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u usrUserDo) Join(table schema.Tabler, on ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u usrUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u usrUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u usrUserDo) Group(cols ...field.Expr) *usrUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u usrUserDo) Having(conds ...gen.Condition) *usrUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u usrUserDo) Limit(limit int) *usrUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u usrUserDo) Offset(offset int) *usrUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u usrUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *usrUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u usrUserDo) Unscoped() *usrUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u usrUserDo) Create(values ...*model.UsrUser) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u usrUserDo) CreateInBatches(values []*model.UsrUser, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u usrUserDo) Save(values ...*model.UsrUser) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u usrUserDo) First() (*model.UsrUser, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsrUser), nil
	}
}

func (u usrUserDo) Take() (*model.UsrUser, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsrUser), nil
	}
}

func (u usrUserDo) Last() (*model.UsrUser, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsrUser), nil
	}
}

func (u usrUserDo) Find() ([]*model.UsrUser, error) {
	result, err := u.DO.Find()
	return result.([]*model.UsrUser), err
}

func (u usrUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UsrUser, err error) {
	buf := make([]*model.UsrUser, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u usrUserDo) FindInBatches(result *[]*model.UsrUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u usrUserDo) Attrs(attrs ...field.AssignExpr) *usrUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u usrUserDo) Assign(attrs ...field.AssignExpr) *usrUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u usrUserDo) Joins(fields ...field.RelationField) *usrUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u usrUserDo) Preload(fields ...field.RelationField) *usrUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u usrUserDo) FirstOrInit() (*model.UsrUser, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsrUser), nil
	}
}

func (u usrUserDo) FirstOrCreate() (*model.UsrUser, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UsrUser), nil
	}
}

func (u usrUserDo) FindByPage(offset int, limit int) (result []*model.UsrUser, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u usrUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u usrUserDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u usrUserDo) Delete(models ...*model.UsrUser) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *usrUserDo) withDO(do gen.Dao) *usrUserDo {
	u.DO = *do.(*gen.DO)
	return u
}
