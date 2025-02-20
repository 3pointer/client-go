// Copyright 2021 TiKV Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: The code in this file is based on code from the
// TiDB project, licensed under the Apache License v 2.0
//
// https://github.com/pingcap/tidb/tree/cc5e161ac06827589c4966674597c137cc9e809c/store/tikv/metrics/shortcuts.go
//

// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import "github.com/prometheus/client_golang/prometheus"

// Shortcuts for performance improvement.
var (
	TxnCmdHistogramWithCommit   prometheus.Observer
	TxnCmdHistogramWithRollback prometheus.Observer
	TxnCmdHistogramWithBatchGet prometheus.Observer
	TxnCmdHistogramWithGet      prometheus.Observer
	TxnCmdHistogramWithLockKeys prometheus.Observer

	RawkvCmdHistogramWithGet           prometheus.Observer
	RawkvCmdHistogramWithBatchGet      prometheus.Observer
	RawkvCmdHistogramWithBatchPut      prometheus.Observer
	RawkvCmdHistogramWithDelete        prometheus.Observer
	RawkvCmdHistogramWithBatchDelete   prometheus.Observer
	RawkvCmdHistogramWithRawScan       prometheus.Observer
	RawkvCmdHistogramWithRawReversScan prometheus.Observer
	RawkvSizeHistogramWithKey          prometheus.Observer
	RawkvSizeHistogramWithValue        prometheus.Observer

	BackoffHistogramRPC                      prometheus.Observer
	BackoffHistogramLock                     prometheus.Observer
	BackoffHistogramLockFast                 prometheus.Observer
	BackoffHistogramPD                       prometheus.Observer
	BackoffHistogramRegionMiss               prometheus.Observer
	BackoffHistogramRegionScheduling         prometheus.Observer
	BackoffHistogramServerBusy               prometheus.Observer
	BackoffHistogramTiKVDiskFull             prometheus.Observer
	BackoffHistogramRegionRecoveryInProgress prometheus.Observer
	BackoffHistogramStaleCmd                 prometheus.Observer
	BackoffHistogramDataNotReady             prometheus.Observer
	BackoffHistogramEmpty                    prometheus.Observer

	TxnRegionsNumHistogramWithSnapshot         prometheus.Observer
	TxnRegionsNumHistogramPrewrite             prometheus.Observer
	TxnRegionsNumHistogramCommit               prometheus.Observer
	TxnRegionsNumHistogramCleanup              prometheus.Observer
	TxnRegionsNumHistogramPessimisticLock      prometheus.Observer
	TxnRegionsNumHistogramPessimisticRollback  prometheus.Observer
	TxnRegionsNumHistogramWithCoprocessor      prometheus.Observer
	TxnRegionsNumHistogramWithBatchCoprocessor prometheus.Observer

	LockResolverCountWithBatchResolve             prometheus.Counter
	LockResolverCountWithExpired                  prometheus.Counter
	LockResolverCountWithNotExpired               prometheus.Counter
	LockResolverCountWithWaitExpired              prometheus.Counter
	LockResolverCountWithResolve                  prometheus.Counter
	LockResolverCountWithResolveForWrite          prometheus.Counter
	LockResolverCountWithResolveAsync             prometheus.Counter
	LockResolverCountWithWriteConflict            prometheus.Counter
	LockResolverCountWithQueryTxnStatus           prometheus.Counter
	LockResolverCountWithQueryTxnStatusCommitted  prometheus.Counter
	LockResolverCountWithQueryTxnStatusRolledBack prometheus.Counter
	LockResolverCountWithQueryCheckSecondaryLocks prometheus.Counter
	LockResolverCountWithResolveLocks             prometheus.Counter
	LockResolverCountWithResolveLockLite          prometheus.Counter

	RegionCacheCounterWithInvalidateRegionFromCacheOK prometheus.Counter
	RegionCacheCounterWithSendFail                    prometheus.Counter
	RegionCacheCounterWithGetRegionByIDOK             prometheus.Counter
	RegionCacheCounterWithGetRegionByIDError          prometheus.Counter
	RegionCacheCounterWithGetCacheMissOK              prometheus.Counter
	RegionCacheCounterWithGetCacheMissError           prometheus.Counter
	RegionCacheCounterWithScanRegionsOK               prometheus.Counter
	RegionCacheCounterWithScanRegionsError            prometheus.Counter
	RegionCacheCounterWithGetStoreOK                  prometheus.Counter
	RegionCacheCounterWithGetStoreError               prometheus.Counter
	RegionCacheCounterWithInvalidateStoreRegionsOK    prometheus.Counter

	LoadRegionCacheHistogramWhenCacheMiss  prometheus.Observer
	LoadRegionCacheHistogramWithRegions    prometheus.Observer
	LoadRegionCacheHistogramWithRegionByID prometheus.Observer
	LoadRegionCacheHistogramWithGetStore   prometheus.Observer

	TxnHeartBeatHistogramOK    prometheus.Observer
	TxnHeartBeatHistogramError prometheus.Observer

	StatusCountWithOK    prometheus.Counter
	StatusCountWithError prometheus.Counter

	SecondaryLockCleanupFailureCounterCommit   prometheus.Counter
	SecondaryLockCleanupFailureCounterRollback prometheus.Counter

	TwoPCTxnCounterOk    prometheus.Counter
	TwoPCTxnCounterError prometheus.Counter

	AsyncCommitTxnCounterOk    prometheus.Counter
	AsyncCommitTxnCounterError prometheus.Counter

	OnePCTxnCounterOk       prometheus.Counter
	OnePCTxnCounterError    prometheus.Counter
	OnePCTxnCounterFallback prometheus.Counter

	BatchRecvHistogramOK    prometheus.Observer
	BatchRecvHistogramError prometheus.Observer

	PrewriteAssertionUsageCounterNone     prometheus.Counter
	PrewriteAssertionUsageCounterExist    prometheus.Counter
	PrewriteAssertionUsageCounterNotExist prometheus.Counter
	PrewriteAssertionUsageCounterUnknown  prometheus.Counter
)

func initShortcuts() {
	TxnCmdHistogramWithCommit = TiKVTxnCmdHistogram.WithLabelValues(LblCommit)
	TxnCmdHistogramWithRollback = TiKVTxnCmdHistogram.WithLabelValues(LblRollback)
	TxnCmdHistogramWithBatchGet = TiKVTxnCmdHistogram.WithLabelValues(LblBatchGet)
	TxnCmdHistogramWithGet = TiKVTxnCmdHistogram.WithLabelValues(LblGet)
	TxnCmdHistogramWithLockKeys = TiKVTxnCmdHistogram.WithLabelValues(LblLockKeys)

	RawkvCmdHistogramWithGet = TiKVRawkvCmdHistogram.WithLabelValues("get")
	RawkvCmdHistogramWithBatchGet = TiKVRawkvCmdHistogram.WithLabelValues("batch_get")
	RawkvCmdHistogramWithBatchPut = TiKVRawkvCmdHistogram.WithLabelValues("batch_put")
	RawkvCmdHistogramWithDelete = TiKVRawkvCmdHistogram.WithLabelValues("delete")
	RawkvCmdHistogramWithBatchDelete = TiKVRawkvCmdHistogram.WithLabelValues("batch_delete")
	RawkvCmdHistogramWithRawScan = TiKVRawkvCmdHistogram.WithLabelValues("raw_scan")
	RawkvCmdHistogramWithRawReversScan = TiKVRawkvCmdHistogram.WithLabelValues("raw_reverse_scan")
	RawkvSizeHistogramWithKey = TiKVRawkvSizeHistogram.WithLabelValues("key")
	RawkvSizeHistogramWithValue = TiKVRawkvSizeHistogram.WithLabelValues("value")

	BackoffHistogramRPC = TiKVBackoffHistogram.WithLabelValues("tikvRPC")
	BackoffHistogramLock = TiKVBackoffHistogram.WithLabelValues("txnLock")
	BackoffHistogramLockFast = TiKVBackoffHistogram.WithLabelValues("tikvLockFast")
	BackoffHistogramPD = TiKVBackoffHistogram.WithLabelValues("pdRPC")
	BackoffHistogramRegionMiss = TiKVBackoffHistogram.WithLabelValues("regionMiss")
	BackoffHistogramRegionScheduling = TiKVBackoffHistogram.WithLabelValues("regionScheduling")
	BackoffHistogramServerBusy = TiKVBackoffHistogram.WithLabelValues("serverBusy")
	BackoffHistogramTiKVDiskFull = TiKVBackoffHistogram.WithLabelValues("tikvDiskFull")
	BackoffHistogramRegionRecoveryInProgress = TiKVBackoffHistogram.WithLabelValues("regionRecoveryInProgress")
	BackoffHistogramStaleCmd = TiKVBackoffHistogram.WithLabelValues("staleCommand")
	BackoffHistogramDataNotReady = TiKVBackoffHistogram.WithLabelValues("dataNotReady")
	BackoffHistogramEmpty = TiKVBackoffHistogram.WithLabelValues("")

	TxnRegionsNumHistogramWithSnapshot = TiKVTxnRegionsNumHistogram.WithLabelValues("snapshot")
	TxnRegionsNumHistogramPrewrite = TiKVTxnRegionsNumHistogram.WithLabelValues("2pc_prewrite")
	TxnRegionsNumHistogramCommit = TiKVTxnRegionsNumHistogram.WithLabelValues("2pc_commit")
	TxnRegionsNumHistogramCleanup = TiKVTxnRegionsNumHistogram.WithLabelValues("2pc_cleanup")
	TxnRegionsNumHistogramPessimisticLock = TiKVTxnRegionsNumHistogram.WithLabelValues("2pc_pessimistic_lock")
	TxnRegionsNumHistogramPessimisticRollback = TiKVTxnRegionsNumHistogram.WithLabelValues("2pc_pessimistic_rollback")
	TxnRegionsNumHistogramWithCoprocessor = TiKVTxnRegionsNumHistogram.WithLabelValues("coprocessor")
	TxnRegionsNumHistogramWithBatchCoprocessor = TiKVTxnRegionsNumHistogram.WithLabelValues("batch_coprocessor")

	LockResolverCountWithBatchResolve = TiKVLockResolverCounter.WithLabelValues("batch_resolve")
	LockResolverCountWithExpired = TiKVLockResolverCounter.WithLabelValues("expired")
	LockResolverCountWithNotExpired = TiKVLockResolverCounter.WithLabelValues("not_expired")
	LockResolverCountWithWaitExpired = TiKVLockResolverCounter.WithLabelValues("wait_expired")
	LockResolverCountWithResolve = TiKVLockResolverCounter.WithLabelValues("resolve")
	LockResolverCountWithResolveForWrite = TiKVLockResolverCounter.WithLabelValues("resolve_for_write")
	LockResolverCountWithResolveAsync = TiKVLockResolverCounter.WithLabelValues("resolve_async_commit")
	LockResolverCountWithWriteConflict = TiKVLockResolverCounter.WithLabelValues("write_conflict")
	LockResolverCountWithQueryTxnStatus = TiKVLockResolverCounter.WithLabelValues("query_txn_status")
	LockResolverCountWithQueryTxnStatusCommitted = TiKVLockResolverCounter.WithLabelValues("query_txn_status_committed")
	LockResolverCountWithQueryTxnStatusRolledBack = TiKVLockResolverCounter.WithLabelValues("query_txn_status_rolled_back")
	LockResolverCountWithQueryCheckSecondaryLocks = TiKVLockResolverCounter.WithLabelValues("query_check_secondary_locks")
	LockResolverCountWithResolveLocks = TiKVLockResolverCounter.WithLabelValues("query_resolve_locks")
	LockResolverCountWithResolveLockLite = TiKVLockResolverCounter.WithLabelValues("query_resolve_lock_lite")

	RegionCacheCounterWithInvalidateRegionFromCacheOK = TiKVRegionCacheCounter.WithLabelValues("invalidate_region_from_cache", "ok")
	RegionCacheCounterWithSendFail = TiKVRegionCacheCounter.WithLabelValues("send_fail", "ok")
	RegionCacheCounterWithGetRegionByIDOK = TiKVRegionCacheCounter.WithLabelValues("get_region_by_id", "ok")
	RegionCacheCounterWithGetRegionByIDError = TiKVRegionCacheCounter.WithLabelValues("get_region_by_id", "err")
	RegionCacheCounterWithGetCacheMissOK = TiKVRegionCacheCounter.WithLabelValues("get_region_when_miss", "ok")
	RegionCacheCounterWithGetCacheMissError = TiKVRegionCacheCounter.WithLabelValues("get_region_when_miss", "err")
	RegionCacheCounterWithScanRegionsOK = TiKVRegionCacheCounter.WithLabelValues("scan_regions", "ok")
	RegionCacheCounterWithScanRegionsError = TiKVRegionCacheCounter.WithLabelValues("scan_regions", "err")
	RegionCacheCounterWithGetStoreOK = TiKVRegionCacheCounter.WithLabelValues("get_store", "ok")
	RegionCacheCounterWithGetStoreError = TiKVRegionCacheCounter.WithLabelValues("get_store", "err")
	RegionCacheCounterWithInvalidateStoreRegionsOK = TiKVRegionCacheCounter.WithLabelValues("invalidate_store_regions", "ok")

	LoadRegionCacheHistogramWhenCacheMiss = TiKVLoadRegionCacheHistogram.WithLabelValues("get_region_when_miss")
	LoadRegionCacheHistogramWithRegionByID = TiKVLoadRegionCacheHistogram.WithLabelValues("get_region_by_id")
	LoadRegionCacheHistogramWithRegions = TiKVLoadRegionCacheHistogram.WithLabelValues("scan_regions")
	LoadRegionCacheHistogramWithGetStore = TiKVLoadRegionCacheHistogram.WithLabelValues("get_store")

	TxnHeartBeatHistogramOK = TiKVTxnHeartBeatHistogram.WithLabelValues("ok")
	TxnHeartBeatHistogramError = TiKVTxnHeartBeatHistogram.WithLabelValues("err")

	StatusCountWithOK = TiKVStatusCounter.WithLabelValues("ok")
	StatusCountWithError = TiKVStatusCounter.WithLabelValues("err")

	SecondaryLockCleanupFailureCounterCommit = TiKVSecondaryLockCleanupFailureCounter.WithLabelValues("commit")
	SecondaryLockCleanupFailureCounterRollback = TiKVSecondaryLockCleanupFailureCounter.WithLabelValues("rollback")

	TwoPCTxnCounterOk = TiKVTwoPCTxnCounter.WithLabelValues("ok")
	TwoPCTxnCounterError = TiKVTwoPCTxnCounter.WithLabelValues("err")

	AsyncCommitTxnCounterOk = TiKVAsyncCommitTxnCounter.WithLabelValues("ok")
	AsyncCommitTxnCounterError = TiKVAsyncCommitTxnCounter.WithLabelValues("err")

	OnePCTxnCounterOk = TiKVOnePCTxnCounter.WithLabelValues("ok")
	OnePCTxnCounterError = TiKVOnePCTxnCounter.WithLabelValues("err")
	OnePCTxnCounterFallback = TiKVOnePCTxnCounter.WithLabelValues("fallback")

	BatchRecvHistogramOK = TiKVBatchRecvLatency.WithLabelValues("ok")
	BatchRecvHistogramError = TiKVBatchRecvLatency.WithLabelValues("err")

	PrewriteAssertionUsageCounterNone = TiKVPrewriteAssertionUsageCounter.WithLabelValues("none")
	PrewriteAssertionUsageCounterExist = TiKVPrewriteAssertionUsageCounter.WithLabelValues("exist")
	PrewriteAssertionUsageCounterNotExist = TiKVPrewriteAssertionUsageCounter.WithLabelValues("not-exist")
	PrewriteAssertionUsageCounterUnknown = TiKVPrewriteAssertionUsageCounter.WithLabelValues("unknown")
}
