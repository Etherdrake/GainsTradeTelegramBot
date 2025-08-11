use std::error::Error;
use primitive_types::H160;
use crate::blockchain_constants::chainids::{ARBITRUM_ID, ARBITRUM_SEPOLIA_ID, POLYGON_ID};
use crate::contracts::arbitrum::GNS_DIAMOND_CONTRACT_ARBITRUM;
use crate::contracts::polygon::GNS_DIAMOND_CONTRACT_POLYGON;
use crate::gains::contracts::GNS_DIAMOND_CONTRACT_SEPOLIA;
use crate::gains::errors::{ChainError, CollateralError};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};

pub fn get_chain_params(
    chain: u64,
    collateral_token: u8,
) -> Result<(H160, u64, &'static str), Box<dyn Error>> {
    let diamond_contract: H160;
    let chain_id: u64;
    let uri: &str;

    // TODO: The native token needs to be handled correctly here as well
    match chain {
        137 => {
            diamond_contract = match collateral_token {
                1 => *GNS_DIAMOND_CONTRACT_POLYGON,
                2 => *GNS_DIAMOND_CONTRACT_POLYGON,
                3 => *GNS_DIAMOND_CONTRACT_POLYGON,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Incorrect collateral value!".to_string(),
                    }));
                }
            };
            chain_id = POLYGON_ID;
            uri = POLYGON_URI;
        }
        42161 => {
            diamond_contract = match collateral_token {
                1 => *GNS_DIAMOND_CONTRACT_ARBITRUM,
                2 => *GNS_DIAMOND_CONTRACT_ARBITRUM,
                3 => *GNS_DIAMOND_CONTRACT_ARBITRUM,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Incorrect collateral value!".to_string(),
                    }));
                }
            };
            chain_id = ARBITRUM_ID;
            uri = ARBITRUM_URI;
        }
        421614 => {
            diamond_contract = match collateral_token {
                1 => *GNS_DIAMOND_CONTRACT_SEPOLIA,
                2 => *GNS_DIAMOND_CONTRACT_SEPOLIA,
                3 => *GNS_DIAMOND_CONTRACT_SEPOLIA,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Incorrect collateral value!".to_string(),
                    }));
                }
            };
            chain_id = ARBITRUM_SEPOLIA_ID;
            uri = ARBITRUM_SEPOLIA_URI;
        }
        _ => {
            return Err(Box::new(ChainError {
                message: "Unknown chain!".to_string(),
            }));
        }
    }
    Ok((diamond_contract, chain_id, uri))
}