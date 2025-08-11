use serde::Deserialize;

#[derive(Deserialize)]
pub struct GetAllPendingOrdersPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetAllTradeInfosPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetAllTradesPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetCollateralsPayload {
    pub(crate) offset: u64,
    pub(crate) limit: u64,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetPendingOrderPayload {
    pub(crate) trader: String,
    pub(crate) order_id: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetPendingOrdersPayload {
    pub(crate) trader: String,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetPnlPercentPayload {
    pub(crate) open_price: u64,
    pub(crate) current_price: u64,
    pub(crate) long: bool,
    pub(crate) leverage: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetTradeInfoPayload {
    pub(crate) trader: String,
    pub(crate) index: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetTradeInfosPayload {
    pub(crate) trader: String,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetTradersPayload {
    pub(crate) offset: u32,
    pub(crate) limit: u32,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}

#[derive(Deserialize)]
pub struct GetTradesPayload {
    pub(crate) trader: String,
    pub(crate) collateral_token: u8,
    pub(crate) chain: u64,
}