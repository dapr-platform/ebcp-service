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


Table: v_ebcp_exhibition_room_info
[ 0] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] room_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] room_remarks                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] hall_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] hall_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] exhibition_start_time                          TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[11] exhibition_end_time                            TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[12] exhibition_status                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[13] item_count                                     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[14] items                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "room_id": "MFMPWqdHeYPvZuGeRJEYxSfQX",    "room_name": "WTOoVWAsjlgouXVNrQXxYYhtg",    "room_floor": "hQDIOpZrjtRfgRgvbhKXmLOtL",    "room_location": "CWxCOiDUaDcFfKqWXeVyWUOha",    "room_status": 74,    "room_remarks": "PXIWCYvggMvedRBPKCvZWwLhf",    "hall_id": "xuEACWSkLkGlvflVMtvPJEPsq",    "hall_name": "xIVBmlMJmuCNegsZlRyMONNkp",    "exhibition_id": "UBmTdkvFnfdQvpsTLLOEyLcmM",    "exhibition_name": "EArxSVCPnxJjakDPnmwwXxHNf",    "exhibition_start_time": 19,    "exhibition_end_time": 84,    "exhibition_status": 66,    "item_count": 98,    "items": 42}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_room_info does not have a primary key defined, setting col position 1 room_id as primary key
Warning table: v_ebcp_exhibition_room_info primary key column room_id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_room_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_room_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_room_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_room_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_room_info_FIELD_NAME_room_status = "room_status"

	Ebcp_exhibition_room_info_FIELD_NAME_room_remarks = "room_remarks"

	Ebcp_exhibition_room_info_FIELD_NAME_hall_id = "hall_id"

	Ebcp_exhibition_room_info_FIELD_NAME_hall_name = "hall_name"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_start_time = "exhibition_start_time"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_end_time = "exhibition_end_time"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_status = "exhibition_status"

	Ebcp_exhibition_room_info_FIELD_NAME_item_count = "item_count"

	Ebcp_exhibition_room_info_FIELD_NAME_items = "items"
)

// Ebcp_exhibition_room_info struct is a row record of the v_ebcp_exhibition_room_info table in the  database
type Ebcp_exhibition_room_info struct {
	RoomID string `json:"room_id"` //room_id

	RoomName string `json:"room_name"` //room_name

	RoomFloor string `json:"room_floor"` //room_floor

	RoomLocation string `json:"room_location"` //room_location

	RoomStatus int32 `json:"room_status"` //room_status

	RoomRemarks string `json:"room_remarks"` //room_remarks

	HallID string `json:"hall_id"` //hall_id

	HallName string `json:"hall_name"` //hall_name

	ExhibitionID string `json:"exhibition_id"` //exhibition_id

	ExhibitionName string `json:"exhibition_name"` //exhibition_name

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //exhibition_start_time

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //exhibition_end_time

	ExhibitionStatus int32 `json:"exhibition_status"` //exhibition_status

	ItemCount int32 `json:"item_count"` //item_count

	Items any `json:"items"` //items

}

var Ebcp_exhibition_room_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_room_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "room_id",
			Comment: `room_id`,
			Notes: `Warning table: v_ebcp_exhibition_room_info does not have a primary key defined, setting col position 1 room_id as primary key
Warning table: v_ebcp_exhibition_room_info primary key column room_id is nullable column, setting it as NOT NULL
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
			Comment:            `room_name`,
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
			Comment:            `room_floor`,
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
			Name:               "room_location",
			Comment:            `room_location`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "room_status",
			Comment:            `room_status`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "room_remarks",
			Comment:            `room_remarks`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "hall_id",
			Comment:            `hall_id`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "hall_name",
			Comment:            `hall_name`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "exhibition_id",
			Comment:            `exhibition_id`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "exhibition_name",
			Comment:            `exhibition_name`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "exhibition_start_time",
			Comment:            `exhibition_start_time`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "exhibition_end_time",
			Comment:            `exhibition_end_time`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "exhibition_status",
			Comment:            `exhibition_status`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "item_count",
			Comment:            `item_count`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "ItemCount",
			GoFieldType:        "int32",
			JSONFieldName:      "item_count",
			ProtobufFieldName:  "item_count",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "items",
			Comment:            `items`,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_room_info) TableName() string {
	return "v_ebcp_exhibition_room_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_room_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_room_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_room_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_room_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_room_infoTableInfo
}
