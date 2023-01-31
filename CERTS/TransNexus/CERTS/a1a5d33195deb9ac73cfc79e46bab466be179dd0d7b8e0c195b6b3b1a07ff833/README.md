# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 31 Jan 23 17:08 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -75 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/0d61f053-0092-48d0-871d-60932b5d47a1.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIQduj60WMEP67eN8Wu9hJ5YjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDI0MDNaFw0yMjExMTYyMDI0MDJaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTa3lTd2l0Y2gxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDczOEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQXy56mgQR07o%2BQRdhevqRYW5p3HDKult%2B4cMzNO%2FubS%2BKw%2FAJ5igA4ZGoRrreQVzjNc%2FpDDSTzQ5QlWF8CaRyoo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFDcytXWGYmP%2FemN2gseN8q8CPhQZMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDczOEowCgYIKoZIzj0EAwIDRwAwRAIgIrAZjH9fCSrHHEKcuzRU2XfdegSMFg2EBpFWVvi%2BDJYCIEiqpBnF8QBcmBnv3zhAplgBHTp%2BCS8udCFh5y0QZGrQ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31 Jan 23 17:51 UTC