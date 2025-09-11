
create TABLE o_ebcp_schedule_job (
    id VARCHAR(32) NOT NULL,
    created_by VARCHAR(32) NOT NULL,
    created_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(32) NOT NULL,
    updated_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    rel_id VARCHAR(32) NOT NULL,
    rel_type VARCHAR(32) NOT NULL,
    start_time VARCHAR(32) NOT NULL,
    stop_time VARCHAR(32) NOT NULL,
    start_date VARCHAR(32) NOT NULL DEFAULT '',
    stop_date VARCHAR(32) NOT NULL DEFAULT '',
    week_days VARCHAR(32) NOT NULL DEFAULT '',
    enabled INTEGER NOT NULL DEFAULT 1,
                         PRIMARY KEY (id)
);
comment on table o_ebcp_schedule_job is '定时任务表';
comment on column o_ebcp_schedule_job.rel_id is '关联ID';
comment on column o_ebcp_schedule_job.rel_type is '关联类型,exhibition,room';
comment on column o_ebcp_schedule_job.start_time is '启动时间,HH:mm';
comment on column o_ebcp_schedule_job.stop_time is '停止时间,HH:mm';
comment on column o_ebcp_schedule_job.start_date is '开始日期,yyyy-mm-dd';
comment on column o_ebcp_schedule_job.stop_date is '停止日期,yyyy-mm-dd';
comment on column o_ebcp_schedule_job.week_days is '周几,逗号分隔,1-7代表周一到周日';
comment on column o_ebcp_schedule_job.enabled is '是否启用(0: 禁用, 1: 启用)';



-- 创建视图
CREATE VIEW v_ebcp_exhibition_room_item_info AS
SELECT 
    e.id AS id,
    e.name AS name,
    COALESCE(e.start_time, '1970-01-01 00:00:00') AS start_time,
    COALESCE(e.end_time,   '1970-01-01 00:00:00') AS end_time,
    e.status AS status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_room WHERE exhibition_id = e.id) AS total_room_count,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE exhibition_id = e.id) AS total_item_count,
    (
        SELECT json_agg(
            json_build_object(
                'id', r.id,
                'name', r.name,
                'floor', r.floor,
                'floor_value', (SELECT dict_value FROM o_ops_dict WHERE id = r.floor),
                'floor_name', (SELECT dict_name FROM o_ops_dict WHERE id = r.floor),
                'location', r.location,
                'location_value', (SELECT dict_value FROM o_ops_dict WHERE id = r.location),
                'location_name', (SELECT dict_name FROM o_ops_dict WHERE id = r.location),
                'status', r.status,
                'items', (
                    SELECT json_agg(
                        json_build_object(
                            'id', i.id,
                            'name', i.name,
                            'type', i.type,
                            'sub_type', i.sub_type,
                            'status', i.status,
                            'remarks', i.remarks,
                            'commands', i.commands,
                            'ip_address', i.ip_address,
                            'port', i.port
                        )
                    )
                    FROM o_ebcp_exhibition_item i
                    WHERE i.room_id = r.id
                )
            )
        )
        FROM o_ebcp_exhibition_room r
        WHERE r.exhibition_id = e.id
    ) AS rooms
FROM o_ebcp_exhibition e;

COMMENT ON VIEW v_ebcp_exhibition_room_item_info IS '展览信息视图';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.id IS '展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.name IS '展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.start_time IS '展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.end_time IS '展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.status IS '展览状态（1: 运行中, 2: 筹备中, 3: 已结束）';    
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.total_room_count IS '展厅总数';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.total_item_count IS '展项总数';
COMMENT ON COLUMN v_ebcp_exhibition_room_item_info.rooms IS '展览使用的所有展厅（包含每个展厅内的展项信息）';


-- 展厅详细视图
CREATE OR REPLACE VIEW v_ebcp_exhibition_room_info AS
SELECT 
    er.id AS id,
    er.name AS name,
    er.floor AS floor,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.floor) AS floor_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.floor) AS floor_name,
    er.location AS location,
    (SELECT dict_value FROM o_ops_dict WHERE id = er.location) AS location_value,
    (SELECT dict_name FROM o_ops_dict WHERE id = er.location) AS location_name,
    er.status AS status,
    er.remarks AS remarks,
    e.id AS exhibition_id,
    e.name AS exhibition_name,
    COALESCE(e.start_time, '1970-01-01 00:00:00') AS exhibition_start_time,
    COALESCE(e.end_time, '1970-01-01 00:00:00') AS exhibition_end_time,
    e.status AS exhibition_status,
    (SELECT COUNT(*) FROM o_ebcp_exhibition_item WHERE room_id = er.id) AS item_count,
    json_agg(
        json_build_object(
            'id', ei.id,
            'name', ei.name,
            'type', ei.type,
            'status', ei.status,
            'remarks', ei.remarks,
            'commands', ei.commands,
            'ip_address', ei.ip_address,
            'port', ei.port
        )
    ) AS items,
    (
        SELECT json_agg(json_build_object(
            'device_id', cd.id,
            'device_name', cd.name,
            'device_type', cd.device_type,
            'status', cd.status,
            'ip_address', cd.ip_address,
            'port', cd.port,
            'version', cd.version,
            'commands', cd.commands
        ))
        FROM o_ebcp_control_device cd 
        WHERE cd.room_id = er.id
    ) AS control_devices    
FROM 
    o_ebcp_exhibition_room er
LEFT JOIN 
    o_ebcp_exhibition e ON er.exhibition_id = e.id
LEFT JOIN 
    o_ebcp_exhibition_item ei ON ei.room_id = er.id
GROUP BY 
    er.id, er.name, er.floor, er.location, er.status, er.remarks, 
    e.id, e.name, e.start_time, e.end_time, e.status;

COMMENT ON VIEW v_ebcp_exhibition_room_info IS '展厅详细视图，包含展厅信息及其关联的展览和展项信息（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.id IS '展厅ID';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.name IS '展厅名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor IS '展厅楼层';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor_value IS '展厅楼层值';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.floor_name IS '展厅楼层名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location IS '展厅位置';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location_value IS '展厅位置值';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.location_name IS '展厅位置名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.status IS '展厅状态';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.remarks IS '展厅备注';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_id IS '所属展览ID';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_name IS '所属展览名称';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_start_time IS '所属展览开始时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_end_time IS '所属展览结束时间';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.exhibition_status IS '所属展览状态';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.item_count IS '展厅内展项数量';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.items IS '展厅内的展项列表（JSON格式）';
COMMENT ON COLUMN v_ebcp_exhibition_room_info.control_devices IS '展厅内的中控设备列表（JSON格式）';
