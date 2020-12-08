# go swoole thinkphp

## 此项目是框架 gin hyperf thinkphp的对比程序

### 可用 ab -n 3000 -c 1000 【http://xx.com/|redis|mysql 进行测试】

#### 进入gin目录
    1. mv .env.example .env
    2. vi .env 改成自己的链接
    2. go mod tidy
    3. go run main.go

#### 进入tp目录
    1. composer install
    2. mv .env.example .env
    3. vi .env 改成自己的连接

#### 进入hyperf
    1. composer install
    2. mv .env.example .env
    3. vi .env 改成自己的连接
    4. ./bin/hyperf.php start
