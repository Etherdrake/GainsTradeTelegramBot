use super::super::super::nodes::constants::{ARBITRUM_URI, POLYGON_URI};

use ethers::core::k256::U256;
use ethers::prelude::gas_oracle::{GasOracle, ProviderOracle};
use ethers::prelude::{Http, Provider};
use crate::nodes::constants::ARBITRUM_SEPOLIA_URI;


//
pub async fn provider_oracle(chain: u64) -> primitive_types::U256 {
    let mut uri: &str;
    if chain == 137 {
        uri = POLYGON_URI;
    } else if chain == 42161 {
        uri = ARBITRUM_URI;
        } else if chain == 421614 {
        uri = ARBITRUM_SEPOLIA_URI;
        } else {
        panic!("Unsupported chain");
    }
    let provider = Provider::<Http>::try_from(uri).unwrap();
    let oracle = ProviderOracle::new(provider);
    match oracle.fetch().await {
        Ok(gas_price) => {
            // println!("[Provider oracle]: Gas price is {gas_price:?}");
            gas_price
        }
        Err(e) => panic!("[Provider oracle]: Cannot estimate gas: {e:?}"),
    }
}
