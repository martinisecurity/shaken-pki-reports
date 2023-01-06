# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 841J

Tested At: 06 Jan 23 02:53 UTC\
Initial Validity Period: 14 day(s)\
Remaining Validity Period: -97 day(s)\
Subject: CN=SHAKEN 841J, OU=SHAKEN, O=Securus Technologies LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/841J/4dd0cc8d-8dd9-49b0-965c-82179fcfaeaa.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIQZAhCUfDS9d3ZkP6%2BYth9VzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTYxODMzNTNaFw0yMjA5MzAxODMzNTJaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhTZWN1cnVzIFRlY2hub2xvZ2llcyBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg0MUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARrvCMp%2Blr9sxl%2BYWrTUJfbIuiPN7DSxHKOz%2BA4R3XtjzrW5UcnLjPOX8kDPRM%2BLFi8XNXSpOUA3RFaOM9AIp0Qo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFNendn3L63mU8GLPydVxdhjyM%2FEoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0MUowCgYIKoZIzj0EAwIDSQAwRgIhAP7D%2BGChbU8nCskzdn8CDdFk07qW%2FrJpeeWuZo2tVUprAiEAsQo5rgPKaNINGzkeF53ALDi0gEuNmgloZi%2B6b3%2F7ym8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 06 Jan 23 03:03 UTC