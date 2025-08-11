use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct TradePayload {
    user: String,
    index: u32,
    pair_index: u16,
    leverage: u32,
    long: bool,
    is_open: bool,
    collateral_index: u8,
    trade_type: u8,
    collateral_amount: u128,
    open_price: u64,
    tp: u64,
    sl: u64,
    __placeholder: u64,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CancelOpenOrderPayload {
    pub(crate) index: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct CloseTradeMarketPayload {
    pub(crate) index: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct OpenTradePayload {
    pub(crate) trade: TradePayload,
    pub(crate) max_slippage_p: u16,
    pub(crate) referrer: String,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UpdateOpenOrderPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UpdateSlPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UpdateTpPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: String,
}
