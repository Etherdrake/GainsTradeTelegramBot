use serde::Serialize;

#[derive(Serialize)]
pub struct ErrorMessageResponse {
    pub(crate) message: String,
}