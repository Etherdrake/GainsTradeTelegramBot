use std::sync::Arc;

// Check_dai_approved_gns returns the amount of dai approved to the GNS contract.
pub async fn check_dai_approved_gns(payload: CheckAllowancePayload) -> Result<String> {
    let uri: &str;
    let chain_id: u64;
    let contract_address: H160;
    let gns_address: H160;
    if payload.chain == "polygon" {
        chain_id = 137;
        uri = POLYGON_URI;
        gns_address = H160::from_str(GFARM_POLY).unwrap();
        contract_address = H160::from_str(DAI_ADDRESS_POLYGON).unwrap();
    } else if payload.chain == "arbitrum" {
        chain_id = 42161;
        uri = ARBITRUM_URI;
        contract_address = H160::from_str(DAI_ADDRESS_ARBITRUM).unwrap();
        gns_address = H160::from_str(GFARM_ARB).unwrap();
    } else {
        return Ok("error".to_string());
    }

    // Now, you can use this endpoint to initialize your provider
    let provider = Provider::<Http>::try_from(uri)?;
    let client = Arc::new(provider);

    abigen!(
        Allowance,
        "./abi/UChildDAI.json",
        event_derives(serde::Deserialize, serde::Serialize)
    );

    let contract = Allowance::new(contract_address, Arc::new(client.clone()));

    let owner_address = payload.trader_address;

    // Create tuple of owner and spender to check allowance
    let value = contract
        .allowance(owner_address, gns_address)
        .call()
        .await?;

    Ok(value.to_string())
}