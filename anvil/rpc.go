package anvil

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
)

// Send transactions impersonating an externally owned account or contract.
func (a Anvil) ImpersonateAccount(address common.Address) error {
	_, err := makeRequest[any](a.url, "anvil_impersonateAccount", []any{address})
	return err
}

// Stops impersonating an account or contract if previously set with ImpersonateAccount
func (a Anvil) StopImpersonatingAccount(address common.Address) error {
	_, err := makeRequest[bool](a.url, "anvil_stopImpersonatingAccount", []any{address})
	return err
}

// AutoImpersonateAccount accepts true to enable auto impersonation of accounts, and false to disable it.
// When enabled, any transaction's sender will be automatically impersonated (same effect as impersonateAccount).
func (a Anvil) AutoImpersonateAccount(enabled bool) error {
	_, err := makeRequest[any](a.url, "anvil_autoImpersonateAccount", []any{enabled})
	return err
}

// GetAutomine returns true if automatic mining is enabled, and false otherwise.
func (a Anvil) GetAutomine() (bool, error) {
	res, err := makeRequest[bool](a.url, "anvil_getAutomine", []any{})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// GetBlobByHash returns the blob for a given KZG commitment versioned hash.
func (a Anvil) GetBlobByHash(hash common.Hash) (string, error) {
	// Return as hex string for flexibility.
	res, err := makeRequest[string](a.url, "anvil_getBlobByHash", []any{hash})
	if err != nil {
		return "", err
	}
	return *res, nil
}

// GetBlobsByTransactionHash returns the blobs for a given transaction hash.
func (a Anvil) GetBlobsByTransactionHash(txHash common.Hash) ([]string, error) {
	res, err := makeRequest[[]string](a.url, "anvil_getBlobsByTransactionHash", []any{txHash})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetBlobSidecarsByBlockId returns the blob sidecars for a given block id.
// Block ID can be a block number, hash, or tag like "latest".
func (a Anvil) GetBlobSidecarsByBlockId(blockId string) (json.RawMessage, error) {
	res, err := makeRequest[json.RawMessage](a.url, "anvil_getBlobSidecarsByBlockId", []any{blockId})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetBlobsByBlockId returns blobs for a given block ID, optionally filtered by a list of versioned hashes.
func (a Anvil) GetBlobsByBlockId(blockId string, versionedHashes []common.Hash) (json.RawMessage, error) {
	params := []any{blockId}
	if len(versionedHashes) > 0 {
		params = append(params, versionedHashes)
	}
	res, err := makeRequest[json.RawMessage](a.url, "anvil_getBlobsByBlockId", params)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// Mine mines a series of blocks.
// If blocks or interval are nil, the defaults (1 block, 1 second) are used by Anvil.
func (a Anvil) Mine(blocks, interval *big.Int) error {
	var params []any
	if blocks != nil {
		params = append(params, toHexQuantityBig(blocks))
	}
	if interval != nil {
		params = append(params, toHexQuantityBig(interval))
	}
	_, err := makeRequest[any](a.url, "anvil_mine", params)
	return err
}

// DropTransaction removes a transaction from the pool and may return the dropped hash.
func (a Anvil) DropTransaction(txHash common.Hash) (*common.Hash, error) {
	res, err := makeRequest[*common.Hash](a.url, "anvil_dropTransaction", []any{txHash})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// Reset resets the fork to a fresh forked state, and optionally updates the fork config.
// Pass nil to disable forking entirely.
func (a Anvil) Reset(forkConfig any) error {
	var params []any
	if forkConfig != nil {
		params = append(params, forkConfig)
	}
	_, err := makeRequest[any](a.url, "anvil_reset", params)
	return err
}

// SetRpcUrl sets the backend RPC URL used for forking.
func (a Anvil) SetRpcUrl(url string) error {
	_, err := makeRequest[any](a.url, "anvil_setRpcUrl", []any{url})
	return err
}

// SetBalance modifies the balance of an account.
func (a Anvil) SetBalance(address common.Address, balance *big.Int) error {
	_, err := makeRequest[any](a.url, "anvil_setBalance", []any{address, toHexQuantityBig(balance)})
	return err
}

// SetCode sets the code of a contract.
func (a Anvil) SetCode(address common.Address, codeHex string) error {
	// codeHex should be 0x-prefixed bytecode.
	_, err := makeRequest[any](a.url, "anvil_setCode", []any{address, codeHex})
	return err
}

// SetNonce sets the nonce of an address.
func (a Anvil) SetNonce(address common.Address, nonce uint64) error {
	_, err := makeRequest[any](a.url, "anvil_setNonce", []any{address, toHexQuantityUint64(nonce)})
	return err
}

// SetStorageAt writes a single storage slot of the account's storage.
func (a Anvil) SetStorageAt(address common.Address, slot common.Hash, value common.Hash) (bool, error) {
	res, err := makeRequest[bool](a.url, "anvil_setStorageAt", []any{address, slot, value})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// SetCoinbase sets the coinbase (block author) address.
func (a Anvil) SetCoinbase(address common.Address) error {
	_, err := makeRequest[any](a.url, "anvil_setCoinbase", []any{address})
	return err
}

// SetLoggingEnabled enables or disables logging.
func (a Anvil) SetLoggingEnabled(enabled bool) error {
	_, err := makeRequest[any](a.url, "anvil_setLoggingEnabled", []any{enabled})
	return err
}

// SetMinGasPrice sets the minimum gas price for the node.
func (a Anvil) SetMinGasPrice(price *big.Int) error {
	_, err := makeRequest[any](a.url, "anvil_setMinGasPrice", []any{toHexQuantityBig(price)})
	return err
}

// SetNextBlockBaseFeePerGas sets the base fee of the next block.
func (a Anvil) SetNextBlockBaseFeePerGas(baseFee *big.Int) error {
	_, err := makeRequest[any](a.url, "anvil_setNextBlockBaseFeePerGas", []any{toHexQuantityBig(baseFee)})
	return err
}

// SetChainID sets the chain ID of the current EVM instance.
func (a Anvil) SetChainID(chainID uint64) error {
	_, err := makeRequest[any](a.url, "anvil_setChainId", []any{toHexQuantityUint64(chainID)})
	return err
}

// DumpState returns a hex string representing the complete state of the chain.
// It can be re-imported into a fresh instance of Anvil to restore the same state.
func (a Anvil) DumpState() (string, error) {
	res, err := makeRequest[string](a.url, "anvil_dumpState", []any{})
	if err != nil {
		return "", err
	}
	return *res, nil
}

// LoadState merges a state snapshot previously returned by DumpState into the current chain state.
// Colliding accounts or storage slots will be overwritten.
func (a Anvil) LoadState(stateHex string) (bool, error) {
	res, err := makeRequest[bool](a.url, "anvil_loadState", []any{stateHex})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// NodeInfo retrieves the configuration parameters for the currently running Anvil node.
func (a Anvil) NodeInfo() (map[string]any, error) {
	res, err := makeRequest[map[string]any](a.url, "anvil_nodeInfo", []any{})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// EvmSetAutomine enables or disables automatic mining of new blocks for each new transaction.
// If disabled, Anvil mines according to the configured interval; if enabled, blocks are mined
// only when transactions arrive.
func (a Anvil) EvmSetAutomine(enabled bool) error {
	_, err := makeRequest[any](a.url, "evm_setAutomine", []any{enabled})
	return err
}

// EvmSetIntervalMining sets the mining behavior to interval mode with the given interval in seconds.
func (a Anvil) EvmSetIntervalMining(intervalSeconds uint64) error {
	_, err := makeRequest[any](a.url, "evm_setIntervalMining", []any{toHexQuantityUint64(intervalSeconds)})
	return err
}

// EvmSnapshot snapshots the state of the blockchain at the current block and returns a snapshot id.
func (a Anvil) EvmSnapshot() (string, error) {
	res, err := makeRequest[string](a.url, "evm_snapshot", []any{})
	if err != nil {
		return "", err
	}
	return *res, nil
}

// EvmRevert reverts the state of the blockchain to a previous snapshot id.
func (a Anvil) EvmRevert(snapshotID string) (bool, error) {
	res, err := makeRequest[bool](a.url, "evm_revert", []any{snapshotID})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// EvmIncreaseTime jumps forward in time by the given number of seconds.
func (a Anvil) EvmIncreaseTime(seconds int64) (int64, error) {
	res, err := makeRequest[int64](a.url, "evm_increaseTime", []any{seconds})
	if err != nil {
		return 0, err
	}
	return *res, nil
}

// EvmSetNextBlockTimestamp sets the exact timestamp to use for the next block.
func (a Anvil) EvmSetNextBlockTimestamp(timestamp uint64) error {
	_, err := makeRequest[any](a.url, "evm_setNextBlockTimestamp", []any{toHexQuantityUint64(timestamp)})
	return err
}

// SetBlockTimestampInterval sets a block timestamp interval; the next block timestamp is
// computed as lastBlockTimestamp + interval.
func (a Anvil) SetBlockTimestampInterval(intervalSeconds uint64) error {
	_, err := makeRequest[any](a.url, "anvil_setBlockTimestampInterval", []any{toHexQuantityUint64(intervalSeconds)})
	return err
}

// EvmSetBlockGasLimit sets the block gas limit for following blocks.
func (a Anvil) EvmSetBlockGasLimit(limit *big.Int) error {
	_, err := makeRequest[any](a.url, "evm_setBlockGasLimit", []any{toHexQuantityBig(limit)})
	return err
}

// RemoveBlockTimestampInterval removes a previously set block timestamp interval, if it exists.
func (a Anvil) RemoveBlockTimestampInterval() (bool, error) {
	res, err := makeRequest[bool](a.url, "anvil_removeBlockTimestampInterval", []any{})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// EvmMine mines a single block. If a timestamp is provided, that timestamp is used for the block.
func (a Anvil) EvmMine(timestamp ...uint64) error {
	params := []any{}
	if len(timestamp) > 0 {
		params = append(params, toHexQuantityUint64(timestamp[0]))
	}
	_, err := makeRequest[any](a.url, "evm_mine", params)
	return err
}

// EnableTraces turns on call traces for transactions returned to the user instead of just tx hash/receipt.
func (a Anvil) EnableTraces() error {
	_, err := makeRequest[any](a.url, "anvil_enableTraces", []any{})
	return err
}

// SendUnsignedTransaction executes a transaction regardless of signature status.
// tx is a standard transaction object encoded as a map, similar to eth_sendTransaction.
func (a Anvil) SendUnsignedTransaction(tx map[string]any) (common.Hash, error) {
	res, err := makeRequest[common.Hash](a.url, "eth_sendUnsignedTransaction", []any{tx})
	if err != nil {
		return common.Hash{}, err
	}
	return *res, nil
}

// TxpoolStatusResult is the standard result of txpool_status:
// number of pending and queued txs, encoded as hex quantities.
type TxpoolStatusResult struct {
	Pending string `json:"pending"`
	Queued  string `json:"queued"`
}

// TxpoolStatus returns the number of transactions currently pending and queued in the txpool.
func (a Anvil) TxpoolStatus() (TxpoolStatusResult, error) {
	res, err := makeRequest[TxpoolStatusResult](a.url, "txpool_status", []any{})
	if err != nil {
		return TxpoolStatusResult{}, err
	}
	return *res, nil
}

// TxpoolInspect returns a human-readable summary of transactions currently pending and queued.
// The exact shape is nested maps as in Geth; we expose it as generic JSON.
func (a Anvil) TxpoolInspect() (map[string]any, error) {
	res, err := makeRequest[map[string]any](a.url, "txpool_inspect", []any{})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// TxpoolContent returns the details of all transactions currently pending and queued.
func (a Anvil) TxpoolContent() (map[string]any, error) {
	res, err := makeRequest[map[string]any](a.url, "txpool_content", []any{})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSGetApiLevel returns the Otterscan API level (simple version number).
func (a Anvil) OTSGetApiLevel() (uint64, error) {
	res, err := makeRequest[uint64](a.url, "ots_getApiLevel", []any{})
	if err != nil {
		return 0, err
	}
	return *res, nil
}

// OTSInternalOperation is a simplified representation of an internal operation from ots_getInternalOperations.
type OTSInternalOperation struct {
	Type uint8          `json:"type"`
	From common.Address `json:"from"`
	To   common.Address `json:"to"`
	// Value is the amount of ETH transferred, as a hex quantity string.
	Value string `json:"value"`
}

// OTSGetInternalOperations returns internal ETH transfers for a transaction.
func (a Anvil) OTSGetInternalOperations(txHash common.Hash) ([]OTSInternalOperation, error) {
	res, err := makeRequest[[]OTSInternalOperation](a.url, "ots_getInternalOperations", []any{txHash.Hex()})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSHasCode checks if an address contains deployed code at a specific block (or "latest").
func (a Anvil) OTSHasCode(address common.Address, blockTag string) (bool, error) {
	res, err := makeRequest[bool](a.url, "ots_hasCode", []any{address.Hex(), blockTag})
	if err != nil {
		return false, err
	}
	return *res, nil
}

// OTSGetTransactionError returns the raw revert data for a transaction, or "0x" on success / no-reason failure.
func (a Anvil) OTSGetTransactionError(txHash common.Hash) (string, error) {
	res, err := makeRequest[string](a.url, "ots_getTransactionError", []any{txHash.Hex()})
	if err != nil {
		return "", err
	}
	return *res, nil
}

// OTSTraceTransaction returns a call tree trace for a transaction.
func (a Anvil) OTSTraceTransaction(txHash common.Hash) (json.RawMessage, error) {
	res, err := makeRequest[json.RawMessage](a.url, "ots_traceTransaction", []any{txHash.Hex()})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSGetBlockDetails returns a tailored block object for a given block number.
func (a Anvil) OTSGetBlockDetails(blockNumber uint64) (json.RawMessage, error) {
	res, err := makeRequest[json.RawMessage](a.url, "ots_getBlockDetails", []any{toHexQuantityUint64(blockNumber)})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSGetBlockTransactions returns paginated transaction + receipt data for a given block.
func (a Anvil) OTSGetBlockTransactions(blockNumber uint64, pageSize uint64) (json.RawMessage, error) {
	params := []any{toHexQuantityUint64(blockNumber), pageSize}
	res, err := makeRequest[json.RawMessage](a.url, "ots_getBlockTransactions", params)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSSearchTransactionsBefore searches paginated inbound/outbound/internal transactions for an address before a block.
func (a Anvil) OTSSearchTransactionsBefore(address common.Address, blockNumber uint64, pageSize uint64) (json.RawMessage, error) {
	params := []any{address.Hex(), blockNumber, pageSize}
	res, err := makeRequest[json.RawMessage](a.url, "ots_searchTransactionsBefore", params)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSSearchTransactionsAfter searches paginated inbound/outbound/internal transactions for an address after a block.
func (a Anvil) OTSSearchTransactionsAfter(address common.Address, blockNumber uint64, pageSize uint64) (json.RawMessage, error) {
	params := []any{address.Hex(), blockNumber, pageSize}
	res, err := makeRequest[json.RawMessage](a.url, "ots_searchTransactionsAfter", params)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// OTSGetTransactionBySenderAndNonce returns the transaction hash for a given sender and nonce, or "" if not found.
func (a Anvil) OTSGetTransactionBySenderAndNonce(sender common.Address, nonce uint64) (string, error) {
	params := []any{sender.Hex(), nonce}
	res, err := makeRequest[string](a.url, "ots_getTransactionBySenderAndNonce", params)
	if err != nil {
		return "", err
	}
	return *res, nil
}

// OTSContractCreator is the result of ots_getContractCreator.
type OTSContractCreator struct {
	Hash    common.Hash    `json:"hash"`
	Creator common.Address `json:"creator"`
}

// OTSGetContractCreator returns the tx hash and creator address that deployed a contract,
// or (nil, nil) if the address is not a contract.
func (a Anvil) OTSGetContractCreator(address common.Address) (*OTSContractCreator, error) {
	res, err := makeRequest[*OTSContractCreator](a.url, "ots_getContractCreator", []any{address.Hex()})
	if err != nil {
		return nil, err
	}
	return *res, nil
}

func makeRequest[T any](url string, method string, params []any) (*T, error) {
	type Response struct {
		Result T `json:"result"`
	}

	client := resty.New().SetJSONMarshaler(json.Marshal).SetJSONUnmarshaler(json.Unmarshal)

	resp, err := client.R().
		SetBody(map[string]any{
			"method":  method,
			"params":  params,
			"id":      1,
			"jsonrpc": "2.0",
		}).
		Post(url)
	if err != nil {
		return nil, err
	}

	defer resp.RawResponse.Body.Close()

	var response Response
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// toHexQuantity turns a uint64 into a 0x-prefixed hex quantity string.
func toHexQuantityUint64(v uint64) string {
	return fmt.Sprintf("0x%x", v)
}

// toHexQuantityBig turns a *big.Int into a 0x-prefixed hex quantity string.
func toHexQuantityBig(v *big.Int) string {
	if v == nil {
		return "0x0"
	}
	return "0x" + v.Text(16)
}
