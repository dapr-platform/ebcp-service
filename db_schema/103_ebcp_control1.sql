-- 插入控制设备数据 - 四厅新时代展馆
INSERT INTO o_ebcp_control_device (
    id, created_by, created_time, updated_by, updated_time,
    name, device_type, item_id, room_id, status, ip_address, port, version, commands
) VALUES
    -- 时序控制设备
    (md5('四厅总控-时序1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-时序1', 'timing_control', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 01"},{"name":"关闭","type":"stop","command":"F5 01 02"}]'),

    (md5('四厅总控-时序2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-时序2', 'timing_control', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 02 01"},{"name":"关闭","type":"stop","command":"F5 02 02"}]'),

    (md5('四厅总控-时序3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-时序3', 'timing_control', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 01"},{"name":"关闭","type":"stop","command":"F5 03 02"}]'),

    -- 投影设备
    (md5('四厅总控-投影1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-投影1', 'projector', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 01"},{"name":"关闭","type":"stop","command":"F6 01 02"}]'),

    (md5('四厅总控-投影2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-投影2', 'projector', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 02 01"},{"name":"关闭","type":"stop","command":"F6 02 02"}]'),

    -- 电脑设备
    (md5('四厅总控-海军'), 'admin', NOW(), 'admin', NOW(), '四厅总控-海军', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 01"},{"name":"关闭","type":"stop","command":"F7 01 02"}]'),

    (md5('四厅总控-陆军'), 'admin', NOW(), 'admin', NOW(), '四厅总控-陆军', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 02 01"},{"name":"关闭","type":"stop","command":"F7 02 02"}]'),

    (md5('四厅总控-空军'), 'admin', NOW(), 'admin', NOW(), '四厅总控-空军', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 03 01"},{"name":"关闭","type":"stop","command":"F7 03 02"}]'),

    (md5('四厅总控-寸土不让'), 'admin', NOW(), 'admin', NOW(), '四厅总控-寸土不让', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 04 01"},{"name":"关闭","type":"stop","command":"F7 04 02"}]'),

    (md5('四厅总控-沉浸式'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸式', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 05 01"},{"name":"关闭","type":"stop","command":"F7 05 02"}]'),

    (md5('四厅总控-致敬英雄'), 'admin', NOW(), 'admin', NOW(), '四厅总控-致敬英雄', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 06 01"},{"name":"关闭","type":"stop","command":"F7 06 02"}]'),

    (md5('四厅总控-军人荣耀'), 'admin', NOW(), 'admin', NOW(), '四厅总控-军人荣耀', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 07 01"},{"name":"关闭","type":"stop","command":"F7 07 02"}]'),

    (md5('四厅总控-阅兵号角'), 'admin', NOW(), 'admin', NOW(), '四厅总控-阅兵号角', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 08 01"},{"name":"关闭","type":"stop","command":"F7 08 02"}]'),

    (md5('四厅总控-拍照打卡'), 'admin', NOW(), 'admin', NOW(), '四厅总控-拍照打卡', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 09 01"},{"name":"关闭","type":"stop","command":"F7 09 02"}]'),

    (md5('四厅总控-沙盘-民兵训练基地'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沙盘-民兵训练基地', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0A 01"},{"name":"关闭","type":"stop","command":"F7 0A 02"}]'),

    (md5('四厅总控-联勤保障'), 'admin', NOW(), 'admin', NOW(), '四厅总控-联勤保障', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0B 01"},{"name":"关闭","type":"stop","command":"F7 0B 02"}]'),

    (md5('四厅总控-武警'), 'admin', NOW(), 'admin', NOW(), '四厅总控-武警', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0C 01"},{"name":"关闭","type":"stop","command":"F7 0C 02"}]'),

    (md5('四厅总控-东海防空'), 'admin', NOW(), 'admin', NOW(), '四厅总控-东海防空', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0D 01"},{"name":"关闭","type":"stop","command":"F7 0D 02"}]'),

    (md5('四厅总控-绕岛巡航'), 'admin', NOW(), 'admin', NOW(), '四厅总控-绕岛巡航', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0E 01"},{"name":"关闭","type":"stop","command":"F7 0E 02"}]'),

    (md5('四厅总控-习差三哨所'), 'admin', NOW(), 'admin', NOW(), '四厅总控-习差三哨所', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 0F 01"},{"name":"关闭","type":"stop","command":"F7 0F 02"}]'),

    (md5('四厅总控-国际合作集锦'), 'admin', NOW(), 'admin', NOW(), '四厅总控-国际合作集锦', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 10 01"},{"name":"关闭","type":"stop","command":"F7 10 02"}]'),

    (md5('四厅总控-国际军事比赛'), 'admin', NOW(), 'admin', NOW(), '四厅总控-国际军事比赛', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 11 01"},{"name":"关闭","type":"stop","command":"F7 11 02"}]'),

    (md5('四厅总控-军队脱贫'), 'admin', NOW(), 'admin', NOW(), '四厅总控-军队脱贫', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 12 01"},{"name":"关闭","type":"stop","command":"F7 12 02"}]'),

    (md5('四厅总控-为了人民'), 'admin', NOW(), 'admin', NOW(), '四厅总控-为了人民', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 13 01"},{"name":"关闭","type":"stop","command":"F7 13 02"}]'),

    (md5('四厅总控-站岗1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-站岗1', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 14 01"},{"name":"关闭","type":"stop","command":"F7 14 02"}]'),

    (md5('四厅总控-站岗2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-站岗2', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 15 01"},{"name":"关闭","type":"stop","command":"F7 15 02"}]'),

    (md5('四厅总控-站岗3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-站岗3', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 16 01"},{"name":"关闭","type":"stop","command":"F7 16 02"}]'),

    (md5('四厅总控-站岗4'), 'admin', NOW(), 'admin', NOW(), '四厅总控-站岗4', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 17 01"},{"name":"关闭","type":"stop","command":"F7 17 02"}]'),

    (md5('四厅总控-模拟射击1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-模拟射击1', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 18 01"},{"name":"关闭","type":"stop","command":"F7 18 02"}]'),

    (md5('四厅总控-模拟射击2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-模拟射击2', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 19 01"},{"name":"关闭","type":"stop","command":"F7 19 02"}]'),

    (md5('四厅总控-抗疫'), 'admin', NOW(), 'admin', NOW(), '四厅总控-抗疫', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 1A 01"},{"name":"关闭","type":"stop","command":"F7 1A 02"}]'),

    (md5('四厅总控-双拥共建模范'), 'admin', NOW(), 'admin', NOW(), '四厅总控-双拥共建模范', 'computer', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 1B 01"},{"name":"关闭","type":"stop","command":"F7 1B 02"}]'),

    -- 灯光电源设备
    (md5('四厅总控-火箭军场景射灯02'), 'admin', NOW(), 'admin', NOW(), '四厅总控-火箭军场景射灯02', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 01"},{"name":"关闭","type":"stop","command":"F8 01 02"}]'),

    (md5('四厅总控-火箭军场景射灯03'), 'admin', NOW(), 'admin', NOW(), '四厅总控-火箭军场景射灯03', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 02 01"},{"name":"关闭","type":"stop","command":"F8 02 02"}]'),

    (md5('四厅总控-祖国在我心射灯02'), 'admin', NOW(), 'admin', NOW(), '四厅总控-祖国在我心射灯02', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 03 01"},{"name":"关闭","type":"stop","command":"F8 03 02"}]'),

    (md5('四厅总控-一起站岗+军队共建村射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-一起站岗+军队共建村射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 04 01"},{"name":"关闭","type":"stop","command":"F8 04 02"}]'),

    (md5('四厅总控-维和部队射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-维和部队射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 05 01"},{"name":"关闭","type":"stop","command":"F8 05 02"}]'),

    (md5('四厅总控-尾厅大好河山洗墙灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-尾厅大好河山洗墙灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 06 01"},{"name":"关闭","type":"stop","command":"F8 06 02"}]'),

    (md5('四厅总控-尾厅射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-尾厅射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 07 01"},{"name":"关闭","type":"stop","command":"F8 07 02"}]'),

    (md5('四厅总控-人民至上射灯1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-人民至上射灯1', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 08 01"},{"name":"关闭","type":"stop","command":"F8 08 02"}]'),

    (md5('四厅总控-人民至上射灯2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-人民至上射灯2', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 01"},{"name":"关闭","type":"stop","command":"F9 01 02"}]'),

    (md5('四厅总控-海岛巡逻射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-海岛巡逻射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 02 01"},{"name":"关闭","type":"stop","command":"F9 02 02"}]'),

    (md5('四厅总控-深入推进备战打仗射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深入推进备战打仗射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 03 01"},{"name":"关闭","type":"stop","command":"F9 03 02"}]'),

    (md5('四厅总控-火箭军场景展柜'), 'admin', NOW(), 'admin', NOW(), '四厅总控-火箭军场景展柜', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 04 01"},{"name":"关闭","type":"stop","command":"F9 04 02"}]'),

    (md5('四厅总控-军民共建村展柜'), 'admin', NOW(), 'admin', NOW(), '四厅总控-军民共建村展柜', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 05 01"},{"name":"关闭","type":"stop","command":"F9 05 02"}]'),

    (md5('四厅总控-人民至上展柜'), 'admin', NOW(), 'admin', NOW(), '四厅总控-人民至上展柜', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 06 01"},{"name":"关闭","type":"stop","command":"F9 06 02"}]'),

    (md5('四厅总控-尾厅灯箱'), 'admin', NOW(), 'admin', NOW(), '四厅总控-尾厅灯箱', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 07 01"},{"name":"关闭","type":"stop","command":"F9 07 02"}]'),

    (md5('四厅总控-沉浸拉膜1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸拉膜1', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 08 01"},{"name":"关闭","type":"stop","command":"F9 08 02"}]'),

    (md5('四厅总控-沉浸拉膜2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸拉膜2', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 01 01"},{"name":"关闭","type":"stop","command":"FA 01 02"}]'),

    (md5('四厅总控-沉浸拉膜3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸拉膜3', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 02 01"},{"name":"关闭","type":"stop","command":"FA 02 02"}]'),

    (md5('四厅总控-沉浸拉膜4'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸拉膜4', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 03 01"},{"name":"关闭","type":"stop","command":"FA 03 02"}]'),

    (md5('四厅总控-深情励三军射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 04 01"},{"name":"关闭","type":"stop","command":"FA 04 02"}]'),

    (md5('四厅总控-深情励三军天花1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花1', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 05 01"},{"name":"关闭","type":"stop","command":"FA 05 02"}]'),

    (md5('四厅总控-深情励三军天花2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花2', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 06 01"},{"name":"关闭","type":"stop","command":"FA 06 02"}]'),

    (md5('四厅总控-深情励三军天花3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花3', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 07 01"},{"name":"关闭","type":"stop","command":"FA 07 02"}]'),

    (md5('四厅总控-深情励三军天花4'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花4', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FA 08 01"},{"name":"关闭","type":"stop","command":"FA 08 02"}]'),

    (md5('四厅总控-深情励三军天花5'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花5', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 01 01"},{"name":"关闭","type":"stop","command":"FB 01 02"}]'),

    (md5('四厅总控-深情励三军天花6'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军天花6', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 02 01"},{"name":"关闭","type":"stop","command":"FB 02 02"}]'),

    (md5('四厅总控-深入推进打仗灯箱'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深入推进打仗灯箱', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 03 01"},{"name":"关闭","type":"stop","command":"FB 03 02"}]'),

    (md5('四厅总控-履行新时代使命任务灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-履行新时代使命任务灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 04 01"},{"name":"关闭","type":"stop","command":"FB 04 02"}]'),

    (md5('四厅总控-备用'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 05 01"},{"name":"关闭","type":"stop","command":"FB 05 02"}]'),

    (md5('四厅总控-礼序乾坤乐和天地灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-礼序乾坤乐和天地灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 06 01"},{"name":"关闭","type":"stop","command":"FB 06 02"}]'),

    (md5('四厅总控-火箭军地台+新时代使命灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-火箭军地台+新时代使命灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 07 01"},{"name":"关闭","type":"stop","command":"FB 07 02"}]'),

    (md5('四厅总控-人民至上生命至上地台灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-人民至上生命至上地台灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FB 08 01"},{"name":"关闭","type":"stop","command":"FB 08 02"}]'),

    (md5('四厅总控-深情励三军灯箱'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军灯箱', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FC 01 01"},{"name":"关闭","type":"stop","command":"FC 01 02"}]'),

    (md5('四厅总控-深情励三军楼梯地台灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军楼梯地台灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FC 02 01"},{"name":"关闭","type":"stop","command":"FC 02 02"}]'),

    (md5('四厅总控-深情励三军立墙灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-深情励三军立墙灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FC 03 01"},{"name":"关闭","type":"stop","command":"FC 03 02"}]'),

    (md5('四厅总控-卫国戍边地台灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-卫国戍边地台灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FC 04 01"},{"name":"关闭","type":"stop","command":"FC 04 02"}]'),

    (md5('四厅总控-礼序乾坤乐和天地射灯+展柜'), 'admin', NOW(), 'admin', NOW(), '四厅总控-礼序乾坤乐和天地射灯+展柜', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FC 05 01"},{"name":"关闭","type":"stop","command":"FC 05 02"}]'),

    (md5('四厅总控-沉浸场景拉膜1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜1', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 01 01"},{"name":"关闭","type":"stop","command":"FD 01 02"}]'),

    (md5('四厅总控-沉浸场景拉膜2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜2', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 02 01"},{"name":"关闭","type":"stop","command":"FD 02 02"}]'),

    (md5('四厅总控-沉浸场景拉膜3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜3', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 03 01"},{"name":"关闭","type":"stop","command":"FD 03 02"}]'),

    (md5('四厅总控-沉浸场景拉膜4'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜4', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 04 01"},{"name":"关闭","type":"stop","command":"FD 04 02"}]'),

    (md5('四厅总控-沉浸场景拉膜5'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜5', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 05 01"},{"name":"关闭","type":"stop","command":"FD 05 02"}]'),

    (md5('四厅总控-沉浸场景拉膜6'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜6', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 06 01"},{"name":"关闭","type":"stop","command":"FD 06 02"}]'),

    (md5('四厅总控-沉浸场景拉膜7'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜7', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 07 01"},{"name":"关闭","type":"stop","command":"FD 07 02"}]'),

    (md5('四厅总控-沉浸场景拉膜8'), 'admin', NOW(), 'admin', NOW(), '四厅总控-沉浸场景拉膜8', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FD 08 01"},{"name":"关闭","type":"stop","command":"FD 08 02"}]'),

    (md5('四厅总控-火箭军场景射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-火箭军场景射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 01 01"},{"name":"关闭","type":"stop","command":"FE 01 02"}]'),

    (md5('四厅总控-预留电源1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-预留电源1', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 02 01"},{"name":"关闭","type":"stop","command":"FE 02 02"}]'),

    (md5('四厅总控-预留电源2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-预留电源2', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 03 01"},{"name":"关闭","type":"stop","command":"FE 03 02"}]'),

    (md5('四厅总控-奖杯上方射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-奖杯上方射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 04 01"},{"name":"关闭","type":"stop","command":"FE 04 02"}]'),

    (md5('四厅总控-献花上方射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-献花上方射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 05 01"},{"name":"关闭","type":"stop","command":"FE 05 02"}]'),

    (md5('四厅总控-祖国在我心射灯'), 'admin', NOW(), 'admin', NOW(), '四厅总控-祖国在我心射灯', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 06 01"},{"name":"关闭","type":"stop","command":"FE 06 02"}]'),

    (md5('四厅总控-祖国在我心地台灯带'), 'admin', NOW(), 'admin', NOW(), '四厅总控-祖国在我心地台灯带', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 07 01"},{"name":"关闭","type":"stop","command":"FE 07 02"}]'),

    (md5('四厅总控-礼序乾坤和天地灯箱'), 'admin', NOW(), 'admin', NOW(), '四厅总控-礼序乾坤和天地灯箱', 'lighting', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FE 08 01"},{"name":"关闭","type":"stop","command":"FE 08 02"}]'),

    (md5('四厅总控-备用1'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用1', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 01 01"},{"name":"关闭","type":"stop","command":"FF 01 02"}]'),

    (md5('四厅总控-备用2'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用2', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 02 01"},{"name":"关闭","type":"stop","command":"FF 02 02"}]'),

    (md5('四厅总控-备用3'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用3', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 03 01"},{"name":"关闭","type":"stop","command":"FF 03 02"}]'),

    (md5('四厅总控-备用4'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用4', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 04 01"},{"name":"关闭","type":"stop","command":"FF 04 02"}]'),

    (md5('四厅总控-备用5'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用5', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 05 01"},{"name":"关闭","type":"stop","command":"FF 05 02"}]'),

    (md5('四厅总控-备用6'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用6', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 06 01"},{"name":"关闭","type":"stop","command":"FF 06 02"}]'),

    (md5('四厅总控-备用7'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用7', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 07 01"},{"name":"关闭","type":"stop","command":"FF 07 02"}]'),

    (md5('四厅总控-备用8'), 'admin', NOW(), 'admin', NOW(), '四厅总控-备用8', 'power', 
     'fs0DR6y4JXouDoTH23GbJ', 'TNloGfa6KxA9hEqp2sEBk', 1, '10.7.65.14', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"FF 08 01"},{"name":"关闭","type":"stop","command":"FF 08 02"}]');
