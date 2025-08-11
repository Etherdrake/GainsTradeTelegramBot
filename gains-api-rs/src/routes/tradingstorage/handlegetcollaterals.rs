use axum::http::StatusCode;
use axum::Json;
use crate::errors::ErrorMessageResponse;
use crate::gains::errors::{ChainError, CollateralError};
use crate::gains::tradingstorage::read::getcollaterals::get_collaterals;
use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::{CollateralStruct, PendingOrderStruct};

pub async fn handle_get_collaterals(
    Json(payload): Json<crate::routes::tradingstorage::payloads::GetCollateralsPayload>,
) -> (StatusCode, Json<Result<Vec<CollateralStruct>, ErrorMessageResponse>>) {
    match get_collaterals(
        payload.collateral_token,
        payload.chain,
    ).await {
        Ok(alltrades) => (StatusCode::OK, Json(Ok(alltrades))),
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
