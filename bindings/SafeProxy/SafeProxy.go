// This file was generated by seer: https://github.com/G7DAO/seer.
// seer version: 0.3.5
// seer command: seer evm generate --package SafeProxy --cli --struct SafeProxy --output bindings/SafeProxy/SafeProxy.go
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SafeProxy

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	"net/http"
	"strings"

	"context"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	// Reference imports to suppress errors if they are not otherwise used.
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/G7DAO/seer/bindings/CreateCall"
	"github.com/G7DAO/seer/bindings/GnosisSafe"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	// SafeProxyMetaData contains all meta data concerning the SafeProxy contract.
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var SafeProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516101d63803806101d68339818101604052602081101561003357600080fd5b8101908080519060200190929190505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156100ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806101b46022913960400191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050609b806101196000396000f3fe60806040526000547fa619486e00000000000000000000000000000000000000000000000000000000600035141560405780600c1b600c1c60005260206000f35b3660008037600080366000845af43d6000803e60008114156060573d6000fd5b3d6000f3fea2646970667358221220bfbe5e66dfccd59d80684323ec36a561ddc5ef3b39a33a941f25cabefff21eb964736f6c63430007060033496e76616c69642073696e676c65746f6e20616464726573732070726f7669646564",
}

// SafeProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyMetaData.ABI instead.
var SafeProxyABI = SafeProxyMetaData.ABI

// SafeProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeProxyMetaData.Bin instead.
var SafeProxyBin = SafeProxyMetaData.Bin

// DeploySafeProxy deploys a new Ethereum contract, binding an instance of SafeProxy to it.
func DeploySafeProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _singleton common.Address) (common.Address, *types.Transaction, *SafeProxy, error) {
	parsed, err := SafeProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeProxyBin), backend, _singleton)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeProxy{SafeProxyCaller: SafeProxyCaller{contract: contract}, SafeProxyTransactor: SafeProxyTransactor{contract: contract}, SafeProxyFilterer: SafeProxyFilterer{contract: contract}}, nil
}

// SafeProxy is an auto generated Go binding around an Ethereum contract.
type SafeProxy struct {
	SafeProxyCaller     // Read-only binding to the contract
	SafeProxyTransactor // Write-only binding to the contract
	SafeProxyFilterer   // Log filterer for contract events
}

// SafeProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxySession struct {
	Contract     *SafeProxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyCallerSession struct {
	Contract *SafeProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyTransactorSession struct {
	Contract     *SafeProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyRaw struct {
	Contract *SafeProxy // Generic contract binding to access the raw methods on
}

// SafeProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyCallerRaw struct {
	Contract *SafeProxyCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyTransactorRaw struct {
	Contract *SafeProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxy creates a new instance of SafeProxy, bound to a specific deployed contract.
func NewSafeProxy(address common.Address, backend bind.ContractBackend) (*SafeProxy, error) {
	contract, err := bindSafeProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxy{SafeProxyCaller: SafeProxyCaller{contract: contract}, SafeProxyTransactor: SafeProxyTransactor{contract: contract}, SafeProxyFilterer: SafeProxyFilterer{contract: contract}}, nil
}

// NewSafeProxyCaller creates a new read-only instance of SafeProxy, bound to a specific deployed contract.
func NewSafeProxyCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyCaller, error) {
	contract, err := bindSafeProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyCaller{contract: contract}, nil
}

// NewSafeProxyTransactor creates a new write-only instance of SafeProxy, bound to a specific deployed contract.
func NewSafeProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyTransactor, error) {
	contract, err := bindSafeProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyTransactor{contract: contract}, nil
}

// NewSafeProxyFilterer creates a new log filterer instance of SafeProxy, bound to a specific deployed contract.
func NewSafeProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFilterer, error) {
	contract, err := bindSafeProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFilterer{contract: contract}, nil
}

// bindSafeProxy binds a generic wrapper to an already deployed contract.
func bindSafeProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxy *SafeProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxy.Contract.SafeProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxy *SafeProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxy.Contract.SafeProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxy *SafeProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxy.Contract.SafeProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxy *SafeProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxy *SafeProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxy *SafeProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SafeProxy *SafeProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SafeProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SafeProxy *SafeProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeProxy.Contract.Fallback(&_SafeProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SafeProxy *SafeProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeProxy.Contract.Fallback(&_SafeProxy.TransactOpts, calldata)
}

func CreateSafeProxyDeploymentCommand() *cobra.Command {
	var keyfile, nonce, password, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas, rpc string
	var gasLimit uint64
	var simulate bool
	var timeout uint
	var safeAddress, safeApi, safeCreateCall, safeSaltRaw, safeNonceRaw string
	var safeOperationType uint8
	var salt [32]byte
	var predictAddress bool
	var safeNonce *big.Int

	var singleton common.Address
	var singletonRaw string

	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a new SafeProxy contract",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if keyfile == "" {
				return fmt.Errorf("--keystore not specified (this should be a path to an Ethereum account keystore file)")
			}

			if rpc == "" {
				return fmt.Errorf("--rpc not specified (this should be a URL to an Ethereum JSONRPC API)")
			}

			if safeAddress != "" {
				if !common.IsHexAddress(safeAddress) {
					return fmt.Errorf("--safe is not a valid Ethereum address")
				}
				if safeApi == "" {
					client, clientErr := NewClient(rpc)
					if clientErr != nil {
						return clientErr
					}
					chainIDCtx, cancelChainIDCtx := NewChainContext(timeout)
					defer cancelChainIDCtx()
					chainID, chainIDErr := client.ChainID(chainIDCtx)
					if chainIDErr != nil {
						return chainIDErr
					}
					safeApi = fmt.Sprintf("https://safe-client.safe.global/v1/chains/%s/transactions/%s/propose", chainID.String(), safeAddress)
					fmt.Println("--safe-api not specified, using default (", safeApi, ")")
				}

				if safeCreateCall == "" {
					fmt.Println("--safe-create-call not specified, using default (0x7cbB62EaA69F79e6873cD1ecB2392971036cFAa4)")
					safeCreateCall = "0x7cbB62EaA69F79e6873cD1ecB2392971036cFAa4"
				}
				if !common.IsHexAddress(safeCreateCall) {
					return fmt.Errorf("--safe-create-call is not a valid Ethereum address")
				}

				if SafeOperationType(safeOperationType).String() == "Unknown" {
					return fmt.Errorf("--safe-operation must be 0 (Call) or 1 (DelegateCall)")
				}

				if safeSaltRaw == "" {
					fmt.Println("--safe-salt not specified, generating random salt")
					_, err := rand.Read(salt[:])
					if err != nil {
						return fmt.Errorf("failed to generate random salt: %v", err)
					}
					// prompt user to accept random salt
					fmt.Println("Generated salt:", common.Bytes2Hex(salt[:]))
					fmt.Println("Please check the salt and confirm (y/n)")
					var confirm string
					fmt.Scanln(&confirm)
					if confirm != "y" && confirm != "Y" && confirm != "\n" && confirm != "" {
						return fmt.Errorf("salt not accepted, please specify a valid salt")
					}
				} else {
					copy(salt[:], safeSaltRaw)
				}

				if safeNonceRaw == "" {
					fmt.Println("--safe-nonce not specified, fetching nonce from Safe contract")
				} else {
					safeNonce = new(big.Int)
					_, ok := safeNonce.SetString(safeNonceRaw, 0)
					if !ok {
						return fmt.Errorf("--safe-nonce is not a valid big integer")
					}
				}
			}

			if singletonRaw == "" {
				return fmt.Errorf("--singleton argument not specified")
			} else if !common.IsHexAddress(singletonRaw) {
				return fmt.Errorf("--singleton argument is not a valid Ethereum address")
			}
			singleton = common.HexToAddress(singletonRaw)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, clientErr := NewClient(rpc)
			if clientErr != nil {
				return clientErr
			}

			key, keyErr := KeyFromFile(keyfile, password)
			if keyErr != nil {
				return keyErr
			}

			chainIDCtx, cancelChainIDCtx := NewChainContext(timeout)
			defer cancelChainIDCtx()
			chainID, chainIDErr := client.ChainID(chainIDCtx)
			if chainIDErr != nil {
				return chainIDErr
			}

			transactionOpts, transactionOptsErr := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
			if transactionOptsErr != nil {
				return transactionOptsErr
			}

			SetTransactionParametersFromArgs(transactionOpts, nonce, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas, gasLimit, simulate)

			if safeAddress != "" {
				// Generate deploy bytecode with constructor arguments
				deployBytecode, err := generateSafeProxyDeployBytecode(
					singleton,
				)
				if err != nil {
					return fmt.Errorf("failed to generate deploy bytecode: %v", err)
				}

				// Create Safe proposal for deployment
				value := transactionOpts.Value
				if value == nil {
					value = big.NewInt(0)
				}

				if predictAddress {
					fmt.Println("Predicting deployment address...")
					from := common.HexToAddress(safeAddress)
					if safeOperationType == 0 {
						from = common.HexToAddress(safeCreateCall)
					}
					deploymentAddress, err := PredictDeploymentAddressSafe(from, salt, deployBytecode)
					if err != nil {
						return fmt.Errorf("failed to predict deployment address: %v", err)
					}
					fmt.Println("Predicted deployment address:", deploymentAddress.Hex())
					return nil
				} else {
					fmt.Println("Creating Safe proposal...")
					err = DeployWithSafe(client, key, common.HexToAddress(safeAddress), common.HexToAddress(safeCreateCall), value, safeApi, deployBytecode, SafeOperationType(safeOperationType), salt, safeNonce)
					if err != nil {
						return fmt.Errorf("failed to create Safe proposal: %v", err)
					}
				}

				return nil
			}

			address, deploymentTransaction, _, deploymentErr := DeploySafeProxy(
				transactionOpts,
				client,
				singleton,
			)
			if deploymentErr != nil {
				return deploymentErr
			}

			cmd.Printf("Transaction hash: %s\nContract address: %s\n", deploymentTransaction.Hash().Hex(), address.Hex())
			if transactionOpts.NoSend {
				estimationMessage := ethereum.CallMsg{
					From: transactionOpts.From,
					Data: deploymentTransaction.Data(),
				}

				gasEstimationCtx, cancelGasEstimationCtx := NewChainContext(timeout)
				defer cancelGasEstimationCtx()

				gasEstimate, gasEstimateErr := client.EstimateGas(gasEstimationCtx, estimationMessage)
				if gasEstimateErr != nil {
					return gasEstimateErr
				}

				transactionBinary, transactionBinaryErr := deploymentTransaction.MarshalBinary()
				if transactionBinaryErr != nil {
					return transactionBinaryErr
				}
				transactionBinaryHex := hex.EncodeToString(transactionBinary)

				cmd.Printf("Transaction: %s\nEstimated gas: %d\n", transactionBinaryHex, gasEstimate)
			} else {
				cmd.Println("Transaction submitted")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&rpc, "rpc", "", "URL of the JSONRPC API to use")
	cmd.Flags().StringVar(&keyfile, "keyfile", "", "Path to the keystore file to use for the transaction")
	cmd.Flags().StringVar(&password, "password", "", "Password to use to unlock the keystore (if not specified, you will be prompted for the password when the command executes)")
	cmd.Flags().StringVar(&nonce, "nonce", "", "Nonce to use for the transaction")
	cmd.Flags().StringVar(&value, "value", "", "Value to send with the transaction")
	cmd.Flags().StringVar(&gasPrice, "gas-price", "", "Gas price to use for the transaction")
	cmd.Flags().StringVar(&maxFeePerGas, "max-fee-per-gas", "", "Maximum fee per gas to use for the (EIP-1559) transaction")
	cmd.Flags().StringVar(&maxPriorityFeePerGas, "max-priority-fee-per-gas", "", "Maximum priority fee per gas to use for the (EIP-1559) transaction")
	cmd.Flags().Uint64Var(&gasLimit, "gas-limit", 0, "Gas limit for the transaction")
	cmd.Flags().BoolVar(&simulate, "simulate", false, "Simulate the transaction without sending it")
	cmd.Flags().UintVar(&timeout, "timeout", 60, "Timeout (in seconds) for interactions with the JSONRPC API")
	cmd.Flags().StringVar(&safeAddress, "safe", "", "Address of the Safe contract")
	cmd.Flags().StringVar(&safeApi, "safe-api", "", "Safe API for the Safe Transaction Service (optional)")
	cmd.Flags().StringVar(&safeCreateCall, "safe-create-call", "", "Address of the CreateCall contract (optional)")
	cmd.Flags().Uint8Var(&safeOperationType, "safe-operation", 1, "Safe operation type: 0 (Call) or 1 (DelegateCall) - default is 1")
	cmd.Flags().StringVar(&safeSaltRaw, "safe-salt", "", "Salt to use for the Safe transaction")
	cmd.Flags().BoolVar(&predictAddress, "safe-predict-address", false, "Predict the deployment address (only works for Safe transactions)")
	cmd.Flags().StringVar(&safeNonceRaw, "safe-nonce", "", "Safe nonce overrider for the transaction (optional)")

	cmd.Flags().StringVar(&singletonRaw, "singleton", "", "singleton argument (common.Address)")

	return cmd
}

func generateSafeProxyDeployBytecode(
	singleton common.Address,
) ([]byte, error) {
	abiPacked, err := SafeProxyMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %v", err)
	}

	constructorArguments, err := abiPacked.Pack("",
		singleton,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack constructor arguments: %v", err)
	}

	deployBytecode := append(common.FromHex(SafeProxyMetaData.Bin), constructorArguments...)
	return deployBytecode, nil
}

func CreateFallbackCommand() *cobra.Command {
	var keyfile, nonce, password, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas, rpc, contractAddressRaw, safeFunction, safeNonceRaw string
	var gasLimit uint64
	var simulate bool
	var timeout uint
	var contractAddress common.Address
	var safeAddress, safeApi string
	var safeOperationType uint8
	var safeNonce *big.Int

	var calldata []byte
	var calldataRaw string

	cmd := &cobra.Command{
		Use:   "fallback",
		Short: "Execute the Fallback method on a SafeProxy contract",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if contractAddressRaw == "" {
				return fmt.Errorf("--contract not specified")
			} else if !common.IsHexAddress(contractAddressRaw) {
				return fmt.Errorf("--contract is not a valid Ethereum address")
			}
			contractAddress = common.HexToAddress(contractAddressRaw)

			if keyfile == "" {
				return fmt.Errorf("--keystore not specified (this should be a path to an Ethereum account keystore file)")
			}

			if rpc == "" {
				return fmt.Errorf("--rpc not specified (this should be a URL to an Ethereum JSONRPC API)")
			}

			if safeAddress != "" {
				if !common.IsHexAddress(safeAddress) {
					return fmt.Errorf("--safe is not a valid Ethereum address")
				}
				if safeApi == "" {
					client, clientErr := NewClient(rpc)
					if clientErr != nil {
						return clientErr
					}
					chainIDCtx, cancelChainIDCtx := NewChainContext(timeout)
					defer cancelChainIDCtx()
					chainID, chainIDErr := client.ChainID(chainIDCtx)
					if chainIDErr != nil {
						return chainIDErr
					}
					safeApi = fmt.Sprintf("https://safe-client.safe.global/v1/chains/%s/transactions/%s/propose", chainID.String(), safeAddress)
					fmt.Println("--safe-api not specified, using default (", safeApi, ")")
				}

				if SafeOperationType(safeOperationType).String() == "Unknown" {
					return fmt.Errorf("--safe-operation must be 0 (Call) or 1 (DelegateCall)")
				}

				if safeNonceRaw == "" {
					fmt.Println("--safe-nonce not specified, fetching nonce from Safe contract")
				} else {
					safeNonce = new(big.Int)
					_, ok := safeNonce.SetString(safeNonceRaw, 0)
					if !ok {
						return fmt.Errorf("--safe-nonce is not a valid big integer")
					}
				}
			}

			var calldataIntermediate []byte

			var calldataIntermediateHexDecodeErr error
			calldataIntermediate, calldataIntermediateHexDecodeErr = hex.DecodeString(calldataRaw)
			if calldataIntermediateHexDecodeErr != nil {
				return calldataIntermediateHexDecodeErr
			}

			copy(calldata[:], calldataIntermediate)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, clientErr := NewClient(rpc)
			if clientErr != nil {
				return clientErr
			}

			key, keyErr := KeyFromFile(keyfile, password)
			if keyErr != nil {
				return keyErr
			}

			chainIDCtx, cancelChainIDCtx := NewChainContext(timeout)
			defer cancelChainIDCtx()
			chainID, chainIDErr := client.ChainID(chainIDCtx)
			if chainIDErr != nil {
				return chainIDErr
			}

			transactionOpts, transactionOptsErr := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
			if transactionOptsErr != nil {
				return transactionOptsErr
			}

			SetTransactionParametersFromArgs(transactionOpts, nonce, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas, gasLimit, simulate)

			contract, contractErr := NewSafeProxy(contractAddress, client)
			if contractErr != nil {
				return contractErr
			}

			session := SafeProxyTransactorSession{
				Contract:     &contract.SafeProxyTransactor,
				TransactOpts: *transactionOpts,
			}

			if safeAddress != "" {
				abi, err := SafeProxyMetaData.GetAbi()
				if err != nil {
					return fmt.Errorf("failed to get ABI: %v", err)
				}

				// Generate transaction data (override method name if safe function is specified)
				methodName := "fallback"
				if safeFunction != "" {
					methodName = safeFunction
				}

				transaction, err := abi.Pack(
					methodName,
					calldata,
				)

				if err != nil {
					return err
				}

				// Create Safe proposal for transaction
				value := transactionOpts.Value
				if value == nil {
					value = big.NewInt(0)
				}

				err = CreateSafeProposal(client, key, common.HexToAddress(safeAddress), contractAddress, transaction, value, safeApi, SafeOperationType(safeOperationType), safeNonce)
				if err != nil {
					return fmt.Errorf("failed to create Safe proposal: %v", err)
				}

				return nil
			}

			transaction, err := session.Fallback(

				calldata,
			)
			if err != nil {
				return err
			}

			cmd.Printf("Transaction hash: %s\n", transaction.Hash().Hex())
			if transactionOpts.NoSend {
				estimationMessage := ethereum.CallMsg{
					From: transactionOpts.From,
					To:   &contractAddress,
					Data: transaction.Data(),
				}

				gasEstimationCtx, cancelGasEstimationCtx := NewChainContext(timeout)
				defer cancelGasEstimationCtx()

				gasEstimate, gasEstimateErr := client.EstimateGas(gasEstimationCtx, estimationMessage)
				if gasEstimateErr != nil {
					return gasEstimateErr
				}

				transactionBinary, transactionBinaryErr := transaction.MarshalBinary()
				if transactionBinaryErr != nil {
					return transactionBinaryErr
				}
				transactionBinaryHex := hex.EncodeToString(transactionBinary)

				cmd.Printf("Transaction: %s\nEstimated gas: %d\n", transactionBinaryHex, gasEstimate)
			} else {
				cmd.Println("Transaction submitted")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&rpc, "rpc", "", "URL of the JSONRPC API to use")
	cmd.Flags().StringVar(&keyfile, "keyfile", "", "Path to the keystore file to use for the transaction")
	cmd.Flags().StringVar(&password, "password", "", "Password to use to unlock the keystore (if not specified, you will be prompted for the password when the command executes)")
	cmd.Flags().StringVar(&nonce, "nonce", "", "Nonce to use for the transaction")
	cmd.Flags().StringVar(&value, "value", "", "Value to send with the transaction")
	cmd.Flags().StringVar(&gasPrice, "gas-price", "", "Gas price to use for the transaction")
	cmd.Flags().StringVar(&maxFeePerGas, "max-fee-per-gas", "", "Maximum fee per gas to use for the (EIP-1559) transaction")
	cmd.Flags().StringVar(&maxPriorityFeePerGas, "max-priority-fee-per-gas", "", "Maximum priority fee per gas to use for the (EIP-1559) transaction")
	cmd.Flags().Uint64Var(&gasLimit, "gas-limit", 0, "Gas limit for the transaction")
	cmd.Flags().BoolVar(&simulate, "simulate", false, "Simulate the transaction without sending it")
	cmd.Flags().UintVar(&timeout, "timeout", 60, "Timeout (in seconds) for interactions with the JSONRPC API")
	cmd.Flags().StringVar(&contractAddressRaw, "contract", "", "Address of the contract to interact with")
	cmd.Flags().StringVar(&safeAddress, "safe", "", "Address of the Safe contract")
	cmd.Flags().StringVar(&safeApi, "safe-api", "", "Safe API for the Safe Transaction Service (optional)")
	cmd.Flags().Uint8Var(&safeOperationType, "safe-operation", 0, "Safe operation type: 0 (Call) or 1 (DelegateCall)")
	cmd.Flags().StringVar(&safeFunction, "safe-function", "", "Safe function overrider to use for the transaction (optional)")
	cmd.Flags().StringVar(&safeNonceRaw, "safe-nonce", "", "Safe nonce overrider for the transaction (optional)")

	cmd.Flags().StringVar(&calldataRaw, "calldata", "", "calldata argument ([]byte)")

	return cmd
}

var ErrNoRPCURL error = errors.New("no RPC URL provided -- please pass an RPC URL from the command line or set the SAFE_PROXY_RPC_URL environment variable")

// Generates an Ethereum client to the JSONRPC API at the given URL. If rpcURL is empty, then it
// attempts to read the RPC URL from the SAFE_PROXY_RPC_URL environment variable. If that is empty,
// too, then it returns an error.
func NewClient(rpcURL string) (*ethclient.Client, error) {
	if rpcURL == "" {
		rpcURL = os.Getenv("SAFE_PROXY_RPC_URL")
	}

	if rpcURL == "" {
		return nil, ErrNoRPCURL
	}

	client, err := ethclient.Dial(rpcURL)
	return client, err
}

// Creates a new context to be used when interacting with the chain client.
func NewChainContext(timeout uint) (context.Context, context.CancelFunc) {
	baseCtx := context.Background()
	parsedTimeout := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(baseCtx, parsedTimeout)
	return ctx, cancel
}

// Unlocks a key from a keystore (byte contents of a keystore file) with the given password.
func UnlockKeystore(keystoreData []byte, password string) (*keystore.Key, error) {
	key, err := keystore.DecryptKey(keystoreData, password)
	return key, err
}

// Loads a key from file, prompting the user for the password if it is not provided as a function argument.
func KeyFromFile(keystoreFile string, password string) (*keystore.Key, error) {
	var emptyKey *keystore.Key
	keystoreContent, readErr := os.ReadFile(keystoreFile)
	if readErr != nil {
		return emptyKey, readErr
	}

	// If password is "", prompt user for password.
	if password == "" {
		fmt.Printf("Please provide a password for keystore (%s): ", keystoreFile)
		passwordRaw, inputErr := term.ReadPassword(int(os.Stdin.Fd()))
		if inputErr != nil {
			return emptyKey, fmt.Errorf("error reading password: %s", inputErr.Error())
		}
		fmt.Print("\n")
		password = string(passwordRaw)
	}

	key, err := UnlockKeystore(keystoreContent, password)
	return key, err
}

// This method is used to set the parameters on a view call from command line arguments (represented mostly as
// strings).
func SetCallParametersFromArgs(opts *bind.CallOpts, pending bool, fromAddress, blockNumber string) {
	if pending {
		opts.Pending = true
	}

	if fromAddress != "" {
		opts.From = common.HexToAddress(fromAddress)
	}

	if blockNumber != "" {
		opts.BlockNumber = new(big.Int)
		opts.BlockNumber.SetString(blockNumber, 0)
	}
}

// This method is used to set the parameters on a transaction from command line arguments (represented mostly as
// strings).
func SetTransactionParametersFromArgs(opts *bind.TransactOpts, nonce, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas string, gasLimit uint64, noSend bool) {
	if nonce != "" {
		opts.Nonce = new(big.Int)
		opts.Nonce.SetString(nonce, 0)
	}

	if value != "" {
		opts.Value = new(big.Int)
		opts.Value.SetString(value, 0)
	}

	if gasPrice != "" {
		opts.GasPrice = new(big.Int)
		opts.GasPrice.SetString(gasPrice, 0)
	}

	if maxFeePerGas != "" {
		opts.GasFeeCap = new(big.Int)
		opts.GasFeeCap.SetString(maxFeePerGas, 0)
	}

	if maxPriorityFeePerGas != "" {
		opts.GasTipCap = new(big.Int)
		opts.GasTipCap.SetString(maxPriorityFeePerGas, 0)
	}

	if gasLimit != 0 {
		opts.GasLimit = gasLimit
	}

	opts.NoSend = noSend
}

func CreateSafeProxyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "safe-proxy",
		Short: "Interact with the SafeProxy contract",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.SetOut(os.Stdout)

	DeployGroup := &cobra.Group{
		ID: "deploy", Title: "Commands which deploy contracts",
	}
	cmd.AddGroup(DeployGroup)
	ViewGroup := &cobra.Group{
		ID: "view", Title: "Commands which view contract state",
	}
	TransactGroup := &cobra.Group{
		ID: "transact", Title: "Commands which submit transactions",
	}
	cmd.AddGroup(ViewGroup, TransactGroup)

	cmdDeploySafeProxy := CreateSafeProxyDeploymentCommand()
	cmdDeploySafeProxy.GroupID = DeployGroup.ID
	cmd.AddCommand(cmdDeploySafeProxy)

	cmdTransactFallback := CreateFallbackCommand()
	cmdTransactFallback.GroupID = TransactGroup.ID
	cmd.AddCommand(cmdTransactFallback)

	return cmd
}

// SafeOperationType represents the type of operation for a Safe transaction
type SafeOperationType uint8

const (
	Call         SafeOperationType = 0
	DelegateCall SafeOperationType = 1
)

// String returns the string representation of the SafeOperationType
func (o SafeOperationType) String() string {
	switch o {
	case Call:
		return "Call"
	case DelegateCall:
		return "DelegateCall"
	default:
		return "Unknown"
	}
}

// SafeTransactionData represents the data for a Safe transaction
type SafeTransactionData struct {
	To             string            `json:"to"`
	Value          string            `json:"value"`
	Data           string            `json:"data"`
	Operation      SafeOperationType `json:"operation"`
	SafeTxGas      uint64            `json:"safeTxGas"`
	BaseGas        uint64            `json:"baseGas"`
	GasPrice       string            `json:"gasPrice"`
	GasToken       string            `json:"gasToken"`
	RefundReceiver string            `json:"refundReceiver"`
	Nonce          *big.Int          `json:"nonce"`
	SafeTxHash     string            `json:"safeTxHash"`
	Sender         string            `json:"sender"`
	Signature      string            `json:"signature"`
	Origin         string            `json:"origin"`
}

const (
	NativeTokenAddress = "0x0000000000000000000000000000000000000000"
)

func DeployWithSafe(client *ethclient.Client, key *keystore.Key, safeAddress common.Address, factoryAddress common.Address, value *big.Int, safeApi string, deployBytecode []byte, safeOperationType SafeOperationType, salt [32]byte, safeNonce *big.Int) error {
	abi, err := CreateCall.CreateCallMetaData.GetAbi()
	if err != nil {
		return fmt.Errorf("failed to get ABI: %v", err)
	}

	safeCreateCallTxData, err := abi.Pack("performCreate2", value, deployBytecode, salt)
	if err != nil {
		return fmt.Errorf("failed to pack performCreate2 transaction: %v", err)
	}

	return CreateSafeProposal(client, key, safeAddress, factoryAddress, safeCreateCallTxData, value, safeApi, SafeOperationType(safeOperationType), safeNonce)
}

func PredictDeploymentAddressSafe(from common.Address, salt [32]byte, deployBytecode []byte) (common.Address, error) {
	// Calculate the hash of the init code (deployment bytecode)
	initCodeHash := crypto.Keccak256(deployBytecode)

	// Calculate the CREATE2 address
	deployedAddress := crypto.CreateAddress2(from, salt, initCodeHash)

	return deployedAddress, nil
}

func CreateSafeProposal(client *ethclient.Client, key *keystore.Key, safeAddress common.Address, to common.Address, data []byte, value *big.Int, safeApi string, safeOperationType SafeOperationType, safeNonce *big.Int) error {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %v", err)
	}

	// Create a new instance of the GnosisSafe contract
	safeInstance, err := GnosisSafe.NewGnosisSafe(safeAddress, client)
	if err != nil {
		return fmt.Errorf("failed to create GnosisSafe instance: %v", err)
	}

	nonce := safeNonce
	if safeNonce == nil {
		// Fetch the current nonce from the Safe contract
		fetchedNonce, err := safeInstance.Nonce(&bind.CallOpts{})
		if err != nil {
			return fmt.Errorf("failed to fetch nonce from Safe contract: %v", err)
		}
		nonce = fetchedNonce
	} else {
		nonce = safeNonce
	}

	safeTransactionData := SafeTransactionData{
		To:             to.Hex(),
		Value:          value.String(),
		Data:           common.Bytes2Hex(data),
		Operation:      safeOperationType,
		SafeTxGas:      0,
		BaseGas:        0,
		GasPrice:       "0",
		GasToken:       NativeTokenAddress,
		RefundReceiver: NativeTokenAddress,
		Nonce:          nonce,
	}

	// Calculate SafeTxHash
	safeTxHash, err := CalculateSafeTxHash(safeAddress, safeTransactionData, chainID)
	if err != nil {
		return fmt.Errorf("failed to calculate SafeTxHash: %v", err)
	}

	// Sign the SafeTxHash
	signature, err := crypto.Sign(safeTxHash.Bytes(), key.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to sign SafeTxHash: %v", err)
	}

	// Adjust V value for Ethereum's replay protection
	signature[64] += 27

	// Convert signature to hex
	senderSignature := "0x" + common.Bytes2Hex(signature)

	// Prepare the request body
	requestBody := map[string]interface{}{
		"to":             safeTransactionData.To,
		"value":          safeTransactionData.Value,
		"data":           "0x" + safeTransactionData.Data,
		"operation":      int(safeTransactionData.Operation),
		"safeTxGas":      fmt.Sprintf("%d", safeTransactionData.SafeTxGas),
		"baseGas":        fmt.Sprintf("%d", safeTransactionData.BaseGas),
		"gasPrice":       safeTransactionData.GasPrice,
		"gasToken":       safeTransactionData.GasToken,
		"refundReceiver": safeTransactionData.RefundReceiver,
		"nonce":          fmt.Sprintf("%d", safeTransactionData.Nonce),
		"safeTxHash":     safeTxHash.Hex(),
		"sender":         key.Address.Hex(),
		"signature":      senderSignature,
		"origin":         fmt.Sprintf("{\"url\":\"%s\",\"name\":\"TokenSender Deployment\"}", safeApi),
	}

	// Marshal the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Send the request to the Safe Transaction Service
	req, err := http.NewRequest("POST", safeApi, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Safe proposal created successfully")
	return nil
}

func CalculateSafeTxHash(safeAddress common.Address, txData SafeTransactionData, chainID *big.Int) (common.Hash, error) {
	domainSeparator := apitypes.TypedDataDomain{
		ChainId:           (*math.HexOrDecimal256)(chainID),
		VerifyingContract: safeAddress.Hex(),
	}

	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"SafeTx": []apitypes.Type{
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "operation", Type: "uint8"},
				{Name: "safeTxGas", Type: "uint256"},
				{Name: "baseGas", Type: "uint256"},
				{Name: "gasPrice", Type: "uint256"},
				{Name: "gasToken", Type: "address"},
				{Name: "refundReceiver", Type: "address"},
				{Name: "nonce", Type: "uint256"},
			},
		},
		Domain:      domainSeparator,
		PrimaryType: "SafeTx",
		Message: apitypes.TypedDataMessage{
			"to":             txData.To,
			"value":          txData.Value,
			"data":           "0x" + txData.Data,
			"operation":      fmt.Sprintf("%d", txData.Operation),
			"safeTxGas":      fmt.Sprintf("%d", txData.SafeTxGas),
			"baseGas":        fmt.Sprintf("%d", txData.BaseGas),
			"gasPrice":       txData.GasPrice,
			"gasToken":       txData.GasToken,
			"refundReceiver": txData.RefundReceiver,
			"nonce":          fmt.Sprintf("%d", txData.Nonce),
		},
	}

	typedDataHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to hash typed data: %v", err)
	}

	return common.BytesToHash(typedDataHash), nil
}
