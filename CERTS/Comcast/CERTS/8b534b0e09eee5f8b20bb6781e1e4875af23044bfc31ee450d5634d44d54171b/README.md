# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 12 Feb 24 19:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/22def39e-474a-4934-afb3-8e79162592b1.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICkzCCAjqgAwIBAgIIaNkeLqeJ3v8wCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTI0MDExOTEwMTc0OVoXDTI0MDIxODEwMTc0OVowYzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MRQwEgYDVQQDEwtTSEFLRU4gMzE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE2xjgIjsu1unLDfMu1%2BSIeEq0LyCoGKjS5Z8te6MKH%2FmgEY8AmPN1CU9AnVRJEtCkaCmu%2FmDe05CgAlOt7GFc6jgdswgdgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFAhffVAKDq0HcWxczXJA8Ro6MS4OMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwFgYIKwYBBQUHARoECjAIoAYWBDMxOEowCgYIKoZIzj0EAwIDRwAwRAIgRTBMbX4f%2B2ZvKycJMVULXVI1OFIeIUMNqnzf3iZMN9cCIFUx3lzSLxyp5vvizHoYOiUKGCTgNzdb%2Fl4CgdvRWlsf)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |


Generated: 12 Feb 24 19:26 UTC