# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 807J

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -124 day(s)\
Subject: CN=SHAKEN 807J, OU=SHAKEN, O=SipPhony LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/807J/99ee724e-0393-41fa-85a3-cb5dfc0ee4de.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQdfXWkjFF2oGnSGHL%2FvZf%2BzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDYyMDQyMjJaFw0yMjA5MDQyMDQyMjFaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxTaXBQaG9ueSBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgwN0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQpxp%2FJ97hBEuFcZXkR3y3CrJ0Sv2xvUqU24ce65TwH01SeoJ3xobyUC%2FC%2FjGAbWn5hj6Cz%2BJSDaPqyWo%2B59I3ho4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAQp2s%2BAaDLJaHbvzMnrTCyqg0Z1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgwN0owCgYIKoZIzj0EAwIDSAAwRQIgek05rAga8iNXquTG%2BcV5bNynguCHi%2FZhqUvVeocdoOYCIQCLsdJKkhBvkDphFBkF1d5AUKq8XQ7W%2F4YxCVWJFgpUpg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 07 Jan 23 19:18 UTC