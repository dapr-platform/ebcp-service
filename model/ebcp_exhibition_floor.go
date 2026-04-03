package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_ebcp_exhibition_floor
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 5] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] short_name                                     VARCHAR(10)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 10      default: []
[ 7] sort_order                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 8] background                                     TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 9] monitor                                        TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "gTjIaMKQMtCLBRYTEcbkLDwMX",    "created_by": "bwAZOELNZhaYUrhivdcVGRtKK",    "created_time": 1,    "updated_by": "RRXgrojxiDIwGLcXDFdYcgYcx",    "updated_time": 36,    "name": "tjPOVVLnGjTalpeBTNhLVydLy",    "short_name": "kSLXuCyItxZXwvdfGPacWdlql",    "sort_order": 5,    "background": "BGlnlTjQfWIJbYXXtHjAivWki",    "monitor": "XGjISbppjCMhdFNVbdUtHmFQq"}



*/

var (
	Ebcp_exhibition_floor_FIELD_NAME_id = "id"

	Ebcp_exhibition_floor_FIELD_NAME_created_by = "created_by"

	Ebcp_exhibition_floor_FIELD_NAME_created_time = "created_time"

	Ebcp_exhibition_floor_FIELD_NAME_updated_by = "updated_by"

	Ebcp_exhibition_floor_FIELD_NAME_updated_time = "updated_time"

	Ebcp_exhibition_floor_FIELD_NAME_name = "name"

	Ebcp_exhibition_floor_FIELD_NAME_short_name = "short_name"

	Ebcp_exhibition_floor_FIELD_NAME_sort_order = "sort_order"

	Ebcp_exhibition_floor_FIELD_NAME_background = "background"

	Ebcp_exhibition_floor_FIELD_NAME_monitor = "monitor"
)

// Ebcp_exhibition_floor struct is a row record of the o_ebcp_exhibition_floor table in the  database
type Ebcp_exhibition_floor struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	Name string `json:"name"` //楼层名称

	ShortName string `json:"short_name"` //楼层简称

	SortOrder int32 `json:"sort_order"` //排序

	Background string `json:"background"` //背景图

	Monitor string `json:"monitor"` //监控信息

}

var Ebcp_exhibition_floorTableInfo = &TableInfo{
	Name: "o_ebcp_exhibition_floor",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "created_by",
			Comment:            `created_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CreatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "created_by",
			ProtobufFieldName:  "created_by",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "created_time",
			Comment:            `created_time`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "updated_by",
			Comment:            `updated_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_time",
			Comment:            `updated_time`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "name",
			Comment:            `楼层名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "short_name",
			Comment:            `楼层简称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       10,
			GoFieldName:        "ShortName",
			GoFieldType:        "string",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sort_order",
			Comment:            `排序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SortOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "sort_order",
			ProtobufFieldName:  "sort_order",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "background",
			Comment:            `背景图`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Background",
			GoFieldType:        "string",
			JSONFieldName:      "background",
			ProtobufFieldName:  "background",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "monitor",
			Comment:            `监控信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Monitor",
			GoFieldType:        "string",
			JSONFieldName:      "monitor",
			ProtobufFieldName:  "monitor",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_floor) TableName() string {
	return "o_ebcp_exhibition_floor"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_floor) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_floor) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_floor) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_floor) TableInfo() *TableInfo {
	return Ebcp_exhibition_floorTableInfo
}
