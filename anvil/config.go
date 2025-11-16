package anvil

// AnvilConfig configures how the Anvil process is started.
// Zero values generally mean "omit this flag and let Anvil use its default".
type Config struct {
	// Number of dev accounts to generate and configure.
	//
	// CLI: -a, --accounts
	// Default: 10
	accounts uint

	// Block time in seconds for interval mining.
	// 0 => omit flag => instant mining / default behavior.
	//
	// CLI: -b, --block-time
	blockTime uint

	// The balance of every dev account in Ether.
	// Empty string => omit flag => default 10000 ETH.
	//
	// CLI: --balance
	// Default: "10000"
	balance string

	// Writes output of `anvil` as JSON to user-specified file.
	// Empty => omit flag.
	//
	// CLI: --config-out
	configOut string

	// Derivation path of the child key to be derived.
	//
	// CLI: --derivation-path
	// Default: "m/44'/60'/0'/0/"
	derivationPath string

	// Dump the state and block environment of chain on exit to the given path.
	// If the value is a directory, state is written to "<value>/state.json".
	//
	// CLI: --dump-state
	dumpStatePath string

	// EVM hardfork to use (e.g. "prague", "cancun", "shanghai", "paris", "london", "latest").
	//
	// CLI: --hardfork
	// Default: "latest"
	hardfork string

	// Initialize the genesis block with the given `genesis.json` file.
	//
	// CLI: --init
	initPath string

	// Launch an ipc server.
	// If ipc is true and ipcPath == "", Anvil's default path is used (/tmp/anvil.ipc).
	//
	// CLI: --ipc [<PATH>]
	ipc     bool
	ipcPath string

	// Number of threads to use.
	// Number of threads to use. Specifying 0 defaults to the number of logical cores
	//
	// CLI: -j, --threads
	threads uint

	// Initialize the chain from a previously saved state snapshot.
	//
	// CLI: --load-state
	loadStatePath string

	// BIP39 mnemonic phrase used for generating accounts.
	// Mutually exclusive with MnemonicRandom and MnemonicSeedUnsafe.
	//
	// CLI: -m, --mnemonic
	mnemonic string

	// Max number of states to persist on disk.
	// Note that `prune_history` will overwrite `max_persisted_states` to 0.
	//
	// CLI: --max-persisted-states
	maxPersistedStates uint

	// Enable mixed mining.
	//
	// CLI: --mixed-mining
	mixedMining bool

	// Automatically generate a BIP39 mnemonic phrase and derive accounts from it.
	// Mutually exclusive with Mnemonic and MnemonicSeedUnsafe.
	//
	// If true and MnemonicRandomWords == 0, Anvil uses its default (12).
	//
	// CLI: --mnemonic-random [<MNEMONIC_RANDOM>]
	mnemonicRandom      bool
	mnemonicRandomWords uint

	// Generate a BIP39 mnemonic phrase from a given seed (UNSAFE).
	// Mutually exclusive with Mnemonic and MnemonicRandom.
	// CAREFUL: This is NOT SAFE and should only be used for testing.
	// Never use the private keys generated in production.
	//
	// CLI: --mnemonic-seed-unsafe
	mnemonicSeedUnsafe string

	// Disable auto and interval mining; mine on demand instead.
	//
	// CLI: --no-mining (alias: --no-mine)
	noMining bool

	// The number of the genesis block.
	//
	// CLI: --number
	number uint64

	// How transactions are sorted in the mempool (e.g. "fees", "fifo").
	//
	// CLI: --order
	// Default: "fees"
	order string

	// port number to listen on.
	//
	// CLI: -p, --port
	// Default: 8545
	port uint

	// Preserve historical state snapshots when dumping state.
	//
	// CLI: --preserve-historical-states
	preserveHistoricalStates bool

	// Don't keep full chain history.
	// If pruneHistory is true and PruneHistoryStates > 0, at most that many
	// states are kept in memory.
	//
	// CLI: --prune-history [<PRUNE_HISTORY>]
	pruneHistory       bool
	pruneHistoryStates uint

	// Interval in seconds at which the state and block environment is dumped to disk.
	//
	// CLI: -s, --state-interval
	stateInterval uint

	// Slots in an epoch.
	//
	// CLI: --slots-in-an-epoch
	// Default: 32
	slotsInEpoch uint

	// This is an alias for both --load-state and --dump-state.
	// It initializes the chain with the state stored at the file (if it exists)
	// and dumps the chain's state on exit.
	//
	// CLI: --state
	statePath string

	// timestamp of the genesis block.
	//
	// CLI: --timestamp
	timestamp uint64

	// Number of blocks with transactions to keep in memory.
	//
	// CLI: --transaction-block-keeper
	transactionBlockKeeper uint64

	// Show logs during execution.
	showLogs bool

	// Log color mode: "auto", "always", or "never".
	//
	// CLI: --color
	color string

	// Format log messages as JSON.
	//
	// CLI: --json
	jsonLogs bool

	// Format log messages as Markdown.
	//
	// CLI: --md
	markdownLogs bool

	// Do not print log messages.
	//
	// CLI: -q, --quiet
	quiet bool

	// verbosity level of log messages.
	// 0 => omit flag; >0 => pass -v multiple times.
	//
	// CLI: -v, --verbosity...
	verbosity int

	// CORS allow_origin header.
	//
	// CLI: --allow-origin
	// Default: "*"
	allowOrigin string

	// Path to the cache directory where states are stored.
	//
	// CLI: --cache-path
	cachePath string

	// Hosts the server will listen on (e.g. "127.0.0.1", "0.0.0.0").
	//
	// CLI: --host
	// Default: "127.0.0.1"
	host string

	// Disable CORS.
	//
	// CLI: --no-cors
	noCors bool

	// Disable default request body size limit.
	//
	// CLI: --no-request-size-limit
	noRequestSizeLimit bool

	// Number of assumed available compute units per second for this provider.
	//
	// CLI: --compute-units-per-second
	// Default: 330
	computeUnitsPerSecond uint

	// Fetch state over a remote endpoint instead of starting from an empty state.
	//
	// CLI: -f, --fork-url (alias: --rpc-url)
	forkURL string

	// Fetch state from a specific block number over a remote endpoint.
	// If negative, the value is subtracted from the latest block.
	//
	// CLI: --fork-block-number
	forkBlockNumber int64

	// Specify chain ID to skip fetching it from remote endpoint.
	//
	// CLI: --fork-chain-id
	forkChainID uint64

	// Headers for the RPC client, e.g. "User-Agent: test-agent".
	// These will be rendered as multiple --fork-header values.
	//
	// CLI: --fork-header
	forkHeaders []string

	// Initial retry backoff on encountering errors with fork provider.
	// Represented as string so you can pass exactly what Anvil expects.
	//
	// CLI: --fork-retry-backoff
	forkRetryBackoff string

	// Fetch state from after a specific transaction hash.
	//
	// CLI: --fork-transaction-hash
	forkTransactionHash string

	// Disable rate limiting for this node's provider.
	//
	// CLI: --no-rate-limit (alias: --no-rpc-rate-limit)
	noRateLimit bool

	// Explicitly disable RPC storage caching.
	//
	// CLI: --no-storage-caching
	noStorageCaching bool

	// Number of retry requests for spurious networks.
	//
	// CLI: --retries
	// Default: 5
	retries uint

	// timeout for requests sent to remote JSON-RPC server in forking mode, in ms.
	// Represented as string for direct pass-through (e.g. "45000").
	//
	// CLI: --timeout
	// Default: "45000"
	timeout string

	// Base fee in a block.
	// String to allow decimal or hex (e.g. "1000000000" or "0x3b9aca00").
	//
	// CLI: --block-base-fee-per-gas (alias: --base-fee)
	blockBaseFeePerGas string

	// Chain ID.
	//
	// CLI: --chain-id
	chainID uint64

	// Contract code size limit in bytes.
	//
	// CLI: --code-size-limit
	// Default: 0x6000
	codeSizeLimit uint64

	// Disable the call.gas_limit <= block.gas_limit constraint.
	//
	// CLI: --disable-block-gas-limit
	disableBlockGasLimit bool

	// Disable EIP-170 contract code size limit.
	//
	// CLI: --disable-code-size-limit
	disableCodeSizeLimit bool

	// Disable enforcement of a minimum suggested priority fee.
	//
	// CLI: --disable-min-priority-fee (alias: --no-priority-fee)
	disableMinPriorityFee bool

	// Block gas limit.
	//
	// CLI: --gas-limit
	gasLimit uint64

	// Gas price (string for decimal or hex).
	//
	// CLI: --gas-price
	gasPrice string

	// Enables automatic impersonation on startup.
	//
	// CLI: --auto-impersonate (alias: --auto-unlock)
	autoImpersonate bool

	// Disable printing of console.log invocations to stdout.
	//
	// CLI: --disable-console-log (alias: --no-console-log)
	disableConsoleLog bool

	// Disable the default CREATE2 deployer.
	//
	// CLI: --disable-default-create2-deployer (alias: --no-create2)
	disableDefaultCreate2Deployer bool

	// Disable pool balance checks.
	//
	// CLI: --disable-pool-balance-checks
	disablePoolBalanceChecks bool

	// Memory limit per EVM execution in bytes.
	//
	// CLI: --memory-limit
	memoryLimit uint64

	// Enable printing of traces for executed txs and eth_call.
	//
	// CLI: --print-traces (alias: --enable-trace-printing)
	printTraces bool

	// Enable steps tracing for debug/geth-style traces.
	//
	// CLI: --steps-tracing (alias: --tracing)
	stepsTracing bool

	// Enable celo network features.
	//
	// CLI: --celo
	celo bool

	// Enable optimism network features.
	//
	// CLI: --optimism
	optimism bool
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	c := Config{}
	return &c
}

// SetAccounts sets Accounts, the number of dev accounts to generate and configure.
// A value of 0 omits the flag and lets Anvil use its default (10).
func (c *Config) SetAccounts(n uint) *Config {
	c.accounts = n
	return c
}

// SetBlockTime sets BlockTime, the block time in seconds for interval mining.
// A value of 0 omits the flag and uses Anvil's instant-mining / default behavior.
func (c *Config) SetBlockTime(seconds uint) *Config {
	c.blockTime = seconds
	return c
}

// SetBalance sets Balance, the balance of every dev account in Ether.
// An empty string omits the flag and uses Anvil's default of 10000 ETH.
func (c *Config) SetBalance(balance string) *Config {
	c.balance = balance
	return c
}

// SetConfigOut sets ConfigOut, the path where Anvil writes JSON output for its configuration.
// An empty string omits the flag.
func (c *Config) SetConfigOut(path string) *Config {
	c.configOut = path
	return c
}

// SetDerivationPath sets DerivationPath, the BIP32 derivation path used for deriving accounts.
// If unset, Anvil uses its default "m/44'/60'/0'/0/".
func (c *Config) SetDerivationPath(path string) *Config {
	c.derivationPath = path
	return c
}

// SetDumpStatePath sets DumpStatePath, the path where Anvil dumps state and block environment on exit.
// If the value is a directory, state is written to "<value>/state.json".
func (c *Config) SetDumpStatePath(path string) *Config {
	c.dumpStatePath = path
	return c
}

// SetHardfork sets Hardfork, the EVM hardfork to use (e.g. "prague", "cancun", "shanghai", "paris", "london", "latest").
// If unset, Anvil uses "latest".
func (c *Config) SetHardfork(hardfork string) *Config {
	c.hardfork = hardfork
	return c
}

// SetInitPath sets InitPath, the path to a genesis.json file used to initialize the genesis block.
func (c *Config) SetInitPath(path string) *Config {
	c.initPath = path
	return c
}

// SetIPC enables or disables the IPC server via IPC, and optionally sets IPCPath.
// If IPC is true and IPCPath is empty, Anvil uses its default (/tmp/anvil.ipc).
func (c *Config) SetIPC(enabled bool, path string) *Config {
	c.ipc = enabled
	c.ipcPath = path
	return c
}

// SetIPCEnabled sets IPC to enable or disable the IPC server, without changing IPCPath.
func (c *Config) SetIPCEnabled(enabled bool) *Config {
	c.ipc = enabled
	return c
}

// SetIPCPath sets IPCPath, the path for the IPC server socket.
// IPC must be true for IPCPath to be used; if empty, Anvil uses its default.
func (c *Config) SetIPCPath(path string) *Config {
	c.ipcPath = path
	return c
}

// SetThreads sets Threads, the number of threads to use.
// A value of 0 omits the flag and uses the number of logical cores.
func (c *Config) SetThreads(n uint) *Config {
	c.threads = n
	return c
}

// SetLoadStatePath sets LoadStatePath, the path to a previously saved state snapshot to initialize the chain.
func (c *Config) SetLoadStatePath(path string) *Config {
	c.loadStatePath = path
	return c
}

// SetMnemonic sets Mnemonic, the BIP39 mnemonic phrase used for generating accounts.
// This is mutually exclusive with MnemonicRandom and MnemonicSeedUnsafe.
func (c *Config) SetMnemonic(mnemonic string) *Config {
	c.mnemonic = mnemonic
	return c
}

// SetMaxPersistedStates sets MaxPersistedStates, the maximum number of states to persist on disk.
func (c *Config) SetMaxPersistedStates(n uint) *Config {
	c.maxPersistedStates = n
	return c
}

// SetMixedMining sets MixedMining, enabling or disabling mixed mining.
func (c *Config) SetMixedMining(enabled bool) *Config {
	c.mixedMining = enabled
	return c
}

// SetMnemonicRandom sets MnemonicRandom, whether Anvil should automatically generate a BIP39 mnemonic.
// This is mutually exclusive with Mnemonic and MnemonicSeedUnsafe.
// If true and MnemonicRandomWords is 0, Anvil uses its default word count (12).
func (c *Config) SetMnemonicRandom(enabled bool) *Config {
	c.mnemonicRandom = enabled
	return c
}

// SetMnemonicRandomWords sets MnemonicRandomWords, the number of BIP39 words to use when MnemonicRandom is true.
func (c *Config) SetMnemonicRandomWords(words uint) *Config {
	c.mnemonicRandomWords = words
	return c
}

// SetMnemonicSeedUnsafe sets MnemonicSeedUnsafe, the seed used to generate a BIP39 mnemonic (UNSAFE).
// This is mutually exclusive with Mnemonic and MnemonicRandom.
func (c *Config) SetMnemonicSeedUnsafe(seed string) *Config {
	c.mnemonicSeedUnsafe = seed
	return c
}

// SetNoMining sets NoMining, disabling auto and interval mining to allow mining on demand only.
func (c *Config) SetNoMining(noMining bool) *Config {
	c.noMining = noMining
	return c
}

// SetNumber sets Number, the number of the genesis block.
func (c *Config) SetNumber(number uint64) *Config {
	c.number = number
	return c
}

// SetOrder sets Order, how transactions are sorted in the mempool (e.g. "fees", "fifo").
// If unset, Anvil uses "fees".
func (c *Config) SetOrder(order string) *Config {
	c.order = order
	return c
}

// SetPort sets Port, the port number Anvil listens on.
// If unset, Anvil uses its default (8545).
func (c *Config) SetPort(port uint) *Config {
	c.port = port
	return c
}

// SetPreserveHistoricalStates sets PreserveHistoricalStates, controlling whether historical state
// snapshots are preserved when dumping state.
func (c *Config) SetPreserveHistoricalStates(preserve bool) *Config {
	c.preserveHistoricalStates = preserve
	return c
}

// SetPruneHistory sets PruneHistory, enabling or disabling history pruning.
// When true and PruneHistoryStates > 0, at most that many states are kept in memory.
func (c *Config) SetPruneHistory(enabled bool) *Config {
	c.pruneHistory = enabled
	return c
}

// SetPruneHistoryStates sets PruneHistoryStates, the maximum number of states kept in memory when pruning history.
func (c *Config) SetPruneHistoryStates(n uint) *Config {
	c.pruneHistoryStates = n
	return c
}

// SetStateInterval sets StateInterval, the interval in seconds at which state and block environment
// are periodically dumped to disk.
func (c *Config) SetStateInterval(seconds uint) *Config {
	c.stateInterval = seconds
	return c
}

// SetSlotsInEpoch sets SlotsInEpoch, the number of slots in an epoch.
// If unset, Anvil uses its default (32).
func (c *Config) SetSlotsInEpoch(slots uint) *Config {
	c.slotsInEpoch = slots
	return c
}

// SetStatePath sets StatePath, a path used as an alias for both load-state and dump-state.
// It initializes the chain with the state stored at the file (if it exists) and dumps state on exit.
func (c *Config) SetStatePath(path string) *Config {
	c.statePath = path
	return c
}

// SetTimestamp sets Timestamp, the timestamp of the genesis block.
func (c *Config) SetTimestamp(timestamp uint64) *Config {
	c.timestamp = timestamp
	return c
}

// SetTransactionBlockKeeper sets TransactionBlockKeeper, the number of blocks with transactions
// to keep in memory.
func (c *Config) SetTransactionBlockKeeper(n uint64) *Config {
	c.transactionBlockKeeper = n
	return c
}

// SetShowLogs sets ShowLogs, enabling or disabling log output.
func (c *Config) SetShowLogs(show bool) *Config {
	c.showLogs = show
	return c
}

// SetColor sets Color, the log color mode ("auto", "always", or "never").
func (c *Config) SetColor(color string) *Config {
	c.color = color
	return c
}

// SetJSONLogs sets JSONLogs, enabling or disabling JSON-formatted log output.
func (c *Config) SetJSONLogs(enabled bool) *Config {
	c.jsonLogs = enabled
	return c
}

// SetMarkdownLogs sets MarkdownLogs, enabling or disabling Markdown-formatted log output.
func (c *Config) SetMarkdownLogs(enabled bool) *Config {
	c.markdownLogs = enabled
	return c
}

// SetQuiet sets Quiet, enabling or disabling log output entirely.
func (c *Config) SetQuiet(quiet bool) *Config {
	c.quiet = quiet
	return c
}

// SetVerbosity sets Verbosity, the verbosity level of log messages.
// A value of 0 omits the flag; values > 0 cause -v to be passed multiple times.
func (c *Config) SetVerbosity(level int) *Config {
	c.verbosity = level
	return c
}

// SetAllowOrigin sets AllowOrigin, the CORS allow_origin header.
// If unset, Anvil uses its default ("*").
func (c *Config) SetAllowOrigin(origin string) *Config {
	c.allowOrigin = origin
	return c
}

// SetCachePath sets CachePath, the path to the cache directory where states are stored.
func (c *Config) SetCachePath(path string) *Config {
	c.cachePath = path
	return c
}

// SetHost sets Host, the host address the server will listen on (e.g. "127.0.0.1", "0.0.0.0").
// If unset, Anvil uses its default ("127.0.0.1").
func (c *Config) SetHost(host string) *Config {
	c.host = host
	return c
}

// SetNoCors sets NoCors, disabling CORS when true.
func (c *Config) SetNoCors(noCors bool) *Config {
	c.noCors = noCors
	return c
}

// SetNoRequestSizeLimit sets NoRequestSizeLimit, disabling the default request body size limit when true.
func (c *Config) SetNoRequestSizeLimit(noLimit bool) *Config {
	c.noRequestSizeLimit = noLimit
	return c
}

// SetComputeUnitsPerSecond sets ComputeUnitsPerSecond, the assumed available compute units per second for this provider.
// If unset, Anvil uses its default (330).
func (c *Config) SetComputeUnitsPerSecond(units uint) *Config {
	c.computeUnitsPerSecond = units
	return c
}

// SetForkURL sets ForkURL, the remote endpoint URL used to fetch state instead of starting from an empty state.
func (c *Config) SetForkURL(url string) *Config {
	c.forkURL = url
	return c
}

// SetForkBlockNumber sets ForkBlockNumber, the block number from which to fetch state.
// If negative, the value is subtracted from the latest block.
func (c *Config) SetForkBlockNumber(block int64) *Config {
	c.forkBlockNumber = block
	return c
}

// SetForkChainID sets ForkChainID, the chain ID to use when forking, skipping fetching it from the remote endpoint.
func (c *Config) SetForkChainID(id uint64) *Config {
	c.forkChainID = id
	return c
}

// SetForkHeaders sets ForkHeaders, the headers used by the RPC client (e.g. "User-Agent: test-agent").
// These are rendered as multiple --fork-header values.
func (c *Config) SetForkHeaders(headers []string) *Config {
	c.forkHeaders = headers
	return c
}

// AddForkHeader appends a single header to ForkHeaders.
func (c *Config) AddForkHeader(header string) *Config {
	c.forkHeaders = append(c.forkHeaders, header)
	return c
}

// SetForkRetryBackoff sets ForkRetryBackoff, the initial retry backoff duration used
// on errors with the fork provider, as a string in Anvil's expected format.
func (c *Config) SetForkRetryBackoff(backoff string) *Config {
	c.forkRetryBackoff = backoff
	return c
}

// SetForkTransactionHash sets ForkTransactionHash, the transaction hash after which
// to fetch state when forking.
func (c *Config) SetForkTransactionHash(hash string) *Config {
	c.forkTransactionHash = hash
	return c
}

// SetNoRateLimit sets NoRateLimit, disabling rate limiting for this node's provider when true.
func (c *Config) SetNoRateLimit(disable bool) *Config {
	c.noRateLimit = disable
	return c
}

// SetNoStorageCaching sets NoStorageCaching, explicitly disabling RPC storage caching when true.
func (c *Config) SetNoStorageCaching(disable bool) *Config {
	c.noStorageCaching = disable
	return c
}

// SetRetries sets Retries, the number of retry requests for spurious networks.
// If unset, Anvil uses its default (5).
func (c *Config) SetRetries(retries uint) *Config {
	c.retries = retries
	return c
}

// SetTimeout sets Timeout, the timeout in milliseconds for requests sent to the remote
// JSON-RPC server in forking mode, as a string (e.g. "45000").
// If unset, Anvil uses its default ("45000").
func (c *Config) SetTimeout(timeout string) *Config {
	c.timeout = timeout
	return c
}

// SetBlockBaseFeePerGas sets BlockBaseFeePerGas, the base fee in a block,
// expressed as a string to allow decimal or hex (e.g. "1000000000" or "0x3b9aca00").
func (c *Config) SetBlockBaseFeePerGas(fee string) *Config {
	c.blockBaseFeePerGas = fee
	return c
}

// SetChainID sets ChainID, the chain ID to use.
func (c *Config) SetChainID(id uint64) *Config {
	c.chainID = id
	return c
}

// SetCodeSizeLimit sets CodeSizeLimit, the contract code size limit in bytes.
// If unset, Anvil uses its default (0x6000).
func (c *Config) SetCodeSizeLimit(limit uint64) *Config {
	c.codeSizeLimit = limit
	return c
}

// SetDisableBlockGasLimit sets DisableBlockGasLimit, disabling the call.gas_limit <= block.gas_limit constraint when true.
func (c *Config) SetDisableBlockGasLimit(disable bool) *Config {
	c.disableBlockGasLimit = disable
	return c
}

// SetDisableCodeSizeLimit sets DisableCodeSizeLimit, disabling the EIP-170 contract code size limit when true.
func (c *Config) SetDisableCodeSizeLimit(disable bool) *Config {
	c.disableCodeSizeLimit = disable
	return c
}

// SetDisableMinPriorityFee sets DisableMinPriorityFee, disabling enforcement of a minimum suggested priority fee when true.
func (c *Config) SetDisableMinPriorityFee(disable bool) *Config {
	c.disableMinPriorityFee = disable
	return c
}

// SetGasLimit sets GasLimit, the block gas limit.
func (c *Config) SetGasLimit(limit uint64) *Config {
	c.gasLimit = limit
	return c
}

// SetGasPrice sets GasPrice, the gas price to use as a string, allowing decimal or hex.
func (c *Config) SetGasPrice(price string) *Config {
	c.gasPrice = price
	return c
}

// SetAutoImpersonate sets AutoImpersonate, enabling or disabling automatic impersonation (auto-unlock) on startup.
func (c *Config) SetAutoImpersonate(enabled bool) *Config {
	c.autoImpersonate = enabled
	return c
}

// SetDisableConsoleLog sets DisableConsoleLog, disabling printing of console.log invocations to stdout when true.
func (c *Config) SetDisableConsoleLog(disable bool) *Config {
	c.disableConsoleLog = disable
	return c
}

// SetDisableDefaultCreate2Deployer sets DisableDefaultCreate2Deployer, disabling the default CREATE2 deployer when true.
func (c *Config) SetDisableDefaultCreate2Deployer(disable bool) *Config {
	c.disableDefaultCreate2Deployer = disable
	return c
}

// SetDisablePoolBalanceChecks sets DisablePoolBalanceChecks, disabling pool balance checks when true.
func (c *Config) SetDisablePoolBalanceChecks(disable bool) *Config {
	c.disablePoolBalanceChecks = disable
	return c
}

// SetMemoryLimit sets MemoryLimit, the memory limit per EVM execution in bytes.
func (c *Config) SetMemoryLimit(limit uint64) *Config {
	c.memoryLimit = limit
	return c
}

// SetPrintTraces sets PrintTraces, enabling printing of traces for executed transactions and eth_call when true.
func (c *Config) SetPrintTraces(enabled bool) *Config {
	c.printTraces = enabled
	return c
}

// SetStepsTracing sets StepsTracing, enabling steps tracing for debug/geth-style traces when true.
func (c *Config) SetStepsTracing(enabled bool) *Config {
	c.stepsTracing = enabled
	return c
}

// SetCelo sets Celo, enabling Celo network features when true.
func (c *Config) SetCelo(enabled bool) *Config {
	c.celo = enabled
	return c
}

// SetOptimism sets Optimism, enabling Optimism network features when true.
func (c *Config) SetOptimism(enabled bool) *Config {
	c.optimism = enabled
	return c
}
