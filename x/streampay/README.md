# Streampay
Streampay module is used to create stream payments on cosmos-sdk based chains

Anyone can create a stream payment to a recipient address with a specified amount with type of stream and end time.
receiver can claim the amount after the end time or can claim streamed amount at anytime if payment type is continuous or periodic.

Types of steam payments:
1. Continuous
2. Delayed
3. Periodic

**Continuous Stream Payment**
  Amount will be streamed continuously based on time and can be claimed at any time.
**Delayed Stream Payment**
  Amount can be claimed only at the end of end time (works as scheduled payment)
**Periodic Stream Payment**
  Amount will be streamed according to stream periods

_Cancellable Stream Payments_
 streampay module supports cancellable streaming payments also cancellable streams can be cancelled before it's endtime
 if stream is in progress remaining amount will be returned to stream payment creator

## Stream Payment
```protobuf
message StreamPayment {
  string id = 1;
  string sender = 2;
  string recipient = 3;
  cosmos.base.v1beta1.Coin total_amount = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags) = "yaml:\"total_amount\""
  ];
  StreamType stream_type = 5 [(gogoproto.moretags) = "yaml:\"stream_type\""];
  repeated Period periods = 6 [(gogoproto.nullable) = true];
  bool cancellable = 7;
  google.protobuf.Timestamp start_time = 8 [
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true,
    (gogoproto.moretags) = "yaml:\"start_time\""
  ];
  google.protobuf.Timestamp end_time = 9 [
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true,
    (gogoproto.moretags) = "yaml:\"end_time\""
  ];
  cosmos.base.v1beta1.Coin streamed_amount = 10 [
    (gogoproto.nullable) = false,
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin",
    (gogoproto.moretags) = "yaml:\"streamed_amount\""
  ];
  google.protobuf.Timestamp last_claimed_at = 11 [
    (gogoproto.nullable) = false,
    (gogoproto.stdtime) = true,
    (gogoproto.moretags) = "yaml:\"last_claimed_at\""
  ];
  StreamStatus status = 12;
}
```

## State
The state of the module is expressed by following fields:

- `1.StreamPayments:` list of stream payments
- `2.Params:` streampay module params

```protobuf
message GenesisState {
  repeated StreamPayment stream_payments      = 1 [(gogoproto.nullable) = false];
  uint64                 next_stream_payment_number = 2;
  Params params                 = 3 [(gogoproto.nullable) = false];
}
```

## Messages 

### Stream Send
`MsgStreamSend` can be submitted by any account to create a stream payment 
```protobuf
message MsgStreamSend {
  string sender = 1;
  string recipient = 2;
  cosmos.base.v1beta1.Coin amount = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin"
  ];
  google.protobuf.Duration duration = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.stdduration) = true
  ];
  StreamType stream_type = 5 [(gogoproto.moretags) = "yaml:\"stream_type\""];
  repeated Period periods = 6 [(gogoproto.nullable) = true];
  bool cancellable = 7;
}
```
### Stop Stream / Cancel Stream
`MsgStopStream` can be submitted by stream-payment creator to stop the stream payment.
```protobuf
message MsgStopStream {
  string stream_id = 1 [(gogoproto.moretags) = "yaml:\"stream_id\""];
  string sender = 2;
}
```

### Claim Streamed Amount
`MsgClaimStreamedAmount` can be submitted by claimer to claim amount from stream payment.
```protobuf
message MsgClaimStreamedAmount {
  string stream_id = 1 [(gogoproto.moretags) = "yaml:\"stream_id\""];
  string claimer = 2;
}
```
## Queries
```protobuf
service Query {
  rpc StreamingPayments(QueryStreamPaymentsRequest) returns (QueryStreamPaymentsResponse) {
    option (google.api.http).get = "/omniflix/streampay/v1/streaming-payments";
  }
  rpc StreamingPayment(QueryStreamPaymentRequest) returns (QueryStreamPaymentResponse) {
    option (google.api.http).get = "/omniflix/streampay/v1/streaming-payments/{id}";
  }
  rpc Params(QueryParamsRequest) returns (QueryParamsResponse) {
    option (google.api.http).get = "/omniflix/streampay/v1/params";
  }
}
```

## CLI Transactions

### Stream Send
Continuous payment
```bash
 streampayd tx streampay stream-send [recipient] [amount] \
 --end-time <end-timestamp> \
 --chain-id <chain-id> \
 --from <sender> \
 --fees <fees>
```

Delayed payment:
```bash
streampayd tx streampay stream-send [recipient] [amount] \
--end-time <end-timestamp> \
--delayed \
--chain-id <chain-id> \
--from <sender> \
--fees <fees>
```
Periodic payment:
```shell
streampayd tx streampay stream-send [recipient] [amount] \
--end-time <end-timestamp> \
--stream-periods-file <stream-periods-file> \
--chain-id <chain-id> \
--from <sender> \
--fees <fees>
```

### Stop Stream
```bash
streampayd tx streampay stop-stream [stream-id] --chain-id <chain-id> --from <sender> --fees <fees>
```
### Claim Streamed Amount
```bash
streampayd tx streampay claim [stream-id] --chain-id <chain-id> --from <sender> --fees <fees>
```

## CLI Queries
1. Query all streaming payments
```bash
streampayd q streampay stream-payments
```
2. Query stream payment by it's id
```bash
streampayd q streampay stream-payment [stream-payment-id]
```
3. Query streampay module params
```bash
streampayd q streampay params
```
