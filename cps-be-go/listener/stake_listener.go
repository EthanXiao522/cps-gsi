package listener

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"staking-interaction/adapter"
	"staking-interaction/common/config"
	"staking-interaction/contracts/stake"
	"staking-interaction/dto"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client    *ethclient.Client //保存客户端实例,用于关闭
	isRunning bool              //标记扫块循环是否运行
)

/*
1. 获取clientInfo(无auth和h账户信息)
2. 获取Stake合约地址和ABI
3. 获取监听的事件ID
4. 获取blockNumber(start和latest)
5. 创建过滤器查询事件,扫块获取事件日志
6. 解析事件日志,根据事件ID区分事件类型,并处理事件数据
7. 循环监听新块,重复步骤5-6
8. 提供停止监听的接口,设置isRunning=false,停止扫块循环,并关闭client连接
*/
func StartStakeEventListener() {
	conf := config.Get()
	clientInfo, err := adapter.NewSyncEthClient()
	if err != nil {
		log.Printf("ListenToStakeEvent: Failed to connect to the BSC network: %v", err)
	}
	defer clientInfo.CloseEthClient()

	contractABI, err := abi.JSON(strings.NewReader(stake.MTKStakeMetaData.ABI))
	if err != nil {
		log.Printf("ListenToStakeEvent: Failed to parse contract ABI: %v", err)
	}

	// 获取Stake事件ID
	stakeEventID := contractABI.Events[config.StakedEventName].ID
	withdrawEventID := contractABI.Events[config.WithdrawnEventName].ID
	log.Printf("Listening for events with stakeEventId:%s , withdrawEventId:%s", stakeEventID, withdrawEventID)

	client := clientInfo.Client
	//BlockNumber returns the most recent block number
	startBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("ListenToStakeEvent: Failed to get block number: %v", err)
	}
	log.Printf("Starting to listen for events from block number: %d", startBlock)

	contractAddress := common.HexToAddress(conf.BlockchainConfig.Contracts.Stake)

	isRunning = true
	// 循环监听事件
	for isRunning {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer func() {
			cancel() //确保每次循环结束时取消上下文，避免资源泄漏
			fmt.Println("一次循环结束")
		}()
		currentBlock, err := client.BlockNumber(ctx)
		if err != nil {
			log.Printf("ListenToStakeEvent: Failed to get block number, will retry after 10 seconds: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		endBlock := startBlock + uint64(conf.BlockchainConfig.Sync.BatchSize)
		if endBlock > currentBlock {
			endBlock = currentBlock
		}

		if endBlock > startBlock {
			log.Printf("Scanning blocks: %d to %d", startBlock, endBlock)
			// 构建过滤器查询事件,扫块获取事件日志
			logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(startBlock)),
				ToBlock:   big.NewInt(int64(endBlock)),
				Addresses: []common.Address{contractAddress},
				Topics:    [][]common.Hash{{stakeEventID, withdrawEventID}},
			})
			if err != nil {
				log.Printf("ListenToStakeEvent: Failed to filter logs, will retry after 10 seconds: %v", err)
				time.Sleep(10 * time.Second)
				continue
			}
			// 解析事件日志,根据事件ID区分事件类型,并处理事件数据
			for _, vLog := range logs {
				handleLog(vLog, dto.StakeEventListener{
					StakedEventId:    stakeEventID,
					WithdrawnEventId: withdrawEventID,
					ContractAddress:  contractAddress,
					ContractABI:      contractABI,
				})
			}

			startBlock = endBlock + 1
		} else {
			time.Sleep(10 * time.Second) //如果没有新块，等待一段时间再继续监听
		}
	}
	fmt.Println("for exit, Stake event listener has stopped.")
}

// 处理事件日志,根据事件ID区分事件类型,并处理事件数据
func handleLog(vLog types.Log, listener dto.StakeEventListener) {
	switch vLog.Topics[0] {
	case listener.StakedEventId:
		var stakedEvent stake.MTKStakeStaked
		err := listener.ContractABI.UnpackIntoInterface(&stakedEvent, config.StakedEventName, vLog.Data)
		if err != nil {
			log.Printf("Failed to unpack Staked event: %v", err)
			return
		}
		stakedEvent.User = common.HexToAddress(vLog.Topics[1].Hex())
		fmt.Printf("Staked event -blockNum:%d, User: %s, Amount: %s, stakeId:%s, Period: %d\n",
			vLog.BlockNumber,
			stakedEvent.User.Hex(),
			stakedEvent.Amount.String(),
			stakedEvent.StakeId.String(),
			stakedEvent.Period)

	case listener.WithdrawnEventId:
		var withdrawnEvent stake.MTKStakeWithdraw
		err := listener.ContractABI.UnpackIntoInterface(&withdrawnEvent, config.WithdrawnEventName, vLog.Data)
		if err != nil {
			log.Printf("Failed to unpack Withdrawn event: %v", err)
			return
		}
		withdrawnEvent.User = common.HexToAddress(vLog.Topics[1].Hex())
		fmt.Printf("Withdrawn event: block number: %d, transaction hash:%v, user:%v, StakedIndex=%s\n",
			vLog.BlockNumber,
			vLog.TxHash.Hex(),
			withdrawnEvent.User,
			withdrawnEvent.StakeId.String(),
		)

	default:
		log.Printf("Unknown event with ID: %s", vLog.Topics[0].Hex())
	}
}

// 停止监听函数,设置isRunning=false,停止扫块循环,并关闭client连接
func StopStakeEventListener() {
	if !isRunning {
		log.Println("Stake event listener is not running.")
		return
	}
	isRunning = false
	if client != nil {
		client.Close()
	}
	log.Println("Stake event listener stopped.")
}
