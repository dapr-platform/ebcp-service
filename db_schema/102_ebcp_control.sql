-- +goose Up
-- +goose StatementBegin

-- 插入展览信息
INSERT INTO o_ebcp_exhibition (
    id, created_by, created_time, updated_by, updated_time,
    name, start_time, end_time, remarks, hall_id, status
) VALUES
    (md5('百年展'), 'admin', NOW(), 'admin', NOW(), 
        '百年展', NOW(), NOW() + INTERVAL '10 year', NULL, md5('中国人民革命军事博物馆'), 1);

-- 插入展览控制项目
INSERT INTO o_ebcp_exhibition_item (
    id, created_by, created_time, updated_by, updated_time,
    name, exhibition_id, room_id, type, sub_type, status, commands, ip_address, port
) VALUES
    -- 百年展总控
    (md5('百年展总控'), 'admin', NOW(), 'admin', NOW(), '百年展总控', 
     md5('百年展'), md5('1F西侧'), 'static', 'master_control', 0, 
     '[{"name":"总控全开","type":"start","command":"F1 01 01"},{"name":"总控全关","type":"stop","command":"F1 01 02"},{"name":"总控互动不开","type":"control","command":"F1 01 03"}]',
     '10.7.66.3', 33333),
    
    -- 照明控制项目
    (md5('百年展-照明总控'), 'admin', NOW(), 'admin', NOW(), '百年展-照明总控', 
     md5('百年展'), md5('1F西侧'), 'static', 'lighting_control', 0, 
     '[{"name":"照明全开","type":"start","command":"F2 01 01"},{"name":"照明全关","type":"stop","command":"F2 01 02"}]',
     '10.7.66.3', 33333),
     
    (md5('百年展-一厅照明'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明', 
     md5('百年展'), md5('1F西侧'), 'static', 'lighting', 0, 
     '[{"name":"一厅照明全开","type":"start","command":"F2 02 01"},{"name":"一厅照明全关","type":"stop","command":"F2 02 04"}]',
     '10.7.66.3', 33333),
    
    (md5('百年展-二厅照明'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明', 
     md5('百年展'), md5('2F北侧'), 'static', 'lighting', 0, 
     '[{"name":"二厅照明全开","type":"start","command":"F2 02 02"},{"name":"二厅照明全关","type":"stop","command":"F2 02 05"}]',
     '10.7.66.3', 33333),
    
    (md5('百年展-三厅照明'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明', 
     md5('百年展'), md5('3F东南侧'), 'static', 'lighting', 0, 
     '[{"name":"三厅照明全开","type":"start","command":"F2 02 03"},{"name":"三厅照明全关","type":"stop","command":"F2 02 06"}]',
     '10.7.66.3', 33333),
    
    -- 设备电源控制项目
    (md5('百年展-设备总控'), 'admin', NOW(), 'admin', NOW(), '百年展-设备总控', 
     md5('百年展'), md5('1F西侧'), 'static', 'power_control', 0, 
     '[{"name":"设备全开","type":"start","command":"F3 01 01"},{"name":"设备全关","type":"stop","command":"F3 01 02"}]',
     '10.7.66.3', 33333),
    
    (md5('百年展-一厅设备电源'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源', 
     md5('百年展'), md5('1F西侧'), 'static', 'power', 0, 
     '[]',
     '10.7.66.3', 33333),
    
    (md5('百年展-二厅设备电源'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源', 
     md5('百年展'), md5('2F北侧'), 'static', 'power', 0, 
     '[]',
     '10.7.66.3', 33333),
    
    (md5('百年展-三厅设备电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源', 
     md5('百年展'), md5('3F东南侧'), 'static', 'power', 0, 
     '[]',
     '10.7.66.3', 33333),
    
    -- 具体设备控制项目
    (md5('百年展-一厅设备'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备', 
     md5('百年展'), md5('1F西侧'), 'static', 'device', 0, 
     '[{"name":"一厅设备全开","type":"start","command":"F3 02 01"},{"name":"一厅设备全关","type":"stop","command":"F3 02 04"}]', '10.7.66.3', 33333),
    
    (md5('百年展-二厅设备'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备', 
     md5('百年展'), md5('2F北侧'), 'static', 'device', 0, 
     '[{"name":"二厅设备全开","type":"start","command":"F3 02 02"},{"name":"二厅设备全关","type":"stop","command":"F3 02 05"}]', '10.7.66.3', 33333),
    
    (md5('百年展-三厅设备'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备', 
     md5('百年展'), md5('3F东南侧'), 'static', 'device', 0, 
     '[{"name":"三厅设备全开","type":"start","command":"F3 02 03"},{"name":"三厅设备全关","type":"stop","command":"F3 02 06"}]', '10.7.66.3', 33333);

-- 插入控制设备数据
INSERT INTO o_ebcp_control_device (
    id, created_by, created_time, updated_by, updated_time,
    name, device_type, item_id, room_id, status, ip_address, port, version, commands
) VALUES
    -- 百年展总控设备
    (md5('百年展-总控系统'), 'admin', NOW(), 'admin', NOW(), '百年展-总控系统', 'master_control', 
     md5('百年展总控'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"总控全开","type":"start","command":"F1 01 01"},{"name":"总控全关","type":"stop","command":"F1 01 02"},{"name":"总控互动不开","type":"control","command":"F1 01 03"}]'),
    
    -- 一厅照明设备
    (md5('百年展-一厅照明-序厅党徽'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-序厅党徽', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 01"},{"name":"关闭","type":"stop","command":"F4 02 01"}]'),
    
    (md5('百年展-一厅照明-序厅圆灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-序厅圆灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 02"},{"name":"关闭","type":"stop","command":"F4 02 02"}]'),
    
    (md5('百年展-一厅照明-序厅发光字'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-序厅发光字', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 03"},{"name":"关闭","type":"stop","command":"F4 02 03"}]'),
    
    (md5('百年展-一厅照明-前言轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-前言轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 04"},{"name":"关闭","type":"stop","command":"F4 02 04"}]'),
    
    (md5('百年展-一厅照明-训令轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-训令轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 05"},{"name":"关闭","type":"stop","command":"F4 02 05"}]'),
    
    (md5('百年展-一厅照明-军队建设轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-军队建设轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 06"},{"name":"关闭","type":"stop","command":"F4 02 06"}]'),
    
    (md5('百年展-一厅照明-灯箱+背景灯+站位灯箱'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-灯箱+背景灯+站位灯箱', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 07"},{"name":"关闭","type":"stop","command":"F4 02 07"}]'),
     
    (md5('百年展-一厅照明-北一展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-北一展柜', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 08"},{"name":"关闭","type":"stop","command":"F4 02 08"}]'),
     
    (md5('百年展-一厅照明-草地轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-草地轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 09"},{"name":"关闭","type":"stop","command":"F4 02 09"}]'),
     
    (md5('百年展-一厅照明-三湾轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-三湾轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 0A"},{"name":"关闭","type":"stop","command":"F4 02 0A"}]'),
     
    (md5('百年展-一厅照明-古田会议+军魂'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-古田会议+军魂', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 0B"},{"name":"关闭","type":"stop","command":"F4 02 0B"}]'),
     
    (md5('百年展-一厅照明-抗疫轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-抗疫轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
          '[{"name":"开启","type":"start","command":"F4 01 0C"},{"name":"关闭","type":"stop","command":"F4 02 0C"}]'),
     
    (md5('百年展-一厅照明-军魂轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-军魂轨道灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 0D"},{"name":"关闭","type":"stop","command":"F4 02 0D"}]'),
     
    (md5('百年展-一厅照明-砺三军灯箱'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-砺三军灯箱', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 0E"},{"name":"关闭","type":"stop","command":"F4 02 0E"}]'),
     
    (md5('百年展-一厅照明-砺三军展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-砺三军展柜', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 0F"},{"name":"关闭","type":"stop","command":"F4 02 0F"}]'),
     
    (md5('百年展-一厅照明-抗疫通柜'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-抗疫通柜', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 10"},{"name":"关闭","type":"stop","command":"F4 02 10"}]'),
     
    (md5('百年展-一厅照明-中间北展柜1'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-中间北展柜1', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 11"},{"name":"关闭","type":"stop","command":"F4 02 11"}]'),
     
    (md5('百年展-一厅照明-中间北展柜2'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-中间北展柜2', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 12"},{"name":"关闭","type":"stop","command":"F4 02 12"}]'),
     
    (md5('百年展-一厅照明-东+中间南展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-东+中间南展柜', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 13"},{"name":"关闭","type":"stop","command":"F4 02 13"}]'),
     
    (md5('百年展-一厅照明-门头灯'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅照明-门头灯', 'lighting', 
     md5('百年展-一厅照明'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 01 14"},{"name":"关闭","type":"stop","command":"F4 02 14"}]'),
     
    -- 一厅设备电源设备
    (md5('百年展-一厅设备电源-签名留言台电源+电视'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-签名留言台电源+电视', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 01"},{"name":"关闭","type":"stop","command":"F4 04 01"}]'),
    
    (md5('百年展-一厅设备电源-签名留言显示器'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-签名留言显示器', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 02"},{"name":"关闭","type":"stop","command":"F4 04 02"}]'),
    
    (md5('百年展-一厅设备电源-北二展柜+火电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-北二展柜+火电脑', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 03"},{"name":"关闭","type":"stop","command":"F4 04 03"}]'),
     
    (md5('百年展-一厅设备电源-三台电视'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-三台电视', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 04"},{"name":"关闭","type":"stop","command":"F4 04 04"}]'),
     
    (md5('百年展-一厅设备电源-南三电视'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-南三电视', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 05"},{"name":"关闭","type":"stop","command":"F4 04 05"}]'),
     
    (md5('百年展-一厅设备电源-草地投影'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备电源-草地投影', 'power', 
     md5('百年展-一厅设备电源'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F4 03 06"},{"name":"关闭","type":"stop","command":"F4 04 06"}]'),
    
    -- 二厅照明设备
    (md5('百年展-二厅照明-三大原则轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-三大原则轨道灯', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 01"},{"name":"关闭","type":"stop","command":"F5 02 01"}]'),
    
    (md5('百年展-二厅照明-十大政策轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-十大政策轨道灯', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 02"},{"name":"关闭","type":"stop","command":"F5 02 02"}]'),
    
    (md5('百年展-二厅照明-西北轨道灯1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西北轨道灯1', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 03"},{"name":"关闭","type":"stop","command":"F5 02 03"}]'),
    
    (md5('百年展-二厅照明-西北轨道灯2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西北轨道灯2', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 04"},{"name":"关闭","type":"stop","command":"F5 02 04"}]'),
    
    (md5('百年展-二厅照明-拉膜灯箱1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-拉膜灯箱1', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 05"},{"name":"关闭","type":"stop","command":"F5 02 05"}]'),
    
    (md5('百年展-二厅照明-拉膜灯箱2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-拉膜灯箱2', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 06"},{"name":"关闭","type":"stop","command":"F5 02 06"}]'),
    
    (md5('百年展-二厅照明-东北轨道灯1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东北轨道灯1', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 07"},{"name":"关闭","type":"stop","command":"F5 02 07"}]'),
    
    (md5('百年展-二厅照明-东北轨道灯2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东北轨道灯2', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 08"},{"name":"关闭","type":"stop","command":"F5 02 08"}]'),
    
    (md5('百年展-二厅照明-东北轨道灯3'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东北轨道灯3', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 09"},{"name":"关闭","type":"stop","command":"F5 02 09"}]'),
    
    (md5('百年展-二厅照明-东南轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东南轨道灯', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0A"},{"name":"关闭","type":"stop","command":"F5 02 0A"}]'),
    
    (md5('百年展-二厅照明-西南轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西南轨道灯', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0B"},{"name":"关闭","type":"stop","command":"F5 02 0B"}]'),
    
    (md5('百年展-二厅照明-灯箱灯带'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-灯箱灯带', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0C"},{"name":"关闭","type":"stop","command":"F5 02 0C"}]'),
    
    (md5('百年展-二厅照明-东北灯箱'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东北灯箱', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0D"},{"name":"关闭","type":"stop","command":"F5 02 0D"}]'),
    
    (md5('百年展-二厅照明-西南展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西南展柜', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0E"},{"name":"关闭","type":"stop","command":"F5 02 0E"}]'),
    
    (md5('百年展-二厅照明-西侧展柜2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西侧展柜2', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 0F"},{"name":"关闭","type":"stop","command":"F5 02 0F"}]'),
    
    (md5('百年展-二厅照明-中间通柜+攻克锦州'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-中间通柜+攻克锦州', 'lighting', 
     md5('百年展-二厅照明'), md5('12F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 10"},{"name":"关闭","type":"stop","command":"F5 02 10"}]'),
    
    (md5('百年展-二厅照明-西北展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-西北展柜', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 11"},{"name":"关闭","type":"stop","command":"F5 02 11"}]'),
    
    (md5('百年展-二厅照明-东北展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东北展柜', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 12"},{"name":"关闭","type":"stop","command":"F5 02 12"}]'),
    
    (md5('百年展-二厅照明-东南展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-东南展柜', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 13"},{"name":"关闭","type":"stop","command":"F5 02 13"}]'),
    
    (md5('百年展-二厅照明-拉膜灯箱'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅照明-拉膜灯箱', 'lighting', 
     md5('百年展-二厅照明'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 01 14"},{"name":"关闭","type":"stop","command":"F5 02 14"}]'),
    
    -- 二厅设备电源设备
    (md5('百年展-二厅设备电源-触摸屏电源'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-触摸屏电源', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 01"},{"name":"关闭","type":"stop","command":"F5 04 01"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里投影'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里投影', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 02"},{"name":"关闭","type":"stop","command":"F5 04 02"}]'),
    
    (md5('百年展-二厅设备电源-电视+一体机+展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-电视+一体机+展柜', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 03"},{"name":"关闭","type":"stop","command":"F5 04 03"}]'),
    
    (md5('百年展-二厅设备电源-二厅平板'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-二厅平板', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 04"},{"name":"关闭","type":"stop","command":"F5 04 04"}]'),
    
    (md5('百年展-二厅设备电源-四个电视'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-四个电视', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 05"},{"name":"关闭","type":"stop","command":"F5 04 05"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里灯箱1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里灯箱1', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 06"},{"name":"关闭","type":"stop","command":"F5 04 06"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里灯箱2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里灯箱2', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 07"},{"name":"关闭","type":"stop","command":"F5 04 07"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里灯箱3'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里灯箱3', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 08"},{"name":"关闭","type":"stop","command":"F5 04 08"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里灯箱4'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里灯箱4', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 09"},{"name":"关闭","type":"stop","command":"F5 04 09"}]'),
    
    (md5('百年展-二厅设备电源-决胜千里灯箱5'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-决胜千里灯箱5', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 0A"},{"name":"关闭","type":"stop","command":"F5 04 0A"}]'),
    
    (md5('百年展-二厅设备电源-背影灯'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备电源-背影灯', 'power', 
     md5('百年展-二厅设备电源'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F5 03 0B"},{"name":"关闭","type":"stop","command":"F5 04 0B"}]'),
    
    -- 三厅照明设备
    (md5('百年展-三厅照明-第二单元灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-第二单元灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 01"},{"name":"关闭","type":"stop","command":"F6 02 01"}]'),
    
    (md5('百年展-三厅照明-革命精神灯带'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-革命精神灯带', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 02"},{"name":"关闭","type":"stop","command":"F6 02 02"}]'),
    
    (md5('百年展-三厅照明-中厅灯带'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-中厅灯带', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 03"},{"name":"关闭","type":"stop","command":"F6 02 03"}]'),
    
    (md5('百年展-三厅照明-东方红轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-东方红轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 04"},{"name":"关闭","type":"stop","command":"F6 02 04"}]'),
    
    (md5('百年展-三厅照明-抗美援朝轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-抗美援朝轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 05"},{"name":"关闭","type":"stop","command":"F6 02 05"}]'),
    
    (md5('百年展-三厅照明-抢险救灾轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-抢险救灾轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 06"},{"name":"关闭","type":"stop","command":"F6 02 06"}]'),
    
    (md5('百年展-三厅照明-核试验灯带'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-核试验灯带', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 07"},{"name":"关闭","type":"stop","command":"F6 02 07"}]'),
    
    (md5('百年展-三厅照明-精兵之路轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-精兵之路轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 08"},{"name":"关闭","type":"stop","command":"F6 02 08"}]'),
    
    (md5('百年展-三厅照明-深圳基建兵灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-深圳基建兵灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 09"},{"name":"关闭","type":"stop","command":"F6 02 09"}]'),
    
    (md5('百年展-三厅照明-军事建设轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-军事建设轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0A"},{"name":"关闭","type":"stop","command":"F6 02 0A"}]'),
    
    (md5('百年展-三厅照明-基建兵灯带+灯箱'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-基建兵灯带+灯箱', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0B"},{"name":"关闭","type":"stop","command":"F6 02 0B"}]'),
    
    (md5('百年展-三厅照明-512轨道灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-512轨道灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0C"},{"name":"关闭","type":"stop","command":"F6 02 0C"}]'),
    
    (md5('百年展-三厅照明-港澳灯'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-港澳灯', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0D"},{"name":"关闭","type":"stop","command":"F6 02 0D"}]'),
    
    (md5('百年展-三厅照明-中厅展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-中厅展柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0E"},{"name":"关闭","type":"stop","command":"F6 02 0E"}]'),
    
    (md5('百年展-三厅照明-西南展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-西南展柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 0F"},{"name":"关闭","type":"stop","command":"F6 02 0F"}]'),
    
    (md5('百年展-三厅照明-西侧中展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-西侧中展柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 10"},{"name":"关闭","type":"stop","command":"F6 02 10"}]'),
    
    (md5('百年展-三厅照明-中间南通柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-中间南通柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 11"},{"name":"关闭","type":"stop","command":"F6 02 11"}]'),
    
    (md5('百年展-三厅照明-北侧展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-北侧展柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 12"},{"name":"关闭","type":"stop","command":"F6 02 12"}]'),
    
    (md5('百年展-三厅照明-东侧展柜'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅照明-东侧展柜', 'lighting', 
     md5('百年展-三厅照明'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 01 13"},{"name":"关闭","type":"stop","command":"F6 02 13"}]'),
    
    -- 三厅设备电源设备
    (md5('百年展-三厅设备电源-留言亭'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-留言亭', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 01"},{"name":"关闭","type":"stop","command":"F6 04 01"}]'),
    
    (md5('百年展-三厅设备电源-中厅一体机+电视电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-中厅一体机+电视电源', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 02"},{"name":"关闭","type":"stop","command":"F6 04 02"}]'),
    
    (md5('百年展-三厅设备电源-开国大典电视电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-开国大典电视电源', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 03"},{"name":"关闭","type":"stop","command":"F6 04 03"}]'),
    
    (md5('百年展-三厅设备电源-东侧3台电视'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-东侧3台电视', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 04"},{"name":"关闭","type":"stop","command":"F6 04 04"}]'),
    
    (md5('百年展-三厅设备电源-西侧3台电视+一体机'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-西侧3台电视+一体机', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 05"},{"name":"关闭","type":"stop","command":"F6 04 05"}]'),
    
    (md5('百年展-三厅设备电源-点歌台'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-点歌台', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 06"},{"name":"关闭","type":"stop","command":"F6 04 06"}]'),
    
    (md5('百年展-三厅设备电源-投影北电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-投影北电源', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 07"},{"name":"关闭","type":"stop","command":"F6 04 07"}]'),
    
    (md5('百年展-三厅设备电源-投影南电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-投影南电源', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 08"},{"name":"关闭","type":"stop","command":"F6 04 08"}]'),
    
    (md5('百年展-三厅设备电源-中厅投影电源'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备电源-中厅投影电源', 'power', 
     md5('百年展-三厅设备电源'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F6 03 09"},{"name":"关闭","type":"stop","command":"F6 04 09"}]'),
    
    -- 一厅设备
    (md5('百年展-一厅设备-在党的旗帜下前进电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-在党的旗帜下前进电脑', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 01"},{"name":"关闭","type":"stop","command":"F7 02 01"}]'),
    
    (md5('百年展-一厅设备-草地党支部电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地党支部电脑', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 02"},{"name":"关闭","type":"stop","command":"F7 02 02"}]'),
     
    (md5('百年展-一厅设备-草地火电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地火电脑', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 03"},{"name":"关闭","type":"stop","command":"F7 02 03"}]'),
    
    (md5('百年展-一厅设备-站位电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-站位电脑', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 04"},{"name":"关闭","type":"stop","command":"F7 02 04"}]'),

    (md5('百年展-一厅设备-签名留言台电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-签名留言台电脑', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 05"},{"name":"关闭","type":"stop","command":"F7 02 05"}]'),
   
    (md5('百年展-一厅设备-签名留言台1开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-签名留言台1开', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 06"},{"name":"关闭","type":"stop","command":"F7 01 06"}]'),
    
    (md5('百年展-一厅设备-签名留言台2开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-签名留言台2开', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 07"},{"name":"关闭","type":"stop","command":"F7 01 07"}]'),
    
    (md5('百年展-一厅设备-草地党支部投影1'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地党支部投影1', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 08"},{"name":"关闭","type":"stop","command":"F7 02 06"}]'),
    
    (md5('百年展-一厅设备-草地党支部投影2'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地党支部投影2', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 09"},{"name":"关闭","type":"stop","command":"F7 02 07"}]'),
    
    (md5('百年展-一厅设备-草地党支部投影3'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地党支部投影3', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0A"},{"name":"关闭","type":"stop","command":"F7 02 08"}]'),
    
    (md5('百年展-一厅设备-草地党支部投影4'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-草地党支部投影4', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0B"},{"name":"关闭","type":"stop","command":"F7 02 09"}]'),
    
    (md5('百年展-一厅设备-时序电源1'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-时序电源1', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0C"},{"name":"关闭","type":"stop","command":"F7 02 0A"}]'),
    
    (md5('百年展-一厅设备-时序电源2'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-时序电源2', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0D"},{"name":"关闭","type":"stop","command":"F7 02 0B"}]'),
    
    (md5('百年展-一厅设备-在党的旗帜下前进'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-在党的旗帜下前进', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0E"},{"name":"关闭","type":"stop","command":"F7 02 0C"}]'),
    
    (md5('百年展-一厅设备-战位LED'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-战位LED', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 0F"},{"name":"关闭","type":"stop","command":"F7 02 0D"}]'),
    
    (md5('百年展-一厅设备-抗疫电视开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-一厅设备-抗疫电视', 'device', 
     md5('百年展-一厅设备'), md5('1F西侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F7 01 10"},{"name":"关闭","type":"stop","command":"F7 01 10"},{"name":"确定","type":"stop","command":"F7 02 0E"}]'),
    -- 二厅设备
    (md5('百年展-二厅设备-决胜千里党指挥电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-决胜千里党指挥电脑', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 01"},{"name":"关闭","type":"stop","command":"F8 02 01"}]'),
    
    (md5('百年展-二厅设备-强军梦电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-强军梦电脑', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 02"},{"name":"关闭","type":"stop","command":"F8 02 02"}]'),
    
    (md5('百年展-二厅设备-铸盾砺剑电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-铸盾砺剑电脑', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 03"},{"name":"关闭","type":"stop","command":"F8 02 03"}]'),
    
    (md5('百年展-二厅设备-铸盾砺剑触摸屏1开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-铸盾砺剑触摸屏1', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 04"},{"name":"关闭","type":"stop","command":"F8 01 04"}]'),
    
    (md5('百年展-二厅设备-铸盾砺剑触摸屏2开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-铸盾砺剑触摸屏2', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 05"},{"name":"关闭","type":"stop","command":"F8 01 05"}]'),
    
    (md5('百年展-二厅设备-铸盾砺剑触摸屏3开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-铸盾砺剑触摸屏3', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 06"},{"name":"关闭","type":"stop","command":"F8 01 06"}]'),
    
    (md5('百年展-二厅设备-电报机开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-电报机', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 07"},{"name":"关闭","type":"stop","command":"F8 01 07"}]'),
    
    (md5('百年展-二厅设备-决胜千里党指挥投影1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-决胜千里党指挥投影1', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 08"},{"name":"关闭","type":"stop","command":"F8 02 04"}]'),
    
    (md5('百年展-二厅设备-决胜千里党指挥投影2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-决胜千里党指挥投影2', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 09"},{"name":"关闭","type":"stop","command":"F8 02 05"}]'),
    
    (md5('百年展-二厅设备-通电玻璃'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-通电玻璃', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0A"},{"name":"关闭","type":"stop","command":"F8 02 06"}]'),
    
    (md5('百年展-二厅设备-时序电源1'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-时序电源1', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0B"},{"name":"关闭","type":"stop","command":"F8 02 07"}]'),
    
    (md5('百年展-二厅设备-时序电源2'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-时序电源2', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0C"},{"name":"关闭","type":"stop","command":"F8 02 08"}]'),
    
    (md5('百年展-二厅设备-强军梦LED'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-强军梦LED', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0D"},{"name":"关闭","type":"stop","command":"F8 02 09"}]'),
    
    (md5('百年展-二厅设备-铸盾砺剑LED'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-铸盾砺剑LED', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0E"},{"name":"关闭","type":"stop","command":"F8 02 0A"}]'),
    
    (md5('百年展-二厅设备-攻克锦州电视开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-二厅设备-攻克锦州电视', 'device', 
     md5('百年展-二厅设备'), md5('2F北侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F8 01 0F"},{"name":"关闭","type":"stop","command":"F8 01 0F"}]'),
    
    -- 三厅设备
    (md5('百年展-三厅设备-老电影电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-老电影电脑', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 01"},{"name":"关闭","type":"stop","command":"F9 02 01"}]'),
    
    (md5('百年展-三厅设备-九八抗洪电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪电脑', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 02"},{"name":"关闭","type":"stop","command":"F9 02 02"}]'),
    
    (md5('百年展-三厅设备-留言亭电脑'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-留言亭电脑', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 03"},{"name":"关闭","type":"stop","command":"F9 02 03"}]'),
    
    (md5('百年展-三厅设备-革命精神答题1开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-革命精神答题1开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 04"},{"name":"关闭","type":"stop","command":"F9 01 04"}]'),
    
    (md5('百年展-三厅设备-革命精神答题2开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-革命精神答题2开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 05"},{"name":"关闭","type":"stop","command":"F9 01 05"}]'),
    
    (md5('百年展-三厅设备-革命精神答题3开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-革命精神答题3开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 06"},{"name":"关闭","type":"stop","command":"F9 01 06"}]'),
    
    (md5('百年展-三厅设备-革命精神答题4开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-革命精神答题4开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 07"},{"name":"关闭","type":"stop","command":"F9 01 07"}]'),
    
    (md5('百年展-三厅设备-开国将帅名录开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-开国将帅名录开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 08"},{"name":"关闭","type":"stop","command":"F9 01 08"}]'),
    
    (md5('百年展-三厅设备-老电影投影'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-老电影投影', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 09"},{"name":"关闭","type":"stop","command":"F9 02 04"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影南1'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影南1', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0A"},{"name":"关闭","type":"stop","command":"F9 02 05"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影南2'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影南2', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0B"},{"name":"关闭","type":"stop","command":"F9 02 06"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影南3'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影南3', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0C"},{"name":"关闭","type":"stop","command":"F9 02 07"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影北4'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影北4', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0D"},{"name":"关闭","type":"stop","command":"F9 02 08"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影北5'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影北5', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0E"},{"name":"关闭","type":"stop","command":"F9 02 09"}]'),
    
    (md5('百年展-三厅设备-九八抗洪投影北6'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-九八抗洪投影北6', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 0F"},{"name":"关闭","type":"stop","command":"F9 02 0A"}]'),
    
    (md5('百年展-三厅设备-时序电源1'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-时序电源1', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 10"},{"name":"关闭","type":"stop","command":"F9 02 0B"}]'),
    
    (md5('百年展-三厅设备-时序电源2'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-时序电源2', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 11"},{"name":"关闭","type":"stop","command":"F9 02 0C"}]'),
    
    (md5('百年展-三厅设备-抗美援朝英模开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-抗美援朝英模开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 12"},{"name":"关闭","type":"stop","command":"F9 01 12"},{"name":"确定","type":"stop","command":"F9 02 0D"}]'),
    
    (md5('百年展-三厅设备-深圳基建工程兵开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-深圳基建工程兵开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 13"},{"name":"关闭","type":"stop","command":"F9 01 13"},{"name":"确定","type":"stop","command":"F9 02 0E"}]'),
    
    (md5('百年展-三厅设备-从胜利走向胜利开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-从胜利走向胜利开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 14"},{"name":"关闭","type":"stop","command":"F9 01 14"}]'),
    
    (md5('百年展-三厅设备-维护海外权益开/关'), 'admin', NOW(), 'admin', NOW(), '百年展-三厅设备-维护海外权益开/关', 'device', 
     md5('百年展-三厅设备'), md5('3F东南侧'), 1, '10.7.66.3', 33333, '1.0', 
     '[{"name":"开启","type":"start","command":"F9 01 15"},{"name":"关闭","type":"stop","command":"F9 01 15"},{"name":"确定","type":"stop","command":"F9 02 0E"}]');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM o_ebcp_control_device WHERE ip_address = '10.7.66.3';
DELETE FROM o_ebcp_exhibition_item WHERE exhibition_id = md5('百年展');
DELETE FROM o_ebcp_exhibition WHERE name = '百年展';
-- +goose StatementEnd