use ethers::types::Address;
use serde::{Deserialize, Serialize};

// BigTypes are what we encode and send of to the blockchain
// all BigTypes have a corresponding type. A payload received
// from HootCore always is a normal type, and then converted into
// a BigType that is then passed as param from the handler function
// to the blockchain interaction function

// This is what we encode and send to the blockchain
// and matches the ITradingStorage defined types.
// https://gains-network.gitbook.io/docs-home/developer/technical-reference/facets/gnstradingstorage
#[derive(Debug, Serialize, Deserialize)]
pub struct TradePayloadBig {
    pub(crate) user: Address,
    pub(crate) index: u32,
    pub(crate) pair_index: u16,
    pub(crate) leverage: u32,
    pub(crate) long: bool,
    pub(crate) is_open: bool,
    pub(crate) collateral_index: u8,
    pub(crate) trade_type: u8,
    pub(crate) collateral_amount: u128,
    pub(crate) open_price: u64,
    pub(crate) tp: u64,
    pub(crate) sl: u64,
    pub(crate) __placeholder: u128,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateOpenOrderPayloadBig {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) trigger_price: u64,
    pub(crate) tp: u64,
    pub(crate) sl: u64,
    pub(crate) max_slippage_p: u16,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateSlPayloadBig {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) new_sl: u64,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateTpPayloadBig {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) new_tp: u64,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}
