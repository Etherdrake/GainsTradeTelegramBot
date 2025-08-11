use axum::http::StatusCode;
use axum::Json;
use num_format::Locale::pa;
use serde::Deserialize;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getpendingorders::get_pending_orders;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{PendingOrderStruct};

pub async fn handle_get_pending_orders(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetPendingOrdersPayload>,
) -> (StatusCode, Json<Result<Vec<PendingOrderStruct>, ErrorMessageResponse>>) {
    match get_pending_orders(
        payload.trader,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(all_pending_orders) => (StatusCode::OK, Json(Ok(all_pending_orders))),
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

