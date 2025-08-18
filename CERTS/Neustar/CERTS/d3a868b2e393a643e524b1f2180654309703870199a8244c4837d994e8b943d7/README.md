# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ***SHAKEN***464D

Tested At: 18 Aug 25 21:09 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 385 day(s)\
Subject: CN=***SHAKEN***464D, OU=VOIP, O=Fibernetics, L=Cambridge, ST=ON, C=CA\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-2, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://stir.fibernetics.ca/prod-cert2025.pem

[View certificate details](https://x509.io/?cert=MIIDdjCCAx2gAwIBAgIUTSQnArnpxFJ2pMvIpCcwDouV2cYwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMjAeFw0yNTA4MDUxNDIxNTlaFw0yNjA5MDcxNDIxNTlaMG4xCzAJBgNVBAYTAkNBMQswCQYDVQQIDAJPTjESMBAGA1UEBwwJQ2FtYnJpZGdlMRQwEgYDVQQKDAtGaWJlcm5ldGljczENMAsGA1UECwwEVk9JUDEZMBcGA1UEAwwQKioqU0hBS0VOKioqNDY0RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEeX6KS9csFjkI3VSWSHC9J8i9jQ%2FK06D9jUYErkdYAUYjig70V3t1qW2JvcksHy9dGMKbUxwoUJU7GCvrulUbqjggF1MIIBcTAWBggrBgEFBQcBGgQKMAigBhYENDY0RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFC0dN2RITw4dbEihZNhZ1abQBsJhMBYGA1UdIAQPMA0wCwYJKwYBBAGDtx8BMIHgBgNVHR8EgdgwgdUwgdKgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybKKBmqSBlzCBlDELMAkGA1UEBhMCQ0ExKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMSEwHwYDVQQLDBhzdGlwYS1jc3RnYS5jY2lkLm5ldXN0YXIxNzA1BgNVBAMMLk5ldXN0YXIgQ2FuYWRhIENlcnRpZmllZCBDYWxsZXIgU1RJLVBBIFJvb3QgQ0EwHQYDVR0OBBYEFG25ynk%2Bd7e07bj2uPT3%2FMAbbqgQMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiA5dbPLspmHXWHSBC7MGN1x%2FRwZn3p6UFH8k51zrUJCLwIgVdKNqWIshZeJidCkrsZMaw4VNq3pKdmC74T2blUk7pc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 464D', but common name is '***SHAKEN***464D' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 1.3.6.1.4.1.56223.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC