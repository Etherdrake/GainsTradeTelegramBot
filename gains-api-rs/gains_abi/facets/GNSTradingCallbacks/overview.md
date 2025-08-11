### Read Functions

#### 1. `getAddresses()`
- **Description**: Retrieves the addresses stored in the contract.
- **Parameters**: None
- **Returns**: Tuple containing addresses of `gns` and `gnsStaking`.
- **Visibility**: Public
- **State Mutability**: View

#### 2. `getPendingGovFeesCollateral(uint8 _collateralIndex)`
- **Description**: Retrieves pending government fees for a specific collateral index.
- **Parameters**: `_collateralIndex` (uint8): Index of the collateral.
- **Returns**: Pending government fees for the specified collateral.
- **Visibility**: Public
- **State Mutability**: View

#### 3. `getVaultClosingFeeP()`
- **Description**: Retrieves the percentage of vault closing fee.
- **Parameters**: None
- **Returns**: Percentage of vault closing fee.
- **Visibility**: Public
- **State Mutability**: View

#### 4. `hasRole(address _account, enum IAddressStore.Role _role)`
- **Description**: Checks if an account has a specific role.
- **Parameters**: `_account` (address): The account to check. `_role` (enum IAddressStore.Role): The role to check.
- **Returns**: True if the account has the specified role, otherwise false.
- **Visibility**: Public
- **State Mutability**: View

### Write Functions

#### 1. `claimPendingGovFees()`
- **Description**: Claims pending government fees.
- **Parameters**: None
- **Returns**: None
- **Visibility**: Public
- **State Mutability**: Non-payable

#### 2. `initialize(address _rolesManager)`
- **Description**: Initializes the contract with roles manager.
- **Parameters**: `_rolesManager` (address): Address of the roles manager contract.
- **Returns**: None
- **Visibility**: Public
- **State Mutability**: Non-payable

#### 3. `initializeCallbacks(uint8 _vaultClosingFeeP)`
- **Description**: Initializes callbacks with a specified vault closing fee percentage.
- **Parameters**: `_vaultClosingFeeP` (uint8): Percentage of vault closing fee.
- **Returns**: None
- **Visibility**: Public
- **State Mutability**: Non-payable

#### 4. `setRoles(address[] _accounts, enum IAddressStore.Role[] _roles, bool[] _values)`
- **Description**: Sets roles for multiple accounts.
- **Parameters**: `_accounts` (address[]): Array of accounts. `_roles` (enum IAddressStore.Role[]): Array of roles. `_values` (bool[]): Array of boolean values indicating whether to set or unset the roles.
- **Returns**: None
- **Visibility**: Public
- **State Mutability**: Non-payable

#### 5. `updateVaultClosingFeeP(uint8 _valueP)`
- **Description**: Updates the percentage of vault closing fee.
- **Parameters**: `_valueP` (uint8): Percentage of vault closing fee to be updated.
- **Returns**: None
- **Visibility**: Public
- **State Mutability**: Non-payable
