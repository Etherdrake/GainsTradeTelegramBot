use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::pairsstorage::getpairmaxleverage::get_pair_max_leverage;
use crate::routes::pairsstorage::payloads::GetPairMaxLeverage;

pub async fn handle_get_pair_max_leverage(
    Json(payload): Json<GetPairMaxLeverage>,
) -> impl IntoResponse {
    match get_pair_max_leverage(
        payload.index,
        payload.chain,
    ).await {
        Ok(orders) => (
            StatusCode::OK,
            Json(Ok(orders)),
        ),
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            error!("Error in handle_get_eth_sepolia: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}