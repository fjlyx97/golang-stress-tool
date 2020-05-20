# Go语言实现的简单压测工具
使用Go语言编写的简单压测工具

# 使用
go run main.go -n 总请求数量 -c 并发数 -f json文件路径

# json文件示例
- timeout默认为30秒
- keep_alive默认为false

## GET请求
```json
{
    "method" : "GET",
    "url" : "https://www.qq.com/",
    "timeout" : 5,
    "keep_alive" : false 
}
```

## POST请求
```json
{
    "method" : "POST",
    "url" : "http://www.xx.com/query/",
    "content_type" : "application/json",
    "timeout" : 5,
    "keep_alive" : false,
    "post_data" : {
        "test": {
            "testText": ["Hello"]
        }
    }
}
```
