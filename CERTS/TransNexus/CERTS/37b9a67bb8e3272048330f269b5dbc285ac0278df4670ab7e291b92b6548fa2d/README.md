# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 6628

Tested At: 11 Jan 23 20:38 UTC\
Initial Validity Period: 124 day(s)\
Remaining Validity Period: -8 day(s)\
Subject: CN=SHAKEN 6628, OU=SHAKEN, O=Merryville Investments LTD Inc dba ClarityTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/6628/e0f14ccb-0cb4-4b46-a7b7-988bdfd30cce.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDDjCCArWgAwIBAgIQenZZqZQaD99tzZyOQ2QERjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MDEwMzI1NTFaFw0yMzAxMDMwMzI1NTBaMGwxCzAJBgNVBAYTAlVTMTYwNAYDVQQKEy1NZXJyeXZpbGxlIEludmVzdG1lbnRzIExURCBJbmMgZGJhIENsYXJpdHlUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDY2MjgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAST11ccDfIZk1Qh%2FZM9L%2BauNjQ5ksJx2DlyWtVvyQp5SnSMagtT5qg91sNAPHE6kpTteTGYYvyzy5FmpN15ArdGo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFM6ZNiv8yMgbT3jI3G2%2B5kOLgUUEMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDY2MjgwCgYIKoZIzj0EAwIDRwAwRAIgcgXxxaAn9Og2z20Xqypba7XsDH7byJ%2BLAtA43OUUDpsCIDYpArDjpeiUtSHjtTEi9JE7vGttM1fPjjdMy2Dmn1eM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:04 UTC