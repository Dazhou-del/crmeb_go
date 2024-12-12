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

func newStoreProductRule(db *gorm.DB, opts ...gen.DOOption) storeProductRule {
	_storeProductRule := storeProductRule{}

	_storeProductRule.storeProductRuleDo.UseDB(db, opts...)
	_storeProductRule.storeProductRuleDo.UseModel(&model.StoreProductRule{})

	tableName := _storeProductRule.storeProductRuleDo.TableName()
	_storeProductRule.ALL = field.NewAsterisk(tableName)
	_storeProductRule.ID = field.NewInt64(tableName, "id")
	_storeProductRule.RuleName = field.NewString(tableName, "rule_name")
	_storeProductRule.RuleValue = field.NewString(tableName, "rule_value")
	_storeProductRule.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductRule.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductRule.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductRule.fillFieldMap()

	return _storeProductRule
}

// storeProductRule 商品规则值(规格)表
type storeProductRule struct {
	storeProductRuleDo storeProductRuleDo

	ALL       field.Asterisk
	ID        field.Int64
	RuleName  field.String // 规格名称
	RuleValue field.String // 规格值
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (s storeProductRule) Table(newTableName string) *storeProductRule {
	s.storeProductRuleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductRule) As(alias string) *storeProductRule {
	s.storeProductRuleDo.DO = *(s.storeProductRuleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductRule) updateTableName(table string) *storeProductRule {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.RuleName = field.NewString(table, "rule_name")
	s.RuleValue = field.NewString(table, "rule_value")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductRule) WithContext(ctx context.Context) IStoreProductRuleDo {
	return s.storeProductRuleDo.WithContext(ctx)
}

func (s storeProductRule) TableName() string { return s.storeProductRuleDo.TableName() }

func (s storeProductRule) Alias() string { return s.storeProductRuleDo.Alias() }

func (s storeProductRule) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductRuleDo.Columns(cols...)
}

func (s *storeProductRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductRule) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["rule_name"] = s.RuleName
	s.fieldMap["rule_value"] = s.RuleValue
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductRule) clone(db *gorm.DB) storeProductRule {
	s.storeProductRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductRule) replaceDB(db *gorm.DB) storeProductRule {
	s.storeProductRuleDo.ReplaceDB(db)
	return s
}

type storeProductRuleDo struct{ gen.DO }

type IStoreProductRuleDo interface {
	gen.SubQuery
	Debug() IStoreProductRuleDo
	WithContext(ctx context.Context) IStoreProductRuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductRuleDo
	WriteDB() IStoreProductRuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductRuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductRuleDo
	Not(conds ...gen.Condition) IStoreProductRuleDo
	Or(conds ...gen.Condition) IStoreProductRuleDo
	Select(conds ...field.Expr) IStoreProductRuleDo
	Where(conds ...gen.Condition) IStoreProductRuleDo
	Order(conds ...field.Expr) IStoreProductRuleDo
	Distinct(cols ...field.Expr) IStoreProductRuleDo
	Omit(cols ...field.Expr) IStoreProductRuleDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo
	Group(cols ...field.Expr) IStoreProductRuleDo
	Having(conds ...gen.Condition) IStoreProductRuleDo
	Limit(limit int) IStoreProductRuleDo
	Offset(offset int) IStoreProductRuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductRuleDo
	Unscoped() IStoreProductRuleDo
	Create(values ...*model.StoreProductRule) error
	CreateInBatches(values []*model.StoreProductRule, batchSize int) error
	Save(values ...*model.StoreProductRule) error
	First() (*model.StoreProductRule, error)
	Take() (*model.StoreProductRule, error)
	Last() (*model.StoreProductRule, error)
	Find() ([]*model.StoreProductRule, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductRule, err error)
	FindInBatches(result *[]*model.StoreProductRule, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductRule) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductRuleDo
	Assign(attrs ...field.AssignExpr) IStoreProductRuleDo
	Joins(fields ...field.RelationField) IStoreProductRuleDo
	Preload(fields ...field.RelationField) IStoreProductRuleDo
	FirstOrInit() (*model.StoreProductRule, error)
	FirstOrCreate() (*model.StoreProductRule, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductRule, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductRuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductRuleDo) Debug() IStoreProductRuleDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductRuleDo) WithContext(ctx context.Context) IStoreProductRuleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductRuleDo) ReadDB() IStoreProductRuleDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductRuleDo) WriteDB() IStoreProductRuleDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductRuleDo) Session(config *gorm.Session) IStoreProductRuleDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductRuleDo) Clauses(conds ...clause.Expression) IStoreProductRuleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductRuleDo) Returning(value interface{}, columns ...string) IStoreProductRuleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductRuleDo) Not(conds ...gen.Condition) IStoreProductRuleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductRuleDo) Or(conds ...gen.Condition) IStoreProductRuleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductRuleDo) Select(conds ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductRuleDo) Where(conds ...gen.Condition) IStoreProductRuleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductRuleDo) Order(conds ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductRuleDo) Distinct(cols ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductRuleDo) Omit(cols ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductRuleDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductRuleDo) Group(cols ...field.Expr) IStoreProductRuleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductRuleDo) Having(conds ...gen.Condition) IStoreProductRuleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductRuleDo) Limit(limit int) IStoreProductRuleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductRuleDo) Offset(offset int) IStoreProductRuleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductRuleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductRuleDo) Unscoped() IStoreProductRuleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductRuleDo) Create(values ...*model.StoreProductRule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductRuleDo) CreateInBatches(values []*model.StoreProductRule, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductRuleDo) Save(values ...*model.StoreProductRule) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductRuleDo) First() (*model.StoreProductRule, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRule), nil
	}
}

func (s storeProductRuleDo) Take() (*model.StoreProductRule, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRule), nil
	}
}

func (s storeProductRuleDo) Last() (*model.StoreProductRule, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRule), nil
	}
}

func (s storeProductRuleDo) Find() ([]*model.StoreProductRule, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductRule), err
}

func (s storeProductRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductRule, err error) {
	buf := make([]*model.StoreProductRule, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductRuleDo) FindInBatches(result *[]*model.StoreProductRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductRuleDo) Attrs(attrs ...field.AssignExpr) IStoreProductRuleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductRuleDo) Assign(attrs ...field.AssignExpr) IStoreProductRuleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductRuleDo) Joins(fields ...field.RelationField) IStoreProductRuleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductRuleDo) Preload(fields ...field.RelationField) IStoreProductRuleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductRuleDo) FirstOrInit() (*model.StoreProductRule, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRule), nil
	}
}

func (s storeProductRuleDo) FirstOrCreate() (*model.StoreProductRule, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductRule), nil
	}
}

func (s storeProductRuleDo) FindByPage(offset int, limit int) (result []*model.StoreProductRule, count int64, err error) {
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

func (s storeProductRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductRuleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductRuleDo) Delete(models ...*model.StoreProductRule) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductRuleDo) withDO(do gen.Dao) *storeProductRuleDo {
	s.DO = *do.(*gen.DO)
	return s
}