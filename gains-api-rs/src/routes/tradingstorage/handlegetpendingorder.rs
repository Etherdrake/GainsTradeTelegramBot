use axum::http::StatusCode;
use axum::Json;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getpendingorder::get_pending_order;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{PendingOrderStruct};

pub async fn handle_get_pending_order(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetPendingOrderPayload>,
) -> (StatusCode, Json<Result<PendingOrderStruct, ErrorMessageResponse>>) {
    match get_pending_order(
        payload.trader,
        payload.order_id,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(pending_order) => (StatusCode::OK, Json(Ok(pending_order))),
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}
