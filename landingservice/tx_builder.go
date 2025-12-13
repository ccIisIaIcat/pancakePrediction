package landingservice

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ccIisIaIcat/pancakePrediction/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TxBuilder 交易构建器
type TxBuilder struct {
	contractAddress common.Address
	rpcURL          string
	chainID         *big.Int
}

// NewTxBuilder 创建交易构建器
func NewTxBuilder(contractAddress string, rpcURL string, chainID int64) *TxBuilder {
	return &TxBuilder{
		contractAddress: common.HexToAddress(contractAddress),
		rpcURL:          rpcURL,
		chainID:         big.NewInt(chainID),
	}
}

// BetBullParams BetBull 交易参数
type BetBullParams struct {
	FromAddress common.Address // 发送者地址
	Epoch       *big.Int       // 期数
	BetAmount   *big.Int       // 下注金额 (wei)
	GasLimit    uint64         // Gas 限制 (可选,0 则自动估算)
	GasPrice    *big.Int       // Gas 价格 (可选,nil 则自动获取)
	Nonce       *uint64        // Nonce (可选,nil 则自动获取)
}

// BetBearParams BetBear 交易参数
type BetBearParams struct {
	FromAddress common.Address // 发送者地址
	Epoch       *big.Int       // 期数
	BetAmount   *big.Int       // 下注金额 (wei)
	GasLimit    uint64         // Gas 限制 (可选,0 则自动估算)
	GasPrice    *big.Int       // Gas 价格 (可选,nil 则自动获取)
	Nonce       *uint64        // Nonce (可选,nil 则自动获取)
}

// ClaimParams Claim 交易参数
type ClaimParams struct {
	FromAddress common.Address // 发送者地址
	Epochs      []*big.Int     // 要领取的期数列表
	GasLimit    uint64         // Gas 限制 (可选,0 则自动估算)
	GasPrice    *big.Int       // Gas 价格 (可选,nil 则自动获取)
	Nonce       *uint64        // Nonce (可选,nil 则自动获取)
}

// BuildBetBullTx 构建 BetBull 未签名交易
func (b *TxBuilder) BuildBetBullTx(params BetBullParams) (*types.Transaction, error) {
	client, err := ethclient.Dial(b.rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc: %w", err)
	}
	defer client.Close()

	// 获取 nonce
	nonce := uint64(0)
	if params.Nonce != nil {
		nonce = *params.Nonce
	} else {
		nonce, err = client.PendingNonceAt(context.Background(), params.FromAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce: %w", err)
		}
	}

	// 获取 gas price
	gasPrice := params.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get gas price: %w", err)
		}
	}

	// 创建合约实例
	contract, err := contracts.NewPancakePrediction(b.contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	// 构建交易选项
	opts := &bind.TransactOpts{
		From:     params.FromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		GasPrice: gasPrice,
		GasLimit: params.GasLimit,
		Value:    params.BetAmount,
		NoSend:   true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	// 构建交易
	tx, err := contract.BetBull(opts, params.Epoch)
	if err != nil {
		return nil, fmt.Errorf("failed to build bet bull tx: %w", err)
	}

	return tx, nil
}

// BuildBetBearTx 构建 BetBear 未签名交易
func (b *TxBuilder) BuildBetBearTx(params BetBearParams) (*types.Transaction, error) {
	client, err := ethclient.Dial(b.rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc: %w", err)
	}
	defer client.Close()

	// 获取 nonce
	nonce := uint64(0)
	if params.Nonce != nil {
		nonce = *params.Nonce
	} else {
		nonce, err = client.PendingNonceAt(context.Background(), params.FromAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce: %w", err)
		}
	}

	// 获取 gas price
	gasPrice := params.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get gas price: %w", err)
		}
	}

	// 创建合约实例
	contract, err := contracts.NewPancakePrediction(b.contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	// 构建交易选项
	opts := &bind.TransactOpts{
		From:     params.FromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		GasPrice: gasPrice,
		GasLimit: params.GasLimit,
		Value:    params.BetAmount,
		NoSend:   true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	// 构建交易
	tx, err := contract.BetBear(opts, params.Epoch)
	if err != nil {
		return nil, fmt.Errorf("failed to build bet bear tx: %w", err)
	}

	return tx, nil
}

// GetChainID 获取配置的 ChainID
func (b *TxBuilder) GetChainID() *big.Int {
	return b.chainID
}

// BuildClaimTx 构建 Claim 未签名交易
func (b *TxBuilder) BuildClaimTx(params ClaimParams) (*types.Transaction, error) {
	client, err := ethclient.Dial(b.rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial rpc: %w", err)
	}
	defer client.Close()

	// 获取 nonce
	nonce := uint64(0)
	if params.Nonce != nil {
		nonce = *params.Nonce
	} else {
		nonce, err = client.PendingNonceAt(context.Background(), params.FromAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce: %w", err)
		}
	}

	// 获取 gas price
	gasPrice := params.GasPrice
	if gasPrice == nil {
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get gas price: %w", err)
		}
	}

	// 创建合约实例
	contract, err := contracts.NewPancakePrediction(b.contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	// 构建交易选项
	opts := &bind.TransactOpts{
		From:     params.FromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		GasPrice: gasPrice,
		GasLimit: params.GasLimit,
		Value:    big.NewInt(0),
		NoSend:   true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	// 构建交易
	tx, err := contract.Claim(opts, params.Epochs)
	if err != nil {
		return nil, fmt.Errorf("failed to build claim tx: %w", err)
	}

	return tx, nil
}
