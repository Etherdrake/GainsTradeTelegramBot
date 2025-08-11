use std::sync::Arc;

// Approve DAI to gains trading
pub async fn approve_dai_gns(body: ApproveToDaiPayload) -> Result<String, Box<dyn std::error::Error>> {
    let mut contract_address: H160;
    let mut chain_id: u64;
    let mut uri: &str;
    let chain: &str;
    let dai_address: &str;

    println!("Chain Received {}", body.chain);

    if body.chain == "polygon" {
        contract_address = *GFARM_V5_CONTRACT_POLY_DAI;
        chain_id = POLYGON_ID;
        uri = POLYGON_URI;
        chain = "polygon";
        dai_address = DAI_ADDRESS_POLYGON;
    } else if body.chain == "arbitrum" {
        contract_address = *GFARM_V5_CONTRACT_ARB_DAI;
        chain_id = ARBITRUM_ID;
        uri = ARBITRUM_URI;
        chain = "arbitrum";
        dai_address = DAI_ADDRESS_ARBITRUM;
    } else {
        return Err(Box::new(ChainError {
            message: "Wrong chain!".parse().unwrap(),
        }));
    }

    let priv_key = get_priv_key_by_guid(&body.guid).await?;
    match priv_key {
        Some(priv_key) => {
            let wallet: LocalWallet = priv_key.parse()?;
            let provider = Provider::<Http>::try_from(uri)?;
            let clone_provider = provider.clone();
            let third_provider = clone_provider.clone();

            let client = Arc::new(SignerMiddleware::new(
                provider,
                wallet.clone().with_chain_id(chain_id),
            ));

            abigen!(
                ERC20,
                "./abi/ERC20.abi.json",
                event_derives(serde::Deserialize, serde::Serialize)
            );

            let dai_address = Address::from_str(DAI_ADDRESS_POLYGON).unwrap();

            let call: ApproveCall = ApproveCall {
                spender: contract_address,
                value: body.amount,
            };

            let nonce: U256 = client.get_transaction_count(client.address(), None).await?;

            let block_number = clone_provider
                .get_block_number()
                .await
                .expect("Failed to get block number");

            let block: Option<Block<Transaction>> = third_provider
                .get_block_with_txs(block_number)
                .await
                .expect("Failed to get block");

            let mut base_fee_clone: U256;

            if let Some(block) = block {
                // Access the baseFeePerGas from the block
                if let Some(base_fee) = block.base_fee_per_gas {
                    base_fee_clone = base_fee.clone();
                    println!("Base Fee: {:?}", base_fee);
                } else {
                    base_fee_clone = U256::from(parse_units("51716033060", "gwei")?);
                    println!("Base Fee not available in the block");
                }
            } else {
                base_fee_clone = U256::from(parse_units("1", "gwei")?);
                println!("Block not found");
            }

            let mut gas_fee: U256;
            let mut max_prio_fee: U256;
            let mut max_fee_per_gas: U256;

            if chain == "polygon" {
                gas_fee = Some(U256::from(3_000_000)).unwrap();
                max_prio_fee = Some(U256::from(parse_units("80", "gwei")?)).unwrap();
                max_fee_per_gas = Some(U256::from(parse_units("100", "gwei")?)).unwrap();
            } else if chain == "arbitrum" {
                gas_fee = base_fee_clone;
                max_prio_fee = U256::from(parse_units("0", "gwei")?);
                max_fee_per_gas = gas_fee + max_prio_fee;
            } else {
                return Err(Box::new(ChainError {
                    message: "Wrong chain!".parse().unwrap(),
                }));
            }

            let tx: Eip1559TransactionRequest = Eip1559TransactionRequest {
                from: Some(client.address()),
                to: Some(NameOrAddress::Address(contract_address)),
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
                println!("OpenTrade Successful - Receipt: {}", tx_receipt);
                std::prelude::rust_2015::Ok(tx_receipt)
            } else {
                println!("OpenTrade Failed - Receipt: {}", tx_receipt);
                Err(Box::new(std::io::Error::new(
                    std::io::ErrorKind::Other,
                    "Transaction status is not 1",
                )))
            }
        }
        None => Err(format!("No private key found for guid: {}", body.trader_address).into()),
    }
}
