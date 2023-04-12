# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID CA-1

Tested At: 12 Apr 23 01:46 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 2357 day(s)\
Subject: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID Root CA

[View certificate details](https://understandingwebpki.com/?cert=MIIE2TCCA8GgAwIBAgIUKLgn5Rb7gesY7PEvnRlFb183ZS0wDQYJKoZIhvcNAQELBQAwgYExLDAqBgNVBAMMI05ldXN0YXIgQ2VydGlmaWVkIENhbGxlciBJRCBSb290IENBMRkwFwYDVQQLDBB3d3cuY2NpZC5uZXVzdGFyMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzELMAkGA1UEBhMCVVMwHhcNMTkwOTIzMTMzMjE1WhcNMjkwOTIzMTMzMjE1WjB%2BMSkwJwYDVQQDDCBOZXVzdGFyIENlcnRpZmllZCBDYWxsZXIgSUQgQ0EtMTEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxCzAJBgNVBAYTAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuQpLqAhLMASoJZU7JqpdMyjEGTmOveIeZOV4bdsw84H9TWX4e9yYPMr6bn0PqxybKqlXdL4odx9vCcRPTK9mVJq%2B1bX3r9SpuzqdeE1%2BivGMPe2gKfsW7vAchUS3VokjoRyzT1gbnm0oaEKipQSThjM3XVLBw46hXixfQd9UQG1KQu6BDFepxH%2FDcMYYimh8z39ZFupHPe1m%2BChEWgydBRruH7eIWVXMbUwmfKbXjwyK%2FHP%2BUjG3zt5YfDPVbC0%2BU4mVNdlDrr4%2FYdCetHdNu%2FmUcfTg%2FHnPohplIorIM0V%2FR3cMoZ5XcNpAinexgi7E%2BTiUi3T3k3HRid1OebvsTQIDAQABo4IBSTCCAUUwEgYDVR0TAQH%2FBAgwBgEB%2FwIBADAfBgNVHSMEGDAWgBQDxFFTZk7gLnzaaVrRHhiUhp4JGzCBgAYIKwYBBQUHAQEEdDByMEoGCCsGAQUFBzAChj5odHRwOi8vY2FjZXJ0cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkUm9vdENBLmNydDAkBggrBgEFBQcwAYYYaHR0cDovL29jc3AuY2NpZC5uZXVzdGFyMA8GA1UdJQQIMAYGBFUdJQAwSwYDVR0fBEQwQjBAoD6gPIY6aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkUm9vdENBLmNybDAdBgNVHQ4EFgQUO7lcyzEXnEjP1Npm6422qNXhCfcwDgYDVR0PAQH%2FBAQDAgGGMA0GCSqGSIb3DQEBCwUAA4IBAQBJroHa0ic%2FT4YYGEaXM%2B1cZNbsmwDSHCa2%2B3s5R37HLLUnoLVMByt4YbyvEP41Ty3Q5Uu9Ahk%2BZxkJGeron9Fxi5xmrSUePbj%2FoTs9MuOTiQUhEHDTUWJi8BmAjCWOZqUmIF%2F%2Fqd1pclEG%2Bs5v0VtBO6H51eCuE%2B423OmGdkZ6vrGe2c1U0JemI1w6bnYeBl4AMXhuuS2BMAojgmCtmSQX9Ung9r4pgilU7yUV0ci0%2FC3zmGgFxynoMSqsqkQG4VZcs83dpxd2S7eAPKZSBbriDccKiTwDTlnRT5eGwLaqnLBlAAnzFZMxO0WzCCCQpFXbJ2p%2FS0J0lj%2FDNZfwhuAu)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_basic_constraints
- e_atis_ca_authority_key_identifier
- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_issuer_dn
- e_atis_ca_key_usage
- e_atis_ca_serial_number
- e_atis_ca_signature_algorithm
- e_atis_ca_subject
- e_atis_ca_subject_cn
- e_atis_ca_subject_key_identifier
- e_atis_ca_subject_public_key
- e_atis_ca_version
- e_us_cp_ca_key_usage_crl_sign

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Apr 23 01:46 UTC