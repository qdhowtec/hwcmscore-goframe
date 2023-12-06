/**
 * 基于goframe框架的逻辑层基类
 * author: liuzongchao
 * time: 2023-12-06
 * email: 328468168@qq.com
 */
package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/qdhowtec/hwcmscore-goframe/hwreq"
	"github.com/qdhowtec/hwcmscore-goframe/hwresp"
	"math"
	"reflect"
)

type BaseLogic struct {
	ModelName    string
	AddEntity    interface{}
	ModifyEntity interface{}
	FindResp     interface{}
	LstResp      interface{}
}

func (b *BaseLogic) Add(args g.Map) (idResp hwresp.IdResp, err error) {
	fmt.Println("add", args)
	ctx := gctx.New()
	args["create_time"] = gtime.Now().Timestamp()
	args["update_time"] = gtime.Now().Timestamp()

	row := reflect.New(reflect.TypeOf(b.AddEntity)) //反射获得类型
	gconv.Scan(args, row)                           //作用是为了排除非数据库字段
	err = g.Model(b.ModelName).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := g.Model(b.ModelName).Data(row).InsertAndGetId()
		if err != nil {
			return err
		}
		idResp.Id = int(id)
		return err
	})

	return
}
func (b *BaseLogic) Modify(args g.Map) (err error) {
	fmt.Println("modify", args)
	ctx := gctx.New()
	args["update_time"] = gtime.Now().Timestamp()
	row := reflect.New(reflect.TypeOf(b.ModifyEntity)) //反射获得类型
	//作用是为了排除非数据库字段
	err = g.Model(b.ModelName).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Model(b.ModelName).Where(g.Map{
			"id":       args["id"],
			"store_id": args["store_id"],
		}).Scan(row)
		if err == sql.ErrNoRows {
			err = errors.New("无效的ID")
			panic(err.Error())
		}
		if err != nil {
			panic(err.Error())
		}
		err = gconv.Scan(args, row)
		if err != nil {
			panic(err.Error())
		}
		_, err = g.Model(b.ModelName).Where(g.Map{
			"id":       args["id"],
			"store_id": args["store_id"],
		}).Data(row).Update()

		return err
	})
	return
}
func (b *BaseLogic) Delete(args g.Map) (err error) {
	fmt.Println("delete", args)
	ctx := gctx.New()
	saveData := reflect.New(reflect.TypeOf(b.ModifyEntity)) //反射获得类型
	gconv.Scan(args, saveData)                              //作用是为了排除非数据库字段
	err = g.Model(b.ModelName).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Model(b.ModelName).Where(g.Map{
			"id":       args["id"],
			"store_id": args["store_id"],
		}).Scan(saveData)
		if err == sql.ErrNoRows {
			err = errors.New("无效的ID")
			panic(err.Error())
		}
		if err != nil {
			panic(err.Error())
		}
		_, err = g.Model(b.ModelName).Where(g.Map{
			"id":      args["id"],
			"shop_id": args["shop_id"],
		}).Delete()

		return err
	})
	return
}
func (b *BaseLogic) Find(args g.Map) (resp interface{}, err error) {
	fmt.Println("find", args)
	row := reflect.New(reflect.TypeOf(b.ModifyEntity)) //反射获得类型
	err = g.Model(b.ModelName).OmitEmpty().Where(g.Map{
		"id":      args["id"],
		"shop_id": args["shop_id"],
	}).Scan(row)
	if err == sql.ErrNoRows {
		err = errors.New("无效的ID")
		return
	}
	if err != nil {
		return
	}
	resp = reflect.New(reflect.TypeOf(b.FindResp)) //反射获得类型
	gconv.Scan(row, resp)
	return
}
func (b *BaseLogic) Lst(args g.Map) (resp interface{}, err error) {
	fmt.Println("lst", args)
	list := reflect.New(reflect.TypeOf(b.FindResp))
	resp = reflect.New(reflect.TypeOf(b.LstResp)) //反射获得类型
	entity := hwreq.LstReq{}
	err = gconv.Scan(args, &entity)
	limit := 10
	if entity.Limit > 0 {
		limit = int(entity.Limit)
	}
	page := int(entity.Page)
	offset := (page - 1) * limit
	total := 0
	where := g.Map{"shop_id": args["shop_id"]}
	for _, v := range entity.Filter {
		switch v.Type {
		case "LIKE":
			where[v.Field+" like ?"] = "%" + v.Value + "%"
			break
		case "GT":
			where[v.Field+" > ?"] = v.Value
			break
		case "GTE":
			where[v.Field+" >= ?"] = v.Value
			break
		case "LT":
			where[v.Field+" < ?"] = v.Value
			break
		case "LTE":
			where[v.Field+" <= ?"] = v.Value
			break
		case "NEQ":
			where[v.Field+" != ?"] = v.Value
			break
		default:
			where[v.Field] = v.Value
		}
	}
	err = g.Model(b.ModelName).
		Where(where).
		Order("id", "desc").
		Limit(offset, limit).
		ScanAndCount(list, &total, true)

	result := g.Map{
		"list":      list,
		"total":     total,
		"curPage":   page,
		"limit":     limit,
		"pageTotal": int(math.Ceil(float64(total) / float64(limit))),
	}
	gconv.Scan(result, resp)
	return
}
