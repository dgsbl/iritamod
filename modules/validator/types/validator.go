package types

import (
	"sort"
	"strings"

	"github.com/tendermint/tendermint/crypto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ stakingtypes.ValidatorI = Validator{}

// DoNotModifyDesc used in flags to indicate that description field should not be updated
const DoNotModifyDesc = "[do-not-modify]"

// NewValidator creates a new MsgCreateValidator instance.
func NewValidator(
	id tmbytes.HexBytes,
	name, description string,
	pubKey crypto.PubKey,
	cert string,
	power int64,
	operator sdk.AccAddress,
) Validator {
	var pkStr string
	if pubKey != nil {
		pkStr = sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, pubKey)
	}

	return Validator{
		Id:          id.String(),
		Name:        name,
		Pubkey:      pkStr,
		Certificate: cert,
		Power:       power,
		Description: description,
		Operator:    operator.String(),
	}
}

// IsJailed implement ValidatorI
func (v Validator) IsJailed() bool {
	return v.Jailed
}

// GetMoniker implement ValidatorI
func (v Validator) GetMoniker() string {
	return v.Name
}

// GetStatus implement ValidatorI
func (v Validator) GetStatus() stakingtypes.BondStatus {
	if v.Jailed {
		return stakingtypes.Unbonded
	} else {
		return stakingtypes.Bonded
	}
}

// IsBonded implement ValidatorI
func (v Validator) IsBonded() bool {
	return !v.Jailed
}

// IsUnbonded implement ValidatorI
func (v Validator) IsUnbonded() bool {
	return v.Jailed
}

// IsUnbonding implement ValidatorI
func (v Validator) IsUnbonding() bool {
	return false
}

// GetOperator implement ValidatorI
func (v Validator) GetOperator() sdk.ValAddress {
	pubkey, err := v.TmConsPubKey()
	if err != nil {
		panic(err)
	}
	return sdk.ValAddress(pubkey.Address())
}

// TmConsPubKey implement ValidatorI
func (v Validator) TmConsPubKey() (crypto.PubKey, error) {
	pk, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.Pubkey)
	if err != nil {
		return nil, err
	}

	// The way things are refactored now, v.ConsensusPubkey is sometimes a TM
	// ed25519 pubkey, sometimes our own ed25519 pubkey. This is very ugly and
	// inconsistent.
	// Luckily, here we coerce it into a TM ed25519 pubkey always, as this
	// pubkey will be passed into TM (eg calling encoding.PubKeyToProto).
	if intoTmPk, ok := pk.(cryptotypes.IntoTmPubKey); ok {
		return intoTmPk.AsTmPubKey(), nil
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "Logic error: ConsensusPubkey must be an SDK key and SDK PubKey types must be convertible to tendermint PubKey; got: %T", pk)
}

// GetConsAddr implement ValidatorI
func (v Validator) GetConsAddr() (sdk.ConsAddress, error) {
	pk, err := v.TmConsPubKey()
	if err != nil {
		return sdk.ConsAddress{}, err
	}
	return sdk.ConsAddress(pk.Address()), nil
}

// GetTokens implement ValidatorI
func (v Validator) GetTokens() sdk.Int {
	return sdk.TokensFromConsensusPower(v.Power)
}

// GetBondedTokens implement ValidatorI
func (v Validator) GetBondedTokens() sdk.Int {
	if v.Jailed {
		return sdk.NewInt(0)
	}
	return sdk.TokensFromConsensusPower(v.Power)
}

// GetConsensusPower implement ValidatorI
func (v Validator) GetConsensusPower() int64 {
	return v.Power
}

// GetCommission implement ValidatorI
func (v Validator) GetCommission() sdk.Dec {
	return sdk.NewDec(0)
}

// GetMinSelfDelegation implement ValidatorI
func (v Validator) GetMinSelfDelegation() sdk.Int {
	return sdk.NewInt(0)
}

// GetDelegatorShares implement ValidatorI
func (v Validator) GetDelegatorShares() sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromShares implement ValidatorI
func (v Validator) TokensFromShares(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromSharesTruncated implement ValidatorI
func (v Validator) TokensFromSharesTruncated(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// TokensFromSharesRoundUp implement ValidatorI
func (v Validator) TokensFromSharesRoundUp(dec sdk.Dec) sdk.Dec {
	return sdk.NewDec(0)
}

// SharesFromTokens implement ValidatorI
func (v Validator) SharesFromTokens(amt sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(0), nil
}

// SharesFromTokensTruncated implement ValidatorI
func (v Validator) SharesFromTokensTruncated(amt sdk.Int) (sdk.Dec, error) {
	return sdk.NewDec(0), nil

}

// Validators is a collection of Validator
type Validators []Validator

// Sort Validators sorts validator array in ascending operator address order
func (v Validators) Sort() {
	sort.Sort(v)
}

// Implements sort interface
func (v Validators) Len() int {
	return len(v)
}

// Implements sort interface
func (v Validators) Less(i, j int) bool {
	return strings.Compare(v[i].Name, v[j].Name) == -1
}

// Implements sort interface
func (v Validators) Swap(i, j int) {
	it := v[i]
	v[i] = v[j]
	v[j] = it
}
