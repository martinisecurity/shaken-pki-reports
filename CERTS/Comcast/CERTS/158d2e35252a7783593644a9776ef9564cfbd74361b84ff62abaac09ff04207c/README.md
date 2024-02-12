# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 12 Feb 24 19:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/7204954e-1eaf-4240-92ea-9540d453dc10.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIClDCCAjqgAwIBAgIINBTAfZiDSB8wCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTI0MDEyMDEwMTc0OFoXDTI0MDIxOTEwMTc0OFowYzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MRQwEgYDVQQDEwtTSEFLRU4gMzE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABNNof5cm90mYwvxp%2FdLghlePyBJI36Che%2BlH%2FGgzTiV13BZ1RUd2mzfMqCaY2H61RWMTLUf%2BuKxWVdCW7PPJIeujgdswgdgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFIUAKfAjLQO4dKVQhyhzD9QOO7zuMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwFgYIKwYBBQUHARoECjAIoAYWBDMxOEowCgYIKoZIzj0EAwIDSAAwRQIgF2TAkkccUQ8BgFckzFKsKO33CaiicBEt44jjsYmMS2ECIQCs7alHJIPwWVlC0LQZEQ4H%2FNoG8WW9ydJoniKdojdLig%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 62 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 12 Feb 24 19:26 UTC