/**
 * (自动生成)CWD缺陷：CWD-1417
 * (自动生成)CWD场景：变量未使用
 * 
 * (必填)用例描述：展示变量未使用缺陷示例
 * (必填)规范参考：声明的变量应当被使用
 * (选填)排查指导：
 * (选填)排查指导：
 * (选填)修复建议指导：
 * (必填)贡献产品线：公共开发部
 * (必填)贡献部门：星火部
 * (必填)贡献者：zhangchao
 * (必填)污点标注：无
 * (必填)规则模型配置：
 */
package main

import "fmt"

// 反例：变量未使用
func bad() {
	unusedVar := "不会被使用"
	fmt.Println("hello")
	_ = unusedVar
}

// 正例：变量被使用
func good() {
	usedVar := "被使用"
	fmt.Println(usedVar)
}

func main() {
	bad()
	good()
}
