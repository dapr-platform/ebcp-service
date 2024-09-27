-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE o_ebcp_device (
                               id VARCHAR(32) NOT NULL,
                               created_by VARCHAR(32) NOT NULL,
                               created_time TIMESTAMP NOT NULL,
                               updated_by VARCHAR(32) NOT NULL,
                               updated_time TIMESTAMP NOT NULL,
                               name VARCHAR(255) NOT NULL,
                               type INTEGER NOT NULL, -- 设备类型（摄像头、音视频、照明等）
                               control_interface VARCHAR(100) NOT NULL, -- 中控接口类型
                               status INTEGER NOT NULL, -- 状态（1: 正常, 2: 故障）
                               PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_device IS '设备表';
COMMENT ON COLUMN o_ebcp_device.id IS '设备唯一标识';
COMMENT ON COLUMN o_ebcp_device.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_device.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_device.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_device.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_device.name IS '设备名称';
COMMENT ON COLUMN o_ebcp_device.type IS '设备类型';
COMMENT ON COLUMN o_ebcp_device.control_interface IS '中控接口类型';
COMMENT ON COLUMN o_ebcp_device.status IS '设备状态';

CREATE TABLE o_ebcp_exhibition_hall (
                                        id VARCHAR(32) NOT NULL,
                                        created_by VARCHAR(32) NOT NULL,
                                        created_time TIMESTAMP NOT NULL,
                                        updated_by VARCHAR(32) NOT NULL,
                                        updated_time TIMESTAMP NOT NULL,
                                        name VARCHAR(255) NOT NULL, -- 展馆名称
                                        description TEXT, -- 展馆描述
                                        PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_hall IS '展馆表';
COMMENT ON COLUMN o_ebcp_exhibition_hall.id IS '展馆唯一标识';
COMMENT ON COLUMN o_ebcp_exhibition_hall.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_exhibition_hall.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_exhibition_hall.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_exhibition_hall.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_exhibition_hall.name IS '展馆名称';
COMMENT ON COLUMN o_ebcp_exhibition_hall.description IS '展馆描述';

CREATE TABLE o_ebcp_exhibition_room(
                                       id VARCHAR(32) NOT NULL,
                                       created_by VARCHAR(32) NOT NULL,
                                       created_time TIMESTAMP NOT NULL,
                                       updated_by VARCHAR(32) NOT NULL,
                                       updated_time TIMESTAMP NOT NULL,
                                       name VARCHAR(255) NOT NULL, -- 展厅名称
                                       location VARCHAR(255), -- 展厅位置
                                       exhibition_hall_id VARCHAR(32) NOT NULL, -- 所属展馆ID
                                       status INTEGER NOT NULL, -- 展厅状态（1: 正常, 2: 未使用, 3: 维修）
                                       remarks VARCHAR(255), -- 备注
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_room IS '展厅表';
COMMENT ON COLUMN o_ebcp_exhibition_room.id IS '展厅唯一标识';
COMMENT ON COLUMN o_ebcp_exhibition_room.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_exhibition_room.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_exhibition_room.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_exhibition_room.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_exhibition_room.name IS '展厅名称';
COMMENT ON COLUMN o_ebcp_exhibition_room.location IS '展厅位置';
COMMENT ON COLUMN o_ebcp_exhibition_room.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition_room.status IS '展厅状态';
COMMENT ON COLUMN o_ebcp_exhibition_room.remarks IS '备注';

CREATE TABLE o_ebcp_exhibition_area(
                                       id VARCHAR(32) NOT NULL,
                                       created_by VARCHAR(32) NOT NULL,
                                       created_time TIMESTAMP NOT NULL,
                                       updated_by VARCHAR(32) NOT NULL,
                                       updated_time TIMESTAMP NOT NULL,
                                       name VARCHAR(255) NOT NULL, -- 展区名称
                                       exhibition_room_id VARCHAR(32) NOT NULL, -- 所属展厅ID
                                       current_exhibition_name VARCHAR(255), -- 当前展览名称
                                       remarks VARCHAR(255), -- 备注
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_area IS '展区表';
COMMENT ON COLUMN o_ebcp_exhibition_area.id IS '展区唯一标识';
COMMENT ON COLUMN o_ebcp_exhibition_area.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_exhibition_area.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_exhibition_area.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.name IS '展区名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.exhibition_room_id IS '所属展厅ID';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_name IS '当前展览名称';
COMMENT ON COLUMN o_ebcp_exhibition_area.remarks IS '备注';

CREATE TABLE o_ebcp_exhibition_item(
                                       id VARCHAR(32) NOT NULL,
                                       created_by VARCHAR(32) NOT NULL,
                                       created_time TIMESTAMP NOT NULL,
                                       updated_by VARCHAR(32) NOT NULL,
                                       updated_time TIMESTAMP NOT NULL,
                                       name VARCHAR(255) NOT NULL, -- 展项名称
                                       exhibition_hall_id VARCHAR(32) NOT NULL, -- 所属展馆ID
                                       type VARCHAR(50), -- 展项类型（多媒体、静态）
                                       status INTEGER NOT NULL, -- 展项状态（1: 启动, 2: 停止, 3: 故障）
                                       remarks VARCHAR(255), -- 备注
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_exhibition_item IS '展项表';
COMMENT ON COLUMN o_ebcp_exhibition_item.id IS '展项唯一标识';
COMMENT ON COLUMN o_ebcp_exhibition_item.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_exhibition_item.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_exhibition_item.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_exhibition_item.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_exhibition_item.name IS '展项名称';
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_hall_id IS '所属展馆ID';
COMMENT ON COLUMN o_ebcp_exhibition_item.type IS '展项类型';
COMMENT ON COLUMN o_ebcp_exhibition_item.status IS '展项状态';
COMMENT ON COLUMN o_ebcp_exhibition_item.remarks IS '备注';

CREATE TABLE o_ebcp_schedule_task(
                                     id VARCHAR(32) NOT NULL, -- 调度任务唯一标识
                                     created_by VARCHAR(32) NOT NULL,
                                     created_time TIMESTAMP NOT NULL,
                                     updated_by VARCHAR(32) NOT NULL,
                                     updated_time TIMESTAMP NOT NULL,
                                     name VARCHAR(255) NOT NULL, -- 调度任务名称
                                     time_setting_id VARCHAR(32) NOT NULL, -- 关联时间配置
                                     action_id VARCHAR(32) NOT NULL, -- 关联动作表
                                     status INTEGER NOT NULL, -- 状态（1: 激活, 0: 停用）
                                     remarks VARCHAR(255), -- 备注
                                     PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_schedule_task IS '调度任务表';
COMMENT ON COLUMN o_ebcp_schedule_task.id IS '调度任务唯一标识';
COMMENT ON COLUMN o_ebcp_schedule_task.created_by IS '创建人';
COMMENT ON COLUMN o_ebcp_schedule_task.created_time IS '创建时间';
COMMENT ON COLUMN o_ebcp_schedule_task.updated_by IS '更新人';
COMMENT ON COLUMN o_ebcp_schedule_task.updated_time IS '更新时间';
COMMENT ON COLUMN o_ebcp_schedule_task.name IS '调度任务名称';
COMMENT ON COLUMN o_ebcp_schedule_task.time_setting_id IS '关联时间配置表';
COMMENT ON COLUMN o_ebcp_schedule_task.action_id IS '关联动作表';
COMMENT ON COLUMN o_ebcp_schedule_task.status IS '调度任务状态';
COMMENT ON COLUMN o_ebcp_schedule_task.remarks IS '备注';


CREATE TABLE o_ebcp_schedule_time(
                                     id VARCHAR(32) NOT NULL, -- 时间配置唯一标识
                                     type INTEGER NOT NULL, -- 时间类型 (1: 节假日, 2: 工作日, 3: 闭馆日)
                                     specific_time TIMESTAMP, -- 具体时间（可选）
                                     repeat_pattern VARCHAR(255), -- 重复模式（如每天, 每周）
                                     PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_schedule_time IS '时间配置表';
COMMENT ON COLUMN o_ebcp_schedule_time.id IS '时间配置唯一标识';
COMMENT ON COLUMN o_ebcp_schedule_time.type IS '时间类型';
COMMENT ON COLUMN o_ebcp_schedule_time.specific_time IS '具体时间';
COMMENT ON COLUMN o_ebcp_schedule_time.repeat_pattern IS '重复模式';

CREATE TABLE o_ebcp_schedule_action(
                                       id VARCHAR(32) NOT NULL, -- 动作唯一标识
                                       action_type INTEGER NOT NULL, -- 动作类型（1: 播放节目, 2: 切换节目, 3: 控制设备）
                                       target_id VARCHAR(32), -- 目标ID（设备ID或者展项ID）
                                       operation_details VARCHAR(255), -- 操作细节（如播放节目名称或设备控制指令）
                                       PRIMARY KEY (id)
);

COMMENT ON TABLE o_ebcp_schedule_action IS '调度动作表';
COMMENT ON COLUMN o_ebcp_schedule_action.id IS '动作唯一标识';
COMMENT ON COLUMN o_ebcp_schedule_action.action_type IS '动作类型';
COMMENT ON COLUMN o_ebcp_schedule_action.target_id IS '目标设备或展项';
COMMENT ON COLUMN o_ebcp_schedule_action.operation_details IS '操作细节';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS o_ebcp_device cascade ;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall cascade ;
DROP TABLE IF EXISTS o_ebcp_exhibition_room cascade ;
DROP TABLE IF EXISTS o_ebcp_exhibition_area cascade ;
DROP TABLE IF EXISTS o_ebcp_exhibition_item cascade ;
DROP TABLE IF EXISTS o_ebcp_schedule_task cascade ;
DROP TABLE IF EXISTS o_ebcp_schedule_time cascade ;
DROP TABLE IF EXISTS o_ebcp_schedule_action cascade ;



-- +goose StatementEnd
