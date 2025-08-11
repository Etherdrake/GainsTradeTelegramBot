use std::convert::Infallible;
use std::error::Error;
use std::sync::Arc;
use ethers::abi::Address;
use ethers::contract::abigen;
use ethers::prelude::{Http, I256, Provider, U64};
use primitive_types::{H160, U256};
use crate::blockchain_constants::chainids::{ARBITRUM_ID, ARBITRUM_SEPOLIA_ID, POLYGON_ID};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};
use crate::contracts::sepolia::GNS_DIAMOND_CONTRACT_SEPOLIA;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_pnl_percent(
    open_price: u64,
    current_price: u64,
    long: bool,
    leverage: u32,
    collateral_token: u8,
    chain: u64,
) -> Result<I256, Box<dyn Error>> { // TODO: Check all the functions and especially how this is calculated
    let mut trading_contract: H160 = Default::default();
    let mut chain_id: u64 = 0;
    let mut uri: &str = "";

    match get_chain_params(chain, collateral_token) {
        Ok((contract, id, uri_str)) => {
            trading_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e),
    }

    abigen!(
        getPnlPercent,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getPnlPercent::new(contract_address, Arc::new(client.clone()));

    let value = contract.get_pnl_percent(open_price, current_price, long, leverage).await?;

    Ok(value)
}
