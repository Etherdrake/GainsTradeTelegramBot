use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradinginteractions::write::bigtypes::UpdateSlPayloadBig;
use crate::gains::tradinginteractions::write::updatesl::update_stop_loss;

pub async fn handle_update_sl(
    Json(payload): Json<crate::routes::tradinginteractions::tradinginteractionspayloads::UpdateSlPayload>,
) -> impl IntoResponse {
    let payload_big: UpdateSlPayloadBig = payload.into();
    match update_stop_loss(
        payload_big.guid,
        payload_big.index,
        payload_big.new_sl,
        payload_big.chain,
        payload_big.collateral_token,
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
            error!("Error in handle_update_sl: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}