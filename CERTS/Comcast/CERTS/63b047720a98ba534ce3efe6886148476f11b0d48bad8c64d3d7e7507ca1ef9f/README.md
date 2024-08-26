# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 26 Aug 24 18:00 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -70 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/6a6e06ee-8c8a-444b-a54b-d158e3b29e76.pem

[View certificate details](https://x509.io/?cert=MIICqjCCAlGgAwIBAgIIQXe7k5UMiHswCgYIKoZIzj0EAwIwdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjQwNTE4MDg1OTAxWhcNMjQwNjE3MDg1OTAxWjBjMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxFDASBgNVBAMTC1NIQUtFTiAzMThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER3XQ9c8rvzWblLSFGB0xOjTQ9kG0dVfgvX1DXvO2JPZaIvULjfrRLyf7je14SN%2BMrK4Zork9MOq%2FLeRafQaSwqOB2zCB2DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUpssTt2dXxxxH%2BhU0gO1tCk1OvrQwHwYDVR0jBBgwFoAU64mTb7rIbYlz7nBF2qD7W9Q5IiEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNHADBEAiAMPikGk%2FA%2BJvrtXmzReKkLk%2Bc%2BCfsU1%2Byv4APUZWZVwwIgUW04kZrB1cMnhwJQU8CZD2krz0DHbUVkyrWxVqhUNFA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |


Generated: 26 Aug 24 18:03 UTC