# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 26 Aug 24 18:00 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -52 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/53a33cea-2134-494a-a50d-9668cd3bf516.pem

[View certificate details](https://x509.io/?cert=MIICqzCCAlGgAwIBAgIIYJPLKVeAR08wCgYIKoZIzj0EAwIwdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjQwNjA0MjA1NTU2WhcNMjQwNzA0MjA1NTU2WjBjMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxFDASBgNVBAMTC1NIQUtFTiAzMThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAERa6iwDc42AMGhZn%2BxQ8xJwG7tHMCFVB8g2%2BV6AVju1ZnO67NRHOehJLP5YIWvh%2BMLQxJl%2BosimZj905kkjrAIqOB2zCB2DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQU%2F1LwV3eaQ%2FB3OD56Fdhx2%2BD8ty0wHwYDVR0jBBgwFoAU64mTb7rIbYlz7nBF2qD7W9Q5IiEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNIADBFAiEA%2FXqjNAcDtnGtNxqYoBY%2Bk58spUHiwL03gCQE7kxVyakCIA471kXUdftw1fCT5oMNY%2FE5VvH9oDOBzyFy2SH%2FnRoN)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 26 Aug 24 18:03 UTC