use std::str::FromStr;
use std::sync::Arc;
use ethers::middleware::Middleware;
use ethers::prelude::{abigen, Address, Http, Provider};
use ethers::utils::format_units;
use primitive_types::H160;
use crate::gains::utils::getchainparams::get_chain_params;
use crate::routes::native::payloads::GetGasBalancePayload;


pub async fn get_gas_balance(payload: GetGasBalancePayload) -> Result<String, Box<dyn std::error::Error>> {
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

    let pubkey: Address = match Address::from_str(&*payload.public_key) {
        Ok(address) => address,
        Err(_) => return Err(format!("Failed to parse public key: {}", &payload.public_key).into()),
    };

    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    abigen!(
        ERC20,
        "./abi/ERC20.abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let contract = ERC20::new(pubkey, Arc::new(client.clone()));

    // Call balanceOf
    let value = match client.get_balance(pubkey, None).await {
        Ok(val) => val,
        Err(e) => return Err(format!("Failed to get balance: {}", e).into()),
    };

    // We use format to get the right string
    let token_decimals = 18;

    let balance_ether_str = match format_units(value, token_decimals) {
        Ok(balance_ether_string) => { ;
            Ok(balance_ether_string)
        }
        Err(err) => {
            eprintln!("Error fetching balance: {}", err);
            Err(Box::new(err))
        }
    }?;

    Ok(balance_ether_str)
}