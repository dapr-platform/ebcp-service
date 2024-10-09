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
                                       PRIMARY KEY (id),
                                       FOREIGN KEY (exhibition_hall_id) REFERENCES o_ebcp_exhibition_hall(id)
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
                                       exhibition_area_id VARCHAR(32) NOT NULL, -- 所属展区ID
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
COMMENT ON COLUMN o_ebcp_exhibition_item.exhibition_area_id IS '所属展区ID';
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

CREATE VIEW v_ebcp_exhibition_hall_details AS
SELECT 
    eh.id AS id,
    eh.name AS hall_name,
    eh.description AS hall_description,
    json_agg(
        json_build_object(
            'room_id', er.id,
            'room_name', er.name,
            'room_location', er.location,
            'room_status', er.status,
            'areas', (
                SELECT json_agg(
                    json_build_object(
                        'area_id', ea.id,
                        'area_name', ea.name,
                        'current_exhibition_name', ea.current_exhibition_name
                    )
                )
                FROM o_ebcp_exhibition_area ea
                WHERE ea.exhibition_room_id = er.id
            )
        )
    ) AS rooms
FROM 
    o_ebcp_exhibition_hall eh
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.exhibition_hall_id = eh.id
GROUP BY 
    eh.id, eh.name, eh.description;

COMMENT ON VIEW v_ebcp_exhibition_hall_details IS '展馆详细视图，包含展馆信息及其关联的展厅和展区信息（JSON格式）';

CREATE VIEW v_ebcp_exhibition_area_details AS
SELECT 
    ea.id AS id,
    ea.name AS area_name,
    ea.current_exhibition_name,
    ea.exhibition_room_id,
    json_agg(
        json_build_object(
            'item_id', ei.id,
            'item_name', ei.name,
            'item_status', ei.status,
            'item_type', ei.type,
            'item_remarks', ei.remarks
        )
    ) AS exhibition_items
FROM 
    o_ebcp_exhibition_area ea
LEFT JOIN 
    o_ebcp_exhibition_item ei ON ei.exhibition_area_id = ea.id
GROUP BY 
    ea.id, ea.name, ea.current_exhibition_name, ea.exhibition_room_id;

COMMENT ON VIEW v_ebcp_exhibition_area_details IS '展区详细视图，包含展区信息及其关联的所有展项信息（JSON格式），展项信息包括名字、状态、类型和备注';

-- Insert sample data for o_ebcp_device
INSERT INTO o_ebcp_device (id, created_by, created_time, updated_by, updated_time, name, type, control_interface, status)
VALUES 
('DEV001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Camera 1', 1, 'HTTP', 1),
('DEV002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Projector 1', 2, 'RS232', 1),
('DEV003', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Light 1', 3, 'DMX', 1);

-- Insert sample data for o_ebcp_exhibition_hall
INSERT INTO o_ebcp_exhibition_hall (id, created_by, created_time, updated_by, updated_time, name, description)
VALUES 
('HALL001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Main Exhibition Hall', 'Our primary exhibition space'),
('HALL002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Science Wing', 'Dedicated to scientific exhibits');

-- Insert sample data for o_ebcp_exhibition_room
INSERT INTO o_ebcp_exhibition_room (id, created_by, created_time, updated_by, updated_time, name, location, status, exhibition_hall_id, remarks)
VALUES 
('ROOM001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Dinosaur Room', 'First Floor', 1, 'HALL001', 'Features prehistoric exhibits'),
('ROOM002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Space Exploration', 'Second Floor', 1, 'HALL002', 'Showcases space technology');

-- Insert sample data for o_ebcp_exhibition_area
INSERT INTO o_ebcp_exhibition_area (id, created_by, created_time, updated_by, updated_time, name, current_exhibition_name, exhibition_room_id, remarks)
VALUES 
('AREA001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Jurassic Period', 'Dinosaurs of the Jurassic', 'ROOM001', 'Focus on Jurassic era dinosaurs'),
('AREA002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Mars Exploration', 'Journey to the Red Planet', 'ROOM002', 'Interactive Mars rover display');

-- Insert sample data for o_ebcp_exhibition_item
INSERT INTO o_ebcp_exhibition_item (id, created_by, created_time, updated_by, updated_time, name, status, type, remarks, exhibition_area_id)
VALUES 
('ITEM001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'T-Rex Skeleton', 1, 'static', 'Life-size replica', 'AREA001'),
('ITEM002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Mars Rover Model', 1, 'interactive', 'Interactive display', 'AREA002');

-- Insert sample data for o_ebcp_schedule_task
INSERT INTO o_ebcp_schedule_task (id, created_by, created_time, updated_by, updated_time, name, time_setting_id, action_id, status, remarks)
VALUES 
('TASK001', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Daily Opening', 'TIME001', 'ACTION001', 1, 'Tasks for opening the exhibition'),
('TASK002', 'SYSTEM', CURRENT_TIMESTAMP, 'SYSTEM', CURRENT_TIMESTAMP, 'Daily Closing', 'TIME002', 'ACTION002', 1, 'Tasks for closing the exhibition');

-- Insert sample data for o_ebcp_schedule_time
INSERT INTO o_ebcp_schedule_time (id, type, specific_time, repeat_pattern)
VALUES 
('TIME001', 2, NULL, 'DAILY'),
('TIME002', 2, NULL, 'DAILY');

-- Insert sample data for o_ebcp_schedule_action
INSERT INTO o_ebcp_schedule_action (id, action_type, target_id, operation_details)
VALUES 
('ACTION001', 1, 'DEV001', 'start_recording'),
('ACTION002', 2, 'DEV002', 'power_off');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP VIEW IF EXISTS v_ebcp_exhibition_hall_details CASCADE;
DROP VIEW IF EXISTS v_ebcp_exhibition_area_details CASCADE;

DROP TABLE IF EXISTS o_ebcp_exhibition_item CASCADE;
DROP TABLE IF EXISTS o_ebcp_exhibition_area CASCADE;
DROP TABLE IF EXISTS o_ebcp_exhibition_room CASCADE;
DROP TABLE IF EXISTS o_ebcp_exhibition_hall CASCADE;
DROP TABLE IF EXISTS o_ebcp_device CASCADE;
DROP TABLE IF EXISTS o_ebcp_schedule_task CASCADE;
DROP TABLE IF EXISTS o_ebcp_schedule_time CASCADE;
DROP TABLE IF EXISTS o_ebcp_schedule_action CASCADE;

-- +goose StatementEnd
