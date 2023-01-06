# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Sansay Intermediate CA US WEST 1

Tested At: 06 Jan 23 03:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -84 day(s)\
Subject: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Root CA US, OU=Sansay CA, O=Sansay Corporation, L=San Diego, ST=California, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIC2jCCAoCgAwIBAgIUFLVfOAX18HsTtfiw3u0g8lFwPpowCgYIKoZIzj0EAwIwgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVMwHhcNMjExMDEzMjA0NjEwWhcNMjIxMDEzMjA0NjEwWjCBhTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMTAwLgYDVQQDDCdTSEFLRU4gU2Fuc2F5IEludGVybWVkaWF0ZSBDQSBVUyBXRVNUIDEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARuhlHmOYoUiuAbIvxDHYUdDmCzFO4NLi4r47NjEzoDYDCdCKjWnCWHepTAG9PxCjNPT3GBwC2wsmLzFVyn8sj6o4HGMIHDMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwHwYDVR0jBBgwFoAUCq7%2FlvCbQaO9332%2FbdpFqOgEG7kwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgKEMAoGCCqGSM49BAMCA0gAMEUCIQCCqgDD%2F5l8ehq4z2sqKGA%2Bi%2BEa%2FojVZcew55gYtCs88wIgScQGeC8085y6yHM6Zvk0ce0OF%2B2Ee0wofPdQB7qg%2BUk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 06 Jan 23 03:03 UTC