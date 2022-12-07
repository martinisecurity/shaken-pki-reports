# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Sansay Intermediate CA Canada 1

Tested At: 07 Dec 22 18:54 UTC\
Initial Validity Period: 3285 day(s)\
Remaining Validity Period: 3034 day(s)\
Subject: CN=SHAKEN Sansay Intermediate CA Canada 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Root CA Canada, OU=Sansay CA, O=Sansay Corporation, L=San Diego, ST=California, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIDZjCCAw2gAwIBAgICEAEwCgYIKoZIzj0EAwIwgY4xCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSUwIwYDVQQDDBxTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgQ2FuYWRhMB4XDTIyMDMzMTE3MTE1OFoXDTMxMDMyOTE3MTE1OFowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEvMC0GA1UEAwwmU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgQ2FuYWRhIDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATkIyPmC%2BLnzHYTq5dExBvTqovPfLf9Q6p1NiR3ghjmoll8vb3k3x30u%2F84LB8yvJZhNz9Ssrugf11XGuW8Cebuo4IBYTCCAV0wFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTjLvlmdFBSOSX%2FJnz4i1fKeEWwPjCBwwYDVR0jBIG7MIG4gBRTPtxiUKCwCos7eKGaWlinLCGd0KGBlKSBkTCBjjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCVNhbiBEaWVnbzEbMBkGA1UECgwSU2Fuc2F5IENvcnBvcmF0aW9uMRIwEAYDVQQLDAlTYW5zYXkgQ0ExJTAjBgNVBAMMHFNIQUtFTiBTYW5zYXkgUm9vdCBDQSBDYW5hZGGCCQCKSSCThs2lOTAPBgNVHRMBAf8EBTADAQH%2FMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHBzOi8vc3RpcGEtY3N0Z2EuY2NpZC5uZXVzdGFyL2FwaS92MS9jcmwwDgYDVR0PAQH%2FBAQDAgKEMAoGCCqGSM49BAMCA0cAMEQCIHt3Q4ufi7NUoY%2FMmF2nhg6%2B2z1seHaxcE7g3y174HT%2FAiBXmU0bEGhS9QsRu28OHJG1B%2FcVXNKLDFSNgG5JeX9m0A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ca_certificate_policies](../../ISSUES/e_atis_ca_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_ca_serial_number](../../ISSUES/e_atis_ca_serial_number/README.md) | error | ATIS1000080 | STI certificates shall include a Serial Number field containing an integer greater than zero. The serial number shall contain at least 64 bits of output from a Cryptographically Secure PseudoRandom Number Generator (CSPRNG) |


Generated: 07 Dec 22 18:54 UTC