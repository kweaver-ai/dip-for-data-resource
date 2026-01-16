create table if not exists `t_knowledge_network_info` (
    `m_id` bigint not null auto_increment comment '主键',
    `f_id` char(36) not null default '' comment '逻辑主键',
    `f_name` varchar(128) not null default '' comment '资源名称',
    `f_version` int not null default 0 comment '资源版本',
    `f_type` tinyint not null default 0 comment '资源类型；1:知识网络；2:数据源；3:图谱；4:图分析服务',
    `f_config_id` varchar(64) not null default '' comment '资源对应的配置ID，配置文件中定义',
    `f_real_id` varchar(36) not null default '' comment '资源在所属平台的ID，eg：知识网络在AD平台上的id',
    `f_created_at` datetime(3) not null default current_timestamp(3) comment '创建时间',
    `f_updated_at` datetime(3) not null default current_timestamp(3) comment '更新时间',
    `f_deleted_at` bigint not null default 0 comment '删除时间戳',
    unique key `ux_id_deleted_at`(`f_id`,`f_deleted_at`),
    PRIMARY KEY (`m_id`)
) comment 'ad相关资源信息表';

create table if not exists `t_knowledge_network_info_detail` (
    `m_id` bigint not null auto_increment comment '主键',
    `f_id` char(36) not null default '' comment '逻辑主键，与t_knowledge_network_info的f_id字段一致',
    `f_detail` text null  comment '详细信息',
    key `ix_id`(`f_id`),
    PRIMARY KEY (`m_id`)
) comment 'ad相关资源信息详情表';
