package service

import (
	"context"
	"fmt"
	"github.com/aman-bansal/approval-queue/config"
	"github.com/aman-bansal/approval-queue/internal/models"
	"github.com/jinzhu/gorm"
)

type StorageService interface {

}

type MysqlStorageService struct {
	db *gorm.DB
}

func NewStorageService(config config.ApprovalQueueConfig) (StorageService, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True",
		config.MysqlUser, config.MysqlPass, config.MysqlHost, config.MysqlDb))
	if err != nil {
		return nil, err
	}
	return &MysqlStorageService{
		db: db,
	}, nil
}

type DbObject interface{}

func (m *MysqlStorageService) StartTransaction(ctx context.Context) *gorm.DB {
	return m.db.Begin()
}

func (m *MysqlStorageService) RollbackTransaction(ctx context.Context) *gorm.DB {
	return m.db.Rollback()
}

func (m *MysqlStorageService) CommitTransaction(ctx context.Context) *gorm.DB {
	return m.db.Commit()
}

func (m *MysqlStorageService) Create(ctx context.Context, obj DbObject) error {
	create := m.db.Create(obj)
	return create.Error
}

func (m *MysqlStorageService) Update(ctx context.Context, obj DbObject) error {
	update := m.db.Save(obj)
	return update.Error
}

func (m *MysqlStorageService) Get(ctx context.Context, queueId int64) (models.ApprovalQueue, error) {
	queue := models.ApprovalQueue{}
	first := m.db.Where("id= ?", queueId).First(&queue)
	return queue, first.Error
}

func (m *MysqlStorageService) Search(ctx context.Context) ([]models.ApprovalQueue, error) {
	queue := make([]models.ApprovalQueue, 0)
	//to write query
	find := m.db.Where("").Find(queue)
	return queue, find.Error
}