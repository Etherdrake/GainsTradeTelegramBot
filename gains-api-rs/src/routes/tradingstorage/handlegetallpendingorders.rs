use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getallpendingsorders::get_all_pending_orders;
use crate::errors::ErrorMessageResponse;
use crate::routes::tradingstorage::payloads::GetAllPendingOrdersPayload;

pub async fn handle_get_all_pending_orders(
    Json(payload): Json<GetAllPendingOrdersPayload>,
) -> impl IntoResponse {
    match get_all_pending_orders(
        payload.offset,
        payload.limit,
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
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}