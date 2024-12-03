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

func newEbSystemStore(db *gorm.DB, opts ...gen.DOOption) ebSystemStore {
	_ebSystemStore := ebSystemStore{}

	_ebSystemStore.ebSystemStoreDo.UseDB(db, opts...)
	_ebSystemStore.ebSystemStoreDo.UseModel(&model.EbSystemStore{})

	tableName := _ebSystemStore.ebSystemStoreDo.TableName()
	_ebSystemStore.ALL = field.NewAsterisk(tableName)
	_ebSystemStore.ID = field.NewInt32(tableName, "id")
	_ebSystemStore.Name = field.NewString(tableName, "name")
	_ebSystemStore.Introduction = field.NewString(tableName, "introduction")
	_ebSystemStore.Phone = field.NewString(tableName, "phone")
	_ebSystemStore.Address = field.NewString(tableName, "address")
	_ebSystemStore.DetailedAddress = field.NewString(tableName, "detailed_address")
	_ebSystemStore.Image = field.NewString(tableName, "image")
	_ebSystemStore.Latitude = field.NewString(tableName, "latitude")
	_ebSystemStore.Longitude = field.NewString(tableName, "longitude")
	_ebSystemStore.ValidTime = field.NewString(tableName, "valid_time")
	_ebSystemStore.DayTime = field.NewString(tableName, "day_time")
	_ebSystemStore.IsShow = field.NewBool(tableName, "is_show")
	_ebSystemStore.IsDel = field.NewBool(tableName, "is_del")
	_ebSystemStore.CreateTime = field.NewTime(tableName, "create_time")
	_ebSystemStore.UpdateTime = field.NewTime(tableName, "update_time")

	_ebSystemStore.fillFieldMap()

	return _ebSystemStore
}

// ebSystemStore 门店自提
type ebSystemStore struct {
	ebSystemStoreDo ebSystemStoreDo

	ALL             field.Asterisk
	ID              field.Int32
	Name            field.String // 门店名称
	Introduction    field.String // 简介
	Phone           field.String // 手机号码
	Address         field.String // 省市区
	DetailedAddress field.String // 详细地址
	Image           field.String // 门店logo
	Latitude        field.String // 纬度
	Longitude       field.String // 经度
	ValidTime       field.String // 核销有效日期
	DayTime         field.String // 每日营业开关时间
	IsShow          field.Bool   // 是否显示
	IsDel           field.Bool   // 是否删除
	CreateTime      field.Time   // 添加时间
	UpdateTime      field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (e ebSystemStore) Table(newTableName string) *ebSystemStore {
	e.ebSystemStoreDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ebSystemStore) As(alias string) *ebSystemStore {
	e.ebSystemStoreDo.DO = *(e.ebSystemStoreDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ebSystemStore) updateTableName(table string) *ebSystemStore {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt32(table, "id")
	e.Name = field.NewString(table, "name")
	e.Introduction = field.NewString(table, "introduction")
	e.Phone = field.NewString(table, "phone")
	e.Address = field.NewString(table, "address")
	e.DetailedAddress = field.NewString(table, "detailed_address")
	e.Image = field.NewString(table, "image")
	e.Latitude = field.NewString(table, "latitude")
	e.Longitude = field.NewString(table, "longitude")
	e.ValidTime = field.NewString(table, "valid_time")
	e.DayTime = field.NewString(table, "day_time")
	e.IsShow = field.NewBool(table, "is_show")
	e.IsDel = field.NewBool(table, "is_del")
	e.CreateTime = field.NewTime(table, "create_time")
	e.UpdateTime = field.NewTime(table, "update_time")

	e.fillFieldMap()

	return e
}

func (e *ebSystemStore) WithContext(ctx context.Context) IEbSystemStoreDo {
	return e.ebSystemStoreDo.WithContext(ctx)
}

func (e ebSystemStore) TableName() string { return e.ebSystemStoreDo.TableName() }

func (e ebSystemStore) Alias() string { return e.ebSystemStoreDo.Alias() }

func (e ebSystemStore) Columns(cols ...field.Expr) gen.Columns {
	return e.ebSystemStoreDo.Columns(cols...)
}

func (e *ebSystemStore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ebSystemStore) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 15)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["introduction"] = e.Introduction
	e.fieldMap["phone"] = e.Phone
	e.fieldMap["address"] = e.Address
	e.fieldMap["detailed_address"] = e.DetailedAddress
	e.fieldMap["image"] = e.Image
	e.fieldMap["latitude"] = e.Latitude
	e.fieldMap["longitude"] = e.Longitude
	e.fieldMap["valid_time"] = e.ValidTime
	e.fieldMap["day_time"] = e.DayTime
	e.fieldMap["is_show"] = e.IsShow
	e.fieldMap["is_del"] = e.IsDel
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["update_time"] = e.UpdateTime
}

func (e ebSystemStore) clone(db *gorm.DB) ebSystemStore {
	e.ebSystemStoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ebSystemStore) replaceDB(db *gorm.DB) ebSystemStore {
	e.ebSystemStoreDo.ReplaceDB(db)
	return e
}

type ebSystemStoreDo struct{ gen.DO }

type IEbSystemStoreDo interface {
	gen.SubQuery
	Debug() IEbSystemStoreDo
	WithContext(ctx context.Context) IEbSystemStoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEbSystemStoreDo
	WriteDB() IEbSystemStoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEbSystemStoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEbSystemStoreDo
	Not(conds ...gen.Condition) IEbSystemStoreDo
	Or(conds ...gen.Condition) IEbSystemStoreDo
	Select(conds ...field.Expr) IEbSystemStoreDo
	Where(conds ...gen.Condition) IEbSystemStoreDo
	Order(conds ...field.Expr) IEbSystemStoreDo
	Distinct(cols ...field.Expr) IEbSystemStoreDo
	Omit(cols ...field.Expr) IEbSystemStoreDo
	Join(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo
	Group(cols ...field.Expr) IEbSystemStoreDo
	Having(conds ...gen.Condition) IEbSystemStoreDo
	Limit(limit int) IEbSystemStoreDo
	Offset(offset int) IEbSystemStoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemStoreDo
	Unscoped() IEbSystemStoreDo
	Create(values ...*model.EbSystemStore) error
	CreateInBatches(values []*model.EbSystemStore, batchSize int) error
	Save(values ...*model.EbSystemStore) error
	First() (*model.EbSystemStore, error)
	Take() (*model.EbSystemStore, error)
	Last() (*model.EbSystemStore, error)
	Find() ([]*model.EbSystemStore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemStore, err error)
	FindInBatches(result *[]*model.EbSystemStore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EbSystemStore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEbSystemStoreDo
	Assign(attrs ...field.AssignExpr) IEbSystemStoreDo
	Joins(fields ...field.RelationField) IEbSystemStoreDo
	Preload(fields ...field.RelationField) IEbSystemStoreDo
	FirstOrInit() (*model.EbSystemStore, error)
	FirstOrCreate() (*model.EbSystemStore, error)
	FindByPage(offset int, limit int) (result []*model.EbSystemStore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEbSystemStoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ebSystemStoreDo) Debug() IEbSystemStoreDo {
	return e.withDO(e.DO.Debug())
}

func (e ebSystemStoreDo) WithContext(ctx context.Context) IEbSystemStoreDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ebSystemStoreDo) ReadDB() IEbSystemStoreDo {
	return e.Clauses(dbresolver.Read)
}

func (e ebSystemStoreDo) WriteDB() IEbSystemStoreDo {
	return e.Clauses(dbresolver.Write)
}

func (e ebSystemStoreDo) Session(config *gorm.Session) IEbSystemStoreDo {
	return e.withDO(e.DO.Session(config))
}

func (e ebSystemStoreDo) Clauses(conds ...clause.Expression) IEbSystemStoreDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ebSystemStoreDo) Returning(value interface{}, columns ...string) IEbSystemStoreDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ebSystemStoreDo) Not(conds ...gen.Condition) IEbSystemStoreDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ebSystemStoreDo) Or(conds ...gen.Condition) IEbSystemStoreDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ebSystemStoreDo) Select(conds ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ebSystemStoreDo) Where(conds ...gen.Condition) IEbSystemStoreDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ebSystemStoreDo) Order(conds ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ebSystemStoreDo) Distinct(cols ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ebSystemStoreDo) Omit(cols ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ebSystemStoreDo) Join(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ebSystemStoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ebSystemStoreDo) RightJoin(table schema.Tabler, on ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ebSystemStoreDo) Group(cols ...field.Expr) IEbSystemStoreDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ebSystemStoreDo) Having(conds ...gen.Condition) IEbSystemStoreDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ebSystemStoreDo) Limit(limit int) IEbSystemStoreDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ebSystemStoreDo) Offset(offset int) IEbSystemStoreDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ebSystemStoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEbSystemStoreDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ebSystemStoreDo) Unscoped() IEbSystemStoreDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ebSystemStoreDo) Create(values ...*model.EbSystemStore) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ebSystemStoreDo) CreateInBatches(values []*model.EbSystemStore, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ebSystemStoreDo) Save(values ...*model.EbSystemStore) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ebSystemStoreDo) First() (*model.EbSystemStore, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemStore), nil
	}
}

func (e ebSystemStoreDo) Take() (*model.EbSystemStore, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemStore), nil
	}
}

func (e ebSystemStoreDo) Last() (*model.EbSystemStore, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemStore), nil
	}
}

func (e ebSystemStoreDo) Find() ([]*model.EbSystemStore, error) {
	result, err := e.DO.Find()
	return result.([]*model.EbSystemStore), err
}

func (e ebSystemStoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EbSystemStore, err error) {
	buf := make([]*model.EbSystemStore, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ebSystemStoreDo) FindInBatches(result *[]*model.EbSystemStore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ebSystemStoreDo) Attrs(attrs ...field.AssignExpr) IEbSystemStoreDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ebSystemStoreDo) Assign(attrs ...field.AssignExpr) IEbSystemStoreDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ebSystemStoreDo) Joins(fields ...field.RelationField) IEbSystemStoreDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ebSystemStoreDo) Preload(fields ...field.RelationField) IEbSystemStoreDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ebSystemStoreDo) FirstOrInit() (*model.EbSystemStore, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemStore), nil
	}
}

func (e ebSystemStoreDo) FirstOrCreate() (*model.EbSystemStore, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EbSystemStore), nil
	}
}

func (e ebSystemStoreDo) FindByPage(offset int, limit int) (result []*model.EbSystemStore, count int64, err error) {
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

func (e ebSystemStoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ebSystemStoreDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ebSystemStoreDo) Delete(models ...*model.EbSystemStore) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ebSystemStoreDo) withDO(do gen.Dao) *ebSystemStoreDo {
	e.DO = *do.(*gen.DO)
	return e
}
