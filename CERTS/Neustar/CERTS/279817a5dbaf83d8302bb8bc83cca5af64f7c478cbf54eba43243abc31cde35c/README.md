# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 468B

Tested At: 01 Mar 23 18:14 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 51 day(s)\
Subject: CN=SHAKEN 468B, O=Panhandle Telecommunications Systems Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/117.115

[View certificate details](https://understandingwebpki.com/?cert=MIIDGDCCAr%2BgAwIBAgIUSn1a%2F0ZSVtKRaNMz%2FrtmcAreaSIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDQyMTE1MzUzOFoXDTIzMDQyMTE1MzUzOFowVjELMAkGA1UEBhMCVVMxMTAvBgNVBAoMKFBhbmhhbmRsZSBUZWxlY29tbXVuaWNhdGlvbnMgU3lzdGVtcyBJbmMxFDASBgNVBAMMC1NIQUtFTiA0NjhCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0kdRVaKoCZRzn2iCEJGST3BPLIXstw36jCMM4sB9vtKEjGy%2FHmdv43ywRSaWilA1NKn%2B5bTA7SqXZbT3PafiPKOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ0NjhCMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUtDGzwQTDvgt9X%2B2e8F5NHwYOzDEwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIFWiTEvICL2m2UnFHuiEunB7XCPwCYGOsq99RGBz5IyxAiA1gTi%2FKHBJF6E%2BZMiDa6Q85p1wKlsJ8uoILmSPMmmkpw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01 Mar 23 18:22 UTC