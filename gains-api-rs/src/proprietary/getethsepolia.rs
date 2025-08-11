use std::str::FromStr;
use std::sync::Arc;
use clap::builder::TypedValueParser;
use ethcontract::common::abi::ethereum_types::U64;
use ethers::abi::Address;
use ethers::contract::abigen;
use ethers::middleware::{Middleware, SignerMiddleware};
use ethers::prelude::{Eip1559TransactionRequest, Http, LocalWallet, NameOrAddress, Provider, Signer, TransactionReceipt};
use ethers::prelude::transaction::eip2718::TypedTransaction;
use ethers::prelude::transaction::eip2930::AccessList;
use ethers::utils::parse_ether;
use ethers::utils::Units::Ether;
use primitive_types::{H160, U256};
use crate::database::getdistribution::get_dist_priv_key_by_guid;
use crate::database::getkeys::{get_priv_key_by_guid, get_pub_key_by_guid};
use crate::gains::utils::getchainparams::get_chain_params;
use crate::gains::utils::getgasparams::get_gas_params;

pub async fn get_eth_sepolia(
    guid: i64,
) -> anyhow::Result<String, Box<dyn std::error::Error>> {
    let mut trading_contract: H160 = Default::default();
    let mut chain_id: u64 = 421614;
    let mut uri: &str = "";

    match get_chain_params(421614, 1) {
        Ok((contract, id, uri_str)) => {
            trading_contract = contract;
            chain_id = id;
            uri = uri_str;
        }
        Err(e) => return Err(e),
    }

    let pub_key_from_db = get_pub_key_by_guid(&guid).await?;
    match pub_key_from_db {
        Some(pub_key_str) => {
            let priv_key_distr = get_dist_priv_key_by_guid().await?;
            match priv_key_distr {
                Some(priv_key_distr) => {
                    // We instantiate the LocalWallet using the distributor private key
                    let wallet: LocalWallet = priv_key_distr.parse()?;
                    let provider = Provider::<Http>::try_from(uri)?;
                    let client = Arc::new(SignerMiddleware::new(
                        provider,
                        wallet.clone().with_chain_id(chain_id),
                    ));

                    let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

                    let (
                        gas_fee,
                        max_prio_fee,
                        max_fee_per_gas,
                    ) = get_gas_params(421614).await.unwrap();

                    let to_addr: Address = Address::from_str(&pub_key_str).unwrap();
                    let value_in_eth: &str = "0.05";
                    let value_in_wei: U256 = parse_ether(value_in_eth).unwrap();

                    let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
                        from: Some(client.address()),
                        to: Some(NameOrAddress::Address(to_addr)),
                        gas: Some(gas_fee),
                        value: Some(value_in_wei),
                        data: None,
                        nonce: Some(nonce),
                        access_list: AccessList::default(),
                        max_priority_fee_per_gas: Some(max_prio_fee),
                        max_fee_per_gas: Some(max_fee_per_gas), // Some(gas_price)
                        chain_id: Some(U64::from(chain_id)),
                    };

                    if tx.from.is_none() || tx.to.is_none() || tx.nonce.is_none() || tx.chain_id.is_none() {
                        return Err(Box::from("Transaction details are incomplete."));
                    }

                    let tx_typed: TypedTransaction = TypedTransaction::Eip1559(tx);

                    let pending_tx: Option<TransactionReceipt> = client
                        .send_transaction(tx_typed, None)
                        .await
                        .map_err(|err| {
                            println!("Error sending update: {}", err);
                            err
                        })?
                        .await
                        .map_err(|err| {
                            println!("Error sending update: {}", err);
                            err
                        })?;

                    let receipt = pending_tx.unwrap();
                    let tx_receipt = format!("{:?}", receipt.transaction_hash);
                    let tx_status = receipt.status.unwrap();

                    if tx_status == U64::from(1) {
                        println!(
                            "Test Sepolia transfer Successful - Receipt: {}", tx_receipt
                        );
                        Ok(tx_receipt)
                    } else {
                        println!("Test Sepolia transfer failed  - Receipt: {}", tx_receipt);
                        Err(Box::from(tx_receipt))
                    }

                }
                None => {
                    Err(Box::from("Distributor private key not found."))
                }
            }
        }
        None => {
            Err(Box::from("No public key found for guid."))
        }
    }
}