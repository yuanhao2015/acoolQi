/*
 Navicat MySQL Data Transfer

 Source Server         : 10.10.88.241(test_ywadmin)
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : 10.10.88.241:30133
 Source Schema         : acoolqi-admin

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 14/10/2021 10:02:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 102 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 'admin', '2021-07-07 18:29:57', '', NULL, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2021-07-07 18:29:57', '', NULL, '初始化密码 123456');
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-light', 'Y', 'admin', '2021-07-07 18:29:57', 'admin', NULL, '深色主题theme-dark，浅色主题theme-light');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 211 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', '泰逢科技', 0, NULL, NULL, NULL, '0', '0', 'admin', '2021-07-07 18:29:57', '', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '公共部', 1, NULL, NULL, NULL, '0', '0', 'admin', '2021-07-07 18:29:57', '', NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '运维部门', 1, NULL, NULL, NULL, '0', '0', 'admin', '2021-07-07 18:29:57', '', NULL);
INSERT INTO `sys_dept` VALUES (200, 100, '0,100', '项目组', 2, '', '', '', '0', '0', 'admin', '2021-09-27 11:04:41', '', NULL);
INSERT INTO `sys_dept` VALUES (201, 200, '0,100,200', '飓风部门', 1, '', '', '', '0', '0', 'admin', '2021-09-27 11:06:18', '', NULL);
INSERT INTO `sys_dept` VALUES (202, 200, '0,100,200', '点聚部门', 2, '', '', '', '0', '0', 'admin', '2021-09-27 11:08:56', '', NULL);
INSERT INTO `sys_dept` VALUES (203, 101, '0,100,101', 'PHP部门', 2, '', '', '', '0', '0', 'admin', '2021-09-27 11:38:52', '', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'N' COMMENT '是否默认（Y是 N否）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 101 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '性别男');
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '性别女');
INSERT INTO `sys_dict_data` VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '性别未知');
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '显示菜单');
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '隐藏菜单');
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '默认分组');
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '系统分组');
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '系统默认是');
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '系统默认否');
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '通知');
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '公告');
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '关闭状态');
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'primary', 'N', '0', 'admin', '2021-07-07 18:29:57', 'admin', '2021-10-09 17:59:42', '新增操作');
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2021-07-07 18:29:57', 'admin', '2021-10-09 19:29:02', '修改操作');
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '删除操作');
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2021-07-07 18:29:57', 'admin', '2021-10-09 19:22:47', '授权操作');
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'success', 'N', '0', 'admin', '2021-07-07 18:29:57', 'admin', '2021-10-09 19:29:11', '导出操作');
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '导入操作');
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '强退操作');
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '生成操作');
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '清空操作');
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '正常状态');
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '停用状态');
INSERT INTO `sys_dict_data` VALUES (100, 10, '其他', '10', 'sys_oper_type', '', 'primary', '', '0', 'admin', '2021-10-11 10:08:24', 'admin', '2021-10-11 10:08:33', '其他类型');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '用户性别列表');
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '菜单状态列表');
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '系统开关列表');
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '任务状态列表');
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '任务分组列表');
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '系统是否列表');
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '通知类型列表');
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '通知状态列表');
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '操作类型列表');
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '登录状态列表');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '3' COMMENT '计划执行错误策略（1立即执行 2执行一次 3放弃执行）',
  `concurrent` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注信息',
  PRIMARY KEY (`job_id`, `job_name`, `job_group`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '系统默认（无参）', 'DEFAULT', 'ryTask.ryNoParams', '0/10 * * * * ?', '3', '1', '1', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_job` VALUES (2, '系统默认（有参）', 'DEFAULT', 'ryTask.ryParams(\'ry\')', '0/15 * * * * ?', '3', '1', '1', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_job` VALUES (3, '系统默认（多参）', 'DEFAULT', 'ryTask.ryMultipleParams(\'ry\', true, 2000L, 316.50D, 100)', '0/20 * * * * ?', '3', '1', '1', 'admin', '2021-07-07 18:29:57', '', NULL, '');

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
  `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '调用目标字符串',
  `job_message` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '异常信息',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '定时任务调度日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_logininfor
-- ----------------------------
DROP TABLE IF EXISTS `sys_logininfor`;
CREATE TABLE `sys_logininfor`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` smallint(2) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime(0) NULL DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 381 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_logininfor
-- ----------------------------
INSERT INTO `sys_logininfor` VALUES (323, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 14:22:53');
INSERT INTO `sys_logininfor` VALUES (324, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-11 14:25:27');
INSERT INTO `sys_logininfor` VALUES (325, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 14:25:37');
INSERT INTO `sys_logininfor` VALUES (326, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-11 14:27:16');
INSERT INTO `sys_logininfor` VALUES (327, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 14:28:04');
INSERT INTO `sys_logininfor` VALUES (328, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 10', 0, '登录成功', '2021-10-11 15:01:35');
INSERT INTO `sys_logininfor` VALUES (329, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 15:29:35');
INSERT INTO `sys_logininfor` VALUES (330, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-11 16:09:27');
INSERT INTO `sys_logininfor` VALUES (331, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 1, '登录失败', '2021-10-11 16:09:34');
INSERT INTO `sys_logininfor` VALUES (332, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 16:09:37');
INSERT INTO `sys_logininfor` VALUES (333, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-11 16:10:02');
INSERT INTO `sys_logininfor` VALUES (334, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 16:10:05');
INSERT INTO `sys_logininfor` VALUES (335, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 17:32:32');
INSERT INTO `sys_logininfor` VALUES (336, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 18:54:32');
INSERT INTO `sys_logininfor` VALUES (337, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-11 18:55:24');
INSERT INTO `sys_logininfor` VALUES (338, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 1, '密码错误', '2021-10-11 18:55:30');
INSERT INTO `sys_logininfor` VALUES (339, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 1, '密码错误', '2021-10-11 18:55:31');
INSERT INTO `sys_logininfor` VALUES (340, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 1, '密码错误', '2021-10-11 18:55:32');
INSERT INTO `sys_logininfor` VALUES (341, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-11 18:55:36');
INSERT INTO `sys_logininfor` VALUES (342, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-12 15:53:24');
INSERT INTO `sys_logininfor` VALUES (343, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-12 15:56:44');
INSERT INTO `sys_logininfor` VALUES (344, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-12 15:56:47');
INSERT INTO `sys_logininfor` VALUES (345, 'admin', '10.10.88.117', '内网IP', 'Chrome', 'Windows 10', 0, '登录成功', '2021-10-12 15:58:57');
INSERT INTO `sys_logininfor` VALUES (346, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-12 16:01:36');
INSERT INTO `sys_logininfor` VALUES (347, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-12 16:01:38');
INSERT INTO `sys_logininfor` VALUES (348, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-12 16:03:07');
INSERT INTO `sys_logininfor` VALUES (349, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-12 16:03:08');
INSERT INTO `sys_logininfor` VALUES (350, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-12 18:16:00');
INSERT INTO `sys_logininfor` VALUES (351, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 09:48:17');
INSERT INTO `sys_logininfor` VALUES (352, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 10:51:08');
INSERT INTO `sys_logininfor` VALUES (353, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 11:30:56');
INSERT INTO `sys_logininfor` VALUES (354, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 11:31:03');
INSERT INTO `sys_logininfor` VALUES (355, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 11:31:13');
INSERT INTO `sys_logininfor` VALUES (356, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 11:31:16');
INSERT INTO `sys_logininfor` VALUES (357, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 11:47:19');
INSERT INTO `sys_logininfor` VALUES (358, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 14:38:01');
INSERT INTO `sys_logininfor` VALUES (359, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 14:56:58');
INSERT INTO `sys_logininfor` VALUES (360, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 14:57:03');
INSERT INTO `sys_logininfor` VALUES (361, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:12:10');
INSERT INTO `sys_logininfor` VALUES (362, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 1, '密码错误', '2021-10-13 15:12:15');
INSERT INTO `sys_logininfor` VALUES (363, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:12:19');
INSERT INTO `sys_logininfor` VALUES (364, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:12:38');
INSERT INTO `sys_logininfor` VALUES (365, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:12:40');
INSERT INTO `sys_logininfor` VALUES (366, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:13:06');
INSERT INTO `sys_logininfor` VALUES (367, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:13:11');
INSERT INTO `sys_logininfor` VALUES (368, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:14:21');
INSERT INTO `sys_logininfor` VALUES (369, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:14:25');
INSERT INTO `sys_logininfor` VALUES (370, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:14:42');
INSERT INTO `sys_logininfor` VALUES (371, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:14:50');
INSERT INTO `sys_logininfor` VALUES (372, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:15:02');
INSERT INTO `sys_logininfor` VALUES (373, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:15:06');
INSERT INTO `sys_logininfor` VALUES (374, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:19:02');
INSERT INTO `sys_logininfor` VALUES (375, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:19:10');
INSERT INTO `sys_logininfor` VALUES (376, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 15:21:49');
INSERT INTO `sys_logininfor` VALUES (377, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:21:54');
INSERT INTO `sys_logininfor` VALUES (378, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 15:59:01');
INSERT INTO `sys_logininfor` VALUES (379, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 16:05:11');
INSERT INTO `sys_logininfor` VALUES (380, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 16:05:15');
INSERT INTO `sys_logininfor` VALUES (381, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 16:15:16');
INSERT INTO `sys_logininfor` VALUES (382, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 16:15:20');
INSERT INTO `sys_logininfor` VALUES (383, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 16:18:57');
INSERT INTO `sys_logininfor` VALUES (384, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 16:19:01');
INSERT INTO `sys_logininfor` VALUES (385, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 16:29:16');
INSERT INTO `sys_logininfor` VALUES (386, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 16:29:24');
INSERT INTO `sys_logininfor` VALUES (387, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 17:17:29');
INSERT INTO `sys_logininfor` VALUES (388, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 17:30:48');
INSERT INTO `sys_logininfor` VALUES (389, 'admin', '10.10.88.117', '内网IP', 'Chrome', 'Windows 10', 0, '登录成功', '2021-10-13 17:59:29');
INSERT INTO `sys_logininfor` VALUES (390, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 18:34:25');
INSERT INTO `sys_logininfor` VALUES (391, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 18:35:02');
INSERT INTO `sys_logininfor` VALUES (392, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 18:35:07');
INSERT INTO `sys_logininfor` VALUES (393, 'test', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登出成功', '2021-10-13 18:36:41');
INSERT INTO `sys_logininfor` VALUES (394, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-13 18:53:09');
INSERT INTO `sys_logininfor` VALUES (395, 'admin', '10.10.88.36', '内网IP', 'Chrome', 'Windows 10', 0, '登录成功', '2021-10-13 18:53:56');
INSERT INTO `sys_logininfor` VALUES (396, 'admin', '10.10.24.5', '内网IP', 'Chrome', 'Windows 7', 0, '登录成功', '2021-10-14 09:51:00');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父菜单ID',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `is_frame` int(1) NULL DEFAULT 1 COMMENT '是否为外链（0是 1否）',
  `is_cache` int(1) NULL DEFAULT 0 COMMENT '是否缓存（0缓存 1不缓存）',
  `menu_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '菜单状态（0显示 1隐藏）',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '菜单状态（0正常 1停用）',
  `perms` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '菜单图标',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2009 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '系统管理', 0, 1, 'system', NULL, 1, 0, 'M', '0', '0', '', 'system', 'admin', '2021-07-07 18:29:57', '', NULL, '系统管理目录');
INSERT INTO `sys_menu` VALUES (100, '用户管理', 1, 1, 'user', 'system/user/index', 1, 0, 'C', '0', '0', 'system:user:list', 'user', 'admin', '2021-07-07 18:29:57', '', NULL, '用户管理菜单');
INSERT INTO `sys_menu` VALUES (101, '角色管理', 1, 2, 'role', 'system/role/index', 1, 0, 'C', '0', '0', 'system:role:list', 'peoples', 'admin', '2021-07-07 18:29:57', '', NULL, '角色管理菜单');
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 1, 3, 'menu', 'system/menu/index', 1, 0, 'C', '0', '0', 'system:menu:list', 'tree-table', 'admin', '2021-07-07 18:29:57', '', NULL, '菜单管理菜单');
INSERT INTO `sys_menu` VALUES (103, '部门管理', 1, 4, 'dept', 'system/dept/index', 1, 0, 'C', '0', '0', 'system:dept:list', 'tree', 'admin', '2021-07-07 18:29:57', '', NULL, '部门管理菜单');
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 1, 5, 'post', 'system/post/index', 1, 0, 'C', '0', '0', 'system:post:list', 'post', 'admin', '2021-07-07 18:29:57', '', NULL, '岗位管理菜单');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 1, 6, 'dict', 'system/dict/index', 1, 0, 'C', '0', '0', 'system:dict:list', 'dict', 'admin', '2021-07-07 18:29:57', '', NULL, '字典管理菜单');
INSERT INTO `sys_menu` VALUES (106, '参数设置', 1, 7, 'config', 'system/config/index', 1, 0, 'C', '0', '0', 'system:config:list', 'edit', 'admin', '2021-07-07 18:29:57', '', NULL, '参数设置菜单');
INSERT INTO `sys_menu` VALUES (107, '通知公告', 1, 8, 'notice', 'system/notice/index', 1, 0, 'C', '0', '0', 'system:notice:list', 'message', 'admin', '2021-07-07 18:29:57', '', NULL, '通知公告菜单');
INSERT INTO `sys_menu` VALUES (1001, '用户查询', 100, 1, '', '', 1, 0, 'F', '0', '0', 'system:user:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1002, '用户新增', 100, 2, '', '', 1, 0, 'F', '0', '0', 'system:user:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1003, '用户修改', 100, 3, '', '', 1, 0, 'F', '0', '0', 'system:user:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1004, '用户删除', 100, 4, '', '', 1, 0, 'F', '0', '0', 'system:user:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1005, '用户导出', 100, 5, '', '', 1, 0, 'F', '0', '0', 'system:user:export', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1006, '用户导入', 100, 6, '', '', 1, 0, 'F', '0', '0', 'system:user:import', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1007, '重置密码', 100, 7, '', '', 1, 0, 'F', '0', '0', 'system:user:resetPwd', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1008, '角色查询', 101, 1, '', '', 1, 0, 'F', '0', '0', 'system:role:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1009, '角色新增', 101, 2, '', '', 1, 0, 'F', '0', '0', 'system:role:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1010, '角色修改', 101, 3, '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1011, '角色删除', 101, 4, '', '', 1, 0, 'F', '0', '0', 'system:role:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1012, '角色导出', 101, 5, '', '', 1, 0, 'F', '0', '0', 'system:role:export', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1013, '菜单查询', 102, 1, '', '', 1, 0, 'F', '0', '0', 'system:menu:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1014, '菜单新增', 102, 2, '', '', 1, 0, 'F', '0', '0', 'system:menu:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1015, '菜单修改', 102, 3, '', '', 1, 0, 'F', '0', '0', 'system:menu:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1016, '菜单删除', 102, 4, '', '', 1, 0, 'F', '0', '0', 'system:menu:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1017, '部门查询', 103, 1, '', '', 1, 0, 'F', '0', '0', 'system:dept:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1018, '部门新增', 103, 2, '', '', 1, 0, 'F', '0', '0', 'system:dept:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1019, '部门修改', 103, 3, '', '', 1, 0, 'F', '0', '0', 'system:dept:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1020, '部门删除', 103, 4, '', '', 1, 0, 'F', '0', '0', 'system:dept:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1021, '岗位查询', 104, 1, '', '', 1, 0, 'F', '0', '0', 'system:post:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1022, '岗位新增', 104, 2, '', '', 1, 0, 'F', '0', '0', 'system:post:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1023, '岗位修改', 104, 3, '', '', 1, 0, 'F', '0', '0', 'system:post:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1024, '岗位删除', 104, 4, '', '', 1, 0, 'F', '0', '0', 'system:post:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1025, '岗位导出', 104, 5, '', '', 1, 0, 'F', '0', '0', 'system:post:export', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1026, '字典查询', 105, 1, NULL, '', 1, 0, 'F', '0', '0', 'system:dict:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1027, '字典新增', 105, 2, NULL, '', 1, 0, 'F', '0', '0', 'system:dict:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1028, '字典修改', 105, 3, NULL, '', 1, 0, 'F', '0', '0', 'system:dict:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1029, '字典删除', 105, 4, NULL, '', 1, 0, 'F', '0', '0', 'system:dict:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1030, '字典导出', 105, 5, NULL, '', 1, 0, 'F', '0', '0', 'system:dict:export', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1031, '参数查询', 106, 1, NULL, '', 1, 0, 'F', '0', '0', 'system:config:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1032, '参数新增', 106, 2, NULL, '', 1, 0, 'F', '0', '0', 'system:config:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1033, '参数修改', 106, 3, NULL, '', 1, 0, 'F', '0', '0', 'system:config:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1034, '参数删除', 106, 4, NULL, '', 1, 0, 'F', '0', '0', 'system:config:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1035, '参数导出', 106, 5, NULL, '', 1, 0, 'F', '0', '0', 'system:config:export', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1036, '公告查询', 107, 1, NULL, '', 1, 0, 'F', '0', '0', 'system:notice:query', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1037, '公告新增', 107, 2, NULL, '', 1, 0, 'F', '0', '0', 'system:notice:add', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1038, '公告修改', 107, 3, NULL, '', 1, 0, 'F', '0', '0', 'system:notice:edit', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (1039, '公告删除', 107, 4, NULL, '', 1, 0, 'F', '0', '0', 'system:notice:remove', '#', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2000, '日志管理', 1, 9, 'log', '', 1, 0, 'M', '0', '0', '', 'log', '', '2021-09-27 11:22:18', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2001, '操作日志', 2000, 1, 'operlog', 'monitor/operlog/index', 1, 0, 'C', '0', '0', 'monitor:operlog:list', 'form', '', '2021-09-27 11:23:47', 'admin', '2021-10-13 16:11:49', '');
INSERT INTO `sys_menu` VALUES (2002, '登录日志', 2000, 2, 'logininfor', 'monitor/logininfor/index', 1, 0, 'C', '0', '0', 'monitor:logininfor:list', 'logininfor', '', '2021-09-27 11:24:31', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2003, '操作查询', 2001, 1, '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:query', '', '', '2021-09-27 11:25:14', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2004, '操作删除', 2001, 2, '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:remove', '', '', '2021-09-27 11:25:39', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2005, '日志导出', 2001, 3, '', '', 1, 0, 'F', '0', '0', 'monitor:operlog:export', '', '', '2021-09-27 11:26:07', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2006, '登录查询', 2002, 1, '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:query', '', '', '2021-09-27 11:26:34', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2007, '登录删除', 2002, 2, '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:remove', '', '', '2021-09-27 11:26:54', '', NULL, '');
INSERT INTO `sys_menu` VALUES (2008, '日志导出', 2002, 3, '', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:export', '', '', '2021-09-27 11:27:09', '', NULL, '');

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
  `notice_id` int(4) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `notice_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告标题',
  `notice_type` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '公告类型（1通知 2公告）',
  `notice_content` longblob NULL COMMENT '公告内容',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '公告状态（0正常 1关闭）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知公告表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '请求参数',
  `json_result` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '返回参数',
  `status` int(1) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime(0) NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 67 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (33, '部门新增', 1, 'system:dept:add', 'POST', 1, 'acool', '运维部门', '/api/v1/system/dept/add', '10.10.24.5', '内网IP', '{\"parentId\":100,\"deptName\":\"阿斯顿\",\"orderNum\":111,\"status\":\"0\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 17:29:19');
INSERT INTO `sys_oper_log` VALUES (34, '部门修改', 2, 'system:dept:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dept/edit', '10.10.24.5', '内网IP', '{\"deptId\":210,\"ancestors\":\"0,100\",\"deptName\":\"阿斯顿11\",\"orderNum\":111,\"leader\":\"\",\"parentId\":100,\"phone\":\"\",\"status\":\"0\",\"email\":\"\",\"delFlag\":\"0\",\"createTime\":\"2021-10-09T17:29:19+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 17:39:42');
INSERT INTO `sys_oper_log` VALUES (35, '部门修改', 2, 'system:dept:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dept/edit', '10.10.24.5', '内网IP', '{\"deptId\":210,\"ancestors\":\"0,100\",\"deptName\":\"阿斯顿11\",\"orderNum\":111,\"leader\":\"\",\"parentId\":200,\"phone\":\"\",\"status\":\"0\",\"email\":\"\",\"delFlag\":\"0\",\"createTime\":\"2021-10-09T17:29:19+08:00\",\"createBy\":\"admin\",\"updateTime\":\"2021-10-09T17:39:42+08:00\",\"updateBy\":\"admin\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 17:39:52');
INSERT INTO `sys_oper_log` VALUES (36, '部门修改', 2, 'system:dept:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dept/edit', '10.10.24.5', '内网IP', '{\"deptId\":210,\"ancestors\":\"0,100\",\"deptName\":\"阿斯顿11\",\"orderNum\":111,\"leader\":\"111\",\"parentId\":100,\"phone\":\"\",\"status\":\"0\",\"email\":\"\",\"delFlag\":\"0\",\"createTime\":\"2021-10-09T17:29:19+08:00\",\"createBy\":\"admin\",\"updateTime\":\"2021-10-09T17:39:52+08:00\",\"updateBy\":\"admin\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 17:39:59');
INSERT INTO `sys_oper_log` VALUES (43, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/type/edit', '10.10.24.5', '内网IP', '{\"dictId\":1,\"dictName\":\"用户性别\",\"dictType\":\"sys_user_sex\",\"status\":\"0\",\"remark\":\"用户性别列表\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 18:09:16');
INSERT INTO `sys_oper_log` VALUES (44, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/data/edit', '10.10.24.5', '内网IP', '{\"dictCode\":19,\"dictSort\":2,\"dictLabel\":\"修改\",\"dictValue\":\"2\",\"dictType\":\"sys_oper_type\",\"isDefault\":\"N\",\"listClass\":\"success\",\"cssClass\":\"\",\"status\":\"0\",\"remark\":\"修改操作\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 19:22:37');
INSERT INTO `sys_oper_log` VALUES (45, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/data/edit', '10.10.24.5', '内网IP', '{\"dictCode\":21,\"dictSort\":4,\"dictLabel\":\"授权\",\"dictValue\":\"4\",\"dictType\":\"sys_oper_type\",\"isDefault\":\"N\",\"listClass\":\"info\",\"cssClass\":\"\",\"status\":\"0\",\"remark\":\"授权操作\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 19:22:47');
INSERT INTO `sys_oper_log` VALUES (47, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/data/edit', '10.10.24.5', '内网IP', '{\"dictCode\":19,\"dictSort\":2,\"dictLabel\":\"修改\",\"dictValue\":\"2\",\"dictType\":\"sys_oper_type\",\"isDefault\":\"N\",\"listClass\":\"warning\",\"cssClass\":\"\",\"status\":\"0\",\"remark\":\"修改操作\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"2021-10-09T19:22:37+08:00\",\"updateBy\":\"admin\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 19:29:02');
INSERT INTO `sys_oper_log` VALUES (48, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/data/edit', '10.10.24.5', '内网IP', '{\"dictCode\":22,\"dictSort\":5,\"dictLabel\":\"导出\",\"dictValue\":\"5\",\"dictType\":\"sys_oper_type\",\"isDefault\":\"N\",\"listClass\":\"success\",\"cssClass\":\"\",\"status\":\"0\",\"remark\":\"导出操作\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-09 19:29:11');
INSERT INTO `sys_oper_log` VALUES (49, '字典新增', 1, 'system:dict:add', 'POST', 1, 'acool', '运维部门', '/api/v1/system/dict/data/add', '10.10.24.5', '内网IP', '{\"dictLabel\":\"其他\",\"dictValue\":\"10\",\"listClass\":\"primary\",\"dictSort\":10,\"status\":\"0\",\"dictType\":\"sys_oper_type\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-11 10:08:24');
INSERT INTO `sys_oper_log` VALUES (50, '字典修改', 2, 'system:dict:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/dict/data/edit', '10.10.24.5', '内网IP', '{\"dictCode\":100,\"dictSort\":10,\"dictLabel\":\"其他\",\"dictValue\":\"10\",\"dictType\":\"sys_oper_type\",\"isDefault\":\"\",\"listClass\":\"primary\",\"cssClass\":\"\",\"status\":\"0\",\"remark\":\"其他类型\",\"createTime\":\"2021-10-11T10:08:24+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-11 10:08:33');
INSERT INTO `sys_oper_log` VALUES (51, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2002],\"deptIds\":null}', '', 1, '', '2021-10-11 10:13:40');
INSERT INTO `sys_oper_log` VALUES (52, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[],\"deptIds\":null}', '', 1, '', '2021-10-11 10:14:55');
INSERT INTO `sys_oper_log` VALUES (53, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2002],\"deptIds\":null}', '', 1, '', '2021-10-11 10:16:38');
INSERT INTO `sys_oper_log` VALUES (54, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[],\"deptIds\":null}', '', 1, '', '2021-10-11 10:21:46');
INSERT INTO `sys_oper_log` VALUES (55, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2002],\"deptIds\":null}', '', 1, '', '2021-10-11 10:28:59');
INSERT INTO `sys_oper_log` VALUES (56, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[],\"deptIds\":null}', '', 1, '', '2021-10-11 10:30:31');
INSERT INTO `sys_oper_log` VALUES (57, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,107,2000,2001,2002],\"deptIds\":null}', '', 1, '', '2021-10-11 10:30:42');
INSERT INTO `sys_oper_log` VALUES (58, '部门删除', 3, 'system:dept:remove', 'DELETE', 1, 'acool', '运维部门', '/api/v1/system/dept/remove/210', '10.10.24.5', '内网IP', '', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-11 10:30:58');
INSERT INTO `sys_oper_log` VALUES (59, '用户新增', 1, 'system:user:add', 'POST', 1, 'acool', '运维部门', '/api/v1/system/user/add', '10.10.24.5', '内网IP', '{\"deptId\":202,\"userName\":\"test\",\"nickName\":\"测试\",\"password\":\"123456\",\"status\":\"0\",\"postIds\":[4],\"roleIds\":[2]}', '', 1, '', '2021-10-11 11:02:08');
INSERT INTO `sys_oper_log` VALUES (60, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2002],\"deptIds\":null}', '', 1, '', '2021-10-13 09:58:41');
INSERT INTO `sys_oper_log` VALUES (61, '', 10, 'system:role:changeStatus', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/changeStatus', '10.10.24.5', '内网IP', '{\"roleId\":1,\"status\":\"1\"}', '', 1, '', '2021-10-13 10:01:59');
INSERT INTO `sys_oper_log` VALUES (62, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2002,2003,2006],\"deptIds\":null}', '', 1, '', '2021-10-13 11:30:40');
INSERT INTO `sys_oper_log` VALUES (63, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,2000,2001,2003,2004,2005,2002,2006,2007,2008],\"deptIds\":null}', '', 1, '', '2021-10-13 15:13:03');
INSERT INTO `sys_oper_log` VALUES (64, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,107,1036,2000,2001,2003,2004,2005,2002,2006,2007,2008],\"deptIds\":null}', '', 1, '', '2021-10-13 15:14:39');
INSERT INTO `sys_oper_log` VALUES (65, '菜单修改', 2, 'system:menu:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/menu/edit', '10.10.24.5', '内网IP', '{\"menuId\":2001,\"parentId\":2000,\"menuName\":\"操作日志\",\"orderNum\":1,\"path\":\"operlog\",\"menuType\":\"C\",\"visible\":\"0\",\"isFrame\":1,\"isCache\":0,\"perms\":\"monitor:operlog:list\",\"icon\":\"form\",\"remark\":\"\",\"createTime\":\"2021-09-27T11:23:47+08:00\",\"createBy\":\"\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"status\":\"0\",\"component\":\"monitor/operlog\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-13 16:04:52');
INSERT INTO `sys_oper_log` VALUES (66, '菜单修改', 2, 'system:menu:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/menu/edit', '10.10.24.5', '内网IP', '{\"menuId\":2001,\"parentId\":2000,\"menuName\":\"操作日志\",\"orderNum\":1,\"path\":\"operlog\",\"menuType\":\"C\",\"visible\":\"0\",\"isFrame\":1,\"isCache\":0,\"perms\":\"monitor:operlog:list\",\"icon\":\"form\",\"remark\":\"\",\"createTime\":\"2021-09-27T11:23:47+08:00\",\"createBy\":\"\",\"updateTime\":\"2021-10-13T16:04:52+08:00\",\"updateBy\":\"admin\",\"status\":\"0\",\"component\":\"monitor/operlog/index\"}', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-13 16:11:49');
INSERT INTO `sys_oper_log` VALUES (67, '角色修改', 2, 'system:role:edit', 'PUT', 1, 'acool', '运维部门', '/api/v1/system/role/edit', '10.10.24.5', '内网IP', '{\"roleId\":2,\"roleName\":\"普通角色\",\"roleKey\":\"common\",\"roleSort\":2,\"dataScope\":\"2\",\"menuCheckStrictly\":true,\"deptCheckStrictly\":true,\"status\":\"0\",\"delFlag\":\"0\",\"createTime\":\"2021-07-07T18:29:57+08:00\",\"createBy\":\"admin\",\"updateTime\":\"0001-01-01T00:00:00Z\",\"updateBy\":\"\",\"remark\":\"普通角色\",\"menuIds\":[1,107,2000,2001,2002,1036,2003,2006],\"deptIds\":null}', '', 1, '', '2021-10-13 18:34:58');
INSERT INTO `sys_oper_log` VALUES (68, '公告删除', 3, 'system:notice:remove', 'DELETE', 1, 'acool', '运维部门', '/api/v1/system/notice/remove/1', '10.10.24.5', '内网IP', '', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-14 10:02:06');
INSERT INTO `sys_oper_log` VALUES (69, '公告删除', 3, 'system:notice:remove', 'DELETE', 1, 'acool', '运维部门', '/api/v1/system/notice/remove/2', '10.10.24.5', '内网IP', '', '{\"status\":200,\"msg\":\"操作成功\",\"data\":null}', 0, '', '2021-10-14 10:02:08');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '岗位信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, '0', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理', 2, '0', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_post` VALUES (3, 'hr', '人力资源', 3, '0', 'admin', '2021-07-07 18:29:57', '', NULL, '');
INSERT INTO `sys_post` VALUES (4, 'user', '普通员工', 4, '0', 'admin', '2021-07-07 18:29:57', '', NULL, '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色权限字符串',
  `role_sort` int(4) NOT NULL COMMENT '显示顺序',
  `data_scope` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '1' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `menu_check_strictly` tinyint(1) NULL DEFAULT 1 COMMENT '菜单树选择项是否关联显示',
  `dept_check_strictly` tinyint(1) NULL DEFAULT 1 COMMENT '部门树选择项是否关联显示',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'admin', 1, '1', 1, 1, '0', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '超级管理员');
INSERT INTO `sys_role` VALUES (2, '普通角色', 'common', 2, '2', 1, 1, '0', '0', 'admin', '2021-07-07 18:29:57', '', NULL, '普通角色');

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (2, 100);
INSERT INTO `sys_role_dept` VALUES (2, 101);
INSERT INTO `sys_role_dept` VALUES (2, 105);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (2, 1);
INSERT INTO `sys_role_menu` VALUES (2, 107);
INSERT INTO `sys_role_menu` VALUES (2, 1036);
INSERT INTO `sys_role_menu` VALUES (2, 2000);
INSERT INTO `sys_role_menu` VALUES (2, 2001);
INSERT INTO `sys_role_menu` VALUES (2, 2002);
INSERT INTO `sys_role_menu` VALUES (2, 2003);
INSERT INTO `sys_role_menu` VALUES (2, 2006);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '部门ID',
  `user_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称',
  `user_type` varchar(2) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '00' COMMENT '用户类型（00系统用户）',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '用户邮箱',
  `phone_number` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '手机号码',
  `sex` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `avatar` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '头像地址',
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '密码',
  `status` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '帐号状态（0正常 1停用）',
  `del_flag` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '删除标志（0代表存在 2代表删除）',
  `login_ip` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime(0) NULL DEFAULT NULL COMMENT '最后登录时间',
  `create_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '创建者',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '更新者',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 103 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 103, 'admin', 'acool', '00', 'acool@163.com', '15888888888', '1', '', '$2a$10$JezvXdxT1grrVZkF51IZXuGDwp7XOS7MOpY9G48EJTaQ/R3hBYZs.', '0', '0', '127.0.0.1', '2021-07-07 18:29:57', 'admin', '2021-07-07 18:29:57', '', NULL, '管理员');
INSERT INTO `sys_user` VALUES (102, 202, 'test', '测试', '00', '', '', '', '', '$2a$10$n2jiNpkqwaKAIsQ8jAfzBeULhiqAaGZewRGXy3kPU/N.QUSz96hby', '0', '', '', NULL, '', '2021-10-11 11:02:08', '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 1);
INSERT INTO `sys_user_post` VALUES (102, 4);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (102, 2);

SET FOREIGN_KEY_CHECKS = 1;
