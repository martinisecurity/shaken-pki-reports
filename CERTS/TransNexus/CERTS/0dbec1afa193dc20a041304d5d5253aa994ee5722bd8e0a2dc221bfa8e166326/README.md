# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 793J

Tested At: 10 Nov 22 06:40 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 3 day(s)\
Subject: CN=SHAKEN 793J, OU=SHAKEN, O=UniTel Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/45c216c2-fbb9-4f33-a3cf-6a64aad19357/ac9e8f52e8a0ec67ce3aa73bc710735f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApSgAwIBAgIQf7fs5j3lKVY3xmRGuvHApzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDUxNTQ1MTdaFw0yMjExMTIxNTQ1MTZaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxVbmlUZWwgVm9pY2UxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc5M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATsBCyQXvW3ETm7duwTIWIczVL8sgjzwpjltZP8yIjK4Wynxqo4P5DTTKw5Vb4RE0kOW94GcGV6MftrUkkdF0PRo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFPcskv6MzIumbRF2%2FxAl4sN18MP3MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc5M0owCgYIKoZIzj0EAwIDRwAwRAIgDyqrs3Wz82FnZGjoKFHgvmd0CKhgmDSi6H374leE53ECICCg4rQclvGXNuMchdSerE9o7qUOTOV4SmCFU8B5j0su)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 10 Nov 22 06:43 UTC