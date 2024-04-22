package database

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(id int, preloads ...string) error
	GetWhere(where map[string]interface{}, preloads ...string) error
	GetDeleted(where map[string]interface{}) error

	FindAllWhere(where map[string]interface{}, preloads ...string) error
	FindAllPaginating(filter Filter, paginator *Paginator, preloads ...string) error

	Create() error
	Save() error
	Delete(id int) error
	Update(fields interface{}, id int) error
}

type repositoryImpl struct {
	db     *gorm.DB
	entity interface{}
	ctx    *fiber.Ctx
}

func New(entity interface{}, ctx *fiber.Ctx) Repository {
	return &repositoryImpl{
		db:     database.GetDB(),
		entity: entity,
		ctx:    ctx,
	}
}

func (r *repositoryImpl) GetById(id int, preloads ...string) error {
	queryDb := r.checkPreloads(preloads...)
	queryDb = queryDb.First(r.entity, id)
	if queryDb.Error != nil {
		return queryDb.Error
	}

	return nil
}

func (r *repositoryImpl) GetWhere(where map[string]interface{}, preloads ...string) error {
	queryDb := r.checkPreloads(preloads...)
	queryDb = queryDb.Where(where).First(r.entity)
	if queryDb.Error != nil {
		return queryDb.Error
	}

	return nil
}

func (r *repositoryImpl) GetDeleted(where map[string]interface{}) error {
	queryDb := r.db
	queryDb = queryDb.Unscoped().Where(where).First(r.entity)
	if queryDb.Error != nil {
		return queryDb.Error
	}

	return nil
}

func (r *repositoryImpl) FindAllWhere(where map[string]interface{}, preloads ...string) error {
	queryDb := r.checkPreloads(preloads...)
	queryDb = queryDb.Where(where).Find(r.entity)

	if queryDb.Error != nil {
		return queryDb.Error
	}

	return nil
}

func (r *repositoryImpl) FindAllPaginating(filter Filter, paginator *Paginator, preloads ...string) error {
	query := r.checkPreloads(preloads...)
	query = filter.Apply(query)
	query = paginator.Paginate(query, r.entity)
	query = query.Find(r.entity)

	if query.Error != nil {
		return query.Error
	}

	if paginator.Rows != nil {
		jsonData, err := json.Marshal(r.entity)
		if err != nil {
			return err
		}

		err = json.Unmarshal(jsonData, paginator.Rows)
		if err != nil {
			return err
		}
	} else {
		paginator.Rows = r.entity
	}

	return nil
}

func (r *repositoryImpl) Create() error {
	if r.ctx.Locals("user") == nil {
		return fmt.Errorf("user operator not found, try auth again")
	}
	user := r.ctx.Locals("user").(*domain.User)

	queryDb := r.db
	queryDb = queryDb.Create(r.entity)

	if queryDb.Error != nil {
		return queryDb.Error
	}

	if queryDb.RowsAffected == 0 {
		return fmt.Errorf("rows affected equals zero")
	}

	rawEntity, err := json.Marshal(r.entity)
	if err != nil {
		return err
	}

	id, err := r.getEntityId()
	if err != nil {
		log.Println("failed to save the id of the entity created in audit table", string(rawEntity))
	}

	audit := domain.Audit{
		UserId:    int(user.ID),
		TableName: queryDb.Statement.Table,
		RowId:     int(id),
		Action:    "CREATE",
		JsonData:  string(rawEntity),
	}

	bigquery := bigquery.GetDB().Create(&audit)
	if bigquery.Error != nil {
		return bigquery.Error
	}

	return nil
}

func (r *repositoryImpl) Save() error {
	if r.ctx.Locals("user") == nil {
		return fmt.Errorf("user operator not found, try auth again")
	}
	user := r.ctx.Locals("user").(*domain.User)

	var action string

	id, err := r.getEntityId()
	if err != nil {
		entityStruct := reflect.ValueOf(r.entity)
		for entityStruct.Kind() == reflect.Ptr {
			entityStruct = entityStruct.Elem()
		}
		log.Println("failed to get the id of the entity", entityStruct)
	}

	if id == 0 {
		action = "INSERT"
	} else {
		action = "UPDATE"
	}

	queryDb := r.db
	queryDb = queryDb.Save(r.entity)
	if queryDb.Error != nil {
		return queryDb.Error
	}

	if queryDb.RowsAffected == 0 {
		return fmt.Errorf("rows affected equals zero")
	}

	rawEntity, err := json.Marshal(r.entity)
	if err != nil {
		return err
	}

	id, err = r.getEntityId()
	if err != nil {
		log.Println("failed to save the id of the entity created in audit table", string(rawEntity))
	}

	audit := domain.Audit{
		UserId:    int(user.ID),
		TableName: queryDb.Statement.Table,
		RowId:     int(id),
		Action:    action,
		JsonData:  string(rawEntity),
	}

	bigquery := bigquery.GetDB().Create(&audit)
	if bigquery.Error != nil {
		return bigquery.Error
	}

	return nil
}

func (r *repositoryImpl) Delete(id int) error {
	if r.ctx.Locals("user") == nil {
		return fmt.Errorf("user operator not found, try auth again")
	}
	user := r.ctx.Locals("user").(*domain.User)

	queryDb := r.db
	queryDb = queryDb.Delete(r.entity, id)
	if queryDb.Error != nil {
		return r.db.Error
	}
	if queryDb.RowsAffected == 0 {
		return fmt.Errorf("rows affected equals zero")
	}

	audit := domain.Audit{
		UserId:    int(user.ID),
		TableName: queryDb.Statement.Table,
		RowId:     id,
		Action:    "DELETE",
	}

	bigquery := bigquery.GetDB().Create(&audit)
	if bigquery.Error != nil {
		return bigquery.Error
	}

	return nil
}

func (r *repositoryImpl) Update(fields interface{}, id int) error {
	if r.ctx.Locals("user") == nil {
		return fmt.Errorf("user operator not found, try auth again")
	}
	user := r.ctx.Locals("user").(*domain.User)

	queryDb := r.db
	queryDb = queryDb.Model(r.entity).Where("id = ?", id).Updates(fields).Find(r.entity)

	if queryDb.Error != nil {
		return queryDb.Error
	}

	if queryDb.RowsAffected == 0 {
		return fmt.Errorf("rows affected equals zero")
	}

	rawEntity, err := json.Marshal(&fields)
	if err != nil {
		return err
	}

	audit := domain.Audit{
		UserId:    int(user.ID),
		TableName: queryDb.Statement.Table,
		RowId:     id,
		Action:    "UPDATE",
		JsonData:  string(rawEntity),
	}

	bigquery := bigquery.GetDB().Create(&audit)
	if bigquery.Error != nil {
		return bigquery.Error
	}

	return nil
}

func (r *repositoryImpl) checkPreloads(args ...string) *gorm.DB {
	if len(args) == 0 {
		return r.db
	}

	queryDb := r.db
	fmt.Printf("Query: %t", queryDb.Statement.Unscoped)
	for _, arg := range args {
		queryDb = queryDb.Preload(arg)
	}
	return queryDb
}

func (r *repositoryImpl) getEntityId() (uint, error) {
	entityStruct := reflect.ValueOf(r.entity)
	for entityStruct.Kind() == reflect.Ptr {
		entityStruct = entityStruct.Elem()
	}
	id, ok := entityStruct.FieldByName("ID").Interface().(uint)
	if !ok {
		return 0, fmt.Errorf("failed to get entity id")
	}
	return id, nil
}
