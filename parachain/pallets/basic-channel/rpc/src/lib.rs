use std::sync::Arc;

use itertools::zip;

use codec::{Codec,Decode};
use jsonrpc_core::{Result, Error as JsonError};
use jsonrpc_derive::rpc;
use sc_rpc::DenyUnsafe;
use sp_core::{H256, offchain::OffchainStorage};
use parking_lot::RwLock;

use artemis_basic_channel::outbound::{CommitmentData, generate_merkle_proofs, offchain_key};

type Proofs<TAccountId> = Vec<(TAccountId, Vec<u8>)>;

#[rpc]
pub trait BasicChannelApi<TAccountId>
{
	#[rpc(name = "get_merkle_proofs")]
	fn get_merkle_proofs(&self, root: H256) -> Result<Proofs<TAccountId>>;
}

pub struct BasicChannel<TStorage, TAccountId> {
	_marker: std::marker::PhantomData<TAccountId>,
	/// Offchain storage
	storage: Arc<RwLock<TStorage>>,
	/// Standard Substrate RPC check
	deny_unsafe: DenyUnsafe,
	/// Used for the storage indexing keys
	indexing_prefix: &'static [u8],
}

impl<TStorage, TAccountId> BasicChannel<TStorage, TAccountId> {
	pub fn new(storage: TStorage, deny_unsafe: DenyUnsafe, indexing_prefix: &'static [u8]) -> Self {
		Self {
			_marker: Default::default(),
			deny_unsafe,
			storage: Arc::new(RwLock::new(storage)),
			indexing_prefix,
		}
	}
}

impl<TStorage, TAccountId> BasicChannelApi<TAccountId> for BasicChannel<TStorage, TAccountId>
where
	TAccountId: Codec + Send + Sync + 'static,
	TStorage: OffchainStorage + Send + Sync + 'static,
{
	fn get_merkle_proofs(&self, root: H256) -> Result<Proofs<TAccountId>> {
		self.deny_unsafe.check_if_safe()?;

		let key = offchain_key(self.indexing_prefix, root);
		if let Some(data) = self.storage.read().get(sp_offchain::STORAGE_PREFIX, &*key) {
			if let Ok(cdata) = <CommitmentData<TAccountId>>::decode(&mut data.as_slice()) {
				let num_coms = cdata.subcommitments.len();
				let mut accounts = Vec::with_capacity(num_coms);
				let mut commitments = Vec::with_capacity(num_coms);
				for (acc, com) in cdata.subcommitments {
					accounts.push(acc);
					commitments.push(com);
				};
				match generate_merkle_proofs(commitments.into_iter()) {
					Ok(proofs) => Ok(zip(accounts, proofs).collect::<Proofs<TAccountId>>()),
					Err(_) => Err(JsonError::invalid_request()),
				}
			} else {
				Err(JsonError::invalid_request())
			}
		} else {
			Err(JsonError::invalid_request())
		}
	}
}
