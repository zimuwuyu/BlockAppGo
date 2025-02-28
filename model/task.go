package model

import (
	"time"
)

type TaskState string

const (
	locked   TaskState = "LOCKED"
	pending  TaskState = "PENDING"
	running  TaskState = "RUNNING"
	failed   TaskState = "FAILED"
	success  TaskState = "SUCCESS"
	canceled TaskState = "CANCELED"
)

type Task struct {
	ID             int                    `gorm:"column:id;primaryKey"`
	Name           string                 `gorm:"column:name;comment:'任务名称'"`
	ModelId        int                    `gorm:"column:model_id;int;not null;comment:'blockModel id'"`
	TaskState      TaskState              `gorm:"column:task_state;varchar(30);not null;comment:'任务状态'"`
	UserId         int                    `gorm:"column:user_id;int;not null;comment:'用户id'"`
	ModelHigh      float32                `gorm:"column:model_high;not null;comment:'模型高度, 指的是模型在最长方向上的高度'"`
	Step           float32                `gorm:"column:step;not null;default:0.2;comment:'积木堆叠步长'"`
	BlockSize      float32                `gorm:"column:block_size;not null;default:0.008;comment:'积木大小，长和宽，长款相同'"`
	CylinderRadius float32                `gorm:"column:cylinder_radius;not null;default:0.001625;comment:'积木块小圆柱体高度'"`
	Split          bool                   `gorm:"column:split;not null;default:false;comment:'是否每个积木块的绘制都添加到场景, 并更新场景, 否则统一绘制一层的积木并更新'"`
	Visible        bool                   `gorm:"column:visible;not null;default:false;comment:'是否显示绘制界面'"`
	AsZType        interface{}            `gorm:"column:as_z_type;type:json;comment:'AsZType字段，可为枚举值或二维数组'"`
	ZReverse       bool                   `gorm:"column:z_reverse;not null;default:false;comment:'是否反转z轴'"`
	UsePvVoxel     bool                   `gorm:"column:use_pv_voxel;not null;default:false;comment:'是否使用pyvista的voxelize方法, 否则使用open3d的VoxelGrid方法'"`
	UsePv3d        bool                   `gorm:"column:use_pv_3d;not null;default:false;comment:'是否使用pyvista绘制3D模型'"`
	WallThickness  float32                `gorm:"column:wall_thickness;comment:'壁厚度, 在这个厚度以内的积木会被删除, 若为None则为积木厚度的一半'"`
	ViewSize       []int                  `gorm:"column:view_size;type:text[];comment:'流程图尺寸，显示窗口大小'"`
	Version        string                 `gorm:"column:version;comment:'task版本信息'"`
	Extend         map[string]interface{} `gorm:"column:extend;type:json;comment:'扩展字段'" ` // 使用 map 存储动态 JSON 数据
	CreateTime     time.Time              `gorm:"column:create_time;type:datetime(6);default:CURRENT_TIMESTAMP(6)"`
	UpdateTime     time.Time              `gorm:"column:update_time;type:datetime(6);default:CURRENT_TIMESTAMP(6);on_update:CURRENT_TIMESTAMP(6)"`
}

type TaskLog struct {
	ID         int       `gorm:"column:id;primaryKey"`
	TaskId     int       `gorm:"column:task_id;not null;comment:'任务id'"`
	LogType    string    `gorm:"column:log_type;default:debug;comment:'日志等级'"`
	LogTime    string    `gorm:"column:log_time;type:datetime(6);comment:'任务日志时间'"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime(6);default:CURRENT_TIMESTAMP(6)"`
}
