package anvil

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestAnvil_Impersonation shows basic account impersonation control.
func TestAnvil_Impersonation(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	addr := Account0

	// anvil_impersonateAccount
	if err := anvl.ImpersonateAccount(addr); err != nil {
		t.Fatalf("ImpersonateAccount failed: %v", err)
	}

	// anvil_autoImpersonateAccount
	if err := anvl.AutoImpersonateAccount(true); err != nil {
		t.Fatalf("AutoImpersonateAccount(true) failed: %v", err)
	}

	// anvil_stopImpersonatingAccount
	if err := anvl.StopImpersonatingAccount(addr); err != nil {
		t.Fatalf("StopImpersonatingAccount failed: %v", err)
	}
}

// TestAnvil_MiningAndTime exercises automine toggles, intervals, and manual mining.
func TestAnvil_MiningAndTime(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	// Query current automine status
	auto, err := anvl.GetAutomine()
	if err != nil {
		t.Fatalf("GetAutomine failed: %v", err)
	}
	t.Logf("automine: %v", auto)

	// Disable EVM automine and mine every 5 seconds
	if err := anvl.EvmSetAutomine(false); err != nil {
		t.Fatalf("EvmSetAutomine(false) failed: %v", err)
	}
	if err := anvl.EvmSetIntervalMining(5); err != nil {
		t.Fatalf("EvmSetIntervalMining failed: %v", err)
	}

	// Mine 3 blocks with 0-second interval via anvil_mine
	if err := anvl.Mine(big.NewInt(3), big.NewInt(0)); err != nil {
		t.Fatalf("Mine failed: %v", err)
	}

	// Single block via evm_mine
	if err := anvl.EvmMine(); err != nil {
		t.Fatalf("EvmMine failed: %v", err)
	}

	// Jump time 1 hour forward
	if _, err := anvl.EvmIncreaseTime(3600); err != nil {
		t.Fatalf("EvmIncreaseTime failed: %v", err)
	}

	// Set next block timestamp explicitly
	if err := anvl.EvmSetNextBlockTimestamp(1_700_000_000); err != nil {
		t.Fatalf("EvmSetNextBlockTimestamp failed: %v", err)
	}

	// Block timestamp interval
	if err := anvl.SetBlockTimestampInterval(12); err != nil {
		t.Fatalf("SetBlockTimestampInterval failed: %v", err)
	}
	if ok, err := anvl.RemoveBlockTimestampInterval(); err != nil {
		t.Fatalf("RemoveBlockTimestampInterval failed: %v", err)
	} else {
		t.Logf("RemoveBlockTimestampInterval ok=%v", ok)
	}
}

// TestAnvil_StateMutations exercises balance, code, nonce, storage, chain & gas config changes.
func TestAnvil_StateMutations(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	addr := Account0

	// Balance
	if err := anvl.SetBalance(addr, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18))); err != nil {
		t.Fatalf("SetBalance failed: %v", err)
	}

	// Code (simple 0x6001... bytecode)
	if err := anvl.SetCode(addr, "0x6001600101"); err != nil {
		t.Fatalf("SetCode failed: %v", err)
	}

	// Nonce
	if err := anvl.SetNonce(addr, 42); err != nil {
		t.Fatalf("SetNonce failed: %v", err)
	}

	// Storage
	slot := common.HexToHash("0x0")
	val := common.HexToHash("0x1")
	if ok, err := anvl.SetStorageAt(addr, slot, val); err != nil {
		t.Fatalf("SetStorageAt failed: %v", err)
	} else if !ok {
		t.Fatalf("SetStorageAt returned false")
	}

	// Coinbase
	if err := anvl.SetCoinbase(addr); err != nil {
		t.Fatalf("SetCoinbase failed: %v", err)
	}

	// Logging + traces
	if err := anvl.SetLoggingEnabled(true); err != nil {
		t.Fatalf("SetLoggingEnabled(true) failed: %v", err)
	}
	if err := anvl.EnableTraces(); err != nil {
		t.Fatalf("EnableTraces failed: %v", err)
	}

	// Gas-related
	if err := anvl.SetMinGasPrice(big.NewInt(1e9)); err != nil { // 1 gwei
		t.Fatalf("SetMinGasPrice failed: %v", err)
	}
	if err := anvl.SetNextBlockBaseFeePerGas(big.NewInt(2e9)); err != nil {
		t.Fatalf("SetNextBlockBaseFeePerGas failed: %v", err)
	}
	if err := anvl.EvmSetBlockGasLimit(big.NewInt(30_000_000)); err != nil {
		t.Fatalf("EvmSetBlockGasLimit failed: %v", err)
	}

	// Chain ID
	if err := anvl.SetChainID(31337); err != nil {
		t.Fatalf("SetChainID failed: %v", err)
	}
}

// TestAnvil_SnapshotsAndForking exercises snapshots, reset and dump/load state.
func TestAnvil_SnapshotsAndForking(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	// Snapshot / revert
	snapID, err := anvl.EvmSnapshot()
	if err != nil {
		t.Fatalf("EvmSnapshot failed: %v", err)
	}
	if ok, err := anvl.EvmRevert(snapID); err != nil {
		t.Fatalf("EvmRevert failed: %v", err)
	} else if !ok {
		t.Fatalf("EvmRevert returned false")
	}

	// Dump / load state
	stateHex, err := anvl.DumpState()
	if err != nil {
		t.Fatalf("DumpState failed: %v", err)
	}
	if ok, err := anvl.LoadState(stateHex); err != nil {
		t.Fatalf("LoadState failed: %v", err)
	} else if !ok {
		t.Fatalf("LoadState returned false")
	}

	// Reset fork (no new config, disables forking if it was enabled)
	if err := anvl.Reset(nil); err != nil {
		t.Fatalf("Reset(nil) failed: %v", err)
	}

	// Change backend RPC URL
	if err := anvl.SetRpcUrl("https://rpc.example"); err != nil {
		t.Fatalf("SetRpcUrl failed: %v", err)
	}

	info, err := anvl.NodeInfo()
	if err != nil {
		t.Fatalf("NodeInfo failed: %v", err)
	}
	t.Logf("NodeInfo chainId: %v", info["chainId"])
}

// TestAnvil_Blobs exercises blob-related helpers.
func TestAnvil_Blobs(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	// Dummy hashes – these will likely not exist; we just assert the calls don't error badly.
	blobHash := common.HexToHash("0x" + "00")
	txHash := common.HexToHash("0x" + "01")

	blob, err := anvl.GetBlobByHash(blobHash)
	if err != nil {
		t.Fatalf("GetBlobByHash failed: %v", err)
	}
	t.Logf("blob length: %d", len(blob))

	blobs, err := anvl.GetBlobsByTransactionHash(txHash)
	if err != nil {
		t.Fatalf("GetBlobsByTransactionHash failed: %v", err)
	}
	t.Logf("tx blobs count: %d", len(blobs))

	sidecars, err := anvl.GetBlobSidecarsByBlockId("latest")
	if err != nil {
		t.Fatalf("GetBlobSidecarsByBlockId failed: %v", err)
	}
	t.Logf("sidecars json length: %d", len(sidecars))

	blockBlobs, err := anvl.GetBlobsByBlockId("latest", nil)
	if err != nil {
		t.Fatalf("GetBlobsByBlockId failed: %v", err)
	}
	t.Logf("block blobs json length: %d", len(blockBlobs))
}

// TestAnvil_UnsignedTx exercises sending an unsigned tx.
func TestAnvil_UnsignedTx(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	tx := map[string]any{
		"from":  Account0.Hex(),
		"to":    common.HexToAddress("0x0000000000000000000000000000000000000001").Hex(),
		"value": "0x1",
	}

	hash, err := anvl.SendUnsignedTransaction(tx)
	if err != nil {
		t.Fatalf("SendUnsignedTransaction failed: %v", err)
	}

	if (hash == common.Hash{}) {
		t.Fatalf("SendUnsignedTransaction returned empty hash")
	}
}

// TestAnvil_Txpool exercises the txpool methods.
func TestAnvil_Txpool(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	status, err := anvl.TxpoolStatus()
	if err != nil {
		t.Fatalf("TxpoolStatus failed: %v", err)
	}
	t.Logf("pending: %s queued: %s", status.Pending, status.Queued)

	inspect, err := anvl.TxpoolInspect()
	if err != nil {
		t.Fatalf("TxpoolInspect failed: %v", err)
	}
	if inspect == nil {
		t.Fatalf("TxpoolInspect returned nil map")
	}
	t.Logf("inspect keys: %v", inspect)

	content, err := anvl.TxpoolContent()
	if err != nil {
		t.Fatalf("TxpoolContent failed: %v", err)
	}
	if content == nil {
		t.Fatalf("TxpoolContent returned nil map")
	}
	t.Logf("content keys: %v", content)
}

// TestAnvil_Otterscan exercises a subset of the ots_* helpers.
func TestAnvil_Otterscan(t *testing.T) {
	anvl, err := NewWithConfig(NewConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer anvl.Close()

	apiLevel, err := anvl.OTSGetApiLevel()
	if err != nil {
		t.Fatalf("OTSGetApiLevel failed: %v", err)
	}
	t.Logf("ots api level: %d", apiLevel)

	addr := Account0
	hasCode, err := anvl.OTSHasCode(addr, "latest")
	if err != nil {
		t.Fatalf("OTSHasCode failed: %v", err)
	}
	t.Logf("ots has code for %s: %v", addr.Hex(), hasCode)

	// The remaining OTS calls depend heavily on real txs/blocks,
	// so we only check that they don't panic / error when given dummy input.
	// You can add more assertions when wiring to a real chain.
}
