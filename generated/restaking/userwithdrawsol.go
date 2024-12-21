// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UserWithdrawSol is the `user_withdraw_sol` instruction.
type UserWithdrawSol struct {
	RequestId *uint64

	// [0] = [WRITE, SIGNER] user
	//
	// [1] = [] system_program
	//
	// [2] = [] receipt_token_program
	//
	// [3] = [] receipt_token_mint
	//
	// [4] = [] user_receipt_token_account
	//
	// [5] = [WRITE] fund_account
	//
	// [6] = [WRITE] fund_withdrawal_batch_account
	// ··········· Users can derive proper account address with target batch id for each withdrawal requests.
	// ··········· And the batch id can be read from a user fund account which the withdrawal requests belong to.
	// ··········· seeds = [FundWithdrawalBatchAccount::SEED, receipt_token_mint.key().as_ref(), Pubkey::default().as_ref(), &fund_withdrawal_batch_account.batch_id.to_le_bytes()],
	// ··········· bump = fund_withdrawal_batch_account.get_bump(),
	//
	// [7] = [WRITE] fund_reserve_account
	//
	// [8] = [WRITE] fund_treasury_account
	//
	// [9] = [WRITE] user_fund_account
	//
	// [10] = [] reward_account
	//
	// [11] = [] user_reward_account
	//
	// [12] = [] event_authority
	//
	// [13] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUserWithdrawSolInstructionBuilder creates a new `UserWithdrawSol` instruction builder.
func NewUserWithdrawSolInstructionBuilder() *UserWithdrawSol {
	nd := &UserWithdrawSol{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"])
	return nd
}

// SetRequestId sets the "request_id" parameter.
func (inst *UserWithdrawSol) SetRequestId(request_id uint64) *UserWithdrawSol {
	inst.RequestId = &request_id
	return inst
}

// SetUserAccount sets the "user" account.
func (inst *UserWithdrawSol) SetUserAccount(user ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UserWithdrawSol) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UserWithdrawSol) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UserWithdrawSol) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenProgramAccount sets the "receipt_token_program" account.
func (inst *UserWithdrawSol) SetReceiptTokenProgramAccount(receiptTokenProgram ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenProgram)
	return inst
}

// GetReceiptTokenProgramAccount gets the "receipt_token_program" account.
func (inst *UserWithdrawSol) GetReceiptTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *UserWithdrawSol) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *UserWithdrawSol) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserReceiptTokenAccountAccount sets the "user_receipt_token_account" account.
func (inst *UserWithdrawSol) SetUserReceiptTokenAccountAccount(userReceiptTokenAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userReceiptTokenAccount)
	return inst
}

func (inst *UserWithdrawSol) findFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// path: user
	seeds = append(seeds, user.Bytes())
	// path: receiptTokenProgram
	seeds = append(seeds, receiptTokenProgram.Bytes())
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	programID := Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"]

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, programID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, programID)
	}
	return
}

// FindUserReceiptTokenAccountAddressWithBumpSeed calculates UserReceiptTokenAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindUserReceiptTokenAccountAddressWithBumpSeed(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserReceiptTokenAccountAddress finds UserReceiptTokenAccount account address with given seeds.
func (inst *UserWithdrawSol) FindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	return
}

func (inst *UserWithdrawSol) MustFindUserReceiptTokenAccountAddress(user ag_solanago.PublicKey, receiptTokenProgram ag_solanago.PublicKey, receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserReceiptTokenAccountAddress(user, receiptTokenProgram, receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserReceiptTokenAccountAccount gets the "user_receipt_token_account" account.
func (inst *UserWithdrawSol) GetUserReceiptTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *UserWithdrawSol) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *UserWithdrawSol) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundAccountAddressWithBumpSeed calculates FundAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *UserWithdrawSol) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWithdrawSol) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *UserWithdrawSol) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetFundWithdrawalBatchAccountAccount sets the "fund_withdrawal_batch_account" account.
// Users can derive proper account address with target batch id for each withdrawal requests.
// And the batch id can be read from a user fund account which the withdrawal requests belong to.
// seeds = [FundWithdrawalBatchAccount::SEED, receipt_token_mint.key().as_ref(), Pubkey::default().as_ref(), &fund_withdrawal_batch_account.batch_id.to_le_bytes()],
// bump = fund_withdrawal_batch_account.get_bump(),
func (inst *UserWithdrawSol) SetFundWithdrawalBatchAccountAccount(fundWithdrawalBatchAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(fundWithdrawalBatchAccount).WRITE()
	return inst
}

// GetFundWithdrawalBatchAccountAccount gets the "fund_withdrawal_batch_account" account.
// Users can derive proper account address with target batch id for each withdrawal requests.
// And the batch id can be read from a user fund account which the withdrawal requests belong to.
// seeds = [FundWithdrawalBatchAccount::SEED, receipt_token_mint.key().as_ref(), Pubkey::default().as_ref(), &fund_withdrawal_batch_account.batch_id.to_le_bytes()],
// bump = fund_withdrawal_batch_account.get_bump(),
func (inst *UserWithdrawSol) GetFundWithdrawalBatchAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFundReserveAccountAccount sets the "fund_reserve_account" account.
func (inst *UserWithdrawSol) SetFundReserveAccountAccount(fundReserveAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(fundReserveAccount).WRITE()
	return inst
}

func (inst *UserWithdrawSol) findFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_reserve
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x72), byte(0x65), byte(0x73), byte(0x65), byte(0x72), byte(0x76), byte(0x65)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundReserveAccountAddressWithBumpSeed calculates FundReserveAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindFundReserveAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundReserveAccountAddress finds FundReserveAccount account address with given seeds.
func (inst *UserWithdrawSol) FindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWithdrawSol) MustFindFundReserveAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundReserveAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundReserveAccountAccount gets the "fund_reserve_account" account.
func (inst *UserWithdrawSol) GetFundReserveAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetFundTreasuryAccountAccount sets the "fund_treasury_account" account.
func (inst *UserWithdrawSol) SetFundTreasuryAccountAccount(fundTreasuryAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(fundTreasuryAccount).WRITE()
	return inst
}

func (inst *UserWithdrawSol) findFindFundTreasuryAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: fund_treasury
	seeds = append(seeds, []byte{byte(0x66), byte(0x75), byte(0x6e), byte(0x64), byte(0x5f), byte(0x74), byte(0x72), byte(0x65), byte(0x61), byte(0x73), byte(0x75), byte(0x72), byte(0x79)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindFundTreasuryAccountAddressWithBumpSeed calculates FundTreasuryAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindFundTreasuryAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundTreasuryAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindFundTreasuryAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTreasuryAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundTreasuryAccountAddress finds FundTreasuryAccount account address with given seeds.
func (inst *UserWithdrawSol) FindFundTreasuryAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundTreasuryAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWithdrawSol) MustFindFundTreasuryAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundTreasuryAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundTreasuryAccountAccount gets the "fund_treasury_account" account.
func (inst *UserWithdrawSol) GetFundTreasuryAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetUserFundAccountAccount sets the "user_fund_account" account.
func (inst *UserWithdrawSol) SetUserFundAccountAccount(userFundAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(userFundAccount).WRITE()
	return inst
}

func (inst *UserWithdrawSol) findFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_fund
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x66), byte(0x75), byte(0x6e), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserFundAccountAddressWithBumpSeed calculates UserFundAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindUserFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserFundAccountAddress finds UserFundAccount account address with given seeds.
func (inst *UserWithdrawSol) FindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserWithdrawSol) MustFindUserFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserFundAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserFundAccountAccount gets the "user_fund_account" account.
func (inst *UserWithdrawSol) GetUserFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *UserWithdrawSol) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rewardAccount)
	return inst
}

func (inst *UserWithdrawSol) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: reward
	seeds = append(seeds, []byte{byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindRewardAccountAddressWithBumpSeed calculates RewardAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *UserWithdrawSol) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *UserWithdrawSol) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *UserWithdrawSol) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetUserRewardAccountAccount sets the "user_reward_account" account.
func (inst *UserWithdrawSol) SetUserRewardAccountAccount(userRewardAccount ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(userRewardAccount)
	return inst
}

func (inst *UserWithdrawSol) findFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: user_reward
	seeds = append(seeds, []byte{byte(0x75), byte(0x73), byte(0x65), byte(0x72), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64)})
	// path: receiptTokenMint
	seeds = append(seeds, receiptTokenMint.Bytes())
	// path: user
	seeds = append(seeds, user.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindUserRewardAccountAddressWithBumpSeed calculates UserRewardAccount account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindUserRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindUserRewardAccountAddress finds UserRewardAccount account address with given seeds.
func (inst *UserWithdrawSol) FindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	return
}

func (inst *UserWithdrawSol) MustFindUserRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindUserRewardAccountAddress(receiptTokenMint, user, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetUserRewardAccountAccount gets the "user_reward_account" account.
func (inst *UserWithdrawSol) GetUserRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *UserWithdrawSol) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *UserWithdrawSol) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: __event_authority
	seeds = append(seeds, []byte{byte(0x5f), byte(0x5f), byte(0x65), byte(0x76), byte(0x65), byte(0x6e), byte(0x74), byte(0x5f), byte(0x61), byte(0x75), byte(0x74), byte(0x68), byte(0x6f), byte(0x72), byte(0x69), byte(0x74), byte(0x79)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindEventAuthorityAddressWithBumpSeed calculates EventAuthority account address with given seeds and a known bump seed.
func (inst *UserWithdrawSol) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *UserWithdrawSol) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *UserWithdrawSol) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *UserWithdrawSol) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *UserWithdrawSol) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetProgramAccount sets the "program" account.
func (inst *UserWithdrawSol) SetProgramAccount(program ag_solanago.PublicKey) *UserWithdrawSol {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *UserWithdrawSol) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst UserWithdrawSol) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UserWithdrawSol,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UserWithdrawSol) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UserWithdrawSol) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RequestId == nil {
			return errors.New("RequestId parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserReceiptTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FundWithdrawalBatchAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FundReserveAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.FundTreasuryAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.UserFundAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.UserRewardAccount is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *UserWithdrawSol) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UserWithdrawSol")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" RequestId", *inst.RequestId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  user", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        system_program", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" receipt_token_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   user_receipt_token_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 fund_", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("fund_withdrawal_batch_", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         fund_reserve_", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        fund_treasury_", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("            user_fund_", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("               reward_", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          user_reward_", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("       event_authority", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("               program", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj UserWithdrawSol) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RequestId` param:
	err = encoder.Encode(obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UserWithdrawSol) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RequestId`:
	err = decoder.Decode(&obj.RequestId)
	if err != nil {
		return err
	}
	return nil
}

// NewUserWithdrawSolInstruction declares a new UserWithdrawSol instruction with the provided parameters and accounts.
func NewUserWithdrawSolInstruction(
	// Parameters:
	request_id uint64,
	// Accounts:
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	userReceiptTokenAccount ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	fundWithdrawalBatchAccount ag_solanago.PublicKey,
	fundReserveAccount ag_solanago.PublicKey,
	fundTreasuryAccount ag_solanago.PublicKey,
	userFundAccount ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	userRewardAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *UserWithdrawSol {
	return NewUserWithdrawSolInstructionBuilder().
		SetRequestId(request_id).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenProgramAccount(receiptTokenProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetUserReceiptTokenAccountAccount(userReceiptTokenAccount).
		SetFundAccountAccount(fundAccount).
		SetFundWithdrawalBatchAccountAccount(fundWithdrawalBatchAccount).
		SetFundReserveAccountAccount(fundReserveAccount).
		SetFundTreasuryAccountAccount(fundTreasuryAccount).
		SetUserFundAccountAccount(userFundAccount).
		SetRewardAccountAccount(rewardAccount).
		SetUserRewardAccountAccount(userRewardAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
