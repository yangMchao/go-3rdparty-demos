

# quick 拆解示例

    定义数据结构：我们创建了一个带有属性和方法的 User 结构体
    创建规则：我们定义了一个带有条件和动作的规则
    注入数据和函数：我们将用户实例和 println 函数添加到数据上下文中
    构建规则：我们使用 RuleBuilder 编译规则字符串
    执行规则：我们运行规则，根据条件修改用户数据

```text
Copy code
rule "rule_name" "rule_description" salience priority_number
begin
    // 规则逻辑在此
end
```
`salience 值确定规则优先级——值越高，越先执行。`

