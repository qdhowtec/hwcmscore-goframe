package hwctl

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/qdhowtec/hwcmscore-goframe/hwreq"
	"github.com/qdhowtec/hwcmscore-goframe/hwutils"
	"reflect"
)

type BaseStoreController struct {
	CurUser   hwreq.CurUser
	ReqFind   interface{}
	RespFind  interface{}
	ReqLst    interface{}
	RespLst   interface{}
	ReqAdd    interface{}
	RespAdd   interface{}
	ReqModify interface{}
	ReqDelete interface{}
	Service   interface{}
}

func (c *BaseStoreController) Add(r *ghttp.Request) {
	c.setCurUser(r)
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	//通过反射类型接口 声明reqBody为预定义的请求结构体，进行空值验证
	reqBody := reflect.New(reflect.TypeOf(c.ReqAdd)) //结构体类型
	err := gconv.Scan(r.GetBody(), reqBody)
	if err != nil {
		panic(err.Error())
	}
	if err = g.Validator().Data(reqBody).Run(r.GetCtx()); err != nil {
		panic(err.Error())
	}
	//定义的一个map方便取参数使用
	req := g.Map{}
	err = gconv.Scan(reqBody, &req)
	if err != nil {
		panic(err.Error())
	}
	req["shop_id"] = c.CurUser.ShopId
	req["store_id"] = c.CurUser.StoreId
	//反射调用logic的方法
	serviceType := reflect.ValueOf(c.Service)
	result := serviceType.MethodByName("Add").Call([]reflect.Value{reflect.ValueOf(req)})
	gconv.Scan(result[1], &err)
	if err != nil {
		panic(err.Error())
	}
	resp := reflect.New(reflect.TypeOf(c.RespAdd)).Interface() //指针类型
	gconv.Scan(result[0], &err)
	if err != nil {
		panic(err.Error())
	}
	err = gconv.Scan(result[0], resp)
	if err != nil {
		panic(err.Error())
	}
	r.Response.WriteJson(hwutils.Success(resp))
}

func (c *BaseStoreController) Modify(r *ghttp.Request) {
	c.setCurUser(r)
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	//通过反射类型接口 声明reqBody为预定义的请求结构体，进行空值验证
	reqBody := reflect.New(reflect.TypeOf(c.ReqModify)) //结构体类型
	err := gconv.Scan(r.GetBody(), reqBody)
	if err != nil {
		panic(err.Error())
	}
	if err = g.Validator().Data(reqBody).Run(r.GetCtx()); err != nil {
		panic(err.Error())
	}
	//定义的一个map方便取参数使用
	req := g.Map{}
	err = gconv.Scan(reqBody, &req)
	if err != nil {
		panic(err.Error())
	}
	req["shop_id"] = c.CurUser.ShopId
	req["store_id"] = c.CurUser.StoreId
	//反射调用logic的方法
	serviceType := reflect.ValueOf(c.Service)
	result := serviceType.MethodByName("Modify").Call([]reflect.Value{reflect.ValueOf(req)})
	gconv.Scan(result[0], &err)
	if err != nil {
		panic(err.Error())
	}
	r.Response.WriteJson(hwutils.Success(g.Map{}))
}

func (c *BaseStoreController) Delete(r *ghttp.Request) {
	c.setCurUser(r)
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	//通过反射类型接口 声明reqBody为预定义的请求结构体，进行空值验证
	reqBody := reflect.New(reflect.TypeOf(c.ReqDelete)) //结构体类型
	err := gconv.Scan(r.GetBody(), reqBody)
	if err != nil {
		panic(err.Error())
	}
	if err = g.Validator().Data(reqBody).Run(r.GetCtx()); err != nil {
		panic(err.Error())
	}
	//定义的一个map方便取参数使用
	req := g.Map{}
	err = gconv.Scan(reqBody, &req)
	if err != nil {
		panic(err.Error())
	}
	req["shop_id"] = c.CurUser.ShopId
	req["store_id"] = c.CurUser.StoreId
	//反射调用logic的方法
	serviceType := reflect.ValueOf(c.Service)
	result := serviceType.MethodByName("Delete").Call([]reflect.Value{reflect.ValueOf(req)})
	gconv.Scan(result[0], &err)
	if err != nil {
		panic(err.Error())
	}
	r.Response.WriteJson(hwutils.Success(g.Map{}))
}

func (c *BaseStoreController) Find(r *ghttp.Request) {
	c.setCurUser(r)
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	//通过反射类型接口 声明reqBody为预定义的请求结构体，进行空值验证
	reqBody := reflect.New(reflect.TypeOf(c.ReqFind)) //结构体类型
	err := gconv.Scan(r.GetBody(), reqBody)
	if err != nil {
		panic(err.Error())
	}
	if err = g.Validator().Data(reqBody).Run(r.GetCtx()); err != nil {
		panic(err.Error())
	}
	//定义的一个map方便取参数使用
	req := g.Map{}
	err = gconv.Scan(reqBody, &req)
	if err != nil {
		panic(err.Error())
	}
	req["shop_id"] = c.CurUser.ShopId
	req["store_id"] = c.CurUser.StoreId
	//反射调用logic的方法
	serviceType := reflect.ValueOf(c.Service)
	result := serviceType.MethodByName("Find").Call([]reflect.Value{reflect.ValueOf(req)})
	resp := reflect.New(reflect.TypeOf(c.RespFind)).Interface() //指针类型
	gconv.Scan(result[1], &err)
	if err != nil {
		panic(err.Error())
	}
	err = gconv.Scan(result[0], resp)
	if err != nil {
		panic(err.Error())
	}
	r.Response.WriteJson(hwutils.Success(resp))
}

func (c *BaseStoreController) Lst(r *ghttp.Request) {
	c.setCurUser(r)
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	//通过反射类型接口 声明reqBody为预定义的请求结构体，进行空值验证
	reqBody := reflect.New(reflect.TypeOf(c.ReqLst)) //结构体类型
	err := gconv.Scan(r.GetBody(), reqBody)
	if err != nil {
		panic(err.Error())
	}
	if err = g.Validator().Data(reqBody).Run(r.GetCtx()); err != nil {
		panic(err.Error())
	}
	//定义的一个map方便取参数使用
	req := g.Map{}
	err = gconv.Scan(reqBody, &req)
	if err != nil {
		panic(err.Error())
	}
	req["shop_id"] = c.CurUser.ShopId
	req["store_id"] = c.CurUser.StoreId
	fmt.Println(req)
	//反射调用logic的方法
	serviceType := reflect.ValueOf(c.Service)
	result := serviceType.MethodByName("Lst").Call([]reflect.Value{reflect.ValueOf(req)})
	resp := reflect.New(reflect.TypeOf(c.RespLst)).Interface() //指针类型
	err = gconv.Scan(result[0], resp)
	if err != nil {
		panic(err.Error())
	}
	r.Response.WriteJson(hwutils.Success(resp))
}

func (c *BaseStoreController) setCurUser(r *ghttp.Request) {
	defer func() {
		if err := recover(); err != nil {
			r.Response.WriteJson(hwutils.Fail(err))
			return
		}
	}()
	gconv.Scan(r.Header.Get("shopAdminInfo"), &c.CurUser)
}
