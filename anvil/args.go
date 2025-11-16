package anvil

import (
	"fmt"
	"strings"
)

func getArgs(c *Config) []string {
	args := []string{}

	if c.Accounts != 0 { // default
		args = append(args, "--accounts", fmt.Sprint(c.Accounts))
	}

	if c.BlockTime != 0 {
		args = append(args, "--block-time", fmt.Sprint(c.BlockTime))
	}

	if c.Balance != "" {
		args = append(args, "--balance", c.Balance)
	}

	if c.ConfigOut != "" {
		args = append(args, "--config-out", c.ConfigOut)
	}

	if c.DerivationPath != "" {
		args = append(args, "--derivation-path", c.DerivationPath)
	}

	if c.DumpStatePath != "" {
		args = append(args, "--dump-state", c.DumpStatePath)
	}

	if c.Hardfork != "" {
		args = append(args, "--hardfork", c.Hardfork)
	}

	if c.InitPath != "" {
		args = append(args, "--init", c.InitPath)
	}

	if c.IPC {
		if c.IPCPath == "" {
			args = append(args, "--ipc")
		} else {
			args = append(args, "--ipc", c.IPCPath)
		}
	}

	if c.Threads != 0 {
		args = append(args, "--threads", fmt.Sprint(c.Threads))
	}

	if c.LoadStatePath != "" {
		args = append(args, "--load-state", c.LoadStatePath)
	}

	if c.Mnemonic != "" {
		args = append(args, "--mnemonic", c.Mnemonic)
	}

	if c.MaxPersistedStates != 0 {
		args = append(args, "--max-persisted-states", fmt.Sprint(c.MaxPersistedStates))
	}

	if c.MixedMining {
		args = append(args, "--mixed-mining")
	}

	if c.MnemonicRandom {
		if c.MnemonicRandomWords == 0 {
			args = append(args, "--mnemonic-random")
		} else {
			args = append(args, "--mnemonic-random", fmt.Sprint(c.MnemonicRandomWords))
		}
	}

	if c.MnemonicSeedUnsafe != "" {
		args = append(args, "--mnemonic-seed-unsafe", c.MnemonicSeedUnsafe)
	}

	if c.NoMining {
		args = append(args, "--no-mining")
	}

	if c.Number != 0 {
		args = append(args, "--number", fmt.Sprint(c.Number))
	}

	if c.Order != "" && c.Order != "fees" {
		args = append(args, "--order", c.Order)
	}

	if c.Port != 0 {
		args = append(args, "--port", fmt.Sprint(c.Port))
	}

	if c.PreserveHistoricalStates {
		args = append(args, "--preserve-historical-states")
	}

	if c.PruneHistory {
		if c.PruneHistoryStates == 0 {
			args = append(args, "--prune-history")
		} else {
			args = append(args, "--prune-history", fmt.Sprint(c.PruneHistoryStates))
		}
	}

	if c.StateInterval != 0 {
		args = append(args, "--state-interval", fmt.Sprint(c.StateInterval))
	}

	if c.SlotsInEpoch != 0 {
		args = append(args, "--slots-in-an-epoch", fmt.Sprint(c.SlotsInEpoch))
	}

	if c.StatePath != "" {
		args = append(args, "--state", c.StatePath)
	}

	if c.Timestamp != 0 {
		args = append(args, "--timestamp", fmt.Sprint(c.Timestamp))
	}

	if c.TransactionBlockKeeper != 0 {
		args = append(args, "--transaction-block-keeper", fmt.Sprint(c.TransactionBlockKeeper))
	}

	if c.Color != "" {
		args = append(args, "--color", c.Color)
	}

	if c.JSONLogs {
		args = append(args, "--json")
	}

	if c.MarkdownLogs {
		args = append(args, "--md")
	}

	if c.Quiet {
		args = append(args, "--quiet")
	}

	if c.Verbosity > 0 {
		vv := "-" + strings.Repeat("v", c.Verbosity)
		args = append(args, vv)
	}

	if c.AllowOrigin != "" {
		args = append(args, "--allow-origin", c.AllowOrigin)
	}

	if c.CachePath != "" {
		args = append(args, "--cache-path", c.CachePath)
	}

	if c.Host != "" {
		args = append(args, "--host", c.Host)
	}

	if c.NoCors {
		args = append(args, "--no-cors")
	}

	if c.NoRequestSizeLimit {
		args = append(args, "--no-request-size-limit")
	}

	if c.ComputeUnitsPerSecond != 0 {
		args = append(args, "--compute-units-per-second", fmt.Sprint(c.ComputeUnitsPerSecond))
	}

	if c.ForkURL != "" {
		args = append(args, "--fork-url", c.ForkURL)
	}

	if c.ForkBlockNumber != 0 {
		args = append(args, "--fork-block-number", fmt.Sprint(c.ForkBlockNumber))
	}

	if c.ForkChainID != 0 {
		args = append(args, "--fork-chain-id", fmt.Sprint(c.ForkChainID))
	}

	if len(c.ForkHeaders) > 0 {
		args = append(args, "--fork-header")
		for _, h := range c.ForkHeaders {
			args = append(args, h)
		}
	}

	if c.ForkRetryBackoff != "" {
		args = append(args, "--fork-retry-backoff", c.ForkRetryBackoff)
	}

	if c.ForkTransactionHash != "" {
		args = append(args, "--fork-transaction-hash", c.ForkTransactionHash)
	}

	if c.NoRateLimit {
		args = append(args, "--no-rate-limit")
	}

	if c.NoStorageCaching {
		args = append(args, "--no-storage-caching")
	}

	if c.Retries != 0 {
		args = append(args, "--retries", fmt.Sprint(c.Retries))
	}

	if c.Timeout != "" {
		args = append(args, "--timeout", c.Timeout)
	}

	if c.BlockBaseFeePerGas != "" {
		args = append(args, "--block-base-fee-per-gas", c.BlockBaseFeePerGas)
	}

	if c.ChainID != 0 {
		args = append(args, "--chain-id", fmt.Sprint(c.ChainID))
	}

	if c.CodeSizeLimit != 0 {
		args = append(args, "--code-size-limit", fmt.Sprintf("0x%x", c.CodeSizeLimit))
	}

	if c.DisableBlockGasLimit {
		args = append(args, "--disable-block-gas-limit")
	}

	if c.DisableCodeSizeLimit {
		args = append(args, "--disable-code-size-limit")
	}

	if c.DisableMinPriorityFee {
		args = append(args, "--disable-min-priority-fee")
	}

	if c.GasLimit != 0 {
		args = append(args, "--gas-limit", fmt.Sprint(c.GasLimit))
	}

	if c.GasPrice != "" {
		args = append(args, "--gas-price", c.GasPrice)
	}

	if c.AutoImpersonate {
		args = append(args, "--auto-impersonate")
	}

	if c.DisableConsoleLog {
		args = append(args, "--disable-console-log")
	}

	if c.DisableDefaultCreate2Deployer {
		args = append(args, "--disable-default-create2-deployer")
	}

	if c.DisablePoolBalanceChecks {
		args = append(args, "--disable-pool-balance-checks")
	}

	if c.MemoryLimit != 0 {
		args = append(args, "--memory-limit", fmt.Sprint(c.MemoryLimit))
	}

	if c.PrintTraces {
		args = append(args, "--print-traces")
	}

	if c.StepsTracing {
		args = append(args, "--steps-tracing")
	}

	if c.Celo {
		args = append(args, "--celo")
	}

	if c.Optimism {
		args = append(args, "--optimism")
	}

	return args
}
