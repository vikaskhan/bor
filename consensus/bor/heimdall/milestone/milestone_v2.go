package milestone

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// MilestoneV2 defines a response object type of bor milestone
type MilestoneV2 struct {
	Proposer        common.Address `json:"proposer"`
	StartBlock      uint64         `json:"start_block"`
	EndBlock        uint64         `json:"end_block"`
	Hash            common.Hash    `json:"hash"`
	BorChainID      string         `json:"bor_chain_id"`
	MilestoneID     string         `json:"milestone_id"`
	Timestamp       uint64         `json:"timestamp"`
	TotalDifficulty uint64         `json:"total_difficulty"`
}

func (m *MilestoneV2) UnmarshalJSON(data []byte) error {
	type Alias MilestoneV2
	temp := &struct {
		StartBlock      string `json:"start_block"`
		EndBlock        string `json:"end_block"`
		Hash            string `json:"hash"`
		Timestamp       string `json:"timestamp"`
		TotalDifficulty string `json:"total_difficulty"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	startBlock, err := strconv.ParseUint(temp.StartBlock, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid start_block: %w", err)
	}
	m.StartBlock = startBlock

	endBlock, err := strconv.ParseUint(temp.EndBlock, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid end_block: %w", err)
	}
	m.EndBlock = endBlock

	decodedHash, err := base64.StdEncoding.DecodeString(temp.Hash)
	if err != nil {
		return fmt.Errorf("failed to decode hash: %w", err)
	}
	m.Hash = common.BytesToHash(decodedHash)

	timestamp, err := strconv.ParseUint(temp.Timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp: %w", err)
	}
	m.Timestamp = timestamp

	totalDifficulty, err := strconv.ParseUint(temp.TotalDifficulty, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid total_difficulty: %w", err)
	}
	m.TotalDifficulty = totalDifficulty

	return nil
}

type MilestoneResponseV2 struct {
	Result MilestoneV2 `json:"milestone"`
}

func (m *MilestoneCountResponseV2) UnmarshalJSON(data []byte) error {
	temp := &struct {
		Count string `json:"count"`
	}{}

	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	count, err := strconv.ParseInt(temp.Count, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid count: %w", err)
	}
	m.Count = count

	return nil
}

type MilestoneCountResponseV2 struct {
	Count int64 `json:"count"`
}
