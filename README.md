# AI 自动代码审查

## 简介
本项目专注于利用LLM实现代码的自动审查。通过先进的大模型的代码审查能力，我们的系统能够快速、准确地分析代码，发现潜在的问题，提升代码质量。

## 项目架构
![image.png](https://raw.githubusercontent.com/lyydsheep/pic/main/20250607160400.png)

## 功能特性
- 语法检查：自动识别代码中的语法错误。
- 风格检查：确保代码符合团队的编码规范。
- 安全审查：检测代码中的安全漏洞。
- 性能分析：评估代码的性能表现

## TODO
- [ ] 功能模块设计与开发  
- [ ] 数据库表结构设计  
- [ ] 接口文档编写  
- [ ] 单元测试覆盖  
- [ ] 部署与持续集成配置  
## 技术栈
- 后端：Go 语言，使用 Gin 框架构建 API。
- 人工智能：集成LLM进行代码分析。
- 依赖管理：Go Modules。

## 快速开始
1. 克隆仓库：`git clone https://github.com/lyydsheep/ai_code_review.git`
2. 安装依赖：`go mod tidy`
3. 启动服务：`go run main.go`

## 贡献指南
如果您想为项目做出贡献，请遵循以下步骤：
1. Fork 本仓库。
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)。
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)。
4. 将更改推送到分支 (`git push origin feature/AmazingFeature`)。
5. 打开一个 Pull Request。