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
[ 4] hall_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] hall_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] room_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[11] item_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] item_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] item_type                                      VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[14] item_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] room_count                                     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[16] item_count                                     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []


JSON Sample
-------------------------------------
{    "exhibition_id": "dfrWPYojMKLRxBUPTNgfSoCBw",    "exhibition_name": "tOBeIRZKjcXZUBlVOwXxeOjFf",    "exhibition_start_time": 6,    "exhibition_end_time": 65,    "hall_id": "cVpOUvkksHeEWdpsxeueCrhDF",    "hall_name": "mEJkNQBattStwMybaXHkeFuil",    "room_id": "dMmDXWPYPktbXXKGHMFWhKuOt",    "room_name": "qNxCtuCfXSenaMjblHLTehvKe",    "room_floor": "oZOqHThbGiYJCmuXdyIaqGmjK",    "room_location": "BfFKloSyPsbLOmEITfDAaSgxr",    "room_status": 44,    "item_id": "PfWEPQDMZARGLGOyOfJYnYVgV",    "item_name": "xUvTrvPgnXBVdtysoFKSfLTTN",    "item_type": "IBRCrkiTGkbRRhEupPCWMIFlx",    "item_status": 22,    "room_count": 65,    "item_count": 99}


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

	Ebcp_exhibition_info_FIELD_NAME_hall_id = "hall_id"

	Ebcp_exhibition_info_FIELD_NAME_hall_name = "hall_name"

	Ebcp_exhibition_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_info_FIELD_NAME_room_status = "room_status"

	Ebcp_exhibition_info_FIELD_NAME_item_id = "item_id"

	Ebcp_exhibition_info_FIELD_NAME_item_name = "item_name"

	Ebcp_exhibition_info_FIELD_NAME_item_type = "item_type"

	Ebcp_exhibition_info_FIELD_NAME_item_status = "item_status"

	Ebcp_exhibition_info_FIELD_NAME_room_count = "room_count"

	Ebcp_exhibition_info_FIELD_NAME_item_count = "item_count"
)

// Ebcp_exhibition_info struct is a row record of the v_ebcp_exhibition_info table in the  database
type Ebcp_exhibition_info struct {
	ExhibitionID string `json:"exhibition_id"` //exhibition_id

	ExhibitionName string `json:"exhibition_name"` //exhibition_name

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //exhibition_start_time

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //exhibition_end_time

	HallID string `json:"hall_id"` //hall_id

	HallName string `json:"hall_name"` //hall_name

	RoomID string `json:"room_id"` //room_id

	RoomName string `json:"room_name"` //room_name

	RoomFloor string `json:"room_floor"` //room_floor

	RoomLocation string `json:"room_location"` //room_location

	RoomStatus int32 `json:"room_status"` //room_status

	ItemID string `json:"item_id"` //item_id

	ItemName string `json:"item_name"` //item_name

	ItemType string `json:"item_type"` //item_type

	ItemStatus int32 `json:"item_status"` //item_status

	RoomCount int32 `json:"room_count"` //room_count

	ItemCount int32 `json:"item_count"` //item_count

}

var Ebcp_exhibition_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "exhibition_id",
			Comment: `exhibition_id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "room_id",
			Comment:            `room_id`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "item_id",
			Comment:            `item_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ItemID",
			GoFieldType:        "string",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "item_name",
			Comment:            `item_name`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ItemName",
			GoFieldType:        "string",
			JSONFieldName:      "item_name",
			ProtobufFieldName:  "item_name",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "item_type",
			Comment:            `item_type`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "ItemType",
			GoFieldType:        "string",
			JSONFieldName:      "item_type",
			ProtobufFieldName:  "item_type",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "item_status",
			Comment:            `item_status`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ItemStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "item_status",
			ProtobufFieldName:  "item_status",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "room_count",
			Comment:            `room_count`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "RoomCount",
			GoFieldType:        "int32",
			JSONFieldName:      "room_count",
			ProtobufFieldName:  "room_count",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
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
