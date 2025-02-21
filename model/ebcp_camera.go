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


Table: o_ebcp_camera
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] device_no                                      VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 7] main_stream_url                                VARCHAR(1024)        null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 8] sub_stream_url                                 VARCHAR(1024)        null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[ 9] status                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]


JSON Sample
-------------------------------------
{    "id": "BHJsvpCWOOZRHoCHlqWkmHdil",    "created_by": "ZGeIMkNrSftNjigsHGtgHfyIQ",    "created_time": 63,    "updated_by": "NfavelRNlgWrMKVlEmsnLeNTs",    "updated_time": 3,    "name": "dqtrJxbJZPNyaCjmanNLuhkyS",    "device_no": "rYILaKuNOVEHPIhjdZBUnrmic",    "main_stream_url": "udTVufvmJHrPUiImFtAchBvuZ",    "sub_stream_url": "faMflYoQOhUPNlrISIjOsNkdl",    "status": 7}



*/

var (
	Ebcp_camera_FIELD_NAME_id = "id"

	Ebcp_camera_FIELD_NAME_created_by = "created_by"

	Ebcp_camera_FIELD_NAME_created_time = "created_time"

	Ebcp_camera_FIELD_NAME_updated_by = "updated_by"

	Ebcp_camera_FIELD_NAME_updated_time = "updated_time"

	Ebcp_camera_FIELD_NAME_name = "name"

	Ebcp_camera_FIELD_NAME_device_no = "device_no"

	Ebcp_camera_FIELD_NAME_main_stream_url = "main_stream_url"

	Ebcp_camera_FIELD_NAME_sub_stream_url = "sub_stream_url"

	Ebcp_camera_FIELD_NAME_status = "status"
)

// Ebcp_camera struct is a row record of the o_ebcp_camera table in the  database
type Ebcp_camera struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	Name string `json:"name"` //摄像头名称

	DeviceNo string `json:"device_no"` //设备编号

	MainStreamURL string `json:"main_stream_url"` //主码流URL

	SubStreamURL string `json:"sub_stream_url"` //辅码流URL

	Status int32 `json:"status"` //状态(1: 正常, 2: 故障)

}

var Ebcp_cameraTableInfo = &TableInfo{
	Name: "o_ebcp_camera",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
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
			Name:               "created_by",
			Comment:            `created_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CreatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "created_by",
			ProtobufFieldName:  "created_by",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "created_time",
			Comment:            `created_time`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "updated_by",
			Comment:            `updated_by`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_time",
			Comment:            `updated_time`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "name",
			Comment:            `摄像头名称`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "device_no",
			Comment:            `设备编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "DeviceNo",
			GoFieldType:        "string",
			JSONFieldName:      "device_no",
			ProtobufFieldName:  "device_no",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "main_stream_url",
			Comment:            `主码流URL`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "MainStreamURL",
			GoFieldType:        "string",
			JSONFieldName:      "main_stream_url",
			ProtobufFieldName:  "main_stream_url",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "sub_stream_url",
			Comment:            `辅码流URL`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "SubStreamURL",
			GoFieldType:        "string",
			JSONFieldName:      "sub_stream_url",
			ProtobufFieldName:  "sub_stream_url",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
			Comment:            `状态(1: 正常, 2: 故障)`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_camera) TableName() string {
	return "o_ebcp_camera"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_camera) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_camera) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_camera) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_camera) TableInfo() *TableInfo {
	return Ebcp_cameraTableInfo
}
