use ethcontract::common::abi::ethereum_types::Address;
use ethers::core::k256::elliptic_curve::consts::U24;
use primitive_types::U256;
use serde::{Deserialize, Serialize};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::TradeStruct;

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct TempOrder {
    user: Address,
    pair_index: U256,
    index: U256,
    position_size: U256,
    spread_reduction_p: U256,
    buy: bool,
    leverage: U256,
    tp: U256,
    sl: U256,
    min_price: U256,
    max_price: U256,
    block: U256,
    token_id: U256,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct OpenTradeBig {
    pub guid: i64,
    pub trade: TradeStruct,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct CancelOpenOrderBig {
    pub guid: i64,
    pub index: u32,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct CloseTradeMarketBig {
    pub guid: i64,
    pub index: u32,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateOpenLimitOrderBig {
    pub guid: i64,
    pub pair_index: U256,
    pub index: U256,
    pub price: U256,
    pub tp: U256,
    pub sl: U256,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateStopLossBig {
    pub guid: i64,
    pub pair_index: U256,
    pub index: U256,
    pub new_sl: U256,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateTakeProfitBig {
    pub guid: i64,
    pub pair_index: U256,
    pub index: U256,
    pub new_tp: U256,
}

#[derive(Debug, serde::Deserialize, serde::Serialize, Clone)]
pub struct OpenLimitOrderBigGNS {
    pub guid: i64,
    pub user: Address,
    pub pair_index: U256,
    pub index: U256,
    pub position_size: U256,
    pub spread_reduction_p: U256,
    pub buy: bool,
    pub leverage: U256,
    pub tp: U256,
    pub sl: U256,
    pub min_price: U256,
    pub max_price: U256,
    pub block: U256,
    pub token_id: U256,
    pub chain: String,
}
