package service

import (
	"fmt"
	"math/big"
	"staking-interaction/adapter"
	"staking-interaction/common/config"
	"staking-interaction/contracts/stake"
	"staking-interaction/dto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var cfg = config.Get()
var contractAddr = cfg.BlockchainConfig.Contracts.Stake

type StakeService struct {
	ethClientInfo *adapter.InitClient
}

func NewStakeService(ethClientInfo *adapter.InitClient) *StakeService {
	return &StakeService{
		ethClientInfo: ethClientInfo,
	}
}

func (s *StakeService) NewStakeContract() (*stake.MTKStake, error) {
	return stake.NewMTKStake(common.HexToAddress(contractAddr), s.ethClientInfo.Client)
}

func (s *StakeService) Stake(amount int64, period uint8) (response *dto.StakeResponse, err error) {
	contract, err := s.NewStakeContract()
	if err != nil {
		return nil, fmt.Errorf("failed to NewStakeContract:%v", err)
	}

	trans, err := contract.Stake(s.ethClientInfo.Auth, big.NewInt(amount), period)

	if trans == nil || err != nil {
		return nil, fmt.Errorf("failed to Stake:%v", err)
	}

	return &dto.StakeResponse{
		Hash:            trans.Hash().String(),
		ContractAddress: contractAddr,
		FromAddress:     s.ethClientInfo.FromAddress,
		Method:          "Stake",
	}, nil
}

func (s *StakeService) Withdraw(stakeId *big.Int) (response *dto.StakeResponse, err error) {
	contract, err := s.NewStakeContract()
	if err != nil {
		return nil, fmt.Errorf("failed to NewStakeContract:%v", err)
	}

	trans, err := contract.Withdraw(s.ethClientInfo.Auth, stakeId)

	if trans == nil || err != nil {
		return nil, fmt.Errorf("failed to Withdraw:%v", err)
	}

	return &dto.StakeResponse{
		Hash:            trans.Hash().String(),
		ContractAddress: contractAddr,
		FromAddress:     s.ethClientInfo.FromAddress,
		Method:          "Withdraw",
	}, nil
}

func (s *StakeService) UserStakes(address string, index int64) (response *dto.StakeInfoResponse, err error) {
	contract, err := s.NewStakeContract()
	if err != nil {
		return nil, fmt.Errorf("failed to NewStakeContract:%v", err)
	}

	stakeInfo, err := contract.UserStakes(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(index))
	if err != nil {
		return nil, fmt.Errorf("failed to UserStakes:%v", err)
	}

	return &dto.StakeInfoResponse{
		StakeId:        stakeInfo.StakeId,
		Amount:         stakeInfo.Amount,
		StartTime:      stakeInfo.StartTime,
		EndTime:        stakeInfo.EndTime,
		Period:         stakeInfo.Period,
		RewardRateYear: stakeInfo.RewardRateYear,
		IsActive:       stakeInfo.IsActive,
	}, nil
}
