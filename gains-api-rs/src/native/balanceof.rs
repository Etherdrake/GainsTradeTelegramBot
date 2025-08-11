use anyhow::{Context, Result};
use clap::Parser;

use std::str::FromStr;
use std::sync::Arc;
use ethers::contract::abigen;
use ethers::prelude::{Address, Http, Provider};
use ethers::utils::{format_units, parse_units};
use crate::blockchain_constants::chainids::{ARBITRUM_ID, POLYGON_ID};
use crate::blockchain_constants::token_constants::{DAI_ADDRESS_ARBITRUM, DAI_ADDRESS_POLYGON, DAI_ADDRESS_SEPOLIA, USDC_ADDRESS_ARBITRUM, USDC_ADDRESS_POLYGON, USDC_ADDRESS_SEPOLIA, WETH_ADDRESS_ARBITRUM, WETH_ADDRESS_POLYGON};
use crate::gains::errors::{ChainError, CollateralError};
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI, ARBITRUM_URI, POLYGON_URI};
use crate::routes::native::payloads::{BalanceOfPayload};


// Check token balance of wallet for one of the collateral tokens
pub async fn balance_of(payload: BalanceOfPayload) -> Result<String, Box<dyn std::error::Error>> {
    let (token_address, uri, chain_id) = match payload.chain {
        137 => {
            let token_address = match payload.collateral {
                1 => *DAI_ADDRESS_POLYGON,
                2 => *WETH_ADDRESS_POLYGON,
                3 => *USDC_ADDRESS_POLYGON,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Wrong collateral!".parse().unwrap(),
                    }))
                }
            };
            (token_address, POLYGON_URI, POLYGON_ID)
        },
        42161 => {
            let token_address = match payload.collateral {
                1 => *DAI_ADDRESS_ARBITRUM,
                2 => *WETH_ADDRESS_ARBITRUM,
                3 => *USDC_ADDRESS_ARBITRUM,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Wrong collateral!".parse().unwrap(),
                    }))
                }
            };
            (token_address, ARBITRUM_URI, ARBITRUM_ID)
        },
        421614 => {
            let token_address = match payload.collateral {
                1 => *DAI_ADDRESS_SEPOLIA,
                2 => *USDC_ADDRESS_SEPOLIA,
                3 => *USDC_ADDRESS_SEPOLIA,
                _ => {
                    return Err(Box::new(CollateralError {
                        message: "Wrong collateral!".parse().unwrap(),
                    }))
                }
            };
            (token_address, ARBITRUM_SEPOLIA_URI, 421614) // replace "sepolia URI" with the actual Sepolia URI constant
        },
        _ => {
            return Err(Box::new(ChainError {
                message: "Wrong chain!".parse().unwrap(),
            }));
        }
    };

    let pubkey: Address = Address::from_str(&*payload.public_key).unwrap();
    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    abigen!(
        ERC20,
        "./abi/ERC20.abi.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let contract = ERC20::new(token_address.clone(), Arc::new(client.clone()));

    // Call balanceOf
    let value = contract.balance_of(pubkey).call().await?;

    // Convert balance to token units
    let token_decimals = match payload.collateral {
        1 => 18,  // DAI has 18 decimals
        2 => 18,  // WETH has 18 decimals
        3 => 6,   // USDC has 6 decimals
        _ => {
            return Err(Box::new(CollateralError {
                message: "Wrong collateral!".parse().unwrap(),
            }))
        }
    };

    // We use format to get the right string
    let token_balance = format_units(value, token_decimals)?.to_string();

    Ok(token_balance)
}