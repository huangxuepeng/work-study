package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/test/:id", test)
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler,
		"NAME_OF_ENV_VARIABLE"))
	r.Run(":8080")
}

// @Summary 获取日志列表
// @Description 获取用户所有的操作日志
// @Tags 操作审计
// @accept json
// @Produce json
// @Param page path []int true "页码"
// @Param limit path int true "每页多少条信息"
// @Param userId path int true "用户id"
// @Param mod query string false "模块 SrcType"
// @Param type path int true "操作 Type"
// @Failure 400 {object} IDD {"Id":100}  "能力模型不存在"
// @Success 200 {object} ListResponse
// @Router /test/{id} [get]
func test(c *gin.Context) {
	var d IDD
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"Msg":  "失败",
		})
		return
	}
	c.JSON(200, Response{
		Id:   23,
		Name: "sjansjk",
	})
}

type IDD struct {
	Id int
}
type Response struct {
	Id   int
	Name string
}

// 列表数据响应内容
type ListResponse struct {
	Code int
	List []*LogRecord
}

type LogRecord struct {
	Uid          int
	LogType      int32
	SrcType      int32
	UserId       string
	DataName     string
	OperatorTime int64
	CreateTime   int64
}

// message LogRecord {
// 	string uid = 1;
// 	enum LogType{
// 		LOGIN = 0;
// 		LOGINOUT = 1;
// 		ADD_SERVICE = 2;
// 		DELETE_SERVICE = 3;
// 		PUBLISH_SERVICE = 4;
// 		UPDATE_SERVICE = 5;
// 		INSTALL_SERVICE = 6;
// 		UNINSTALL_SERVICE = 7;
// 		UPDATE_RUNNING_CONFIG = 8;
// 		STOP_SERVICE = 9;
// 		START_SERVICE = 10;
// 		ADD_IMAGE = 11;
// 		DELETE_IMAGE = 12;
// 		UPDATE_IMAGE = 13;
// 		SET_BASE_IMAGE = 14;
// 		CANCEL_BASE_IMAGE = 15;
// 		ADD_OUTER_SERVICE = 16;
// 		DELETE_OUTER_SERVICE = 17;
// 		MODIFY_OUTER_SERVICE = 18;
// 		ADD_CONTRACT  = 19;
// 		DELETE_CONTRACT = 20;
// 		MODIFY_CONTRACT = 21;
// 		ADD_PACK_SERVICE = 22;
// 		MODIFY_PACK_SERVICE = 23;
// 		DELETE_PACK_SERVICE = 24;
// 		ADD_DATA = 25;
// 		MODIFY_DATA = 26;
// 		DELETE_DATA = 27;
// 		ADD_DATA_RESOURCE = 28;
// 		DELETE_DATA_RESOURCE = 29;
// 		MODIFY_DATA_RESOURCE = 30;
// 		EXPORT_DATA = 31;
// 		UPDATE_LABLE = 32;
// 		EMPTY_NODE = 33;
// 		NOT_SCHEDULABLE = 34;
// 		SCHEDULABLE = 35;
// 		ADD_IP_POOL = 36;
// 		MODIFY_IP_POOL = 37;
// 		DELETE_IP_POOL = 38;
// 		MODIFY_USRINFO = 39;
// 		DELETE_USER = 40;
// 		GENERATE_AUTH_CODE = 41;
// 		PASSWORD_RESET = 42;
// 		ADD_ROLE = 43;
// 		MODIFY_ROLE = 44;
// 		DELETE_ROLE = 45;
// 		ADD_USER_ABILITY = 46;
// 		MODIFY_USER_ABILITY = 47;
// 		DELETE_USER_ABILITY = 48;
// 		PLANTFORM_BASE_SET = 49;
// 		ANNOUNCEMENT_MANAGEMENT = 50;
// 		ADD_AUTH = 51;
// 		DELETE_AUTH = 52;
// 		UPLOAD_LICENSE = 53;
// 		REQUEST_DOWNLOAD_FILE = 54;
// 	}
// 	LogType type = 2;
// 	// 来源
// 	enum SrcType{
// 	  UC = 0;
// 	  SERVICE_MGR = 1;
// 	  INSTANCE_MGR = 2;
// 	  IMAGE_MGR = 3;
// 	  SERVICE_CATEGORY= 4;
// 	  OFFLINE_PKG = 5;
// 	  DATAUNIT = 6;
// 	  RESOURCE_MGR = 7;
// 	  NETWORK_MGR = 8;
// 	  UM = 9;
// 	  ACCESS_MGR = 10;
// 	  ABILITY_MGR = 11;
// 	  PLATROM_MGR = 12;
// 	  LICENSE_MGR = 13;
// 	}
// 	SrcType srcType = 3;
// 	// 用户标识
// 	string userId = 4;
// 	// 操作数据的名称(例如安装应用app1，则dataName为app1[最好是显示的title便于用户识别理解，当然只是建议],若无操作数据，则为空)
// 	string dataName = 5;
// 	string msg = 6;
// 	// 执行操作时间
// 	int64 operatorTime = 7;
// 	// 操作日志创建时间
// 	int64 createTime = 8;
//   }
