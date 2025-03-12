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


Table: o_ebcp_exhibition_item
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 5] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] exhibition_id                                  VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] room_id                                        VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] type                                           VARCHAR(50)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 9] export_info                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]
[11] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "tQSnIqsGWQcATFgKOKnvfKhVu",    "created_by": "DGAABIoJqXYwMscrpCXnjhhaK",    "created_time": 59,    "updated_by": "GJxKJfhYaDRiqIeamGBtBUmfu",    "updated_time": 45,    "name": "tIbqYfkfymGMBFCpHcOYwkCcv",    "exhibition_id": "PaWGXApHaYOkSYRHcCdEhJEnE",    "room_id": "SMNJnIexlUQJGTOpWCEoJkjCb",    "type": "rIsIyQCafipsJnSvupCOWfiTO",    "export_info": "cEsslJIKjsOArrMsagFZelsfQ",    "status": 68,    "remarks": "apvNEMrtZlFHNfGmHBObqkSbN"}



*/

var (
	Ebcp_exhibition_item_FIELD_NAME_id = "id"

	Ebcp_exhibition_item_FIELD_NAME_created_by = "created_by"

	Ebcp_exhibition_item_FIELD_NAME_created_time = "created_time"

	Ebcp_exhibition_item_FIELD_NAME_updated_by = "updated_by"

	Ebcp_exhibition_item_FIELD_NAME_updated_time = "updated_time"

	Ebcp_exhibition_item_FIELD_NAME_name = "name"

	Ebcp_exhibition_item_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_item_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_item_FIELD_NAME_type = "type"

	Ebcp_exhibition_item_FIELD_NAME_export_info = "export_info"

	Ebcp_exhibition_item_FIELD_NAME_status = "status"

	Ebcp_exhibition_item_FIELD_NAME_remarks = "remarks"
)

// Ebcp_exhibition_item struct is a row record of the o_ebcp_exhibition_item table in the  database
type Ebcp_exhibition_item struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	Name string `json:"name"` //展项名称

	ExhibitionID string `json:"exhibition_id"` //所属展览ID

	RoomID string `json:"room_id"` //所属展厅ID

	Type string `json:"type"` //展项类型（media、static）

	ExportInfo string `json:"export_info"` //输出信息

	Status int32 `json:"status"` //状态（1: 启动, 2: 停止, 3: 故障）

	Remarks string `json:"remarks"` //备注

}

var Ebcp_exhibition_itemTableInfo = &TableInfo{
	Name: "o_ebcp_exhibition_item",
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
			Comment:            `展项名称`,
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
			Name:               "exhibition_id",
			Comment:            `所属展览ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_id",
			ProtobufFieldName:  "exhibition_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "room_id",
			Comment:            `所属展厅ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RoomID",
			GoFieldType:        "string",
			JSONFieldName:      "room_id",
			ProtobufFieldName:  "room_id",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "type",
			Comment:            `展项类型（media、static）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "export_info",
			Comment:            `输出信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ExportInfo",
			GoFieldType:        "string",
			JSONFieldName:      "export_info",
			ProtobufFieldName:  "export_info",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "status",
			Comment:            `状态（1: 启动, 2: 停止, 3: 故障）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "remarks",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Remarks",
			GoFieldType:        "string",
			JSONFieldName:      "remarks",
			ProtobufFieldName:  "remarks",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_item) TableName() string {
	return "o_ebcp_exhibition_item"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_item) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_item) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_item) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_item) TableInfo() *TableInfo {
	return Ebcp_exhibition_itemTableInfo
}
