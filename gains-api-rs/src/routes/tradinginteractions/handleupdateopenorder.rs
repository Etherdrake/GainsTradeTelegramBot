use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradinginteractions::write::bigtypes::UpdateOpenOrderPayloadBig;
use crate::gains::tradinginteractions::write::updateopenorder::update_open_order;

pub async fn handle_update_open_order(
    Json(payload): Json<crate::routes::tradinginteractions::tradinginteractionspayloads::UpdateOpenOrderPayload>,
) -> impl IntoResponse {
    payload.print();
    let payload_big: UpdateOpenOrderPayloadBig = payload.into();
    match update_open_order(
        payload_big.guid,
        payload_big.index,
        payload_big.trigger_price,
        payload_big.tp,
        payload_big.sl,
        payload_big.max_slippage_p,
        payload_big.collateral_token,
        payload_big.chain,
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
            error!("Error in handle_update_open_order: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}