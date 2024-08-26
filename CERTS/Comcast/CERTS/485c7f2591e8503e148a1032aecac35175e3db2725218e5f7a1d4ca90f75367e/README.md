# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 26 Aug 24 18:00 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -11 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/4244a23c-d07e-4894-9447-4930791486ca.pem

[View certificate details](https://x509.io/?cert=MIICrDCCAlGgAwIBAgIINpO%2B87LCiYwwCgYIKoZIzj0EAwIwdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjQwNzE1MjA1NTUxWhcNMjQwODE0MjA1NTUxWjBjMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxFDASBgNVBAMTC1NIQUtFTiAzMThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbTw3CTjiQSauxBy2qrg3kD2wEURh2ZpfoOThAiw%2Bvla1HhzL%2FCKzqfseBS7JR8bTbJqs%2BW0xzdvVYiYivU1266OB2zCB2DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUdPNyTHBhn8E0YLaS7SaoJjvCtUYwHwYDVR0jBBgwFoAU64mTb7rIbYlz7nBF2qD7W9Q5IiEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNJADBGAiEA1TVqaQYNqG0JRxdbBWooeK%2BeETwPe2a5cnLKv5PAWGkCIQCMcwXuhnW3OiDYJ75%2FoFkH8J9YBOtzlfl721JLAmRCBw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 62 |


Generated: 26 Aug 24 18:03 UTC