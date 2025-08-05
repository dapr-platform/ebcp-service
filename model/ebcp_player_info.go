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


Table: v_ebcp_player_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] port                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 4] version                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 6] current_program_id                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] current_program_state                          INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 8] volume                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] sound_state                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[10] item_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] item_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] item_type                                      VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[13] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[14] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[15] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[16] room_floor_value                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[17] room_floor_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[18] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[19] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[20] programs                                       JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "VLfOdLSpPsKGRDTvGGuYEIbWP",    "name": "nBLZWcuCrunOXqnADRVAWeMXD",    "ip_address": "ahRgRaFSXEllMngUsHkLwJMQR",    "port": 74,    "version": "OgBwrVKoWuUfUVgJcShhQSjFx",    "status": 13,    "current_program_id": "ysgCVSQxFDPTbnvcOTvqNkGtS",    "current_program_state": 96,    "volume": 55,    "sound_state": 45,    "item_id": "LFbuBmDLpOAwnMfqakgcXBHlG",    "item_name": "EIkCDkImUdygIngXVwectvrmi",    "item_type": "uyqGOrJPyfArBerEbXTIgaZoW",    "room_id": "NFAmGtQFvdawgRnxTwdtekZhc",    "room_name": "KSHbBqWGdyhErjwqGbUpJqTVo",    "room_floor": "bneZQARqcILqLldgoutfHvlfs",    "room_floor_value": "FgjYaUkCJbdbpZGjmevCuaUhp",    "room_floor_name": "GWHaYrYjWYivgWePhVydRltYB",    "exhibition_id": "YJTUiWKBNQHiWVUvbvQfCRsuO",    "exhibition_name": "rbtunuOgVtJXiBcyoOleoGRjP",    "programs": 26}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_player_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_player_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_player_info_FIELD_NAME_id = "id"

	Ebcp_player_info_FIELD_NAME_name = "name"

	Ebcp_player_info_FIELD_NAME_ip_address = "ip_address"

	Ebcp_player_info_FIELD_NAME_port = "port"

	Ebcp_player_info_FIELD_NAME_version = "version"

	Ebcp_player_info_FIELD_NAME_status = "status"

	Ebcp_player_info_FIELD_NAME_current_program_id = "current_program_id"

	Ebcp_player_info_FIELD_NAME_current_program_state = "current_program_state"

	Ebcp_player_info_FIELD_NAME_volume = "volume"

	Ebcp_player_info_FIELD_NAME_sound_state = "sound_state"

	Ebcp_player_info_FIELD_NAME_item_id = "item_id"

	Ebcp_player_info_FIELD_NAME_item_name = "item_name"

	Ebcp_player_info_FIELD_NAME_item_type = "item_type"

	Ebcp_player_info_FIELD_NAME_room_id = "room_id"

	Ebcp_player_info_FIELD_NAME_room_name = "room_name"

	Ebcp_player_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_player_info_FIELD_NAME_room_floor_value = "room_floor_value"

	Ebcp_player_info_FIELD_NAME_room_floor_name = "room_floor_name"

	Ebcp_player_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_player_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_player_info_FIELD_NAME_programs = "programs"
)

// Ebcp_player_info struct is a row record of the v_ebcp_player_info table in the  database
type Ebcp_player_info struct {
	ID string `json:"id"` //设备ID

	Name string `json:"name"` //设备名称

	IPAddress string `json:"ip_address"` //IP地址

	Port int32 `json:"port"` //端口

	Version string `json:"version"` //版本

	Status int32 `json:"status"` //状态,1:正常,2:离线,3:故障

	CurrentProgramID string `json:"current_program_id"` //当前节目ID

	CurrentProgramState int32 `json:"current_program_state"` //当前节目状态,0:播放,1:暂停,2:停止

	Volume int32 `json:"volume"` //音量(0-100)

	SoundState int32 `json:"sound_state"` //音量状态,0:静音,1:非静音

	ItemID string `json:"item_id"` //所属展项ID

	ItemName string `json:"item_name"` //所属展项名称

	ItemType string `json:"item_type"` //所属展项类型

	RoomID string `json:"room_id"` //所属展厅ID

	RoomName string `json:"room_name"` //所属展厅名称

	RoomFloor string `json:"room_floor"` //所属展厅楼层

	RoomFloorValue string `json:"room_floor_value"` //所属展厅楼层值

	RoomFloorName string `json:"room_floor_name"` //所属展厅楼层名称

	ExhibitionID string `json:"exhibition_id"` //所属展览ID

	ExhibitionName string `json:"exhibition_name"` //所属展览名称

	Programs any `json:"programs"` //关联的节目信息

}

var Ebcp_player_infoTableInfo = &TableInfo{
	Name: "v_ebcp_player_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `设备ID`,
			Notes: `Warning table: v_ebcp_player_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_player_info primary key column id is nullable column, setting it as NOT NULL
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
			Comment:            `设备名称`,
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
			Name:               "ip_address",
			Comment:            `IP地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "IPAddress",
			GoFieldType:        "string",
			JSONFieldName:      "ip_address",
			ProtobufFieldName:  "ip_address",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "port",
			Comment:            `端口`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Port",
			GoFieldType:        "int32",
			JSONFieldName:      "port",
			ProtobufFieldName:  "port",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "version",
			Comment:            `版本`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Version",
			GoFieldType:        "string",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "status",
			Comment:            `状态,1:正常,2:离线,3:故障`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "current_program_id",
			Comment:            `当前节目ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CurrentProgramID",
			GoFieldType:        "string",
			JSONFieldName:      "current_program_id",
			ProtobufFieldName:  "current_program_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "current_program_state",
			Comment:            `当前节目状态,0:播放,1:暂停,2:停止`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "CurrentProgramState",
			GoFieldType:        "int32",
			JSONFieldName:      "current_program_state",
			ProtobufFieldName:  "current_program_state",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "volume",
			Comment:            `音量(0-100)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Volume",
			GoFieldType:        "int32",
			JSONFieldName:      "volume",
			ProtobufFieldName:  "volume",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "sound_state",
			Comment:            `音量状态,0:静音,1:非静音`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "SoundState",
			GoFieldType:        "int32",
			JSONFieldName:      "sound_state",
			ProtobufFieldName:  "sound_state",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "item_id",
			Comment:            `所属展项ID`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "item_name",
			Comment:            `所属展项名称`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "item_type",
			Comment:            `所属展项类型`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "room_id",
			Comment:            `所属展厅ID`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "room_name",
			Comment:            `所属展厅名称`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "room_floor",
			Comment:            `所属展厅楼层`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "room_floor_value",
			Comment:            `所属展厅楼层值`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "room_floor_name",
			Comment:            `所属展厅楼层名称`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "exhibition_id",
			Comment:            `所属展览ID`,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "exhibition_name",
			Comment:            `所属展览名称`,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "programs",
			Comment:            `关联的节目信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "Programs",
			GoFieldType:        "any",
			JSONFieldName:      "programs",
			ProtobufFieldName:  "programs",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_player_info) TableName() string {
	return "v_ebcp_player_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_player_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_player_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_player_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_player_info) TableInfo() *TableInfo {
	return Ebcp_player_infoTableInfo
}
