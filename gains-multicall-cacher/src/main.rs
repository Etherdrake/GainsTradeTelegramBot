use std::collections::HashMap;
use std::time::Duration;
use redis::{AsyncCommands, Client, IntoConnectionInfo};
use redis::aio::MultiplexedConnection;
use tokio::time::sleep;
use crate::multicalls::{get_pair_max_leverage, get_pair_min_leverage};

mod multicalls;

#[tokio::main]
async fn main() {
    let rdb_client = create_redis_connection().await.unwrap();
    loop {
        let min_map = get_pair_min_leverage("arbitrum").await.unwrap();
        write_to_redis(rdb_client.clone(), min_map, "arbitrum").await.expect("Writing min_map to Redis failed! Exiting...");
        let max_map = get_pair_max_leverage("arbitrum").await.unwrap();
        write_to_redis(rdb_client.clone(), max_map, "arbitrum").await.expect("Writing min_map to Redis failed! Exiting...");
        let min_map = get_pair_min_leverage("polygon").await.unwrap();
        write_to_redis(rdb_client.clone(), min_map, "polygon").await.expect("Writing min_map to Redis failed! Exiting...");
        let max_map = get_pair_max_leverage("polygon").await.unwrap();
        write_to_redis(rdb_client.clone(), max_map, "polygon").await.expect("Writing min_map to Redis failed! Exiting...");

        sleep(Duration::from_secs(300)).await;
    }
}


async fn write_to_redis(mut client: redis::aio::MultiplexedConnection, map: HashMap<i32, u64>, chain: &str) -> redis::RedisResult<()> {
    let map_str = map.clone();
    for (key, value) in map {
        client.set(chain.to_owned() + ":" + &*key.to_string(), value).await?;
    }
    println!("Writes complete, written {:?} keys!", map_str.len());
    Ok(())
}

async fn create_redis_connection() -> Result<MultiplexedConnection, redis::RedisError> {
    // We are going to connect to logical database 6
    let connection_info = "redis://127.0.0.1:6379/6".to_string().into_connection_info()?;
    let client = Client::open(connection_info)?;
    let connection = client.get_multiplexed_tokio_connection().await?;
    Ok(connection)
}