use axum::http::StatusCode;
use axum::{async_trait, Json};
use axum::handler::Handler;
use axum::response::IntoResponse;
use axum_macros::debug_handler;
use log::error;
use rust_decimal::prelude::ToPrimitive;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::testnet::getfreedai::get_free_dai;
use crate::native::getgasbalance::get_gas_balance;
use crate::proprietary::getethsepolia::get_eth_sepolia;
use crate::routes::proprietary::payloads::GetFundsPayload;

pub async fn handle_get_funds(
    Json(payload): Json<GetFundsPayload>,
) -> impl IntoResponse {
    let chain: u64 = 421614.to_u64().unwrap();
    match get_free_dai(payload.guid, chain).await {
        Ok(_) => {
            match get_eth_sepolia(payload.guid).await {
                Ok(tx_receipt) => {
                    (StatusCode::OK,
                     Json(Ok(tx_receipt)))
                }
                Err(e) => {
                    let error_message = match e.downcast_ref::<CollateralError>() {
                        Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                        None => match e.downcast_ref::<ChainError>() {
                            Some(chain_error) => format!("Chain error: {}", chain_error),
                            None => format!("An unknown error occurred: {}", e),
                        },
                    };
                    error!("Error in get_eth_sepolia: {}", e.to_string());
                    (StatusCode::BAD_REQUEST, Json(Err(ErrorMessageResponse { message: error_message })))
                }
            }
        },
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            error!("Error in handle_get_funds: {}", e.to_string());
            (StatusCode::BAD_REQUEST, Json(Err(ErrorMessageResponse { message: error_message })))
        }
    }
}
