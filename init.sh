#!/bin/bash

# 提示用户输入项目信息
read -p "请输入项目名称: " PROJECT_NAME
read -p "请输入模块名称: " MODULE_NAME

# 定义一个函数来生成分隔符
get_delimiter() {
    local input="$1"
    local delimiters='@#%^&*-_+=|~'
    for (( i=0; i<${#delimiters}; i++ )); do
        local char="${delimiters:$i:1}"
        if [[ "$input" != *"$char"* ]]; then
            echo "$char"
            return
        fi
    done
    echo '@'
}

# 获取项目名称的分隔符
PROJECT_DELIMITER=$(get_delimiter "$PROJECT_NAME")
# 递归查找项目中的所有文件，并替换 your-app-name 为 $PROJECT_NAME
find . -type f -not -path './.git/*' -not -path './init.sh' -exec sed -i '' "s${PROJECT_DELIMITER}your-app-name${PROJECT_DELIMITER}$PROJECT_NAME${PROJECT_DELIMITER}g" {} +

# 获取模块名称的分隔符
MODULE_DELIMITER=$(get_delimiter "$MODULE_NAME")
# 递归查找项目中的所有文件，并替换 your-module-name 为 $MODULE_NAME
find . -type f -not -path './.git/*' -not -path './init.sh' -exec sed -i '' "s${MODULE_DELIMITER}your-module-name${MODULE_DELIMITER}$MODULE_NAME${MODULE_DELIMITER}g" {} +

# 初始化 Go 模块
go mod tidy

echo "项目初始化完成！"