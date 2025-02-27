package model

import "time"

type ImageType string

const (
	Carousel   ImageType = "CAROUSEL"
	Process    ImageType = "PROCESS"
	ProcessGif ImageType = "PROCESS_GIF"
	Show3D     ImageType = "SHOW_3D"
	ShowGif    ImageType = "SHOW_GIF"
	Step       ImageType = "STEP"
)

type StorageType string

const (
	BlockMinio StorageType = "BLOCKMINIO"
	BlockYhs   StorageType = "BLOCKYHS"
)

type PictureStorage struct {
	ID          int       `gorm:"column:id;primaryKey"`
	ModelId     int       `gorm:"column:model_id;not null;comment:'模型表主键id'"`
	FileName    string    `gorm:"column:file_name;not null;comment:'文件名'"`
	ImagePath   string    `gorm:"column:image_path;not null;comment:'在minio的存储路径'"`
	BucketName  string    `gorm:"column:bucket_name;not null;comment:'在minio的存储的桶名'"`
	Url         string    `gorm:"column:url;comment:'存在公共桶的此自动才有值'"`
	ImageType   ImageType `gorm:"column:image_type;not null;comment:'图片类型'"`
	StorageType string    `gorm:"column:storage_type;not null;comment:'存储类型'"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}
