use std::convert::Infallible;
use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethcontract::common::abi::ethereum_types::Address;
use ethers::contract::abigen;
use ethers::core::k256::elliptic_curve::consts::U24;
use ethers::prelude::{Http, Provider};
use primitive_types::H160;
use crate::blockchain_constants::chainids::{ARBITRUM_ID, ARBITRUM_SEPOLIA_ID, POLYGON_ID};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};use crate::contracts::sepolia::GNS_DIAMOND_CONTRACT_SEPOLIA;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{TradeStruct, PendingOrderStruct, TradeInfoStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_trades_infos(
    trader: String,
    collateral_token: u8,
    chain: u64,
) -> Result<Vec<TradeInfoStruct>, Box<dyn Error>> {
    let mut trading_contract: H160 = Default::default();
    let mut chain_id: u64 = 0;
    let mut uri: &str = "";

    match get_chain_params(chain, collateral_token){
        Ok((contract, id, uri_str)) => {
            trading_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e),
    }

    abigen!(
        getTradesInfos,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getTradesInfos::new(contract_address, Arc::new(client.clone()));

    let trader_address: Address = Address::from_str(&*trader)
        .map_err(|e| Box::new(e) as Box<dyn Error>)?;

    let value = contract.get_trade_infos(trader_address).await?;

    let trade_infos: Vec<TradeInfoStruct> = value
        .into_iter()
        .map(|trade_info| TradeInfoStruct {
            created_block: trade_info.created_block,
            tp_last_updated_block: trade_info.tp_last_updated_block,
            sl_last_updated_block: trade_info.sl_last_updated_block,
            max_slippage_p:trade_info.max_slippage_p,
            last_oi_update_ts: trade_info.last_oi_update_ts,
            collateral_price_usd: trade_info.collateral_price_usd,
            __placeholder: trade_info.placeholder,
        })
        .collect();

    Ok(trade_infos)
}




















