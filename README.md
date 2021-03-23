# leetcode-algo

[![Github-Actions](https://github.com/Saodd/leetcode-algo/workflows/Go/badge.svg?branch=master)](https://saodd.github.io/leetcode-algo/coverage.html)

（点击上面的icon可以跳转测试覆盖度报告哦）

## 版权声明：

项目代码中包含LeetCode算法原题和解答，版权归力扣和相关答题人所有。解题代码部分均为本人原创，使用MIT许可。

## 本地生成测试报告

```shell
go test -coverprofile=coverage.out ./...
go tool cover -html coverage.out -o coverage.html
```
