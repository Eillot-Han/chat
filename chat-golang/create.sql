CREATE TABLE `user`
(
    `id`        INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `sex`       TINYINT(1)       NOT NULL DEFAULT '0' COMMENT '1男2女0未知',
    `name`      VARCHAR(32)      NOT NULL COMMENT '用户名',
    `password`  VARCHAR(32)      NOT NULL COMMENT '密码',
    `phone`     VARCHAR(11)      NOT NULL COMMENT '手机号码',
    `email`     VARCHAR(64)      NOT NULL COMMENT 'email',
    `sign_info` VARCHAR(128)     NOT NULL COMMENT '个性签名',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '用户';

CREATE TABLE `relation_ship`
(
    `id`       INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `ship_id`  INT(11)          NULL COMMENT '用户关系id',
    `small_id` INT(11)          NULL COMMENT '用户A的id',
    `big_id`   INT(11)          NULL COMMENT '用户B的id',
    `status`   TINYINT(1)       NULL DEFAULT '0' COMMENT '用户:0-正常, 1-删除',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '好友关系';

drop table `group`;
CREATE TABLE `group`
(
    `id`           INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `account`      INT(11)          NOT NULL COMMENT '群id',
    `name`         VARCHAR(256)     NOT NULL COMMENT '群名称',
    `creator`      INT(11)          NOT NULL DEFAULT '0' COMMENT '创建者用户id',
    `user_cnt`     INT(11)          NOT NULL DEFAULT '0' COMMENT '成员人数',
    `status`       TINYINT(3)       NOT NULL DEFAULT '0' COMMENT '是否删除,0-正常，1-删除',
    `last_chatted` varchar(100) COMMENT '最后聊天时间',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '群组';

drop table `group_member`;
CREATE TABLE `group_member`
(
    `id`                  INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`            INT(10)          NOT NULL COMMENT '群id',
    `member_id`           INT(10)          NOT NULL COMMENT '成员id',
    `member_unread`       INT(11)          NOT NULL DEFAULT '0' COMMENT '未读消息数量',
    `member_send`         INT(11)          NOT NULL DEFAULT '0' COMMENT '用户发送消息数量',
    `member_last_chatted` varchar(100) COMMENT '用户最后聊天时间',
    `permissions`         TINYINT(4)       NOT NULL DEFAULT '0' COMMENT '权限,0-正常，1-管理员，2-群主',
    `status`              TINYINT(3)       NOT NULL DEFAULT '0' COMMENT '是否删除,0-正常，1-删除',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '群组成员';

drop table `friend_message`;
create table friend_message
(
    id        BIGINT UNSIGNED auto_increment,
    relate_id INT                    not null comment '用户的关系id',
    from_id   INT                    not null comment '发送用户的id',
    to_id     INT                    not null comment '接收用户的id',
    content   VARCHAR(4096)          not null comment '消息内容',
    type      TINYINT(2) default '1' not null comment '消息类型 0文件 1文本',
    send_time varchar(100)           null comment '消息发送时间',
    status    TINYINT(1) default '0' not null comment '0正常 1被删除',
    primary key (id)
) comment '好友消息' engine = INNODB;


drop table `group_message`;
CREATE TABLE `group_message`
(
    `id`        BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `group_id`  INT(11)             NOT NULL COMMENT '用户的群关系id',
    `user_id`   INT(11)             NOT NULL COMMENT '发送用户的id',
    `content`   VARCHAR(4096)       NOT NULL COMMENT '消息内容',
    `type`      TINYINT(2)          NOT NULL DEFAULT '1' COMMENT '消息类型 0文件 1文本',
    `send_time` varchar(100)        NOT NULL COMMENT '消息发送时间',
    `status`    INT(11)             NOT NULL DEFAULT '0' COMMENT '0正常 1被删除',
    PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT '群消息';