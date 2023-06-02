# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 726J

Tested At: 02 Jun 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -205 day(s)\
Subject: CN=SHAKEN 726J, OU=SHAKEN, O=Votacall Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88dbb645-62bf-4d40-960c-4e8e3223880f/aeecdd08bbcfd4e4e6150f7704e02c62.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQVkiRZiBWHBIjsdIUVrvqdjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDIzMjNaFw0yMjExMDgyMDIzMjJaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxWb3RhY2FsbCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATkA4p8awYWEVn9W3d14SvWmbGDMl8QGuOsR13QoQk8wS3ZAlH6QuSHDXeW0riVeqe%2BCDcSUO3OVq7S7vUeaQT4o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFGMzeH5IOBAmlSNT23LQuupDjLAjMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyNkowCgYIKoZIzj0EAwIDSAAwRQIgeDNdnrvzGiZkDNtjFoyUDJfbV8Hd2L8c6asE8swqteUCIQDYEU6SxnteefoRrebzOs7%2BTcgLBYOV2nPJ48PUHFVxHA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC