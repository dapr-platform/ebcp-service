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


Table: o_ebcp_exhibition_area
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 6] exhibition_room_id                             VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] current_exhibition_name                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 8] current_exhibition_start_time                  TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 9] current_exhibition_end_time                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[10] remarks                                        TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "iJfLPFEhuZVkrXdWIcMYLdTkG",    "created_by": "xuaEWVpHMCjFSjcmMOjJlMAXc",    "created_time": 13,    "updated_by": "WxPPHEOeDYxSBiojwaRElMbso",    "updated_time": 66,    "name": "lfUqLLAnLFrUHPvJRyXKrbSYu",    "exhibition_room_id": "aqrMrpKHGkAPEfuyGdEklaAOy",    "current_exhibition_name": "uaABWNWnFppVHdMROrnZYcjRC",    "current_exhibition_start_time": 43,    "current_exhibition_end_time": 66,    "remarks": "GVTpDrJUwrvlccPegWotMiQbM"}



*/

var (
	Ebcp_exhibition_area_FIELD_NAME_id = "id"

	Ebcp_exhibition_area_FIELD_NAME_created_by = "created_by"

	Ebcp_exhibition_area_FIELD_NAME_created_time = "created_time"

	Ebcp_exhibition_area_FIELD_NAME_updated_by = "updated_by"

	Ebcp_exhibition_area_FIELD_NAME_updated_time = "updated_time"

	Ebcp_exhibition_area_FIELD_NAME_name = "name"

	Ebcp_exhibition_area_FIELD_NAME_exhibition_room_id = "exhibition_room_id"

	Ebcp_exhibition_area_FIELD_NAME_current_exhibition_name = "current_exhibition_name"

	Ebcp_exhibition_area_FIELD_NAME_current_exhibition_start_time = "current_exhibition_start_time"

	Ebcp_exhibition_area_FIELD_NAME_current_exhibition_end_time = "current_exhibition_end_time"

	Ebcp_exhibition_area_FIELD_NAME_remarks = "remarks"
)

// Ebcp_exhibition_area struct is a row record of the o_ebcp_exhibition_area table in the  database
type Ebcp_exhibition_area struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	Name string `json:"name"` //展区名称

	ExhibitionRoomID string `json:"exhibition_room_id"` //所属展厅ID

	CurrentExhibitionName string `json:"current_exhibition_name"` //当前展览名称

	CurrentExhibitionStartTime common.LocalTime `json:"current_exhibition_start_time"` //当前展览开始时间

	CurrentExhibitionEndTime common.LocalTime `json:"current_exhibition_end_time"` //当前展览结束时间

	Remarks string `json:"remarks"` //备注

}

var Ebcp_exhibition_areaTableInfo = &TableInfo{
	Name: "o_ebcp_exhibition_area",
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
			Comment:            `展区名称`,
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
			Name:               "exhibition_room_id",
			Comment:            `所属展厅ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "ExhibitionRoomID",
			GoFieldType:        "string",
			JSONFieldName:      "exhibition_room_id",
			ProtobufFieldName:  "exhibition_room_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "current_exhibition_name",
			Comment:            `当前展览名称`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "CurrentExhibitionName",
			GoFieldType:        "string",
			JSONFieldName:      "current_exhibition_name",
			ProtobufFieldName:  "current_exhibition_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "current_exhibition_start_time",
			Comment:            `当前展览开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CurrentExhibitionStartTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "current_exhibition_start_time",
			ProtobufFieldName:  "current_exhibition_start_time",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "current_exhibition_end_time",
			Comment:            `当前展览结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CurrentExhibitionEndTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "current_exhibition_end_time",
			ProtobufFieldName:  "current_exhibition_end_time",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "remarks",
			Comment:            `备注`,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_exhibition_area) TableName() string {
	return "o_ebcp_exhibition_area"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_exhibition_area) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_exhibition_area) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_exhibition_area) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_exhibition_area) TableInfo() *TableInfo {
	return Ebcp_exhibition_areaTableInfo
}
