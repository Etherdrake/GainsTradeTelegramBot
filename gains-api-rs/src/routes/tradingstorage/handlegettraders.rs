use axum::http::StatusCode;
use axum::Json;
use ethers::prelude::Address;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::gettraders::get_traders;

pub async fn handle_get_traders(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetTradersPayload>,
) -> (StatusCode, Json<Result<Vec<Address>, ErrorMessageResponse>>) {
    match get_traders(
        payload.offset,
        payload.limit,
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(traders) => (StatusCode::OK, Json(Ok(traders))),
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
