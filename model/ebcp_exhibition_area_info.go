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
[ 0] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] exhibition_start_time                          TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] exhibition_end_time                            TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] hall_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] hall_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] room_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[11] total_room_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[12] total_item_count                               INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[13] room_item_count                                INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "exhibition_id": "vivmIXeBCpuhxAjeqCoVeBrat",    "exhibition_name": "EnwCgikWmoUxoxGUFHgOCHhPu",    "exhibition_start_time": 86,    "exhibition_end_time": 93,    "hall_id": "WpNUnIFUvwYpbrDTjOyDUVCLC",    "hall_name": "lHDAQSXbFDnFaaNTBwKqqNIMq",    "room_id": "redIdDJYuIxNseAihTdJwmHPh",    "room_name": "oruWbAfdDuALkfmaBUFoNGwhD",    "room_floor": "yBonFKIWNBoogaOgBtjFHaMAU",    "room_location": "YwUHQbZhabNTXFJfBlMUAavrZ",    "room_status": 44,    "total_room_count": 53,    "total_item_count": 70,    "room_item_count": 24}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_area_info does not have a primary key defined, setting col position 1 exhibition_id as primary key
Warning table: v_ebcp_exhibition_area_info primary key column exhibition_id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_start_time = "exhibition_start_time"

	Ebcp_exhibition_area_info_FIELD_NAME_exhibition_end_time = "exhibition_end_time"

	Ebcp_exhibition_area_info_FIELD_NAME_hall_id = "hall_id"

	Ebcp_exhibition_area_info_FIELD_NAME_hall_name = "hall_name"

	Ebcp_exhibition_area_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_area_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_area_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_area_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_area_info_FIELD_NAME_room_status = "room_status"

	Ebcp_exhibition_area_info_FIELD_NAME_total_room_count = "total_room_count"

	Ebcp_exhibition_area_info_FIELD_NAME_total_item_count = "total_item_count"

	Ebcp_exhibition_area_info_FIELD_NAME_room_item_count = "room_item_count"
)

// Ebcp_exhibition_area_info struct is a row record of the v_ebcp_exhibition_area_info table in the  database
type Ebcp_exhibition_area_info struct {
	ExhibitionID string `json:"exhibition_id"` //展览ID

	ExhibitionName string `json:"exhibition_name"` //展览名称

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //展览开始时间

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //展览结束时间

	HallID string `json:"hall_id"` //展馆ID

	HallName string `json:"hall_name"` //展馆名称

	RoomID string `json:"room_id"` //展厅ID

	RoomName string `json:"room_name"` //展厅名称

	RoomFloor string `json:"room_floor"` //展厅楼层

	RoomLocation string `json:"room_location"` //展厅位置

	RoomStatus int32 `json:"room_status"` //展厅状态

	TotalRoomCount int32 `json:"total_room_count"` //展厅总数

	TotalItemCount int32 `json:"total_item_count"` //展项总数

	RoomItemCount int32 `json:"room_item_count"` //展厅展项总数

}

var Ebcp_exhibition_area_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_area_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "exhibition_id",
			Comment: `展览ID`,
			Notes: `Warning table: v_ebcp_exhibition_area_info does not have a primary key defined, setting col position 1 exhibition_id as primary key
Warning table: v_ebcp_exhibition_area_info primary key column exhibition_id is nullable column, setting it as NOT NULL
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
			Name:               "hall_id",
			Comment:            `展馆ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "HallID",
			GoFieldType:        "string",
			JSONFieldName:      "hall_id",
			ProtobufFieldName:  "hall_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "hall_name",
			Comment:            `展馆名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "HallName",
			GoFieldType:        "string",
			JSONFieldName:      "hall_name",
			ProtobufFieldName:  "hall_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "room_id",
			Comment:            `展厅ID`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "room_item_count",
			Comment:            `展厅展项总数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "RoomItemCount",
			GoFieldType:        "int32",
			JSONFieldName:      "room_item_count",
			ProtobufFieldName:  "room_item_count",
			ProtobufType:       "int32",
			ProtobufPos:        14,
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
