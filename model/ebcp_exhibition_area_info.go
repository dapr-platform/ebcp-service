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


Table: v_ebcp_exhibition_area_info
[ 0] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] room_floor_value                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] room_floor_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] room_location_value                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] room_location_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] room_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] room_remarks                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] exhibition_start_time                          TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[13] exhibition_end_time                            TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[14] exhibition_status                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] items                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "room_id": "kZaTumKfLxVCCjsWgBxDnuNky",    "room_name": "IWhMDwiRJgTbEjvvTgOgAFlvq",    "room_floor": "ViUrECZoJvRaatyKhjZDxZDbX",    "room_floor_value": "HQuJvNySgJMTVygjoSpBuGThg",    "room_floor_name": "rJCtUXBImihErjqgVGKNVgLwI",    "room_location": "GlmiNFFxcIQxWupDVsLBxrOMt",    "room_location_value": "KdekFAdMruOTOCHxDFAmroQRB",    "room_location_name": "bgYxiVeNRvICqvQYenrQmAVhj",    "room_status": 70,    "room_remarks": "mtmnqHYESIahAIoyVkwqcBdeC",    "exhibition_id": "ZUEflpxYJdPnuSGuodcirRLUJ",    "exhibition_name": "nNTnnfkKMFaOBhEgSjLtCAJoL",    "exhibition_start_time": 18,    "exhibition_end_time": 48,    "exhibition_status": 46,    "items": 16}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_area_info does not have a primary key defined, setting col position 1 room_id as primary key
Warning table: v_ebcp_exhibition_area_info primary key column room_id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_area_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_area_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_area_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_area_info_FIELD_NAME_room_floor_value = "room_floor_value"

	Ebcp_exhibition_area_info_FIELD_NAME_room_floor_name = "room_floor_name"

	Ebcp_exhibition_area_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_area_info_FIELD_NAME_room_location_value = "room_location_value"

	Ebcp_exhibition_area_info_FIELD_NAME_room_location_name = "room_location_name"

	Ebcp_exhibition_area_info_FIELD_NAME_room_status = "room_status"

	Ebcp_exhibition_area_info_FIELD_NAME_room_remarks = "room_remarks"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_start_time = "exhibition_start_time"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_end_time = "exhibition_end_time"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_status = "exhibition_status"

	Ebcp_exhibition_area_info_FIELD_NAME_items = "items"
)

// Ebcp_exhibition_area_info struct is a row record of the v_ebcp_exhibition_area_info table in the  database
type Ebcp_exhibition_area_info struct {
	RoomID string `json:"room_id"` //展厅ID

	RoomName string `json:"room_name"` //展厅名称

	RoomFloor string `json:"room_floor"` //展厅楼层

	RoomFloorValue string `json:"room_floor_value"` //展厅楼层值

	RoomFloorName string `json:"room_floor_name"` //展厅楼层名称

	RoomLocation string `json:"room_location"` //展厅位置

	RoomLocationValue string `json:"room_location_value"` //展厅位置值

	RoomLocationName string `json:"room_location_name"` //展厅位置名称

	RoomStatus int32 `json:"room_status"` //展厅状态

	RoomRemarks string `json:"room_remarks"` //展厅备注

	ExhibitionID string `json:"exhibition_id"` //展览ID

	ExhibitionName string `json:"exhibition_name"` //展览名称

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //展览开始时间

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //展览结束时间

	ExhibitionStatus int32 `json:"exhibition_status"` //展览状态

	Items any `json:"items"` //展厅内的展项列表

}

var Ebcp_exhibition_area_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_area_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "room_id",
			Comment: `展厅ID`,
			Notes: `Warning table: v_ebcp_exhibition_area_info does not have a primary key defined, setting col position 1 room_id as primary key
Warning table: v_ebcp_exhibition_area_info primary key column room_id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RoomID",
			GoFieldType:        "string",
			JSONFieldName:      "room_id",
			ProtobufFieldName:  "room_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "room_name",
			Comment:            `展厅名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomName",
			GoFieldType:        "string",
			JSONFieldName:      "room_name",
			ProtobufFieldName:  "room_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "room_floor",
			Comment:            `展厅楼层`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RoomFloor",
			GoFieldType:        "string",
			JSONFieldName:      "room_floor",
			ProtobufFieldName:  "room_floor",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "room_floor_value",
			Comment:            `展厅楼层值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomFloorValue",
			GoFieldType:        "string",
			JSONFieldName:      "room_floor_value",
			ProtobufFieldName:  "room_floor_value",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "room_floor_name",
			Comment:            `展厅楼层名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomFloorName",
			GoFieldType:        "string",
			JSONFieldName:      "room_floor_name",
			ProtobufFieldName:  "room_floor_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "room_location",
			Comment:            `展厅位置`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RoomLocation",
			GoFieldType:        "string",
			JSONFieldName:      "room_location",
			ProtobufFieldName:  "room_location",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "room_location_value",
			Comment:            `展厅位置值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomLocationValue",
			GoFieldType:        "string",
			JSONFieldName:      "room_location_value",
			ProtobufFieldName:  "room_location_value",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "room_location_name",
			Comment:            `展厅位置名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "RoomLocationName",
			GoFieldType:        "string",
			JSONFieldName:      "room_location_name",
			ProtobufFieldName:  "room_location_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "room_status",
			Comment:            `展厅状态`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "RoomStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "room_status",
			ProtobufFieldName:  "room_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "room_remarks",
			Comment:            `展厅备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "RoomRemarks",
			GoFieldType:        "string",
			JSONFieldName:      "room_remarks",
			ProtobufFieldName:  "room_remarks",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "exhibition_id",
			Comment:            `展览ID`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "exhibition_status",
			Comment:            `展览状态`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "items",
			Comment:            `展厅内的展项列表`,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_area_info) TableName() string {
	return "v_ebcp_exhibition_area_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_area_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_area_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_area_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_area_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_area_infoTableInfo
}
