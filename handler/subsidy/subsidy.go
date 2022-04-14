package subsidy

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"time"
// 	"unicorn-files/fastjson"
// 	"unicorn-files/models/auth"
// 	"unicorn-files/models/namespace"
// 	ticket "unicorn-files/models/ticket"
// 	"unicorn-files/pkg/connection"
// 	"unicorn-files/pkg/logger"
// 	"unicorn-files/pkg/pagination"
// 	"unicorn-files/pkg/response/code"
// 	. "unicorn-files/pkg/response/response"
// 	"unicorn-files/pkg/tools"

// 	"github.com/360EntSecGroup-Skylar/excelize/v2"
// 	"github.com/gin-gonic/gin"
// 	uuid "github.com/satori/go.uuid"
// 	"github.com/spf13/viper"
// )

// /*
//   @Author : Mustang Gong
// */

// func GetGroupUserListHandler(c *gin.Context) {

// 	SearchParams := map[string]map[string]interface{}{
// 		"like": pagination.RequestParams(c),
// 	}

// 	type projectModel struct {
// 		namespace.Project
// 		NamespaceName string `json:"namespace_name"`
// 	}

// 	var projectList []*projectModel

// 	db := connection.DB.Self.Table("namespace_project").
// 		Joins("left join namespace_namespace on namespace_namespace.id = namespace_project.namespace_id").
// 		Select("namespace_project.id, namespace_project.namespace_id, namespace_project.name, namespace_namespace.name as namespace_name, namespace_project.description, namespace_project.twf_created, namespace_project.twf_modified").Scan(&projectList)

// 	result, err := pagination.Paging(&pagination.Param{
// 		C:  c,
// 		DB: db,
// 	}, &projectList, SearchParams, "namespace_project")

// 	if err != nil {
// 		Response(c, code.SelectProjectError, nil, err.Error())
// 		return
// 	}

// 	Response(c, nil, result, "获取所有项目列表成功")
// }

// /**
// 获取所有申请成功租房补贴的员工列表
// */
// func GetUserListHandler(c *gin.Context) {
// 	workflowId := c.Param("id")
// 	start, _ := c.GetQuery("startDate")
// 	end, _ := c.GetQuery("endDate")
// 	//fmt.Print(start)
// 	//fmt.Print(len(end))
// 	layout := "2006-01-02 15:04:05"
// 	startDate, err := time.Parse(layout, start)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		Response(c, code.ProcessingTicketDateError, nil, err.Error())
// 		return
// 	}
// 	endDate, err := time.Parse(layout, end)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		Response(c, code.ProcessingTicketDateError, nil, err.Error())
// 		return
// 	}
// 	//projectId := 1	// 人力资源
// 	//workflowId := 28 // 租房补贴
// 	type UserInfoModel struct {
// 		auth.User
// 		DepartName string `json:"depart_name"`
// 	}

// 	type ticketModel struct {
// 		ticket.Ticket
// 		UserInfo UserInfoModel `json:"user_info"`
// 	}

// 	var ticketList []*ticketModel
// 	err = connection.DB.Self.Table("ticket_ticket").Where("ticket_ticket.workflow_id = ? AND twf_created >= ? AND twf_created <= ?", workflowId, startDate, endDate).
// 		Scan(&ticketList).Error
// 	if err != nil {
// 		logger.Errorf("查询ticket数据失败: %s", err.Error())
// 		Response(c, code.SelectTicketError, nil, err.Error())
// 		return
// 	}

// 	if len(ticketList) <= 0 {
// 		logger.Error("没有可以下载的数据")
// 		Response(c, code.ProcessingTicketDataError, nil, "")
// 		return
// 	}

// 	// 准备文件
// 	baseFilePath := viper.GetString("filePath")
// 	fmt.Print(baseFilePath)
// 	// 创建当次主目录
// 	// 每个人一个目录
// 	// 把文件复制到个人目录
// 	// 数据导出成 exel
// 	// 打包要是文件
// 	uuidString := uuid.NewV4().String()
// 	uuidString = strings.ReplaceAll(uuidString, "-", "")[0:10]
// 	path := baseFilePath + "/" + uuidString + "/"
// 	_, err = tools.CreateDictByPath(path)
// 	if err != nil {
// 		logger.Error("mkdir failed![%v]\n", err)
// 		Response(c, code.ProcessingTicketCreateDictError, nil, err.Error())
// 		return
// 	}

// 	var sbList []*ticket.Subsidy
// 	for _, data := range ticketList {
// 		// 获取创建人信息
// 		err = connection.DB.Self.Model(&auth.User{}).Where("username = ?", data.CreateUser).Find(&data.UserInfo).Error
// 		if err != nil {
// 			logger.Errorf("获取创建人信息: %s", err.Error())
// 			Response(c, code.SelectUserError, nil, err.Error())
// 			return
// 		}

// 		// 获取部门数据
// 		departCount := 0
// 		err = connection.DB.Self.Model(&auth.Depart{}).Where("id = ?", data.UserInfo.Depart).Count(&departCount).Error
// 		if err != nil {
// 			logger.Errorf("获取部门数据: %s", err.Error())
// 			Response(c, code.SelectDepartError, nil, err.Error())
// 			return
// 		}

// 		if departCount > 0 {
// 			var department auth.Depart
// 			err = connection.DB.Self.Model(&auth.Depart{}).Where("id = ?", data.UserInfo.Depart).Find(&department).Error
// 			if err != nil {
// 				logger.Errorf("获取部门数据: %s", err.Error())
// 				Response(c, code.SelectDepartError, nil, err.Error())
// 				return
// 			}
// 			data.UserInfo.DepartName = department.Name
// 		} else {
// 			data.UserInfo.DepartName = ""
// 		}

// 		isEnd := "否"
// 		if data.IsEnd {
// 			isEnd = "是"
// 		}
// 		//tm := strconv.FormatInt(data.CreatedAt.Unix(), 10)
// 		tm := fmt.Sprint(data.Id)
// 		// 创建用户目录
// 		subPath := path + data.UserInfo.Nickname + "-" + tm + "/"
// 		tools.CreateDictByPath(subPath)

// 		// 读取 json 数据
// 		jsonArr, _ := fastjson.ParseArrayRawMessage(data.Fields)
// 		item := jsonArr.GetJsonObject(0).GetJsonObject("data")
// 		subsidyType, _ := item.GetInt("subsidy_type")
// 		subsidyMessage, _ := item.GetString("subsidy_message")

// 		subsidyFile1 := item.GetJsonArray("subsidy_file1")
// 		lenArr := len(subsidyFile1.Values())
// 		var fileAArr []string
// 		if subsidyFile1 != nil && lenArr > 0 {
// 			for i := 0; i < lenArr; i++ {
// 				file1 := subsidyFile1.GetJsonObject(i)
// 				url, _ := file1.GetString("url")
// 				var ss []string
// 				if url != "" {
// 					ss = strings.Split(url, "/")
// 				}

// 				fileId := ss[len(ss)-1]
// 				//fileId, _ := file1.GetString("uid")
// 				fileAArr = append(fileAArr, url)

// 				tools.CopyFile(baseFilePath+"/"+fileId, subPath+fileId)
// 			}
// 		}
// 		//fmt.Println('\n')
// 		subsidyFile2 := item.GetJsonArray("subsidy_file2")
// 		lenArr = len(subsidyFile2.Values())
// 		var fileBArr []string
// 		if subsidyFile2 != nil && lenArr > 0 {
// 			for i := 0; i < lenArr; i++ {
// 				file1 := subsidyFile2.GetJsonObject(i)
// 				url, _ := file1.GetString("url")
// 				var ss []string
// 				if url != "" {
// 					ss = strings.Split(url, "/")
// 				}

// 				fileId := ss[len(ss)-1]
// 				//fileId, _ := file1.GetString("uid")
// 				fileBArr = append(fileBArr, url)

// 				tools.CopyFile(baseFilePath+"/"+fileId, subPath+fileId)
// 			}
// 		}
// 		//fmt.Println('\n')
// 		subsidyFile3 := item.GetJsonArray("subsidy_file3")
// 		lenArr = len(subsidyFile3.Values())
// 		var fileCArr []string
// 		if subsidyFile3 != nil && lenArr > 0 {
// 			for i := 0; i < lenArr; i++ {
// 				file1 := subsidyFile3.GetJsonObject(i)
// 				url, _ := file1.GetString("url")
// 				var ss []string
// 				if url != "" {
// 					ss = strings.Split(url, "/")
// 				}

// 				fileId := ss[len(ss)-1]
// 				//fileId, _ := file1.GetString("uid")
// 				fileCArr = append(fileCArr, url)

// 				tools.CopyFile(baseFilePath+"/"+fileId, subPath+fileId)
// 			}
// 		}
// 		//fmt.Println('\n')
// 		rentalType := "中介"
// 		if subsidyType == 2 {
// 			rentalType = "房东"
// 		}
// 		denid := "否"
// 		if data.IsDenied {
// 			denid = "是"
// 		}
// 		sb := ticket.Subsidy{
// 			UserName:     data.UserInfo.Nickname,
// 			Depart:       data.UserInfo.DepartName,
// 			Leader:       data.UserInfo.ReportTo,
// 			Job:          data.UserInfo.Position,
// 			DateTime:     data.CreatedAt,
// 			IsEnd:        isEnd,
// 			IsDenied:     denid,
// 			RentalType:   rentalType,
// 			InfoMessage:  subsidyMessage,
// 			IntervalFile: strings.Join(fileAArr, " "),
// 			ContractFile: strings.Join(fileBArr, " "),
// 			RoomFile:     strings.Join(fileCArr, " "),
// 			ID:           tm,
// 		}

// 		sbList = append(sbList, &sb)
// 	}

// 	// 创建数据表
// 	if sbList != nil && len(sbList) > 0 {
// 		f := excelize.NewFile()
// 		// 设置表头
// 		head := []string{"申请人", "部门", "直属领导", "职位", "申请日期", "是否结束", "是否驳回", "租房类型", "申请说明", "距离截图", "租房合同", "房本照片", "ID"}

// 		for i, s := 1, 65; i < len(head)+1; i++ {
// 			ss := string(rune(s)) + strconv.Itoa(1)
// 			f.SetCellValue("Sheet1", ss, head[i-1])
// 			s++
// 		}
// 		// 填入数据
// 		for i := 0; i < len(sbList); i++ {
// 			item := sbList[i]
// 			s := i + 2
// 			f.SetCellValue("Sheet1", "A"+strconv.Itoa(s), item.UserName)
// 			f.SetCellValue("Sheet1", "B"+strconv.Itoa(s), item.Depart)
// 			f.SetCellValue("Sheet1", "C"+strconv.Itoa(s), item.Leader)
// 			f.SetCellValue("Sheet1", "D"+strconv.Itoa(s), item.Job)
// 			f.SetCellValue("Sheet1", "E"+strconv.Itoa(s), item.DateTime)
// 			f.SetCellValue("Sheet1", "F"+strconv.Itoa(s), item.IsEnd)
// 			f.SetCellValue("Sheet1", "G"+strconv.Itoa(s), item.IsDenied)
// 			f.SetCellValue("Sheet1", "H"+strconv.Itoa(s), item.RentalType)
// 			f.SetCellValue("Sheet1", "I"+strconv.Itoa(s), item.InfoMessage)
// 			f.SetCellValue("Sheet1", "J"+strconv.Itoa(s), item.IntervalFile)
// 			f.SetCellValue("Sheet1", "K"+strconv.Itoa(s), item.ContractFile)
// 			f.SetCellValue("Sheet1", "L"+strconv.Itoa(s), item.RoomFile)
// 			f.SetCellValue("Sheet1", "M"+strconv.Itoa(s), item.ID)
// 		}
// 		// 根据指定路径保存文件
// 		if err := f.SaveAs(path + "subsidy.xlsx"); err != nil {
// 			fmt.Println(err)
// 		}
// 	}

// 	p := baseFilePath + "/" + uuidString + ".zip"
// 	tools.Zip(path, p)
// 	tools.RemoveDictPath(path)

// 	Response(c, nil, viper.GetString("uploadUrl")+uuidString+".zip", "")
// }
