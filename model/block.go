package model

import (
	"github.com/lib/pq"
	"time"
)

type BlockModel struct {
	ID           int            `gorm:"column:id;type:integer;primaryKey"`
	Url          string         `gorm:"column:url;type:varchar(200);comment:'图片minio存储地址'"`
	Title        string         `gorm:"column:title;type:varchar(30);not null;comment:'模型名称'"`
	ModelType    int            `gorm:"column:model_type;type:integer;not null;comment:'模型类别对模型分类表主键id'"`
	ModelFile    string         `gorm:"column:model_file;type:varchar(200);comment:'图片minio存储地址'"`
	BucketName   string         `gorm:"column:bucket_name;type:varchar(20);comment:'图片minio存储桶名'"`
	Hot          int            `gorm:"column:hot;type:integer;default:0;comment:'热度'"`
	Tags         pq.StringArray `gorm:"column:tags;type:varchar(20)[];comment:'模型标签'"`
	Introduction string         `gorm:"column:introduction;type:text;comment:'模型描述信息'"`
	IsRecommend  bool           `gorm:"column:is_recommend;type:bool;default:false;comment:'是否是推荐模型'"`
	TotalCount   int            `gorm:"column:total_count;type:integer;comment:'积木总数量'"`
	IsShow       bool           `gorm:"column:is_show;type:bool;default:false;comment:'是否展示'"`
	CreateTime   time.Time      `gorm:"column:create_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime   time.Time      `gorm:"column:update_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}

func (BlockModel) TableName() string {
	return "block_model"
}

type BlockModelType struct {
	ID         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name;type:varchar(30);not null;uniqueIndex;comment:'类别名字'"`
	CreateTime string `gorm:"column:create_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime string `gorm:"column:update_time;type:timestamp(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}

func (BlockModelType) TableName() string {
	return "model_type"
}
