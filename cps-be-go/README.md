# go-staking-interaction 项目说明
# 一.项目目录
## 1. common通用文件
1.1. config配置:配置yaml,.env 和配置加载逻辑(获取项目根目录)
1.2. logger 初始化日志逻辑
1.3. redis锁实现

## 2. adapter
存放创建redis/mysql/eth各自client的代码；

## 3. cmd
存放stake, airdroperc 等的简单测试代码；

## 4. contracts
合约代码 combined-json 和生成的合约对应Go文件；


## 5. controller
控制层:处理请求和响应
## 6. service
业务层:处理业务逻辑
## 7. repository
持久层:数据存储和查询
## 8. model
数据对象
## 9. router
路由配置
## 10. middleware
中间件:JWT权限

## 11. utils
工具类

## 12. listener
### 12.1. 事件监听；
### 12.2. 扫块；sync_block 
- 获取区块
- 多协程处理transactions
- 分析trax的receipt, 分开处理ERC20交易和BNB交易
- 同步更新DB数据

### 12.3. 区块重组处理方式:
1. 扫块,保持差距；--30个左右
2. hash--receipt 状态

## 13. dto
数据传输对象（Data Transfer Object）: 
- 请求 DTO,
- 响应 DTO,
- 内部传输结构

## 14. docker
docker-compose.yml
- todo:部署文档
启动: docker-compose --env-file=../.env up -d <指定服务>
停止: docker-compose down <可选:指定服务>
注意:MySQL 容器首次启动时才会读取 MYSQL_ROOT_PASSWORD 并初始化密码。如果你之前启动过，数据卷（volume）里已经保存了旧数据，之后的启动会直接跳过初始化，沿用旧密码
停止: docker-compose down -v <可选:指定服务>

---
# 二.操作文档:
1. 准备部署合约地址 
2. 生成合约对应的Go文件
3. 准备redis和mysql数据库,基础数据表(自动?)
4. 启动最简单的cmd下main；
5. debug观察业务逻辑
6. eth-client信息
7. contractService信息
8. DB持久化操作,数据更新对比
9. 扫块逻辑(块信息,tran信息,event信息)