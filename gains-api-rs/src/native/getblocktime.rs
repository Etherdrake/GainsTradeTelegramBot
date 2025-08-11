use ethers::middleware::Middleware;
use ethers::prelude::{BlockId, Http, Provider};
use primitive_types::H160;
use crate::gains::utils::getchainparams::get_chain_params;

pub async fn get_block_time(block_number: u64, chain: u64) -> Result<u64, Box<dyn std::error::Error>> {
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

    // Connect to an Ethereum node
    let provider = Provider::<Http>::try_from(uri)?;

    // Get the block information
    let block = provider.get_block(BlockId::Number(block_number.into())).await?.unwrap();

    // Extract the timestamp from the block
    let timestamp = block.timestamp;

    let timestamp_conv: u64 = timestamp.as_u64();

    Ok(timestamp_conv)
}