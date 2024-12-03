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

func newEbSystemGroupDatum(db *gorm.DB, opts ...gen.DOOption) ebSystemGroupDatum {
	_ebSystemGroupDatum := ebSystemGroupDatum{}

	_ebSystemGroupDatum.ebSystemGroupDatumDo.UseDB(db, opts...)
	_ebSystemGroupDatum.ebSystemGroupDatumDo.UseModel(&model.EbSystemGroupDatum{})

	tableName := _ebSystemGroupDatum.ebSystemGroupDatumDo.TableName()
	_ebSystemGroupDatum.ALL = field.NewAsterisk(tableName)
	_ebSystemGroupDatum.ID = field.NewInt32(tableName, "id")
	_ebSystemGroupDatum.Gid = field.NewInt32(tableName, "gid")
	_ebSystemGroupDatum.Value = field.NewString(tableName, "value")
	_ebSystemGroupDatum.Sort = field.NewInt32(tableName, "sort")
	_ebSystemGroupDatum.Status = field.NewBool(tableName, "status")
	_ebSystemGroupDatum.CreateTime = field.NewTime(tableName, "create_time")
	_ebSystemGroupDatum.UpdateTime = field.NewTime(tableName, "update_time")

	_ebSystemGroupDatum.fillFieldMap()

	return _ebSystemGroupDatum
}

// ebSystemGroupDatum 组合数据详情表
type ebSystemGroupDatum struct {
	ebSystemGroupDatumDo ebSystemGroupDatumDo

	ALL        field.Asterisk
	ID         field.Int32  // 组合数据详情ID
	Gid        field.Int32  // 对应的数据组id
	Value      field.String // 数据组对应的数据值（json数据）
	Sort       field.Int32  // 数据排序
	Status     field.Bool   // 状态（1：开启；2：关闭；）
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebSystemGroupDatum) Table(newTableName string) *ebSystemGroupDatum {
	e.ebSystemGroupDatumDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebSystemGroupDatum) As(alias string) *ebSystemGroupDatum {
	e.ebSystemGroupDatumDo.DO = *(e.ebSystemGroupDatumDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebSystemGroupDatum) updateTableName(table string) *ebSystemGroupDatum {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Gid = field.NewInt32(table, "gid")
	e.Value = field.NewString(table, "value")
	e.Sort = field.NewInt32(table, "sort")
	e.Status = field.NewBool(table, "status")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebSystemGroupDatum) WithContext(ctx context.Context) IEbSystemGroupDatumDo {
	return e.ebSystemGroupDatumDo.WithContext(ctx)
}

func (e ebSystemGroupDatum) TableName() string { return e.ebSystemGroupDatumDo.TableName() }

func (e ebSystemGroupDatum) Alias() string { return e.ebSystemGroupDatumDo.Alias() }

func (e ebSystemGroupDatum) Columns(cols ...field.Expr) gen.Columns {
	return e.ebSystemGroupDatumDo.Columns(cols...)
}

func (e *ebSystemGroupDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebSystemGroupDatum) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 7)
	e.fieldMap["id"] = e.ID
	e.fieldMap["gid"] = e.Gid
	e.fieldMap["value"] = e.Value
	e.fieldMap["sort"] = e.Sort
	e.fieldMap["status"] = e.Status
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebSystemGroupDatum) clone(db *gorm.DB) ebSystemGroupDatum {
	e.ebSystemGroupDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebSystemGroupDatum) replaceDB(db *gorm.DB) ebSystemGroupDatum {
	e.ebSystemGroupDatumDo.ReplaceDB(db)
	return e
}

type ebSystemGroupDatumDo struct{ gen.DO }

type IEbSystemGroupDatumDo interface {
	gen.SubQuery
	Debug() IEbSystemGroupDatumDo
	WithContext(ctx context.Context) IEbSystemGroupDatumDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbSystemGroupDatumDo
	WriteDB() IEbSystemGroupDatumDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbSystemGroupDatumDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbSystemGroupDatumDo
	Not(conds ...gen.Condition) IEbSystemGroupDatumDo
	Or(conds ...gen.Condition) IEbSystemGroupDatumDo
	Select(conds ...field.Expr) IEbSystemGroupDatumDo
	Where(conds ...gen.Condition) IEbSystemGroupDatumDo
	Order(conds ...field.Expr) IEbSystemGroupDatumDo
	Distinct(cols ...field.Expr) IEbSystemGroupDatumDo
	Omit(cols ...field.Expr) IEbSystemGroupDatumDo
	Join(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo
	Group(cols ...field.Expr) IEbSystemGroupDatumDo
	Having(conds ...gen.Condition) IEbSystemGroupDatumDo
	Limit(limit int) IEbSystemGroupDatumDo
	Offset(offset int) IEbSystemGroupDatumDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemGroupDatumDo
	Unscoped() IEbSystemGroupDatumDo
	Create(values ...*model.EbSystemGroupDatum) error
	CreateInBatches(values []*model.EbSystemGroupDatum, batchSize int) error
	Save(values ...*model.EbSystemGroupDatum) error
	First() (*model.EbSystemGroupDatum, error)
	Take() (*model.EbSystemGroupDatum, error)
	Last() (*model.EbSystemGroupDatum, error)
	Find() ([]*model.EbSystemGroupDatum, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemGroupDatum, err error)
	FindInBatches(result *[]*model.EbSystemGroupDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbSystemGroupDatum) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbSystemGroupDatumDo
	Assign(attrs ...field.AssignExpr) IEbSystemGroupDatumDo
	Joins(fields ...field.RelationField) IEbSystemGroupDatumDo
	Preload(fields ...field.RelationField) IEbSystemGroupDatumDo
	FirstOrInit() (*model.EbSystemGroupDatum, error)
	FirstOrCreate() (*model.EbSystemGroupDatum, error)
	FindByPage(offset int, limit int) (result []*model.EbSystemGroupDatum, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbSystemGroupDatumDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebSystemGroupDatumDo) Debug() IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Debug())
}

func (e ebSystemGroupDatumDo) WithContext(ctx context.Context) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebSystemGroupDatumDo) ReadDB() IEbSystemGroupDatumDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebSystemGroupDatumDo) WriteDB() IEbSystemGroupDatumDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebSystemGroupDatumDo) Session(config *gorm.Session) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebSystemGroupDatumDo) Clauses(conds ...clause.Expression) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebSystemGroupDatumDo) Returning(value interface{}, columns ...string) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebSystemGroupDatumDo) Not(conds ...gen.Condition) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebSystemGroupDatumDo) Or(conds ...gen.Condition) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebSystemGroupDatumDo) Select(conds ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebSystemGroupDatumDo) Where(conds ...gen.Condition) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebSystemGroupDatumDo) Order(conds ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebSystemGroupDatumDo) Distinct(cols ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebSystemGroupDatumDo) Omit(cols ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebSystemGroupDatumDo) Join(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebSystemGroupDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebSystemGroupDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebSystemGroupDatumDo) Group(cols ...field.Expr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebSystemGroupDatumDo) Having(conds ...gen.Condition) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebSystemGroupDatumDo) Limit(limit int) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebSystemGroupDatumDo) Offset(offset int) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebSystemGroupDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebSystemGroupDatumDo) Unscoped() IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebSystemGroupDatumDo) Create(values ...*model.EbSystemGroupDatum) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebSystemGroupDatumDo) CreateInBatches(values []*model.EbSystemGroupDatum, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebSystemGroupDatumDo) Save(values ...*model.EbSystemGroupDatum) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebSystemGroupDatumDo) First() (*model.EbSystemGroupDatum, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemGroupDatum), nil
	}
}

func (e ebSystemGroupDatumDo) Take() (*model.EbSystemGroupDatum, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemGroupDatum), nil
	}
}

func (e ebSystemGroupDatumDo) Last() (*model.EbSystemGroupDatum, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemGroupDatum), nil
	}
}

func (e ebSystemGroupDatumDo) Find() ([]*model.EbSystemGroupDatum, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbSystemGroupDatum), err
}

func (e ebSystemGroupDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemGroupDatum, err error) {
	buf := make([]*model.EbSystemGroupDatum, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebSystemGroupDatumDo) FindInBatches(result *[]*model.EbSystemGroupDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebSystemGroupDatumDo) Attrs(attrs ...field.AssignExpr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebSystemGroupDatumDo) Assign(attrs ...field.AssignExpr) IEbSystemGroupDatumDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebSystemGroupDatumDo) Joins(fields ...field.RelationField) IEbSystemGroupDatumDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebSystemGroupDatumDo) Preload(fields ...field.RelationField) IEbSystemGroupDatumDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebSystemGroupDatumDo) FirstOrInit() (*model.EbSystemGroupDatum, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemGroupDatum), nil
	}
}

func (e ebSystemGroupDatumDo) FirstOrCreate() (*model.EbSystemGroupDatum, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemGroupDatum), nil
	}
}

func (e ebSystemGroupDatumDo) FindByPage(offset int, limit int) (result []*model.EbSystemGroupDatum, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e ebSystemGroupDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebSystemGroupDatumDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebSystemGroupDatumDo) Delete(models ...*model.EbSystemGroupDatum) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebSystemGroupDatumDo) withDO(do gen.Dao) *ebSystemGroupDatumDo {
	e.DO = *do.(*gen.DO)
	return e
}
