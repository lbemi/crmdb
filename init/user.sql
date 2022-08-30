
SET FOREIGN_KEY_CHECKS=0;


create table user
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime                    not null,
    updated_at datetime                    not null,
    created_by bigint unsigned default '0' not null,
    updated_by bigint unsigned default '0' not null,
    memo       varchar(64)                 null,
    user_name  varchar(32)                 not null,
    real_name  varchar(32)                 null,
    password   char(128)                   not null,
    email      varchar(64)                 null,
    mobile     char(20)                    null,
    status     tinyint(1)                  not null
);


-- ----------------------------
-- Table structure for p_admin
-- ----------------------------
DROP TABLE IF EXISTS `p_admin`;
CREATE TABLE `p_admin` (
`admin_id` varchar(255) NOT NULL,
`realname` varchar(255) NOT NULL,
`mobile` varchar(255) NOT NULL,
`password` varchar(255) NOT NULL,
`image_url` varchar(255) DEFAULT NULL,
`create_time` bigint(20) NOT NULL,
`status` int(2) NOT NULL COMMENT '状态（1-启用 2-禁用）',
PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of p_admin
-- ----------------------------
INSERT INTO `p_admin` VALUES ('07d6e1f810f14da48e883d8c31c4dba8', '河南', '18600000001', 'e10adc3949ba59abbe56e057f20f883e', null, '10', '1592904335389', '1');
INSERT INTO `p_admin` VALUES ('25849d2c3e864573afb779a23e618a08', '萧炎', '13222602507', 'e10adc3949ba59abbe56e057f20f883e', null, '1', '1573021896640', '1');

-- ----------------------------
-- Table structure for p_menu
-- ----------------------------
DROP TABLE IF EXISTS `p_menu`;
CREATE TABLE `p_menu` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`icon` varchar(255) DEFAULT NULL COMMENT '小图标',
`menu_name` varchar(32) NOT NULL COMMENT '菜单名称',
`menu_url` varchar(32) NOT NULL COMMENT '菜单访问地址',
`file_name` varchar(32) NOT NULL COMMENT '文件名',
`parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级id（默认0）',
`order_id` int(11) NOT NULL COMMENT '排序',
`create_time` bigint(20) NOT NULL COMMENT '创建时间',
`status` int(11) NOT NULL COMMENT '状态（1-启用 2-停用）',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=126 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of p_menu
-- ----------------------------
INSERT INTO `p_menu` VALUES ('1', 'el-icon-setting', '系统配置', '', '', '0', '18', '1552288997274', '1');
INSERT INTO `p_menu` VALUES ('2', '', '字典管理', 'dictionary', 'Dictionary', '1', '1', '1552289059230', '1');
INSERT INTO `p_menu` VALUES ('3', null, '活动管理', 'active-manage', 'ActiveManage', '1', '9', '1552289458808', '1');
INSERT INTO `p_menu` VALUES ('4', null, '管理员管理', 'administrator-manage', 'AdministratorManage', '1', '3', '1552357679596', '1');
INSERT INTO `p_menu` VALUES ('5', null, '用户等级管理', 'userLevel', 'UserLevel', '1', '10', '1552357700208', '1');
INSERT INTO `p_menu` VALUES ('6', null, 'banner管理', 'banner-manage', 'BannerManage', '1', '8', '1552357760951', '1');


-- ----------------------------
-- Table structure for p_menu_btn
-- ----------------------------
DROP TABLE IF EXISTS `p_menu_btn`;
CREATE TABLE `p_menu_btn` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`menu_id` int(11) DEFAULT NULL,
`btn_name` varchar(32) DEFAULT NULL,
`btn_code` varchar(32) DEFAULT NULL COMMENT '按钮代码',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20260 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of p_menu_btn
-- ----------------------------
INSERT INTO `p_menu_btn` VALUES ('20000', '2', '添加', 'add');
INSERT INTO `p_menu_btn` VALUES ('20001', '2', '删除', 'del');
INSERT INTO `p_menu_btn` VALUES ('20002', '2', '修改', 'upd');
INSERT INTO `p_menu_btn` VALUES ('20003', '3', '添加', 'add');
INSERT INTO `p_menu_btn` VALUES ('20004', '3', '删除', 'del');
INSERT INTO `p_menu_btn` VALUES ('20005', '3', '修改', 'upd');
INSERT INTO `p_menu_btn` VALUES ('20009', '21', '添加', 'add');
INSERT INTO `p_menu_btn` VALUES ('20010', '21', '删除', 'del');
INSERT INTO `p_menu_btn` VALUES ('20011', '21', '修改', 'upd');
INSERT INTO `p_menu_btn` VALUES ('20012', '19', '添加', 'add');
INSERT INTO `p_menu_btn` VALUES ('20013', '19', '删除', 'del');
INSERT INTO `p_menu_btn` VALUES ('20014', '19', '修改', 'upd');
INSERT INTO `p_menu_btn` VALUES ('20015', '18', '添加', 'add');
INSERT INTO `p_menu_btn` VALUES ('20016', '18', '修改', 'upd');
INSERT INTO `p_menu_btn` VALUES ('20017', '18', '删除', 'del');
INSERT INTO `p_menu_btn` VALUES ('20018', '16', '添加', 'add');


-- ----------------------------
-- Table structure for p_role
-- ----------------------------
DROP TABLE IF EXISTS `p_role`;
CREATE TABLE `p_role` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`role_name` varchar(32) NOT NULL COMMENT '角色名称',
`remark` varchar(32) DEFAULT NULL COMMENT '描述',
`create_time` bigint(20) NOT NULL COMMENT '创建时间',
`status` int(2) NOT NULL COMMENT '状态（1-启用 2-禁用）',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of p_role
-- ----------------------------
INSERT INTO `p_role` VALUES ('1', '超级管理员', '啥权限都有', '1552289099975', '1');

INSERT INTO `p_role` VALUES ('10', '渠道管理员', '渠道管理员', '1573029091948', '1');

-- ----------------------------
-- Table structure for p_role_menu_btn
-- ----------------------------
DROP TABLE IF EXISTS `p_role_menu_btn`;
CREATE TABLE `p_role_menu_btn` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`role_id` int(11) NOT NULL,
`type` int(255) DEFAULT NULL COMMENT '类型（1-menu 2-button）',
`ref_id` int(11) DEFAULT NULL COMMENT '关联id（1-menuid 2-btnid）',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38531 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of p_role_menu_btn
-- ----------------------------

