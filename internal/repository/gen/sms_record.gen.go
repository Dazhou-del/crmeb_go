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

func newSmsRecord(db *gorm.DB, opts ...gen.DOOption) smsRecord {
	_smsRecord := smsRecord{}

	_smsRecord.smsRecordDo.UseDB(db, opts...)
	_smsRecord.smsRecordDo.UseModel(&model.SmsRecord{})

	tableName := _smsRecord.smsRecordDo.TableName()
	_smsRecord.ALL = field.NewAsterisk(tableName)
	_smsRecord.ID = field.NewInt64(tableName, "id")
	_smsRecord.UID = field.NewString(tableName, "uid")
	_smsRecord.Phone = field.NewString(tableName, "phone")
	_smsRecord.Content = field.NewString(tableName, "content")
	_smsRecord.AddIP = field.NewString(tableName, "add_ip")
	_smsRecord.Template = field.NewString(tableName, "template")
	_smsRecord.Resultcode = field.NewInt64(tableName, "resultcode")
	_smsRecord.RecordID = field.NewInt64(tableName, "record_id")
	_smsRecord.Memo = field.NewString(tableName, "memo")
	_smsRecord.CreatedAt = field.NewInt64(tableName, "created_at")
	_smsRecord.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_smsRecord.DeletedAt = field.NewField(tableName, "deleted_at")

	_smsRecord.fillFieldMap()

	return _smsRecord
}

// smsRecord 短信发送记录表
type smsRecord struct {
	smsRecordDo smsRecordDo

	ALL        field.Asterisk
	ID         field.Int64  // 短信发送记录编号
	UID        field.String // 短信平台账号
	Phone      field.String // 接受短信的手机号
	Content    field.String // 短信内容
	AddIP      field.String // 添加记录ip
	Template   field.String // 短信模板ID
	Resultcode field.Int64  // 状态码 100=成功,130=失败,131=空号,132=停机,133=关机,134=无状态
	RecordID   field.Int64  // 发送记录id
	Memo       field.String // 短信平台返回信息
	CreatedAt  field.Int64
	UpdatedAt  field.Int64
	DeletedAt  field.Field

	fieldMap map[string]field.Expr
}

func (s smsRecord) Table(newTableName string) *smsRecord {
	s.smsRecordDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsRecord) As(alias string) *smsRecord {
	s.smsRecordDo.DO = *(s.smsRecordDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsRecord) updateTableName(table string) *smsRecord {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UID = field.NewString(table, "uid")
	s.Phone = field.NewString(table, "phone")
	s.Content = field.NewString(table, "content")
	s.AddIP = field.NewString(table, "add_ip")
	s.Template = field.NewString(table, "template")
	s.Resultcode = field.NewInt64(table, "resultcode")
	s.RecordID = field.NewInt64(table, "record_id")
	s.Memo = field.NewString(table, "memo")
	s.CreatedAt = field.NewInt64(table, "created_at")
	s.UpdatedAt = field.NewInt64(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *smsRecord) WithContext(ctx context.Context) ISmsRecordDo {
	return s.smsRecordDo.WithContext(ctx)
}

func (s smsRecord) TableName() string { return s.smsRecordDo.TableName() }

func (s smsRecord) Alias() string { return s.smsRecordDo.Alias() }

func (s smsRecord) Columns(cols ...field.Expr) gen.Columns { return s.smsRecordDo.Columns(cols...) }

func (s *smsRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsRecord) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["uid"] = s.UID
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["content"] = s.Content
	s.fieldMap["add_ip"] = s.AddIP
	s.fieldMap["template"] = s.Template
	s.fieldMap["resultcode"] = s.Resultcode
	s.fieldMap["record_id"] = s.RecordID
	s.fieldMap["memo"] = s.Memo
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
}

func (s smsRecord) clone(db *gorm.DB) smsRecord {
	s.smsRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsRecord) replaceDB(db *gorm.DB) smsRecord {
	s.smsRecordDo.ReplaceDB(db)
	return s
}

type smsRecordDo struct{ gen.DO }

type ISmsRecordDo interface {
	gen.SubQuery
	Debug() ISmsRecordDo
	WithContext(ctx context.Context) ISmsRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsRecordDo
	WriteDB() ISmsRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsRecordDo
	Not(conds ...gen.Condition) ISmsRecordDo
	Or(conds ...gen.Condition) ISmsRecordDo
	Select(conds ...field.Expr) ISmsRecordDo
	Where(conds ...gen.Condition) ISmsRecordDo
	Order(conds ...field.Expr) ISmsRecordDo
	Distinct(cols ...field.Expr) ISmsRecordDo
	Omit(cols ...field.Expr) ISmsRecordDo
	Join(table schema.Tabler, on ...field.Expr) ISmsRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsRecordDo
	Group(cols ...field.Expr) ISmsRecordDo
	Having(conds ...gen.Condition) ISmsRecordDo
	Limit(limit int) ISmsRecordDo
	Offset(offset int) ISmsRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsRecordDo
	Unscoped() ISmsRecordDo
	Create(values ...*model.SmsRecord) error
	CreateInBatches(values []*model.SmsRecord, batchSize int) error
	Save(values ...*model.SmsRecord) error
	First() (*model.SmsRecord, error)
	Take() (*model.SmsRecord, error)
	Last() (*model.SmsRecord, error)
	Find() ([]*model.SmsRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsRecord, err error)
	FindInBatches(result *[]*model.SmsRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsRecordDo
	Assign(attrs ...field.AssignExpr) ISmsRecordDo
	Joins(fields ...field.RelationField) ISmsRecordDo
	Preload(fields ...field.RelationField) ISmsRecordDo
	FirstOrInit() (*model.SmsRecord, error)
	FirstOrCreate() (*model.SmsRecord, error)
	FindByPage(offset int, limit int) (result []*model.SmsRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsRecordDo) Debug() ISmsRecordDo {
	return s.withDO(s.DO.Debug())
}

func (s smsRecordDo) WithContext(ctx context.Context) ISmsRecordDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsRecordDo) ReadDB() ISmsRecordDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsRecordDo) WriteDB() ISmsRecordDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsRecordDo) Session(config *gorm.Session) ISmsRecordDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsRecordDo) Clauses(conds ...clause.Expression) ISmsRecordDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsRecordDo) Returning(value interface{}, columns ...string) ISmsRecordDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsRecordDo) Not(conds ...gen.Condition) ISmsRecordDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsRecordDo) Or(conds ...gen.Condition) ISmsRecordDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsRecordDo) Select(conds ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsRecordDo) Where(conds ...gen.Condition) ISmsRecordDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsRecordDo) Order(conds ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsRecordDo) Distinct(cols ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsRecordDo) Omit(cols ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsRecordDo) Join(table schema.Tabler, on ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsRecordDo) Group(cols ...field.Expr) ISmsRecordDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsRecordDo) Having(conds ...gen.Condition) ISmsRecordDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsRecordDo) Limit(limit int) ISmsRecordDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsRecordDo) Offset(offset int) ISmsRecordDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsRecordDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsRecordDo) Unscoped() ISmsRecordDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsRecordDo) Create(values ...*model.SmsRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsRecordDo) CreateInBatches(values []*model.SmsRecord, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsRecordDo) Save(values ...*model.SmsRecord) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsRecordDo) First() (*model.SmsRecord, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsRecord), nil
	}
}

func (s smsRecordDo) Take() (*model.SmsRecord, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsRecord), nil
	}
}

func (s smsRecordDo) Last() (*model.SmsRecord, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsRecord), nil
	}
}

func (s smsRecordDo) Find() ([]*model.SmsRecord, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsRecord), err
}

func (s smsRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsRecord, err error) {
	buf := make([]*model.SmsRecord, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsRecordDo) FindInBatches(result *[]*model.SmsRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsRecordDo) Attrs(attrs ...field.AssignExpr) ISmsRecordDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsRecordDo) Assign(attrs ...field.AssignExpr) ISmsRecordDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsRecordDo) Joins(fields ...field.RelationField) ISmsRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsRecordDo) Preload(fields ...field.RelationField) ISmsRecordDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsRecordDo) FirstOrInit() (*model.SmsRecord, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsRecord), nil
	}
}

func (s smsRecordDo) FirstOrCreate() (*model.SmsRecord, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsRecord), nil
	}
}

func (s smsRecordDo) FindByPage(offset int, limit int) (result []*model.SmsRecord, count int64, err error) {
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

func (s smsRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsRecordDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsRecordDo) Delete(models ...*model.SmsRecord) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsRecordDo) withDO(do gen.Dao) *smsRecordDo {
	s.DO = *do.(*gen.DO)
	return s
}