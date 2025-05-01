# fee-sink-rewards

This code determines the recipient of TERABRO's node reward. It uses the block hash of the fee sink transaction as the seed for the random picker, ensuring that the reward distribution is both fair and verifiable.

## Verifying a Payout

Anyone can verify the payout on-chain using the script with the following command:

```shell
go run cmd/run_local/main.go
```

Simply change the round number in `cmd/run_local/main.go` to verify a different round.