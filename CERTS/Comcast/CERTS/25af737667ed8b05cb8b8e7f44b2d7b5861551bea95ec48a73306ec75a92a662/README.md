# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Comcast SHAKEN Intermediate CA

Tested At: 26 Aug 24 18:03 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 7072 day(s)\
Subject: CN=Comcast SHAKEN Intermediate CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Root CA, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US

[View certificate details](https://x509.io/?cert=MIICoDCCAkegAwIBAgIIBwibJ3x2okIwCgYIKoZIzj0EAwIwbjELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEVMBMGA1UEBwwMUGhpbGFkZWxwaGlhMRAwDgYDVQQKDAdDb21jYXN0MR8wHQYDVQQDDBZDb21jYXN0IFNIQUtFTiBSb290IENBMB4XDTI0MDExMTAyNTkwNFoXDTQ0MDEwNjAyNTkwNFowdjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MScwJQYDVQQDEx5Db21jYXN0IFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASE3Qryjt%2FXTwCLArOuczF84j3Steuk6%2B8D7i3phiakqxxLzqsFxue4MPlmFQ%2FVW%2Bibbzq9mDVg5Ee6UHUsZCrno4HGMIHDMA4GA1UdDwEB%2FwQEAwICBDAPBgNVHRMBAf8EBTADAQH%2FMB0GA1UdDgQWBBTriZNvushtiXPucEXaoPtb1DkiITAfBgNVHSMEGDAWgBSRkMqxhg5PFl6%2BtTdRP2l55SMbHDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMAoGCCqGSM49BAMCA0cAMEQCIHHnWMJziPTFEKRMOuvSIJaP6PghQs2tn1ir776b%2BrvkAiAfOfHefB%2BUIfaexjKQaCKLu9lGopAyKWiR2QgMNieWDg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size_ca](../../ISSUES/e_atis_serial_number_size_ca/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 59 |
| [e_atis_ext_crl_distribution_struct_ca](../../ISSUES/e_atis_ext_crl_distribution_struct_ca/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 26 Aug 24 18:03 UTC