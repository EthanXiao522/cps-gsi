package dto

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type StakeResponse struct {
	Hash            string         `json:"hash"`
	ContractAddress string         `json:"contractAddress"`
	FromAddress     common.Address `json:"fromAddress"`
	Method          string         `json:"method"`
}
type StakeRequest struct {
	Amount int64 `json:"amount"`
	Period uint8 `json:"period"`
}

type WithDrawnRequest struct {
	Index big.Int `json:"index"`
}

type StakeEventListener struct {
	StakedEventId    common.Hash
	WithdrawnEventId common.Hash
	ContractAddress  common.Address
	ContractABI      abi.ABI
}

type StakeInfoResponse struct {
	StakeId        *big.Int `json:"stakeId"`
	Amount         *big.Int `json:"amount"`
	StartTime      *big.Int `json:"startTime"`
	EndTime        *big.Int `json:"endTime"`
	Period         uint8    `json:"period"`
	RewardRateYear *big.Int `json:"rewardRateYear"`
	IsActive       bool     `json:"isActive"`
}
