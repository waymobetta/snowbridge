package parachain

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/snowfork/go-substrate-rpc-client/v3/types"
	"github.com/snowfork/snowbridge/relayer/contracts/basic"
	"github.com/snowfork/snowbridge/relayer/contracts/incentivized"
)

type ParaheadPartialLog struct {
	ParentHash     string `json:"parentHash"`
	Number         uint32 `json:"number"`
	StateRoot      string `json:"stateRoot"`
	ExtrinsicsRoot string `json:"extrinsicsRoot"`
}

type ParaHeadProofLog struct {
	Pos   *big.Int `json:"pos"`
	Width *big.Int `json:"width"`
	Proof []string `json:"proof"`
}

type BeefyMMRLeafPartialLog struct {
	ParentNumber         uint32 `json:"parentNumber"`
	ParentHash           string `json:"parentHash"`
	NextAuthoritySetId   uint64 `json:"nextAuthoritySetId"` // revive:disable-line
	NextAuthoritySetLen  uint32 `json:"nextAuthoritySetLen"`
	NextAuthoritySetRoot string `json:"nextAuthoritySetRoot"`
}

type BasicInboundChannelMessageLog struct {
	Target  common.Address `json:"target"`
	Nonce   uint64         `json:"nonce"`
	Payload string         `json:"payload"`
}

type IncentivizedInboundChannelMessageLog struct {
	Target  common.Address `json:"target"`
	Nonce   uint64         `json:"nonce"`
	Fee     *big.Int       `json:"fee"`
	Payload string         `json:"payload"`
}

type BasicSubmitInput struct {
	Messages            []BasicInboundChannelMessageLog `json:"_messages"`
	ParaheadPartial     ParaheadPartialLog              `json:"_ownParachainHeadPartial"`
	ParaHeadProof       ParaHeadProofLog                `json:"_parachainHeadProof"`
	BeefyMMRLeafPartial BeefyMMRLeafPartialLog          `json:"_beefyMMRLeafPartial"`
	BeefyMMRLeafIndex   int64                           `json:"_beefyMMRLeafIndex"`
	BeefyLeafCount      int64                           `json:"_beefyLeafCount"`
	BeefyMMRProof       []string                        `json:"_beefyMMRProof"`
}

type IncentivizedSubmitInput struct {
	Messages            []IncentivizedInboundChannelMessageLog `json:"_messages"`
	ParaheadPartial     ParaheadPartialLog                     `json:"_ownParachainHeadPartial"`
	ParaHeadProof       ParaHeadProofLog                       `json:"_parachainHeadProof"`
	BeefyMMRLeafPartial BeefyMMRLeafPartialLog                 `json:"_beefyMMRLeafPartial"`
	BeefyMMRLeafIndex   int64                                  `json:"_beefyMMRLeafIndex"`
	BeefyLeafCount      int64                                  `json:"_beefyLeafCount"`
	BeefyMMRProof       []string                               `json:"_beefyMMRProof"`
}

func (wr *EthereumChannelWriter) logBasicTx(
	messages []basic.BasicInboundChannelMessage,
	paraheadPartial basic.ParachainLightClientOwnParachainHeadPartial,
	paraHeadProof basic.ParachainLightClientParachainHeadProof,
	beefyMMRLeafPartial basic.ParachainLightClientBeefyMMRLeafPartial,
	beefyMMRLeafIndex int64, beefyLeafCount int64, beefyMMRProof [][32]byte,
	paraHead types.Header, paraHeadLeaf []byte, paraHeadProofRoot []byte, mmrLeaf types.MMRLeaf,
	commitmentHash types.H256,
) error {

	var basicMessagesLog []BasicInboundChannelMessageLog
	for _, item := range messages {
		basicMessagesLog = append(basicMessagesLog, BasicInboundChannelMessageLog{
			Target:  item.Target,
			Nonce:   item.Nonce,
			Payload: "0x" + hex.EncodeToString(item.Payload),
		})
	}
	var paraHeadProofString []string
	for _, item := range paraHeadProof.Proof {
		paraHeadProofString = append(paraHeadProofString, "0x"+hex.EncodeToString(item[:]))
	}
	var beefyMMRProofString []string
	for _, item := range beefyMMRProof {
		beefyMMRProofString = append(beefyMMRProofString, "0x"+hex.EncodeToString(item[:]))
	}
	input := &BasicSubmitInput{
		Messages: basicMessagesLog,
		ParaheadPartial: ParaheadPartialLog{
			ParentHash:     "0x" + hex.EncodeToString(paraheadPartial.ParentHash[:]),
			Number:         paraheadPartial.Number,
			StateRoot:      "0x" + hex.EncodeToString(paraheadPartial.StateRoot[:]),
			ExtrinsicsRoot: "0x" + hex.EncodeToString(paraheadPartial.ExtrinsicsRoot[:]),
		},
		ParaHeadProof: ParaHeadProofLog{
			Pos:   paraHeadProof.Pos,
			Width: paraHeadProof.Width,
			Proof: paraHeadProofString,
		},
		BeefyMMRLeafPartial: BeefyMMRLeafPartialLog{
			ParentNumber:         beefyMMRLeafPartial.ParentNumber,
			ParentHash:           "0x" + hex.EncodeToString(beefyMMRLeafPartial.ParentHash[:]),
			NextAuthoritySetId:   beefyMMRLeafPartial.NextAuthoritySetId,
			NextAuthoritySetLen:  beefyMMRLeafPartial.NextAuthoritySetLen,
			NextAuthoritySetRoot: "0x" + hex.EncodeToString(beefyMMRLeafPartial.NextAuthoritySetRoot[:]),
		},
		BeefyMMRLeafIndex: beefyMMRLeafIndex,
		BeefyLeafCount:    beefyLeafCount,
		BeefyMMRProof:     beefyMMRProofString,
	}
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	mmrLeafEncoded, _ := types.EncodeToBytes(mmrLeaf)
	mmrLeafOpaqueEncoded, _ := types.EncodeToHexString(mmrLeafEncoded)
	paraHeadEncoded, _ := types.EncodeToHexString(paraHead)
	log.WithFields(log.Fields{
		"input":                       string(b),
		"commitmentHash":              "0x" + hex.EncodeToString(commitmentHash[:]),
		"paraHeadEncoded":             paraHeadEncoded,
		"paraHeadLeaf":                "0x" + hex.EncodeToString(paraHeadLeaf),
		"paraHeadProofRootCalculated": "0x" + hex.EncodeToString(paraHeadProofRoot),
		"paraHeadProofRootMMRLeaf":    "0x" + hex.EncodeToString(mmrLeaf.ParachainHeads[:]),
		"mmrLeafOpaqueEncoded":        mmrLeafOpaqueEncoded,
	}).Info("Submitting tx")
	return nil
}

func (wr *EthereumChannelWriter) logIncentivizedTx(
	messages []incentivized.IncentivizedInboundChannelMessage,
	paraheadPartial incentivized.ParachainLightClientOwnParachainHeadPartial,
	paraHeadProof incentivized.ParachainLightClientParachainHeadProof,
	beefyMMRLeafPartial incentivized.ParachainLightClientBeefyMMRLeafPartial,
	beefyMMRLeafIndex int64, beefyLeafCount int64, beefyMMRProof [][32]byte,
	paraHead types.Header, paraHeadLeaf []byte, paraHeadProofRoot []byte, mmrLeaf types.MMRLeaf,
	commitmentHash types.H256,
) error {
	var incentivizedMessagesLog []IncentivizedInboundChannelMessageLog
	for _, item := range messages {
		incentivizedMessagesLog = append(incentivizedMessagesLog, IncentivizedInboundChannelMessageLog{
			Target:  item.Target,
			Nonce:   item.Nonce,
			Fee:     item.Fee,
			Payload: "0x" + hex.EncodeToString(item.Payload),
		})
	}

	var paraHeadProofString []string
	for _, item := range paraHeadProof.Proof {
		paraHeadProofString = append(paraHeadProofString, "0x"+hex.EncodeToString(item[:]))
	}
	var beefyMMRProofString []string
	for _, item := range beefyMMRProof {
		beefyMMRProofString = append(beefyMMRProofString, "0x"+hex.EncodeToString(item[:]))
	}
	input := &IncentivizedSubmitInput{
		Messages: incentivizedMessagesLog,
		ParaheadPartial: ParaheadPartialLog{
			ParentHash:     "0x" + hex.EncodeToString(paraheadPartial.ParentHash[:]),
			Number:         paraheadPartial.Number,
			StateRoot:      "0x" + hex.EncodeToString(paraheadPartial.StateRoot[:]),
			ExtrinsicsRoot: "0x" + hex.EncodeToString(paraheadPartial.ExtrinsicsRoot[:]),
		},
		ParaHeadProof: ParaHeadProofLog{
			Pos:   paraHeadProof.Pos,
			Width: paraHeadProof.Width,
			Proof: paraHeadProofString,
		},
		BeefyMMRLeafPartial: BeefyMMRLeafPartialLog{
			ParentNumber:         beefyMMRLeafPartial.ParentNumber,
			ParentHash:           "0x" + hex.EncodeToString(beefyMMRLeafPartial.ParentHash[:]),
			NextAuthoritySetId:   beefyMMRLeafPartial.NextAuthoritySetId,
			NextAuthoritySetLen:  beefyMMRLeafPartial.NextAuthoritySetLen,
			NextAuthoritySetRoot: "0x" + hex.EncodeToString(beefyMMRLeafPartial.NextAuthoritySetRoot[:]),
		},
		BeefyMMRLeafIndex: beefyMMRLeafIndex,
		BeefyLeafCount:    beefyLeafCount,
		BeefyMMRProof:     beefyMMRProofString,
	}
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	mmrLeafEncoded, _ := types.EncodeToBytes(mmrLeaf)
	mmrLeafOpaqueEncoded, _ := types.EncodeToHexString(mmrLeafEncoded)
	paraHeadEncoded, _ := types.EncodeToHexString(paraHead)
	log.WithFields(log.Fields{
		"input":                       string(b),
		"commitmentHash":              "0x" + hex.EncodeToString(commitmentHash[:]),
		"paraHeadEncoded":             paraHeadEncoded,
		"paraHeadLeaf":                "0x" + hex.EncodeToString(paraHeadLeaf),
		"paraHeadProofRootCalculated": "0x" + hex.EncodeToString(paraHeadProofRoot),
		"paraHeadProofRootMMRLeaf":    "0x" + hex.EncodeToString(mmrLeaf.ParachainHeads[:]),
		"mmrLeafOpaqueEncoded":        mmrLeafOpaqueEncoded,
	}).Info("Submitting tx")

	return nil
}
