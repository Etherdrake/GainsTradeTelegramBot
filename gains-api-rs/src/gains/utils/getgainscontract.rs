use crate::gains::errors::ChainError;
use primitive_types::H160;
use crate::contracts::arbitrum::GNS_DIAMOND_CONTRACT_ARBITRUM;
use crate::contracts::polygon::GNS_DIAMOND_CONTRACT_POLYGON;
use crate::gains::contracts::GNS_DIAMOND_CONTRACT_SEPOLIA;

// TODO: Add the different collaterals that exist for every single diamond contract
pub fn get_diamond_contract_address(chain: u64) -> Result<H160, Box<ChainError>> {
    let diamond_contract: H160;
    if chain == 137 {
        diamond_contract = *GNS_DIAMOND_CONTRACT_POLYGON;
    } else if chain == 42161 {
        diamond_contract = *GNS_DIAMOND_CONTRACT_ARBITRUM;
    } else if chain == 421614 {
        diamond_contract = *GNS_DIAMOND_CONTRACT_SEPOLIA
    } else {
        return Err(Box::new(ChainError {
            message: "Wrong chain!".parse().unwrap(),
        }));
    }
    Ok(diamond_contract)
}
