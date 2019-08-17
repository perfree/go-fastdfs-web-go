## GoFastDfs数据库设计文档

### 用户表 

t_user

|    字段    |   类型   | 长度 | 空值 | 主键 |           注释            |
| :--------: | :------: | :--: | :--: | :--: | :-----------------------: |
|   userId   | INTEGER  |  0   |  否  |  是  |           主键            |
|  userName  | VARCHAR2 |  32  |  否  |  否  |      用户名也是账户       |
|  password  | VARCHAR2 |  64  |  否  |  否  |       加密后的密码        |
|    salt    | VARCHAR2 |  32  |  否  |  否  |           盐值            |
|   email    | VARCHAR2 |  32  |  是  |  否  |       邮箱也是账户        |
|   state    | INTEGER  |  1   |  否  |  否  | 状态，0正常，1禁用，默认0 |
| createTime | VARCHAR2 |  32  |  否  |  否  |         创建时间          |
| updateTime | VARCHAR2 |  32  |  是  |  否  |         更新时间          |

### 用户角色关联表

t_user_role

|  字段  |  类型   | 长度 | 空值 | 主键 |  注释  |
| :----: | :-----: | :--: | :--: | :--: | :----: |
| userId | INTEGER |  0   |  否  |  否  | 用户id |
| roleId | INTEGER |  0   |  否  |  否  | 角色id |

### 角色表

t_role

|    字段    |   类型   | 长度 | 空值 | 主键 |                 注释                  |
| :--------: | :------: | :--: | :--: | :--: | :-----------------------------------: |
|   roleId   | INTEGER  |  0   |  否  |  是  |                 主键                  |
|  roleName  | VARCHAR2 |  32  |  否  |  否  |               角色名称                |
| roleValue  | VARCHAR2 |  32  |  否  |  否  |                角色值                 |
|  isDelete  | integer  |  1   |  否  |  否  | 是否删除，0：未删除，1：已删除，默认0 |
| createTime | VARCHAR2 |  32  |  否  |  否  |               创建时间                |
| updateTime | VARCHAR2 |  32  |  是  |  否  |               更新时间                |

### 角色权限关联表

t_role_permission

|     字段     |  类型   | 长度 | 空值 | 主键 |  注释  |
| :----------: | :-----: | :--: | :--: | :--: | :----: |
|    roleId    | INTEGER |  0   |  否  |  否  | 角色ID |
| permissionId | INTEGER |  0   |  否  |  否  | 权限ID |

### 权限表

t_permission

|      字段       |   类型   | 长度 | 空值 | 主键 |                      注释                      |
| :-------------: | :------: | :--: | :--: | :--: | :--------------------------------------------: |
|  permissionId   | INTEGER  |  0   |  否  |  是  |                      主键                      |
| permissionName  | VARCHAR2 |  32  |  否  |  否  |                     权限名                     |
| permissionValue | VARCHAR2 |  32  |  否  |  否  |                     权限值                     |
|      type       | INTEGER  |  1   |  否  |  否  | 类型，0：一级菜单，1：二级菜单，2：按钮，默认0 |
|    parentId     | INTEGER  |  0   |  否  |  否  |           父权限ID，0位父权限，默认0           |
|    isDelete     | INTEGER  |  1   |  否  |  否  |     是否删除，0：未删除，1：已删除，默认0      |
|   createTime    | VARCHAR2 |  32  |  否  |  否  |                    创建时间                    |
|   updateTime    | VARCHAR2 |  32  |  是  |  否  |                    更新时间                    |

### 角色集群关联表

t_role_peers

|  字段   |  类型   | 长度 | 空值 | 主键 |  注释  |
| :-----: | :-----: | :--: | :--: | :--: | :----: |
| roleId  | INTEGER |  0   |  否  |  否  | 角色ID |
| peersId | INTEGER |  0   |  否  |  否  | 集群ID |

## 集群表

t_peers

|      字段       |   类型   | 长度 | 空值 | 主键 |    注释     |
| :-------------: | :------: | :--: | :--: | :--: | :---------: |
|     peeerId     | INTEGER  |  0   |  否  |  是  |    主键     |
|    peersName    | VARCHAR2 |  32  |  否  |  否  |  集群名称   |
|    groupName    | VARCHAR2 |  32  |  是  |  否  |   组名称    |
|  peersAddress   | VARCHAR2 | 128  |  否  |  否  |  集群地址   |
| peersUrlMapping | VARCHAR2 | 128  |  是  |  否  | 集群映射url |
|   createTime    | VARCHAR2 |  32  |  否  |  否  |  创建时间   |
|   updateTime    | VARCHAR2 |  32  |  是  |  否  |  更新时间   |

### 日志表

t_log

|     字段     |   类型   | 长度 | 空值 | 主键 |         注释          |
| :----------: | :------: | :--: | :--: | :--: | :-------------------: |
|    logId     | INTEGER  |  0   |  否  |  是  |         主键          |
|     type     | INTEGER  |  1   |  否  |  否  | 类型,默认0，具体待定  |
|    action    | VARCHAR2 | 256  |  否  |  否  | 动作，默认0，具体待定 |
|   message    | VARCHAR2 | 256  |  否  |  否  |       提示信息        |
| errorMessage |   TEXT   |  0   |  是  |  否  |       错误信息        |
|  createTime  | VARCHAR2 |  32  |  否  |  否  |       创建时间        |

