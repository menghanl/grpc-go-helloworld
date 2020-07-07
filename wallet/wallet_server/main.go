package main

import (
	"fmt"

	walletpb "google.golang.org/grpc/grpc-wallet/grpc/examples/wallet"
)

func main() {
	fmt.Printf("hello wallet, %+v\n", walletpb.BalancePerAddress{Address: "abc", Balance: 123})
}
