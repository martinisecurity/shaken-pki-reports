# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 224C

Tested At: 15 Nov 23 16:49 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 299 day(s)\
Subject: CN=SHAKEN 224C, OU=Voice Services, O=Onvoy LLC, L=Chicago, ST=IL, C=US\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://sticr-cstga.ccid.neustar/api/v1/certificate/925f707cc3de6c01ae333d764bdc21ae.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDeDCCAx6gAwIBAgIUMbPKRMGq%2Bd7NqxoE8ItTcQttYJQwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yMzA4MDgxNjQ1MDZaFw0yNDA5MDkxNjQ1MDZaMG8xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJTDEQMA4GA1UEBwwHQ2hpY2FnbzESMBAGA1UECgwJT252b3kgTExDMRcwFQYDVQQLDA5Wb2ljZSBTZXJ2aWNlczEUMBIGA1UEAwwLU0hBS0VOIDIyNEMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASGsUbSNrHcHurSaLzW8B8nILXWoL9n8zaap7wYDHgmhwaRr1aRqRLXcup5nhS5GuaOktxePQ0iS4d9cSNFkYtto4IBdTCCAXEwFgYIKwYBBQUHARoECjAIoAYWBDIyNEMwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBQtHTdkSE8OHWxIoWTYWdWm0AbCYTAWBgNVHSAEDzANMAsGCSsGAQQBg7cfATCB4AYDVR0fBIHYMIHVMIHSoDOgMYYvaHR0cHM6Ly9zdGlwYS1jcmwtY3N0Z2EuY2NpZC5uZXVzdGFyL2FwaS92MS9jcmyigZqkgZcwgZQxCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEhMB8GA1UECwwYc3RpcGEtY3N0Z2EuY2NpZC5uZXVzdGFyMTcwNQYDVQQDDC5OZXVzdGFyIENhbmFkYSBDZXJ0aWZpZWQgQ2FsbGVyIFNUSS1QQSBSb290IENBMB0GA1UdDgQWBBQepW0Z2rBjsuOSJker8yJYsMjq8jAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhALuGCnJpelDS865n5fHszVge%2BFNmCIaFkcYGcyJR%2BlrdAiBEr%2Blr81fMC6Pe%2FS1TxWNwaQbt07ZZv8J7SADU8Y2cPw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 16:51 UTC