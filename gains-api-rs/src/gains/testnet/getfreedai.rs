use std::sync::Arc;
use ethcontract::common::abi::ethereum_types::U64;
use ethers::abi::AbiEncode;
use ethers::contract::abigen;
use ethers::middleware::SignerMiddleware;
use ethers::prelude::{Eip1559TransactionRequest, Http, LocalWallet, Middleware, NameOrAddress, Provider, Signer, TransactionReceipt};
use ethers::prelude::transaction::eip2718::TypedTransaction;
use ethers::prelude::transaction::eip2930::AccessList;
use primitive_types::{H160, U256};
use crate::blockchain_constants::chainids::{ARBITRUM_SEPOLIA_ID};
use crate::blockchain_constants::token_constants::DAI_ADDRESS_SEPOLIA;
use crate::database::getkeys::get_priv_key_by_guid;
use crate::gains::utils::getgasparams::get_gas_params;
use crate::nodes::constants::{ARBITRUM_SEPOLIA_URI};

pub async fn get_free_dai(
    guid: i64,
    chain: u64,
) -> anyhow::Result<String, Box<dyn std::error::Error>> {
    let mut trading_contract: H160;
    let mut chain_id: u64;
    let mut uri: &str;

    trading_contract = *DAI_ADDRESS_SEPOLIA;
    chain_id = ARBITRUM_SEPOLIA_ID;
    uri = ARBITRUM_SEPOLIA_URI;

    // This is the wallet PrivateKey
    let priv_key_from_db = get_priv_key_by_guid(&guid).await?;
    match priv_key_from_db {
        Some(priv_key_str) => {
            let wallet: LocalWallet = priv_key_str.parse()?;
            let provider = Provider::<Http>::try_from(uri)?;
            let client = Arc::new(SignerMiddleware::new(
                provider,
                wallet.clone().with_chain_id(chain_id),
            ));

            abigen!(GainsNetworkCollateralDAI, "./gains_abi/testnet/GainsNetworkCollateralDAI_abi.json");

            let call: GetFreeDaiCall = GetFreeDaiCall {
            };

            let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

            let (
                gas_fee,
                max_prio_fee,
                max_fee_per_gas) = get_gas_params(chain).await.unwrap();


            let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
                from: Some(client.address()),
                to: Some(NameOrAddress::Address(trading_contract)),
                gas: Some(gas_fee),
                value: None,
                data: Some(call.encode().into()),
                nonce: Some(nonce),
                access_list: AccessList::default(),
                max_priority_fee_per_gas: Some(max_prio_fee),
                max_fee_per_gas: Some(max_fee_per_gas), // Some(gas_price)
                chain_id: Some(U64::from(chain_id)),
            };

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
                    "Get Free DAI Successful - Receipt: {},", tx_receipt
                );
                Ok(tx_receipt)
            } else {
                println!("Get Free DAI Failed - Receipt: {}", tx_receipt);
                Err(Box::from(tx_receipt))
            }
        }
        None => Err(Box::new(std::io::Error::new(
            std::io::ErrorKind::Other,
            format!("No private key found for guid."),
        ))),
    }
}
