# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 318J

Tested At: 12 Feb 24 19:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 8 day(s)\
Subject: CN=SHAKEN 318J, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US\
Link: https://sticr.stir.comcast.com/3bb20b6b-a37a-4350-8847-12856a5b9008.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIClDCCAjqgAwIBAgIIPMv2eRZelEQwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTI0MDEyMTEwMTc1M1oXDTI0MDIyMDEwMTc1M1owYzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MRQwEgYDVQQDEwtTSEFLRU4gMzE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGRpjIZH7AAPvHUiHE75CZmaRUiPfGQHKX3KNzFbGSDBfDl5T%2BG0TLojki%2BdD89WkRJynOvo%2FQ7Y1XKx%2FzmXFySjgdswgdgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFAynrvRX96u1RvegTHmn5TtWd%2B35MB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwFgYIKwYBBQUHARoECjAIoAYWBDMxOEowCgYIKoZIzj0EAwIDSAAwRQIgI6SK0SpQa%2F%2BKiR6M2W1uW%2BFj3wwPxaznsbTViQKO08YCIQDnAecS94cUnWaNF5G75tgAmLEQ%2Bd87YzoDS2wKMHTqpw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 62 |


Generated: 12 Feb 24 19:26 UTC