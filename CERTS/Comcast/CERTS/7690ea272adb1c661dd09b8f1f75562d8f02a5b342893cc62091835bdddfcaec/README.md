# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 12 Feb 24 19:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/068a4bb3-d1b0-4b26-b14b-b6b4c99511a7.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIClDCCAjqgAwIBAgIIC01bbRoURrQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTI0MDExNTEwMTc0OVoXDTI0MDIxNDEwMTc0OVowYzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MRQwEgYDVQQDEwtTSEFLRU4gMzE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPUXNiRWGHK%2FBRQYxQL%2BjQBPWsgIJeHqSr7aVl7mpZADaYY6ommGJG%2BMQSUP2L1nn3PNKAj5xtswvLq%2FEcr8VaajgdswgdgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFGiOsRvUfcJiNRN7arIg%2Bjvz8tKMMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwFgYIKwYBBQUHARoECjAIoAYWBDMxOEowCgYIKoZIzj0EAwIDSAAwRQIhALi8g%2Bk9MnWYpy%2BJziEvNsmuziklHX5MrAbx%2BciojgbdAiBn6ZJOvXfg2iiDWxd1Dz%2BrNA7Rz85uEs8DM%2B5noQB0qw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 60 |


Generated: 12 Feb 24 19:26 UTC