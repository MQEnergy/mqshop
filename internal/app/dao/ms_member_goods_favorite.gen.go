// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/MQEnergy/mqshop/internal/app/model"
)

func newMemberGoodsFavorite(db *gorm.DB, opts ...gen.DOOption) memberGoodsFavorite {
	_memberGoodsFavorite := memberGoodsFavorite{}

	_memberGoodsFavorite.memberGoodsFavoriteDo.UseDB(db, opts...)
	_memberGoodsFavorite.memberGoodsFavoriteDo.UseModel(&model.MemberGoodsFavorite{})

	tableName := _memberGoodsFavorite.memberGoodsFavoriteDo.TableName()
	_memberGoodsFavorite.ALL = field.NewAsterisk(tableName)
	_memberGoodsFavorite.ID = field.NewInt64(tableName, "id")
	_memberGoodsFavorite.MemberID = field.NewInt64(tableName, "member_id")
	_memberGoodsFavorite.GoodsID = field.NewInt64(tableName, "goods_id")
	_memberGoodsFavorite.CreatedAt = field.NewInt64(tableName, "created_at")

	_memberGoodsFavorite.fillFieldMap()

	return _memberGoodsFavorite
}

// memberGoodsFavorite 用户收藏商品表
type memberGoodsFavorite struct {
	memberGoodsFavoriteDo

	ALL       field.Asterisk
	ID        field.Int64
	MemberID  field.Int64 // 用户ID
	GoodsID   field.Int64 // 商品ID
	CreatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (m memberGoodsFavorite) Table(newTableName string) *memberGoodsFavorite {
	m.memberGoodsFavoriteDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m memberGoodsFavorite) As(alias string) *memberGoodsFavorite {
	m.memberGoodsFavoriteDo.DO = *(m.memberGoodsFavoriteDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *memberGoodsFavorite) updateTableName(table string) *memberGoodsFavorite {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.MemberID = field.NewInt64(table, "member_id")
	m.GoodsID = field.NewInt64(table, "goods_id")
	m.CreatedAt = field.NewInt64(table, "created_at")

	m.fillFieldMap()

	return m
}

func (m *memberGoodsFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *memberGoodsFavorite) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["id"] = m.ID
	m.fieldMap["member_id"] = m.MemberID
	m.fieldMap["goods_id"] = m.GoodsID
	m.fieldMap["created_at"] = m.CreatedAt
}

func (m memberGoodsFavorite) clone(db *gorm.DB) memberGoodsFavorite {
	m.memberGoodsFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m memberGoodsFavorite) replaceDB(db *gorm.DB) memberGoodsFavorite {
	m.memberGoodsFavoriteDo.ReplaceDB(db)
	return m
}

type memberGoodsFavoriteDo struct{ gen.DO }

type IMemberGoodsFavoriteDo interface {
	gen.SubQuery
	Debug() IMemberGoodsFavoriteDo
	WithContext(ctx context.Context) IMemberGoodsFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMemberGoodsFavoriteDo
	WriteDB() IMemberGoodsFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMemberGoodsFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMemberGoodsFavoriteDo
	Not(conds ...gen.Condition) IMemberGoodsFavoriteDo
	Or(conds ...gen.Condition) IMemberGoodsFavoriteDo
	Select(conds ...field.Expr) IMemberGoodsFavoriteDo
	Where(conds ...gen.Condition) IMemberGoodsFavoriteDo
	Order(conds ...field.Expr) IMemberGoodsFavoriteDo
	Distinct(cols ...field.Expr) IMemberGoodsFavoriteDo
	Omit(cols ...field.Expr) IMemberGoodsFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo
	Group(cols ...field.Expr) IMemberGoodsFavoriteDo
	Having(conds ...gen.Condition) IMemberGoodsFavoriteDo
	Limit(limit int) IMemberGoodsFavoriteDo
	Offset(offset int) IMemberGoodsFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberGoodsFavoriteDo
	Unscoped() IMemberGoodsFavoriteDo
	Create(values ...*model.MemberGoodsFavorite) error
	CreateInBatches(values []*model.MemberGoodsFavorite, batchSize int) error
	Save(values ...*model.MemberGoodsFavorite) error
	First() (*model.MemberGoodsFavorite, error)
	Take() (*model.MemberGoodsFavorite, error)
	Last() (*model.MemberGoodsFavorite, error)
	Find() ([]*model.MemberGoodsFavorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberGoodsFavorite, err error)
	FindInBatches(result *[]*model.MemberGoodsFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MemberGoodsFavorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMemberGoodsFavoriteDo
	Assign(attrs ...field.AssignExpr) IMemberGoodsFavoriteDo
	Joins(fields ...field.RelationField) IMemberGoodsFavoriteDo
	Preload(fields ...field.RelationField) IMemberGoodsFavoriteDo
	FirstOrInit() (*model.MemberGoodsFavorite, error)
	FirstOrCreate() (*model.MemberGoodsFavorite, error)
	FindByPage(offset int, limit int) (result []*model.MemberGoodsFavorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMemberGoodsFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m memberGoodsFavoriteDo) Debug() IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Debug())
}

func (m memberGoodsFavoriteDo) WithContext(ctx context.Context) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberGoodsFavoriteDo) ReadDB() IMemberGoodsFavoriteDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberGoodsFavoriteDo) WriteDB() IMemberGoodsFavoriteDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberGoodsFavoriteDo) Session(config *gorm.Session) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberGoodsFavoriteDo) Clauses(conds ...clause.Expression) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberGoodsFavoriteDo) Returning(value interface{}, columns ...string) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberGoodsFavoriteDo) Not(conds ...gen.Condition) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberGoodsFavoriteDo) Or(conds ...gen.Condition) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberGoodsFavoriteDo) Select(conds ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberGoodsFavoriteDo) Where(conds ...gen.Condition) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberGoodsFavoriteDo) Order(conds ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberGoodsFavoriteDo) Distinct(cols ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberGoodsFavoriteDo) Omit(cols ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberGoodsFavoriteDo) Join(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberGoodsFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberGoodsFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberGoodsFavoriteDo) Group(cols ...field.Expr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberGoodsFavoriteDo) Having(conds ...gen.Condition) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberGoodsFavoriteDo) Limit(limit int) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberGoodsFavoriteDo) Offset(offset int) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberGoodsFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberGoodsFavoriteDo) Unscoped() IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberGoodsFavoriteDo) Create(values ...*model.MemberGoodsFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberGoodsFavoriteDo) CreateInBatches(values []*model.MemberGoodsFavorite, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberGoodsFavoriteDo) Save(values ...*model.MemberGoodsFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberGoodsFavoriteDo) First() (*model.MemberGoodsFavorite, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberGoodsFavorite), nil
	}
}

func (m memberGoodsFavoriteDo) Take() (*model.MemberGoodsFavorite, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberGoodsFavorite), nil
	}
}

func (m memberGoodsFavoriteDo) Last() (*model.MemberGoodsFavorite, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberGoodsFavorite), nil
	}
}

func (m memberGoodsFavoriteDo) Find() ([]*model.MemberGoodsFavorite, error) {
	result, err := m.DO.Find()
	return result.([]*model.MemberGoodsFavorite), err
}

func (m memberGoodsFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberGoodsFavorite, err error) {
	buf := make([]*model.MemberGoodsFavorite, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberGoodsFavoriteDo) FindInBatches(result *[]*model.MemberGoodsFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberGoodsFavoriteDo) Attrs(attrs ...field.AssignExpr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberGoodsFavoriteDo) Assign(attrs ...field.AssignExpr) IMemberGoodsFavoriteDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberGoodsFavoriteDo) Joins(fields ...field.RelationField) IMemberGoodsFavoriteDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberGoodsFavoriteDo) Preload(fields ...field.RelationField) IMemberGoodsFavoriteDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberGoodsFavoriteDo) FirstOrInit() (*model.MemberGoodsFavorite, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberGoodsFavorite), nil
	}
}

func (m memberGoodsFavoriteDo) FirstOrCreate() (*model.MemberGoodsFavorite, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberGoodsFavorite), nil
	}
}

func (m memberGoodsFavoriteDo) FindByPage(offset int, limit int) (result []*model.MemberGoodsFavorite, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberGoodsFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberGoodsFavoriteDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberGoodsFavoriteDo) Delete(models ...*model.MemberGoodsFavorite) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberGoodsFavoriteDo) withDO(do gen.Dao) *memberGoodsFavoriteDo {
	m.DO = *do.(*gen.DO)
	return m
}