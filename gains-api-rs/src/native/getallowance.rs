use std::str::FromStr;
use std::sync::Arc;
use ethcontract::common::abi::ethereum_types::Address;
use ethers::contract::abigen;
use ethers::prelude::{Http, Provider};
use ethers::utils::format_units;
use primitive_types::H160;
use crate::gains::utils::getchainparams::get_chain_params;
use crate::native::getdecimals::{ERC20, get_decimals};
use crate::routes::native::payloads::{GetAllowancePayload};

pub async fn get_allowance(payload: GetAllowancePayload) -> Result<String, Box<dyn std::error::Error>> {
    let mut diamond_contract: H160 = Default::default();
    let mut chain_id: u64 = 0;
    let mut uri: &str = "";

    match get_chain_params(payload.chain, payload.collateral) {
        Ok((contract, id, uri_str)) => {
            diamond_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e.into()),
    }

    let owner: Address = Address::from_str(&payload.owner).map_err(|e| format!("fn get_allowance: Failed to parse owner: {}", e))?;
    let spender: Address = Address::from_str(&payload.spender).map_err(|e| format!("fn get_allowance: Failed to parse spender: {}", e))?;
    let src_token: Address = Address::from_str(&payload.token).map_err(|e| format!("fn get_allowance: Failed to parse token: {}", e))?;

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    let contract_src_token = ERC20::new(src_token, Arc::clone(&client));

    // Retrieve decimals
    let token_decimals = get_decimals(&contract_src_token).await?;

    // Call allowance
    let value = contract_src_token.allowance(owner, spender).call().await?;

    // Format the allowance value
    let allowance_str = match format_units(value, token_decimals) {
        Ok(allowance_string) => {
            Ok(allowance_string)
        }
        Err(err) => {
            eprintln!("Error fetching allowance: {}", err);
            Err(Box::new(err))
        }
    }?;

    Ok(allowance_str)
}