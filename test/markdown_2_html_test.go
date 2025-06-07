package test

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"os"
	"testing"
)

func TestPkg(t *testing.T) {
	// 1. Markdown 内容
	md := "# Code Review\n\n## 修改的文件\n\n本次代码变更涉及以下文件：\n\n1. `api/controller/build.go` - 删除\n2. `api/controller/webhook.go` - 新增\n3. `api/request/demo.go` - 删除\n4. `api/request/webhook.go` - 新增\n5. `common/app/response.go` - 修改\n6. `common/errcode/err.go` - 修改\n7. `go.mod` - 修改\n8. `go.sum` - 修改\n9. `logic/service/demo.go` - 删除\n10. `logic/service/client.go` - 修改\n11. `logic/service/webhook.go` - 新增\n12. `test/kafka_test.go` - 新增\n13. `wire.go` - 修改\n\n## 潜在风险分析\n\n1. **删除旧功能引入的风险**：\n   - 删除了`build.go`和相关的`demo.go`文件，但没有看到替代功能的完整实现，可能导致现有功能缺失\n   - `wire.go`中移除了依赖注入配置，可能导致服务启动失败\n\n2. **错误处理变更风险**：\n   - `response.go`中的`Error`方法现在直接接受`error`接口而非`*errcode.AppError`，可能导致某些错误处理逻辑不一致\n   - 新的错误处理方式可能无法正确捕获和转换所有类型的错误\n\n3. **Kafka集成风险**：\n   - 新增的`kafka_test.go`引入了Kafka依赖，但未看到生产环境配置和错误处理机制\n   - 测试代码中的硬编码IP地址(\":9092\")不利于维护和不同环境的部署\n\n4. **Webhook实现不完整**：\n   - `webhook.go`中的`ProcessHook`方法只有panic占位符，没有实际实现\n   - 缺少必要的验证和安全性考虑\n\n## 优化建议\n\n1. **错误处理改进**：\n   - 在`response.go`中，考虑添加对原始错误的日志记录，便于调试\n   - 可以添加更详细的错误分类处理逻辑\n\n2. **Kafka配置优化**：\n   - 将Kafka broker地址提取到配置文件中\n   - 添加连接重试和错误恢复机制\n   - 考虑添加异步生产者模式以提高性能\n\n3. **Webhook实现建议**：\n   - 添加请求签名验证\n   - 实现完整的业务逻辑处理\n   - 考虑添加速率限制防止滥用\n\n4. **代码组织改进**：\n   - 将Kafka相关代码移动到专门的包中，而不是放在test目录下\n   - 考虑添加接口文档和示例\n\n5. **依赖管理**：\n   - 检查新增的Kafka依赖是否都是必要的，避免引入过多间接依赖\n\n## 总结\n\n本次变更主要重构了错误处理机制，引入了Webhook支持，并添加了Kafka集成。需要注意错误处理方式的变更可能影响现有代码，同时新功能的实现还不完整。建议完善Webhook和Kafka的实现，并确保错误处理的一致性和可靠性。"

	// 2. 转换为 HTML
	html := markdown.ToHTML([]byte(md), nil, nil)

	// 3. 包装为完整 HTML 页面（加上基本结构）
	fullHTML := `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Markdown to HTML</title>
</head>
<body>
` + string(html) + `
</body>
</html>`

	// 4. 保存为本地文件
	err := os.WriteFile("output.html", []byte(fullHTML), 0644)
	if err != nil {
		fmt.Println("❌ 写入文件失败:", err)
		return
	}

	fmt.Println("✅ 成功将 HTML 文件保存为 output.html")
}
