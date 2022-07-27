### 功能概述
fhmc-operator用于读写用户操作日志，目前主要记录用户的登入登出操作以及用户对服务的启停、修改配置等操作
![](./%E5%8A%9F%E8%83%BD%E5%88%97%E8%A1%A8.jpg)
注:  
```
对接时除传递模块标识及操作标识外，还需要传递操作用户ID、操作数据名称（最好是title:version格式，便于用户阅读）、模块内对该操作定义的消息详情（msg），以及操作时间等字段，具体字段名可查看fhmc-apis/structs/opertor-log/v1
```


### http接口
http接口主要提供以下两种能力
- 读取操作记录
  - 操作记录列表将分页进行返回
  - 请求参数包括页码、还有limit
- 写操作记录
  - 提供特殊情况，可以通过http请求写入操作记录
同时提供rpc接口
- 写操作记录
  - 主要建议用户通过grpc接口写入操作日志

### 模块标识数据结构
```
enum SrcType{
    UC = 0;
    APPLICATION = 1;
    SERVICE_CATEGORY= 2;
    OFFLINE_PKG = 3;
    DATAUNIT = 4;
    RESOURCE_MGR = 5;
    NETWORK_MGR = 6;
    UM = 7;
    ACCESS_MGR = 8;
    ABILITY_MGR = 9;
    PLATROM_MGR = 10;
    LICENSE_MGR = 11;
}

```
### 操作分类数据结构
```
enum LogType{
    LOGIN = 0;
    LOGINOUT = 1;
    ADD_SERVICE = 2;
    DELETE_SERVICE = 3;
    PUBLISH_SERVICE = 4;
    UPDATE_SERVICE = 5;
    INSTALL_SERVICE = 6;
    UNINSTALL_SERVICE = 7;
    UPDATE_RUNNING_CONFIG = 8;
    STOP_SERVICE = 9;
    START_SERVICE = 10;
    ADD_IMAGE = 11;
    DELETE_IMAGE = 12;
    UPDATE_IMAGE = 13;
    SET_BASE_IMAGE = 14;
    CANCEL_BASE_IMAGE = 15;
    ADD_OUTER_SERVICE = 16;
    DELETE_OUTER_SERVICE = 17;
    MODIFY_OUTER_SERVICE = 18;
    ADD_CONTRACT  = 19;
    DELETE_CONTRACT = 20;
    MODIFY_CONTRACT = 21;
    ADD_PACK_SERVICE = 22;
    MODIFY_PACK_SERVICE = 23;
    DELETE_PACK_SERVICE = 24;
    ADD_DATA = 25;
    MODIFY_DATA = 26;
    DELETE_DATA = 27;
    ADD_DATA_RESOURCE = 28;
    DELETE_DATA_RESOURCE = 29;
    MODIFY_DATA_RESOURCE = 30;
    EXPORT_DATA = 31;
    UPDATE_LABLE = 32;
    EMPTY_NODE = 33;
    NOT_SCHEDULABLE = 34;
    SCHEDULABLE = 35;
    ADD_IP_POOL = 36;
    MODIFY_IP_POOL = 37;
    DELETE_IP_POOL = 38;
    MODIFY_USRINFO = 39;
    DELETE_USER = 40;
    GENERATE_AUTH_CODE = 41;
    PASSWORD_RESET = 42;
    ADD_ROLE = 43;
    MODIFY_ROLE = 44;
    DELETE_ROLE = 45;
    ADD_USER_ABILITY = 46;
    MODIFY_USER_ABILITY = 47;
    DELETE_USER_ABILITY = 48;
    PLANTFORM_BASE_SET = 49;
    ANNOUNCEMENT_MANAGEMENT = 50;
    ADD_AUTH = 51;
    DELETE_AUTH = 52;
    UPLOAD_LICENSE = 53;
    REQUEST_DOWNLOAD_FILE = 54;
}
```