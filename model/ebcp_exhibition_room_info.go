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
[ 0] id                                             VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] name                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] floor                                          VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] floor_value                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 4] floor_name                                     VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 5] location                                       VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] location_value                                 VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] location_name                                  VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] status                                         INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 9] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[10] exhibition_id                                  VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] exhibition_name                                VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] exhibition_start_time                          TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[13] exhibition_end_time                            TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[14] exhibition_status                              INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[15] item_count                                     INT8                 null: true   primary: false  isArray: false  auto: false  col: INT8            len: -1      default: []
[16] items                                          JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []
[17] control_devices                                JSON                 null: true   primary: false  isArray: false  auto: false  col: JSON            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "hIPdTbFmfgPALUbxiBUquoqKY",    "name": "lqbWjjVpstXZJDtRFEuxcqQnW",    "floor": "WwGcJifIwihmbInJoHQLZIDld",    "floor_value": "avACfDGiplPMNjiyqPRqWcUuV",    "floor_name": "utXnsBPDaVNvkDwcFuqcAgcMx",    "location": "TxPODXRwhNhppQOSUWcIXqHBP",    "location_value": "LRINxPuUNoxolgqVlAWubdNHu",    "location_name": "VnqqkeZtVBIMueFYSclOuUnLd",    "status": 26,    "remarks": "kGAtHfUuAsZmoNoVrWbpdLrNb",    "exhibition_id": "nhjbmFpjEXeKhZwvSmSpsMZeI",    "exhibition_name": "wtTixqnsFDeYjptwPQgYiWRmZ",    "exhibition_start_time": 88,    "exhibition_end_time": 59,    "exhibition_status": 32,    "item_count": 28,    "items": 7,    "control_devices": 95}


Comments
-------------------------------------
[ 0] Warning table: v_ebcp_exhibition_room_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_room_info primary key column id is nullable column, setting it as NOT NULL




*/

var (
	Ebcp_exhibition_room_info_FIELD_NAME_id = "id"

	Ebcp_exhibition_room_info_FIELD_NAME_name = "name"

	Ebcp_exhibition_room_info_FIELD_NAME_floor = "floor"

	Ebcp_exhibition_room_info_FIELD_NAME_floor_value = "floor_value"

	Ebcp_exhibition_room_info_FIELD_NAME_floor_name = "floor_name"

	Ebcp_exhibition_room_info_FIELD_NAME_location = "location"

	Ebcp_exhibition_room_info_FIELD_NAME_location_value = "location_value"

	Ebcp_exhibition_room_info_FIELD_NAME_location_name = "location_name"

	Ebcp_exhibition_room_info_FIELD_NAME_status = "status"

	Ebcp_exhibition_room_info_FIELD_NAME_remarks = "remarks"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_id = "exhibition_id"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_name = "exhibition_name"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_start_time = "exhibition_start_time"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_end_time = "exhibition_end_time"

	Ebcp_exhibition_room_info_FIELD_NAME_exhibition_status = "exhibition_status"

	Ebcp_exhibition_room_info_FIELD_NAME_item_count = "item_count"

	Ebcp_exhibition_room_info_FIELD_NAME_items = "items"

	Ebcp_exhibition_room_info_FIELD_NAME_control_devices = "control_devices"
)

// Ebcp_exhibition_room_info struct is a row record of the v_ebcp_exhibition_room_info table in the  database
type Ebcp_exhibition_room_info struct {
	ID string `json:"id"` //展厅ID

	Name string `json:"name"` //展厅名称

	Floor string `json:"floor"` //展厅楼层

	FloorValue string `json:"floor_value"` //展厅楼层值

	FloorName string `json:"floor_name"` //展厅楼层名称

	Location string `json:"location"` //展厅位置

	LocationValue string `json:"location_value"` //展厅位置值

	LocationName string `json:"location_name"` //展厅位置名称

	Status int32 `json:"status"` //展厅状态

	Remarks string `json:"remarks"` //展厅备注

	ExhibitionID string `json:"exhibition_id"` //所属展览ID

	ExhibitionName string `json:"exhibition_name"` //所属展览名称

	ExhibitionStartTime common.LocalTime `json:"exhibition_start_time"` //所属展览开始时间

	ExhibitionEndTime common.LocalTime `json:"exhibition_end_time"` //所属展览结束时间

	ExhibitionStatus int32 `json:"exhibition_status"` //所属展览状态

	ItemCount int32 `json:"item_count"` //展厅内展项数量

	Items any `json:"items"` //展厅内的展项列表（JSON格式）

	ControlDevices any `json:"control_devices"` //展厅内的中控设备列表（JSON格式）

}

var Ebcp_exhibition_room_infoTableInfo = &TableInfo{
	Name: "v_ebcp_exhibition_room_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "id",
			Comment: `展厅ID`,
			Notes: `Warning table: v_ebcp_exhibition_room_info does not have a primary key defined, setting col position 1 id as primary key
Warning table: v_ebcp_exhibition_room_info primary key column id is nullable column, setting it as NOT NULL
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "floor",
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
			GoFieldName:        "Floor",
			GoFieldType:        "string",
			JSONFieldName:      "floor",
			ProtobufFieldName:  "floor",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "floor_value",
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
			GoFieldName:        "FloorValue",
			GoFieldType:        "string",
			JSONFieldName:      "floor_value",
			ProtobufFieldName:  "floor_value",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "floor_name",
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
			GoFieldName:        "FloorName",
			GoFieldType:        "string",
			JSONFieldName:      "floor_name",
			ProtobufFieldName:  "floor_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "location",
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
			GoFieldName:        "Location",
			GoFieldType:        "string",
			JSONFieldName:      "location",
			ProtobufFieldName:  "location",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "location_value",
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
			GoFieldName:        "LocationValue",
			GoFieldType:        "string",
			JSONFieldName:      "location_value",
			ProtobufFieldName:  "location_value",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "location_name",
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
			GoFieldName:        "LocationName",
			GoFieldType:        "string",
			JSONFieldName:      "location_name",
			ProtobufFieldName:  "location_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "remarks",
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
			GoFieldName:        "Remarks",
			GoFieldType:        "string",
			JSONFieldName:      "remarks",
			ProtobufFieldName:  "remarks",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "exhibition_start_time",
			Comment:            `所属展览开始时间`,
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
			Comment:            `所属展览结束时间`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "item_count",
			Comment:            `展厅内展项数量`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "items",
			Comment:            `展厅内的展项列表（JSON格式）`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "control_devices",
			Comment:            `展厅内的中控设备列表（JSON格式）`,
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
