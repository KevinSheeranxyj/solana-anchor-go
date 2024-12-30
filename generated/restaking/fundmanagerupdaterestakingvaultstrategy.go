// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundManagerUpdateRestakingVaultStrategy is the `fund_manager_update_restaking_vault_strategy` instruction.
type FundManagerUpdateRestakingVaultStrategy struct {
	Vault                       *ag_solanago.PublicKey
	SolAllocationWeight         *uint64
	SolAllocationCapacityAmount *uint64

	// [0] = [SIGNER] fund_manager
	//
	// [1] = [] receipt_token_mint
	//
	// [2] = [WRITE] fund_account
	//
	// [3] = [] event_authority
	//
	// [4] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundManagerUpdateRestakingVaultStrategyInstructionBuilder creates a new `FundManagerUpdateRestakingVaultStrategy` instruction builder.
func NewFundManagerUpdateRestakingVaultStrategyInstructionBuilder() *FundManagerUpdateRestakingVaultStrategy {
	nd := &FundManagerUpdateRestakingVaultStrategy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[0] = ag_solanago.Meta(Addresses["5UpLTLA7Wjqp7qdfjuTtPcUw3aVtbqFA5Mgm34mxPNg2"]).SIGNER()
	return nd
}

// SetVault sets the "vault" parameter.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetVault(vault ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.Vault = &vault
	return inst
}

// SetSolAllocationWeight sets the "sol_allocation_weight" parameter.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetSolAllocationWeight(sol_allocation_weight uint64) *FundManagerUpdateRestakingVaultStrategy {
	inst.SolAllocationWeight = &sol_allocation_weight
	return inst
}

// SetSolAllocationCapacityAmount sets the "sol_allocation_capacity_amount" parameter.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetSolAllocationCapacityAmount(sol_allocation_capacity_amount uint64) *FundManagerUpdateRestakingVaultStrategy {
	inst.SolAllocationCapacityAmount = &sol_allocation_capacity_amount
	return inst
}

// SetFundManagerAccount sets the "fund_manager" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetFundManagerAccount(fundManager ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(fundManager).SIGNER()
	return inst
}

// GetFundManagerAccount gets the "fund_manager" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) GetFundManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFundAccountAccount sets the "fund_account" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetFundAccountAccount(fundAccount ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(fundAccount).WRITE()
	return inst
}

func (inst *FundManagerUpdateRestakingVaultStrategy) findFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerUpdateRestakingVaultStrategy) FindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *FundManagerUpdateRestakingVaultStrategy) MustFindFundAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindFundAccountAddress finds FundAccount account address with given seeds.
func (inst *FundManagerUpdateRestakingVaultStrategy) FindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindFundAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *FundManagerUpdateRestakingVaultStrategy) MustFindFundAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindFundAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetFundAccountAccount gets the "fund_account" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) GetFundAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *FundManagerUpdateRestakingVaultStrategy) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *FundManagerUpdateRestakingVaultStrategy) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *FundManagerUpdateRestakingVaultStrategy) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *FundManagerUpdateRestakingVaultStrategy) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *FundManagerUpdateRestakingVaultStrategy) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) SetProgramAccount(program ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundManagerUpdateRestakingVaultStrategy) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst FundManagerUpdateRestakingVaultStrategy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundManagerUpdateRestakingVaultStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundManagerUpdateRestakingVaultStrategy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundManagerUpdateRestakingVaultStrategy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Vault == nil {
			return errors.New("Vault parameter is not set")
		}
		if inst.SolAllocationWeight == nil {
			return errors.New("SolAllocationWeight parameter is not set")
		}
		if inst.SolAllocationCapacityAmount == nil {
			return errors.New("SolAllocationCapacityAmount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.FundManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FundAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *FundManagerUpdateRestakingVaultStrategy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundManagerUpdateRestakingVaultStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("                         Vault", *inst.Vault))
						paramsBranch.Child(ag_format.Param("           SolAllocationWeight", *inst.SolAllocationWeight))
						paramsBranch.Child(ag_format.Param("   SolAllocationCapacityAmount", *inst.SolAllocationCapacityAmount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      fund_manager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("             fund_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj FundManagerUpdateRestakingVaultStrategy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Vault` param:
	err = encoder.Encode(obj.Vault)
	if err != nil {
		return err
	}
	// Serialize `SolAllocationWeight` param:
	err = encoder.Encode(obj.SolAllocationWeight)
	if err != nil {
		return err
	}
	// Serialize `SolAllocationCapacityAmount` param:
	err = encoder.Encode(obj.SolAllocationCapacityAmount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundManagerUpdateRestakingVaultStrategy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Vault`:
	err = decoder.Decode(&obj.Vault)
	if err != nil {
		return err
	}
	// Deserialize `SolAllocationWeight`:
	err = decoder.Decode(&obj.SolAllocationWeight)
	if err != nil {
		return err
	}
	// Deserialize `SolAllocationCapacityAmount`:
	err = decoder.Decode(&obj.SolAllocationCapacityAmount)
	if err != nil {
		return err
	}
	return nil
}

// NewFundManagerUpdateRestakingVaultStrategyInstruction declares a new FundManagerUpdateRestakingVaultStrategy instruction with the provided parameters and accounts.
func NewFundManagerUpdateRestakingVaultStrategyInstruction(
	// Parameters:
	vault ag_solanago.PublicKey,
	sol_allocation_weight uint64,
	sol_allocation_capacity_amount uint64,
	// Accounts:
	fundManager ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	fundAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundManagerUpdateRestakingVaultStrategy {
	return NewFundManagerUpdateRestakingVaultStrategyInstructionBuilder().
		SetVault(vault).
		SetSolAllocationWeight(sol_allocation_weight).
		SetSolAllocationCapacityAmount(sol_allocation_capacity_amount).
		SetFundManagerAccount(fundManager).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetFundAccountAccount(fundAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
