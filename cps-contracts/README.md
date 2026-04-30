## Foundry

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

-   **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
-   **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
-   **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
-   **Chisel**: Fast, utilitarian, and verbose solidity REPL.

## Documentation

https://book.getfoundry.sh/

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ forge script script/Counter.s.sol:CounterScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

---

### 1.部署命令
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

