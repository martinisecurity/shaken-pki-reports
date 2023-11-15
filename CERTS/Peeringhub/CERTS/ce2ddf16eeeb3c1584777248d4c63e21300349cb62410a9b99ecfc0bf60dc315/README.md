# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Losh Communications, Inc SHAKEN 149K 2023-11-15_000001

Tested At: 15 Nov 23 18:04 UTC\
Initial Validity Period: 284 day(s)\
Remaining Validity Period: 284 day(s)\
Subject: CN=Losh Communications\\, Inc SHAKEN 149K 2023-11-15_000001, O=Losh Communications\\, Inc, L=Springfield, ST=MO, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://nyc01.trunks2.calldecibel.com:5000/stirshaken_certs/149K.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDTDCCAvOgAwIBAgIQO98M7w4MEyKuqxiZTfWAnTAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMzExMTUwMDAwMTZaFw0yNDA4MjQxOTU2NTlaMIGUMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTU8xFDASBgNVBAcMC1NwcmluZ2ZpZWxkMSEwHwYDVQQKDBhMb3NoIENvbW11bmljYXRpb25zLCBJbmMxPzA9BgNVBAMMNkxvc2ggQ29tbXVuaWNhdGlvbnMsIEluYyBTSEFLRU4gMTQ5SyAyMDIzLTExLTE1XzAwMDAwMTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAU2lGwxw3b%2BUTKcMZ%2FLkqc%2BWEJh3kJHvRC6VGXqBi4UMEefb3BCCwQ318XmcH1YS1%2FUjhIVTtL7KJZarlhXXuujggE8MIIBODAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUTlOYyjEBkpqehaA9cOuwfbWxEEQwHwYDVR0jBBgwFoAUrqFzUYgpVxHKDKn0sQpuTrhLTQcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMBYGCCsGAQUFBwEaBAowCKAGFgQxNDlLMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAKBggqhkjOPQQDAgNHADBEAiB9ord1le2i3m6ye2o6xRVAX3V5untRcOw9VniMaladWgIgaJn7du2%2BgWgZcyqiT%2BtfZPS%2FW0ecbKlcLuieyX1STpg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 18:10 UTC