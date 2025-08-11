use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use log::error;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::native::approve::approve;
use crate::native::balanceof::balance_of;

pub async fn handle_approve(
    Json(payload): Json<crate::routes::native::payloads::ApprovePayload>,
) -> impl IntoResponse {
    match approve(
        payload.guid,
        payload.owner,
        payload.token,
        payload.spender,
        payload.amount,
        payload.collateral,
        payload.chain
    ).await {
        Ok(approve) => (
            StatusCode::OK,
            Json(Ok(approve)),
        ),
        Err(e) => {
            let error_message = match e.downcast_ref::<CollateralError>() {
                Some(collateral_error) => format!("Collateral error: {}", collateral_error),
                None => match e.downcast_ref::<ChainError>() {
                    Some(chain_error) => format!("Chain error: {}", chain_error),
                    None => format!("An unknown error occurred: {}", e),
                },
            };
            error!("Error in handle_approve: {}", error_message);
            (
                StatusCode::BAD_REQUEST,
                Json(Err(ErrorMessageResponse { message: error_message })),
            )
        }
    }
}