// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"crmeb_go/internal/model"
)

func newWechatExceptions(db *gorm.DB, opts ...gen.DOOption) wechatExceptions {
	_wechatExceptions := wechatExceptions{}

	_wechatExceptions.wechatExceptionsDo.UseDB(db, opts...)
	_wechatExceptions.wechatExceptionsDo.UseModel(&model.WechatExceptions{})

	tableName := _wechatExceptions.wechatExceptionsDo.TableName()
	_wechatExceptions.ALL = field.NewAsterisk(tableName)
	_wechatExceptions.ID = field.NewInt64(tableName, "id")
	_wechatExceptions.Errcode = field.NewString(tableName, "errcode")
	_wechatExceptions.Errmsg = field.NewString(tableName, "errmsg")
	_wechatExceptions.Data = field.NewString(tableName, "data")
	_wechatExceptions.Remark = field.NewString(tableName, "remark")
	_wechatExceptions.CreatedAt = field.NewInt64(tableName, "created_at")
	_wechatExceptions.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_wechatExceptions.DeletedAt = field.NewField(tableName, "deleted_at")

	_wechatExceptions.fillFieldMap()

	return _wechatExceptions
}

// wechatExceptions 微信异常表
type wechatExceptions struct {
	wechatExceptionsDo wechatExceptionsDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	Errcode   field.String // 错误码
	Errmsg    field.String // 错误信息
	Data      field.String // 回复数据
	Remark    field.String // 备注
	CreatedAt field.Int64  // 创建时间
	UpdatedAt field.Int64  // 修改时间
	DeletedAt field.Field  // 是否删除

	fieldMap map[string]field.Expr
}

func (w wechatExceptions) Table(newTableName string) *wechatExceptions {
	w.wechatExceptionsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wechatExceptions) As(alias string) *wechatExceptions {
	w.wechatExceptionsDo.DO = *(w.wechatExceptionsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wechatExceptions) updateTableName(table string) *wechatExceptions {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Errcode = field.NewString(table, "errcode")
	w.Errmsg = field.NewString(table, "errmsg")
	w.Data = field.NewString(table, "data")
	w.Remark = field.NewString(table, "remark")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")

	w.fillFieldMap()

	return w
}

func (w *wechatExceptions) WithContext(ctx context.Context) IWechatExceptionsDo {
	return w.wechatExceptionsDo.WithContext(ctx)
}

func (w wechatExceptions) TableName() string { return w.wechatExceptionsDo.TableName() }

func (w wechatExceptions) Alias() string { return w.wechatExceptionsDo.Alias() }

func (w wechatExceptions) Columns(cols ...field.Expr) gen.Columns {
	return w.wechatExceptionsDo.Columns(cols...)
}

func (w *wechatExceptions) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wechatExceptions) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["errcode"] = w.Errcode
	w.fieldMap["errmsg"] = w.Errmsg
	w.fieldMap["data"] = w.Data
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
}

func (w wechatExceptions) clone(db *gorm.DB) wechatExceptions {
	w.wechatExceptionsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wechatExceptions) replaceDB(db *gorm.DB) wechatExceptions {
	w.wechatExceptionsDo.ReplaceDB(db)
	return w
}

type wechatExceptionsDo struct{ gen.DO }

type IWechatExceptionsDo interface {
	gen.SubQuery
	Debug() IWechatExceptionsDo
	WithContext(ctx context.Context) IWechatExceptionsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWechatExceptionsDo
	WriteDB() IWechatExceptionsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWechatExceptionsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWechatExceptionsDo
	Not(conds ...gen.Condition) IWechatExceptionsDo
	Or(conds ...gen.Condition) IWechatExceptionsDo
	Select(conds ...field.Expr) IWechatExceptionsDo
	Where(conds ...gen.Condition) IWechatExceptionsDo
	Order(conds ...field.Expr) IWechatExceptionsDo
	Distinct(cols ...field.Expr) IWechatExceptionsDo
	Omit(cols ...field.Expr) IWechatExceptionsDo
	Join(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo
	Group(cols ...field.Expr) IWechatExceptionsDo
	Having(conds ...gen.Condition) IWechatExceptionsDo
	Limit(limit int) IWechatExceptionsDo
	Offset(offset int) IWechatExceptionsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatExceptionsDo
	Unscoped() IWechatExceptionsDo
	Create(values ...*model.WechatExceptions) error
	CreateInBatches(values []*model.WechatExceptions, batchSize int) error
	Save(values ...*model.WechatExceptions) error
	First() (*model.WechatExceptions, error)
	Take() (*model.WechatExceptions, error)
	Last() (*model.WechatExceptions, error)
	Find() ([]*model.WechatExceptions, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatExceptions, err error)
	FindInBatches(result *[]*model.WechatExceptions, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WechatExceptions) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWechatExceptionsDo
	Assign(attrs ...field.AssignExpr) IWechatExceptionsDo
	Joins(fields ...field.RelationField) IWechatExceptionsDo
	Preload(fields ...field.RelationField) IWechatExceptionsDo
	FirstOrInit() (*model.WechatExceptions, error)
	FirstOrCreate() (*model.WechatExceptions, error)
	FindByPage(offset int, limit int) (result []*model.WechatExceptions, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWechatExceptionsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wechatExceptionsDo) Debug() IWechatExceptionsDo {
	return w.withDO(w.DO.Debug())
}

func (w wechatExceptionsDo) WithContext(ctx context.Context) IWechatExceptionsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wechatExceptionsDo) ReadDB() IWechatExceptionsDo {
	return w.Clauses(dbresolver.Read)
}

func (w wechatExceptionsDo) WriteDB() IWechatExceptionsDo {
	return w.Clauses(dbresolver.Write)
}

func (w wechatExceptionsDo) Session(config *gorm.Session) IWechatExceptionsDo {
	return w.withDO(w.DO.Session(config))
}

func (w wechatExceptionsDo) Clauses(conds ...clause.Expression) IWechatExceptionsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wechatExceptionsDo) Returning(value interface{}, columns ...string) IWechatExceptionsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wechatExceptionsDo) Not(conds ...gen.Condition) IWechatExceptionsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wechatExceptionsDo) Or(conds ...gen.Condition) IWechatExceptionsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wechatExceptionsDo) Select(conds ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wechatExceptionsDo) Where(conds ...gen.Condition) IWechatExceptionsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wechatExceptionsDo) Order(conds ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wechatExceptionsDo) Distinct(cols ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wechatExceptionsDo) Omit(cols ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wechatExceptionsDo) Join(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wechatExceptionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wechatExceptionsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wechatExceptionsDo) Group(cols ...field.Expr) IWechatExceptionsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wechatExceptionsDo) Having(conds ...gen.Condition) IWechatExceptionsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wechatExceptionsDo) Limit(limit int) IWechatExceptionsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wechatExceptionsDo) Offset(offset int) IWechatExceptionsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wechatExceptionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWechatExceptionsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wechatExceptionsDo) Unscoped() IWechatExceptionsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wechatExceptionsDo) Create(values ...*model.WechatExceptions) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wechatExceptionsDo) CreateInBatches(values []*model.WechatExceptions, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wechatExceptionsDo) Save(values ...*model.WechatExceptions) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wechatExceptionsDo) First() (*model.WechatExceptions, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatExceptions), nil
	}
}

func (w wechatExceptionsDo) Take() (*model.WechatExceptions, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatExceptions), nil
	}
}

func (w wechatExceptionsDo) Last() (*model.WechatExceptions, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatExceptions), nil
	}
}

func (w wechatExceptionsDo) Find() ([]*model.WechatExceptions, error) {
	result, err := w.DO.Find()
	return result.([]*model.WechatExceptions), err
}

func (w wechatExceptionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WechatExceptions, err error) {
	buf := make([]*model.WechatExceptions, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wechatExceptionsDo) FindInBatches(result *[]*model.WechatExceptions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wechatExceptionsDo) Attrs(attrs ...field.AssignExpr) IWechatExceptionsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wechatExceptionsDo) Assign(attrs ...field.AssignExpr) IWechatExceptionsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wechatExceptionsDo) Joins(fields ...field.RelationField) IWechatExceptionsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wechatExceptionsDo) Preload(fields ...field.RelationField) IWechatExceptionsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wechatExceptionsDo) FirstOrInit() (*model.WechatExceptions, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatExceptions), nil
	}
}

func (w wechatExceptionsDo) FirstOrCreate() (*model.WechatExceptions, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WechatExceptions), nil
	}
}

func (w wechatExceptionsDo) FindByPage(offset int, limit int) (result []*model.WechatExceptions, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wechatExceptionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wechatExceptionsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wechatExceptionsDo) Delete(models ...*model.WechatExceptions) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wechatExceptionsDo) withDO(do gen.Dao) *wechatExceptionsDo {
	w.DO = *do.(*gen.DO)
	return w
}
