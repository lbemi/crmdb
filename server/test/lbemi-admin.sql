/*
 Navicat Premium Data Transfer

 Source Server         : tx-测试服务器-mysql8.0.30
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : 42.192.54.11:3333
 Source Schema         : lbemi-admin

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 08/06/2023 14:14:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for clusters
-- ----------------------------
DROP TABLE IF EXISTS `clusters`;
CREATE TABLE `clusters` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '集群名称',
  `version` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'kubernetes版本',
  `kube_config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'cloud config',
  `run_time` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '运行时',
  `service_cidr` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'service cloud ip',
  `pod_cidr` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'pod id',
  `cni` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'cni网络插件',
  `proxy_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '网络模式（iptables，ipvs)',
  `status` tinyint(1) DEFAULT NULL COMMENT '集群状态',
  `nodes` bigint DEFAULT NULL COMMENT '节点数量',
  `internal_ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cpu` double DEFAULT NULL,
  `memory` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of clusters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for hosts
-- ----------------------------
DROP TABLE IF EXISTS `hosts`;
CREATE TABLE `hosts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `label` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `port` bigint NOT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `auth_method` tinyint DEFAULT NULL,
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `secret` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint(1) NOT NULL,
  `enable_ssh` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of hosts
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for log_login
-- ----------------------------
DROP TABLE IF EXISTS `log_login`;
CREATE TABLE `log_login` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `username` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户名',
  `status` varchar(2) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'ip地址',
  `loginLocation` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '固件',
  `loginTime` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `remark` text COLLATE utf8mb4_general_ci,
  `msg` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of log_login
-- ----------------------------
BEGIN;
INSERT INTO `log_login` VALUES (107, '2023-06-07 03:17:43', '2023-06-07 03:17:43', 'lbemi_admin', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 03:17:43', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '登录成功');
INSERT INTO `log_login` VALUES (108, '2023-06-07 03:21:35', '2023-06-07 03:21:35', 'lbemi_admin', '1', '172.17.0.1:49414', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 03:21:35', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '退出登录');
INSERT INTO `log_login` VALUES (109, '2023-06-07 03:21:46', '2023-06-07 03:21:46', 'lbemi', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 03:21:46', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '登录成功');
INSERT INTO `log_login` VALUES (110, '2023-06-07 03:32:21', '2023-06-07 03:32:21', 'lbemi', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 03:32:21', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '退出登录');
INSERT INTO `log_login` VALUES (111, '2023-06-07 03:32:30', '2023-06-07 03:32:30', 'lbemi_admin', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 03:32:30', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '登录成功');
INSERT INTO `log_login` VALUES (112, '2023-06-07 06:23:21', '2023-06-07 06:23:21', 'lbemi', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-07 06:23:21', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '登录成功');
INSERT INTO `log_login` VALUES (113, '2023-06-08 02:28:05', '2023-06-08 02:28:05', 'lbemi', '1', '125.121.227.29', '', 'Edge:113.0.1774.50', 'Intel Mac OS X 10_15_7', 'Macintosh', '2023-06-08 02:28:05', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.50', '登录成功');
COMMIT;

-- ----------------------------
-- Table structure for log_operator
-- ----------------------------
DROP TABLE IF EXISTS `log_operator`;
CREATE TABLE `log_operator` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `title` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作的模块',
  `businessType` varchar(2) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '04其它 01新增 02修改 03删除',
  `method` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作人员',
  `url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作url',
  `ip` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作IP',
  `location` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '操作地点',
  `param` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求参数',
  `status` smallint DEFAULT NULL COMMENT '请求状态吗',
  `errMsg` text COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of log_operator
-- ----------------------------
BEGIN;
INSERT INTO `log_operator` VALUES (320, '2023-06-07 03:17:26', '2023-06-07 03:17:26', 'operatorLog', '03', 'DELETE', 'lbemi_admin', '/api/v1/logs/operator/all', '125.121.227.29', '', '', 200, '');
INSERT INTO `log_operator` VALUES (321, '2023-06-07 03:17:30', '2023-06-07 03:17:30', 'logout', '01', 'POST', 'lbemi_admin', '/api/v1/users/logout', '125.121.227.29', '', '', 200, '');
INSERT INTO `log_operator` VALUES (322, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 'roles', '01', 'POST', 'lbemi_admin', '/api/v1/roles/5439801874/menus', '125.121.227.29', '', '', 200, '');
INSERT INTO `log_operator` VALUES (323, '2023-06-07 03:21:35', '2023-06-07 03:21:35', 'logout', '01', 'POST', 'lbemi_admin', '/api/v1/users/logout', '125.121.227.29', '', '', 200, '');
INSERT INTO `log_operator` VALUES (324, '2023-06-07 03:32:21', '2023-06-07 03:32:21', 'logout', '01', 'POST', 'lbemi', '/api/v1/users/logout', '125.121.227.29', '', '', 200, '');
INSERT INTO `log_operator` VALUES (325, '2023-06-07 03:33:11', '2023-06-07 03:33:11', 'loginLog', '03', 'DELETE', 'lbemi_admin', '/api/v1/logs/login', '125.121.227.29', '', '', 200, '');
COMMIT;

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(1:启用 2:不启用)',
  `memo` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `parentID` bigint unsigned NOT NULL COMMENT '父级ID',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '菜单URL',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `redirect` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `sequence` bigint NOT NULL COMMENT '排序值',
  `menuType` tinyint(1) NOT NULL COMMENT '菜单类型(1 左侧菜单,2 按钮, 3 非展示权限)',
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'icon图标',
  `method` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作类型 none/GET/POST/PUT/DELETE',
  `code` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `isLink` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `isHide` tinyint(1) DEFAULT NULL,
  `isKeepAlive` tinyint(1) DEFAULT NULL,
  `isAffix` tinyint(1) DEFAULT NULL,
  `isIframe` tinyint(1) DEFAULT NULL,
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '标题',
  `group` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `isK8S` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1667235409 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of menus
-- ----------------------------
BEGIN;
INSERT INTO `menus` VALUES (1667235310, '2022-10-29 15:26:46', '2023-04-01 17:08:11', 1, '系统设置', 0, '/system', 'layout/routerView/parent', '/system/menu', 'system', 97, 1, 'iconfont icon-xitongshezhi', '', '', '', 0, 1, 0, 0, 'message.router.system', '', 0);
INSERT INTO `menus` VALUES (1667235311, '2022-10-29 15:32:06', '2023-03-29 17:08:00', 1, '用户管理', 1667235310, '/system/user', 'system/user/index', NULL, 'systemuser', 10, 1, 'iconfont icon-icon-', '', '', NULL, 0, 1, 0, 0, 'message.router.systemUser', '', 0);
INSERT INTO `menus` VALUES (1667235312, '2022-10-29 15:32:29', '2023-03-30 20:16:54', 1, '角色管理', 1667235310, '/system/role', 'system/role/index', NULL, 'systemRole', 9, 1, 'iconfont icon-icon-', '', '', '', 0, 0, 0, 0, 'message.router.systemRole', '', 0);
INSERT INTO `menus` VALUES (1667235313, '2022-10-29 15:33:01', '2022-11-08 11:34:51', 1, '菜单权限', 1667235310, '/system/menu', 'system/menu/index', NULL, 'systemMenu', 8, 1, 'iconfont icon-caidan', '', '', NULL, 0, 0, 0, 0, 'message.router.systemMenu', '', 0);
INSERT INTO `menus` VALUES (1667235314, '2022-10-29 15:35:33', '2023-04-01 17:07:51', 1, '首页', 0, '/home', 'home/index', '', 'home', 100, 1, 'iconfont icon-shouye', '', '', '', 0, 1, 1, 0, 'message.router.home', '', 0);
INSERT INTO `menus` VALUES (1667235315, '2022-10-29 15:40:32', '2023-06-06 13:43:19', 1, '添加用户按钮', 1667235311, '/api/v1/users/register', '', '', '添加用户', 10, 2, '', 'POST', 'sys:user:add', '', 0, 1, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235316, '2022-10-29 15:41:28', '2023-06-06 13:45:22', 1, '删除用户按钮', 1667235311, '/api/v1/users/:id', '', '', '删除用户按钮', 8, 2, '', 'DELETE', 'sys:user:del', '', 0, 1, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235318, '2022-10-29 15:43:33', '2023-06-06 15:32:14', 1, '查看基础权限', 1667235311, '/api/v1/user*|/api/v1/role*', '', '', '查看基础权限', 11, 3, '', 'GET', '', '', 0, 1, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235319, '2022-10-29 15:50:04', '2023-06-06 13:43:44', 1, '查看角色基础权限', 1667235312, '/api/v1/role*|/api/v1/menu*', '', '', '查看角色基础权限', 10, 3, '', 'GET', '', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235320, '2022-10-29 15:50:42', '2023-06-06 13:43:59', 1, '查看菜单基础权限', 1667235313, '/api/v1/menus*', '', '', '查看菜单基础权限', 10, 3, '', 'GET', '', '', 0, 1, 0, 0, '', '菜单', 0);
INSERT INTO `menus` VALUES (1667235321, '2022-10-29 15:52:23', '2023-06-06 13:45:01', 1, '添加角色按钮权限', 1667235312, '/api/v1/roles', '', '', '添加角色按钮权限', 9, 2, '', 'POST', 'sys:role:add', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235322, '2022-10-29 15:52:53', '2023-06-06 13:48:37', 1, '编辑角色按钮权限', 1667235312, '/api/v1/roles/:id', '', '', '编辑角色', 8, 2, '', 'PUT', 'sys:role:edit', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235323, '2022-10-29 15:53:19', '2023-06-06 13:54:10', 1, '删除角色按钮', 1667235312, '/api/v1/roles/:id', '', '', '删除角色', 7, 2, '', 'DELETE', 'sys:role:del', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235324, '2022-10-29 15:54:36', '2023-06-06 13:50:01', 1, '角色分配权限按钮', 1667235312, '/api/v1/roles/:id/menus', '', '', '分配权限按钮', 7, 2, '', 'POST', 'sys:role:set', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235325, '2022-10-29 15:58:43', '2023-06-06 13:51:40', 1, '添加权限按钮权限按钮', 1667235313, '/api/v1/menus', '', '', '添加权限按钮权限按钮', 7, 2, '', 'POST', 'sys:menu:add', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235326, '2022-10-29 15:59:58', '2023-06-06 13:52:25', 1, '删除权限/菜单/按钮', 1667235313, '/api/v1/menus/:id', '', '', '删除权限/菜单/按钮', 7, 2, '', 'DELETE', 'sys:menu:del', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235327, '2022-10-29 16:00:18', '2023-06-06 13:52:59', 1, '编辑/按钮/权限/按钮', 1667235313, '/api/v1/menus/:id', '', '', '编辑/按钮/权限/按钮', 7, 2, '', 'PUT', 'sys:menu:edit', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235328, '2022-10-29 16:01:15', '2023-06-06 13:54:35', 1, '分配用户角色按钮', 1667235311, '/api/v1/users/:id/roles', '', '', '分配用户角色', 6, 2, '', 'POST', 'sys:user:set', '', 0, 1, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235333, '2022-11-06 14:29:05', '2023-06-06 16:22:54', 1, '允许所有接口POST', 0, '/api/v1/*', '', '', '基础权限', 1, 3, '', 'POST', '', '', 0, 1, 0, 0, '', '特殊', 0);
INSERT INTO `menus` VALUES (1667235334, '2022-11-06 16:27:57', '2023-06-06 13:59:29', 1, '菜单/权限/按钮启用禁用', 1667235313, '/api/v1/menus/:id/status/:status', '', '', '菜单/权限/按钮启用禁用', 1, 2, '', 'PUT', 'sys:menu:status', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235335, '2022-11-07 17:44:07', '2023-06-06 13:55:13', 1, '用户启用停用', 1667235311, '/api/v1/users/:id/status/:status', '', '', '用户启用停用', 1, 2, '', 'PUT', 'sys:user:status', '', 0, 1, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235336, '2022-11-07 18:10:33', '2023-06-06 14:19:02', 1, '启用禁用', 1667235312, '/api/v1/roles/:id/status/:status', '', '', '启用禁用', 1, 2, '', 'PUT', 'sys:role:status', '', 0, 1, 0, 0, '', '角色', 0);
INSERT INTO `menus` VALUES (1667235338, '2022-11-08 13:44:48', '2023-03-30 18:08:36', 2, '资产管理', 0, '/asset', 'layout/routerView/parent', NULL, '资产管理', 4, 1, 'iconfont icon-icon-', '', '', NULL, 0, 1, 0, 0, NULL, '', 0);
INSERT INTO `menus` VALUES (1667235339, '2022-11-08 13:45:39', '2023-05-14 15:18:13', 1, 'kubernetes容器管理', 0, '/kubernetes', 'layout/routerView/parent', '/kubernetes/cluster', 'kubernetes', 98, 1, 'iconfont icon-kubernetes', '', '', '', 0, 1, 0, 0, 'message.router.kubernetes', '', 0);
INSERT INTO `menus` VALUES (1667235340, '2022-11-08 13:47:00', '2022-11-16 14:28:34', 2, '主机管理', 1667235338, '/asset/host', 'asset/host', NULL, '主机管理', 1, 1, 'iconfont icon-icon-', '', '', NULL, 0, 1, 0, 0, NULL, '', 0);
INSERT INTO `menus` VALUES (1667235341, '2022-11-09 10:14:15', '2023-06-06 14:19:12', 1, '添加主机', 1667235340, '/api/v1/host', '', '', '添加', 1, 2, '', 'POST', 'asset:host:add', '', 0, 1, 0, 0, '', '主机', 0);
INSERT INTO `menus` VALUES (1667235342, '2022-11-09 10:14:56', '2023-06-06 13:55:47', 1, '删除主机', 1667235340, '/api/v1/host/:id', '', '', '删除主机', 1, 2, '', 'DELETE', 'asset:host:del', '', 0, 1, 0, 0, '', '主机', 0);
INSERT INTO `menus` VALUES (1667235343, '2022-11-09 10:16:01', '2023-06-06 13:56:01', 1, 'Terminal', 1667235340, '/api/v1/host/:id/ws', '', '', 'SSH', 1, 2, '', 'GET', 'asset:host:ssh', '', 0, 1, 0, 0, '', '主机', 0);
INSERT INTO `menus` VALUES (1667235344, '2022-11-09 10:16:36', '2023-06-06 16:22:33', 1, '查看基础权限', 1667235340, '/api/v1/host*', '', '', '查看基础权限', 1, 3, '', 'GET', '', '', 0, 1, 0, 0, '', '主机', 0);
INSERT INTO `menus` VALUES (1667235346, '2022-12-07 17:49:15', '2023-05-14 15:23:32', 1, 'deployment', 1667235395, '/kubernetes/workload/deployment', 'kubernetes/deployment/index', '', 'k8sDeployment', 10, 1, '', '', '', '', 0, 0, 0, 0, 'message.router.k8sDeployment', '', 0);
INSERT INTO `menus` VALUES (1667235347, '2022-12-27 16:11:11', '2023-06-06 13:57:11', 2, '删除', 1667235346, '/namespace/:name', '', '', '删除', 1, 2, '', 'DELETE', 'k8s:deploy:delete', '', 0, 1, 0, 0, '', 'deployment', 0);
INSERT INTO `menus` VALUES (1667235348, '2023-03-21 20:12:40', '2023-05-14 15:14:55', 1, 'pods', 1667235395, '/kubernetes/workload/pod', 'kubernetes/pod/index', '', 'k8sPod', 1, 1, '', '', '', '', 0, 0, 0, 0, 'message.router.k8sPod', '', 1);
INSERT INTO `menus` VALUES (1667235355, '2023-03-29 22:36:47', '2023-03-30 10:33:28', 1, 'API管理', 1667235310, '/system/commService', 'system/commService/index', NULL, 'systemApi', 1, 1, 'iconfont icon-shouye_dongtaihui', '', '', '', 0, 0, 0, 0, 'message.router.systemApi', '', 0);
INSERT INTO `menus` VALUES (1667235357, '2023-03-30 21:18:40', '2023-03-30 21:18:40', 1, 'limit', 0, '/limits', 'layout/routerView/parent', NULL, 'limits', 1, 1, 'iconfont icon-quanxian', '', '', '', 0, 0, 0, 0, 'message.router.limits', '', 0);
INSERT INTO `menus` VALUES (1667235358, '2023-03-30 21:20:28', '2023-03-30 21:20:28', 1, 'BackEnd', 1667235357, '/limits/backEnd', 'layout/routerView/parent', NULL, 'limitsBackEnd', 1, 1, '', '', '', '', 0, 1, 0, 0, 'message.router.limitsBackEnd', '', 0);
INSERT INTO `menus` VALUES (1667235359, '2023-03-30 21:22:49', '2023-03-30 21:22:49', 1, 'limitsBackEndEndPage', 1667235358, '/limits/backEnd/page', 'limits/backEnd/page/index', NULL, 'limitsBackEndEndPage', 1, 1, '', '', '', '', 0, 1, 0, 0, 'message.router.limitsBackEndEndPage', '', 0);
INSERT INTO `menus` VALUES (1667235360, '2023-03-30 21:24:26', '2023-03-30 21:56:02', 1, '功能', 0, '/fun', 'layout/routerView/parent', NULL, 'funIndex', 1, 1, 'iconfont icon-crew_feature', '', '', '', 0, 1, 0, 0, 'message.router.funIndex', '', 0);
INSERT INTO `menus` VALUES (1667235361, '2023-03-30 21:25:48', '2023-03-30 21:25:48', 1, 'funTagsView', 1667235360, '/fun/tagsView', 'fun/tagsView/index', NULL, 'funTagsView', 1, 1, 'ele-Bell', '', '', '', 0, 1, 0, 0, 'message.router.funTagsView', '', 0);
INSERT INTO `menus` VALUES (1667235362, '2023-03-30 21:29:23', '2023-03-30 21:29:23', 1, '大屏', 0, '/chart', 'chart/index', NULL, 'chartIndex', 8, 1, 'iconfont icon-ico_shuju', '', '', '', 0, 0, 0, 0, 'message.router.chartIndex', '', 0);
INSERT INTO `menus` VALUES (1667235363, '2023-03-30 21:30:53', '2023-06-02 21:00:27', 1, '个人中心', 0, '/personal', 'personal/index', '', 'personal', 6, 1, 'iconfont icon-gerenzhongxin', '', '', '', 0, 0, 0, 0, 'message.router.personal', '', 0);
INSERT INTO `menus` VALUES (1667235364, '2023-03-30 21:31:53', '2023-03-30 21:31:53', 1, '工具中心', 0, '/tools', 'ools/index', NULL, 'tools', 6, 1, 'iconfont icon-gongju', '', '', '', 0, 0, 0, 0, 'message.router.tools', '', 0);
INSERT INTO `menus` VALUES (1667235365, '2023-03-30 21:33:17', '2023-03-30 21:33:17', 1, '外链接', 0, '/link', 'layout/routerView/link', NULL, 'layoutLinkView', 6, 1, 'iconfont icon-caozuo-wailian', '', '', 'https://www.baidu.com', 0, 0, 0, 0, 'message.router.layoutLinkView', '', 0);
INSERT INTO `menus` VALUES (1667235366, '2023-03-30 21:34:46', '2023-04-14 13:54:53', 1, 'grafana', 0, '/iframes', 'layout/routerView/iframe', '', 'layoutIfameView', 6, 1, 'iconfont icon-neiqianshujuchucun', '', '', 'http://192.168.3.120:3000', 0, 0, 0, 1, 'message.router.grafana', '', 0);
INSERT INTO `menus` VALUES (1667235367, '2023-03-30 22:03:09', '2023-03-31 16:54:19', 1, '部门管理', 1667235310, '/system/dept', 'system/dept/index', NULL, 'systemDept', 44, 1, 'ele-Basketball', '', '', '', 0, 0, 0, 0, 'message.router.systemDept', '', 0);
INSERT INTO `menus` VALUES (1667235368, '2023-03-30 22:05:56', '2023-03-31 16:53:29', 1, '字典管理', 1667235310, '/system/dic', 'system/dic/index', NULL, 'systemDic', 44, 1, 'ele-Basketball', '', '', '', 0, 0, 0, 0, 'message.router.systemDic', '', 0);
INSERT INTO `menus` VALUES (1667235369, '2023-03-30 22:14:22', '2023-03-30 22:14:22', 1, '数字滚动', 1667235360, '/fun/countup', 'fun/countup/index', NULL, 'funCountup', 1, 1, '', '', '', '', 0, 1, 0, 0, 'message.router.funCountup', '', 0);
INSERT INTO `menus` VALUES (1667235370, '2023-03-30 22:15:41', '2023-03-30 22:15:41', 1, '数字滚动', 1667235360, '/fun/wangEditor', 'fun/wangEditor/index', NULL, 'funWangEditor', 1, 1, 'ele-AlarmClock', '', '', '', 0, 1, 0, 0, 'message.router.funWangEditor', '', 0);
INSERT INTO `menus` VALUES (1667235372, '2023-03-31 14:55:13', '2023-06-06 13:58:03', 1, '添加菜单按钮权限', 1667235355, '/api/v1/menu', '', '', '添加菜单按钮权限', 1, 2, '', 'POST', 'sys:commService:add', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235373, '2023-03-31 15:03:23', '2023-06-06 13:58:26', 1, '删除API', 1667235355, '/api/v1/menu/:id', '', '', '删除API', 1, 2, '', 'DELETE', 'sys:commService:del', '', 0, 1, 0, 0, '', '通用', 0);
INSERT INTO `menus` VALUES (1667235375, '2023-03-31 22:17:26', '2023-05-14 15:11:21', 1, 'dashboard', 1667235339, '/kubernetes/dashboard', 'kubernetes/dashboard/index', '', 'kubernetesDashboard', 89, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.kubernetesDashboard', '', 1);
INSERT INTO `menus` VALUES (1667235376, '2023-03-31 22:45:49', '2023-05-14 15:11:53', 1, 'kubernetes集群', 1667235339, '/kubernetes/cluster', 'kubernetes/cluster/index', '', 'kubernetesCluster', 92, 1, '', 'GET', '', '', 0, 1, 0, 0, 'message.router.kubernetesCluster', '', 1);
INSERT INTO `menus` VALUES (1667235377, '2023-04-01 16:45:47', '2023-05-14 15:11:05', 1, 'Node节点', 1667235339, '/kubernetes/node', 'kubernetes/node/index', '', 'k8sNode', 81, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sNode', '', 1);
INSERT INTO `menus` VALUES (1667235378, '2023-04-05 17:30:24', '2023-04-05 20:08:12', 1, 'k8sDeploymentDetail', 1667235346, '/kubernetes/deployment/:name/detail', 'kubernetes/deployment/detail/index', '', 'k8sDeploymentDetail', 1, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.k8sDeploymentDetail', '', 1);
INSERT INTO `menus` VALUES (1667235379, '2023-04-05 22:11:41', '2023-04-05 22:12:49', 1, 'pod日志', 1667235348, '/kubernetes/pod/log', 'kubernetes/pod/log/index', '', 'podLog', 10, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.podLog', '', 1);
INSERT INTO `menus` VALUES (1667235380, '2023-04-06 21:02:56', '2023-04-06 21:02:56', 1, 'podShell', 1667235348, '/kubernetes/pod/shell', 'kubernetes/pod/shell/index', '', 'podShell', 1, 1, '', 'GET', '', '', 1, 1, 0, 0, 'message.router.podShell', '', 1);
INSERT INTO `menus` VALUES (1667235381, '2023-04-07 19:23:42', '2023-04-17 16:22:17', 1, '创建deployment', 1667235346, '/kubernetes/deployment/create', 'kubernetes/deployment/create/index', '', 'deploymentCreate', 1, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.deploymentCreate', '', 1);
INSERT INTO `menus` VALUES (1667235382, '2023-04-14 15:24:19', '2023-04-14 15:29:05', 1, 'kibana', 0, '/kibana', 'layout/routerView/iframe', '', 'layoutIfameGrafana', 6, 1, 'iconfont icon-putong', '', '', 'http://192.168.3.130:5601/app/dashboards#/view/72383cd0-0a37-11ec-ada7-fb8814d13cc8?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))', 0, 0, 0, 1, 'message.router.kibana', '', 0);
INSERT INTO `menus` VALUES (1667235383, '2023-04-18 17:08:07', '2023-04-18 17:08:07', 1, '监控', 0, '/monitor', 'layout/routerView/parent', '/monitor/host', 'monitor', 99, 1, 'iconfont icon-putong', 'GET', '', '', 0, 1, 0, 0, 'message.router.monitor', '', 0);
INSERT INTO `menus` VALUES (1667235384, '2023-04-18 17:10:30', '2023-04-18 17:58:29', 1, '主机监控', 1667235383, '/monitor/host', 'layout/routerView/iframe', '', 'monitorHost', 98, 1, '', 'GET', '', 'http://192.168.3.120:3000/d/u4bLF3cnz/fu-wu-qi-ji-chu-xin-xi?orgId=1&var-job=All&var-hostname=iZbp1i0pbbmo42yyyztn3mZ&var-instance=172.16.10.2:9100&var-show_hostname=iZbp1i0pbbmo42yyyztn3mZ&kiosk', 0, 0, 0, 1, 'message.router.monitorHost', '', 0);
INSERT INTO `menus` VALUES (1667235385, '2023-04-18 17:17:10', '2023-04-18 17:40:02', 1, 'SSL证书监控', 1667235383, '/monitor/ssl', 'layout/routerView/iframe', '', 'monitorSSL', 99, 1, '', 'GET', '', 'http://192.168.3.120:3000/d/r8eWoHpGz/ssl-certificate-monitor?orgId=1&kiosk', 0, 0, 0, 1, 'message.router.monitorSSL', '', 0);
INSERT INTO `menus` VALUES (1667235386, '2023-04-18 17:25:13', '2023-04-18 17:40:12', 1, '阿里云流量访问看板', 1667235383, '/monitor/aliyunVistor', 'layout/routerView/iframe', '', 'monitorHost', 95, 1, '', 'GET', '', 'http://192.168.3.130:5601/app/dashboards#/view/72383cd0-0a37-11ec-ada7-fb8814d13cc8?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))', 0, 0, 0, 1, 'message.router.monitorAliyunVistor', '', 0);
INSERT INTO `menus` VALUES (1667235387, '2023-05-04 13:47:57', '2023-05-04 13:48:06', 1, 'podDetail', 1667235348, '/kubernetes/pod/detail', 'kubernetes/pod/detail', '', 'podDetail', 1, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.podDetail', '', 1);
INSERT INTO `menus` VALUES (1667235388, '2023-05-05 13:44:28', '2023-05-05 13:49:39', 1, 'nodeDetail', 1667235377, '/kubernetes/node/detail', 'kubernetes/node/detail', '', 'nodeDetail', 1, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.nodeDetail', '', 1);
INSERT INTO `menus` VALUES (1667235389, '2023-05-09 11:43:13', '2023-05-14 15:11:30', 1, 'namespace', 1667235339, '/kubernetes/namespace', 'kubernetes/namespace/index', '', 'k8sNamespace', 88, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sNamespace', '', 1);
INSERT INTO `menus` VALUES (1667235390, '2023-05-09 17:00:13', '2023-05-14 15:21:00', 1, 'k8sService', 1667235396, '/kubernetes/k8sNetwork/service', 'kubernetes/service/index', '', 'k8sService', 10, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sService', '', 1);
INSERT INTO `menus` VALUES (1667235391, '2023-05-09 17:01:06', '2023-05-14 15:16:46', 1, 'k8sIngress', 1667235396, '/kubernetes/k8sNetwork/ingress', 'kubernetes/ingress/index', '', 'k8sIngress', 9, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sIngress', '', 1);
INSERT INTO `menus` VALUES (1667235392, '2023-05-09 17:01:52', '2023-05-14 15:17:06', 1, 'k8sConfigmap', 1667235397, '/kubernetes/k8sConfig/configmap', 'kubernetes/configmap/index', '', 'k8sConfigmap', 8, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sConfigmap', '', 1);
INSERT INTO `menus` VALUES (1667235393, '2023-05-09 17:02:12', '2023-05-14 15:23:49', 1, 'k8sSecret', 1667235397, '/kubernetes/k8sConfig/secret', 'kubernetes/secret/index', '', 'k8sSecret', 7, 1, '', 'GET', '', '', 0, 0, 0, 0, 'message.router.k8sSecret', '', 1);
INSERT INTO `menus` VALUES (1667235394, '2023-05-11 15:21:19', '2023-05-11 15:21:19', 1, 'k8sServiceDetail', 1667235390, '/kubernetes/service/detail', 'kubernetes/service/detail/index', '', 'k8sServiceDetail', 1, 1, '', 'GET', '', '', 1, 0, 0, 0, 'message.router.k8sServiceDetail', '', 1);
INSERT INTO `menus` VALUES (1667235395, '2023-05-14 14:45:22', '2023-05-14 15:15:30', 1, '工作负载', 1667235339, '/kubernetes/workload', 'layout/routerView/parent', '/kubernetes/workload/deployment', 'kubernetes', 78, 1, 'iconfont icon-shuxingtu', '', '', '', 0, 1, 0, 0, 'message.router.k8sWorkload', '', 1);
INSERT INTO `menus` VALUES (1667235396, '2023-05-14 14:51:13', '2023-05-14 15:21:13', 1, '网络', 1667235339, '/kubernetes/k8sNetwork', 'layout/routerView/parent', '/kubernetes/k8sNetwork/service', 'k8sNetwork', 77, 1, 'iconfont icon-diqiu1', '', '', '', 0, 1, 0, 0, 'message.router.k8sNetwork', '', 1);
INSERT INTO `menus` VALUES (1667235397, '2023-05-14 14:53:06', '2023-05-14 15:19:17', 1, '配置', 1667235339, '/kubernetes/k8sConfig', 'layout/routerView/parent', '/kubernetes/k8sConfig/configmap', 'k8sConfig', 74, 1, 'iconfont icon-xitongguanli', '', '', '', 0, 1, 0, 0, 'message.router.k8sConfig', '', 1);
INSERT INTO `menus` VALUES (1667235398, '2023-06-02 06:41:58', '2023-06-06 16:22:12', 1, '退出登录', 0, '/api/v1/users/logout', '', '', '登出', 1, 3, '', 'POST', '', '', 0, 0, 0, 0, '', '用户', 0);
INSERT INTO `menus` VALUES (1667235399, '2023-06-05 10:54:49', '2023-06-05 10:54:49', 1, '系统日志', 0, '/logs', 'layout/routerView/parent', '/logs/loginLog', 'logs', 88, 1, 'ele-DataBoard', 'GET', '', '', 0, 1, 0, 0, 'message.router.logs', '', 0);
INSERT INTO `menus` VALUES (1667235400, '2023-06-05 10:56:18', '2023-06-05 10:56:18', 1, '登录日志', 1667235399, '/logs/login', 'logsys/loginLog/index', '', 'logLogin', 3, 1, 'iconfont icon-jiaoseguanli', 'GET', '', '', 0, 1, 0, 0, 'message.router.loginLog', '', 0);
INSERT INTO `menus` VALUES (1667235401, '2023-06-05 10:58:39', '2023-06-05 18:09:06', 1, '操作日志', 1667235399, '/logs/operator', 'logsys/operatorLog/index', '', 'operatorLog', 2, 1, 'iconfont icon-chazhaobiaodanliebiao', 'GET', '', '', 0, 1, 0, 0, 'message.router.operatorLog', '', 0);
INSERT INTO `menus` VALUES (1667235402, '2023-06-05 13:54:53', '2023-06-06 13:42:44', 1, '获取logs权限', 0, '/api/v1/logs/*', '', '', '获取logs权限', 19, 3, '', 'GET', '', '', 0, 0, 0, 0, '', '日志', 0);
INSERT INTO `menus` VALUES (1667235403, '2023-06-05 16:07:46', '2023-06-06 13:44:30', 1, '批量删除登录日志', 1667235400, '/api/v1/logs/login', '', '', '批量删除登录日志', 10, 2, '', 'DELETE', 'logs:login:del', '', 0, 1, 0, 0, '', '日志', 0);
INSERT INTO `menus` VALUES (1667235404, '2023-06-05 16:08:52', '2023-06-06 16:14:27', 1, '删除所有登录日志', 1667235400, '/api/v1/logs/login/all', '', '', '删除所有登录日志', 8, 2, '', 'DELETE', 'logs:login:del-all', '', 0, 1, 0, 0, '', '日志', 0);
INSERT INTO `menus` VALUES (1667235405, '2023-06-05 18:10:17', '2023-06-06 13:54:48', 1, '删除操作日志', 1667235401, '/api/v1/logs/operator', '', '', '删除操作日志', 2, 2, '', 'DELETE', 'logs:operator:del', '', 0, 1, 0, 0, '', '日志', 0);
INSERT INTO `menus` VALUES (1667235406, '2023-06-05 18:10:35', '2023-06-06 13:55:00', 1, '删除所有操作日志', 1667235401, '/api/v1/logs/operator/all', '', '', '删除所有操作日志', 1, 2, '', 'DELETE', 'logs:operator:del-all', '', 0, 1, 0, 0, '', '日志', 0);
INSERT INTO `menus` VALUES (1667235408, '2023-06-06 16:15:58', '2023-06-06 16:15:58', 1, '编辑用户', 1667235400, '/api/v1/users/:id', '', '', '编辑用户', 8, 2, '', 'PUT', 'sys:user:edit', '', 0, 1, 0, 0, '', '用户', 0);
COMMIT;

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `roleID` bigint unsigned NOT NULL COMMENT '角色ID',
  `menuID` bigint unsigned NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7439808866 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
BEGIN;
INSERT INTO `role_menus` VALUES (7439806912, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806913, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806914, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806915, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806916, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806917, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806918, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806919, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806920, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806921, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806922, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806923, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806924, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806925, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806926, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806927, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806928, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806929, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806930, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806931, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806932, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806933, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806934, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806935, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806936, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806937, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806938, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806939, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806940, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806941, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806942, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806943, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806944, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806945, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806946, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806947, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806948, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806949, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806950, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806951, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806952, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806953, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806954, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806955, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806956, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806957, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806958, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806959, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806960, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806961, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806962, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806963, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806964, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806965, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806966, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806967, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806968, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806969, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806970, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806971, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806972, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806973, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806974, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806975, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806976, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806977, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806978, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806979, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806980, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806981, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806982, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806983, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439806984, '2023-06-05 10:33:10', '2023-06-05 10:33:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807058, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807059, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807060, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807061, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807062, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807063, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807064, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807065, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807066, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807067, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807068, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807069, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807070, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807071, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807072, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807073, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807074, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807075, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807076, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807077, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807078, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807079, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807080, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807081, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807082, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807083, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807084, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807085, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807086, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807087, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807088, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807089, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807090, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807091, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807092, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807093, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807094, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807095, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807096, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807097, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807098, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807099, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807100, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807101, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807102, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807103, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807104, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807105, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807106, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807107, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807108, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807109, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807110, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807111, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807112, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807113, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807114, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807115, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807116, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807117, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807118, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807119, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807120, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807121, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807122, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807123, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807124, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807125, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807126, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807127, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807128, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807129, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807130, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807131, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807132, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807133, '2023-06-05 10:59:04', '2023-06-05 10:59:04', 0, 0);
INSERT INTO `role_menus` VALUES (7439807210, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807211, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807212, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807213, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807214, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807215, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807216, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807217, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807218, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807219, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807220, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807221, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807222, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807223, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807224, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807225, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807226, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807227, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807228, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807229, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807230, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807231, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807232, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807233, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807234, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807235, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807236, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807237, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807238, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807239, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807240, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807241, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807242, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807243, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807244, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807245, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807246, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807247, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807248, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807249, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807250, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807251, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807252, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807253, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807254, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807255, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807256, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807257, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807258, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807259, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807260, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807261, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807262, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807263, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807264, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807265, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807266, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807267, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807268, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807269, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807270, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807271, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807272, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807273, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807274, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807275, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807276, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807277, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807278, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807279, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807280, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807281, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807282, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807283, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807284, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807285, '2023-06-05 13:42:43', '2023-06-05 13:42:43', 0, 0);
INSERT INTO `role_menus` VALUES (7439807362, '2023-06-05 13:51:47', '2023-06-05 13:51:47', 0, 0);
INSERT INTO `role_menus` VALUES (7439807363, '2023-06-05 13:51:47', '2023-06-05 13:51:47', 0, 0);
INSERT INTO `role_menus` VALUES (7439807364, '2023-06-05 13:51:47', '2023-06-05 13:51:47', 0, 0);
INSERT INTO `role_menus` VALUES (7439807368, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807369, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807370, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807371, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807372, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807373, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807374, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807375, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807376, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807377, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807378, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807379, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807380, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807381, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807382, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807383, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807384, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807385, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807386, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807387, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807388, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807389, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807390, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807391, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807392, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807393, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807394, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807395, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807396, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807397, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807398, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807399, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807400, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807401, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807402, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807403, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807404, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807405, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807406, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807407, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807408, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807409, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807410, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807411, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807412, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807413, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807414, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807415, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807416, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807417, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807418, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807419, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807420, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807421, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807422, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807423, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807424, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807425, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807426, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807427, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807428, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807429, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807430, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807431, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807432, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807433, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807434, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807435, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807436, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807437, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807438, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807439, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807440, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807441, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807442, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807443, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807444, '2023-06-05 13:55:14', '2023-06-05 13:55:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807522, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807523, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807524, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807525, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807526, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807527, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807528, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807529, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807530, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807531, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807532, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807533, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807534, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807535, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807536, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807537, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807538, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807539, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807540, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807541, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807542, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807543, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807544, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807545, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807546, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807547, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807548, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807549, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807550, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807551, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807552, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807553, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807554, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807555, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807556, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807557, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807558, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807559, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807560, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807561, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807562, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807563, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807564, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807565, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807566, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807567, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807568, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807569, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807570, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807571, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807572, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807573, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807574, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807575, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807576, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807577, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807578, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807579, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807580, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807581, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807582, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807583, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807584, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807585, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807586, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807587, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807588, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807589, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807590, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807591, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807592, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807593, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807594, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807595, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807596, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807597, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807598, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807599, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807600, '2023-06-05 16:09:10', '2023-06-05 16:09:10', 0, 0);
INSERT INTO `role_menus` VALUES (7439807680, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807681, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807682, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807683, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807684, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807685, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807686, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807687, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807688, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807689, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807690, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807691, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807692, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807693, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807694, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807695, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807696, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807697, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807698, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807699, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807700, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807701, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807702, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807703, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807704, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807705, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807706, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807707, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807708, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807709, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807710, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807711, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807712, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807713, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807714, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807715, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807716, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807717, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807718, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807719, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807720, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807721, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807722, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807723, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807724, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807725, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807726, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807727, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807728, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807729, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807730, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807731, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807732, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807733, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807734, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807735, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807736, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807737, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807738, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807739, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807740, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807741, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807742, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807743, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807744, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807745, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807746, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807747, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807748, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807749, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807750, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807751, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807752, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807753, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807754, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807755, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807756, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807757, '2023-06-05 16:43:14', '2023-06-05 16:43:14', 0, 0);
INSERT INTO `role_menus` VALUES (7439807836, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807837, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807838, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807839, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807840, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807841, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807842, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807843, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807844, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807845, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807846, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807847, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807848, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807849, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807850, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807851, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807852, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807853, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807854, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807855, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807856, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807857, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807858, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807859, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807860, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807861, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807862, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807863, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807864, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807865, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807866, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807867, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807868, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807869, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807870, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807871, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807872, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807873, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807874, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807875, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807876, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807877, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807878, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807879, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807880, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807881, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807882, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807883, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807884, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807885, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807886, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807887, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807888, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807889, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807890, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807891, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807892, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807893, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807894, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807895, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807896, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807897, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807898, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807899, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807900, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807901, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807902, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807903, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807904, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807905, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807906, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807907, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807908, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807909, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807910, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807911, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807912, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807913, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807914, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807915, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807916, '2023-06-05 18:11:25', '2023-06-05 18:11:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439807998, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439807999, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808000, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808001, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808002, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808003, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808004, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808005, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808006, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808007, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808008, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808009, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808010, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808011, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808012, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808013, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808014, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808015, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808016, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808017, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808018, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808019, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808020, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808021, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808022, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808023, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808024, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808025, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808026, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808027, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808028, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808029, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808030, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808031, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808032, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808033, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808034, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808035, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808036, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808037, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808038, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808039, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808040, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808041, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808042, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808043, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808044, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808045, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808046, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808047, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808048, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808049, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808050, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808051, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808052, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808053, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808054, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808055, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808056, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808057, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808058, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808059, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808060, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808061, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808062, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808063, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808064, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808065, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808066, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808067, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808068, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808069, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808070, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808071, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808072, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808073, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808074, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808075, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808076, '2023-06-05 19:39:59', '2023-06-05 19:39:59', 0, 0);
INSERT INTO `role_menus` VALUES (7439808156, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808157, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808158, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808159, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808160, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808161, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808162, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808163, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808164, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808165, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808166, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808167, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808168, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808169, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808170, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808171, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808172, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808173, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808174, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808175, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808176, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808177, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808178, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808179, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808180, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808181, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808182, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808183, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808184, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808185, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808186, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808187, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808188, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808189, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808190, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808191, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808192, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808193, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808194, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808195, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808196, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808197, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808198, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808199, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808200, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808201, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808202, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808203, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808204, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808205, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808206, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808207, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808208, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808209, '2023-06-06 10:33:25', '2023-06-06 10:33:25', 0, 0);
INSERT INTO `role_menus` VALUES (7439808264, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808265, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808266, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808267, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808268, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808269, '2023-06-06 11:11:39', '2023-06-06 11:11:39', 0, 0);
INSERT INTO `role_menus` VALUES (7439808276, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808277, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808278, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808279, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808280, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808281, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808282, '2023-06-06 11:12:41', '2023-06-06 11:12:41', 0, 0);
INSERT INTO `role_menus` VALUES (7439808290, '2023-06-06 11:13:30', '2023-06-06 11:13:30', 0, 0);
INSERT INTO `role_menus` VALUES (7439808291, '2023-06-06 11:13:30', '2023-06-06 11:13:30', 0, 0);
INSERT INTO `role_menus` VALUES (7439808292, '2023-06-06 11:13:30', '2023-06-06 11:13:30', 0, 0);
INSERT INTO `role_menus` VALUES (7439808293, '2023-06-06 11:13:30', '2023-06-06 11:13:30', 0, 0);
INSERT INTO `role_menus` VALUES (7439808298, '2023-06-06 11:14:40', '2023-06-06 11:14:40', 0, 0);
INSERT INTO `role_menus` VALUES (7439808299, '2023-06-06 11:14:40', '2023-06-06 11:14:40', 0, 0);
INSERT INTO `role_menus` VALUES (7439808300, '2023-06-06 11:14:40', '2023-06-06 11:14:40', 0, 0);
INSERT INTO `role_menus` VALUES (7439808301, '2023-06-06 11:14:40', '2023-06-06 11:14:40', 0, 0);
INSERT INTO `role_menus` VALUES (7439808306, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808307, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808308, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808309, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808310, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808311, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808312, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808313, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808314, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808315, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808316, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808317, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808318, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808319, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808320, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808321, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808322, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808323, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808324, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808325, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808326, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808327, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808328, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808329, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808330, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808331, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808332, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808333, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808334, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808335, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808336, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808337, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808338, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808339, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808340, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808341, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808342, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808343, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808344, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808345, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808346, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808347, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808348, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808349, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808350, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808351, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808352, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808353, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808354, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808355, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808356, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808357, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808358, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808359, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808360, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808361, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808362, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808363, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808364, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808365, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808366, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808367, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808368, '2023-06-06 11:26:55', '2023-06-06 11:26:55', 0, 0);
INSERT INTO `role_menus` VALUES (7439808432, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808433, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808434, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808435, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808436, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808437, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808438, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808439, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808440, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808441, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808442, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808443, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808444, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808445, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808446, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808447, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808448, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808449, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808450, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808451, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808452, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808453, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808454, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808455, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808456, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808457, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808458, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808459, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808460, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808461, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808462, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808463, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808464, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808465, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808466, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808467, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808468, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808469, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808470, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808471, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808472, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808473, '2023-06-06 15:05:33', '2023-06-06 15:05:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808516, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808517, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808518, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808519, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808520, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808521, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808522, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808523, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808524, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808525, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808526, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808527, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808528, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808529, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808530, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808531, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808532, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808533, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808534, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808535, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808536, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808537, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808538, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808539, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808540, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808541, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808542, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808543, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808544, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808545, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808546, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808547, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808548, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808549, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808550, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808551, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808552, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808553, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808554, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808555, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808556, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808557, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808558, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808559, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808560, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808561, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808562, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808563, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808564, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808565, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808566, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808567, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808568, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808569, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808570, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808571, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808572, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808573, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808574, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808575, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808576, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808577, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808578, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808579, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808580, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808581, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808582, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808583, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808584, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808585, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808586, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808587, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808588, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808589, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808590, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808591, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808592, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808593, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808594, '2023-06-06 16:25:33', '2023-06-06 16:25:33', 0, 0);
INSERT INTO `role_menus` VALUES (7439808674, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808675, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808676, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808677, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808678, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808679, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808680, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808681, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808682, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808683, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808684, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808685, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808686, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808687, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808688, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808689, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808690, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808691, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808692, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808693, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808694, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808695, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808696, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808697, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808698, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808699, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808700, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808701, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808702, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808703, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808704, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808705, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808706, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808707, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808708, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808709, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808710, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808711, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808712, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808713, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808714, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808715, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808716, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808717, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808718, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808719, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808720, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808721, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808722, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808723, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808724, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808725, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808726, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808727, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808728, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808729, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808730, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808731, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808732, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808733, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808734, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808735, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808736, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808737, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808738, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808739, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808740, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808741, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808742, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808743, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808744, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808745, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808746, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808747, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808748, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808749, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808750, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808751, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808752, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808753, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808754, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 0, 0);
INSERT INTO `role_menus` VALUES (7439808755, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235314);
INSERT INTO `role_menus` VALUES (7439808756, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235383);
INSERT INTO `role_menus` VALUES (7439808757, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235385);
INSERT INTO `role_menus` VALUES (7439808758, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235384);
INSERT INTO `role_menus` VALUES (7439808759, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235386);
INSERT INTO `role_menus` VALUES (7439808760, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235339);
INSERT INTO `role_menus` VALUES (7439808761, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235376);
INSERT INTO `role_menus` VALUES (7439808762, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235375);
INSERT INTO `role_menus` VALUES (7439808763, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235389);
INSERT INTO `role_menus` VALUES (7439808764, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235377);
INSERT INTO `role_menus` VALUES (7439808765, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235388);
INSERT INTO `role_menus` VALUES (7439808766, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235395);
INSERT INTO `role_menus` VALUES (7439808767, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235346);
INSERT INTO `role_menus` VALUES (7439808768, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235381);
INSERT INTO `role_menus` VALUES (7439808769, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235378);
INSERT INTO `role_menus` VALUES (7439808770, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235348);
INSERT INTO `role_menus` VALUES (7439808771, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235379);
INSERT INTO `role_menus` VALUES (7439808772, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235387);
INSERT INTO `role_menus` VALUES (7439808773, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235380);
INSERT INTO `role_menus` VALUES (7439808774, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235396);
INSERT INTO `role_menus` VALUES (7439808775, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235390);
INSERT INTO `role_menus` VALUES (7439808776, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235394);
INSERT INTO `role_menus` VALUES (7439808777, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235391);
INSERT INTO `role_menus` VALUES (7439808778, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235397);
INSERT INTO `role_menus` VALUES (7439808779, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235392);
INSERT INTO `role_menus` VALUES (7439808780, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235393);
INSERT INTO `role_menus` VALUES (7439808781, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235310);
INSERT INTO `role_menus` VALUES (7439808782, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235368);
INSERT INTO `role_menus` VALUES (7439808783, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235367);
INSERT INTO `role_menus` VALUES (7439808784, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235311);
INSERT INTO `role_menus` VALUES (7439808785, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235312);
INSERT INTO `role_menus` VALUES (7439808786, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235313);
INSERT INTO `role_menus` VALUES (7439808787, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235355);
INSERT INTO `role_menus` VALUES (7439808788, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235399);
INSERT INTO `role_menus` VALUES (7439808789, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235400);
INSERT INTO `role_menus` VALUES (7439808790, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235401);
INSERT INTO `role_menus` VALUES (7439808791, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235362);
INSERT INTO `role_menus` VALUES (7439808792, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235365);
INSERT INTO `role_menus` VALUES (7439808793, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235364);
INSERT INTO `role_menus` VALUES (7439808794, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235366);
INSERT INTO `role_menus` VALUES (7439808795, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235363);
INSERT INTO `role_menus` VALUES (7439808796, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235382);
INSERT INTO `role_menus` VALUES (7439808797, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235338);
INSERT INTO `role_menus` VALUES (7439808798, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235340);
INSERT INTO `role_menus` VALUES (7439808799, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235357);
INSERT INTO `role_menus` VALUES (7439808800, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235358);
INSERT INTO `role_menus` VALUES (7439808801, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235359);
INSERT INTO `role_menus` VALUES (7439808802, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235360);
INSERT INTO `role_menus` VALUES (7439808803, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235361);
INSERT INTO `role_menus` VALUES (7439808804, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235370);
INSERT INTO `role_menus` VALUES (7439808805, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235369);
INSERT INTO `role_menus` VALUES (7439808806, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235402);
INSERT INTO `role_menus` VALUES (7439808807, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235403);
INSERT INTO `role_menus` VALUES (7439808808, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235404);
INSERT INTO `role_menus` VALUES (7439808809, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235405);
INSERT INTO `role_menus` VALUES (7439808810, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235406);
INSERT INTO `role_menus` VALUES (7439808811, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235318);
INSERT INTO `role_menus` VALUES (7439808812, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235315);
INSERT INTO `role_menus` VALUES (7439808813, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235316);
INSERT INTO `role_menus` VALUES (7439808814, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235408);
INSERT INTO `role_menus` VALUES (7439808815, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235328);
INSERT INTO `role_menus` VALUES (7439808816, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235335);
INSERT INTO `role_menus` VALUES (7439808817, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235398);
INSERT INTO `role_menus` VALUES (7439808818, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235319);
INSERT INTO `role_menus` VALUES (7439808819, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235321);
INSERT INTO `role_menus` VALUES (7439808820, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235322);
INSERT INTO `role_menus` VALUES (7439808821, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235323);
INSERT INTO `role_menus` VALUES (7439808822, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235324);
INSERT INTO `role_menus` VALUES (7439808823, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235336);
INSERT INTO `role_menus` VALUES (7439808824, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235320);
INSERT INTO `role_menus` VALUES (7439808825, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235327);
INSERT INTO `role_menus` VALUES (7439808826, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235325);
INSERT INTO `role_menus` VALUES (7439808827, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235326);
INSERT INTO `role_menus` VALUES (7439808828, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235372);
INSERT INTO `role_menus` VALUES (7439808829, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235373);
INSERT INTO `role_menus` VALUES (7439808830, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235334);
INSERT INTO `role_menus` VALUES (7439808831, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235341);
INSERT INTO `role_menus` VALUES (7439808832, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235342);
INSERT INTO `role_menus` VALUES (7439808833, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235343);
INSERT INTO `role_menus` VALUES (7439808834, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235344);
INSERT INTO `role_menus` VALUES (7439808835, '2023-06-06 16:41:16', '2023-06-06 16:41:16', 5439801856, 1667235347);
INSERT INTO `role_menus` VALUES (7439808836, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808837, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808838, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808839, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808840, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808841, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808842, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808843, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808844, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808845, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808846, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808847, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808848, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808849, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808850, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 0, 0);
INSERT INTO `role_menus` VALUES (7439808851, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235310);
INSERT INTO `role_menus` VALUES (7439808852, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235314);
INSERT INTO `role_menus` VALUES (7439808853, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235311);
INSERT INTO `role_menus` VALUES (7439808854, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235312);
INSERT INTO `role_menus` VALUES (7439808855, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235313);
INSERT INTO `role_menus` VALUES (7439808856, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235355);
INSERT INTO `role_menus` VALUES (7439808857, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235399);
INSERT INTO `role_menus` VALUES (7439808858, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235400);
INSERT INTO `role_menus` VALUES (7439808859, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235401);
INSERT INTO `role_menus` VALUES (7439808860, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235402);
INSERT INTO `role_menus` VALUES (7439808861, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235318);
INSERT INTO `role_menus` VALUES (7439808862, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235398);
INSERT INTO `role_menus` VALUES (7439808863, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235319);
INSERT INTO `role_menus` VALUES (7439808864, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235320);
INSERT INTO `role_menus` VALUES (7439808865, '2023-06-07 03:21:27', '2023-06-07 03:21:27', 5439801874, 1667235344);
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `memo` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `sequence` bigint NOT NULL COMMENT '排序值',
  `parent_id` bigint unsigned NOT NULL COMMENT '父级ID',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0 表示禁用，1 表示启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5439801879 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO `roles` VALUES (5439801856, '2022-10-29 18:45:09', '2022-11-05 16:36:43', '管理员', '管理员', 10, 0, 1);
INSERT INTO `roles` VALUES (5439801866, '2022-11-06 22:02:12', '2022-11-07 17:41:52', '', '啊啊啊', 1, 5439801856, 1);
INSERT INTO `roles` VALUES (5439801874, '2023-03-30 16:46:58', '2023-06-06 15:09:05', '只有查看权限', 'gust', 9, 0, 1);
COMMIT;

-- ----------------------------
-- Table structure for rules
-- ----------------------------
DROP TABLE IF EXISTS `rules`;
CREATE TABLE `rules` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4439804820 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of rules
-- ----------------------------
BEGIN;
INSERT INTO `rules` VALUES (4439804720, 'g', '1439801856', '5439801856', '', '', '', '');
INSERT INTO `rules` VALUES (4439804755, 'g', '1439801858', '5439801874', '', '', '', '');
INSERT INTO `rules` VALUES (4439804784, 'p', '5439801856', '/api/v1/users/register', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804785, 'p', '5439801856', '/api/v1/users/:id', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804786, 'p', '5439801856', '/api/v1/user*|/api/v1/role*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804787, 'p', '5439801856', '/api/v1/role*|/api/v1/menu*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804788, 'p', '5439801856', '/api/v1/menus*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804789, 'p', '5439801856', '/api/v1/roles', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804790, 'p', '5439801856', '/api/v1/roles/:id', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804791, 'p', '5439801856', '/api/v1/roles/:id', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804792, 'p', '5439801856', '/api/v1/roles/:id/menus', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804793, 'p', '5439801856', '/api/v1/menus', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804794, 'p', '5439801856', '/api/v1/menus/:id', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804795, 'p', '5439801856', '/api/v1/menus/:id', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804796, 'p', '5439801856', '/api/v1/users/:id/roles', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804797, 'p', '5439801856', '/api/v1/menus/:id/status/:status', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804798, 'p', '5439801856', '/api/v1/users/:id/status/:status', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804799, 'p', '5439801856', '/api/v1/roles/:id/status/:status', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804800, 'p', '5439801856', '/api/v1/host', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804801, 'p', '5439801856', '/api/v1/host/:id', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804802, 'p', '5439801856', '/api/v1/host/:id/ws', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804803, 'p', '5439801856', '/api/v1/host*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804804, 'p', '5439801856', '/namespace/:name', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804805, 'p', '5439801856', '/api/v1/menu', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804806, 'p', '5439801856', '/api/v1/menu/:id', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804807, 'p', '5439801856', '/api/v1/users/logout', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804808, 'p', '5439801856', '/api/v1/logs/*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804809, 'p', '5439801856', '/api/v1/logs/login', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804810, 'p', '5439801856', '/api/v1/logs/login/all', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804811, 'p', '5439801856', '/api/v1/logs/operator', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804812, 'p', '5439801856', '/api/v1/logs/operator/all', 'DELETE', '', '', '');
INSERT INTO `rules` VALUES (4439804813, 'p', '5439801856', '/api/v1/users/:id', 'PUT', '', '', '');
INSERT INTO `rules` VALUES (4439804814, 'p', '5439801874', '/api/v1/user*|/api/v1/role*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804815, 'p', '5439801874', '/api/v1/role*|/api/v1/menu*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804816, 'p', '5439801874', '/api/v1/menus*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804817, 'p', '5439801874', '/api/v1/host*', 'GET', '', '', '');
INSERT INTO `rules` VALUES (4439804818, 'p', '5439801874', '/api/v1/users/logout', 'POST', '', '', '');
INSERT INTO `rules` VALUES (4439804819, 'p', '5439801874', '/api/v1/logs/*', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for tenants
-- ----------------------------
DROP TABLE IF EXISTS `tenants`;
CREATE TABLE `tenants` (
  `tenant_id` int NOT NULL AUTO_INCREMENT,
  `tenant_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of tenants
-- ----------------------------
BEGIN;
INSERT INTO `tenants` VALUES (1, 'domain1');
INSERT INTO `tenants` VALUES (2, 'domain2');
COMMIT;

-- ----------------------------
-- Table structure for user_resources
-- ----------------------------
DROP TABLE IF EXISTS `user_resources`;
CREATE TABLE `user_resources` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '资源类型',
  `resource_id` bigint unsigned DEFAULT NULL COMMENT '资源ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_resources
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '管理员ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2439801901 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` VALUES (2439801891, '2023-06-06 10:12:41', '2023-06-06 10:12:41', 1439801868, 5439801856);
INSERT INTO `user_roles` VALUES (2439801892, '2023-06-06 10:12:41', '2023-06-06 10:12:41', 1439801868, 5439801874);
INSERT INTO `user_roles` VALUES (2439801898, '2023-06-06 11:20:50', '2023-06-06 11:20:50', 1439801856, 5439801856);
INSERT INTO `user_roles` VALUES (2439801900, '2023-06-06 15:09:15', '2023-06-06 15:09:15', 1439801858, 5439801874);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` char(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mobile` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态(1:正常 2:未激活 3:暂停使用)',
  `description` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1439801903 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1439801856, '2022-10-29 18:42:49', '2023-06-06 16:26:55', 'lbemi_admin', '$2a$10$OZytyuVMCELLJhowlQSRTubSf3etwC2ey/G0wLO0pd6o4WOyCoVe6', 'admin@admin.com', '18088888888', 1, '啊啊');
INSERT INTO `users` VALUES (1439801858, '2022-10-29 20:35:08', '2023-06-06 16:02:33', 'lbemi', '$2a$10$iyPEYrsdaBTQ5jq3w7exIeKqCRHfTFHJemk5t03o01XR01rlLm/E6', 'lbemi@test.com', '15558888888', 1, '只有查看权限');
INSERT INTO `users` VALUES (1439801902, '2023-06-06 16:40:32', '2023-06-06 16:40:32', 'test', '$2a$10$TjjgLGH5khc4XHX4tNj7yOqgGAS5U0fIKwuvgLN4dEppNuyzjJBxm', 'test@test.com', '', 2, '测试用户');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
