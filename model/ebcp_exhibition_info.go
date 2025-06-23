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


Table: v_ebcp_exhibition_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] start_time                                     TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] end_time                                       TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] total_room_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] total_item_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 7] rooms                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[ 8] items                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "yosiawnyUdhlHnmrhNowmpnFP",    "name": "oFDmdqbcHOhcoeubAnPlwxcrp",    "start_time": 77,    "end_time": 46,    "status": 59,    "total_room_count": 67,    "total_item_count": 83,    "rooms": 68,    "items": 89}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_info_FIELD_NAME_id = "id"

	Ebcp_exhibition_info_FIELD_NAME_name = "name"

	Ebcp_exhibition_info_FIELD_NAME_start_time = "start_time"

	Ebcp_exhibition_info_FIELD_NAME_end_time = "end_time"

	Ebcp_exhibition_info_FIELD_NAME_status = "status"

	Ebcp_exhibition_info_FIELD_NAME_total_room_count = "total_room_count"

	Ebcp_exhibition_info_FIELD_NAME_total_item_count = "total_item_count"

	Ebcp_exhibition_info_FIELD_NAME_rooms = "rooms"

	Ebcp_exhibition_info_FIELD_NAME_items = "items"
)

// Ebcp_exhibition_info struct is a row record of the v_ebcp_exhibition_info table in the  database
type Ebcp_exhibition_info struct {
	ID string `json:"id"` //展览ID

	Name string `json:"name"` //展览名称

	StartTime common.LocalTime `json:"start_time"` //展览开始时间

	EndTime common.LocalTime `json:"end_time"` //展览结束时间

	Status int32 `json:"status"` //展览状态（1: 运行中, 2: 筹备中, 3: 已结束）

	TotalRoomCount int32 `json:"total_room_count"` //展厅总数

	TotalItemCount int32 `json:"total_item_count"` //展项总数

	Rooms any `json:"rooms"` //展览使用的所有展厅

	Items any `json:"items"` //展览的所有展项

}

var Ebcp_exhibition_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `展览ID`,
			Notes: `Warning table: v_ebcp_exhibition_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_info primary key column id is nullable column, setting it as NOT NULL
`,
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
			Name:               "name",
			Comment:            `展览名称`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "start_time",
			Comment:            `展览开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "end_time",
			Comment:            `展览结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "status",
			Comment:            `展览状态（1: 运行中, 2: 筹备中, 3: 已结束）`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "total_room_count",
			Comment:            `展厅总数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "TotalRoomCount",
			GoFieldType:        "int32",
			JSONFieldName:      "total_room_count",
			ProtobufFieldName:  "total_room_count",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "total_item_count",
			Comment:            `展项总数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "TotalItemCount",
			GoFieldType:        "int32",
			JSONFieldName:      "total_item_count",
			ProtobufFieldName:  "total_item_count",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "rooms",
			Comment:            `展览使用的所有展厅`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Rooms",
			GoFieldType:        "any",
			JSONFieldName:      "rooms",
			ProtobufFieldName:  "rooms",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "items",
			Comment:            `展览的所有展项`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Items",
			GoFieldType:        "any",
			JSONFieldName:      "items",
			ProtobufFieldName:  "items",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_info) TableName() string {
	return "v_ebcp_exhibition_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_infoTableInfo
}
