package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func CreateSafeProposalCmd() *cobra.Command {
	proposalCmd := &cobra.Command{
		Use:   "proposal",
		Short: "Manage proposals for a Safe",
		Long:  `Manage Safe proposals — create new ones, list existing ones, approve pending ones, and execute approved proposals.`,
	}

	proposalCmd.AddCommand(createListProposalsCmd())
	proposalCmd.SetOut(os.Stdout)

	return proposalCmd
}

func createListProposalsCmd() *cobra.Command {
	var safe string
	listProposalsCmd := &cobra.Command{
		Use:   "list",
		Short: "List proposals for a safe",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !common.IsHexAddress(safe) {
				return fmt.Errorf("invalid safe address: %s", safe)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := ethclient.Dial(rpcURL)
			if err != nil {
				return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
			}

			chainID, err := client.ChainID(context.Background())
			if err != nil {
				return fmt.Errorf("failed to get chain ID: %v", err)
			}
			if safeAPIURL == "" {
				safeAPIURL = fmt.Sprintf("https://safe-client.safe.global/v1/chains/%s/safes/%s/transactions/history/", chainID.String(), safe)
				fmt.Println("safe-api is not set, using default: ", safeAPIURL)
			}

			proposals, err := GetProposals(safeAPIURL)
			if err != nil {
				return fmt.Errorf("error retrieving delegates: %v", err)
			}
			if len(proposals) == 0 {
				return fmt.Errorf("no proposals found")
			} else {
				var count int64
				for _, d := range proposals {
					if d.Type != "TRANSACTION" {
						continue
					}
					count += 1
					cmd.Println("------------------------------------------------------------------------------------------------")
					cmd.Printf("Type: %s\n", d.Type)
					cmd.Printf("Proposal Count #%d:\n", count)

					if d.Transaction != nil {
						cmd.Printf("  Transaction ID: %s\n", d.Transaction.ID)
						cmd.Printf("  Transaction Hash: %s\n", d.Transaction.TxHash)
						cmd.Printf("  Transaction Status: %s\n", d.Transaction.TxStatus)
						cmd.Printf("  Transaction Timestamp: %d\n", d.Transaction.Timestamp)

						cmd.Println("\n  Transaction Info:")
						cmd.Printf("    Type: %s\n", d.Transaction.TxInfo.Type)
						cmd.Printf("    Description: %s\n", d.Transaction.TxInfo.HumanDescription)
						cmd.Printf("    Creator: %s (%s)\n", d.Transaction.TxInfo.Creator.Name, d.Transaction.TxInfo.Creator.Value)
						cmd.Printf("    Creator Logo URI: %s\n", d.Transaction.TxInfo.Creator.LogoUri)
						cmd.Printf("    Implementation: %s (%s)\n", d.Transaction.TxInfo.Implementation.Name, d.Transaction.TxInfo.Implementation.Value)
						cmd.Printf("    Implementation Logo URI: %s\n", d.Transaction.TxInfo.Implementation.LogoUri)
						cmd.Printf("    Factory: %s (%s)\n", d.Transaction.TxInfo.Factory.Name, d.Transaction.TxInfo.Factory.Value)
						cmd.Printf("    Factory Logo URI: %s\n", d.Transaction.TxInfo.Factory.LogoUri)
						cmd.Printf("    Salt Nonce: %s\n", d.Transaction.TxInfo.SaltNonce)

						cmd.Println("\n  Execution Info:")
						cmd.Printf("    Type: %s\n", d.Transaction.ExecutionInfo.Type)
						cmd.Printf("    Nonce: %d\n", d.Transaction.ExecutionInfo.Nonce)
						cmd.Printf("    Required Confirmations: %d/%d\n",
							d.Transaction.ExecutionInfo.ConfirmationsSubmitted,
							d.Transaction.ExecutionInfo.ConfirmationsRequired)

						if len(d.Transaction.ExecutionInfo.MissingSigners) > 0 {
							cmd.Println("    Missing Signers:")
							for _, signer := range d.Transaction.ExecutionInfo.MissingSigners {
								cmd.Printf("      - Name: %s\n", signer.Name)
								cmd.Printf("        Value: %s\n", signer.Value)
								cmd.Printf("        LogoUri: %s\n", signer.LogoUri)
							}
						}

						if d.Transaction.SafeAppInfo.Name != "" {
							cmd.Println("\n  Safe App Info:")
							cmd.Printf("    Name: %s\n", d.Transaction.SafeAppInfo.Name)
							cmd.Printf("    URL: %s\n", d.Transaction.SafeAppInfo.URL)
							cmd.Printf("    LogoUri: %s\n", d.Transaction.SafeAppInfo.LogoUri)
						}
					} else {
						cmd.Printf("  Timestamp: %d\n", d.Timestamp)
					}

					cmd.Printf("  Conflict Type: %s\n", d.ConflictType)
					cmd.Println("------------------------------------------------------------------------------------------------")
				}

			}
			return nil
		},
	}
	listProposalsCmd.Flags().StringVar(&safe, "safe", "", "Safe address")
	listProposalsCmd.Flags().StringVar(&safeAPIURL, "safe-api", "", "Override default Safe API URL")
	listProposalsCmd.Flags().StringVar(&rpcURL, "rpc", "", "RPC URL to retrieve chain ID")
	listProposalsCmd.MarkFlagRequired("safe")
	listProposalsCmd.MarkFlagRequired("rpc")

	return listProposalsCmd
}
