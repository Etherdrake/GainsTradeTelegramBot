use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::testnet::getfreedai::get_free_dai;

pub async fn handle_get_free_dai(
    Json(payload): Json<crate::routes::testnet::payloads::GetFreeDaiPayload>,
) -> impl IntoResponse {
    match get_free_dai(
        payload.guid,
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
            error!("Error in handle_get_free_dai: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}