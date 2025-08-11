use std::convert::Infallible;
use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethers::abi::Address;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use primitive_types::{H160, U256};
use crate::blockchain_constants::chainids::{ARBITRUM_ID, ARBITRUM_SEPOLIA_ID, POLYGON_ID};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};
use crate::contracts::sepolia::GNS_DIAMOND_CONTRACT_SEPOLIA;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{PendingOrderStruct, TradeStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_pending_orders(
    trader: String,
    collateral_token: u8,
    chain: u64,
) -> Result<Vec<PendingOrderStruct>, Box<dyn Error>> {
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
        getPendingOrders,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getPendingOrders::new(contract_address, Arc::new(client.clone()));

    let value = contract.get_pending_orders(Address::from_str(&trader).unwrap()).await?;

    let all_trades: Vec<PendingOrderStruct> = value
        .into_iter()
        .map(|order| PendingOrderStruct {
            trade: TradeStruct {
                user: Default::default(),
                index: 0,
                pair_index: 0,
                leverage: 0,
                long: false,
                is_open: false,
                collateral_index: 0,
                trade_type: 0,
                collateral_amount: 0,
                open_price: 0,
                tp: 0,
                sl: 0,
                __placeholder: Default::default(),
            },
            user: order.user,
            index: order.index,
            is_open: order.is_open,
            order_type: 0,
            created_block: 0,
            max_slippage_p: 0,
        })
        .collect();

    Ok(all_trades)
}