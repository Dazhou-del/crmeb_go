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

func newStoreProductRelation(db *gorm.DB, opts ...gen.DOOption) storeProductRelation {
	_storeProductRelation := storeProductRelation{}

	_storeProductRelation.storeProductRelationDo.UseDB(db, opts...)
	_storeProductRelation.storeProductRelationDo.UseModel(&model.StoreProductRelation{})

	tableName := _storeProductRelation.storeProductRelationDo.TableName()
	_storeProductRelation.ALL = field.NewAsterisk(tableName)
	_storeProductRelation.ID = field.NewInt64(tableName, "id")
	_storeProductRelation.UID = field.NewInt64(tableName, "uid")
	_storeProductRelation.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductRelation.Type = field.NewString(tableName, "type")
	_storeProductRelation.Category = field.NewString(tableName, "category")
	_storeProductRelation.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductRelation.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductRelation.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductRelation.fillFieldMap()

	return _storeProductRelation
}

// storeProductRelation 商品点赞和收藏表
type storeProductRelation struct {
	storeProductRelationDo storeProductRelationDo

	ALL       field.Asterisk
	ID        field.Int64  // id
	UID       field.Int64  // 用户ID
	ProductID field.Int64  // 商品ID
	Type      field.String // 类型(收藏(collect）、点赞(like))
	Category  field.String // 某种类型的商品(普通商品、秒杀商品)
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s storeProductRelation) Table(newTableName string) *storeProductRelation {
	s.storeProductRelationDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductRelation) As(alias string) *storeProductRelation {
	s.storeProductRelationDo.DO = *(s.storeProductRelationDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductRelation) updateTableName(table string) *storeProductRelation {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.ProductID = field.NewInt64(table, "product_id")
	s.Type = field.NewString(table, "type")
	s.Category = field.NewString(table, "category")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductRelation) WithContext(ctx context.Context) IStoreProductRelationDo {
	return s.storeProductRelationDo.WithContext(ctx)
}

func (s storeProductRelation) TableName() string { return s.storeProductRelationDo.TableName() }

func (s storeProductRelation) Alias() string { return s.storeProductRelationDo.Alias() }

func (s storeProductRelation) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductRelationDo.Columns(cols...)
}

func (s *storeProductRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductRelation) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["type"] = s.Type
	s.fieldMap["category"] = s.Category
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductRelation) clone(db *gorm.DB) storeProductRelation {
	s.storeProductRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductRelation) replaceDB(db *gorm.DB) storeProductRelation {
	s.storeProductRelationDo.ReplaceDB(db)
	return s
}

type storeProductRelationDo struct{ gen.DO }

type IStoreProductRelationDo interface {
	gen.SubQuery
	Debug() IStoreProductRelationDo
	WithContext(ctx context.Context) IStoreProductRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductRelationDo
	WriteDB() IStoreProductRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductRelationDo
	Not(conds ...gen.Condition) IStoreProductRelationDo
	Or(conds ...gen.Condition) IStoreProductRelationDo
	Select(conds ...field.Expr) IStoreProductRelationDo
	Where(conds ...gen.Condition) IStoreProductRelationDo
	Order(conds ...field.Expr) IStoreProductRelationDo
	Distinct(cols ...field.Expr) IStoreProductRelationDo
	Omit(cols ...field.Expr) IStoreProductRelationDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo
	Group(cols ...field.Expr) IStoreProductRelationDo
	Having(conds ...gen.Condition) IStoreProductRelationDo
	Limit(limit int) IStoreProductRelationDo
	Offset(offset int) IStoreProductRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductRelationDo
	Unscoped() IStoreProductRelationDo
	Create(values ...*model.StoreProductRelation) error
	CreateInBatches(values []*model.StoreProductRelation, batchSize int) error
	Save(values ...*model.StoreProductRelation) error
	First() (*model.StoreProductRelation, error)
	Take() (*model.StoreProductRelation, error)
	Last() (*model.StoreProductRelation, error)
	Find() ([]*model.StoreProductRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductRelation, err error)
	FindInBatches(result *[]*model.StoreProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductRelationDo
	Assign(attrs ...field.AssignExpr) IStoreProductRelationDo
	Joins(fields ...field.RelationField) IStoreProductRelationDo
	Preload(fields ...field.RelationField) IStoreProductRelationDo
	FirstOrInit() (*model.StoreProductRelation, error)
	FirstOrCreate() (*model.StoreProductRelation, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductRelationDo) Debug() IStoreProductRelationDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductRelationDo) WithContext(ctx context.Context) IStoreProductRelationDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductRelationDo) ReadDB() IStoreProductRelationDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductRelationDo) WriteDB() IStoreProductRelationDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductRelationDo) Session(config *gorm.Session) IStoreProductRelationDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductRelationDo) Clauses(conds ...clause.Expression) IStoreProductRelationDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductRelationDo) Returning(value interface{}, columns ...string) IStoreProductRelationDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductRelationDo) Not(conds ...gen.Condition) IStoreProductRelationDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductRelationDo) Or(conds ...gen.Condition) IStoreProductRelationDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductRelationDo) Select(conds ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductRelationDo) Where(conds ...gen.Condition) IStoreProductRelationDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductRelationDo) Order(conds ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductRelationDo) Distinct(cols ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductRelationDo) Omit(cols ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductRelationDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductRelationDo) Group(cols ...field.Expr) IStoreProductRelationDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductRelationDo) Having(conds ...gen.Condition) IStoreProductRelationDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductRelationDo) Limit(limit int) IStoreProductRelationDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductRelationDo) Offset(offset int) IStoreProductRelationDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductRelationDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductRelationDo) Unscoped() IStoreProductRelationDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductRelationDo) Create(values ...*model.StoreProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductRelationDo) CreateInBatches(values []*model.StoreProductRelation, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductRelationDo) Save(values ...*model.StoreProductRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductRelationDo) First() (*model.StoreProductRelation, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRelation), nil
	}
}

func (s storeProductRelationDo) Take() (*model.StoreProductRelation, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRelation), nil
	}
}

func (s storeProductRelationDo) Last() (*model.StoreProductRelation, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRelation), nil
	}
}

func (s storeProductRelationDo) Find() ([]*model.StoreProductRelation, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductRelation), err
}

func (s storeProductRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductRelation, err error) {
	buf := make([]*model.StoreProductRelation, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductRelationDo) FindInBatches(result *[]*model.StoreProductRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductRelationDo) Attrs(attrs ...field.AssignExpr) IStoreProductRelationDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductRelationDo) Assign(attrs ...field.AssignExpr) IStoreProductRelationDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductRelationDo) Joins(fields ...field.RelationField) IStoreProductRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductRelationDo) Preload(fields ...field.RelationField) IStoreProductRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductRelationDo) FirstOrInit() (*model.StoreProductRelation, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRelation), nil
	}
}

func (s storeProductRelationDo) FirstOrCreate() (*model.StoreProductRelation, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRelation), nil
	}
}

func (s storeProductRelationDo) FindByPage(offset int, limit int) (result []*model.StoreProductRelation, count int64, err error) {
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

func (s storeProductRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductRelationDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductRelationDo) Delete(models ...*model.StoreProductRelation) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductRelationDo) withDO(do gen.Dao) *storeProductRelationDo {
	s.DO = *do.(*gen.DO)
	return s
}
