package main

//go:generate abigen --sol ../../contracts/Betting.sol --pkg contract --out ./go-bindings/betting/betting.go

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	contract "betting-cli/go-bindings/betting"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


func connect() *ethclient.Client {
	backend, err := ethclient.Dial("ws://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	return backend
}

func (c Client) getAuth(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.backend.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.backend.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice
	return auth
}

func (c Client) deploy(auth *bind.TransactOpts) (common.Address, *types.Transaction, *contract.Betting, error) {
	return contract.DeployBetting(auth, c.backend,initialOutcome)
}

func (c Client) getBalance(address common.Address) *big.Int {
	balance, err := c.backend.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

// subscribe and listen to contract events
func (c Client) listenBettingEvents(address common.Address, quit chan struct{}) {
	errs := make(chan error, 1)
	logs := make(chan types.Log)

	sub, err := c.backend.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			errs <- err
		case evLog := <-logs:
			abi, err := contract.BettingMetaData.GetAbi()
			if err != nil {
				errs <- err
			}
			eventType, err := abi.EventByID(evLog.Topics[0])
			if err != nil {
				errs <- err
			}
			fmt.Println("Events:")
			switch eventType.Name {
			case "OracleChanged":
				event, err := c.bettingInstance.ParseOracleChanged(evLog)
				if err != nil {
					errs <- err
				}
				fmt.Printf("Previous oracle %v new oracle: %v\n", event.PreviousOracle.Hex(), event.NewOracle.Hex())
			case "BetMade":
				event, err := c.bettingInstance.ParseBetMade(evLog)
				if err != nil {
					errs <- err
				}
				fmt.Printf("Bet made: from %v , outcome %x , amount %v\n", event.Gambler.Hex(), event.Outcome, event.Amount)
			case "Withdrawn":
				event, err := c.bettingInstance.ParseWithdrawn(evLog)
				if err != nil {
					errs <- err
				}
				fmt.Printf("Withdraw of %v wei to: %v\n", event.Amount, event.Gambler.Hex())
			case "Winners":
				event, err := c.bettingInstance.ParseWinners(evLog)
				if err != nil {
					errs <- err
				}
				fmt.Printf("Total Prize for all winners: %v\n", event.TotalPrize)
				for _, winner := range event.Wins{
					fmt.Printf("Winner: %v\n", winner.Hex())
				}
				
			}
		case err := <-errs:
			log.Fatal(err)
		case <-quit:
			return
		}
	}
}

func isZeroAddress(address common.Address) bool {
	return reflect.DeepEqual(address.Bytes(), common.FromHex("0x0000000000000000000000000000000000000000"))
}

func getPrivateKey(privateKeyHex string) *ecdsa.PrivateKey{
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func stringToKeccak256(s string) [32]byte {
	var output [32]byte
	copy(output[:], crypto.Keccak256([]byte(s))[:])
	return output
}

func main() {
	var cmd int
	var wg sync.WaitGroup
	addMapKey:=addressMapKey(address,PKHEX)

	client := &Client{}
	client.backend = connect()
	defer client.backend.Close()

	ownerPrivKey:=getPrivateKey(defaultPKHex)

	client.scanner = bufio.NewScanner(os.Stdin)
	quit := make(chan struct{})
	for {
		fmt.Println("------------------\nChoose a command:")
		for i, c := range commands {
			fmt.Printf("(%v) %s\n", i+1, c)
		}
		fmt.Scanln(&cmd)
		fmt.Println("------------------")
		switch cmd {
		case 1:
			if !isZeroAddress(client.bettingAddress) {
				fmt.Printf("Contract already deployed at: %v\n", client.bettingAddress)
				continue
			}
			wg.Add(1)
			auth := client.getAuth(ownerPrivKey)
			address, _, instance, err := client.deploy(auth)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			client.bettingAddress = address
			client.bettingInstance = instance
			fmt.Println("Contract deployed at:", address.Hex())

			//listen to all contract events
			go func() {
				defer wg.Done()
				client.listenBettingEvents(address, quit)
			}()

		case 2:
			fmt.Println("Enter the address")
			client.scanner.Scan()
			account := common.HexToAddress(client.scanner.Text())
			balance := client.getBalance(account)
			fmt.Println("------------------")
			fmt.Printf("Balance of account %s : %v\n", account.Hex(), balance)
		//Choose Oracle
		case 3:
			fmt.Println("Enter number(1/2/3) to choose oracle:")
			client.scanner.Scan()
			number, err := strconv.Atoi(client.scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			oracle := common.HexToAddress(oracleAddress[number-1])
			auth := client.getAuth(ownerPrivKey)
			tx,err:=client.bettingInstance.ChooseOracle(&bind.TransactOpts{From: auth.From, Signer: auth.Signer}, oracle)
			fmt.Println("The address of the chosen Oracle is: ",oracleAddress[number-1])
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		//Set Outcomes
		case 4:
			fmt.Println("Enter the outcomes you want to set(seperate with comma):")
			client.scanner.Scan()
			if err := client.scanner.Err(); err != nil {
				fmt.Println("Error reading text:", err)
				continue
			}
			outcomes := [][32]byte{}
			s := strings.Split(client.scanner.Text(), ",")
			for _, v := range s {
				outcomes = append(outcomes, stringToKeccak256(v))
			}

			auth := client.getAuth(ownerPrivKey)
			tx,err:=client.bettingInstance.SetOutcomes(&bind.TransactOpts{From: auth.From, Signer: auth.Signer},outcomes)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		//Contract Reset
		case 5:
			auth := client.getAuth(ownerPrivKey)
			tx,err:=client.bettingInstance.ContractReset(&bind.TransactOpts{From: auth.From, Signer: auth.Signer})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		//make bet
		case 6:
			fmt.Println("Enter number(1~6) to choose gambler:")
			client.scanner.Scan()
			number, err := strconv.Atoi(client.scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Enter the outcome(in hex):")
			client.scanner.Scan()
			outcome64:=common.Hex2Bytes(client.scanner.Text())
			var outcome [32]byte
			copy(outcome[:], outcome64)
			fmt.Println("Enter the amount:")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			privateKey:=getPrivateKey(gamblerPKHEX[number-1])
			auth := client.getAuth(privateKey)
			tx, err :=client.bettingInstance.MakeBet(&bind.TransactOpts{From: auth.From, Signer: auth.Signer, Value: amount}, outcome)

			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		//Get Outcomes
		case 7:
			outcomesBytes, err :=client.bettingInstance.GetOutcomes(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			for _, outcome := range outcomesBytes{		
				fmt.Printf("Valid Outcomes: %x \n", outcome)
			}
		//Make Decision
		case 8:
			fmt.Println("Enter the address(only for oracle):\n")
			client.scanner.Scan()
			privateKey:=getPrivateKey(addMapKey[client.scanner.Text()])
			auth := client.getAuth(privateKey)
			fmt.Println("Enter the decided outcome(in hex):\n")
			client.scanner.Scan()
			outcome64:=common.Hex2Bytes(client.scanner.Text())
			var outcome [32]byte
			copy(outcome[:], outcome64)
			tx,err:=client.bettingInstance.MakeDecision(&bind.TransactOpts{From: auth.From, Signer: auth.Signer},outcome)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		//Get Gamblers
		case 9:
			addresses, err :=client.bettingInstance.GetGamblers(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			for _, address := range addresses{		
				fmt.Printf("Gambler:%v \n", address.Hex())
			}

		//Get Winners
		case 10:
			addresses, err :=client.bettingInstance.GetWinners(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			for _, address := range addresses{		
				fmt.Printf("Winner:%v \n", address.Hex())
			}

		//check Winnings
		case 11:
			winning, err :=client.bettingInstance.CheckWinnings(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("winning: %v\n", winning)

		//withdraw
		case 12:
			fmt.Println("Enter the address(only for winner):\n")
			client.scanner.Scan()
			privateKey:=getPrivateKey(addMapKey[client.scanner.Text()])
			auth := client.getAuth(privateKey)
			fmt.Println("Enter the amount you want to withdraw:\n")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			tx,err:=client.bettingInstance.Withdraw(&bind.TransactOpts{From: auth.From, Signer: auth.Signer},amount)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())

		case 13:
			quit <- struct{}{}
			wg.Wait()
			return
		}
	}
}
