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


Table: o_ebcp_item_schedule
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 5] item_id                                        VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] start_time                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] stop_time                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] cycle_type                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "ggbnLrCwZwiklQYfVjvHKPvuJ",    "created_by": "fXIPZfLTYLNmwACPuLNyHAwSR",    "created_time": 52,    "updated_by": "LQlyZxjaSlIXSkErLmBmyuDdV",    "updated_time": 78,    "item_id": "TIGkKqlYNgQtyfVEHLOsonAMu",    "start_time": "pbMJJisROKwqIuTtwEAyCUIbd",    "stop_time": "gISiofpEAhdYbnrxoDWyqnrxD",    "cycle_type": 22}



*/

var (
	Ebcp_item_schedule_FIELD_NAME_id = "id"

	Ebcp_item_schedule_FIELD_NAME_created_by = "created_by"

	Ebcp_item_schedule_FIELD_NAME_created_time = "created_time"

	Ebcp_item_schedule_FIELD_NAME_updated_by = "updated_by"

	Ebcp_item_schedule_FIELD_NAME_updated_time = "updated_time"

	Ebcp_item_schedule_FIELD_NAME_item_id = "item_id"

	Ebcp_item_schedule_FIELD_NAME_start_time = "start_time"

	Ebcp_item_schedule_FIELD_NAME_stop_time = "stop_time"

	Ebcp_item_schedule_FIELD_NAME_cycle_type = "cycle_type"
)

// Ebcp_item_schedule struct is a row record of the o_ebcp_item_schedule table in the  database
type Ebcp_item_schedule struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	ItemID string `json:"item_id"` //展项ID

	StartTime string `json:"start_time"` //开始时间

	StopTime string `json:"stop_time"` //停止时间

	CycleType int32 `json:"cycle_type"` //循环方式(1:工作日, 2:周末, 3:节假日, 4:闭馆日, 5:每天)

}

var Ebcp_item_scheduleTableInfo = &TableInfo{
	Name: "o_ebcp_item_schedule",
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
			Name:               "item_id",
			Comment:            `展项ID`,
			Notes:              ``,
			Nullable:           false,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "start_time",
			Comment:            `开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "StartTime",
			GoFieldType:        "string",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "stop_time",
			Comment:            `停止时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "StopTime",
			GoFieldType:        "string",
			JSONFieldName:      "stop_time",
			ProtobufFieldName:  "stop_time",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "cycle_type",
			Comment:            `循环方式(1:工作日, 2:周末, 3:节假日, 4:闭馆日, 5:每天)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "CycleType",
			GoFieldType:        "int32",
			JSONFieldName:      "cycle_type",
			ProtobufFieldName:  "cycle_type",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_item_schedule) TableName() string {
	return "o_ebcp_item_schedule"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_item_schedule) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_item_schedule) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_item_schedule) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_item_schedule) TableInfo() *TableInfo {
	return Ebcp_item_scheduleTableInfo
}
