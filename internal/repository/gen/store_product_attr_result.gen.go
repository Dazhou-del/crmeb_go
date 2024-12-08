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

func newStoreProductAttrResult(db *gorm.DB, opts ...gen.DOOption) storeProductAttrResult {
	_storeProductAttrResult := storeProductAttrResult{}

	_storeProductAttrResult.storeProductAttrResultDo.UseDB(db, opts...)
	_storeProductAttrResult.storeProductAttrResultDo.UseModel(&model.StoreProductAttrResult{})

	tableName := _storeProductAttrResult.storeProductAttrResultDo.TableName()
	_storeProductAttrResult.ALL = field.NewAsterisk(tableName)
	_storeProductAttrResult.ID = field.NewInt64(tableName, "id")
	_storeProductAttrResult.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductAttrResult.Result = field.NewString(tableName, "result")
	_storeProductAttrResult.ChangeTime = field.NewInt64(tableName, "change_time")
	_storeProductAttrResult.Type = field.NewInt64(tableName, "type")
	_storeProductAttrResult.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductAttrResult.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductAttrResult.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductAttrResult.fillFieldMap()

	return _storeProductAttrResult
}

// storeProductAttrResult 商品属性详情表
type storeProductAttrResult struct {
	storeProductAttrResultDo storeProductAttrResultDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	ProductID  field.Int64  // 商品ID
	Result     field.String // 商品属性参数
	ChangeTime field.Int64  // 上次修改时间
	Type       field.Int64  // 活动类型 0=商品，1=秒杀，2=砍价，3=拼团
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s storeProductAttrResult) Table(newTableName string) *storeProductAttrResult {
	s.storeProductAttrResultDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductAttrResult) As(alias string) *storeProductAttrResult {
	s.storeProductAttrResultDo.DO = *(s.storeProductAttrResultDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductAttrResult) updateTableName(table string) *storeProductAttrResult {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.ProductID = field.NewInt64(table, "product_id")
	s.Result = field.NewString(table, "result")
	s.ChangeTime = field.NewInt64(table, "change_time")
	s.Type = field.NewInt64(table, "type")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductAttrResult) WithContext(ctx context.Context) IStoreProductAttrResultDo {
	return s.storeProductAttrResultDo.WithContext(ctx)
}

func (s storeProductAttrResult) TableName() string { return s.storeProductAttrResultDo.TableName() }

func (s storeProductAttrResult) Alias() string { return s.storeProductAttrResultDo.Alias() }

func (s storeProductAttrResult) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductAttrResultDo.Columns(cols...)
}

func (s *storeProductAttrResult) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductAttrResult) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["result"] = s.Result
	s.fieldMap["change_time"] = s.ChangeTime
	s.fieldMap["type"] = s.Type
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductAttrResult) clone(db *gorm.DB) storeProductAttrResult {
	s.storeProductAttrResultDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductAttrResult) replaceDB(db *gorm.DB) storeProductAttrResult {
	s.storeProductAttrResultDo.ReplaceDB(db)
	return s
}

type storeProductAttrResultDo struct{ gen.DO }

type IStoreProductAttrResultDo interface {
	gen.SubQuery
	Debug() IStoreProductAttrResultDo
	WithContext(ctx context.Context) IStoreProductAttrResultDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductAttrResultDo
	WriteDB() IStoreProductAttrResultDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductAttrResultDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductAttrResultDo
	Not(conds ...gen.Condition) IStoreProductAttrResultDo
	Or(conds ...gen.Condition) IStoreProductAttrResultDo
	Select(conds ...field.Expr) IStoreProductAttrResultDo
	Where(conds ...gen.Condition) IStoreProductAttrResultDo
	Order(conds ...field.Expr) IStoreProductAttrResultDo
	Distinct(cols ...field.Expr) IStoreProductAttrResultDo
	Omit(cols ...field.Expr) IStoreProductAttrResultDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo
	Group(cols ...field.Expr) IStoreProductAttrResultDo
	Having(conds ...gen.Condition) IStoreProductAttrResultDo
	Limit(limit int) IStoreProductAttrResultDo
	Offset(offset int) IStoreProductAttrResultDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductAttrResultDo
	Unscoped() IStoreProductAttrResultDo
	Create(values ...*model.StoreProductAttrResult) error
	CreateInBatches(values []*model.StoreProductAttrResult, batchSize int) error
	Save(values ...*model.StoreProductAttrResult) error
	First() (*model.StoreProductAttrResult, error)
	Take() (*model.StoreProductAttrResult, error)
	Last() (*model.StoreProductAttrResult, error)
	Find() ([]*model.StoreProductAttrResult, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductAttrResult, err error)
	FindInBatches(result *[]*model.StoreProductAttrResult, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductAttrResult) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductAttrResultDo
	Assign(attrs ...field.AssignExpr) IStoreProductAttrResultDo
	Joins(fields ...field.RelationField) IStoreProductAttrResultDo
	Preload(fields ...field.RelationField) IStoreProductAttrResultDo
	FirstOrInit() (*model.StoreProductAttrResult, error)
	FirstOrCreate() (*model.StoreProductAttrResult, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductAttrResult, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductAttrResultDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductAttrResultDo) Debug() IStoreProductAttrResultDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductAttrResultDo) WithContext(ctx context.Context) IStoreProductAttrResultDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductAttrResultDo) ReadDB() IStoreProductAttrResultDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductAttrResultDo) WriteDB() IStoreProductAttrResultDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductAttrResultDo) Session(config *gorm.Session) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductAttrResultDo) Clauses(conds ...clause.Expression) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductAttrResultDo) Returning(value interface{}, columns ...string) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductAttrResultDo) Not(conds ...gen.Condition) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductAttrResultDo) Or(conds ...gen.Condition) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductAttrResultDo) Select(conds ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductAttrResultDo) Where(conds ...gen.Condition) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductAttrResultDo) Order(conds ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductAttrResultDo) Distinct(cols ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductAttrResultDo) Omit(cols ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductAttrResultDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductAttrResultDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductAttrResultDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductAttrResultDo) Group(cols ...field.Expr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductAttrResultDo) Having(conds ...gen.Condition) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductAttrResultDo) Limit(limit int) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductAttrResultDo) Offset(offset int) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductAttrResultDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductAttrResultDo) Unscoped() IStoreProductAttrResultDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductAttrResultDo) Create(values ...*model.StoreProductAttrResult) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductAttrResultDo) CreateInBatches(values []*model.StoreProductAttrResult, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductAttrResultDo) Save(values ...*model.StoreProductAttrResult) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductAttrResultDo) First() (*model.StoreProductAttrResult, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrResult), nil
	}
}

func (s storeProductAttrResultDo) Take() (*model.StoreProductAttrResult, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrResult), nil
	}
}

func (s storeProductAttrResultDo) Last() (*model.StoreProductAttrResult, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrResult), nil
	}
}

func (s storeProductAttrResultDo) Find() ([]*model.StoreProductAttrResult, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductAttrResult), err
}

func (s storeProductAttrResultDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductAttrResult, err error) {
	buf := make([]*model.StoreProductAttrResult, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductAttrResultDo) FindInBatches(result *[]*model.StoreProductAttrResult, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductAttrResultDo) Attrs(attrs ...field.AssignExpr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductAttrResultDo) Assign(attrs ...field.AssignExpr) IStoreProductAttrResultDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductAttrResultDo) Joins(fields ...field.RelationField) IStoreProductAttrResultDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductAttrResultDo) Preload(fields ...field.RelationField) IStoreProductAttrResultDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductAttrResultDo) FirstOrInit() (*model.StoreProductAttrResult, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrResult), nil
	}
}

func (s storeProductAttrResultDo) FirstOrCreate() (*model.StoreProductAttrResult, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductAttrResult), nil
	}
}

func (s storeProductAttrResultDo) FindByPage(offset int, limit int) (result []*model.StoreProductAttrResult, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s storeProductAttrResultDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductAttrResultDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductAttrResultDo) Delete(models ...*model.StoreProductAttrResult) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductAttrResultDo) withDO(do gen.Dao) *storeProductAttrResultDo {
	s.DO = *do.(*gen.DO)
	return s
}
