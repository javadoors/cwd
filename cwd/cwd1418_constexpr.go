/**
 * (自动生成)CWD缺陷：CWD-1418
 * (自动生成)CWD场景：表达式始终为True或False
 * 
 * (必填)用例描述：展示表达式始终为True或False缺陷示例
 * (必填)规范参考：条件表达式应有实际判断意义
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

// 反例：表达式始终为True
func bad() {
	if 1 == 1 {
		fmt.Println("始终为真")
	}
}

// 正例：表达式有实际判断
func good() {
	value := 10
	if value > 5 {
		fmt.Println("条件判断有意义")
	}
}

func main() {
	bad()
	good()
}
