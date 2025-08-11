use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getalltrades::get_all_trades;

pub async fn handle_get_all_trades(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetAllTradesPayload>,
) -> impl IntoResponse {
    match get_all_trades(
        payload.offset,
        payload.limit,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(alltrades) => (
            StatusCode::OK,
            Json(Ok(alltrades)),
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
