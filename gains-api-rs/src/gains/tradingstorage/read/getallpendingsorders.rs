use std::error::Error;
use std::sync::Arc;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use ethers::prelude::types::{U256, H160};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{TradeStruct, PendingOrderStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;


pub async fn get_all_pending_orders(
    offset: u64,
    limit: u64,
    collateral_token: u8,
    chain: u64,
) ->  Result<Vec<PendingOrderStruct>, Box<dyn Error>> {
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
        getAllPendingOrders,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );
    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getAllPendingOrders::new(contract_address, Arc::new(client.clone()));


    let value = contract.get_all_pending_orders(U256::from(offset), U256::from(limit)).await?;

    println!("Pending Orders: {:?}", value);

    let pending_orders: Vec<PendingOrderStruct> = value
        .into_iter()
        .map(|order| PendingOrderStruct {
            trade: TradeStruct {
                user: order.trade.user,
                index: order.trade.index,
                pair_index: order.trade.pair_index,
                leverage: order.trade.leverage,
                long: order.trade.long,
                is_open: order.trade.is_open,
                collateral_index: order.trade.collateral_index,
                trade_type: order.trade.trade_type,
                collateral_amount: order.trade.collateral_amount,
                open_price: order.trade.open_price,
                tp: order.trade.tp,
                sl: order.trade.sl,
                __placeholder: order.trade.placeholder,
            },
            user: order.user,
            index: order.index,
            is_open: order.trade.is_open,
            order_type: order.order_type,
            created_block: order.created_block,
            max_slippage_p: order.max_slippage_p,
        })
        .collect();

    Ok(pending_orders)
}