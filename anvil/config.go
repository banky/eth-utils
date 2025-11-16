package anvil

// AnvilConfig configures how the Anvil process is started.
// Zero values generally mean "omit this flag and let Anvil use its default".
type Config struct {
	// Number of dev accounts to generate and configure.
	//
	// CLI: -a, --accounts
	// Default: 10
	Accounts uint

	// Block time in seconds for interval mining.
	// 0 => omit flag => instant mining / default behavior.
	//
	// CLI: -b, --block-time
	BlockTime uint

	// The balance of every dev account in Ether.
	// Empty string => omit flag => default 10000 ETH.
	//
	// CLI: --balance
	// Default: "10000"
	Balance string

	// Writes output of `anvil` as JSON to user-specified file.
	// Empty => omit flag.
	//
	// CLI: --config-out
	ConfigOut string

	// Derivation path of the child key to be derived.
	//
	// CLI: --derivation-path
	// Default: "m/44'/60'/0'/0/"
	DerivationPath string

	// Dump the state and block environment of chain on exit to the given path.
	// If the value is a directory, state is written to "<value>/state.json".
	//
	// CLI: --dump-state
	DumpStatePath string

	// EVM hardfork to use (e.g. "prague", "cancun", "shanghai", "paris", "london", "latest").
	//
	// CLI: --hardfork
	// Default: "latest"
	Hardfork string

	// Initialize the genesis block with the given `genesis.json` file.
	//
	// CLI: --init
	InitPath string

	// Launch an IPC server.
	// If IPC is true and IPCPath == "", Anvil's default path is used (/tmp/anvil.ipc).
	//
	// CLI: --ipc [<PATH>]
	IPC     bool
	IPCPath string

	// Number of threads to use.
	// Number of threads to use. Specifying 0 defaults to the number of logical cores
	//
	// CLI: -j, --threads
	Threads uint

	// Initialize the chain from a previously saved state snapshot.
	//
	// CLI: --load-state
	LoadStatePath string

	// BIP39 mnemonic phrase used for generating accounts.
	// Mutually exclusive with MnemonicRandom and MnemonicSeedUnsafe.
	//
	// CLI: -m, --mnemonic
	Mnemonic string

	// Max number of states to persist on disk.
	// Note that `prune_history` will overwrite `max_persisted_states` to 0.
	//
	// CLI: --max-persisted-states
	MaxPersistedStates uint

	// Enable mixed mining.
	//
	// CLI: --mixed-mining
	MixedMining bool

	// Automatically generate a BIP39 mnemonic phrase and derive accounts from it.
	// Mutually exclusive with Mnemonic and MnemonicSeedUnsafe.
	//
	// If true and MnemonicRandomWords == 0, Anvil uses its default (12).
	//
	// CLI: --mnemonic-random [<MNEMONIC_RANDOM>]
	MnemonicRandom      bool
	MnemonicRandomWords uint

	// Generate a BIP39 mnemonic phrase from a given seed (UNSAFE).
	// Mutually exclusive with Mnemonic and MnemonicRandom.
	// CAREFUL: This is NOT SAFE and should only be used for testing.
	// Never use the private keys generated in production.
	//
	// CLI: --mnemonic-seed-unsafe
	MnemonicSeedUnsafe string

	// Disable auto and interval mining; mine on demand instead.
	//
	// CLI: --no-mining (alias: --no-mine)
	NoMining bool

	// The number of the genesis block.
	//
	// CLI: --number
	Number uint64

	// How transactions are sorted in the mempool (e.g. "fees", "fifo").
	//
	// CLI: --order
	// Default: "fees"
	Order string

	// Port number to listen on.
	//
	// CLI: -p, --port
	// Default: 8545
	Port uint

	// Preserve historical state snapshots when dumping state.
	//
	// CLI: --preserve-historical-states
	PreserveHistoricalStates bool

	// Don't keep full chain history.
	// If PruneHistory is true and PruneHistoryStates > 0, at most that many
	// states are kept in memory.
	//
	// CLI: --prune-history [<PRUNE_HISTORY>]
	PruneHistory       bool
	PruneHistoryStates uint

	// Interval in seconds at which the state and block environment is dumped to disk.
	//
	// CLI: -s, --state-interval
	StateInterval uint

	// Slots in an epoch.
	//
	// CLI: --slots-in-an-epoch
	// Default: 32
	SlotsInEpoch uint

	// This is an alias for both --load-state and --dump-state.
	// It initializes the chain with the state stored at the file (if it exists)
	// and dumps the chain's state on exit.
	//
	// CLI: --state
	StatePath string

	// Timestamp of the genesis block.
	//
	// CLI: --timestamp
	Timestamp uint64

	// Number of blocks with transactions to keep in memory.
	//
	// CLI: --transaction-block-keeper
	TransactionBlockKeeper uint64

	// Show logs during execution.
	ShowLogs bool

	// Log color mode: "auto", "always", or "never".
	//
	// CLI: --color
	Color string

	// Format log messages as JSON.
	//
	// CLI: --json
	JSONLogs bool

	// Format log messages as Markdown.
	//
	// CLI: --md
	MarkdownLogs bool

	// Do not print log messages.
	//
	// CLI: -q, --quiet
	Quiet bool

	// Verbosity level of log messages.
	// 0 => omit flag; >0 => pass -v multiple times.
	//
	// CLI: -v, --verbosity...
	Verbosity int

	// CORS allow_origin header.
	//
	// CLI: --allow-origin
	// Default: "*"
	AllowOrigin string

	// Path to the cache directory where states are stored.
	//
	// CLI: --cache-path
	CachePath string

	// Hosts the server will listen on (e.g. "127.0.0.1", "0.0.0.0").
	//
	// CLI: --host
	// Default: "127.0.0.1"
	Host string

	// Disable CORS.
	//
	// CLI: --no-cors
	NoCors bool

	// Disable default request body size limit.
	//
	// CLI: --no-request-size-limit
	NoRequestSizeLimit bool

	// Number of assumed available compute units per second for this provider.
	//
	// CLI: --compute-units-per-second
	// Default: 330
	ComputeUnitsPerSecond uint

	// Fetch state over a remote endpoint instead of starting from an empty state.
	//
	// CLI: -f, --fork-url (alias: --rpc-url)
	ForkURL string

	// Fetch state from a specific block number over a remote endpoint.
	// If negative, the value is subtracted from the latest block.
	//
	// CLI: --fork-block-number
	ForkBlockNumber int64

	// Specify chain ID to skip fetching it from remote endpoint.
	//
	// CLI: --fork-chain-id
	ForkChainID uint64

	// Headers for the RPC client, e.g. "User-Agent: test-agent".
	// These will be rendered as multiple --fork-header values.
	//
	// CLI: --fork-header
	ForkHeaders []string

	// Initial retry backoff on encountering errors with fork provider.
	// Represented as string so you can pass exactly what Anvil expects.
	//
	// CLI: --fork-retry-backoff
	ForkRetryBackoff string

	// Fetch state from after a specific transaction hash.
	//
	// CLI: --fork-transaction-hash
	ForkTransactionHash string

	// Disable rate limiting for this node's provider.
	//
	// CLI: --no-rate-limit (alias: --no-rpc-rate-limit)
	NoRateLimit bool

	// Explicitly disable RPC storage caching.
	//
	// CLI: --no-storage-caching
	NoStorageCaching bool

	// Number of retry requests for spurious networks.
	//
	// CLI: --retries
	// Default: 5
	Retries uint

	// Timeout for requests sent to remote JSON-RPC server in forking mode, in ms.
	// Represented as string for direct pass-through (e.g. "45000").
	//
	// CLI: --timeout
	// Default: "45000"
	Timeout string

	// Base fee in a block.
	// String to allow decimal or hex (e.g. "1000000000" or "0x3b9aca00").
	//
	// CLI: --block-base-fee-per-gas (alias: --base-fee)
	BlockBaseFeePerGas string

	// Chain ID.
	//
	// CLI: --chain-id
	ChainID uint64

	// Contract code size limit in bytes.
	//
	// CLI: --code-size-limit
	// Default: 0x6000
	CodeSizeLimit uint64

	// Disable the call.gas_limit <= block.gas_limit constraint.
	//
	// CLI: --disable-block-gas-limit
	DisableBlockGasLimit bool

	// Disable EIP-170 contract code size limit.
	//
	// CLI: --disable-code-size-limit
	DisableCodeSizeLimit bool

	// Disable enforcement of a minimum suggested priority fee.
	//
	// CLI: --disable-min-priority-fee (alias: --no-priority-fee)
	DisableMinPriorityFee bool

	// Block gas limit.
	//
	// CLI: --gas-limit
	GasLimit uint64

	// Gas price (string for decimal or hex).
	//
	// CLI: --gas-price
	GasPrice string

	// Enables automatic impersonation on startup.
	//
	// CLI: --auto-impersonate (alias: --auto-unlock)
	AutoImpersonate bool

	// Disable printing of console.log invocations to stdout.
	//
	// CLI: --disable-console-log (alias: --no-console-log)
	DisableConsoleLog bool

	// Disable the default CREATE2 deployer.
	//
	// CLI: --disable-default-create2-deployer (alias: --no-create2)
	DisableDefaultCreate2Deployer bool

	// Disable pool balance checks.
	//
	// CLI: --disable-pool-balance-checks
	DisablePoolBalanceChecks bool

	// Memory limit per EVM execution in bytes.
	//
	// CLI: --memory-limit
	MemoryLimit uint64

	// Enable printing of traces for executed txs and eth_call.
	//
	// CLI: --print-traces (alias: --enable-trace-printing)
	PrintTraces bool

	// Enable steps tracing for debug/geth-style traces.
	//
	// CLI: --steps-tracing (alias: --tracing)
	StepsTracing bool

	// Enable Celo network features.
	//
	// CLI: --celo
	Celo bool

	// Enable Optimism network features.
	//
	// CLI: --optimism
	Optimism bool
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	c := Config{}
	return &c
}

// SetAccounts sets Accounts, the number of dev accounts to generate and configure.
// A value of 0 omits the flag and lets Anvil use its default (10).
func (c *Config) SetAccounts(n uint) *Config {
	c.Accounts = n
	return c
}

// SetBlockTime sets BlockTime, the block time in seconds for interval mining.
// A value of 0 omits the flag and uses Anvil's instant-mining / default behavior.
func (c *Config) SetBlockTime(seconds uint) *Config {
	c.BlockTime = seconds
	return c
}

// SetBalance sets Balance, the balance of every dev account in Ether.
// An empty string omits the flag and uses Anvil's default of 10000 ETH.
func (c *Config) SetBalance(balance string) *Config {
	c.Balance = balance
	return c
}

// SetConfigOut sets ConfigOut, the path where Anvil writes JSON output for its configuration.
// An empty string omits the flag.
func (c *Config) SetConfigOut(path string) *Config {
	c.ConfigOut = path
	return c
}

// SetDerivationPath sets DerivationPath, the BIP32 derivation path used for deriving accounts.
// If unset, Anvil uses its default "m/44'/60'/0'/0/".
func (c *Config) SetDerivationPath(path string) *Config {
	c.DerivationPath = path
	return c
}

// SetDumpStatePath sets DumpStatePath, the path where Anvil dumps state and block environment on exit.
// If the value is a directory, state is written to "<value>/state.json".
func (c *Config) SetDumpStatePath(path string) *Config {
	c.DumpStatePath = path
	return c
}

// SetHardfork sets Hardfork, the EVM hardfork to use (e.g. "prague", "cancun", "shanghai", "paris", "london", "latest").
// If unset, Anvil uses "latest".
func (c *Config) SetHardfork(hardfork string) *Config {
	c.Hardfork = hardfork
	return c
}

// SetInitPath sets InitPath, the path to a genesis.json file used to initialize the genesis block.
func (c *Config) SetInitPath(path string) *Config {
	c.InitPath = path
	return c
}

// SetIPC enables or disables the IPC server via IPC, and optionally sets IPCPath.
// If IPC is true and IPCPath is empty, Anvil uses its default (/tmp/anvil.ipc).
func (c *Config) SetIPC(enabled bool, path string) *Config {
	c.IPC = enabled
	c.IPCPath = path
	return c
}

// SetIPCEnabled sets IPC to enable or disable the IPC server, without changing IPCPath.
func (c *Config) SetIPCEnabled(enabled bool) *Config {
	c.IPC = enabled
	return c
}

// SetIPCPath sets IPCPath, the path for the IPC server socket.
// IPC must be true for IPCPath to be used; if empty, Anvil uses its default.
func (c *Config) SetIPCPath(path string) *Config {
	c.IPCPath = path
	return c
}

// SetThreads sets Threads, the number of threads to use.
// A value of 0 omits the flag and uses the number of logical cores.
func (c *Config) SetThreads(n uint) *Config {
	c.Threads = n
	return c
}

// SetLoadStatePath sets LoadStatePath, the path to a previously saved state snapshot to initialize the chain.
func (c *Config) SetLoadStatePath(path string) *Config {
	c.LoadStatePath = path
	return c
}

// SetMnemonic sets Mnemonic, the BIP39 mnemonic phrase used for generating accounts.
// This is mutually exclusive with MnemonicRandom and MnemonicSeedUnsafe.
func (c *Config) SetMnemonic(mnemonic string) *Config {
	c.Mnemonic = mnemonic
	return c
}

// SetMaxPersistedStates sets MaxPersistedStates, the maximum number of states to persist on disk.
func (c *Config) SetMaxPersistedStates(n uint) *Config {
	c.MaxPersistedStates = n
	return c
}

// SetMixedMining sets MixedMining, enabling or disabling mixed mining.
func (c *Config) SetMixedMining(enabled bool) *Config {
	c.MixedMining = enabled
	return c
}

// SetMnemonicRandom sets MnemonicRandom, whether Anvil should automatically generate a BIP39 mnemonic.
// This is mutually exclusive with Mnemonic and MnemonicSeedUnsafe.
// If true and MnemonicRandomWords is 0, Anvil uses its default word count (12).
func (c *Config) SetMnemonicRandom(enabled bool) *Config {
	c.MnemonicRandom = enabled
	return c
}

// SetMnemonicRandomWords sets MnemonicRandomWords, the number of BIP39 words to use when MnemonicRandom is true.
func (c *Config) SetMnemonicRandomWords(words uint) *Config {
	c.MnemonicRandomWords = words
	return c
}

// SetMnemonicSeedUnsafe sets MnemonicSeedUnsafe, the seed used to generate a BIP39 mnemonic (UNSAFE).
// This is mutually exclusive with Mnemonic and MnemonicRandom.
func (c *Config) SetMnemonicSeedUnsafe(seed string) *Config {
	c.MnemonicSeedUnsafe = seed
	return c
}

// SetNoMining sets NoMining, disabling auto and interval mining to allow mining on demand only.
func (c *Config) SetNoMining(noMining bool) *Config {
	c.NoMining = noMining
	return c
}

// SetNumber sets Number, the number of the genesis block.
func (c *Config) SetNumber(number uint64) *Config {
	c.Number = number
	return c
}

// SetOrder sets Order, how transactions are sorted in the mempool (e.g. "fees", "fifo").
// If unset, Anvil uses "fees".
func (c *Config) SetOrder(order string) *Config {
	c.Order = order
	return c
}

// SetPort sets Port, the port number Anvil listens on.
// If unset, Anvil uses its default (8545).
func (c *Config) SetPort(port uint) *Config {
	c.Port = port
	return c
}

// SetPreserveHistoricalStates sets PreserveHistoricalStates, controlling whether historical state
// snapshots are preserved when dumping state.
func (c *Config) SetPreserveHistoricalStates(preserve bool) *Config {
	c.PreserveHistoricalStates = preserve
	return c
}

// SetPruneHistory sets PruneHistory, enabling or disabling history pruning.
// When true and PruneHistoryStates > 0, at most that many states are kept in memory.
func (c *Config) SetPruneHistory(enabled bool) *Config {
	c.PruneHistory = enabled
	return c
}

// SetPruneHistoryStates sets PruneHistoryStates, the maximum number of states kept in memory when pruning history.
func (c *Config) SetPruneHistoryStates(n uint) *Config {
	c.PruneHistoryStates = n
	return c
}

// SetStateInterval sets StateInterval, the interval in seconds at which state and block environment
// are periodically dumped to disk.
func (c *Config) SetStateInterval(seconds uint) *Config {
	c.StateInterval = seconds
	return c
}

// SetSlotsInEpoch sets SlotsInEpoch, the number of slots in an epoch.
// If unset, Anvil uses its default (32).
func (c *Config) SetSlotsInEpoch(slots uint) *Config {
	c.SlotsInEpoch = slots
	return c
}

// SetStatePath sets StatePath, a path used as an alias for both load-state and dump-state.
// It initializes the chain with the state stored at the file (if it exists) and dumps state on exit.
func (c *Config) SetStatePath(path string) *Config {
	c.StatePath = path
	return c
}

// SetTimestamp sets Timestamp, the timestamp of the genesis block.
func (c *Config) SetTimestamp(timestamp uint64) *Config {
	c.Timestamp = timestamp
	return c
}

// SetTransactionBlockKeeper sets TransactionBlockKeeper, the number of blocks with transactions
// to keep in memory.
func (c *Config) SetTransactionBlockKeeper(n uint64) *Config {
	c.TransactionBlockKeeper = n
	return c
}

// SetShowLogs sets ShowLogs, enabling or disabling log output.
func (c *Config) SetShowLogs(show bool) *Config {
	c.ShowLogs = show
	return c
}

// SetColor sets Color, the log color mode ("auto", "always", or "never").
func (c *Config) SetColor(color string) *Config {
	c.Color = color
	return c
}

// SetJSONLogs sets JSONLogs, enabling or disabling JSON-formatted log output.
func (c *Config) SetJSONLogs(enabled bool) *Config {
	c.JSONLogs = enabled
	return c
}

// SetMarkdownLogs sets MarkdownLogs, enabling or disabling Markdown-formatted log output.
func (c *Config) SetMarkdownLogs(enabled bool) *Config {
	c.MarkdownLogs = enabled
	return c
}

// SetQuiet sets Quiet, enabling or disabling log output entirely.
func (c *Config) SetQuiet(quiet bool) *Config {
	c.Quiet = quiet
	return c
}

// SetVerbosity sets Verbosity, the verbosity level of log messages.
// A value of 0 omits the flag; values > 0 cause -v to be passed multiple times.
func (c *Config) SetVerbosity(level int) *Config {
	c.Verbosity = level
	return c
}

// SetAllowOrigin sets AllowOrigin, the CORS allow_origin header.
// If unset, Anvil uses its default ("*").
func (c *Config) SetAllowOrigin(origin string) *Config {
	c.AllowOrigin = origin
	return c
}

// SetCachePath sets CachePath, the path to the cache directory where states are stored.
func (c *Config) SetCachePath(path string) *Config {
	c.CachePath = path
	return c
}

// SetHost sets Host, the host address the server will listen on (e.g. "127.0.0.1", "0.0.0.0").
// If unset, Anvil uses its default ("127.0.0.1").
func (c *Config) SetHost(host string) *Config {
	c.Host = host
	return c
}

// SetNoCors sets NoCors, disabling CORS when true.
func (c *Config) SetNoCors(noCors bool) *Config {
	c.NoCors = noCors
	return c
}

// SetNoRequestSizeLimit sets NoRequestSizeLimit, disabling the default request body size limit when true.
func (c *Config) SetNoRequestSizeLimit(noLimit bool) *Config {
	c.NoRequestSizeLimit = noLimit
	return c
}

// SetComputeUnitsPerSecond sets ComputeUnitsPerSecond, the assumed available compute units per second for this provider.
// If unset, Anvil uses its default (330).
func (c *Config) SetComputeUnitsPerSecond(units uint) *Config {
	c.ComputeUnitsPerSecond = units
	return c
}

// SetForkURL sets ForkURL, the remote endpoint URL used to fetch state instead of starting from an empty state.
func (c *Config) SetForkURL(url string) *Config {
	c.ForkURL = url
	return c
}

// SetForkBlockNumber sets ForkBlockNumber, the block number from which to fetch state.
// If negative, the value is subtracted from the latest block.
func (c *Config) SetForkBlockNumber(block int64) *Config {
	c.ForkBlockNumber = block
	return c
}

// SetForkChainID sets ForkChainID, the chain ID to use when forking, skipping fetching it from the remote endpoint.
func (c *Config) SetForkChainID(id uint64) *Config {
	c.ForkChainID = id
	return c
}

// SetForkHeaders sets ForkHeaders, the headers used by the RPC client (e.g. "User-Agent: test-agent").
// These are rendered as multiple --fork-header values.
func (c *Config) SetForkHeaders(headers []string) *Config {
	c.ForkHeaders = headers
	return c
}

// AddForkHeader appends a single header to ForkHeaders.
func (c *Config) AddForkHeader(header string) *Config {
	c.ForkHeaders = append(c.ForkHeaders, header)
	return c
}

// SetForkRetryBackoff sets ForkRetryBackoff, the initial retry backoff duration used
// on errors with the fork provider, as a string in Anvil's expected format.
func (c *Config) SetForkRetryBackoff(backoff string) *Config {
	c.ForkRetryBackoff = backoff
	return c
}

// SetForkTransactionHash sets ForkTransactionHash, the transaction hash after which
// to fetch state when forking.
func (c *Config) SetForkTransactionHash(hash string) *Config {
	c.ForkTransactionHash = hash
	return c
}

// SetNoRateLimit sets NoRateLimit, disabling rate limiting for this node's provider when true.
func (c *Config) SetNoRateLimit(disable bool) *Config {
	c.NoRateLimit = disable
	return c
}

// SetNoStorageCaching sets NoStorageCaching, explicitly disabling RPC storage caching when true.
func (c *Config) SetNoStorageCaching(disable bool) *Config {
	c.NoStorageCaching = disable
	return c
}

// SetRetries sets Retries, the number of retry requests for spurious networks.
// If unset, Anvil uses its default (5).
func (c *Config) SetRetries(retries uint) *Config {
	c.Retries = retries
	return c
}

// SetTimeout sets Timeout, the timeout in milliseconds for requests sent to the remote
// JSON-RPC server in forking mode, as a string (e.g. "45000").
// If unset, Anvil uses its default ("45000").
func (c *Config) SetTimeout(timeout string) *Config {
	c.Timeout = timeout
	return c
}

// SetBlockBaseFeePerGas sets BlockBaseFeePerGas, the base fee in a block,
// expressed as a string to allow decimal or hex (e.g. "1000000000" or "0x3b9aca00").
func (c *Config) SetBlockBaseFeePerGas(fee string) *Config {
	c.BlockBaseFeePerGas = fee
	return c
}

// SetChainID sets ChainID, the chain ID to use.
func (c *Config) SetChainID(id uint64) *Config {
	c.ChainID = id
	return c
}

// SetCodeSizeLimit sets CodeSizeLimit, the contract code size limit in bytes.
// If unset, Anvil uses its default (0x6000).
func (c *Config) SetCodeSizeLimit(limit uint64) *Config {
	c.CodeSizeLimit = limit
	return c
}

// SetDisableBlockGasLimit sets DisableBlockGasLimit, disabling the call.gas_limit <= block.gas_limit constraint when true.
func (c *Config) SetDisableBlockGasLimit(disable bool) *Config {
	c.DisableBlockGasLimit = disable
	return c
}

// SetDisableCodeSizeLimit sets DisableCodeSizeLimit, disabling the EIP-170 contract code size limit when true.
func (c *Config) SetDisableCodeSizeLimit(disable bool) *Config {
	c.DisableCodeSizeLimit = disable
	return c
}

// SetDisableMinPriorityFee sets DisableMinPriorityFee, disabling enforcement of a minimum suggested priority fee when true.
func (c *Config) SetDisableMinPriorityFee(disable bool) *Config {
	c.DisableMinPriorityFee = disable
	return c
}

// SetGasLimit sets GasLimit, the block gas limit.
func (c *Config) SetGasLimit(limit uint64) *Config {
	c.GasLimit = limit
	return c
}

// SetGasPrice sets GasPrice, the gas price to use as a string, allowing decimal or hex.
func (c *Config) SetGasPrice(price string) *Config {
	c.GasPrice = price
	return c
}

// SetAutoImpersonate sets AutoImpersonate, enabling or disabling automatic impersonation (auto-unlock) on startup.
func (c *Config) SetAutoImpersonate(enabled bool) *Config {
	c.AutoImpersonate = enabled
	return c
}

// SetDisableConsoleLog sets DisableConsoleLog, disabling printing of console.log invocations to stdout when true.
func (c *Config) SetDisableConsoleLog(disable bool) *Config {
	c.DisableConsoleLog = disable
	return c
}

// SetDisableDefaultCreate2Deployer sets DisableDefaultCreate2Deployer, disabling the default CREATE2 deployer when true.
func (c *Config) SetDisableDefaultCreate2Deployer(disable bool) *Config {
	c.DisableDefaultCreate2Deployer = disable
	return c
}

// SetDisablePoolBalanceChecks sets DisablePoolBalanceChecks, disabling pool balance checks when true.
func (c *Config) SetDisablePoolBalanceChecks(disable bool) *Config {
	c.DisablePoolBalanceChecks = disable
	return c
}

// SetMemoryLimit sets MemoryLimit, the memory limit per EVM execution in bytes.
func (c *Config) SetMemoryLimit(limit uint64) *Config {
	c.MemoryLimit = limit
	return c
}

// SetPrintTraces sets PrintTraces, enabling printing of traces for executed transactions and eth_call when true.
func (c *Config) SetPrintTraces(enabled bool) *Config {
	c.PrintTraces = enabled
	return c
}

// SetStepsTracing sets StepsTracing, enabling steps tracing for debug/geth-style traces when true.
func (c *Config) SetStepsTracing(enabled bool) *Config {
	c.StepsTracing = enabled
	return c
}

// SetCelo sets Celo, enabling Celo network features when true.
func (c *Config) SetCelo(enabled bool) *Config {
	c.Celo = enabled
	return c
}

// SetOptimism sets Optimism, enabling Optimism network features when true.
func (c *Config) SetOptimism(enabled bool) *Config {
	c.Optimism = enabled
	return c
}
