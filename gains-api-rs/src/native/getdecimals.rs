
use ethers::prelude::*;

// We use this novel method to pass the instance itself.
abigen!(
    ERC20,
    "./abi/ERC20.abi.json",
    event_derives(serde::Deserialize, serde::Serialize)
);

pub async fn get_decimals(contract: &ERC20<Provider<Http>>) -> Result<u32, Box<dyn std::error::Error>> {
    let decimals = contract.decimals().call().await?;
    Ok(decimals.into())
}