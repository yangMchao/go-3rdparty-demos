package engine

import (
	"fmt"
	"testing"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

// User 代表我们系统中的客户
type User struct {
	Name           string
	Age            int64
	PurchaseAmount float64
	Discount       float64
	HasDiscount    bool
}

// SetHasDiscount 设置 HasDiscount 标志
func (u *User) SetHasDiscount(hasDiscount bool) {
	u.HasDiscount = hasDiscount
}

// 计算折扣后的最终金额
func (u *User) FinalAmount() float64 {
	return u.PurchaseAmount * (1 - u.Discount)
}

func TestQuickStart(t *testing.T) {
	// 定义我们的规则
	const rule = `
    rule "discount_rule" "确定用户是否有资格享受折扣" salience 10
    begin
        if User.Age > 60 || User.PurchaseAmount > 100 {
            User.Discount = 0.1
            User.SetHasDiscount(true)
            println("用户有资格享受折扣！")
        } else {
            User.Discount = 0
        }
        println("rule name :",@name)
        println("最终金额：", User.FinalAmount())
    end
    `

	// 创建一个用户
	user := &User{
		Name:           "John",
		Age:            65,
		PurchaseAmount: 80.0,
		Discount:       0,
		HasDiscount:    false,
	}

	// 创建数据上下文并添加数据/函数
	dataContext := context.NewDataContext()
	dataContext.Add("User", user)
	dataContext.Add("println", fmt.Println)

	// 创建规则构建器
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	// 从字符串构建规则
	err := ruleBuilder.BuildRuleFromString(rule)
	if err != nil {
		panic(err)
	}

	// 创建 gengine 实例
	eng := engine.NewGengine()

	// 执行规则
	err = eng.Execute(ruleBuilder, true)
	if err != nil {
		panic(err)
	}

	// 检查结果
	fmt.Printf("用户：%s，年龄：%d，购买金额：$%.2f\n",
		user.Name, user.Age, user.PurchaseAmount)
	fmt.Printf("折扣：%.0f%%，最终金额：$%.2f\n",
		user.Discount*100, user.FinalAmount())
}
