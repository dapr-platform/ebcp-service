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


Table: o_ebcp_item_device_relation
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 5] exhibition_item_id                             VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] device_type                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] device_sub_type                                VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 8] device_id                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []


JSON Sample
-------------------------------------
{    "id": "ycACpFgktKJIHPSdvRrBWCPxS",    "created_by": "GdkbLVKnfhSiZJvcvVfUHBSXZ",    "created_time": 61,    "updated_by": "yCxNuyndvKGBPvZOKQjjKetRC",    "updated_time": 83,    "exhibition_item_id": "DUHgkAXXUKWeQaSOvSciMCEoE",    "device_type": 39,    "device_sub_type": "rUbZkyKgFAnWbfdcoqvBlSMZJ",    "device_id": "HVIxwphqkhXgahkFeMPTwmfZx"}



*/

var (
	Ebcp_item_device_relation_FIELD_NAME_id = "id"

	Ebcp_item_device_relation_FIELD_NAME_created_by = "created_by"

	Ebcp_item_device_relation_FIELD_NAME_created_time = "created_time"

	Ebcp_item_device_relation_FIELD_NAME_updated_by = "updated_by"

	Ebcp_item_device_relation_FIELD_NAME_updated_time = "updated_time"

	Ebcp_item_device_relation_FIELD_NAME_exhibition_item_id = "exhibition_item_id"

	Ebcp_item_device_relation_FIELD_NAME_device_type = "device_type"

	Ebcp_item_device_relation_FIELD_NAME_device_sub_type = "device_sub_type"

	Ebcp_item_device_relation_FIELD_NAME_device_id = "device_id"
)

// Ebcp_item_device_relation struct is a row record of the o_ebcp_item_device_relation table in the  database
type Ebcp_item_device_relation struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	ExhibitionItemID string `json:"exhibition_item_id"` //展项ID

	DeviceType int32 `json:"device_type"` //关联设备类型(1: 播放设备, 2: 中控设备)

	DeviceSubType string `json:"device_sub_type"` //关联设备子类型(中控设备时需要)

	DeviceID string `json:"device_id"` //关联设备ID

}

var Ebcp_item_device_relationTableInfo = &TableInfo{
	Name: "o_ebcp_item_device_relation",
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
			Name:               "exhibition_item_id",
			Comment:            `展项ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionItemID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_item_id",
			ProtobufFieldName:  "exhibition_item_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "device_type",
			Comment:            `关联设备类型(1: 播放设备, 2: 中控设备)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "DeviceType",
			GoFieldType:        "int32",
			JSONFieldName:      "device_type",
			ProtobufFieldName:  "device_type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "device_sub_type",
			Comment:            `关联设备子类型(中控设备时需要)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "DeviceSubType",
			GoFieldType:        "string",
			JSONFieldName:      "device_sub_type",
			ProtobufFieldName:  "device_sub_type",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "device_id",
			Comment:            `关联设备ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "DeviceID",
			GoFieldType:        "string",
			JSONFieldName:      "device_id",
			ProtobufFieldName:  "device_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_item_device_relation) TableName() string {
	return "o_ebcp_item_device_relation"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_item_device_relation) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_item_device_relation) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_item_device_relation) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_item_device_relation) TableInfo() *TableInfo {
	return Ebcp_item_device_relationTableInfo
}
