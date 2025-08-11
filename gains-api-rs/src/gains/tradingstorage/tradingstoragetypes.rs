pub mod trading_storage {
    use ethcontract::U256;
    use ethers::types::Address;
    use serde::{Deserialize, Serialize};

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct TradeStruct {
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
        pub(crate) __placeholder: U256,
    }

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct TradeInfoStruct {
        pub(crate) created_block: u32,
        pub(crate) tp_last_updated_block: u32,
        pub(crate) sl_last_updated_block: u32,
        pub(crate) max_slippage_p: u16,
        pub(crate) last_oi_update_ts: u64,
        pub(crate) collateral_price_usd: u64,
        pub(crate) __placeholder: u64,
    }

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct PendingOrderStruct {
        pub(crate) trade: TradeStruct,
        pub(crate) user: Address,
        pub(crate) index: u32,
        pub(crate) is_open: bool,
        pub(crate) order_type: u8,
        pub(crate) created_block: u32,
        pub(crate) max_slippage_p: u16,
    }

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct CollateralStruct {
        pub(crate) collateral: Address,
        pub(crate) is_active: bool,
        pub(crate) __placeholder: u128,
        pub(crate) precision: u128,
        pub(crate) precision_delta: u128,
    }

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct OrderId {
        pub(crate) trader: Address,
        pub(crate) order_id: u32,
    }

    #[derive(Clone, Debug, Deserialize, Serialize)]
    pub struct Counter {
        pub(crate) current_index: u32,
        pub (crate) open_count: u32,
        pub (crate) __placeholder: U256
    }
}
