use std::error::Error;
use std::str::FromStr;
use std::sync::Arc;
use ethers::contract::abigen;
use ethers::core::k256::elliptic_curve::consts::U2;
use ethers::prelude::{Http, Provider};
use ethers::utils::{format_units, parse_ether, parse_units, ParseUnits};
use num_format::Locale::qu;
use primitive_types::{H160, U256};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgainscontract::get_diamond_contract_address;

pub async fn get_pair_min_pos_size_usd(
    index: u64,
    chain: u64
) -> Result<f64, Box<dyn Error>> {
    let mut trading_contract: H160 = Default::default();
    let mut chain_id: u64 = 0;
    let mut uri: &str = "";

    match get_chain_params(chain, 1) {
        Ok((contract, id, uri_str)) => {
            trading_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e),
    }

    abigen!(
        pairMinPositionSizeUsd,
        "./gains_abi/facets/GNSPairsStorage/GNSPairsStorage_abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_address = get_diamond_contract_address(chain).unwrap();
    let contract = pairMinPositionSizeUsd::new(contract_address, Arc::new(client.clone()));

    let index_conv: U256 = U256::from(index);
    let query: U256 = contract.pair_min_position_size_usd(index_conv).await?;

    let query_parsed = format_units(query, 18).unwrap();
    let query_float: f64 = f64::from_str(&*query_parsed).unwrap();

    Ok(query_float)
}