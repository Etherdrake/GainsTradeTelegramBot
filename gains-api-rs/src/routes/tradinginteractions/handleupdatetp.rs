use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradinginteractions::write::bigtypes::UpdateTpPayloadBig;
use crate::gains::tradinginteractions::write::updatesl::update_stop_loss;
use crate::routes::tradinginteractions::handleupdatesl::handle_update_sl;

pub async fn handle_update_tp(
    Json(payload): Json<crate::routes::tradinginteractions::tradinginteractionspayloads::UpdateTpPayload>,
) -> impl IntoResponse {
    let payload_big: UpdateTpPayloadBig = payload.into();
    match update_stop_loss(
        payload_big.guid,
        payload_big.index,
        payload_big.new_tp,
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
            error!("Error in handle_update_tp: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}