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
[ 0] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] exhibition_start_time                          TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] exhibition_end_time                            TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] exhibition_status                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] total_room_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[ 6] total_item_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "exhibition_id": "JRtoGMotxBYfuDZyOReIHrYUC",    "exhibition_name": "ooZCDesSYmsupvcwOZyGYxWNu",    "exhibition_start_time": 20,    "exhibition_end_time": 64,    "exhibition_status": 0,    "total_room_count": 94,    "total_item_count": 93}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_info does not have a primary key defined, setting col position 1 exhibition_id as primary key
Warning table: v_ebcp_exhibition_info primary key column exhibition_id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_info_FIELD_NAME_exhibition_start_time = "exhibition_start_time"

	Ebcp_exhibition_info_FIELD_NAME_exhibition_end_time = "exhibition_end_time"

	Ebcp_exhibition_info_FIELD_NAME_exhibition_status = "exhibition_status"

	Ebcp_exhibition_info_FIELD_NAME_total_room_count = "total_room_count"

	Ebcp_exhibition_info_FIELD_NAME_total_item_count = "total_item_count"
)

// Ebcp_exhibition_info struct is a row record of the v_ebcp_exhibition_info table in the  database
type Ebcp_exhibition_info struct {
	ExhibitionID string `json:"exhibition_id"` //展览ID

	ExhibitionName string `json:"exhibition_name"` //展览名称

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //展览开始时间

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //展览结束时间

	ExhibitionStatus int32 `json:"exhibition_status"` //展览状态（1: 运行中, 2: 筹备中, 3: 已结束）

	TotalRoomCount int32 `json:"total_room_count"` //展厅总数

	TotalItemCount int32 `json:"total_item_count"` //展项总数

}

var Ebcp_exhibition_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "exhibition_id",
			Comment: `展览ID`,
			Notes: `Warning table: v_ebcp_exhibition_info does not have a primary key defined, setting col position 1 exhibition_id as primary key
Warning table: v_ebcp_exhibition_info primary key column exhibition_id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_id",
			ProtobufFieldName:  "exhibition_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "exhibition_name",
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
			GoFieldName:        "ExhibitionName",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_name",
			ProtobufFieldName:  "exhibition_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "exhibition_start_time",
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
			GoFieldName:        "ExhibitionStartTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "exhibition_start_time",
			ProtobufFieldName:  "exhibition_start_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "exhibition_end_time",
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
			GoFieldName:        "ExhibitionEndTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "exhibition_end_time",
			ProtobufFieldName:  "exhibition_end_time",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "exhibition_status",
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
			GoFieldName:        "ExhibitionStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "exhibition_status",
			ProtobufFieldName:  "exhibition_status",
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
