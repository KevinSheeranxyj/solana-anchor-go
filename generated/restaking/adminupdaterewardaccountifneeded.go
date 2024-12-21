// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminUpdateRewardAccountIfNeeded is the `admin_update_reward_account_if_needed` instruction.
type AdminUpdateRewardAccountIfNeeded struct {
	DesiredAccountSize *uint32 `bin:"optional"`

	// [0] = [WRITE, SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] system_program
	//
	// [3] = [] receipt_token_mint
	//
	// [4] = [WRITE] reward_account
	//
	// [5] = [] event_authority
	//
	// [6] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminUpdateRewardAccountIfNeededInstructionBuilder creates a new `AdminUpdateRewardAccountIfNeeded` instruction builder.
func NewAdminUpdateRewardAccountIfNeededInstructionBuilder() *AdminUpdateRewardAccountIfNeeded {
	nd := &AdminUpdateRewardAccountIfNeeded{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["9b2RSMDYskVvjVbwF4cVwEhZUaaaUgyYSxvESmnoS4LL"]).SIGNER()
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	return nd
}

// SetDesiredAccountSize sets the "desired_account_size" parameter.
func (inst *AdminUpdateRewardAccountIfNeeded) SetDesiredAccountSize(desired_account_size uint32) *AdminUpdateRewardAccountIfNeeded {
	inst.DesiredAccountSize = &desired_account_size
	return inst
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetPayerAccount(payer ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetAdminAccount(admin ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRewardAccountAccount sets the "reward_account" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetRewardAccountAccount(rewardAccount ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rewardAccount).WRITE()
	return inst
}

func (inst *AdminUpdateRewardAccountIfNeeded) findFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateRewardAccountIfNeeded) FindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateRewardAccountIfNeeded) MustFindRewardAccountAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindRewardAccountAddress finds RewardAccount account address with given seeds.
func (inst *AdminUpdateRewardAccountIfNeeded) FindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateRewardAccountIfNeeded) MustFindRewardAccountAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindRewardAccountAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetRewardAccountAccount gets the "reward_account" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *AdminUpdateRewardAccountIfNeeded) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateRewardAccountIfNeeded) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *AdminUpdateRewardAccountIfNeeded) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *AdminUpdateRewardAccountIfNeeded) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *AdminUpdateRewardAccountIfNeeded) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetProgramAccount sets the "program" account.
func (inst *AdminUpdateRewardAccountIfNeeded) SetProgramAccount(program ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AdminUpdateRewardAccountIfNeeded) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst AdminUpdateRewardAccountIfNeeded) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminUpdateRewardAccountIfNeeded,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminUpdateRewardAccountIfNeeded) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminUpdateRewardAccountIfNeeded) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.RewardAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AdminUpdateRewardAccountIfNeeded) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminUpdateRewardAccountIfNeeded")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  DesiredAccountSize (OPT)", inst.DesiredAccountSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("             admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    system_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("receipt_token_mint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           reward_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   event_authority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj AdminUpdateRewardAccountIfNeeded) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DesiredAccountSize` param (optional):
	{
		if obj.DesiredAccountSize == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *AdminUpdateRewardAccountIfNeeded) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DesiredAccountSize` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.DesiredAccountSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewAdminUpdateRewardAccountIfNeededInstruction declares a new AdminUpdateRewardAccountIfNeeded instruction with the provided parameters and accounts.
func NewAdminUpdateRewardAccountIfNeededInstruction(
	// Parameters:
	desired_account_size uint32,
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	rewardAccount ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *AdminUpdateRewardAccountIfNeeded {
	return NewAdminUpdateRewardAccountIfNeededInstructionBuilder().
		SetDesiredAccountSize(desired_account_size).
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetRewardAccountAccount(rewardAccount).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
