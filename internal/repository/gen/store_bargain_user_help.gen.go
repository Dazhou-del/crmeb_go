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

func newStoreBargainUserHelp(db *gorm.DB, opts ...gen.DOOption) storeBargainUserHelp {
	_storeBargainUserHelp := storeBargainUserHelp{}

	_storeBargainUserHelp.storeBargainUserHelpDo.UseDB(db, opts...)
	_storeBargainUserHelp.storeBargainUserHelpDo.UseModel(&model.StoreBargainUserHelp{})

	tableName := _storeBargainUserHelp.storeBargainUserHelpDo.TableName()
	_storeBargainUserHelp.ALL = field.NewAsterisk(tableName)
	_storeBargainUserHelp.ID = field.NewInt64(tableName, "id")
	_storeBargainUserHelp.UID = field.NewInt64(tableName, "uid")
	_storeBargainUserHelp.BargainID = field.NewInt64(tableName, "bargain_id")
	_storeBargainUserHelp.BargainUserID = field.NewInt64(tableName, "bargain_user_id")
	_storeBargainUserHelp.Price = field.NewString(tableName, "price")
	_storeBargainUserHelp.AddTime = field.NewInt64(tableName, "add_time")
	_storeBargainUserHelp.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeBargainUserHelp.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeBargainUserHelp.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeBargainUserHelp.fillFieldMap()

	return _storeBargainUserHelp
}

// storeBargainUserHelp 砍价用户帮助表
type storeBargainUserHelp struct {
	storeBargainUserHelpDo storeBargainUserHelpDo

	ALL           field.Asterisk
	ID            field.Int64  // 砍价用户帮助表ID
	UID           field.Int64  // 帮助的用户id
	BargainID     field.Int64  // 砍价商品ID
	BargainUserID field.Int64  // 用户参与砍价表id
	Price         field.String // 帮助砍价多少金额
	AddTime       field.Int64  // 添加时间
	CreatedAt     field.Int64
	UpdatedAt     field.Int64
	DeletedAt     field.Field

	fieldMap map[string]field.Expr
}

func (s storeBargainUserHelp) Table(newTableName string) *storeBargainUserHelp {
	s.storeBargainUserHelpDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeBargainUserHelp) As(alias string) *storeBargainUserHelp {
	s.storeBargainUserHelpDo.DO = *(s.storeBargainUserHelpDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeBargainUserHelp) updateTableName(table string) *storeBargainUserHelp {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.BargainID = field.NewInt64(table, "bargain_id")
	s.BargainUserID = field.NewInt64(table, "bargain_user_id")
	s.Price = field.NewString(table, "price")
	s.AddTime = field.NewInt64(table, "add_time")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeBargainUserHelp) WithContext(ctx context.Context) IStoreBargainUserHelpDo {
	return s.storeBargainUserHelpDo.WithContext(ctx)
}

func (s storeBargainUserHelp) TableName() string { return s.storeBargainUserHelpDo.TableName() }

func (s storeBargainUserHelp) Alias() string { return s.storeBargainUserHelpDo.Alias() }

func (s storeBargainUserHelp) Columns(cols ...field.Expr) gen.Columns {
	return s.storeBargainUserHelpDo.Columns(cols...)
}

func (s *storeBargainUserHelp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeBargainUserHelp) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["bargain_id"] = s.BargainID
	s.fieldMap["bargain_user_id"] = s.BargainUserID
	s.fieldMap["price"] = s.Price
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeBargainUserHelp) clone(db *gorm.DB) storeBargainUserHelp {
	s.storeBargainUserHelpDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeBargainUserHelp) replaceDB(db *gorm.DB) storeBargainUserHelp {
	s.storeBargainUserHelpDo.ReplaceDB(db)
	return s
}

type storeBargainUserHelpDo struct{ gen.DO }

type IStoreBargainUserHelpDo interface {
	gen.SubQuery
	Debug() IStoreBargainUserHelpDo
	WithContext(ctx context.Context) IStoreBargainUserHelpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreBargainUserHelpDo
	WriteDB() IStoreBargainUserHelpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreBargainUserHelpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreBargainUserHelpDo
	Not(conds ...gen.Condition) IStoreBargainUserHelpDo
	Or(conds ...gen.Condition) IStoreBargainUserHelpDo
	Select(conds ...field.Expr) IStoreBargainUserHelpDo
	Where(conds ...gen.Condition) IStoreBargainUserHelpDo
	Order(conds ...field.Expr) IStoreBargainUserHelpDo
	Distinct(cols ...field.Expr) IStoreBargainUserHelpDo
	Omit(cols ...field.Expr) IStoreBargainUserHelpDo
	Join(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo
	Group(cols ...field.Expr) IStoreBargainUserHelpDo
	Having(conds ...gen.Condition) IStoreBargainUserHelpDo
	Limit(limit int) IStoreBargainUserHelpDo
	Offset(offset int) IStoreBargainUserHelpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainUserHelpDo
	Unscoped() IStoreBargainUserHelpDo
	Create(values ...*model.StoreBargainUserHelp) error
	CreateInBatches(values []*model.StoreBargainUserHelp, batchSize int) error
	Save(values ...*model.StoreBargainUserHelp) error
	First() (*model.StoreBargainUserHelp, error)
	Take() (*model.StoreBargainUserHelp, error)
	Last() (*model.StoreBargainUserHelp, error)
	Find() ([]*model.StoreBargainUserHelp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargainUserHelp, err error)
	FindInBatches(result *[]*model.StoreBargainUserHelp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreBargainUserHelp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreBargainUserHelpDo
	Assign(attrs ...field.AssignExpr) IStoreBargainUserHelpDo
	Joins(fields ...field.RelationField) IStoreBargainUserHelpDo
	Preload(fields ...field.RelationField) IStoreBargainUserHelpDo
	FirstOrInit() (*model.StoreBargainUserHelp, error)
	FirstOrCreate() (*model.StoreBargainUserHelp, error)
	FindByPage(offset int, limit int) (result []*model.StoreBargainUserHelp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreBargainUserHelpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeBargainUserHelpDo) Debug() IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Debug())
}

func (s storeBargainUserHelpDo) WithContext(ctx context.Context) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeBargainUserHelpDo) ReadDB() IStoreBargainUserHelpDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeBargainUserHelpDo) WriteDB() IStoreBargainUserHelpDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeBargainUserHelpDo) Session(config *gorm.Session) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeBargainUserHelpDo) Clauses(conds ...clause.Expression) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeBargainUserHelpDo) Returning(value interface{}, columns ...string) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeBargainUserHelpDo) Not(conds ...gen.Condition) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeBargainUserHelpDo) Or(conds ...gen.Condition) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeBargainUserHelpDo) Select(conds ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeBargainUserHelpDo) Where(conds ...gen.Condition) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeBargainUserHelpDo) Order(conds ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeBargainUserHelpDo) Distinct(cols ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeBargainUserHelpDo) Omit(cols ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeBargainUserHelpDo) Join(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeBargainUserHelpDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeBargainUserHelpDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeBargainUserHelpDo) Group(cols ...field.Expr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeBargainUserHelpDo) Having(conds ...gen.Condition) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeBargainUserHelpDo) Limit(limit int) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeBargainUserHelpDo) Offset(offset int) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeBargainUserHelpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeBargainUserHelpDo) Unscoped() IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeBargainUserHelpDo) Create(values ...*model.StoreBargainUserHelp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeBargainUserHelpDo) CreateInBatches(values []*model.StoreBargainUserHelp, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeBargainUserHelpDo) Save(values ...*model.StoreBargainUserHelp) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeBargainUserHelpDo) First() (*model.StoreBargainUserHelp, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUserHelp), nil
	}
}

func (s storeBargainUserHelpDo) Take() (*model.StoreBargainUserHelp, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUserHelp), nil
	}
}

func (s storeBargainUserHelpDo) Last() (*model.StoreBargainUserHelp, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUserHelp), nil
	}
}

func (s storeBargainUserHelpDo) Find() ([]*model.StoreBargainUserHelp, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreBargainUserHelp), err
}

func (s storeBargainUserHelpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargainUserHelp, err error) {
	buf := make([]*model.StoreBargainUserHelp, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeBargainUserHelpDo) FindInBatches(result *[]*model.StoreBargainUserHelp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeBargainUserHelpDo) Attrs(attrs ...field.AssignExpr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeBargainUserHelpDo) Assign(attrs ...field.AssignExpr) IStoreBargainUserHelpDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeBargainUserHelpDo) Joins(fields ...field.RelationField) IStoreBargainUserHelpDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeBargainUserHelpDo) Preload(fields ...field.RelationField) IStoreBargainUserHelpDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeBargainUserHelpDo) FirstOrInit() (*model.StoreBargainUserHelp, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUserHelp), nil
	}
}

func (s storeBargainUserHelpDo) FirstOrCreate() (*model.StoreBargainUserHelp, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUserHelp), nil
	}
}

func (s storeBargainUserHelpDo) FindByPage(offset int, limit int) (result []*model.StoreBargainUserHelp, count int64, err error) {
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

func (s storeBargainUserHelpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeBargainUserHelpDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeBargainUserHelpDo) Delete(models ...*model.StoreBargainUserHelp) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeBargainUserHelpDo) withDO(do gen.Dao) *storeBargainUserHelpDo {
	s.DO = *do.(*gen.DO)
	return s
}
