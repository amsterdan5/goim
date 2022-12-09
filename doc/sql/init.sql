CREATE DATABASE IF NOT EXISTS `im`;

CREATE TABLE IF NOT EXISTS `admin`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username` varchar(30) NOT NULL DEFAULT '' COMMENT '账号',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
    `last_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账号状态: 0:初始化, 1:正常, 2:禁用',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_username` (`username`),
    UNIQUE KEY `uni_mobile` (`mobile`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='管理员';

CREATE TABLE IF NOT EXISTS `user`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '账号id',
    `username` varchar(30) NOT NULL DEFAULT '' COMMENT '昵称',
    `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `email` varchar(30) NOT NULL DEFAULT '' COMMENT '邮箱',
    `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
    `public_key` text NOT NULL COMMENT '公钥',
    `personal_sign` varchar(255) NOT NULL DEFAULT '' COMMENT '个性签名',
    `total_used_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '在线时长',
    `last_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
    `last_login_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录时间',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账号状态: 0:初始化, 1:待激活, 2:已激活, 3:禁用, 4:已注销',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_mobile` (`mobile`),
    UNIQUE KEY `uni_email` (`email`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='用户';

CREATE TABLE IF NOT EXISTS `user_friend`
(
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账号id',
    `followee` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注者id',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`uid`, `followee`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='用户好友';

CREATE TABLE IF NOT EXISTS `user_black_friend`
(
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '账号id',
    `black_uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '黑名单uid',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`uid`, `black_uid`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='用户黑名单好友';

CREATE TABLE IF NOT EXISTS `group`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '群id',
    `name` varchar(30) NOT NULL DEFAULT '' COMMENT '群名',
    `remakr` varchar(255) NOT NULL DEFAULT '' COMMENT '群备注',
    `notice` text NOT NULL COMMENT '群通知',
    `is_banned` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群聊状态: 0:正常, 1:禁言',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群状态: 0:初始化, 1:正常, 2:解散',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='群';

CREATE TABLE IF NOT EXISTS `group_member`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群id',
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `is_manager` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否管理员, 0:否, 1:是',
    `banned_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '禁言截止时间',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 1:正常, 2:禁言',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_g_u` (`group_id`, `uid`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='群成员';

CREATE TABLE IF NOT EXISTS `notice`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
    `content_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '内容类型,0:普通消息,1:链接消息',
    `content` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
    `notice_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息类型, 0:系统消息,10:好友申请,11:群申请,12:群退出,20:群发,21:指定发送',
    `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 0:已创建, 1:已发送, 2:发送失败',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='消息通知';

CREATE TABLE IF NOT EXISTS `notice_member`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `notice_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息id',
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `is_read` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已读, 0:未读,1:已读',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_uid` (`uid`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='消息接收者';

CREATE TABLE IF NOT EXISTS `group_message`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `group_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '群id',
    `msg_type` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '消息类型,0:普通消息,1:新人入群,2:设置管理员',
    `content` varchar(255) NOT NULL DEFAULT '' COMMENT '内容',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_gid` (`group_id`),
    INDEX `idx_ctime` (`create_time`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='群消息';

CREATE TABLE IF NOT EXISTS `login_log`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '登录ip',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_uid` (`uid`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='登录日志';

CREATE TABLE IF NOT EXISTS `user_config`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `config_name` varchar(30) NOT NULL DEFAULT '' COMMENT '配置名',
    `value` varchar(255) NOT NULL DEFAULT '' COMMENT '值',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_uid_c` (`uid`, `config_name`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='用户配置';

CREATE TABLE IF NOT EXISTS `sys_config`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `config_name` varchar(30) NOT NULL DEFAULT '' COMMENT '配置名',
    `value` varchar(255) NOT NULL DEFAULT '' COMMENT '值',
    `create_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `update_by` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者id',
    `is_delete` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '已删除, 0:否, 1:是',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_config` (`config_name`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='系统配置';

CREATE TABLE IF NOT EXISTS `sys_operate_log`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `table` varchar(50) NOT NULL DEFAULT '' COMMENT '表名',
    `action` varchar(50) NOT NULL DEFAULT '' COMMENT '操作行为',
    `content` text NOT NULL COMMENT '内容',
    `ip` varchar(15) NOT NULL DEFAULT '' COMMENT 'ip',
    `admin_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者id',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_aid` (`admin_id`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='系统操作日志';

CREATE TABLE IF NOT EXISTS `user_operate_log`
(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '记录id',
    `table` varchar(50) NOT NULL DEFAULT '' COMMENT '表名',
    `action` varchar(50) NOT NULL DEFAULT '' COMMENT '操作行为',
    `content` text NOT NULL COMMENT '内容',
    `ip` varchar(15) NOT NULL DEFAULT '' COMMENT 'ip',
    `uid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
    `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX `idx_uid` (`uid`)
)ENGINE=Innodb DEFAULT CHARSET=UTF8mb4 COMMENT='用户操作日志';
