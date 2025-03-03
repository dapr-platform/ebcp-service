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


Table: v_ebcp_exhibition_item_info
[ 0] item_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] item_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] item_type                                      VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 3] item_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] item_remarks                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] hall_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] hall_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] player_devices                                 JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[14] control_devices                                JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[15] schedules                                      JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "item_id": "aCsUNkZrnFCKUWRblvIVnEKrT",    "item_name": "JsnpTrMomPeMwFJJHXMailWjY",    "item_type": "hSvYdoXMevdhCMDKYhxrhKLid",    "item_status": 64,    "item_remarks": "tFshxJOvBhaBmdUBmgelYgifo",    "room_id": "ZlXqhZXfBrGrGAWCNQvrIWYKl",    "room_name": "fpGRxJWHKRjZukCeUGQdjQCNy",    "room_floor": "MLIkqLSUVAmZbSOwYRrllubCm",    "room_location": "WncZfkHAHKHjGLAtWxWqBHZGO",    "hall_id": "mZGSpvjvRTCjXXXUIybTkgUqQ",    "hall_name": "tLbohYYVsRMYAIgdtHrMxeaar",    "exhibition_id": "HYkIjUPXTTlabgNUChCWVShDX",    "exhibition_name": "SbmisqaaQMIdiQbMRmPgpLmoB",    "player_devices": 52,    "control_devices": 68,    "schedules": 48}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_item_info does not have a primary key defined, setting col position 1 item_id as primary key
Warning table: v_ebcp_exhibition_item_info primary key column item_id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_item_info_FIELD_NAME_item_id = "item_id"

	Ebcp_exhibition_item_info_FIELD_NAME_item_name = "item_name"

	Ebcp_exhibition_item_info_FIELD_NAME_item_type = "item_type"

	Ebcp_exhibition_item_info_FIELD_NAME_item_status = "item_status"

	Ebcp_exhibition_item_info_FIELD_NAME_item_remarks = "item_remarks"

	Ebcp_exhibition_item_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_item_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_item_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_item_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_item_info_FIELD_NAME_hall_id = "hall_id"

	Ebcp_exhibition_item_info_FIELD_NAME_hall_name = "hall_name"

	Ebcp_exhibition_item_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_item_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_item_info_FIELD_NAME_player_devices = "player_devices"

	Ebcp_exhibition_item_info_FIELD_NAME_control_devices = "control_devices"

	Ebcp_exhibition_item_info_FIELD_NAME_schedules = "schedules"
)

// Ebcp_exhibition_item_info struct is a row record of the v_ebcp_exhibition_item_info table in the  database
type Ebcp_exhibition_item_info struct {
	ItemID string `json:"item_id"` //item_id

	ItemName string `json:"item_name"` //item_name

	ItemType string `json:"item_type"` //item_type

	ItemStatus int32 `json:"item_status"` //item_status

	ItemRemarks string `json:"item_remarks"` //item_remarks

	RoomID string `json:"room_id"` //room_id

	RoomName string `json:"room_name"` //room_name

	RoomFloor string `json:"room_floor"` //room_floor

	RoomLocation string `json:"room_location"` //room_location

	HallID string `json:"hall_id"` //hall_id

	HallName string `json:"hall_name"` //hall_name

	ExhibitionID string `json:"exhibition_id"` //exhibition_id

	ExhibitionName string `json:"exhibition_name"` //exhibition_name

	PlayerDevices any `json:"player_devices"` //player_devices

	ControlDevices any `json:"control_devices"` //control_devices

	Schedules any `json:"schedules"` //schedules

}

var Ebcp_exhibition_item_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_item_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "item_id",
			Comment: `item_id`,
			Notes: `Warning table: v_ebcp_exhibition_item_info does not have a primary key defined, setting col position 1 item_id as primary key
Warning table: v_ebcp_exhibition_item_info primary key column item_id is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ItemID",
			GoFieldType:        "string",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "item_remarks",
			Comment:            `item_remarks`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ItemRemarks",
			GoFieldType:        "string",
			JSONFieldName:      "item_remarks",
			ProtobufFieldName:  "item_remarks",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "player_devices",
			Comment:            `player_devices`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "PlayerDevices",
			GoFieldType:        "any",
			JSONFieldName:      "player_devices",
			ProtobufFieldName:  "player_devices",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "control_devices",
			Comment:            `control_devices`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "ControlDevices",
			GoFieldType:        "any",
			JSONFieldName:      "control_devices",
			ProtobufFieldName:  "control_devices",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "schedules",
			Comment:            `schedules`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Schedules",
			GoFieldType:        "any",
			JSONFieldName:      "schedules",
			ProtobufFieldName:  "schedules",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_item_info) TableName() string {
	return "v_ebcp_exhibition_item_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_item_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_item_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_item_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_item_info) TableInfo() *TableInfo {
	return Ebcp_exhibition_item_infoTableInfo
}
