package main

import (
	"flag"
	"fmt"
	"math/big"
	"staking-interaction/adapter"
	"staking-interaction/service"
	"time"
)

/*
*
测试stake和withdraw方法(通过StakeService调用MTKStake合约的Stake和Withdraw方法进行测试)
1. 运行前请先部署MTKStake合约，并将合约地址更新到config.yaml中
2. 运行前请确保账户中有足够的MTK代币进行质押测试(需要先approve授权MTKStake合约可以使用你的MTK代币)
3. stake前,查询账户UserStake(address,index)
4. 执行stake方法,输入质押数量和质押周期,运行后会在控制台输出交易哈希、合约地址、调用方法等信息，方便验证交易是否成功
5. stake后,查询账户UserStake(address,index)验证质押信息是否正确
6. 执行withdraw方法,输入质押ID,运行后会在控制台输出交易哈希、合约地址、调用方法等信息，方便验证交易是否成功
7. withdraw后,查询账户UserStake(address,index)验证质押信息是否正确
*/
func main() {

	stakeAmountFlag := flag.Int64("stakeAmount", 0, "质押数量（最小单位，如wei）")
	stakePeriodFlag := flag.Uint("stakePeriod", 0, "质押周期（单位：天）")
	stakeIndexFlag := flag.Int64("stakeIndex", 1, "质押索引（查询UserStake时使用）")
	flag.Parse()
	fmt.Printf("stakeAmount: %d, stakePeriod: %d, stakeIndex: %d\n", *stakeAmountFlag, *stakePeriodFlag, *stakeIndexFlag)

	//校验参数
	if *stakeAmountFlag <= 0 || (*stakePeriodFlag < 0 && *stakePeriodFlag > 3) {
		panic("Invalid input parameters. Please provide valid stakeAmount(>0) and stakePeriod([0-3]).")
	}

	ethClientInfo, err := adapter.NewInitEthClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		ethClientInfo.CloseEthClient()
	}()
	//approve
	transService := service.NewTransactionService(ethClientInfo)
	transService.ApproveMTKStake(big.NewInt(*stakeAmountFlag))

	time.Sleep(15 * time.Second) //等待approve交易完成

	stakeService := service.NewStakeService(ethClientInfo)
	stakeResponse, err := stakeService.Stake(*stakeAmountFlag, uint8(*stakePeriodFlag))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stake response: %s\n", stakeResponse)

	stakeInfo, err := stakeService.UserStakes(ethClientInfo.FromAddress.String(), *stakeIndexFlag)
	if err != nil {
		panic(err)
	}
	fmt.Printf("After staking, UserStake[%d]: %+v\n", *stakeIndexFlag, stakeInfo)

	_withdraw(stakeInfo.StakeId, stakeService)

}

func _withdraw(stakeId *big.Int, stakeService *service.StakeService) {
	withdrawResponse, err := stakeService.Withdraw(stakeId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Withdraw response: %s\n", withdrawResponse)
}
