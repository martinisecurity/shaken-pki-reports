# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NetNumber SHAKEN Root Intermediate CA 1

Tested At: 07 Dec 22 18:54 UTC\
Initial Validity Period: 4380 day(s)\
Remaining Validity Period: 3946 day(s)\
Subject: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1

[View certificate details](https://understandingwebpki.com/?cert=MIIDNzCCApigAwIBAgIJALHQvycN%2Fmz1MAoGCCqGSM49BAMDMIGBMSMwIQYDVQQDDBpOZXROdW1iZXIgU0hBS0VOIFJvb3QgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIxMDkyOTEzMjI1N1oXDTMzMDkyNjEzMjI1N1owgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE8tCNFPuiCKvElj73ghrBgo%2F44JzXhjDtL2L8hptKs%2Fa%2BbxF5mYquRHsD0ckmgpD4Mz28XrZ9ZFu0DRdBxE9%2BQcKvtJ7CzXD3LCvqWLqROg1Icp3HzmnzXJA%2Ft6iRpVARo4HMMIHJMBIGA1UdEwEB%2FwQIMAYBAf8CAQEwDgYDVR0PAQH%2FBAQDAgIEMB8GA1UdIwQYMBaAFH%2FSxmcvsw47rbQTPz8vahAgvo8UMB0GA1UdDgQWBBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMDA4GMADCBiAJCAZArUIt30QwopGwY9J9tJe22P7Q45Hk76w7tvw0S9o2vxrKZK34w91hOj5ebS6zhHEp7cSLC4QQ5MuP4h9nUKbxOAkIBNqSI0Nk3GxM%2FARBaQXHwOZqysCF1L9%2Ft%2FsM%2F8eqT1AchVXg6RQvn3zJ0zH8mKwt8mN57vX3DWLcMZAlSsqkMLEI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_ca_signature_algorithm](../../ISSUES/e_atis_ca_signature_algorithm/README.md) | error | ATIS1000080 | STI certificates shall contain a Signature Algorithm field with the value 'ecdsa-with-SHA256' |
| [e_atis_ca_subject_public_key](../../ISSUES/e_atis_ca_subject_public_key/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Public Key Info field specifying a Public Key Algorithm of "id-ecPublicKey" and containing a 256-bit public key |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 07 Dec 22 18:54 UTC