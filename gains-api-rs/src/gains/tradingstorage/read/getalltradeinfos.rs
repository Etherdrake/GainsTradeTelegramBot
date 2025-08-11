use std::error::Error;
use std::sync::Arc;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use primitive_types::{H160, U256};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{TradeStruct, PendingOrderStruct, TradeInfoStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_all_trade_infos(
    offset: u64,
    limit: u64,
    collateral_token: u8,
    chain: u64,
) -> Result<Vec<TradeInfoStruct>,  Box<dyn Error>>  {
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
        getAllTradeInfos,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getAllTradeInfos::new(contract_address, Arc::new(client.clone()));

    let value = contract.get_all_trade_infos(U256::from(offset), U256::from(limit)).call().await?;

    let all_trades: Vec<TradeInfoStruct> = value.into_iter()
        .map(|trade_info| TradeInfoStruct {
            created_block: trade_info.created_block,
            tp_last_updated_block: trade_info.tp_last_updated_block,
            sl_last_updated_block: trade_info.sl_last_updated_block,
            max_slippage_p: trade_info.max_slippage_p,
            last_oi_update_ts: trade_info.last_oi_update_ts,
            collateral_price_usd: trade_info.collateral_price_usd,
            __placeholder: trade_info.placeholder,
        })
        .collect();

    Ok(all_trades)
}