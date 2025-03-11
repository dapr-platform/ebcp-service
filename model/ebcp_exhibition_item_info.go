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
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] type                                           VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 3] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 5] export_info                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 6] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] room_floor_value                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] room_floor_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] room_location_value                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] room_location_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[14] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[15] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[16] player_devices                                 JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[17] control_devices                                JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[18] schedules                                      JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "WlncbCmNOJJLQdUCRempljWFe",    "name": "GxWSPKAsTQJmZLvTroJloWKVL",    "type": "oOpnrdAJSekKSNapZtUJgNCod",    "status": 19,    "remarks": "OwdXngmxvUVqShUEXpaPIUHpw",    "export_info": "FZlHZELLyiTRARPCboLBmFxaE",    "room_id": "XxKCCTufYmnBpUApbfjKPvdRo",    "room_name": "ngQprheZdCeGlbVFJAXLIohds",    "room_floor": "lNMDvTbffMjPMsTRwJwGovQPj",    "room_floor_value": "eRJQxYmBwngPNRYgOTJLUbwnD",    "room_floor_name": "VunFdojUDOMoRceuNQTuwCXUv",    "room_location": "ZfUfgHALBwtRhNPPUqJkYnoUe",    "room_location_value": "SosSABLLoZTPkZIWnpZwQfdUB",    "room_location_name": "pKpGEvpCYahQgcEcmLTkrRnpZ",    "exhibition_id": "mkpYnQHjkCdlhZwvBAoeREQEf",    "exhibition_name": "mWPhYtuSWruERtdHjSIycwURV",    "player_devices": 7,    "control_devices": 85,    "schedules": 32}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_item_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_item_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_item_info_FIELD_NAME_id = "id"

	Ebcp_exhibition_item_info_FIELD_NAME_name = "name"

	Ebcp_exhibition_item_info_FIELD_NAME_type = "type"

	Ebcp_exhibition_item_info_FIELD_NAME_status = "status"

	Ebcp_exhibition_item_info_FIELD_NAME_remarks = "remarks"

	Ebcp_exhibition_item_info_FIELD_NAME_export_info = "export_info"

	Ebcp_exhibition_item_info_FIELD_NAME_room_id = "room_id"

	Ebcp_exhibition_item_info_FIELD_NAME_room_name = "room_name"

	Ebcp_exhibition_item_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_exhibition_item_info_FIELD_NAME_room_floor_value = "room_floor_value"

	Ebcp_exhibition_item_info_FIELD_NAME_room_floor_name = "room_floor_name"

	Ebcp_exhibition_item_info_FIELD_NAME_room_location = "room_location"

	Ebcp_exhibition_item_info_FIELD_NAME_room_location_value = "room_location_value"

	Ebcp_exhibition_item_info_FIELD_NAME_room_location_name = "room_location_name"

	Ebcp_exhibition_item_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_item_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_item_info_FIELD_NAME_player_devices = "player_devices"

	Ebcp_exhibition_item_info_FIELD_NAME_control_devices = "control_devices"

	Ebcp_exhibition_item_info_FIELD_NAME_schedules = "schedules"
)

// Ebcp_exhibition_item_info struct is a row record of the v_ebcp_exhibition_item_info table in the  database
type Ebcp_exhibition_item_info struct {
	ID string `json:"id"` //id

	Name string `json:"name"` //name

	Type string `json:"type"` //type

	Status int32 `json:"status"` //status

	Remarks string `json:"remarks"` //remarks

	ExportInfo string `json:"export_info"` //export_info

	RoomID string `json:"room_id"` //room_id

	RoomName string `json:"room_name"` //room_name

	RoomFloor string `json:"room_floor"` //room_floor

	RoomFloorValue string `json:"room_floor_value"` //room_floor_value

	RoomFloorName string `json:"room_floor_name"` //room_floor_name

	RoomLocation string `json:"room_location"` //room_location

	RoomLocationValue string `json:"room_location_value"` //room_location_value

	RoomLocationName string `json:"room_location_name"` //room_location_name

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
			Name:    "id",
			Comment: `id`,
			Notes: `Warning table: v_ebcp_exhibition_item_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_item_info primary key column id is nullable column, setting it as NOT NULL
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
			Comment:            `name`,
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
			Name:               "type",
			Comment:            `type`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `status`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "remarks",
			Comment:            `remarks`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "export_info",
			Comment:            `export_info`,
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
			Name:               "room_floor_value",
			Comment:            `room_floor_value`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "room_floor_name",
			Comment:            `room_floor_name`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "room_location_value",
			Comment:            `room_location_value`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "room_location_name",
			Comment:            `room_location_name`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
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
