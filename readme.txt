
创建USER表

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


创建community 表
create TABLE `community`(
                            `id` int(11) not null auto_increment,
                            `community_id` int(10) unsigned not null ,
                            `community_name` varchar(128) collate utf8mb4_general_ci not null ,
                            `introduction` varchar(256) collate utf8mb4_general_ci not null ,
                            `create_time` timestamp not null default current_timestamp,
                            `update_time` timestamp not null default current_timestamp on update CURRENT_TIMESTAMP,
                            primary key (`id`),
                            unique key `idx_community_id` (`community_id`),
                            unique key `idx_community_name` (`community_name`)
)engine=InnoDB default charset=utf8mb4 collate =utf8mb4_general_ci;

//插入数据
insert into `community` values ('1','1','go','golang','2012-01-01 01:01:01','2012-01-01 01:01:02');
insert into `community` values ('2','2','c','c++','2012-01-01 01:01:03','2012-01-01 01:01:04');
insert into `community` values ('3','3','java','java','2012-01-01 01:01:05','2012-01-01 01:01:06');
insert into `community` values ('4','4','rust','大神使用','2012-01-01 01:01:07','2012-01-01 01:01:08');

帖子表
drop TABLE if exists `post`;
create table `post`
(
    `id`           bigint(20)                               not null auto_increment,
    `post_id`      bigint(20)                               not null comment '帖子id',
    `title`        varchar(128) collate utf8mb4_general_ci  not null comment '标题',
    `content`      varchar(8192) collate utf8mb4_general_ci not null comment '内容',
    `author_id`    bigint(20)                               not null comment '作者的用户id',
    `community_id` bigint(20)                               not null comment '所属社区',
    `status`       tinyint(4)                               not null default '1' comment '帖子状态',
    `create_time`  timestamp                                null     default current_timestamp comment '创建时间',
    `update_time`  timestamp                                null     default current_timestamp on update current_timestamp comment '更新时间',
    primary key (`id`),
    unique key `idx_post_id` (`post_id`),
    key `idx_author_id` (`author_id`),
    key `idx_community_id` (`community_id`)
)engine=InnoDB DEFAULT charset = utf8mb4 collate = utf8mb4_general_ci


