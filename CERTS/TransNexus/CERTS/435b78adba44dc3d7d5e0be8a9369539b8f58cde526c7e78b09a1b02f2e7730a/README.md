# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 30 Jan 23 23:01 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -103 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/158b2599-ad08-444d-b19c-d7fad37c803f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApGgAwIBAgIQQbqSkX8YUBuRpDSoL5c0ATAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAyMDIwMDRaFw0yMjEwMTkyMDIwMDNaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTa3lTd2l0Y2gxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDczOEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATLSDWpYXiornxv0gCDsdPdmpx4GJxuC21SoiA3Xp%2BqlTd7ldUd1hOpX4PPfe6wUcGc3kYeU4qsBqsIMS%2F4UnCTo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFGXLcjfTuCw5njIJ2BgDL7Ubsd9eMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDczOEowCgYIKoZIzj0EAwIDSQAwRgIhANGHFQydgsCU7fm1HChdiXi9Yk8vDIaH%2BejjiYmhNZESAiEA%2FM2Sb0CQGqzO6PlEZDh8rUxH4PhYXfE2L5DFU%2Bu%2FmkM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 30 Jan 23 23:10 UTC