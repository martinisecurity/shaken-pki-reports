# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -8 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/d3a3853c-9d20-4d5f-b587-d79e15c5675b.pem

[View certificate details](https://x509.io/?cert=MIICrDCCAlGgAwIBAgIINpjjoh2kT%2F4wCgYIKoZIzj0EAwIwdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwHhcNMjQwNzE4MjA1NTUxWhcNMjQwODE3MjA1NTUxWjBjMQswCQYDVQQGEwJVUzEVMBMGA1UECBMMUGVubnN5bHZhbmlhMRUwEwYDVQQHEwxQaGlsYWRlbHBoaWExEDAOBgNVBAoTB0NvbWNhc3QxFDASBgNVBAMTC1NIQUtFTiAzMThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAExOhyGbLBnUOvQz5k5QHDGW3r6F8JjMgDyLKzuBhbbhoY2qACS6GctO%2Bkzfz%2FK%2FbbysjAdH9kkHLh2vAoGMYU2qOB2zCB2DAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUm5sTMkZrNJ3GWqilVOTaHfp0rzgwHwYDVR0jBBgwFoAU64mTb7rIbYlz7nBF2qD7W9Q5IiEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNJADBGAiEA2mmZjJccVQtmXY1jvp4Oom%2BjbRliQ%2Fi%2F%2Fe4JulqGv%2BwCIQC6MELx1OBSmdgC94RQjbRVyyZGuKxJK1rnGu8p7RfVXA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 62 |


Generated: 26 Aug 24 18:03 UTC