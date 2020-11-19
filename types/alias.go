// Package types nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-project/core/types/assets/
// ALIASGEN: github.com/terra-project/core/types/util/
package types

import (
	"github.com/terra-project/core/types/assets"
	"github.com/terra-project/core/types/util"
)

const (
	MicroLunaDenom       = assets.MicroLunaDenom
	MicroUSDDenom        = assets.MicroUSDDenom
	MicroKRWDenom        = assets.MicroKRWDenom
	MicroSDRDenom        = assets.MicroSDRDenom
	MicroCNYDenom        = assets.MicroCNYDenom
	MicroJPYDenom        = assets.MicroJPYDenom
	MicroEURDenom        = assets.MicroEURDenom
	MicroGBPDenom        = assets.MicroGBPDenom
	MicroMNTDenom        = assets.MicroMNTDenom
	MicroUnit            = assets.MicroUnit
	BlocksPerMinute      = util.BlocksPerMinute
	BlocksPerHour        = util.BlocksPerHour
	BlocksPerDay         = util.BlocksPerDay
	BlocksPerWeek        = util.BlocksPerWeek
	BlocksPerMonth       = util.BlocksPerMonth
	BlocksPerYear        = util.BlocksPerYear
	CoinType             = util.CoinType
	FullFundraiserPath   = util.FullFundraiserPath
	Bech32PrefixAccAddr  = util.Bech32PrefixAccAddr
	Bech32PrefixAccPub   = util.Bech32PrefixAccPub
	Bech32PrefixValAddr  = util.Bech32PrefixValAddr
	Bech32PrefixValPub   = util.Bech32PrefixValPub
	Bech32PrefixConsAddr = util.Bech32PrefixConsAddr
	Bech32PrefixConsPub  = util.Bech32PrefixConsPub
)

var (
	// functions aliases
	IsPeriodLastBlock    = util.IsPeriodLastBlock
	IsWaitingForSoftfork = util.IsWaitingForSoftfork
)

type (
	Base64Bytes = util.Base64Bytes
)
