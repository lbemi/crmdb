INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235310, '2022-10-29 15:26:46', '2022-11-08 13:47:55', 1, '系统管理菜单', 1667235314, '/system', 'layout/routerView/parent', null, 'system', 1, 1, 'icon-shezhi', '', '', null, 0, 1, 0, 1, 'message.router.system');
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235311, '2022-10-29 15:32:06', '2022-11-08 11:34:12', 1, '用户管理菜单', 1667235310, '/user', 'system/user/index', null, 'systemUser', 10, 1, 'iconfont icon-icon-', '', '', null, 0, 1, 0, 1, 'message.router.systemUser');
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235312, '2022-10-29 15:32:29', '2022-11-08 11:34:37', 1, '角色管理菜单', 1667235310, '/role', 'system/role/index', null, 'systemRole', 9, 1, 'iconfont icon-caidan', '', '', null, 0, 1, 0, 1, 'message.router.systemRole');
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235313, '2022-10-29 15:33:01', '2022-11-08 11:34:51', 1, '菜单权限管理菜单', 1667235310, '/menu', 'system/menu/index', null, 'systemMenu', 8, 1, 'iconfont icon-caidan', '', '', null, 0, 1, 0, 1, 'message.router.systemMenu');
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235314, '2022-10-29 15:35:33', '2022-11-09 10:03:32', 1, '首页', 0, '/home', 'home/index', null, 'home', 10, 1, 'iconfont icon-shouye', '', '', null, 0, 1, 0, 1, 'message.router.home');
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235315, '2022-10-29 15:40:32', '2022-11-05 13:44:06', 1, '添加用户按钮', 1667235311, '/api/v1/user', null, null, '添加用户', 10, 2, '', 'POST', 'sys:user:add', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235316, '2022-10-29 15:41:28', '2022-11-07 16:09:17', 1, '删除用户按钮', 1667235311, '/api/v1/user/:id', null, null, '删除用户', 8, 2, '', 'DELETE', 'sys:user:del', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235317, '2022-10-29 15:42:04', '2022-10-29 15:42:04', 1, '编辑用户按钮', 1667235311, '/api/v1/user', null, null, '编辑用户', 8, 2, '', 'PUT', 'sys:user:edit', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235318, '2022-10-29 15:43:33', '2022-11-05 13:43:51', 1, '查看基础权限', 1667235311, '/api/v1/user*|/api/v1/role*', null, null, '基础权限', 11, 3, '', 'GET', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235319, '2022-10-29 15:50:04', '2022-10-29 15:50:04', 1, '查看基础权限', 1667235312, '/api/v1/role*|/api/v1/menu*', null, null, '基础权限', 10, 3, '', 'GET', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235320, '2022-10-29 15:50:42', '2022-10-29 15:50:42', 1, '查看基础权限', 1667235313, '/api/v1/menu*', null, null, '基础权限', 10, 3, '', 'GET', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235321, '2022-10-29 15:52:23', '2022-10-29 15:52:23', 1, '添加按钮权限', 1667235312, '/api/v1/role', null, null, '添加角色', 9, 3, '', 'POST', 'sys:role:add', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235322, '2022-10-29 15:52:53', '2022-11-05 19:07:16', 1, '编辑角色按钮权限', 1667235312, '/api/v1/role/:id', null, null, '编辑角色', 8, 2, '', 'PUT', 'sys:role:edit', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235323, '2022-10-29 15:53:19', '2022-11-05 19:07:25', 1, '删除角色按钮权限', 1667235312, '/api/v1/role/:id', null, null, '删除角色', 7, 2, '', 'DELETE', 'sys:role:del', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235324, '2022-10-29 15:54:36', '2022-10-29 15:54:36', 1, '角色分配权限按钮', 1667235312, '/api/v1/role/:id/menus', null, null, '分配权限按钮', 7, 2, '', 'POST', 'sys:role:set', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235325, '2022-10-29 15:58:43', '2022-11-05 14:04:12', 1, '添加按钮权限按钮', 1667235313, '/api/v1/menu', null, null, '添加按钮', 7, 2, '', 'POST', 'sys:menu:add', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235326, '2022-10-29 15:59:58', '2022-11-06 10:59:52', 1, '删除按钮权限按钮', 1667235313, '/api/v1/menu/:id', null, null, '删除按钮', 7, 2, '', 'DELETE', 'sys:menu:del', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235327, '2022-10-29 16:00:18', '2022-10-29 16:00:18', 1, '编辑按钮权限按钮', 1667235313, '/api/v1/menu/:id', null, null, '编辑', 7, 2, '', 'PUT', 'sys:menu:edit', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235328, '2022-10-29 16:01:15', '2022-11-06 14:18:51', 1, '分配角色按钮', 1667235311, '/api/v1/user/:id/roles', null, null, '分配角色', 6, 2, '', 'POST', 'sys:user:set', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235332, '2022-11-06 11:06:50', '2022-11-06 11:06:50', 1, '测试api', 1667235313, '/api/v1/menu/:id/test', null, null, '测试api', 1, 3, '', 'PUT', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235333, '2022-11-06 14:29:05', '2022-11-07 17:42:38', 1, '允许所有接口GET', 0, '/api/v1/*', null, null, '基础权限', 1, 3, '', 'POST', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235334, '2022-11-06 16:27:57', '2022-11-06 16:28:57', 1, '启用禁用按钮', 1667235313, '/api/v1/menu/:id/status/:status', null, null, '启停', 1, 2, '', 'PUT', 'sys:menu:status', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235335, '2022-11-07 17:44:07', '2022-11-07 17:44:07', 1, '启用停用', 1667235311, '/api/v1/user/:id/status/:status', null, null, '启用停用', 1, 2, '', 'PUT', 'sys:user:status', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235336, '2022-11-07 18:10:33', '2022-11-07 18:10:33', 1, '启用禁用角色', 1667235312, '/api/v1/role/:id/status/:status', null, null, '启用禁用', 1, 2, '', 'PUT', 'sys:role:status', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235338, '2022-11-08 13:44:48', '2022-11-16 14:29:09', 1, '资产管理', 1667235314, '/asset', null, null, '资产管理', 4, 1, 'icon-zichanguanli', '', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235339, '2022-11-08 13:45:39', '2023-03-21 20:15:16', 1, 'kubernetes容器管理', 1667235314, '/container', null, null, 'Kubernetes', 1, 1, 'icon-kubernetes', '', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235340, '2022-11-08 13:47:00', '2022-11-16 14:28:34', 1, '主机管理', 1667235338, '/host', null, null, '主机管理', 1, 1, 'icon-zhujiguanli', '', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235341, '2022-11-09 10:14:15', '2022-11-09 10:14:15', 1, '添加主机', 1667235340, '/api/v1/host', null, null, '添加', 1, 2, '', 'POST', 'asset:host:add', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235342, '2022-11-09 10:14:56', '2022-11-09 10:14:56', 1, '删除', 1667235340, '/api/v1/host/:id', null, null, '删除', 1, 2, '', 'DELETE', 'asset:host:del', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235343, '2022-11-09 10:16:01', '2022-11-10 16:06:54', 1, 'Terminal', 1667235340, '/api/v1/host/:id/ws', null, null, 'SSH', 1, 2, '', 'GET', 'asset:host:ssh', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235344, '2022-11-09 10:16:36', '2022-11-09 10:16:36', 1, '基础权限', 1667235340, '/api/v1/host*', null, null, '基础权限', 1, 3, '', 'GET', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235346, '2022-12-07 17:49:15', '2023-03-21 20:19:11', 2, 'deployment', 1667235339, '/deployment', null, null, 'deployment', 1, 1, '', '', '', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235347, '2022-12-27 16:11:11', '2022-12-27 16:11:11', 1, 'k8s:deploy:delete', 1667235346, '/namespace/:name', null, null, '删除', 1, 2, '', 'DELETE', 'k8s:deploy:delete', null, 0, 1, 0, 1, null);
INSERT INTO lbemi.menus (id, created_at, updated_at, status, memo, parent_id, url, component, redirect, name, sequence, menu_type, icon, method, code, is_link, is_hide, is_keepalive, is_affix, is_iframe, title) VALUES (1667235348, '2023-03-21 20:12:40', '2023-03-21 20:19:14', 2, 'pods', 1667235339, '/pod', null, null, '容器组', 1, 1, '', '', '', null, 0, 1, 0, 1, null);