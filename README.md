# 项目说明书
项目分为2个子项目:

## 1. cps-contracts 存放solidity-foundry框架的智能合约项目
### 1.1. 在bsc测试链上部署一个ERC20合约; --完成
### 1.2. 在bsc测试链上实现一个质押合约;


## 2. cps-be 存放go-gin框架的后端项目,与智能合约交互





---

## 1.cps-contracts部署命令
```Bash
forge script script/DMyToken.s.sol \
--broadcast \
--rpc-url testnet-bsc \
--private-key $PRIVATE_KEY \
--verify \
--verifier etherscan \
--etherscan-api-key $ETHERSCAN_API_KEY
```


### 任务一. MyToken部署bsc-testnet地址:
0x5e6BBd0266c1E69F0F384ed0EEd7BB982C732d1E

### 任务二. MTKStake部署bsc-testnet地址:
0xBa5C7119278d38BeE07027a9318F73fA3561eD42

### 质押MTK测试的操作步骤:
1. 部署MyToken合约,得到合约地址
2. 部署MTKStake合约(使用MyToken合约地址)
3. 给MTKStake转移10*10**18代币, 给MTKStake授权(后续account质押)
4. 查询account质押前的BalanceOf
5. 质押部分MTK
6. 到期前提取MTK(Revert)
7. 到期后提取MTK
8. 查询account质押后的BalanceOf, 对比前后获取的质押奖励；

**其他测试结果见bsc-testnet上合约的Transactions列表详情；**

### vscode提示solidity错误,但编译一切正常问题解决
1. 生成 remappings.txt 文件
`forge remappings > remappings.txt`

这个命令会把 Foundry 的依赖路径映射（比如 @openzeppelin/=lib/openzeppelin-contracts/）导出到 remappings.txt 文件中。VS Code 的 Solidity 插件会自动读取这个文件，从此就能正确解析 import 语句。


### 任务三.go代码实现与质押合约交互
#### 1. 根据智能合约编译out,生成对应go合约
1.1. 编写utils/Convert.go 用于生成合约的combined-json
1.2. 安装abigen工具,使用v2版本g命令,h使用combined-json生成合约对应go文件
```Bash
cd utils

#生成合约 MyToken 的combined-json和go合约文件
go run Convert.go MyToken
abigen --v2 --combined-json ../contracts/MyToken-combined.json --pkg contracts --type MyToken --out ../contracts/MyToken.go

#生成合约MTKStake的combined-json和go合约文件
go run Convert.go MTKStake
abigen --v2 --combined-json ../contracts/MTKStake-combined.json --pkg contracts --type MTKStake --out ../contracts/MTKStake.go
```

