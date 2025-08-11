use std::str::FromStr;
use ethers::prelude::Address;
use ethers::utils::{parse_units, ParseUnits};
use primitive_types::{U128, U256};
use serde::{Deserialize, Serialize};
use crate::gains::tradinginteractions::write::bigtypes::{TradePayloadBig, UpdateOpenOrderPayloadBig, UpdateSlPayloadBig, UpdateTpPayloadBig};

// This is what we receive from HootCore
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TradePayload {
    pub(crate) user: String,
    pub(crate) index: u32,
    pub(crate) pair_index: u16,
    pub(crate) leverage: f32,
    pub(crate) long: bool,
    pub(crate) is_open: bool,
    pub(crate) collateral_index: u8,
    pub(crate) trade_type: u8,
    pub(crate) collateral_amount: f64,
    pub(crate) open_price: f64,
    pub(crate) tp: f64,
    pub(crate) sl: f64,
    pub(crate) __placeholder: u128,
}

impl TradePayload {
    pub fn format(&self) -> String {
        format!(
            "TradePayload {{\n    user: {},\n    index: {},\n    pair_index: {},\n    leverage: {},\n    long: {},\n    is_open: {},\n    collateral_index: {},\n    trade_type: {},\n    collateral_amount: {},\n    open_price: {},\n    tp: {},\n    sl: {},\n    __placeholder: {}\n}}",
            self.user,
            self.index,
            self.pair_index,
            self.leverage,
            self.long,
            self.is_open,
            self.collateral_index,
            self.trade_type,
            self.collateral_amount,
            self.open_price,
            self.tp,
            self.sl,
            self.__placeholder
        )
    }
}

impl From<TradePayload> for TradePayloadBig {
    fn from(payload: TradePayload) -> Self {
        let conv_leverage: u32 = (payload.leverage * 1000.0) as u32;

        // TODO: This is dependent on token, add logic once
        // TODO: UserCache CollateralIndex is equalized vis-a-vis Gains
        // TODO: USDC is 6
        let pu_collateral_amount: ParseUnits =
        parse_units(payload.collateral_amount.clone(), 18).unwrap();
        let collateral_amount_parsed = U256::from(pu_collateral_amount).as_u128();

        let pu_open: ParseUnits = parse_units(payload.open_price.clone(), 10).unwrap();
        let pu_tp: ParseUnits = parse_units(payload.tp.clone(), 10).unwrap();
        let pu_sl: ParseUnits = parse_units(payload.sl.clone(), 10).unwrap();

        // Panics if the number is larger than u64::max_value() -> kills thread
        // TODO: Make sure this times-out correctly if it happens, maybe check for type.max_value on HootCore already
        let open_parsed = U256::from(pu_open).as_u64();
        let tp_parsed = U256::from(pu_tp).as_u64();
        let sl_parsed = U256::from(pu_sl).as_u64();

        let user: Address = match Address::from_str(&*payload.user) {
            Ok(address) => address,
            Err(err) => {
                eprintln!("Error parsing address: {}", err);
                Address::default()
            }
        };

        TradePayloadBig {
            user: user,
            index: payload.index,
            pair_index: payload.pair_index,
            leverage: conv_leverage,
            long: payload.long,
            is_open: payload.is_open,
            collateral_index: payload.collateral_index,
            trade_type: payload.trade_type,
            collateral_amount: collateral_amount_parsed,
            open_price: open_parsed,
            tp: tp_parsed,
            sl: sl_parsed,
            __placeholder: payload.__placeholder,
        }
    }
}


#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct OpenTradePayload {
    pub(crate) guid: i64,
    pub(crate) trade: TradePayload,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl OpenTradePayload {
    pub fn print(&self) {
        println!("OpenTradeSuccess {{");
        println!("    guid: {},", self.guid);
        println!("    trade: {:?}", self.trade); // Assuming TradePayload implements Debug
        println!("    chain: {},", self.chain);
        println!("    collateral_token: {},", self.collateral_token);
        println!("}}");
    }
    pub fn format_success(&self) -> String {
        format!(
            "OpenTradePayload {{\n    guid: {},\n    trader: {:?}\n    trade: TradePayload {{\n        user: {},\n        index: {},\n        pair_index: {},\n        leverage: {},\n        long: {},\n        is_open: {},\n        collateral_index: {},\n        trade_type: {},\n        collateral_amount: {},\n        open_price: {},\n        tp: {},\n        sl: {},\n        __placeholder: {}\n    }}\n    chain: {},\n    collateral_token: {}\n}}",
            self.guid,
            self.trade.user,
            self.trade.user,
            self.trade.index,
            self.trade.pair_index,
            self.trade.leverage,
            self.trade.long,
            self.trade.is_open,
            self.trade.collateral_index,
            self.trade.trade_type,
            self.trade.collateral_amount,
            self.trade.open_price,
            self.trade.tp,
            self.trade.sl,
            self.trade.__placeholder,
            self.chain,
            self.collateral_token
        )
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct CancelOpenOrderPayload {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl CancelOpenOrderPayload {
    pub fn print(&self) {
        println!("CancelOpenOrderPayload {{");
        println!("    guid: {},", self.guid);
        println!("    index: {},", self.index);
        println!("    chain: {},", self.chain);
        println!("    collateral_token: {},", self.collateral_token);
        println!("}}");
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct CloseTradeMarketPayload {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl CloseTradeMarketPayload {
    pub fn print(&self) {
        println!("CloseTradeMarketPayload {{");
        println!("    guid: {},", self.guid);
        println!("    index: {},", self.index);
        println!("    chain: {},", self.chain);
        println!("    collateral_token: {},", self.collateral_token);
        println!("}}");
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateOpenOrderPayload {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) trigger_price: u64,
    pub(crate) tp: u64,
    pub(crate) sl: u64,
    pub(crate) max_slippage_p: u16,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl UpdateOpenOrderPayload {
    pub fn print(&self) {
        println!("UpdateOpenOrderPayload {{");
        println!("    guid: {},", self.guid);
        println!("    index: {},", self.index);
        println!("    trigger_price: {},", self.trigger_price);
        println!("    tp: {},", self.tp);
        println!("    sl: {},", self.sl);
        println!("    max_slippage_p: {},", self.max_slippage_p);
        println!("    chain: {},", self.chain);
        println!("    collateral_token: {},", self.collateral_token);
        println!("}}");
    }
}

impl From<UpdateOpenOrderPayload> for UpdateOpenOrderPayloadBig {
    fn from(payload: UpdateOpenOrderPayload) -> Self {
        let pu_trigger: ParseUnits = parse_units(payload.trigger_price.clone(), 10).unwrap();
        let pu_tp: ParseUnits = parse_units(payload.tp.clone(), 10).unwrap();
        let pu_sl: ParseUnits = parse_units(payload.sl.clone(), 10).unwrap();

        // Panics if the number is larger than u64::max_value() -> kills thread
        // TODO: Make sure this times-out correctly if it happens, maybe check for type.max_value on HootCore already
        let trigger_parsed = U256::from(pu_trigger).as_u64();
        let tp_parsed = U256::from(pu_tp).as_u64();
        let sl_parsed = U256::from(pu_sl).as_u64();

        UpdateOpenOrderPayloadBig {
            guid: payload.guid,
            index: payload.index,
            trigger_price: trigger_parsed,
            tp: tp_parsed,
            sl: sl_parsed,
            max_slippage_p: payload.max_slippage_p,
            chain: payload.chain,
            collateral_token: payload.collateral_token,
        }
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateSlPayload {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) new_sl: u64,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl From<UpdateSlPayload> for UpdateSlPayloadBig {
    fn from(payload: UpdateSlPayload) -> Self {
        let pu_sl: ParseUnits = parse_units(payload.new_sl.clone(), 10).unwrap();

        // Panics if the number is larger than u64::max_value() -> kills thread
        // TODO: Make sure this times-out correctly if it happens, maybe check for type.max_value on HootCore already
        let sl_parsed = U256::from(pu_sl).as_u64();

        UpdateSlPayloadBig {
            guid: payload.guid,
            index: payload.index,
            new_sl: sl_parsed,
            chain: payload.chain,
            collateral_token: payload.collateral_token,
        }
    }
}

#[derive(Clone, Debug, Deserialize, Serialize)]
pub struct UpdateTpPayload {
    pub(crate) guid: i64,
    pub(crate) index: u32,
    pub(crate) new_tp: u64,
    pub(crate) chain: u64,
    pub(crate) collateral_token: u8,
}

impl From<UpdateTpPayload> for UpdateTpPayloadBig {
    fn from(payload: UpdateTpPayload) -> Self {
        let pu_tp: ParseUnits = parse_units(payload.new_tp.clone(), 10).unwrap();

        // Panics if the number is larger than u64::max_value() -> kills thread
        // TODO: Make sure this times-out correctly if it happens, maybe check for type.max_value on HootCore already
        let tp_parsed = U256::from(pu_tp).as_u64();

        UpdateTpPayloadBig {
            guid: payload.guid,
            index: payload.index,
            new_tp: tp_parsed,
            chain: payload.chain,
            collateral_token: payload.collateral_token,

        }
    }
}

#[cfg(test)]
mod tests {
    use std::str::FromStr;
    use ethers::types::Address;
    use super::*;
    #[test]
    fn test_trade_payload_conversion() {
        // Given values
        let payload = TradePayload {
            user: String::from("0xbc9b019aa5885ba50f3770b6fd7e5da981007068"),
            index: 2,
            pair_index: 0,
            leverage: 38.5,
            long: true,
            is_open: true,
            collateral_index: 1,
            trade_type: 1,
            collateral_amount: 500.0,
            open_price: 50960.59,
            tp: 50828.2,
            sl: 62873.5,
            __placeholder: 0,
        };

        // Convert TradePayload to TradePayloadBig
        let payload_big: TradePayloadBig = payload.into();

        // Expected values
        let expected_user = Address::from_str("0xbc9b019aa5885ba50f3770b6fd7e5da981007068").unwrap(); // TODO: Add expect for proper error handling
        let expected_index = 2;
        let expected_pair_index = 0;
        let expected_leverage = 38500;
        let expected_long = true;
        let expected_is_open = true;
        let expected_collateral_index = 1;
        let expected_trade_type = 1;
        let expected_collateral_amount = 500000000000000000000;
        let expected_open_price = 509605900000000;
        let expected_tp = 508282248311688;
        let expected_sl = 628734551948051;
        let expected_placeholder = 0;

        // Assert values
        assert_eq!(payload_big.user, expected_user);
        assert_eq!(payload_big.index, expected_index);
        assert_eq!(payload_big.pair_index, expected_pair_index);
        assert_eq!(payload_big.leverage, expected_leverage);
        assert_eq!(payload_big.long, expected_long);
        assert_eq!(payload_big.is_open, expected_is_open);
        assert_eq!(payload_big.collateral_index, expected_collateral_index);
        assert_eq!(payload_big.trade_type, expected_trade_type);
        assert_eq!(payload_big.collateral_amount, expected_collateral_amount);
        assert_eq!(payload_big.open_price, expected_open_price);
        assert_eq!(payload_big.tp, expected_tp);
        assert_eq!(payload_big.sl, expected_sl);
        assert_eq!(payload_big.__placeholder, expected_placeholder);
    }
}