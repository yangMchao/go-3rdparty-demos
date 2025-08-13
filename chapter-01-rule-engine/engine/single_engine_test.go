package engine

import (
	"testing"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

func TestSingleEngine(t *testing.T) {
	// 创建一个新的规则引擎实例
	singleEngine := engine.NewGengine()

	// 创建一个规则构建器
	dataContext := context.NewDataContext()
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	// 从字符串构建规则
	err := ruleBuilder.BuildRuleFromString(`
		rule "rule1" "description" salience 10 
		begin
			// 规则内容在此
		end
    `)
	if err != nil {
		// 处理错误
	}

	err = singleEngine.Execute(ruleBuilder, true)
	if err != nil {
		// 处理错误
	}

}
