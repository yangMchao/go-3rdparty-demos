package engine

import (
	"fmt"
	"testing"

	"github.com/bilibili/gengine/engine"
)

// 定义一个简单的数据结构
type Request struct {
	UserAge  int
	UserName string
}

type Response struct {
	Score int
	Level string
}

func TestPool(t *testing.T) {
	// 定义一个简单规则
	rule := `
    rule "calculate_score" "根据年龄计算分数" salience 10
    begin
        if Request.UserAge > 20 {
            Response.Score = 100
            Response.Level = "高级"
        } else {
            Response.Score = 50
            Response.Level = "初级"
        }
    end
    `

	// 定义规则中使用的 API
	apis := make(map[string]interface{})

	// 创建一个最小 10 个、最大 20 个引擎的 GenginePool
	// 使用 SortModel (1) 作为执行模型
	pool, err := engine.NewGenginePool(10, 20, engine.ConcurrentModel, rule, apis)
	if err != nil {
		panic(err)
	}

	// 处理请求
	req := &Request{UserAge: 25, UserName: "Alice"}
	resp := &Response{}

	// 使用池执行规则
	err, _ = pool.ExecuteRulesWithSpecifiedEM("Request", req, "Response", resp)
	if err != nil {
		panic(err)
	}
	//println(resultMap)

	// 检查结果
	fmt.Printf("用户 %s 获得分数 %d 和级别 %s\n",
		req.UserName, resp.Score, resp.Level)
}
