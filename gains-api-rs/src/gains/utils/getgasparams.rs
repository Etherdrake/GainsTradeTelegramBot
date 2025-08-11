use std::error::Error;
use primitive_types::{H160, U256};
use ethers::utils::parse_units;
use crate::gains::errors::ChainError;
use crate::gains::helpers::oracle::provider_oracle;

// Modified function to get gas params
pub async fn get_gas_params(chain: u64) -> Result<(U256, U256, U256), Box<dyn Error>> {
    let gas_fee = provider_oracle(chain).await;
    let max_prio_fee;
    let max_fee_per_gas;

    if chain == 137 {
        max_prio_fee = U256::from(parse_units("80", "gwei")?);
        max_fee_per_gas = U256::from(parse_units("100", "gwei")?);
    } else if chain == 42161 {
        max_prio_fee = U256::from(parse_units("0.1", "gwei")?);
        max_fee_per_gas = gas_fee + max_prio_fee;
    } else if chain == 421614 {
        max_prio_fee = U256::from(parse_units("0.2", "gwei")?);
        max_fee_per_gas = gas_fee + max_prio_fee;
    } else {
        return Err(Box::new(ChainError {
            message: "Wrong chain!".to_string(),
        }));
    }
    Ok((gas_fee, max_prio_fee, max_fee_per_gas))
}