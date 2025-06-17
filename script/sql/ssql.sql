create table if not exists usr_users
(
    id         bigint auto_increment
    primary key,
    username   varchar(36)                             null comment '用户名',
    token      varchar(2048) default ''                not null comment 'github token （加密存储）',
    email      varchar(128)  default ''                not null comment '邮箱',
    aes_key    varchar(2048) default ''                not null comment '对称加密密钥',
    gmt_create timestamp     default CURRENT_TIMESTAMP not null,
    gmt_update timestamp     default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP

    );

create unique index usr_users_username_index
    on usr_users (username);

create table push_info
(
    id         bigint auto_increment
        primary key,
    event_id   varchar(36)                            not null comment '事件唯一 ID',
    username   varchar(255)                           not null comment '用户名',
    repository varchar(255) default ''                not null,
    diff       text                                   not null,
    event_time timestamp                              not null,
    status     varchar(8)   default 'init'            not null comment '事件状态(init,success,fail)',
    gmt_create timestamp    default CURRENT_TIMESTAMP not null,
    gmt_update timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

create unique index push_info_event_id_uindex
    on push_info (event_id);

alter table push_info
    modify event_id varchar(64) not null comment '事件唯一 ID';

