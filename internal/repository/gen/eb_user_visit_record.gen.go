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

func newEbUserVisitRecord(db *gorm.DB, opts ...gen.DOOption) ebUserVisitRecord {
	_ebUserVisitRecord := ebUserVisitRecord{}

	_ebUserVisitRecord.ebUserVisitRecordDo.UseDB(db, opts...)
	_ebUserVisitRecord.ebUserVisitRecordDo.UseModel(&model.EbUserVisitRecord{})

	tableName := _ebUserVisitRecord.ebUserVisitRecordDo.TableName()
	_ebUserVisitRecord.ALL = field.NewAsterisk(tableName)
	_ebUserVisitRecord.ID = field.NewInt32(tableName, "id")
	_ebUserVisitRecord.Date = field.NewString(tableName, "date")
	_ebUserVisitRecord.UID = field.NewInt32(tableName, "uid")
	_ebUserVisitRecord.VisitType = field.NewInt32(tableName, "visit_type")

	_ebUserVisitRecord.fillFieldMap()

	return _ebUserVisitRecord
}

// ebUserVisitRecord 用户访问记录表
type ebUserVisitRecord struct {
	ebUserVisitRecordDo ebUserVisitRecordDo

	ALL       field.Asterisk
	ID        field.Int32
	Date      field.String // 日期
	UID       field.Int32  // 用户uid
	VisitType field.Int32  // 访问类型

	fieldMap map[string]field.Expr
}

func (e ebUserVisitRecord) Table(newTableName string) *ebUserVisitRecord {
	e.ebUserVisitRecordDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebUserVisitRecord) As(alias string) *ebUserVisitRecord {
	e.ebUserVisitRecordDo.DO = *(e.ebUserVisitRecordDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebUserVisitRecord) updateTableName(table string) *ebUserVisitRecord {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Date = field.NewString(table, "date")
	e.UID = field.NewInt32(table, "uid")
	e.VisitType = field.NewInt32(table, "visit_type")

	e.fillFieldMap()

	return e
}

func (e *ebUserVisitRecord) WithContext(ctx context.Context) IEbUserVisitRecordDo {
	return e.ebUserVisitRecordDo.WithContext(ctx)
}

func (e ebUserVisitRecord) TableName() string { return e.ebUserVisitRecordDo.TableName() }

func (e ebUserVisitRecord) Alias() string { return e.ebUserVisitRecordDo.Alias() }

func (e ebUserVisitRecord) Columns(cols ...field.Expr) gen.Columns {
	return e.ebUserVisitRecordDo.Columns(cols...)
}

func (e *ebUserVisitRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebUserVisitRecord) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 4)
	e.fieldMap["id"] = e.ID
	e.fieldMap["date"] = e.Date
	e.fieldMap["uid"] = e.UID
	e.fieldMap["visit_type"] = e.VisitType
}

func (e ebUserVisitRecord) clone(db *gorm.DB) ebUserVisitRecord {
	e.ebUserVisitRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebUserVisitRecord) replaceDB(db *gorm.DB) ebUserVisitRecord {
	e.ebUserVisitRecordDo.ReplaceDB(db)
	return e
}

type ebUserVisitRecordDo struct{ gen.DO }

type IEbUserVisitRecordDo interface {
	gen.SubQuery
	Debug() IEbUserVisitRecordDo
	WithContext(ctx context.Context) IEbUserVisitRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbUserVisitRecordDo
	WriteDB() IEbUserVisitRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbUserVisitRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbUserVisitRecordDo
	Not(conds ...gen.Condition) IEbUserVisitRecordDo
	Or(conds ...gen.Condition) IEbUserVisitRecordDo
	Select(conds ...field.Expr) IEbUserVisitRecordDo
	Where(conds ...gen.Condition) IEbUserVisitRecordDo
	Order(conds ...field.Expr) IEbUserVisitRecordDo
	Distinct(cols ...field.Expr) IEbUserVisitRecordDo
	Omit(cols ...field.Expr) IEbUserVisitRecordDo
	Join(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo
	Group(cols ...field.Expr) IEbUserVisitRecordDo
	Having(conds ...gen.Condition) IEbUserVisitRecordDo
	Limit(limit int) IEbUserVisitRecordDo
	Offset(offset int) IEbUserVisitRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserVisitRecordDo
	Unscoped() IEbUserVisitRecordDo
	Create(values ...*model.EbUserVisitRecord) error
	CreateInBatches(values []*model.EbUserVisitRecord, batchSize int) error
	Save(values ...*model.EbUserVisitRecord) error
	First() (*model.EbUserVisitRecord, error)
	Take() (*model.EbUserVisitRecord, error)
	Last() (*model.EbUserVisitRecord, error)
	Find() ([]*model.EbUserVisitRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserVisitRecord, err error)
	FindInBatches(result *[]*model.EbUserVisitRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbUserVisitRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbUserVisitRecordDo
	Assign(attrs ...field.AssignExpr) IEbUserVisitRecordDo
	Joins(fields ...field.RelationField) IEbUserVisitRecordDo
	Preload(fields ...field.RelationField) IEbUserVisitRecordDo
	FirstOrInit() (*model.EbUserVisitRecord, error)
	FirstOrCreate() (*model.EbUserVisitRecord, error)
	FindByPage(offset int, limit int) (result []*model.EbUserVisitRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbUserVisitRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebUserVisitRecordDo) Debug() IEbUserVisitRecordDo {
	return e.withDO(e.DO.Debug())
}

func (e ebUserVisitRecordDo) WithContext(ctx context.Context) IEbUserVisitRecordDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebUserVisitRecordDo) ReadDB() IEbUserVisitRecordDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebUserVisitRecordDo) WriteDB() IEbUserVisitRecordDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebUserVisitRecordDo) Session(config *gorm.Session) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebUserVisitRecordDo) Clauses(conds ...clause.Expression) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebUserVisitRecordDo) Returning(value interface{}, columns ...string) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebUserVisitRecordDo) Not(conds ...gen.Condition) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebUserVisitRecordDo) Or(conds ...gen.Condition) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebUserVisitRecordDo) Select(conds ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebUserVisitRecordDo) Where(conds ...gen.Condition) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebUserVisitRecordDo) Order(conds ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebUserVisitRecordDo) Distinct(cols ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebUserVisitRecordDo) Omit(cols ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebUserVisitRecordDo) Join(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebUserVisitRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebUserVisitRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebUserVisitRecordDo) Group(cols ...field.Expr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebUserVisitRecordDo) Having(conds ...gen.Condition) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebUserVisitRecordDo) Limit(limit int) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebUserVisitRecordDo) Offset(offset int) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebUserVisitRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebUserVisitRecordDo) Unscoped() IEbUserVisitRecordDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebUserVisitRecordDo) Create(values ...*model.EbUserVisitRecord) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebUserVisitRecordDo) CreateInBatches(values []*model.EbUserVisitRecord, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebUserVisitRecordDo) Save(values ...*model.EbUserVisitRecord) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebUserVisitRecordDo) First() (*model.EbUserVisitRecord, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserVisitRecord), nil
	}
}

func (e ebUserVisitRecordDo) Take() (*model.EbUserVisitRecord, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserVisitRecord), nil
	}
}

func (e ebUserVisitRecordDo) Last() (*model.EbUserVisitRecord, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserVisitRecord), nil
	}
}

func (e ebUserVisitRecordDo) Find() ([]*model.EbUserVisitRecord, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbUserVisitRecord), err
}

func (e ebUserVisitRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbUserVisitRecord, err error) {
	buf := make([]*model.EbUserVisitRecord, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebUserVisitRecordDo) FindInBatches(result *[]*model.EbUserVisitRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebUserVisitRecordDo) Attrs(attrs ...field.AssignExpr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebUserVisitRecordDo) Assign(attrs ...field.AssignExpr) IEbUserVisitRecordDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebUserVisitRecordDo) Joins(fields ...field.RelationField) IEbUserVisitRecordDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebUserVisitRecordDo) Preload(fields ...field.RelationField) IEbUserVisitRecordDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebUserVisitRecordDo) FirstOrInit() (*model.EbUserVisitRecord, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserVisitRecord), nil
	}
}

func (e ebUserVisitRecordDo) FirstOrCreate() (*model.EbUserVisitRecord, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbUserVisitRecord), nil
	}
}

func (e ebUserVisitRecordDo) FindByPage(offset int, limit int) (result []*model.EbUserVisitRecord, count int64, err error) {
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

func (e ebUserVisitRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebUserVisitRecordDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebUserVisitRecordDo) Delete(models ...*model.EbUserVisitRecord) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebUserVisitRecordDo) withDO(do gen.Dao) *ebUserVisitRecordDo {
	e.DO = *do.(*gen.DO)
	return e
}
