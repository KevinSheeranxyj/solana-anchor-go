// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package restaking

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminUpdateExtraAccountMetaListIfNeeded is the `admin_update_extra_account_meta_list_if_needed` instruction.
type AdminUpdateExtraAccountMetaListIfNeeded struct {

	// [0] = [SIGNER] payer
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] receipt_token_mint
	//
	// [3] = [WRITE] extra_account_meta_list
	//
	// [4] = [] event_authority
	//
	// [5] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminUpdateExtraAccountMetaListIfNeededInstructionBuilder creates a new `AdminUpdateExtraAccountMetaListIfNeeded` instruction builder.
func NewAdminUpdateExtraAccountMetaListIfNeededInstructionBuilder() *AdminUpdateExtraAccountMetaListIfNeeded {
	nd := &AdminUpdateExtraAccountMetaListIfNeeded{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[1] = ag_solanago.Meta(Addresses["9b2RSMDYskVvjVbwF4cVwEhZUaaaUgyYSxvESmnoS4LL"]).SIGNER()
	return nd
}

// SetPayerAccount sets the "payer" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetPayerAccount(payer ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetAdminAccount(admin ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReceiptTokenMintAccount sets the "receipt_token_mint" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetReceiptTokenMintAccount(receiptTokenMint ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(receiptTokenMint)
	return inst
}

// GetReceiptTokenMintAccount gets the "receipt_token_mint" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetReceiptTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetExtraAccountMetaListAccount sets the "extra_account_meta_list" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetExtraAccountMetaListAccount(extraAccountMetaList ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(extraAccountMetaList).WRITE()
	return inst
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) findFindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: extra-account-metas
	seeds = append(seeds, []byte{byte(0x65), byte(0x78), byte(0x74), byte(0x72), byte(0x61), byte(0x2d), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74), byte(0x2d), byte(0x6d), byte(0x65), byte(0x74), byte(0x61), byte(0x73)})
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

// FindExtraAccountMetaListAddressWithBumpSeed calculates ExtraAccountMetaList account address with given seeds and a known bump seed.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) FindExtraAccountMetaListAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindExtraAccountMetaListAddress(receiptTokenMint, bumpSeed)
	return
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) MustFindExtraAccountMetaListAddressWithBumpSeed(receiptTokenMint ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindExtraAccountMetaListAddress(receiptTokenMint, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindExtraAccountMetaListAddress finds ExtraAccountMetaList account address with given seeds.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) FindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindExtraAccountMetaListAddress(receiptTokenMint, 0)
	return
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) MustFindExtraAccountMetaListAddress(receiptTokenMint ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindExtraAccountMetaListAddress(receiptTokenMint, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetExtraAccountMetaListAccount gets the "extra_account_meta_list" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetExtraAccountMetaListAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetEventAuthorityAccount sets the "event_authority" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(eventAuthority)
	return inst
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) findFindEventAuthorityAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) FindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindEventAuthorityAddress(bumpSeed)
	return
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) MustFindEventAuthorityAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindEventAuthorityAddress finds EventAuthority account address with given seeds.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) FindEventAuthorityAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindEventAuthorityAddress(0)
	return
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) MustFindEventAuthorityAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindEventAuthorityAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetEventAuthorityAccount gets the "event_authority" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProgramAccount sets the "program" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) SetProgramAccount(program ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AdminUpdateExtraAccountMetaListIfNeeded) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst AdminUpdateExtraAccountMetaListIfNeeded) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminUpdateExtraAccountMetaListIfNeeded,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminUpdateExtraAccountMetaListIfNeeded) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReceiptTokenMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ExtraAccountMetaList is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AdminUpdateExtraAccountMetaListIfNeeded) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminUpdateExtraAccountMetaListIfNeeded")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                  payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                  admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     receipt_token_mint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("extra_account_meta_list", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        event_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj AdminUpdateExtraAccountMetaListIfNeeded) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AdminUpdateExtraAccountMetaListIfNeeded) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAdminUpdateExtraAccountMetaListIfNeededInstruction declares a new AdminUpdateExtraAccountMetaListIfNeeded instruction with the provided parameters and accounts.
func NewAdminUpdateExtraAccountMetaListIfNeededInstruction(
	// Accounts:
	payer ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	receiptTokenMint ag_solanago.PublicKey,
	extraAccountMetaList ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *AdminUpdateExtraAccountMetaListIfNeeded {
	return NewAdminUpdateExtraAccountMetaListIfNeededInstructionBuilder().
		SetPayerAccount(payer).
		SetAdminAccount(admin).
		SetReceiptTokenMintAccount(receiptTokenMint).
		SetExtraAccountMetaListAccount(extraAccountMetaList).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
