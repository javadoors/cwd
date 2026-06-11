# cwd
代码缺陷字典（CWD：Code Weakness Dictionary）旨在识别和分类软件代码的安全性、可靠性、规范性。

需要给出golang语言的CWD缺陷正例与反例，缺陷列表有：
- CWD-1415:空代码块
- CWD-1416:死代码
- CWD-1417:变量未使用
- CWD-1418:表达式始终为True或False
- CWD-1429:使用被禁止的代码

```go
/**
 * (自动生成)CWD缺陷：
 * (自动生成)CWD场景：
 * 
 * (必填)用例描述：对外提供的查询接口，返回的响应体中包含密码等第三信息
 * (必填)规范参考：界面显示的接口响应中不能包含密码等敏感信息
 * (选填)排查指导：
 * (选填)排查指导：
 * (选填)修复建议指导：
 * (必填)贡献产品线：公共开发部
 * (必填)贡献部门：星火部
 * (必填)贡献者：zhangchao
 * (必填)污点标注：无
 * (必填)规则模型配置：
 */
```
## 代码样例
cwd1415_emptyblock.go
```go
/**
 * (自动生成)CWD缺陷：
 * (自动生成)CWD场景：
 * 
 * (必填)用例描述：xxx
 * (必填)规范参考：xxx
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

// 反例：空代码块
func bad() {
    if true {
        // 什么都没做
    }
}

// 正例：代码块有实际逻辑
func good() {
    if true {
        fmt.Println("条件满足，执行逻辑")
    }
}

func main() {
    bad()
    good()
}
```
- 正例函数命名为good()
- 反例函数命名为bad()

## 目录结构
```txt
project-name/
├── cmd/                  # 主程序入口
│   └── app/
│       └── main.go
├── pkg/                  # 公共库
├── internal/             # 内部库
├── api/                  # API 定义
├── configs/              # 配置文件
├── deployments/          # 部署文件
├── scripts/              # 构建/检测脚本
│   └── lint.sh           # 静态检查脚本
├── docs/                 # 文档
│   └── CWD.md            # 缺陷字典说明
├── cwd/                  # ✅ 缺陷检测与示例代码
│   ├── cwd1415_emptyblock.go
│   ├── cwd1416_deadcode.go
│   ├── cwd1417_unusedvar.go
│   ├── cwd1418_constexpr.go
│   └── cwd1429_forbidden.go
└── tools/                # 辅助工具
    └── cwd-checker/      # 自定义缺陷检测工具
```

