# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 035K

Tested At: 15 Dec 22 18:11 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -24 day(s)\
Subject: CN=SHAKEN 035K, OU=SHAKEN, O=NetDiverse\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/ba9a1ccc-cb72-4d11-a5b3-b1211e6dea1b/698301eb576be26ebf75aefbdb9424bb.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApegAwIBAgIQWaIfISfI0ylrM6px84ydqDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTQxNjIzNDlaFw0yMjExMjExNjIzNDhaME4xCzAJBgNVBAYTAlVTMRgwFgYDVQQKEw9OZXREaXZlcnNlLCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDAzNUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvUaVtclFdhJRWfbjitkKzUBlnZ%2Fs01wQbBKou5RQCkJl3mJ%2BCbWgSm6VkL5vzUfupD1cZkn%2FcaPIL4h0KB6aSo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFM80INqeB5aqEvMvtlkRM3MssIBUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDAzNUswCgYIKoZIzj0EAwIDSAAwRQIhAKl7d8tQIgw9f2y4vuKR%2FVfP%2BpPK1JatRxMybj5epKKBAiB2252u4DM0gwQ6S286cyAQ4ctglgVC0bRxptgTjUGF0A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Dec 22 18:35 UTC