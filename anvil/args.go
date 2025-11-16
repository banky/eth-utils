package anvil

import (
	"fmt"
	"strings"
)

func getArgs(c *Config) []string {
	args := []string{}

	if c.accounts != 0 { // default
		args = append(args, "--accounts", fmt.Sprint(c.accounts))
	}

	if c.blockTime != 0 {
		args = append(args, "--block-time", fmt.Sprint(c.blockTime))
	}

	if c.balance != "" {
		args = append(args, "--balance", c.balance)
	}

	if c.configOut != "" {
		args = append(args, "--config-out", c.configOut)
	}

	if c.derivationPath != "" {
		args = append(args, "--derivation-path", c.derivationPath)
	}

	if c.dumpStatePath != "" {
		args = append(args, "--dump-state", c.dumpStatePath)
	}

	if c.hardfork != "" {
		args = append(args, "--hardfork", c.hardfork)
	}

	if c.initPath != "" {
		args = append(args, "--init", c.initPath)
	}

	if c.ipc {
		if c.ipcPath == "" {
			args = append(args, "--ipc")
		} else {
			args = append(args, "--ipc", c.ipcPath)
		}
	}

	if c.threads != 0 {
		args = append(args, "--threads", fmt.Sprint(c.threads))
	}

	if c.loadStatePath != "" {
		args = append(args, "--load-state", c.loadStatePath)
	}

	if c.mnemonic != "" {
		args = append(args, "--mnemonic", c.mnemonic)
	}

	if c.maxPersistedStates != 0 {
		args = append(args, "--max-persisted-states", fmt.Sprint(c.maxPersistedStates))
	}

	if c.mixedMining {
		args = append(args, "--mixed-mining")
	}

	if c.mnemonicRandom {
		if c.mnemonicRandomWords == 0 {
			args = append(args, "--mnemonic-random")
		} else {
			args = append(args, "--mnemonic-random", fmt.Sprint(c.mnemonicRandomWords))
		}
	}

	if c.mnemonicSeedUnsafe != "" {
		args = append(args, "--mnemonic-seed-unsafe", c.mnemonicSeedUnsafe)
	}

	if c.noMining {
		args = append(args, "--no-mining")
	}

	if c.number != 0 {
		args = append(args, "--number", fmt.Sprint(c.number))
	}

	if c.order != "" && c.order != "fees" {
		args = append(args, "--order", c.order)
	}

	if c.port != 0 {
		args = append(args, "--port", fmt.Sprint(c.port))
	}

	if c.preserveHistoricalStates {
		args = append(args, "--preserve-historical-states")
	}

	if c.pruneHistory {
		if c.pruneHistoryStates == 0 {
			args = append(args, "--prune-history")
		} else {
			args = append(args, "--prune-history", fmt.Sprint(c.pruneHistoryStates))
		}
	}

	if c.stateInterval != 0 {
		args = append(args, "--state-interval", fmt.Sprint(c.stateInterval))
	}

	if c.slotsInEpoch != 0 {
		args = append(args, "--slots-in-an-epoch", fmt.Sprint(c.slotsInEpoch))
	}

	if c.statePath != "" {
		args = append(args, "--state", c.statePath)
	}

	if c.timestamp != 0 {
		args = append(args, "--timestamp", fmt.Sprint(c.timestamp))
	}

	if c.transactionBlockKeeper != 0 {
		args = append(args, "--transaction-block-keeper", fmt.Sprint(c.transactionBlockKeeper))
	}

	if c.color != "" {
		args = append(args, "--color", c.color)
	}

	if c.jsonLogs {
		args = append(args, "--json")
	}

	if c.markdownLogs {
		args = append(args, "--md")
	}

	if c.quiet {
		args = append(args, "--quiet")
	}

	if c.verbosity > 0 {
		vv := "-" + strings.Repeat("v", c.verbosity)
		args = append(args, vv)
	}

	if c.allowOrigin != "" {
		args = append(args, "--allow-origin", c.allowOrigin)
	}

	if c.cachePath != "" {
		args = append(args, "--cache-path", c.cachePath)
	}

	if c.host != "" {
		args = append(args, "--host", c.host)
	}

	if c.noCors {
		args = append(args, "--no-cors")
	}

	if c.noRequestSizeLimit {
		args = append(args, "--no-request-size-limit")
	}

	if c.computeUnitsPerSecond != 0 {
		args = append(args, "--compute-units-per-second", fmt.Sprint(c.computeUnitsPerSecond))
	}

	if c.forkURL != "" {
		args = append(args, "--fork-url", c.forkURL)
	}

	if c.forkBlockNumber != 0 {
		args = append(args, "--fork-block-number", fmt.Sprint(c.forkBlockNumber))
	}

	if c.forkChainID != 0 {
		args = append(args, "--fork-chain-id", fmt.Sprint(c.forkChainID))
	}

	if len(c.forkHeaders) > 0 {
		args = append(args, "--fork-header")
		for _, h := range c.forkHeaders {
			args = append(args, h)
		}
	}

	if c.forkRetryBackoff != "" {
		args = append(args, "--fork-retry-backoff", c.forkRetryBackoff)
	}

	if c.forkTransactionHash != "" {
		args = append(args, "--fork-transaction-hash", c.forkTransactionHash)
	}

	if c.noRateLimit {
		args = append(args, "--no-rate-limit")
	}

	if c.noStorageCaching {
		args = append(args, "--no-storage-caching")
	}

	if c.retries != 0 {
		args = append(args, "--retries", fmt.Sprint(c.retries))
	}

	if c.timeout != "" {
		args = append(args, "--timeout", c.timeout)
	}

	if c.blockBaseFeePerGas != "" {
		args = append(args, "--block-base-fee-per-gas", c.blockBaseFeePerGas)
	}

	if c.chainID != 0 {
		args = append(args, "--chain-id", fmt.Sprint(c.chainID))
	}

	if c.codeSizeLimit != 0 {
		args = append(args, "--code-size-limit", fmt.Sprintf("0x%x", c.codeSizeLimit))
	}

	if c.disableBlockGasLimit {
		args = append(args, "--disable-block-gas-limit")
	}

	if c.disableCodeSizeLimit {
		args = append(args, "--disable-code-size-limit")
	}

	if c.disableMinPriorityFee {
		args = append(args, "--disable-min-priority-fee")
	}

	if c.gasLimit != 0 {
		args = append(args, "--gas-limit", fmt.Sprint(c.gasLimit))
	}

	if c.gasPrice != "" {
		args = append(args, "--gas-price", c.gasPrice)
	}

	if c.autoImpersonate {
		args = append(args, "--auto-impersonate")
	}

	if c.disableConsoleLog {
		args = append(args, "--disable-console-log")
	}

	if c.disableDefaultCreate2Deployer {
		args = append(args, "--disable-default-create2-deployer")
	}

	if c.disablePoolBalanceChecks {
		args = append(args, "--disable-pool-balance-checks")
	}

	if c.memoryLimit != 0 {
		args = append(args, "--memory-limit", fmt.Sprint(c.memoryLimit))
	}

	if c.printTraces {
		args = append(args, "--print-traces")
	}

	if c.stepsTracing {
		args = append(args, "--steps-tracing")
	}

	if c.celo {
		args = append(args, "--celo")
	}

	if c.optimism {
		args = append(args, "--optimism")
	}

	return args
}
