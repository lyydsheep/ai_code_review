create table if not exists usr_users(
    id       bigint auto_increment
        primary key,
    username varchar(36)              null comment '用户名',
    token    varchar(2048) default '' not null comment 'github token （加密存储）',
    email    varchar(128)  default '' not null comment '邮箱',
    aes_key  varchar(2048) default '' not null comment '对称加密密钥'
);

alter table usr_users
    add gmt_create timestamp default CURRENT_TIMESTAMP not null;

alter table usr_users
    add gmt_update timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP;

create unique index usr_users_username_index
    on usr_users (username);

