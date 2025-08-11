### Read Functions

1. **getAddresses**
    - **Description:** Retrieves the addresses of certain contracts from the trading contract.
    - **Parameters:** None
    - **Returns:** A tuple containing addresses of certain contracts.

2. **getFeeTier**
    - **Description:** Retrieves fee tier information based on the provided index.
    - **Parameters:** `_feeTierIndex` (uint256): Index of the fee tier to retrieve.
    - **Returns:** Fee tier details including fee multiplier and points threshold.

3. **getFeeTiersCount**
    - **Description:** Retrieves the total number of fee tiers.
    - **Parameters:** None
    - **Returns:** Total count of fee tiers.

4. **getFeeTiersTraderDailyInfo**
    - **Description:** Retrieves daily fee tier information for a specific trader.
    - **Parameters:**
        - `_trader` (address): Address of the trader.
        - `_day` (uint32): Specific day to retrieve information for.
    - **Returns:** Fee multiplier cache and points for the trader on the specified day.

5. **getFeeTiersTraderInfo**
    - **Description:** Retrieves overall fee tier information for a specific trader.
    - **Parameters:** `_trader` (address): Address of the trader.
    - **Returns:** Last updated day and trailing points for the trader.

6. **getGroupVolumeMultiplier**
    - **Description:** Retrieves the volume multiplier for a specific group.
    - **Parameters:** `_groupIndex` (uint256): Index of the group.
    - **Returns:** Volume multiplier for the group.

7. **hasRole**
    - **Description:** Checks whether an account has a specific role.
    - **Parameters:**
        - `_account` (address): The address of the account to check.
        - `_role` (enum IAddressStore.Role): The role to check.
    - **Returns:** A boolean indicating whether the account has the specified role.

### Write Functions

1. **initialize**
    - **Description:** Initializes the fee tiers contract.
    - **Parameters:** None
    - **Returns:** None

2. **initializeFeeTiers**
    - **Description:** Initializes fee tiers with provided data.
    - **Parameters:**
        - `_groupIndices` (uint256[]): Indices of groups for fee tiers.
        - `_groupVolumeMultipliers` (uint256[]): Volume multipliers for groups.
        - `_feeTiersIndices` (uint256[]): Indices of fee tiers.
        - `_feeTiers` (struct IFeeTiers.FeeTier[]): Fee tiers data.
    - **Returns:** None

3. **setFeeTiers**
    - **Description:** Sets fee tiers with provided data.
    - **Parameters:**
        - `_feeTiersIndices` (uint256[]): Indices of fee tiers.
        - `_feeTiers` (struct IFeeTiers.FeeTier[]): Fee tiers data.
    - **Returns:** None

4. **setGroupVolumeMultipliers**
    - **Description:** Sets volume multipliers for groups.
    - **Parameters:**
        - `_groupIndices` (uint256[]): Indices of groups.
        - `_groupVolumeMultipliers` (uint256[]): Volume multipliers for groups.
    - **Returns:** None

5. **setRoles**
    - **Description:** Sets roles for multiple accounts.
    - **Parameters:**
        - `_accounts` (address[]): List of account addresses.
        - `_roles` (enum IAddressStore.Role[]): List of role values.
        - `_values` (bool[]): List of boolean values indicating whether the role is assigned.
    - **Returns:** None

6. **updateTraderPoints**
    - **Description:** Updates points for a trader based on trading activity.
    - **Parameters:**
        - `_trader` (address): Address of the trader.
        - `_volumeUsd` (uint256): Volume in USD for the trader.
        - `_pairIndex` (uint256): Index of the trading pair.
    - **Returns:** None
