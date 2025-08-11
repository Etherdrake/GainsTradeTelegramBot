use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethers::abi::Address;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use primitive_types::{H160, U256};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{OrderId, TradeStruct, PendingOrderStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_pending_order(
    trader: String,
    order_id: u32,
    collateral_token: u8,
    chain: u64,
) -> Result<PendingOrderStruct, Box<dyn Error>> {
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
        getPendingOrder,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getPendingOrder::new(contract_address, Arc::new(client.clone()));

    let id_instance = Id {
        user: Address::from_str(&trader).unwrap(),
        index: order_id,
    };

    let query = contract.get_pending_order(id_instance).await?;

    let pending_order = PendingOrderStruct {
        trade: TradeStruct {
            user: query.trade.user,
            index: query.trade.index,
            pair_index: query.trade.pair_index,
            leverage: query.trade.leverage,
            long: query.trade.long,
            is_open: query.trade.is_open,
            collateral_index: query.trade.collateral_index,
            trade_type: query.trade.trade_type,
            collateral_amount: query.trade.collateral_amount,
            open_price: query.trade.open_price,
            tp: query.trade.tp,
            sl: query.trade.sl,
            __placeholder: Default::default(),
        },
        user: query.user,
        index: query.index,
        is_open: query.is_open,
        order_type: query.order_type,
        created_block: query.created_block,
        max_slippage_p: query.max_slippage_p,
    };

    Ok(pending_order)
}