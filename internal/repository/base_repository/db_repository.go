package base_repository

import (
	"context"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/utils/izap"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
)

// DBRepository 表示数据访问对象【mysql】
type DBRepository struct {
	//Model   model.DBModelImpl
	DB      *gorm.DB
	tx      map[string]*gorm.DB
	isNew   bool
	Name    string
	locked  int
	IsMyCat bool
	Gen     *gen.Query
}

func newRepository(r *DBRepository, txm map[string]*gorm.DB) *DBRepository {
	return &DBRepository{DB: r.DB, tx: txm}
}

func NewRepository(db *gorm.DB, gen *gen.Query) *DBRepository {
	return &DBRepository{DB: db, Gen: gen}
}

func (r *DBRepository) NewDB(txm map[string]*gorm.DB) *DBRepository {
	r.isNew = true
	return newRepository(r, txm)
}

// LockForUpdate 更新锁 其他的都不能读写，只能lockForUpdate 更新完再读写。
func (r *DBRepository) LockForUpdate() *DBRepository {
	r.locked = 1
	return r
}

// SharedLock 分享锁 别的事务可以读，但不能写，只能sharedLock 更新完再写
func (r *DBRepository) SharedLock() *DBRepository {
	r.locked = 2
	return r
}

func (r *DBRepository) db() *gorm.DB {
	//if r.tx != nil && r.tx[r.Model.Connection()] != nil {
	//	return r.tx[r.Model.Connection()].Table(r.Model.Table())
	//}
	return r.DB
}

// QueryAll 查询多条数据
func (r *DBRepository) QueryAll(ctx context.Context, condition string, args []interface{}, dest interface{}, options ...map[string]interface{}) error {
	query := r.QueryBuilder(ctx, condition, args, options...)
	return query.Find(dest).Error
}

// QueryAll 查询多条数据
func (r *DBRepository) QueryAll2(ctx context.Context, condition string, args []interface{}, dest interface{}, options ...map[string]interface{}) error {
	query := r.QueryBuilder(ctx, condition, args, options...)
	return query.WithContext(ctx).Find(dest).Error
}

// QueryOne 查询单条数据
func (r *DBRepository) QueryOne(ctx context.Context, condition string, args []interface{}, dest interface{}, options ...map[string]interface{}) error {
	query := r.QueryBuilder(ctx, condition, args, options...)
	var err error
	//判断options是否有orderBy字段，如果有则不使用First方法,First方法会自动加上主键排序
	if len(options) != 0 && len(options[0]) != 0 {
		option := options[0]
		if orderBy, ok := option["order_by"].(string); ok && orderBy != "" {
			err = query.Take(dest).Limit(1).Error
		} else {
			err = query.First(dest).Error
		}
	} else if r.IsMyCat {
		err = query.Take(dest).Limit(1).Error
	} else {
		err = query.First(dest).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

// Count 查询数据条数
func (r *DBRepository) Count(ctx context.Context, condition string, args []interface{}, options ...map[string]interface{}) (int64, error) {
	query := r.QueryBuilder(ctx, condition, args, options...)

	var count int64
	err := query.Count(&count).Error

	return count, err
}

// Exists 判断数据是否存在
func (r *DBRepository) Exists(ctx context.Context, condition interface{}, args []interface{}) (bool, error) {
	result := make(map[string]interface{})
	err := r.db().WithContext(ctx).Where(condition, args...).Take(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if len(result) == 0 {
		return false, err
	}
	return true, err
}

// Create 创建数据
func (r *DBRepository) Create(ctx context.Context, value interface{}) error {
	return r.db().WithContext(ctx).Create(value).Error
}

// CreateOrUpdate 创建或更新数据
func (r *DBRepository) CreateOrUpdate(ctx context.Context, values interface{}, conflict clause.OnConflict) error {
	query := r.db()
	query.Clauses(conflict)
	return query.WithContext(ctx).Create(values).Error
}

// Update 更新数据
func (r *DBRepository) Update(ctx context.Context, condition interface{}, args []interface{}, values interface{}) error {
	return r.db().WithContext(ctx).Where(condition, args...).Updates(values).Error
}

// Delete 删除数据
func (r *DBRepository) Delete(ctx context.Context, condition interface{}, conds []interface{}) error {
	var values interface{}
	return r.db().WithContext(ctx).Where(condition, conds...).Delete(values).Error
}

// QueryBuilder 构建查询条件
func (r *DBRepository) QueryBuilder(ctx context.Context, condition string, args []interface{}, options ...map[string]interface{}) *gorm.DB {
	query := r.db().Where(condition, args...)
	if len(options) != 0 && len(options[0]) != 0 {
		option := options[0]

		//表名别名设置
		if tableAlias, ok := option["alias"].(string); ok && tableAlias != "" {
			query = query.Table(tableAlias)
		}

		//连表查询
		if join, ok := option["join"]; ok && join != "" {
			//判断是否是数组
			if joinArr, ok := join.([]string); ok {
				for _, j := range joinArr {
					query = query.Joins(j)
				}
			} else {
				query = query.Joins(join.(string))
			}
		}

		//指定字段
		if selectField, ok := option["select"].(string); ok && selectField != "" {
			query = query.Select(selectField)
		}

		var limit int
		if pageSize, ok := option["page_size"].(int); ok && pageSize > 0 {
			limit = pageSize
			query = query.Limit(limit)
		}

		if page, ok := option["page"].(int); ok && page > 0 {
			query = query.Offset((page - 1) * limit)
		}

		// 处理排序参数
		if orderBy, ok := option["order_by"].(string); ok && orderBy != "" {
			query = query.Order(orderBy)
		}

		// 处理分组参数
		if groupBy, ok := option["group_by"].(string); ok && groupBy != "" {
			query = query.Group(groupBy)
		}
	}

	if r.isNew && r.locked == 1 {
		query.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	if r.isNew && r.locked == 2 {
		query.Clauses(clause.Locking{
			Strength: "SHARE",
			Table:    clause.Table{Name: clause.CurrentTable},
		})
	}

	return query.WithContext(ctx)
}

// ExecSql 执行原生sql
func (r *DBRepository) ExecSql(ctx context.Context, sql string, args ...interface{}) error {
	return r.db().WithContext(ctx).Exec(sql, args...).Error
}

// ScanStruct 执行原生sql并将结果映射到结构体
func (r *DBRepository) ScanStruct(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	return r.db().WithContext(ctx).Raw(sql, args...).Scan(dest).Error
}

// GetFullTableName 获取数据表全称，用于跨库连接查询
//func (r *DBRepository) GetFullTableName() string {
//	conn := r.Model.Connection()
//	// default 转换为 zulin
//	if conn == "default" {
//		conn = "zulin"
//	}
//	table := r.Model.Table()
//	return fmt.Sprintf("%s.%s", conn, table)
//}

func (r *DBRepository) Transaction(ctx context.Context, fn func(query *gen.Query) error) (err error) {
	return r.db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var (
			flag error
			wg   sync.WaitGroup
		)

		done := make(chan struct{})
		wg.Add(1)
		go func() {
			defer wg.Done()
			flag = fn(gen.Use(tx))
		}()

		go func() {
			defer close(done)
			wg.Wait()
		}()

		select {
		case <-ctx.Done():
			izap.Log.Error("Transaction Rollback due to timeout", zap.Error(ctx.Err()))

			return ctx.Err()
		case <-done:
			return flag
		}
	})
}
