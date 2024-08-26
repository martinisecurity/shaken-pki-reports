# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -24 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/f9c54bf4-5ef4-4992-aa9c-15ded4a0062a.pem

[View certificate details](https://x509.io/?cert=MIICrDCCAlGgAwIBAgIIGGgnL%2FJSUygwCgYIKoZIzj0EAwIwdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjQwNzAyMjA1NTUyWhcNMjQwODAxMjA1NTUyWjBjMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxFDASBgNVBAMTC1NIQUtFTiAzMThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEt3J9dxD8Cy270RtRxTZLRwjjr0yqK9qBeBqr82gbR80LNhMbhKNFIJsUxKR9uJNJ9S55iUKe4Pq8G4kz08C3IaOB2zCB2DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUNItg7idfprvVpW1m2E0XVeIHVC4wHwYDVR0jBBgwFoAU64mTb7rIbYlz7nBF2qD7W9Q5IiEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNJADBGAiEA%2B4IylKZuKpXCr9mDLQxN4spShqblwPO1O1jXZpVK8KYCIQCe9MWIpPkKHLPorsZ3bCjcmQcvVuAWYTTUAecfd8F2ew%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 61 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 26 Aug 24 18:03 UTC