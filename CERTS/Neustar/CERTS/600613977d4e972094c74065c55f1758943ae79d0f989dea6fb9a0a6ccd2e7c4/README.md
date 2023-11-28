# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ***SHAKEN***464D

Tested At: 28 Nov 23 20:17 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 229 day(s)\
Subject: CN=***SHAKEN***464D, OU=VOIP, O=Fibernetics, L=Cambridge, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://stir.fibernetics.ca/prod-cert2023.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDdzCCAx2gAwIBAgIUVWBSFKpe%2Bmd9EYny5F61KFq6UBgwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yMzA2MTIxNTU5MjVaFw0yNDA3MTQxNTU5MjVaMG4xCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjESMBAGA1UEBwwJQ2FtYnJpZGdlMRQwEgYDVQQKDAtGaWJlcm5ldGljczENMAsGA1UECwwEVk9JUDEZMBcGA1UEAwwQKioqU0hBS0VOKioqNDY0RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKxX69RZ1Zx4s84TQq8TcSRVDgrQLFWw5mk0ODHRmUYayoKBTaqf4b3LKh4tyeBSkUTqhgxfhlx2TVuAvTKpCBejggF1MIIBcTAWBggrBgEFBQcBGgQKMAigBhYENDY0RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFC0dN2RITw4dbEihZNhZ1abQBsJhMBYGA1UdIAQPMA0wCwYJKwYBBAGDtx8BMIHgBgNVHR8EgdgwgdUwgdKgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybKKBmqSBlzCBlDELMAkGA1UEBhMCQ0ExKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMSEwHwYDVQQLDBhzdGlwYS1jc3RnYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgU1RJLVBBIFJvb3QgQ0EwHQYDVR0OBBYEFHSlOA6kvnkRC4SVM0f6ah8rwfqMMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA0TXbTxmsu998DyBmNGE2rBfpVExEHq8vPFcGd%2BWfDpICICo1%2BAIKsnkr5KcAsk5cFm2sovMbuGwbXDU3NCIC3ecV)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 464D', but common name is '***SHAKEN***464D' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.56223.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 28 Nov 23 20:21 UTC