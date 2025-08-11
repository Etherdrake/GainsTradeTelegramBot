use std::error::Error;
use std::sync::Arc;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use primitive_types::{H160, U256};
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{CollateralStruct, TradeStruct};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_collaterals(
    collateral_token: u8,
    chain: u64,
) -> Result<Vec<CollateralStruct>, Box<dyn Error>> {
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
        getCollaterals,
        "./gains_abi/facets/GNSTradingStorage/GNSTradingStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = getCollaterals::new(contract_address, Arc::new(client.clone()));

    let value = contract.get_collaterals().await?;

    let collaterals: Vec<CollateralStruct> = value
        .into_iter()
        .map(|data| CollateralStruct {
            collateral: data.collateral,
            is_active: data.is_active,
            __placeholder: data.placeholder,
            precision: data.precision,
            precision_delta: data.precision_delta
        })
        .collect();

    Ok(collaterals)
}