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

func newStoreBargainUser(db *gorm.DB, opts ...gen.DOOption) storeBargainUser {
	_storeBargainUser := storeBargainUser{}

	_storeBargainUser.storeBargainUserDo.UseDB(db, opts...)
	_storeBargainUser.storeBargainUserDo.UseModel(&model.StoreBargainUser{})

	tableName := _storeBargainUser.storeBargainUserDo.TableName()
	_storeBargainUser.ALL = field.NewAsterisk(tableName)
	_storeBargainUser.ID = field.NewInt64(tableName, "id")
	_storeBargainUser.UID = field.NewInt64(tableName, "uid")
	_storeBargainUser.BargainID = field.NewInt64(tableName, "bargain_id")
	_storeBargainUser.BargainPriceMin = field.NewString(tableName, "bargain_price_min")
	_storeBargainUser.BargainPrice = field.NewString(tableName, "bargain_price")
	_storeBargainUser.Price = field.NewString(tableName, "price")
	_storeBargainUser.Status = field.NewInt64(tableName, "status")
	_storeBargainUser.AddTime = field.NewInt64(tableName, "add_time")
	_storeBargainUser.IsDel = field.NewInt64(tableName, "is_del")
	_storeBargainUser.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeBargainUser.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeBargainUser.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeBargainUser.fillFieldMap()

	return _storeBargainUser
}

// storeBargainUser 用户参与砍价表
type storeBargainUser struct {
	storeBargainUserDo storeBargainUserDo

	ALL             field.Asterisk
	ID              field.Int64  // 用户参与砍价表ID
	UID             field.Int64  // 用户ID
	BargainID       field.Int64  // 砍价商品id
	BargainPriceMin field.String // 砍价的最低价
	BargainPrice    field.String // 砍价金额
	Price           field.String // 砍掉的价格
	Status          field.Int64  // 状态 1参与中 2 活动结束参与失败 3活动结束参与成功
	AddTime         field.Int64  // 参与时间
	IsDel           field.Int64  // 是否取消
	CreatedAt       field.Int64
	UpdatedAt       field.Int64
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (s storeBargainUser) Table(newTableName string) *storeBargainUser {
	s.storeBargainUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeBargainUser) As(alias string) *storeBargainUser {
	s.storeBargainUserDo.DO = *(s.storeBargainUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeBargainUser) updateTableName(table string) *storeBargainUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewInt64(table, "uid")
	s.BargainID = field.NewInt64(table, "bargain_id")
	s.BargainPriceMin = field.NewString(table, "bargain_price_min")
	s.BargainPrice = field.NewString(table, "bargain_price")
	s.Price = field.NewString(table, "price")
	s.Status = field.NewInt64(table, "status")
	s.AddTime = field.NewInt64(table, "add_time")
	s.IsDel = field.NewInt64(table, "is_del")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeBargainUser) WithContext(ctx context.Context) IStoreBargainUserDo {
	return s.storeBargainUserDo.WithContext(ctx)
}

func (s storeBargainUser) TableName() string { return s.storeBargainUserDo.TableName() }

func (s storeBargainUser) Alias() string { return s.storeBargainUserDo.Alias() }

func (s storeBargainUser) Columns(cols ...field.Expr) gen.Columns {
	return s.storeBargainUserDo.Columns(cols...)
}

func (s *storeBargainUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeBargainUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["bargain_id"] = s.BargainID
	s.fieldMap["bargain_price_min"] = s.BargainPriceMin
	s.fieldMap["bargain_price"] = s.BargainPrice
	s.fieldMap["price"] = s.Price
	s.fieldMap["status"] = s.Status
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["is_del"] = s.IsDel
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeBargainUser) clone(db *gorm.DB) storeBargainUser {
	s.storeBargainUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeBargainUser) replaceDB(db *gorm.DB) storeBargainUser {
	s.storeBargainUserDo.ReplaceDB(db)
	return s
}

type storeBargainUserDo struct{ gen.DO }

type IStoreBargainUserDo interface {
	gen.SubQuery
	Debug() IStoreBargainUserDo
	WithContext(ctx context.Context) IStoreBargainUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreBargainUserDo
	WriteDB() IStoreBargainUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreBargainUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreBargainUserDo
	Not(conds ...gen.Condition) IStoreBargainUserDo
	Or(conds ...gen.Condition) IStoreBargainUserDo
	Select(conds ...field.Expr) IStoreBargainUserDo
	Where(conds ...gen.Condition) IStoreBargainUserDo
	Order(conds ...field.Expr) IStoreBargainUserDo
	Distinct(cols ...field.Expr) IStoreBargainUserDo
	Omit(cols ...field.Expr) IStoreBargainUserDo
	Join(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo
	Group(cols ...field.Expr) IStoreBargainUserDo
	Having(conds ...gen.Condition) IStoreBargainUserDo
	Limit(limit int) IStoreBargainUserDo
	Offset(offset int) IStoreBargainUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainUserDo
	Unscoped() IStoreBargainUserDo
	Create(values ...*model.StoreBargainUser) error
	CreateInBatches(values []*model.StoreBargainUser, batchSize int) error
	Save(values ...*model.StoreBargainUser) error
	First() (*model.StoreBargainUser, error)
	Take() (*model.StoreBargainUser, error)
	Last() (*model.StoreBargainUser, error)
	Find() ([]*model.StoreBargainUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargainUser, err error)
	FindInBatches(result *[]*model.StoreBargainUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreBargainUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreBargainUserDo
	Assign(attrs ...field.AssignExpr) IStoreBargainUserDo
	Joins(fields ...field.RelationField) IStoreBargainUserDo
	Preload(fields ...field.RelationField) IStoreBargainUserDo
	FirstOrInit() (*model.StoreBargainUser, error)
	FirstOrCreate() (*model.StoreBargainUser, error)
	FindByPage(offset int, limit int) (result []*model.StoreBargainUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreBargainUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeBargainUserDo) Debug() IStoreBargainUserDo {
	return s.withDO(s.DO.Debug())
}

func (s storeBargainUserDo) WithContext(ctx context.Context) IStoreBargainUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeBargainUserDo) ReadDB() IStoreBargainUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeBargainUserDo) WriteDB() IStoreBargainUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeBargainUserDo) Session(config *gorm.Session) IStoreBargainUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeBargainUserDo) Clauses(conds ...clause.Expression) IStoreBargainUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeBargainUserDo) Returning(value interface{}, columns ...string) IStoreBargainUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeBargainUserDo) Not(conds ...gen.Condition) IStoreBargainUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeBargainUserDo) Or(conds ...gen.Condition) IStoreBargainUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeBargainUserDo) Select(conds ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeBargainUserDo) Where(conds ...gen.Condition) IStoreBargainUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeBargainUserDo) Order(conds ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeBargainUserDo) Distinct(cols ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeBargainUserDo) Omit(cols ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeBargainUserDo) Join(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeBargainUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeBargainUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeBargainUserDo) Group(cols ...field.Expr) IStoreBargainUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeBargainUserDo) Having(conds ...gen.Condition) IStoreBargainUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeBargainUserDo) Limit(limit int) IStoreBargainUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeBargainUserDo) Offset(offset int) IStoreBargainUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeBargainUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreBargainUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeBargainUserDo) Unscoped() IStoreBargainUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeBargainUserDo) Create(values ...*model.StoreBargainUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeBargainUserDo) CreateInBatches(values []*model.StoreBargainUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeBargainUserDo) Save(values ...*model.StoreBargainUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeBargainUserDo) First() (*model.StoreBargainUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUser), nil
	}
}

func (s storeBargainUserDo) Take() (*model.StoreBargainUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUser), nil
	}
}

func (s storeBargainUserDo) Last() (*model.StoreBargainUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUser), nil
	}
}

func (s storeBargainUserDo) Find() ([]*model.StoreBargainUser, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreBargainUser), err
}

func (s storeBargainUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreBargainUser, err error) {
	buf := make([]*model.StoreBargainUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeBargainUserDo) FindInBatches(result *[]*model.StoreBargainUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeBargainUserDo) Attrs(attrs ...field.AssignExpr) IStoreBargainUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeBargainUserDo) Assign(attrs ...field.AssignExpr) IStoreBargainUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeBargainUserDo) Joins(fields ...field.RelationField) IStoreBargainUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeBargainUserDo) Preload(fields ...field.RelationField) IStoreBargainUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeBargainUserDo) FirstOrInit() (*model.StoreBargainUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUser), nil
	}
}

func (s storeBargainUserDo) FirstOrCreate() (*model.StoreBargainUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreBargainUser), nil
	}
}

func (s storeBargainUserDo) FindByPage(offset int, limit int) (result []*model.StoreBargainUser, count int64, err error) {
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

func (s storeBargainUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeBargainUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeBargainUserDo) Delete(models ...*model.StoreBargainUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeBargainUserDo) withDO(do gen.Dao) *storeBargainUserDo {
	s.DO = *do.(*gen.DO)
	return s
}
