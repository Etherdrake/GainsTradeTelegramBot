use crate::gains::tradingstorage::tradingstoragetypes::trading_storage::TradeStruct;
use bson::{doc, Bson};
use mongodb::options::FindOneOptions;
use mongodb::{Client, Collection};
use std::env;
pub async fn get_priv_key_by_guid(guid: &i64) -> mongodb::error::Result<Option<String>> {
    let mongo_uri = env::var("MONGO_URI").expect("MONGO_URI environment variable not set");
    let client = Client::with_uri_str(mongo_uri).await?;
    let database = client.database("hooterdb");
    let collection: Collection<bson::Document> = database.collection("users");

    // Query for the user with the given guid
    let filter = doc! {"_id": guid};

    // Create projection to only retrieve the "priv_key" field
    let projection = doc! {"private_key": 1};

    // Create a FindOneOptions with the projection
    let find_options = FindOneOptions::builder().projection(projection).build();

    let document = collection.find_one(filter, Some(find_options)).await?;

    match document {
        Some(doc) => {
            if let Some(priv_key) = doc.get("private_key").and_then(Bson::as_str) {
                Ok(Some(priv_key.to_string()))
            } else {
                Ok(None)
            }
        }
        None => Ok(None),
    }
}

// TODO: This uses some kind of localhost privkey which is incorrect, this should use an .env defined client uri instead.
pub async fn get_pub_key_by_guid(guid: &i64) -> mongodb::error::Result<Option<String>> {
    let mongo_uri = env::var("MONGO_URI").expect("MONGO_URI environment variable not set");
    let client = Client::with_uri_str(mongo_uri).await?;
    let database = client.database("hooterdb");
    let collection: Collection<bson::Document> = database.collection("users");

    // Query for the user with the given guid
    let filter = doc! {"_id": guid};

    // Create projection to only retrieve the "pub_key" field
    let projection = doc! {"public_key": 1};

    // Create a FindOneOptions with the projection
    let find_options = FindOneOptions::builder().projection(projection).build();

    let document = collection.find_one(filter, Some(find_options)).await?;

    match document {
        Some(doc) => {
            if let Some(pub_key) = doc.get("public_key").and_then(Bson::as_str) {
                Ok(Some(pub_key.to_string()))
            } else {
                Ok(None)
            }
        }
        None => Ok(None),
    }
}

pub async fn get_trade_settings_by_guid(guid: i64) -> mongodb::error::Result<Option<TradeStruct>> {
    let mongo_uri = env::var("MONGO_URI").expect("MONGO_URI environment variable not set");
    let client = Client::with_uri_str(mongo_uri).await?;
    let database = client.database("hooterdb");
    let collection: Collection<bson::document::Document> = database.collection("tradesettings");

    // Query for the trade with the given guid
    let filter = doc! {"_id": guid};
    let document = collection.find_one(filter, None).await?;

    match document {
        Some(doc) => {
            // Deserialize the MongoDB document directly into a TradeStruct
            let trade: TradeStruct = bson::from_document(doc)?;
            Ok(Some(trade))
        }
        None => Ok(None),
    }
}
