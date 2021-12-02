package keeper

import (
	"context"

	"github.com/OmniFlix/payment-stream/x/paymentstream/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PaymentStreamsAll(c context.Context, req *types.QueryAllPaymentStreamRequest) (*types.QueryAllPaymentStreamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var paymentstreams []types.PaymentStream
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	paymentstreamStore := prefix.NewStore(store, types.KeyPrefix(""))

	pageRes, err := query.Paginate(paymentstreamStore, req.Pagination, func(key []byte, value []byte) error {
		var paymentstream types.PaymentStream
		if err := k.cdc.Unmarshal(value, &paymentstream); err != nil {
			return err
		}

		paymentstreams = append(paymentstreams, paymentstream)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPaymentStreamResponse{PaymentStreams: paymentstreams, Pagination: pageRes}, nil
}
