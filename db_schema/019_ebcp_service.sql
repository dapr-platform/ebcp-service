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
                                       current_exhibition_start_time TIMESTAMP, -- 当前展览开始时间
                                       current_exhibition_end_time TIMESTAMP, -- 当前展览结束时间
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
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_start_time IS '当前展览开始时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.current_exhibition_end_time IS '当前展览结束时间';
COMMENT ON COLUMN o_ebcp_exhibition_area.remarks IS '备注';

CREATE TABLE o_ebcp_exhibition_item(
                                       id VARCHAR(32) NOT NULL,
                                       created_by VARCHAR(32) NOT NULL,
                                       created_time TIMESTAMP NOT NULL,
                                       updated_by VARCHAR(32) NOT NULL,
                                       updated_time TIMESTAMP NOT NULL,
                                       name VARCHAR(255) NOT NULL, -- 展项名称
                                       exhibition_area_id VARCHAR(32) NOT NULL, -- 所属展区ID
                                       type VARCHAR(50), -- 展项类型（多媒体、static）
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
    ea.current_exhibition_start_time,
    ea.current_exhibition_end_time,
    ea.exhibition_room_id,
    er.location,
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
LEFT JOIN 
    o_ebcp_exhibition_room er ON er.id = ea.exhibition_room_id
GROUP BY 
    ea.id, ea.name, ea.current_exhibition_name, ea.exhibition_room_id,ea.current_exhibition_start_time,ea.current_exhibition_end_time, er.location;

COMMENT ON VIEW v_ebcp_exhibition_area_details IS '展区详细视图，包含展区信息及其关联的所有展项信息（JSON格式），展项信息包括名字、状态、类型和备注';

-- 插入展馆数据
INSERT INTO o_ebcp_exhibition_hall (id, created_by, created_time, updated_by, updated_time, name, description)
VALUES ('EH001', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '中国革命军事博物馆', '展示中国革命军事历史的重要场所');

-- 插入展厅数据
INSERT INTO o_ebcp_exhibition_room (id, created_by, created_time, updated_by, updated_time, name, location, exhibition_hall_id, status, remarks)
VALUES 
('ER001', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '第一展厅', '博物馆东翼', 'EH001', 1, '革命战争史'),
('ER002', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '第二展厅', '博物馆西翼', 'EH001', 1, '现代国防建设');

-- 插入展区数据
INSERT INTO o_ebcp_exhibition_area (id, created_by, created_time, updated_by, updated_time, name, exhibition_room_id, current_exhibition_name, current_exhibition_start_time, current_exhibition_end_time, remarks)
VALUES 
('EA001', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '一展区', 'ER001', '鸦片战争与太平天国运动', '2023-01-01', '2023-12-31', '19世纪中期重大历史事件'),
('EA002', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '二展区', 'ER001', '辛亥革命与北伐战争', '2023-01-01', '2023-12-31', '20世纪初期重要革命'),
('EA003', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '三展区', 'ER001', '抗日战争', '2023-01-01', '2023-12-31', '中国人民抗击日本侵略'),
('EA004', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '一展区', 'ER002', '解放战争', '2023-01-01', '2023-12-31', '国共内战时期'),
('EA005', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '二展区', 'ER002', '抗美援朝', '2023-01-01', '2023-12-31', '保家卫国的伟大战争'),
('EA006', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '三展区', 'ER002', '现代化国防建设', '2023-01-01', '2023-12-31', '中国国防现代化进程');

-- 插入展项数据
INSERT INTO o_ebcp_exhibition_item (id, created_by, created_time, updated_by, updated_time, name, exhibition_area_id, type, status, remarks)
VALUES 
-- 第一展厅一展区
('EI001', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '林则徐销烟', 'EA001', 'static', 1, '鸦片战争时期重要事件'),
('EI002', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '虎门销烟多媒体展示', 'EA001', 'interactive', 1, '虎门销烟场景重现'),
('EI003', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '第一次鸦片战争地图', 'EA001', 'static', 1, '战争进程地图'),
('EI004', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '太平天国运动领袖介绍', 'EA001', 'static', 1, '洪秀全等人物介绍'),
('EI005', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '天京事变影像资料', 'EA001', 'interactive', 1, '太平天国内部冲突'),

-- 第一展厅二展区
('EI006', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '武昌起义纪念碑', 'EA002', 'static', 1, '辛亥革命起点'),
('EI007', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '北伐战争路线图', 'EA002', 'interactive', 1, '互动式战争进程展示'),
('EI008', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '孙中山先生遗物', 'EA002', 'static', 1, '革命先驱生平展示'),
('EI009', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '黄埔军校成立影像', 'EA002', 'interactive', 1, '军事教育发展历程'),
('EI010', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '北伐军誓师大会场景', 'EA002', 'static', 1, '北伐战争重要时刻'),

-- 第一展厅三展区
('EI011', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '卢沟桥事变沙盘', 'EA003', 'static', 1, '全面抗战爆发'),
('EI012', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '平型关大捷多媒体', 'EA003', 'interactive', 1, '抗日战争重大胜利'),
('EI013', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '南京大屠杀遇难同胞纪念墙', 'EA003', 'static', 1, '勿忘国耻'),
('EI014', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '百团大战战役过程', 'EA003', 'interactive', 1, '抗日战争重要战役'),
('EI015', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '抗日英雄事迹展', 'EA003', 'static', 1, '杨靖宇等抗日英雄事迹'),

-- 第二展厅一展区
('EI016', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '渡江战役场景重现', 'EA004', 'interactive', 1, '解放战争关键战役'),
('EI017', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '解放军武器装备展', 'EA004', 'static', 1, '解放战争时期武器'),
('EI018', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '辽沈战役沙盘', 'EA004', 'static', 1, '三大战役之一'),
('EI019', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '淮海战役影像资料', 'EA004', 'interactive', 1, '决定性战役过程'),
('EI020', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '解放战争英雄事迹', 'EA004', 'static', 1, '刘伯承等将领事迹'),

-- 第二展厅二展区
('EI021', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '抗美援朝出国作战部队誓师', 'EA005', 'static', 1, '志愿军出征场景'),
('EI022', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '上甘岭战役多媒体', 'EA005', 'interactive', 1, '抗美援朝重要战役'),
('EI023', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '志愿军装备展示', 'EA005', 'static', 1, '抗美援朝时期武器装备'),
('EI024', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '抗美援朝英雄事迹', 'EA005', 'static', 1, '黄继光等英雄事迹'),
('EI025', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '停战协定签署场景', 'EA005', 'interactive', 1, '抗美援朝胜利'),

-- 第二展厅三展区
('EI026', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '现代化陆军装备展', 'EA006', 'static', 1, '陆军现代化成果'),
('EI027', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '海军发展历程', 'EA006', 'interactive', 1, '海军建设成就'),
('EI028', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '空军装备模型展', 'EA006', 'static', 1, '空军现代化装备'),
('EI029', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '火箭军发展历程', 'EA006', 'interactive', 1, '战略导弹部队建设'),
('EI030', 'admin', CURRENT_TIMESTAMP, 'admin', CURRENT_TIMESTAMP, '国防科技创新成果展', 'EA006', 'static', 1, '军事科技最新成果');




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
