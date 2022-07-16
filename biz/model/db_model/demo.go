package db_model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type DemoStruct struct {
	Id       int64  `gorm:"id" json:"ID"`
	NAME     string `gorm:"name" json:"NAME"`
	AGE      int64  `gorm:"age" json:"AGE"`
	BIRTHDAY int64  `gorm:"birthday" json:"BIRTHDAY"`
}

func GetDemoRecord(ctx context.Context, db *gorm.DB, tableName string) (de *DemoStruct, err error) {
	// 查询单条数据示例
	err = db.WithContext(ctx).Debug().Table(tableName).First(&de).Limit(1).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	return de, err
}
