
//! Autogenerated weights for pallet_collective
//!
//! THIS FILE WAS AUTO-GENERATED USING THE SUBSTRATE BENCHMARK CLI VERSION 3.0.0
//! DATE: 2021-05-08, STEPS: `[50, ]`, REPEAT: 20, LOW RANGE: `[]`, HIGH RANGE: `[]`
//! EXECUTION: Some(Wasm), WASM-EXECUTION: Compiled, CHAIN: Some("/tmp/snowbridge-benchmark-tce/spec.json"), DB CACHE: 128

// Executed Command:
// target/release/snowbridge
// benchmark
// --chain
// /tmp/snowbridge-benchmark-tce/spec.json
// --execution
// wasm
// --wasm-execution
// compiled
// --pallet
// pallet_collective
// --extrinsic
// *
// --repeat
// 20
// --steps
// 50
// --output
// runtime/snowbridge/src/weights/pallet_collective_weights.rs


#![allow(unused_parens)]
#![allow(unused_imports)]

use frame_support::{traits::Get, weights::Weight};
use sp_std::marker::PhantomData;

/// Weight functions for pallet_collective.
pub struct WeightInfo<T>(PhantomData<T>);
impl<T: frame_system::Config> pallet_collective::WeightInfo for WeightInfo<T> {
	fn set_members(m: u32, _n: u32, p: u32, ) -> Weight {
		(0 as Weight)
			// Standard Error: 237_000
			.saturating_add((19_985_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 8_000
			.saturating_add((10_603_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(2 as Weight))
			.saturating_add(T::DbWeight::get().reads((1 as Weight).saturating_mul(p as Weight)))
			.saturating_add(T::DbWeight::get().writes(2 as Weight))
			.saturating_add(T::DbWeight::get().writes((1 as Weight).saturating_mul(p as Weight)))
	}
	fn execute(b: u32, m: u32, ) -> Weight {
		(32_188_000 as Weight)
			// Standard Error: 0
			.saturating_add((4_000 as Weight).saturating_mul(b as Weight))
			// Standard Error: 6_000
			.saturating_add((118_000 as Weight).saturating_mul(m as Weight))
			.saturating_add(T::DbWeight::get().reads(1 as Weight))
	}
	fn propose_execute(b: u32, m: u32, ) -> Weight {
		(39_471_000 as Weight)
			// Standard Error: 0
			.saturating_add((4_000 as Weight).saturating_mul(b as Weight))
			// Standard Error: 8_000
			.saturating_add((302_000 as Weight).saturating_mul(m as Weight))
			.saturating_add(T::DbWeight::get().reads(2 as Weight))
	}
	fn propose_proposed(b: u32, m: u32, p: u32, ) -> Weight {
		(61_614_000 as Weight)
			// Standard Error: 0
			.saturating_add((10_000 as Weight).saturating_mul(b as Weight))
			// Standard Error: 101_000
			.saturating_add((139_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 2_000
			.saturating_add((709_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(4 as Weight))
			.saturating_add(T::DbWeight::get().writes(4 as Weight))
	}
	fn vote(m: u32, ) -> Weight {
		(55_811_000 as Weight)
			// Standard Error: 218_000
			.saturating_add((699_000 as Weight).saturating_mul(m as Weight))
			.saturating_add(T::DbWeight::get().reads(2 as Weight))
			.saturating_add(T::DbWeight::get().writes(1 as Weight))
	}
	fn close_early_disapproved(m: u32, p: u32, ) -> Weight {
		(56_088_000 as Weight)
			// Standard Error: 133_000
			.saturating_add((873_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 2_000
			.saturating_add((692_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(3 as Weight))
			.saturating_add(T::DbWeight::get().writes(3 as Weight))
	}
	fn close_early_approved(b: u32, m: u32, p: u32, ) -> Weight {
		(85_526_000 as Weight)
			// Standard Error: 0
			.saturating_add((6_000 as Weight).saturating_mul(b as Weight))
			// Standard Error: 188_000
			.saturating_add((617_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 2_000
			.saturating_add((704_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(4 as Weight))
			.saturating_add(T::DbWeight::get().writes(3 as Weight))
	}
	fn close_disapproved(m: u32, p: u32, ) -> Weight {
		(58_989_000 as Weight)
			// Standard Error: 110_000
			.saturating_add((1_273_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 1_000
			.saturating_add((716_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(4 as Weight))
			.saturating_add(T::DbWeight::get().writes(3 as Weight))
	}
	fn close_approved(b: u32, m: u32, p: u32, ) -> Weight {
		(91_350_000 as Weight)
			// Standard Error: 0
			.saturating_add((6_000 as Weight).saturating_mul(b as Weight))
			// Standard Error: 181_000
			.saturating_add((686_000 as Weight).saturating_mul(m as Weight))
			// Standard Error: 2_000
			.saturating_add((709_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(5 as Weight))
			.saturating_add(T::DbWeight::get().writes(3 as Weight))
	}
	fn disapprove_proposal(p: u32, ) -> Weight {
		(36_629_000 as Weight)
			// Standard Error: 2_000
			.saturating_add((666_000 as Weight).saturating_mul(p as Weight))
			.saturating_add(T::DbWeight::get().reads(1 as Weight))
			.saturating_add(T::DbWeight::get().writes(3 as Weight))
	}
}
