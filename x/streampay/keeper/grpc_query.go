package keeper

import (
	"context"

	"github.com/OmniFlix/streampay/v2/x/streampay/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StreamingPayments(c context.Context,
	req *types.QueryStreamPaymentsRequest,
) (*types.QueryStreamPaymentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	var filteredStreamPayments []types.StreamPayment
	var pageRes *query.PageResponse
	store := ctx.KVStore(k.storeKey)
	streamPaymentsStore := prefix.NewStore(store, types.PrefixPaymentStreamId)
	pageRes, err := query.FilteredPaginate(streamPaymentsStore,
		req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
			var streamPayment types.StreamPayment
			k.cdc.MustUnmarshal(value, &streamPayment)
			matchSender, matchRecipient := true, true

			// match sender address (if supplied)
			if len(req.Sender) > 0 {
				sender, err := sdk.AccAddressFromBech32(req.Sender)
				if err != nil {
					return false, err
				}

				matchSender = streamPayment.Sender == sender.String()
			}
			// match sender address (if supplied)
			if len(req.Recipient) > 0 {
				recipient, err := sdk.AccAddressFromBech32(req.Recipient)
				if err != nil {
					return false, err
				}

				matchRecipient = streamPayment.Recipient == recipient.String()
			}

			if matchSender && matchRecipient {
				if accumulate {
					filteredStreamPayments = append(filteredStreamPayments, streamPayment)
				}

				return true, nil
			}

			return false, nil
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryStreamPaymentsResponse{StreamPayments: filteredStreamPayments, Pagination: pageRes}, nil
}

// StreamingPayment returns details of the stream payment
func (k Keeper) StreamingPayment(goCtx context.Context,
	req *types.QueryStreamPaymentRequest,
) (*types.QueryStreamPaymentResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	streamPayment, found := k.GetStreamPayment(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "stream payment %s not found", req.Id)
	}

	return &types.QueryStreamPaymentResponse{StreamPayment: &streamPayment}, nil
}

func (k Keeper) Params(goCtx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}
