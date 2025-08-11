### Read Functions

1. **getAddresses**
    - **Description**: Retrieves addresses stored in the contract.
    - **Parameters**: None
    - **Returns**: Tuple containing addresses (`gns` and `gnsStaking`).
    - **State Mutability**: View

2. **getAllyDetails**
    - **Description**: Retrieves details of a specific ally.
    - **Parameters**: `_ally` (address)
    - **Returns**: Tuple containing ally details (referrers referred, volume referred, pending rewards, total rewards, total rewards value, active status).
    - **State Mutability**: View

3. **getReferralsAllyFeeP**
    - **Description**: Retrieves the ally fee percentage for referrals.
    - **Parameters**: None
    - **Returns**: Ally fee percentage.
    - **State Mutability**: View

4. **getReferralsOpenFeeP**
    - **Description**: Retrieves the open fee percentage for referrals.
    - **Parameters**: None
    - **Returns**: Open fee percentage.
    - **State Mutability**: View

5. **getReferralsStartReferrerFeeP**
    - **Description**: Retrieves the start referrer fee percentage for referrals.
    - **Parameters**: None
    - **Returns**: Start referrer fee percentage.
    - **State Mutability**: View

6. **getReferralsTargetVolumeUsd**
    - **Description**: Retrieves the target volume in USD for referrals.
    - **Parameters**: None
    - **Returns**: Target volume in USD.
    - **State Mutability**: View

7. **getReferrerDetails**
    - **Description**: Retrieves details of a specific referrer.
    - **Parameters**: `_referrer` (address)
    - **Returns**: Tuple containing referrer details (ally, traders referred, volume referred, pending rewards, total rewards, total rewards value, active status).
    - **State Mutability**: View

8. **getReferrerFeeP**
    - **Description**: Retrieves the fee percentage for a specific referrer.
    - **Parameters**: `_pairOpenFeeP` (uint256), `_volumeReferredUsd` (uint256)
    - **Returns**: Fee percentage.
    - **State Mutability**: View

9. **getReferrersReferred**
    - **Description**: Retrieves addresses of referrers referred by a specific ally.
    - **Parameters**: `_ally` (address)
    - **Returns**: Array of referrer addresses.
    - **State Mutability**: View

10. **getTraderActiveReferrer**
    - **Description**: Retrieves the active referrer of a specific trader.
    - **Parameters**: `_trader` (address)
    - **Returns**: Address of the active referrer.
    - **State Mutability**: View

11. **getTraderLastReferrer**
    - **Description**: Retrieves the last referrer of a specific trader.
    - **Parameters**: `_trader` (address)
    - **Returns**: Address of the last referrer.
    - **State Mutability**: View

12. **getTradersReferred**
    - **Description**: Retrieves addresses of traders referred by a specific referrer.
    - **Parameters**: `_account` (address)
    - **Returns**: Array of trader addresses.
    - **State Mutability**: View

13. **hasRole**
    - **Description**: Checks if an account has a specific role.
    - **Parameters**: `_account` (address), `_role` (enum IAddressStore.Role)
    - **Returns**: Boolean indicating role existence.
    - **State Mutability**: View

### Write Functions

1. **claimAllyRewards**
    - **Description**: Claims ally rewards.
    - **Parameters**: None
    - **State Mutability**: Non-payable

2. **claimReferrerRewards**
    - **Description**: Claims referrer rewards.
    - **Parameters**: None
    - **State Mutability**: Non-payable

3. **distributeReferralReward**
    - **Description**: Distributes referral rewards to a specific trader.
    - **Parameters**: `_trader` (address), `_volumeUsd` (uint256), `_pairOpenFeeP` (uint256), `_gnsPriceUsd` (uint256)
    - **Returns**: Distributed reward amount.
    - **State Mutability**: Non-payable

4. **initialize**
    - **Description**: Initializes the contract.
    - **Parameters**: `_rolesManager` (address)
    - **State Mutability**: Non-payable

5. **initializeReferrals**
    - **Description**: Initializes referral fee settings.
    - **Parameters**: `_allyFeeP` (uint256), `_startReferrerFeeP` (uint256), `_openFeeP` (uint256), `_targetVolumeUsd` (uint256)
    - **State Mutability**: Non-payable

6. **registerPotentialReferrer**
    - **Description**: Registers a potential referrer for a trader.
    - **Parameters**: `_trader` (address), `_referrer` (address)
    - **State Mutability**: Non-payable

7. **setRoles**
    - **Description**: Sets roles for multiple accounts.
    - **Parameters**: `_accounts` (address[]), `_roles` (enum IAddressStore.Role[]), `_values` (bool[])
    - **State Mutability**: Non-payable

8. **unwhitelistAllies**
    - **Description**: Removes whitelist status from multiple allies.
    - **Parameters**: `_allies` (address[])
    - **State Mutability**: Non-payable

9. **unwhitelistReferrers**
    - **Description**: Removes whitelist status from multiple referrers.
    - **Parameters**: `_referrers` (address[])
    - **State Mutability**: Non-payable

10. **updateAllyFeeP**
    - **Description**: Updates the ally fee percentage for referrals.
    - **Parameters**: `_value` (uint256)
    - **State Mutability**: Non-payable

11. **updateReferralsOpenFeeP**
    - **Description**: Updates the open fee percentage for referrals.
    - **Parameters**: `_value` (uint256)
    - **State Mutability**: Non-payable

12. **updateReferralsTargetVolumeUsd**
    - **Description**: Updates the target volume in USD for referrals.
    - **Parameters**: `_value` (uint256)
    - **State Mutability**: Non-payable

13. **updateStartReferrerFeeP**
    - **Description**: Updates the start referrer fee percentage for referrals.
    - **Parameters**: `_value` (uint256)
    - **State Mutability**: Non-payable

14. **whitelistAllies**
    - **Description**: Whitelists multiple allies.
    - **Parameters**: `_allies` (address[])
    - **State Mutability**: Non-payable

15. **whitelistReferrers**
    - **Description**: Whitelists multiple referrers with specified allies.
    - **Parameters**: `_referrers` (address[]), `_allies` (address[])
    - **State Mutability**: Non-payable
