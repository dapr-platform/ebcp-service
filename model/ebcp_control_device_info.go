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


Table: v_ebcp_control_device_info
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] device_type                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[ 3] ip_address                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] port                                           INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] version                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 7] commands                                       TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 9] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[10] item_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] index_num                                      INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[12] item_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[13] item_type                                      VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
[14] room_id                                        VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[15] room_name                                      VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[16] room_status                                    INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[17] room_remarks                                   TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[18] room_floor                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[19] room_floor_value                               VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[20] room_floor_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[21] room_location                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[22] room_location_value                            VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[23] room_location_name                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[24] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[25] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[26] exhibition_status                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[27] exhibition_hall_id                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[28] exhibition_hall_name                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[29] exhibition_hall_remarks                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[30] linked_item                                    JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "JHWexaTATtNKQGwwkrMuupASr",    "name": "ZHmmEZAZKnbiXnuNWIQJlIXrG",    "device_type": "BKiwACvkvbZEErrwAZwldTyJy",    "ip_address": "RFsGvIolVIIMDtjaeHrScTCyb",    "port": 14,    "version": "clYMMYToeoSkIthExWrkrJUlw",    "status": 51,    "commands": "BMVGWoUGRoNQnFMxasCMpwZiO",    "created_time": 13,    "updated_time": 56,    "item_id": "OfcxsbGDKJtwJgRFrlgCoXtmM",    "index_num": 26,    "item_name": "kwnwxHdncXHZhmpgWrQTsomsM",    "item_type": "XcZLBJcbgUyDYADYLcUqaqOxC",    "room_id": "XpxyxwBJUCZUSpRKXJFAWcYSL",    "room_name": "KsxjgPUAGTthGbaCkYWEVwXtC",    "room_status": 60,    "room_remarks": "xFTBBCoMdgvTRywyWFoxyFoIZ",    "room_floor": "FTqUhvobOULLymSkbBxrQpqBj",    "room_floor_value": "SJLwLPiIABZSEumaYfWoYKyYR",    "room_floor_name": "HZiOEqaYZTjgtSanFORgRDwgg",    "room_location": "mrGIUxUoULqDbtMQoTDYFgUaT",    "room_location_value": "xLtqmFkvJNDvWxajElFCcYXus",    "room_location_name": "QbEBFhhtYPeryBWagXjHNoDDw",    "exhibition_id": "CpvDaVvbmkfNcRQTvWMkHCUEx",    "exhibition_name": "gLnoePtTSLQZiBogPdmhMlHcb",    "exhibition_status": 31,    "exhibition_hall_id": "RekWQdiWSjqWAtkMaLeqcsTsg",    "exhibition_hall_name": "YQUwKGitdNUmBmjmWwCrIvpLF",    "exhibition_hall_remarks": "DXLxIBusYPEcukJpTQLgEYrtx",    "linked_item": 96}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_control_device_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_control_device_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_control_device_info_FIELD_NAME_id = "id"

	Ebcp_control_device_info_FIELD_NAME_name = "name"

	Ebcp_control_device_info_FIELD_NAME_device_type = "device_type"

	Ebcp_control_device_info_FIELD_NAME_ip_address = "ip_address"

	Ebcp_control_device_info_FIELD_NAME_port = "port"

	Ebcp_control_device_info_FIELD_NAME_version = "version"

	Ebcp_control_device_info_FIELD_NAME_status = "status"

	Ebcp_control_device_info_FIELD_NAME_commands = "commands"

	Ebcp_control_device_info_FIELD_NAME_created_time = "created_time"

	Ebcp_control_device_info_FIELD_NAME_updated_time = "updated_time"

	Ebcp_control_device_info_FIELD_NAME_item_id = "item_id"

	Ebcp_control_device_info_FIELD_NAME_index_num = "index_num"

	Ebcp_control_device_info_FIELD_NAME_item_name = "item_name"

	Ebcp_control_device_info_FIELD_NAME_item_type = "item_type"

	Ebcp_control_device_info_FIELD_NAME_room_id = "room_id"

	Ebcp_control_device_info_FIELD_NAME_room_name = "room_name"

	Ebcp_control_device_info_FIELD_NAME_room_status = "room_status"

	Ebcp_control_device_info_FIELD_NAME_room_remarks = "room_remarks"

	Ebcp_control_device_info_FIELD_NAME_room_floor = "room_floor"

	Ebcp_control_device_info_FIELD_NAME_room_floor_value = "room_floor_value"

	Ebcp_control_device_info_FIELD_NAME_room_floor_name = "room_floor_name"

	Ebcp_control_device_info_FIELD_NAME_room_location = "room_location"

	Ebcp_control_device_info_FIELD_NAME_room_location_value = "room_location_value"

	Ebcp_control_device_info_FIELD_NAME_room_location_name = "room_location_name"

	Ebcp_control_device_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_control_device_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_control_device_info_FIELD_NAME_exhibition_status = "exhibition_status"

	Ebcp_control_device_info_FIELD_NAME_exhibition_hall_id = "exhibition_hall_id"

	Ebcp_control_device_info_FIELD_NAME_exhibition_hall_name = "exhibition_hall_name"

	Ebcp_control_device_info_FIELD_NAME_exhibition_hall_remarks = "exhibition_hall_remarks"

	Ebcp_control_device_info_FIELD_NAME_linked_item = "linked_item"
)

// Ebcp_control_device_info struct is a row record of the v_ebcp_control_device_info table in the  database
type Ebcp_control_device_info struct {
	ID string `json:"id"` //设备ID

	Name string `json:"name"` //设备名称

	DeviceType string `json:"device_type"` //设备类型

	IPAddress string `json:"ip_address"` //IP地址

	Port int32 `json:"port"` //端口

	Version string `json:"version"` //版本

	Status int32 `json:"status"` //设备状态(1: 正常, 2: 故障)

	Commands string `json:"commands"` //命令列表

	CreatedTime common.LocalTime `json:"created_time"` //创建时间

	UpdatedTime common.LocalTime `json:"updated_time"` //更新时间

	ItemID string `json:"item_id"` //item_id

	IndexNum int32 `json:"index_num"` //index_num

	ItemName string `json:"item_name"` //item_name

	ItemType string `json:"item_type"` //item_type

	RoomID string `json:"room_id"` //所属展厅ID

	RoomName string `json:"room_name"` //所属展厅名称

	RoomStatus int32 `json:"room_status"` //所属展厅状态

	RoomRemarks string `json:"room_remarks"` //所属展厅备注

	RoomFloor string `json:"room_floor"` //所属展厅楼层

	RoomFloorValue string `json:"room_floor_value"` //所属展厅楼层值

	RoomFloorName string `json:"room_floor_name"` //所属展厅楼层名称

	RoomLocation string `json:"room_location"` //所属展厅位置

	RoomLocationValue string `json:"room_location_value"` //所属展厅位置值

	RoomLocationName string `json:"room_location_name"` //所属展厅位置名称

	ExhibitionID string `json:"exhibition_id"` //所属展览ID

	ExhibitionName string `json:"exhibition_name"` //所属展览名称

	ExhibitionStatus int32 `json:"exhibition_status"` //所属展览状态

	ExhibitionHallID string `json:"exhibition_hall_id"` //所属展馆ID

	ExhibitionHallName string `json:"exhibition_hall_name"` //所属展馆名称

	ExhibitionHallRemarks string `json:"exhibition_hall_remarks"` //所属展馆备注

	LinkedItem any `json:"linked_item"` //直接关联的展项（JSON格式）

}

var Ebcp_control_device_infoTableInfo = &TableInfo{
	Name: "v_ebcp_control_device_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `设备ID`,
			Notes: `Warning table: v_ebcp_control_device_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_control_device_info primary key column id is nullable column, setting it as NOT NULL
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
			Name:               "device_type",
			Comment:            `设备类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       50,
			GoFieldName:        "DeviceType",
			GoFieldType:        "string",
			JSONFieldName:      "device_type",
			ProtobufFieldName:  "device_type",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "status",
			Comment:            `设备状态(1: 正常, 2: 故障)`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "commands",
			Comment:            `命令列表`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Commands",
			GoFieldType:        "string",
			JSONFieldName:      "commands",
			ProtobufFieldName:  "commands",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "created_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "index_num",
			Comment:            `index_num`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IndexNum",
			GoFieldType:        "int32",
			JSONFieldName:      "index_num",
			ProtobufFieldName:  "index_num",
			ProtobufType:       "int32",
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "room_status",
			Comment:            `所属展厅状态`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "room_remarks",
			Comment:            `所属展厅备注`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "room_location",
			Comment:            `所属展厅位置`,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "room_location_value",
			Comment:            `所属展厅位置值`,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "room_location_name",
			Comment:            `所属展厅位置名称`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "exhibition_status",
			Comment:            `所属展览状态`,
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
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "exhibition_hall_id",
			Comment:            `所属展馆ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionHallID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_hall_id",
			ProtobufFieldName:  "exhibition_hall_id",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "exhibition_hall_name",
			Comment:            `所属展馆名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ExhibitionHallName",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_hall_name",
			ProtobufFieldName:  "exhibition_hall_name",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "exhibition_hall_remarks",
			Comment:            `所属展馆备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "ExhibitionHallRemarks",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_hall_remarks",
			ProtobufFieldName:  "exhibition_hall_remarks",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "linked_item",
			Comment:            `直接关联的展项（JSON格式）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "JSON",
			DatabaseTypePretty: "JSON",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "JSON",
			ColumnLength:       -1,
			GoFieldName:        "LinkedItem",
			GoFieldType:        "any",
			JSONFieldName:      "linked_item",
			ProtobufFieldName:  "linked_item",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_control_device_info) TableName() string {
	return "v_ebcp_control_device_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_control_device_info) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_control_device_info) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_control_device_info) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_control_device_info) TableInfo() *TableInfo {
	return Ebcp_control_device_infoTableInfo
}
