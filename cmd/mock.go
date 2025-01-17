package cmd

import (
	"crypto/ecdsa"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

type UtilsMock struct{}

type TokenManagerMock struct{}

type TransactionMock struct{}

type AssetManagerMock struct{}

type StakeManagerMock struct{}

type AccountMock struct{}

type KeystoreMock struct{}

type FlagSetMock struct{}

type ProposeUtilsMock struct{}

type UtilsCmdMock struct{}

type VoteManagerMock struct{}

type BlockManagerMock struct{}

type CryptoMock struct{}

var GetOptionsMock func(bool, string, string) bind.CallOpts

var GetTxnOptsMock func(types.TransactionOptions) *bind.TransactOpts

var WaitForBlockCompletionMock func(*ethclient.Client, string) int

var WaitForAppropriateStateMock func(*ethclient.Client, string, string, ...int) (uint32, error)

var WaitIfCommitStateMock func(*ethclient.Client, string, string) (uint32, error)

var GetDefaultPathMock func() (string, error)

var AssignPasswordMock func(*pflag.FlagSet) string

var FetchBalanceMock func(*ethclient.Client, string) (*big.Int, error)

var AssignAmountInWeiMock func(flagSet *pflag.FlagSet) *big.Int

var CheckAmountAndBalanceMock func(amountInWei *big.Int, balance *big.Int) *big.Int

var GetAmountInDecimalMock func(*big.Int) *big.Float

var ConnectToClientMock func(string) *ethclient.Client

var GetStakerIdMock func(*ethclient.Client, string) (uint32, error)

var GetStakerMock func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)

var GetUpdatedStakerMock func(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)

var GetConfigDataMock func() (types.Configurations, error)

var ParseBoolMock func(string) (bool, error)

var GetDelayedStateMock func(*ethclient.Client, int32) (int64, error)

var GetEpochMock func(*ethclient.Client) (uint32, error)

var GetNumberOfStakersMock func(*ethclient.Client, string) (uint32, error)

var GetRandaoHashMock func(*ethclient.Client, string) ([32]byte, error)

var GetNumberOfProposedBlocksMock func(*ethclient.Client, string, uint32) (uint8, error)

var GetMaxAltBlocksMock func(*ethclient.Client, string) (uint8, error)

var GetProposedBlockMock func(*ethclient.Client, string, uint32, uint8) (bindings.StructsBlock, error)

var GetCommitmentsMock func(*ethclient.Client, string) ([32]byte, error)

var AllZeroMock func([32]byte) bool

var GetEpochLastCommittedMock func(*ethclient.Client, string, uint32) (uint32, error)

var GetActiveAssetsDataMock func(*ethclient.Client, string, uint32) ([]*big.Int, error)

var ConvertUintArrayToUint8ArrayMock func([]uint) []uint8

var WaitForDisputeOrConfirmStateMock func(*ethclient.Client, string, string) (uint32, error)

var PrivateKeyPromptMock func() string

var PasswordPromptMock func() string

var GetEpochLastRevealedMock func(*ethclient.Client, string, uint32) (uint32, error)

var GetVoteValueMock func(*ethclient.Client, string, uint8, uint32) (*big.Int, error)

var GetInfluenceSnapshotMock func(*ethclient.Client, string, uint32, uint32) (*big.Int, error)

var GetNumActiveAssetsMock func(*ethclient.Client, string) (*big.Int, error)

var GetTotalInfluenceRevealedMock func(*ethclient.Client, string, uint32) (*big.Int, error)

var ConvertBigIntArrayToUint32ArrayMock func([]*big.Int) []uint32

var GetLockMock func(*ethclient.Client, string, uint32) (types.Locks, error)

var GetWithdrawReleasePeriodMock func(*ethclient.Client, string) (uint8, error)

var GetConfigFilePathMock func() (string, error)

var ViperWriteConfigAsMock func(string) error

var IsEqualMock func(arr1 []uint32, arr2 []uint32) (bool, int)

var GetActiveAssetIdsMock func(*ethclient.Client, string, uint32) ([]uint8, error)

var GetBlockManagerMock func(*ethclient.Client) *bindings.BlockManager

var GetVotesMock func(*ethclient.Client, string, uint32) (bindings.StructsVote, error)

var ContainsMock func([]int, int) bool

var GetLatestBlockMock func(*ethclient.Client) (*Types.Header, error)

var GetSortedProposedBlockIdsMock func(*ethclient.Client, string, uint32) ([]uint8, error)

var GetUpdatedEpochMock func(*ethclient.Client) (uint32, error)

var AllowanceMock func(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)

var ApproveMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var TransferMock func(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)

var HashMock func(*Types.Transaction) common.Hash

var StakeMock func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)

var ExtendLockMock func(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)

var DelegateMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, *big.Int) (*Types.Transaction, error)

var WithdrawContractMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32) (*Types.Transaction, error)

var SetDelegationAcceptanceMock func(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error)

var SetCommissionContractMock func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)

var DecreaseCommissionContractMock func(*ethclient.Client, *bind.TransactOpts, uint8) (*Types.Transaction, error)

var StakerInfoMock func(*ethclient.Client, *bind.CallOpts, uint32) (types.Staker, error)

var GetMaturityMock func(*ethclient.Client, *bind.CallOpts, uint32) (uint16, error)

var CreateAccountMock func(string, string) accounts.Account

var CreateJobMock func(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) (*Types.Transaction, error)

var UpdateJobMock func(*ethclient.Client, *bind.TransactOpts, uint8, uint8, int8, uint8, string, string) (*Types.Transaction, error)

var UpdateCollectionMock func(*ethclient.Client, *bind.TransactOpts, uint8, uint32, int8, []uint8) (*Types.Transaction, error)

var CreateCollectionMock func(*ethclient.Client, *bind.TransactOpts, []uint8, uint32, int8, string) (*Types.Transaction, error)

var AccountsMock func(string) []accounts.Account

var ImportECDSAMock func(string, *ecdsa.PrivateKey, string) (accounts.Account, error)

var GetStringFromMock func(*pflag.FlagSet) (string, error)

var GetStringToMock func(*pflag.FlagSet) (string, error)

var SetCollectionStatusMock func(*ethclient.Client, *bind.TransactOpts, bool, uint8) (*Types.Transaction, error)

var GetActiveStatusMock func(*ethclient.Client, *bind.CallOpts, uint8) (bool, error)

var GetStringAddressMock func(*pflag.FlagSet) (string, error)

var GetUint32StakerIdMock func(*pflag.FlagSet) (uint32, error)

var GetStringNameMock func(*pflag.FlagSet) (string, error)

var GetStringUrlMock func(*pflag.FlagSet) (string, error)

var GetStringSelectorMock func(*pflag.FlagSet) (string, error)

var GetInt8PowerMock func(*pflag.FlagSet) (int8, error)

var getBiggestInfluenceAndIdMock func(*ethclient.Client, string, uint32, UtilsStruct) (*big.Int, uint32, error)

var getIterationMock func(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) int

var isElectedProposerMock func(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) bool

var pseudoRandomNumberGeneratorMock func([]byte, uint32, []byte) *big.Int

var MakeBlockMock func(*ethclient.Client, string, bool, UtilsStruct) ([]uint32, error)

var getSortedVotesMock func(*ethclient.Client, string, uint8, uint32, UtilsStruct) ([]*big.Int, error)

var influencedMedianMock func([]*big.Int, *big.Int) *big.Int

var ProposeMock func(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error)

var GetUint8WeightMock func(*pflag.FlagSet) (uint8, error)

var GetUint8AssetIdMock func(*pflag.FlagSet) (uint8, error)

var CheckCurrentStatusMock func(*ethclient.Client, string, uint8, UtilsStruct) (bool, error)

var DisputeMock func(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int, UtilsStruct) error

var GiveSortedMock func(*ethclient.Client, *bindings.BlockManager, *bind.TransactOpts, uint32, uint8, []uint32)

var GetStringProviderMock func(*pflag.FlagSet) (string, error)

var GetFloat32GasMultiplierMock func(set *pflag.FlagSet) (float32, error)

var GetInt32BufferMock func(*pflag.FlagSet) (int32, error)

var GetInt32WaitMock func(*pflag.FlagSet) (int32, error)

var GetInt32GasPriceMock func(*pflag.FlagSet) (int32, error)

var GetFloat32GasLimitMock func(*pflag.FlagSet) (float32, error)

var GetStringLogLevelMock func(*pflag.FlagSet) (string, error)

var GetStringStatusMock func(*pflag.FlagSet) (string, error)

var GetUint8CommissionMock func(*pflag.FlagSet) (uint8, error)

var SetCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error

var DecreaseCommissionMock func(*ethclient.Client, uint32, *bind.TransactOpts, uint8, UtilsStruct) error

var DecreaseCommissionPromptMock func() bool

var RevealMock func(*ethclient.Client, *bind.TransactOpts, uint32, []*big.Int, [32]byte) (*Types.Transaction, error)

var CommitMock func(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error)

var ClaimBlockRewardMock func(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error)

var FinalizeDisputeMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint8) (*Types.Transaction, error)

var DisputeBiggestInfluenceProposedMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) (*Types.Transaction, error)

var GetUintSliceJobIdsMock func(*pflag.FlagSet) ([]uint, error)

var GetUint32AggregationMock func(*pflag.FlagSet) (uint32, error)

var GetUint8JobIdMock func(*pflag.FlagSet) (uint8, error)

var GetUint8CollectionIdMock func(*pflag.FlagSet) (uint8, error)

var HexToECDSAMock func(string) (*ecdsa.PrivateKey, error)

var WithdrawMock func(*ethclient.Client, *bind.TransactOpts, uint32, uint32, UtilsStruct) (common.Hash, error)

func (u UtilsMock) GetOptions(pending bool, from string, blockNumber string) bind.CallOpts {
	return GetOptionsMock(pending, from, blockNumber)
}

func (u UtilsMock) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return GetTxnOptsMock(transactionData)
}

func (u UtilsMock) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) int {
	return WaitForBlockCompletionMock(client, hashToRead)
}

func (u UtilsMock) WaitForAppropriateState(client *ethclient.Client, accountAddress string, action string, states ...int) (uint32, error) {
	return WaitForAppropriateStateMock(client, accountAddress, action, states...)
}

func (u UtilsMock) WaitIfCommitState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitIfCommitStateMock(client, accountAddress, action)
}

func (u UtilsMock) AssignPassword(flagSet *pflag.FlagSet) string {
	return AssignPasswordMock(flagSet)
}

func (u UtilsMock) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return FetchBalanceMock(client, accountAddress)
}

func (u UtilsMock) AssignAmountInWei(flagSet *pflag.FlagSet) *big.Int {
	return AssignAmountInWeiMock(flagSet)
}

func (u UtilsMock) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return CheckAmountAndBalanceMock(amountInWei, balance)
}

func (u UtilsMock) GetDefaultPath() (string, error) {
	return GetDefaultPathMock()
}

func (u UtilsMock) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return GetAmountInDecimalMock(amountInWei)
}

func (u UtilsMock) ConnectToClient(provider string) *ethclient.Client {
	return ConnectToClientMock(provider)
}

func (u UtilsMock) PasswordPrompt() string {
	return PasswordPromptMock()
}

func (u UtilsMock) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return GetStakerIdMock(client, address)
}

func (u UtilsMock) GetStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return GetStakerMock(client, address, stakerId)
}

func (u UtilsMock) GetUpdatedStaker(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsStaker, error) {
	return GetUpdatedStakerMock(client, address, stakerId)
}

func (u UtilsMock) GetConfigData() (types.Configurations, error) {
	return GetConfigDataMock()
}

func (u UtilsMock) ParseBool(str string) (bool, error) {
	return ParseBoolMock(str)
}

func (u UtilsMock) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	return GetDelayedStateMock(client, buffer)
}

func (u UtilsMock) GetEpoch(client *ethclient.Client) (uint32, error) {
	return GetEpochMock(client)
}

func (u UtilsMock) GetNumberOfStakers(client *ethclient.Client, address string) (uint32, error) {
	return GetNumberOfStakersMock(client, address)
}

func (u UtilsMock) GetRandaoHash(client *ethclient.Client, address string) ([32]byte, error) {
	return GetRandaoHashMock(client, address)
}

func (u UtilsMock) GetNumberOfProposedBlocks(client *ethclient.Client, address string, epoch uint32) (uint8, error) {
	return GetNumberOfProposedBlocksMock(client, address, epoch)
}

func (u UtilsMock) GetMaxAltBlocks(client *ethclient.Client, address string) (uint8, error) {
	return GetMaxAltBlocksMock(client, address)
}

func (u UtilsMock) GetProposedBlock(client *ethclient.Client, address string, epoch uint32, proposedBlockId uint8) (bindings.StructsBlock, error) {
	return GetProposedBlockMock(client, address, epoch, proposedBlockId)
}

func (u UtilsMock) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return GetCommitmentsMock(client, address)
}

func (u UtilsMock) AllZero(bytesValue [32]byte) bool {
	return AllZeroMock(bytesValue)
}

func (u UtilsMock) GetEpochLastCommitted(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return GetEpochLastCommittedMock(client, address, stakerId)
}

func (u UtilsMock) GetActiveAssetsData(client *ethclient.Client, address string, epoch uint32) ([]*big.Int, error) {
	return GetActiveAssetsDataMock(client, address, epoch)
}

func (u UtilsMock) ConvertUintArrayToUint8Array(uintArr []uint) []uint8 {
	return ConvertUintArrayToUint8ArrayMock(uintArr)
}

func (u UtilsMock) WaitForConfirmState(client *ethclient.Client, accountAddress string, action string) (uint32, error) {
	return WaitForDisputeOrConfirmStateMock(client, accountAddress, action)
}

func (u UtilsMock) PrivateKeyPrompt() string {
	return PrivateKeyPromptMock()
}

func (u UtilsMock) GetConfigFilePath() (string, error) {
	return GetConfigFilePathMock()
}

func (u UtilsMock) ViperWriteConfigAs(path string) error {
	return ViperWriteConfigAsMock(path)
}

func (u UtilsMock) GetEpochLastRevealed(client *ethclient.Client, address string, stakerId uint32) (uint32, error) {
	return GetEpochLastRevealedMock(client, address, stakerId)
}

func (u UtilsMock) GetVoteValue(client *ethclient.Client, address string, assetId uint8, stakerId uint32) (*big.Int, error) {
	return GetVoteValueMock(client, address, assetId, stakerId)
}

func (u UtilsMock) GetInfluenceSnapshot(client *ethclient.Client, address string, stakerId uint32, epoch uint32) (*big.Int, error) {
	return GetInfluenceSnapshotMock(client, address, stakerId, epoch)
}

func (u UtilsMock) GetNumActiveAssets(client *ethclient.Client, address string) (*big.Int, error) {
	return GetNumActiveAssetsMock(client, address)
}

func (u UtilsMock) GetTotalInfluenceRevealed(client *ethclient.Client, address string, epoch uint32) (*big.Int, error) {
	return GetTotalInfluenceRevealedMock(client, address, epoch)
}

func (u UtilsMock) ConvertBigIntArrayToUint32Array(bigIntArray []*big.Int) []uint32 {
	return ConvertBigIntArrayToUint32ArrayMock(bigIntArray)
}

func (u UtilsMock) GetLock(client *ethclient.Client, address string, stakerId uint32) (types.Locks, error) {
	return GetLockMock(client, address, stakerId)
}

func (u UtilsMock) GetWithdrawReleasePeriod(client *ethclient.Client, address string) (uint8, error) {
	return GetWithdrawReleasePeriodMock(client, address)
}

func (u UtilsMock) IsEqual(arr1 []uint32, arr2 []uint32) (bool, int) {
	return IsEqualMock(arr1, arr2)
}

func (u UtilsMock) GetActiveAssetIds(client *ethclient.Client, address string, epoch uint32) ([]uint8, error) {
	return GetActiveAssetIdsMock(client, address, epoch)
}

func (u UtilsMock) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	return GetBlockManagerMock(client)
}

func (u UtilsMock) GetVotes(client *ethclient.Client, address string, stakerId uint32) (bindings.StructsVote, error) {
	return GetVotesMock(client, address, stakerId)
}

func (u UtilsMock) Contains(arr []int, val int) bool {
	return ContainsMock(arr, val)
}

func (u UtilsMock) GetLatestBlock(client *ethclient.Client) (*Types.Header, error) {
	return GetLatestBlockMock(client)
}

func (u UtilsMock) GetSortedProposedBlockIds(client *ethclient.Client, address string, epoch uint32) ([]uint8, error) {
	return GetSortedProposedBlockIdsMock(client, address, epoch)
}

func (u UtilsMock) GetUpdatedEpoch(client *ethclient.Client) (uint32, error) {
	return GetUpdatedEpochMock(client)
}

func (tokenManagerMock TokenManagerMock) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	return AllowanceMock(client, opts, owner, spender)
}

func (tokenManagerMock TokenManagerMock) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	return ApproveMock(client, opts, spender, amount)
}

func (tokenManagerMock TokenManagerMock) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	return TransferMock(client, opts, recipient, amount)
}

func (transactionMock TransactionMock) Hash(txn *Types.Transaction) common.Hash {
	return HashMock(txn)
}

func (assetManagerMock AssetManagerMock) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	return CreateJobMock(client, opts, weight, power, selectorType, name, selector, url)
}

func (assetManagerMock AssetManagerMock) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint8, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error) {
	return CreateCollectionMock(client, opts, jobIDs, aggregationMethod, power, name)
}

func (assetManagerMock AssetManagerMock) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint8, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	return UpdateJobMock(client, opts, jobId, weight, power, selectorType, selector, url)
}

func (assetManagerMock AssetManagerMock) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint8, aggregationMethod uint32, power int8, jobIds []uint8) (*Types.Transaction, error) {
	return UpdateCollectionMock(client, opts, collectionId, aggregationMethod, power, jobIds)
}

func (assetManagerMock AssetManagerMock) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint8) (*Types.Transaction, error) {
	return SetCollectionStatusMock(client, opts, assetStatus, id)
}

func (assetManagerMock AssetManagerMock) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint8) (bool, error) {
	return GetActiveStatusMock(client, opts, id)
}

func (stakeManagerMock StakeManagerMock) Stake(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	return StakeMock(client, opts, epoch, amount)
}

func (stakeManagerMock StakeManagerMock) ExtendLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	return ExtendLockMock(client, opts, stakerId)
}

func (stakeManagerMock StakeManagerMock) Delegate(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	return DelegateMock(client, opts, epoch, stakerId, amount)
}

func (stakeManagerMock StakeManagerMock) Withdraw(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, stakerId uint32) (*Types.Transaction, error) {
	return WithdrawContractMock(client, opts, epoch, stakerId)
}

func (stakeManagerMock StakeManagerMock) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	return SetDelegationAcceptanceMock(client, opts, status)
}

func (stakeManagerMock StakeManagerMock) SetCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	return SetCommissionContractMock(client, opts, commission)
}

func (stakeManagerMock StakeManagerMock) DecreaseCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	return DecreaseCommissionContractMock(client, opts, commission)
}

func (stakeManagerMock StakeManagerMock) StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (types.Staker, error) {
	return StakerInfoMock(client, opts, stakerId)
}

func (stakeManagerMock StakeManagerMock) GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error) {
	return GetMaturityMock(client, opts, age)
}

func (account AccountMock) CreateAccount(path string, password string) accounts.Account {
	return CreateAccountMock(path, password)
}

func (ks KeystoreMock) Accounts(path string) []accounts.Account {
	return AccountsMock(path)
}

func (ks KeystoreMock) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (accounts.Account, error) {
	return ImportECDSAMock(path, priv, passphrase)
}

func (flagSetMock FlagSetMock) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	return GetStringFromMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	return GetStringToMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return GetStringNameMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint32StakerId(flagset *pflag.FlagSet) (uint32, error) {
	return GetUint32StakerIdMock(flagset)
}

func (flagSetMock FlagSetMock) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return GetStringAddressMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return GetStringUrlMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return GetStringSelectorMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return GetInt8PowerMock(flagSet)
}

func (proposeUtilsMock ProposeUtilsMock) getBiggestInfluenceAndId(client *ethclient.Client, address string, epoch uint32, utilsStruct UtilsStruct) (*big.Int, uint32, error) {
	return getBiggestInfluenceAndIdMock(client, address, epoch, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) getIteration(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) int {
	return getIterationMock(client, address, proposer, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) isElectedProposer(client *ethclient.Client, address string, proposer types.ElectedProposer, utilsStruct UtilsStruct) bool {
	return isElectedProposerMock(client, address, proposer, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) pseudoRandomNumberGenerator(seed []byte, max uint32, blockHashes []byte) *big.Int {
	return pseudoRandomNumberGeneratorMock(seed, max, blockHashes)
}

func (proposeUtilsMock ProposeUtilsMock) MakeBlock(client *ethclient.Client, address string, rogueMode bool, utilsStruct UtilsStruct) ([]uint32, error) {
	return MakeBlockMock(client, address, rogueMode, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) getSortedVotes(client *ethclient.Client, address string, assetId uint8, epoch uint32, utilsStruct UtilsStruct) ([]*big.Int, error) {
	return getSortedVotesMock(client, address, assetId, epoch, utilsStruct)
}

func (proposeUtilsMock ProposeUtilsMock) influencedMedian(sortedVotes []*big.Int, totalInfluenceRevealed *big.Int) *big.Int {
	return influencedMedianMock(sortedVotes, totalInfluenceRevealed)
}

func (flagSetMock FlagSetMock) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8WeightMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8AssetId(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8AssetIdMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	return GetStringProviderMock(flagSet)
}

func (flagSetMock FlagSetMock) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	return GetFloat32GasMultiplierMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	return GetInt32BufferMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	return GetInt32WaitMock(flagSet)
}

func (flagSetMock FlagSetMock) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	return GetInt32GasPriceMock(flagSet)
}

func (flagSetMock FlagSetMock) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	return GetFloat32GasLimitMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	return GetStringLogLevelMock(flagSet)
}

func (flagSetMock FlagSetMock) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return GetStringStatusMock(flagSet)
}

func (voteManagerMock VoteManagerMock) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, values []*big.Int, secret [32]byte) (*Types.Transaction, error) {
	return RevealMock(client, opts, epoch, values, secret)
}

func (voteManagerMock VoteManagerMock) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	return CommitMock(client, opts, epoch, commitment)
}

func (blockManagerMock BlockManagerMock) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, medians []uint32, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	return ProposeMock(client, opts, epoch, medians, iteration, biggestInfluencerId)
}

func (flagSetMock FlagSetMock) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8CommissionMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return GetUintSliceJobIdsMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return GetUint32AggregationMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8JobId(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8JobIdMock(flagSet)
}

func (flagSetMock FlagSetMock) GetUint8CollectionId(flagSet *pflag.FlagSet) (uint8, error) {
	return GetUint8CollectionIdMock(flagSet)
}

func (utilsCmdMock UtilsCmdMock) SetCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return SetCommissionMock(client, stakerId, opts, commission, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) DecreaseCommission(client *ethclient.Client, stakerId uint32, opts *bind.TransactOpts, commission uint8, utilsStruct UtilsStruct) error {
	return DecreaseCommissionMock(client, stakerId, opts, commission, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) DecreaseCommissionPrompt() bool {
	return DecreaseCommissionPromptMock()
}

func (utilsCmdMock UtilsCmdMock) Withdraw(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, stakerId uint32, utilsStruct UtilsStruct) (common.Hash, error) {
	return WithdrawMock(client, txnOpts, epoch, stakerId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) CheckCurrentStatus(client *ethclient.Client, address string, assetId uint8, utilsStruct UtilsStruct) (bool, error) {
	return CheckCurrentStatusMock(client, address, assetId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) Dispute(client *ethclient.Client, config types.Configurations, account types.Account, epoch uint32, blockId uint8, assetId int, utilsStruct UtilsStruct) error {
	return DisputeMock(client, config, account, epoch, blockId, assetId, utilsStruct)
}

func (utilsCmdMock UtilsCmdMock) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnOpts *bind.TransactOpts, epoch uint32, assetId uint8, sortedStakers []uint32) {
	GiveSortedMock(client, blockManager, txnOpts, epoch, assetId, sortedStakers)
}

func (blockManagerMock BlockManagerMock) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	return ClaimBlockRewardMock(client, opts)
}

func (blockManagerMock BlockManagerMock) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8) (*Types.Transaction, error) {
	return FinalizeDisputeMock(client, opts, epoch, blockIndex)
}

func (blockManagerMock BlockManagerMock) DisputeBiggestInfluenceProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestInfluencerId uint32) (*Types.Transaction, error) {
	return DisputeBiggestInfluenceProposedMock(client, opts, epoch, blockIndex, correctBiggestInfluencerId)
}

func (c CryptoMock) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return HexToECDSAMock(hexKey)
}
