# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 755J

Tested At: 26 Aug 24 18:11 UTC\
Initial Validity Period: 367 day(s)\
Remaining Validity Period: -65 day(s)\
Subject: serialNumber=ba7cdaca-6d97-4141-9ea5-0b358f4cb3f2, CN=SHAKEN 755J, O=CMSInter.net LLC, C=US\
Issuer: CN=Telonium STI-CA Intermediate CA, O=Telonium STI-CA\
Link: https://certs.telonium.net/23/d477d20f.crt

[View certificate details](https://x509.io/?cert=MIIC6jCCApCgAwIBAgIQUqxJ2Rjz0avpKS93bkM8YjAKBggqhkjOPQQDAjBEMRgwFgYDVQQKEw9UZWxvbml1bSBTVEktQ0ExKDAmBgNVBAMTH1RlbG9uaXVtIFNUSS1DQSBJbnRlcm1lZGlhdGUgQ0EwHhcNMjMwNjIyMTgzMjU2WhcNMjQwNjIyMTgzMzU2WjBtMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQQ01TSW50ZXIubmV0IExMQzEUMBIGA1UEAxMLU0hBS0VOIDc1NUoxLTArBgNVBAUTJGJhN2NkYWNhLTZkOTctNDE0MS05ZWE1LTBiMzU4ZjRjYjNmMjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJuMRwDqbk2ntcKfTJL3YETdgpfOIGNmTbrIcZM4HR63rE%2FwObc8iEMZkMsEKZHVKTxXVFSX5PENk7Mg1O5UytijggE5MIIBNTAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUwbfmLm6JrYx3CWq0r11Ct%2F2SKuUwHwYDVR0jBBgwFoAUqiS7%2FxR1QHkth2%2FoDUF3yrvNiLAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAUBggrBgEFBQcBGgQIMAYTBDc1NUowXQYMKwYBBAGCpGTGKEABBE0wSwIBAQQZVGVsb25pdW0gU1RJLUNBIFNQQyBUb2tlbgQrWFZDVmJzWDZpa3dKbGl2QnZ5WWhleEphYjIzYlRxZnIyYXRsZDNiVS12RTAKBggqhkjOPQQDAgNIADBFAiEAi3Bk9KtR50Hdq4zE8vD2940Tdkj0Rvq3wbGnURbk19sCIDf977I8QZrj9RAalcKOhb%2Frv%2Ft8Sl%2B9jR3Dm2bBSPfw)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, TNAuthorizationList shall have only one TN Entry |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | TNAuthorizationList shall have only one TN Entry |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_not_specified](../../ISSUES/e_atis_ext_not_specified/README.md) | error | ATIS1000080 | Certificate contains extensions that are not specified: 1.3.6.1.4.1.37476.9000.64.1 |


Generated: 26 Aug 24 18:49 UTC