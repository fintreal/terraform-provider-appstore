# Terraform Provider for App Store Connect

This Terraform provider enables you to manage App Store Connect resources using Infrastructure as Code. It provides resources for managing bundle identifiers and provisioning profiles in your Apple Developer account.

### Using Terraform Registry

Add the following to your Terraform configuration:

```hcl
terraform {
  required_providers {
    appstore = {
      source  = "fintreal/appstore"
    }
  }
}
```

## Provider Configuration

To use the provider, you need to configure it with your App Store Connect API credentials:

```hcl
provider "appstore" {
  appstore_key           = "APPSTORE_KEY"           # Your App Store API Key
  appstore_key_id        = "APPSTORE_KEY_ID"        # Your App Store API Key ID
  appstore_key_issuer_id = "APPSTORE_KEY_ISSUER_ID" # Your App Store Issuer ID
}
```

You can also set these values using environment variables:
- `APPSTORE_KEY`
- `APPSTORE_KEY_ID`
- `APPSTORE_KEY_ISSUER_ID`

## Resources

### Bundle Identifier

The `appstore_bundle_identifier` resource allows you to manage bundle identifiers for your apps.

```hcl
resource "appstore_bundle_identifier" "example" {
  name       = "My App Bundle Identifier"
  identifier = "com.example.app"
  platform   = "IOS"  # Options: IOS, MAC_OS
  capabilities = [
    "APPLE_ID_AUTH",
    "PUSH_NOTIFICATIONS",
    "IN_APP_PURCHASE"
    # See documentation for all available capabilities
  ]
}
```

#### Arguments

- `name` (Required) - The name of the bundle identifier
- `identifier` (Required) - The bundle identifier (e.g., "com.example.app")
- `platform` (Required) - The platform type (IOS, MAC_OS, or UNIVERSAL)
- `capabilities` (Optional) - List of app capabilities to enable

#### Attributes

- `id` - The unique identifier of the bundle identifier

### Provisioning Profile

The `appstore_provisioning_profile` resource allows you to manage provisioning profiles for your apps.

```hcl
resource "appstore_provisioning_profile" "example" {
  name                = "My App Provisioning Profile"
  type               = "IOS_APP_STORE"  # See documentation for all available types
  bundle_identifier_id = appstore_bundle_identifier.example.id
  certificate_ids    = ["CERTIFICATE_ID_1", "CERTIFICATE_ID_2"]
}
```

#### Arguments

- `name` (Required) - The name of the provisioning profile
- `type` (Required) - The type of provisioning profile (e.g., IOS_APP_STORE, IOS_APP_DEVELOPMENT)
- `bundle_identifier_id` (Required) - The ID of the associated bundle identifier
- `certificate_ids` (Required) - List of certificate IDs to include in the profile

#### Attributes

- `id` - The unique identifier of the provisioning profile
- `content_base64` - The base64-encoded content of the provisioning profile (sensitive)

## Supported Provisioning Profile Types

- IOS_APP_DEVELOPMENT
- IOS_APP_STORE
- IOS_APP_ADHOC
- IOS_APP_INHOUSE
- MAC_APP_DEVELOPMENT
- MAC_APP_STORE
- MAC_APP_DIRECT
- TVOS_APP_DEVELOPMENT
- TVOS_APP_STORE
- TVOS_APP_ADHOC
- TVOS_APP_INHOUSE
- MAC_CATALYST_APP_DEVELOPMENT
- MAC_CATALYST_APP_STORE
- MAC_CATALYST_APP_DIRECT

## Supported Bundle Identifier Capabilities

- ICLOUD
- IN_APP_PURCHASE
- GAME_CENTER
- PUSH_NOTIFICATIONS
- WALLET
- INTER_APP_AUDIO
- MAPS
- ASSOCIATED_DOMAINS
- PERSONAL_VPN
- APP_GROUPS
- HEALTHKIT
- HOMEKIT
- WIRELESS_ACCESSORY_CONFIGURATION
- APPLE_PAY
- DATA_PROTECTION
- SIRIKIT
- NETWORK_EXTENSIONS
- MULTIPATH
- HOT_SPOT
- NFC_TAG_READING
- CLASSKIT
- AUTOFILL_CREDENTIAL_PROVIDER
- ACCESS_WIFI_INFORMATION
- NETWORK_CUSTOM_PROTOCOL
- COREMEDIA_HLS_LOW_LATENCY
- SYSTEM_EXTENSION_INSTALL
- USER_MANAGEMENT
- APPLE_ID_AUTH
