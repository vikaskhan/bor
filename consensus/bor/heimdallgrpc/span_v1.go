package heimdallgrpc

import (
	"context"

	"github.com/ethereum/go-ethereum/consensus/bor/heimdall/span"
	"github.com/ethereum/go-ethereum/consensus/bor/valset"
	"github.com/ethereum/go-ethereum/log"

	proto "github.com/maticnetwork/polyproto/heimdall"
	protoutils "github.com/maticnetwork/polyproto/utils"
)

func (h *HeimdallGRPCClient) GetSpanV1(ctx context.Context, spanID uint64) (*span.HeimdallSpan, error) {
	req := &proto.SpanRequest{
		ID: spanID,
	}

	log.Info("Fetching span", "spanID", spanID)

	ctxWithTimeout, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	res, err := h.client.Span(ctxWithTimeout, req)
	if err != nil {
		return nil, err
	}

	log.Info("Fetched span", "spanID", spanID)

	return parseSpan(res.Result), nil
}

func (h *HeimdallGRPCClient) GetLatestSpanV1(ctx context.Context) (*span.HeimdallSpan, error) {
	req := &proto.LatestSpanRequest{}

	log.Info("Fetching latest span")

	ctxWithTimeout, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	res, err := h.client.LatestSpan(ctxWithTimeout, req)
	if err != nil {
		return nil, err
	}

	log.Info("Fetched latest span")

	return parseSpan(res.Result), nil
}

func parseSpan(protoSpan *proto.Span) *span.HeimdallSpan {
	resp := &span.HeimdallSpan{
		Span: span.Span{
			Id:         protoSpan.ID,
			StartBlock: protoSpan.StartBlock,
			EndBlock:   protoSpan.EndBlock,
		},
		ValidatorSet:      valset.ValidatorSet{},
		SelectedProducers: []valset.Validator{},
		ChainID:           protoSpan.ChainID,
	}

	for _, validator := range protoSpan.ValidatorSet.Validators {
		resp.ValidatorSet.Validators = append(resp.ValidatorSet.Validators, parseValidator(validator))
	}

	resp.ValidatorSet.Proposer = parseValidator(protoSpan.ValidatorSet.Proposer)

	for _, validator := range protoSpan.SelectedProducers {
		resp.SelectedProducers = append(resp.SelectedProducers, *parseValidator(validator))
	}

	return resp
}

func parseValidator(validator *proto.Validator) *valset.Validator {
	return &valset.Validator{
		ID:               validator.ID,
		Address:          protoutils.ConvertH160toAddress(validator.Address),
		VotingPower:      validator.VotingPower,
		ProposerPriority: validator.ProposerPriority,
	}
}
