# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NetNumber SHAKEN Root Intermediate CA 1

Tested At: 12 Feb 24 17:00 UTC\
Initial Validity Period: 4380 day(s)\
Remaining Validity Period: 3514 day(s)\
Subject: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root CA 1

[View certificate details](https://understandingwebpki.com/?cert=MIIDNzCCApigAwIBAgIJALHQvycN%2Fmz1MAoGCCqGSM49BAMDMIGBMSMwIQYDVQQDDBpOZXROdW1iZXIgU0hBS0VOIFJvb3QgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIxMDkyOTEzMjI1N1oXDTMzMDkyNjEzMjI1N1owgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE8tCNFPuiCKvElj73ghrBgo%2F44JzXhjDtL2L8hptKs%2Fa%2BbxF5mYquRHsD0ckmgpD4Mz28XrZ9ZFu0DRdBxE9%2BQcKvtJ7CzXD3LCvqWLqROg1Icp3HzmnzXJA%2Ft6iRpVARo4HMMIHJMBIGA1UdEwEB%2FwQIMAYBAf8CAQEwDgYDVR0PAQH%2FBAQDAgIEMB8GA1UdIwQYMBaAFH%2FSxmcvsw47rbQTPz8vahAgvo8UMB0GA1UdDgQWBBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwGgYDVR0gAQH%2FBBAwDjAMBgpghkgBhv8JAQEBMAoGCCqGSM49BAMDA4GMADCBiAJCAZArUIt30QwopGwY9J9tJe22P7Q45Hk76w7tvw0S9o2vxrKZK34w91hOj5ebS6zhHEp7cSLC4QQ5MuP4h9nUKbxOAkIBNqSI0Nk3GxM%2FARBaQXHwOZqysCF1L9%2Ft%2FsM%2F8eqT1AchVXg6RQvn3zJ0zH8mKwt8mN57vX3DWLcMZAlSsqkMLEI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_signature_algorithm_ca](../../ISSUES/e_atis_signature_algorithm_ca/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.10045.4.3.3 |
| [e_atis_ext_certificate_policies_ca](../../ISSUES/e_atis_ext_certificate_policies_ca/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_shaken_certificate_policies_id_ca](../../ISSUES/e_shaken_certificate_policies_id_ca/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_public_key_ca](../../ISSUES/e_atis_subject_public_key_ca/README.md) | error | ATIS1000080 | Subject Public Key Info field contains a public key that is not 256 bits |

### Not Effective

- e_atis_ext_not_specified_ca
- e_atis_serial_number_size_ca
- e_atis_subject_c_iso_ca
- e_atis_subject_key_identifier_size_ca
- e_atis_subject_o_required_ca

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC