# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Canada Certified Caller ID SHAKEN CA-2

Tested At: 07 Jan 23 19:18 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3438 day(s)\
Subject: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN Root CA, OU=cms-ca.ccid.neustar, O=Neustar Information Services Inc, C=CA

[View certificate details](https://understandingwebpki.com/?cert=MIIDfzCCAyWgAwIBAgIUPNQVV%2FkBmEngiLlXN100hXPzjXMwCgYIKoZIzj0EAwIwgZIxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTY21zLWNhLmNjaWQubmV1c3RhcjE6MDgGA1UEAwwxTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gUm9vdCBDQTAeFw0yMjA2MDYxNzA1MzRaFw0zMjA2MDYxNzA1MzRaMIGPMQswCQYDVQQGEwJDQTEpMCcGA1UECgwgTmV1c3RhciBJbmZvcm1hdGlvbiBTZXJ2aWNlcyBJbmMxHDAaBgNVBAsME3d3dy5jYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIENBLTIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQMY69o1XBFgwhQ3ja9xPBcy%2BKe2tbsyXA8Imsd7SYH%2FFXv2OFlmSJeF%2BjQeAYOaLAXJyLfcT1U%2BT0UXupmWWpFo4IBWDCCAVQwDwYDVR0TAQH%2FBAUwAwEB%2FzAfBgNVHSMEGDAWgBRIA3%2FvHO2MLj%2BLxARyFhAD%2BtrZLjAWBgNVHSAEDzANMAsGCSsGAQQBg7cfATCB2AYDVR0fBIHQMIHNMIHKoDOgMYYvaHR0cHM6Ly9zdGlwYS1jcmwtY3N0Z2EuY2NpZC5uZXVzdGFyL2FwaS92MS9jcmyigZKkgY8wgYwxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEhMB8GA1UECwwYc3RpcGEtY3N0Z2EuY2NpZC5uZXVzdGFyMS8wLQYDVQQDDCZOZXVzdGFyIENhbmFkYSBDZXJ0aWZpZWQgQ2FsbGVyIFNUSS1QQTAdBgNVHQ4EFgQULR03ZEhPDh1sSKFk2FnVptAGwmEwDgYDVR0PAQH%2FBAQDAgGGMAoGCCqGSM49BAMCA0gAMEUCIHcfpEsabVMgcWcdOwGrRTFgMGKaMNHxzTxHyqZz9A59AiEA7k5OsjycdpWHjzhKn8WVyVXp%2BaEFlz8nvcC36DwTbVY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ca_certificate_policies](../../ISSUES/e_atis_ca_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |


Generated: 07 Jan 23 19:18 UTC