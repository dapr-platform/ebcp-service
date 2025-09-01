package model

import (
	"database/sql"
	"time"

	"github.com/dapr-platform/common"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_ebcp_schedule_job
[ 0] id                                             VARCHAR(32)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 1] created_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 2] created_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 3] updated_by                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] updated_time                                   TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 5] rel_id                                         VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] rel_type                                       VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 7] start_time                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 8] stop_time                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 9] start_date                                     VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[10] stop_date                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[11] week_days                                      VARCHAR(32)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[12] enabled                                        INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [1]


JSON Sample
-------------------------------------
{    "id": "HVBuXALDsnulaUJSbmGaVtZjQ",    "created_by": "PJRXWAxqLVdiAEFLrNYhBpduH",    "created_time": 35,    "updated_by": "bfTaCvdJrkXojmTqcIlFcoJIn",    "updated_time": 5,    "rel_id": "lJwhxfLbPbOSTxnhWcBfAuFMI",    "rel_type": "jDTuCgyiHYvNYKVaATpIggfFp",    "start_time": "vxFYlLefVVwTccpheYUCRntwa",    "stop_time": "MELpfpQYWUBNaHbmvdYMkZjoi",    "start_date": "ZBohbfiRAZugKiBLnxfLAutlc",    "stop_date": "MIDaMQQuZcPVkwPbfMZnFMEDO",    "week_days": "qfnlTfAVPOguYRtDBOFxYifje",    "enabled": 51}



*/

var (
	Ebcp_schedule_job_FIELD_NAME_id = "id"

	Ebcp_schedule_job_FIELD_NAME_created_by = "created_by"

	Ebcp_schedule_job_FIELD_NAME_created_time = "created_time"

	Ebcp_schedule_job_FIELD_NAME_updated_by = "updated_by"

	Ebcp_schedule_job_FIELD_NAME_updated_time = "updated_time"

	Ebcp_schedule_job_FIELD_NAME_rel_id = "rel_id"

	Ebcp_schedule_job_FIELD_NAME_rel_type = "rel_type"

	Ebcp_schedule_job_FIELD_NAME_start_time = "start_time"

	Ebcp_schedule_job_FIELD_NAME_stop_time = "stop_time"

	Ebcp_schedule_job_FIELD_NAME_start_date = "start_date"

	Ebcp_schedule_job_FIELD_NAME_stop_date = "stop_date"

	Ebcp_schedule_job_FIELD_NAME_week_days = "week_days"

	Ebcp_schedule_job_FIELD_NAME_enabled = "enabled"
)

// Ebcp_schedule_job struct is a row record of the o_ebcp_schedule_job table in the  database
type Ebcp_schedule_job struct {
	ID string `json:"id"` //id

	CreatedBy string `json:"created_by"` //created_by

	CreatedTime common.LocalTime `json:"created_time"` //created_time

	UpdatedBy string `json:"updated_by"` //updated_by

	UpdatedTime common.LocalTime `json:"updated_time"` //updated_time

	RelID string `json:"rel_id"` //关联ID

	RelType string `json:"rel_type"` //关联类型,exhibition,room

	StartTime string `json:"start_time"` //启动时间,HH:mm

	StopTime string `json:"stop_time"` //停止时间,HH:mm

	StartDate string `json:"start_date"` //开始日期,yyyy-mm-dd

	StopDate string `json:"stop_date"` //停止日期,yyyy-mm-dd

	WeekDays string `json:"week_days"` //周几,逗号分隔,1-7代表周一到周日

	Enabled int32 `json:"enabled"` //是否启用(0: 禁用, 1: 启用)

}

var Ebcp_schedule_jobTableInfo = &TableInfo{
	Name: "o_ebcp_schedule_job",
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
			Name:               "rel_id",
			Comment:            `关联ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RelID",
			GoFieldType:        "string",
			JSONFieldName:      "rel_id",
			ProtobufFieldName:  "rel_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "rel_type",
			Comment:            `关联类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "RelType",
			GoFieldType:        "string",
			JSONFieldName:      "rel_type",
			ProtobufFieldName:  "rel_type",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "start_time",
			Comment:            `启动时间`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "start_date",
			Comment:            `开始日期`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "StartDate",
			GoFieldType:        "string",
			JSONFieldName:      "start_date",
			ProtobufFieldName:  "start_date",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "stop_date",
			Comment:            `停止日期`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "StopDate",
			GoFieldType:        "string",
			JSONFieldName:      "stop_date",
			ProtobufFieldName:  "stop_date",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "week_days",
			Comment:            `周几,逗号分隔,1-7代表周一到周日`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "WeekDays",
			GoFieldType:        "string",
			JSONFieldName:      "week_days",
			ProtobufFieldName:  "week_days",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "enabled",
			Comment:            `是否启用(0: 禁用, 1: 启用)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Enabled",
			GoFieldType:        "int32",
			JSONFieldName:      "enabled",
			ProtobufFieldName:  "enabled",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *Ebcp_schedule_job) TableName() string {
	return "o_ebcp_schedule_job"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *Ebcp_schedule_job) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *Ebcp_schedule_job) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *Ebcp_schedule_job) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *Ebcp_schedule_job) TableInfo() *TableInfo {
	return Ebcp_schedule_jobTableInfo
}
