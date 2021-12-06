package keeper

import (
	"context"
	"fmt"
	gogotypes "github.com/gogo/protobuf/types"

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

func (k Keeper) PaymentStreams(c context.Context, req *types.QueryPaymentStreamsRequest) (*types.QueryPaymentStreamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var streams []types.PaymentStream
	var pageRes *query.PageResponse
	var err error
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	if req.Status == types.StatusOpen.String() {
		activePaymentStreamsStore := prefix.NewStore(store, types.ActivePaymentPrefix(""))
		pageRes, err = query.Paginate(activePaymentStreamsStore, req.Pagination, func(key []byte, value []byte) error {
			var pId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &pId)
			stream, found := k.GetPaymentStream(ctx, pId.Value)
			if found {
				streams = append(streams, stream)
			}
			return nil
		})
	} else if req.Status == types.StatusCompleted.String() {
		inActivePaymentStreamsStore := prefix.NewStore(store, types.InActivePaymentPrefix(""))
		pageRes, err = query.Paginate(inActivePaymentStreamsStore, req.Pagination, func(key []byte, value []byte) error {
			var pId gogotypes.StringValue
			k.cdc.MustUnmarshal(value, &pId)
			stream, found := k.GetPaymentStream(ctx, pId.Value)
			if found {
				streams = append(streams, stream)
			}
			return nil
		})
	} else {
		paymentstreamsStore := prefix.NewStore(store, types.KeyPrefix(""))

		pageRes, err = query.Paginate(paymentstreamsStore, req.Pagination, func(key []byte, value []byte) error {
			var stream types.PaymentStream
			if err := k.cdc.Unmarshal(value, &stream); err != nil {
				return err
			}
			streams = append(streams, stream)
			return nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPaymentStreamsResponse{PaymentStreams: streams, Pagination: pageRes}, nil
}

func (k Keeper) SenderPaymentStreams(c context.Context, req *types.QuerySenderPaymentStreamsRequest) (*types.QueryPaymentStreamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var streams []types.PaymentStream
	var pageRes *query.PageResponse
	var err error
	var sender sdk.AccAddress

	if len(req.Sender) > 0 {
		sender, err = sdk.AccAddressFromBech32(req.Sender)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid sender address (%s)", err))
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("empty sender address is not allowed"))
	}
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	SenderPaymentStreamsStore := prefix.NewStore(store, types.KeySenderPaymentStream(sender.String(), ""))
	pageRes, err = query.Paginate(SenderPaymentStreamsStore, req.Pagination, func(key []byte, value []byte) error {
		var pId gogotypes.StringValue
		k.cdc.MustUnmarshal(value, &pId)
		stream, found := k.GetPaymentStream(ctx, pId.Value)
		if found {
			streams = append(streams, stream)
		}
		return nil
	})

	return &types.QueryPaymentStreamsResponse{PaymentStreams: streams, Pagination: pageRes}, nil
}

func (k Keeper) RecipientPaymentStreams(c context.Context, req *types.QueryRecipientPaymentStreamsRequest) (*types.QueryPaymentStreamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var streams []types.PaymentStream
	var pageRes *query.PageResponse
	var err error
	var recipient sdk.AccAddress

	if len(req.Recipient) > 0 {
		recipient, err = sdk.AccAddressFromBech32(req.Recipient)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid sender address (%s)", err))
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("empty recipient address is not allowed"))
	}
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	recipientPaymentStreamsStore := prefix.NewStore(store, types.KeyRecipientPaymentStream(recipient.String(), ""))
	pageRes, err = query.Paginate(recipientPaymentStreamsStore, req.Pagination, func(key []byte, value []byte) error {
		var pId gogotypes.StringValue
		k.cdc.MustUnmarshal(value, &pId)
		stream, found := k.GetPaymentStream(ctx, pId.Value)
		if found {
			streams = append(streams, stream)
		}
		return nil
	})

	return &types.QueryPaymentStreamsResponse{PaymentStreams: streams, Pagination: pageRes}, nil
}
