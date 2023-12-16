package handle

import (
	"Auto/model"
	"Auto/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataHandler struct {
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Time    string `json:"time"`
}

// Success 成功响应
func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Message = "请求成功！"
	r.Data = data
	r.Time = util.GetResTime()
	return r
}

// Fail 失败响应
func (r *Result) Fail(code int, msg string) *Result {
	r.Code = code
	r.Message = msg
	r.Time = util.GetResTime()
	return r
}

func New() *DataHandler {
	return &DataHandler{}
}

// GetUserDataBase GetUserLogs 动态查询用户字段
func (*DataHandler) GetUserDataBase(ctx *gin.Context) {
	res := &Result{}
	var req map[string]interface{}

	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}

	Table, ok := req["table_name"].(string)
	if ok {
		fmt.Println("断言成功")
		delete(req, "table_name") // 删除名为 table_name 的字段及其值
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换为 string 类型")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	value, err := model.SearchUserDataBase(req, Table) // 这里用你的用户名调用了之前的函数
	if err != nil {
		fmt.Println("获取不到数据，model异常", err)
		ctx.JSON(200, res.Fail(400, "获取不到数据，请重启服务或联系管理员"))
		return
	}
	ctx.JSON(200, res.Success(value))
}

// SetUserDataBase 动态更新用户表中字段 - 限定条件为表中必须有这个字段
func (*DataHandler) SetUserDataBase(ctx *gin.Context) {
	res := &Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	Table, ok := req["table_name"].(string)
	IDCache, ok := req["id"].(float64)
	var ID int
	if ok {
		ID = int(IDCache)
		fmt.Println("断言成功")
		delete(req, "table_name") // 删除名为 table_name 的字段及其值
		delete(req, "id")
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换类型，请查看断言函数")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	err := model.SetDataBase(req, Table, ID) // 这里用你的用户名调用了之前的函数
	if err != nil {
		fmt.Println("更新失败 - model异常", err)
		ctx.JSON(200, res.Fail(400, "更新失败，请重启服务或联系管理员"))
		return
	}
	ctx.JSON(200, res.Success("更新成功"))
}

// DeleteUserDataBase 实现动态删除操作
func (*DataHandler) DeleteUserDataBase(ctx *gin.Context) {
	res := &Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	Table, ok := req["table_name"].(string)
	if ok {
		fmt.Println("断言成功")
		delete(req, "table_name") // 删除名为 table_name 的字段及其值
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换类型，请查看断言函数")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	err := model.DeleteDataBase(req, Table)
	if err != nil {
		fmt.Println("删除失败 - model异常", err)
		ctx.JSON(200, res.Fail(400, "删除失败，请重启服务或联系管理员"))
		return
	}
	ctx.JSON(200, res.Success("删除成功"))
}
func (*DataHandler) InsetUserDataBase(ctx *gin.Context) {
	res := &Result{}
	var req map[string]interface{}
	// 通过 BindJSON 将 JSON 数据绑定到 map 中
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Printf("json数据错误 err => %s", err)
		ctx.JSON(http.StatusOK, res.Fail(400, "json数据错误！"))
		return
	}
	Table, ok := req["table_name"].(string)
	if ok {
		fmt.Println("断言成功")
		delete(req, "table_name") // 删除名为 table_name 的字段及其值
	} else {
		// 处理类型断言失败的情况
		fmt.Println("无法转换类型，请查看断言函数")
		ctx.JSON(200, res.Fail(400, "断言失败"))
		return
	}
	err := model.InSetDataBase(req, Table)
	if err != nil {
		fmt.Println("插入失败 - model异常", err)
		ctx.JSON(200, res.Fail(400, "插入新数据失败，请重启服务或联系管理员"))
		return
	}
	ctx.JSON(200, res.Success("插入新数据成功"))
}
