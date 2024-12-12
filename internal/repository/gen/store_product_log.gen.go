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

func newStoreProductLog(db *gorm.DB, opts ...gen.DOOption) storeProductLog {
	_storeProductLog := storeProductLog{}

	_storeProductLog.storeProductLogDo.UseDB(db, opts...)
	_storeProductLog.storeProductLogDo.UseModel(&model.StoreProductLog{})

	tableName := _storeProductLog.storeProductLogDo.TableName()
	_storeProductLog.ALL = field.NewAsterisk(tableName)
	_storeProductLog.ID = field.NewInt64(tableName, "id")
	_storeProductLog.Type = field.NewString(tableName, "type")
	_storeProductLog.ProductID = field.NewInt64(tableName, "product_id")
	_storeProductLog.UID = field.NewInt64(tableName, "uid")
	_storeProductLog.VisitNum = field.NewInt64(tableName, "visit_num")
	_storeProductLog.CartNum = field.NewInt64(tableName, "cart_num")
	_storeProductLog.OrderNum = field.NewInt64(tableName, "order_num")
	_storeProductLog.PayNum = field.NewInt64(tableName, "pay_num")
	_storeProductLog.PayPrice = field.NewField(tableName, "pay_price")
	_storeProductLog.CostPrice = field.NewField(tableName, "cost_price")
	_storeProductLog.PayUID = field.NewInt64(tableName, "pay_uid")
	_storeProductLog.RefundNum = field.NewInt64(tableName, "refund_num")
	_storeProductLog.RefundPrice = field.NewField(tableName, "refund_price")
	_storeProductLog.CollectNum = field.NewInt64(tableName, "collect_num")
	_storeProductLog.AddTime = field.NewInt64(tableName, "add_time")
	_storeProductLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_storeProductLog.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_storeProductLog.DeletedAt = field.NewField(tableName, "deleted_at")

	_storeProductLog.fillFieldMap()

	return _storeProductLog
}

// storeProductLog 商品日志表
type storeProductLog struct {
	storeProductLogDo storeProductLogDo

	ALL         field.Asterisk
	ID          field.Int64  // 统计ID
	Type        field.String // 类型visit,cart,order,pay,collect,refund
	ProductID   field.Int64  // 商品ID
	UID         field.Int64  // 用户ID
	VisitNum    field.Int64  // 是否浏览
	CartNum     field.Int64  // 加入购物车数量
	OrderNum    field.Int64  // 下单数量
	PayNum      field.Int64  // 支付数量
	PayPrice    field.Field  // 支付金额
	CostPrice   field.Field  // 商品成本价
	PayUID      field.Int64  // 支付用户ID
	RefundNum   field.Int64  // 退款数量
	RefundPrice field.Field  // 退款金额
	CollectNum  field.Int64  // 收藏
	AddTime     field.Int64  // 添加时间
	CreatedAt   field.Int64
	UpdatedAt   field.Int64
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (s storeProductLog) Table(newTableName string) *storeProductLog {
	s.storeProductLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storeProductLog) As(alias string) *storeProductLog {
	s.storeProductLogDo.DO = *(s.storeProductLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storeProductLog) updateTableName(table string) *storeProductLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Type = field.NewString(table, "type")
	s.ProductID = field.NewInt64(table, "product_id")
	s.UID = field.NewInt64(table, "uid")
	s.VisitNum = field.NewInt64(table, "visit_num")
	s.CartNum = field.NewInt64(table, "cart_num")
	s.OrderNum = field.NewInt64(table, "order_num")
	s.PayNum = field.NewInt64(table, "pay_num")
	s.PayPrice = field.NewField(table, "pay_price")
	s.CostPrice = field.NewField(table, "cost_price")
	s.PayUID = field.NewInt64(table, "pay_uid")
	s.RefundNum = field.NewInt64(table, "refund_num")
	s.RefundPrice = field.NewField(table, "refund_price")
	s.CollectNum = field.NewInt64(table, "collect_num")
	s.AddTime = field.NewInt64(table, "add_time")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *storeProductLog) WithContext(ctx context.Context) IStoreProductLogDo {
	return s.storeProductLogDo.WithContext(ctx)
}

func (s storeProductLog) TableName() string { return s.storeProductLogDo.TableName() }

func (s storeProductLog) Alias() string { return s.storeProductLogDo.Alias() }

func (s storeProductLog) Columns(cols ...field.Expr) gen.Columns {
	return s.storeProductLogDo.Columns(cols...)
}

func (s *storeProductLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storeProductLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 18)
	s.fieldMap["id"] = s.ID
	s.fieldMap["type"] = s.Type
	s.fieldMap["product_id"] = s.ProductID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["visit_num"] = s.VisitNum
	s.fieldMap["cart_num"] = s.CartNum
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["pay_num"] = s.PayNum
	s.fieldMap["pay_price"] = s.PayPrice
	s.fieldMap["cost_price"] = s.CostPrice
	s.fieldMap["pay_uid"] = s.PayUID
	s.fieldMap["refund_num"] = s.RefundNum
	s.fieldMap["refund_price"] = s.RefundPrice
	s.fieldMap["collect_num"] = s.CollectNum
	s.fieldMap["add_time"] = s.AddTime
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s storeProductLog) clone(db *gorm.DB) storeProductLog {
	s.storeProductLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s storeProductLog) replaceDB(db *gorm.DB) storeProductLog {
	s.storeProductLogDo.ReplaceDB(db)
	return s
}

type storeProductLogDo struct{ gen.DO }

type IStoreProductLogDo interface {
	gen.SubQuery
	Debug() IStoreProductLogDo
	WithContext(ctx context.Context) IStoreProductLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStoreProductLogDo
	WriteDB() IStoreProductLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStoreProductLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStoreProductLogDo
	Not(conds ...gen.Condition) IStoreProductLogDo
	Or(conds ...gen.Condition) IStoreProductLogDo
	Select(conds ...field.Expr) IStoreProductLogDo
	Where(conds ...gen.Condition) IStoreProductLogDo
	Order(conds ...field.Expr) IStoreProductLogDo
	Distinct(cols ...field.Expr) IStoreProductLogDo
	Omit(cols ...field.Expr) IStoreProductLogDo
	Join(table schema.Tabler, on ...field.Expr) IStoreProductLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductLogDo
	Group(cols ...field.Expr) IStoreProductLogDo
	Having(conds ...gen.Condition) IStoreProductLogDo
	Limit(limit int) IStoreProductLogDo
	Offset(offset int) IStoreProductLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductLogDo
	Unscoped() IStoreProductLogDo
	Create(values ...*model.StoreProductLog) error
	CreateInBatches(values []*model.StoreProductLog, batchSize int) error
	Save(values ...*model.StoreProductLog) error
	First() (*model.StoreProductLog, error)
	Take() (*model.StoreProductLog, error)
	Last() (*model.StoreProductLog, error)
	Find() ([]*model.StoreProductLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductLog, err error)
	FindInBatches(result *[]*model.StoreProductLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.StoreProductLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStoreProductLogDo
	Assign(attrs ...field.AssignExpr) IStoreProductLogDo
	Joins(fields ...field.RelationField) IStoreProductLogDo
	Preload(fields ...field.RelationField) IStoreProductLogDo
	FirstOrInit() (*model.StoreProductLog, error)
	FirstOrCreate() (*model.StoreProductLog, error)
	FindByPage(offset int, limit int) (result []*model.StoreProductLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStoreProductLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storeProductLogDo) Debug() IStoreProductLogDo {
	return s.withDO(s.DO.Debug())
}

func (s storeProductLogDo) WithContext(ctx context.Context) IStoreProductLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storeProductLogDo) ReadDB() IStoreProductLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s storeProductLogDo) WriteDB() IStoreProductLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s storeProductLogDo) Session(config *gorm.Session) IStoreProductLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s storeProductLogDo) Clauses(conds ...clause.Expression) IStoreProductLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storeProductLogDo) Returning(value interface{}, columns ...string) IStoreProductLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storeProductLogDo) Not(conds ...gen.Condition) IStoreProductLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storeProductLogDo) Or(conds ...gen.Condition) IStoreProductLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storeProductLogDo) Select(conds ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storeProductLogDo) Where(conds ...gen.Condition) IStoreProductLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storeProductLogDo) Order(conds ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storeProductLogDo) Distinct(cols ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storeProductLogDo) Omit(cols ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storeProductLogDo) Join(table schema.Tabler, on ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storeProductLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storeProductLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storeProductLogDo) Group(cols ...field.Expr) IStoreProductLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storeProductLogDo) Having(conds ...gen.Condition) IStoreProductLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storeProductLogDo) Limit(limit int) IStoreProductLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storeProductLogDo) Offset(offset int) IStoreProductLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storeProductLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStoreProductLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storeProductLogDo) Unscoped() IStoreProductLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storeProductLogDo) Create(values ...*model.StoreProductLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storeProductLogDo) CreateInBatches(values []*model.StoreProductLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storeProductLogDo) Save(values ...*model.StoreProductLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storeProductLogDo) First() (*model.StoreProductLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductLog), nil
	}
}

func (s storeProductLogDo) Take() (*model.StoreProductLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductLog), nil
	}
}

func (s storeProductLogDo) Last() (*model.StoreProductLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductLog), nil
	}
}

func (s storeProductLogDo) Find() ([]*model.StoreProductLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.StoreProductLog), err
}

func (s storeProductLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StoreProductLog, err error) {
	buf := make([]*model.StoreProductLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storeProductLogDo) FindInBatches(result *[]*model.StoreProductLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storeProductLogDo) Attrs(attrs ...field.AssignExpr) IStoreProductLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storeProductLogDo) Assign(attrs ...field.AssignExpr) IStoreProductLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storeProductLogDo) Joins(fields ...field.RelationField) IStoreProductLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storeProductLogDo) Preload(fields ...field.RelationField) IStoreProductLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storeProductLogDo) FirstOrInit() (*model.StoreProductLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductLog), nil
	}
}

func (s storeProductLogDo) FirstOrCreate() (*model.StoreProductLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StoreProductLog), nil
	}
}

func (s storeProductLogDo) FindByPage(offset int, limit int) (result []*model.StoreProductLog, count int64, err error) {
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

func (s storeProductLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storeProductLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s storeProductLogDo) Delete(models ...*model.StoreProductLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *storeProductLogDo) withDO(do gen.Dao) *storeProductLogDo {
	s.DO = *do.(*gen.DO)
	return s
}