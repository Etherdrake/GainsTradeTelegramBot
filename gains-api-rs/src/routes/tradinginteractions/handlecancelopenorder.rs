use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradinginteractions::write::cancelopenorder::cancel_open_order;

pub async fn handle_cancel_open_order(
    Json(payload): Json<crate::routes::tradinginteractions::tradinginteractionspayloads::CancelOpenOrderPayload>,
) -> impl IntoResponse {
    match cancel_open_order(
        payload.guid,
        payload.index,
        payload.collateral_token,
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
            error!("Error in handle_cancel_open_order: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}