# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 12 Feb 24 19:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -4 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/183f377e-2810-4532-88a8-4402603e2133.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIClTCCAjqgAwIBAgIIfOWfZBJeNFgwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTI0MDEwOTEwMTc0OFoXDTI0MDIwODEwMTc0OFowYzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MRQwEgYDVQQDEwtTSEFLRU4gMzE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABK3xUoZoaJAJpKJfbX3LdAkWlUgZfwAOhj0FCgJqSBNlu43jrCkkxQHb62LuVU7frMKlObJIw5YoOV9utdExS5ujgdswgdgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFNsvsxgr%2BhENlgak%2BOp6VW0zA15QMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwFgYIKwYBBQUHARoECjAIoAYWBDMxOEowCgYIKoZIzj0EAwIDSQAwRgIhAKqvNeeO8uUseZogFZiq%2B7XsbpZV9yFTjmQ8KP1e4zEAAiEA7Rs35JZicAoSIx67odJvOgsNQ5MTyK8o7AdcLjtVutk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |


Generated: 12 Feb 24 19:26 UTC