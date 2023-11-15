# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 726J

Tested At: 15 Nov 23 18:00 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -374 day(s)\
Subject: CN=SHAKEN 726J, OU=SHAKEN, O=Votacall Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88dbb645-62bf-4d40-960c-4e8e3223880f/6f4c58e2b5fa4cc73657b70dc9cd0f4d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQa7%2BN5nMpkn3C%2BFBWOc2mkjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjkyMDIzMTJaFw0yMjExMDUyMDIzMTFaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxWb3RhY2FsbCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASm2A0wN9xhf5tPvEpAAWJJi72uyWPAb%2FBPtkco%2BCSlnRrVEhJbuKlEDPvLBAfYtYYcZ%2BbXvR%2BGG0Ys27fz79Guo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFMsEKKv%2F0YVHUuK%2FEgne7lLxlWIrMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyNkowCgYIKoZIzj0EAwIDSAAwRQIhANKeEX5FSOGt4cUlaR%2B8r%2FslEpVhAGVBLP%2FYpy5b6ixoAiAPOHUz23%2BeMWtirMbvwC2BFbx5pc3V1c0hFyE4kBupgQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 18:10 UTC