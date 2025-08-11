use bson::{Bson, doc};
use mongodb::{Client, Collection};
use mongodb::options::FindOneOptions;
use std::env;

pub async fn get_dist_priv_key_by_guid() -> mongodb::error::Result<Option<String>> {
    let mongo_uri = env::var("MONGO_URI").expect("MONGO_URI environment variable not set");
    let client = Client::with_uri_str(mongo_uri).await?; // TODO: These values are hardcodeed to a localhost instance
    let database = client.database("secret");
    let collection: Collection<bson::Document> = database.collection("wallet");

    // Query for the user with the given guid
    let filter = doc! {"name": "HootSepoliaDistribution"};

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
