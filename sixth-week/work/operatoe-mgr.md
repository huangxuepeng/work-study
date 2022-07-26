### 功能概述
fhmc-operator用于读写用户操作日志，目前主要记录用户的登入登出操作以及用户对服务的启停、修改配置等操作
![](./%E5%9B%BE%E7%89%87.jpg)

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

  message LogRecord {
  string uid = 1;
  enum LogType{
    LOGIN = 0;
    LOGOUT = 1;
    ADD_APP = 2;
    DELETE_APP = 3;
    PUBLISH_APP = 4;
    INSTALL_SERVICE = 5;
    UNINSTALL_SERVICE = 6;
    UPDATE_SERVICE = 7;
    UPDATE_CONFIG = 8;
    STOP_SERVICE = 9;
    START_SERVICE = 10;
    ADD_IMAGE = 11;
    DELETE_IMAGE = 12;
    UPDATE_IMAGE = 13;
    UPDATE_APP = 14;
    LIST_SERVICE = 15;
	  ADD_DATA = 16;
	  ADD_EXTERNAL_DATA = 17;
	  UPDATA_DATA = 18;
	  DELETE_DATA = 19;
	  UPDATE_LABLE = 20;
	  EMPTY_NODE  = 21;
	  ADD_NET = 22;
	  UPDATE_NET = 23;
	  DELETE_NET = 24;
	  UPDATE_USRINFO = 25;
	  DELETE_USER = 26;
	  PASSWORD_RESETS = 27;
	  ADD_ROLE = 28;
	  UPDATE_ROLE = 29;
	  DELETE_ROLE = 30;
	  ADD_USER_PERMISSION = 31;
	  UPDATE_USER_PERMISSION = 32;
	  DELETE_USER_PERMISSION = 33;
	  PLANTFORM_BASE_SET = 34;
	  ANNOUNCEMENT_MANAGEMENT = 35;
	  CERTIFICATION_CENTER = 36;
	  UPLOAD_LICENSE = 37;
	  REQUEST_DOWNLOAD_FILE = 38;
    NOT_SCHEDULABLE = 39;
    SCHEDULABLE = 40;
  }
  LogType type = 2;
  // 来源，可以是uc/um/application
  enum SrcType{
    UC = 0;
    UM = 1;
    APPLICATION = 2;
    SERVICE_DIRECTORY = 3;
	  SERVICE_PACK = 4;
	  LOCAL_DATA = 5;
	  COMPUTER_RESOURCE_MANAGEMENT = 6;
	  NET_RESOURCE_MANAGEMENT = 7;
	  USER_MANAGEMENT = 8;
	  ROLE_MANAGEMENT = 9;
	  ABLE_MODELS = 10;
	  PLATFORM_SET = 11;
	  PLATFORM_INFO = 12;
  }
  SrcType srcType = 3;
  // 用户标识
  string userId = 4;
  // 操作数据的名称(例如安装应用app1，则dataName为app1[最好是显示的title便于用户识别理解，当然只是建议],若无操作数据，则为空)
  string dataName = 5;
  string msg = 6;
  // 执行操作时间
  int64 operatorTime = 7;
  // 操作日志创建时间
  int64 createTime = 8;
}
