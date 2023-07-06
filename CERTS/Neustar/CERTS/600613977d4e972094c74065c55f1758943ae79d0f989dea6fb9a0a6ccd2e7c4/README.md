# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ***SHAKEN***464D

Tested At: 06 Jul 23 14:07 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 375 day(s)\
Subject: CN=***SHAKEN***464D, OU=VOIP, O=Fibernetics, L=Cambridge, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://stir.fibernetics.ca/prod-cert2023.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDdzCCAx2gAwIBAgIUVWBSFKpe%2Bmd9EYny5F61KFq6UBgwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yMzA2MTIxNTU5MjVaFw0yNDA3MTQxNTU5MjVaMG4xCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjESMBAGA1UEBwwJQ2FtYnJpZGdlMRQwEgYDVQQKDAtGaWJlcm5ldGljczENMAsGA1UECwwEVk9JUDEZMBcGA1UEAwwQKioqU0hBS0VOKioqNDY0RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKxX69RZ1Zx4s84TQq8TcSRVDgrQLFWw5mk0ODHRmUYayoKBTaqf4b3LKh4tyeBSkUTqhgxfhlx2TVuAvTKpCBejggF1MIIBcTAWBggrBgEFBQcBGgQKMAigBhYENDY0RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFC0dN2RITw4dbEihZNhZ1abQBsJhMBYGA1UdIAQPMA0wCwYJKwYBBAGDtx8BMIHgBgNVHR8EgdgwgdUwgdKgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybKKBmqSBlzCBlDELMAkGA1UEBhMCQ0ExKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMSEwHwYDVQQLDBhzdGlwYS1jc3RnYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgU1RJLVBBIFJvb3QgQ0EwHQYDVR0OBBYEFHSlOA6kvnkRC4SVM0f6ah8rwfqMMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA0TXbTxmsu998DyBmNGE2rBfpVExEHq8vPFcGd%2BWfDpICICo1%2BAIKsnkr5KcAsk5cFm2sovMbuGwbXDU3NCIC3ecV)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 464D' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC