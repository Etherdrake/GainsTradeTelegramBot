use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethers::abi::Address;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use primitive_types::H160;
use crate::blockchain_constants::chainids::{ARBITRUM_ID, ARBITRUM_SEPOLIA_ID, POLYGON_ID};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};
use crate::contracts::sepolia::GNS_DIAMOND_CONTRACT_SEPOLIA;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{TradeStruct, PendingOrderStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;


pub async fn get_trades(
    trader: String,
    collateral_token: u8,
    chain: u64,
) -> Result<Vec<TradeStruct>, Box<dyn Error>> {
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
        getTrades,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getTrades::new(contract_address, Arc::new(client.clone()));

    let value = contract.get_trades(Address::from_str(&trader).unwrap()).await?;

    let all_trades: Vec<TradeStruct> = value
        .into_iter()
        .map(|trade| TradeStruct {
            user: trade.user,
            index: trade.index,
            pair_index: Default::default(),
            leverage: Default::default(),
            long: Default::default(),
            is_open: trade.is_open,
            collateral_index: Default::default(),
            trade_type: Default::default(),
            collateral_amount: Default::default(),
            open_price: Default::default(),
            tp: Default::default(),
            sl: Default::default(),
            __placeholder: Default::default(),
        })
        .collect();
    Ok(all_trades)
}











































