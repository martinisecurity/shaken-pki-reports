# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 224C

Tested At: 21 Aug 23 20:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 354 day(s)\
Subject: CN=SHAKEN 224C, OU=Inteliquent, O=Onvoy LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/1.1

[View certificate details](https://understandingwebpki.com/?cert=MIIDFDCCArmgAwIBAgIUYveh2ncZiK9HV3pIhnFTJgVHbUswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDgxMDE5MDc1MFoXDTI0MDgwOTE5MDc1MFowTTELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCU9udm95IExMQzEUMBIGA1UECwwLSW50ZWxpcXVlbnQxFDASBgNVBAMMC1NIQUtFTiAyMjRDMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWXVd5czc1874rIu0EdkBPWapxRRD2SYxTkapztr1LX1LyisIp9sMt%2FC4hBiR2RmRLQx5DmZXs%2FB%2FXiZowywiKqOCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQyMjRDMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUB72sVbKyC4Ogxs4pCCHtQ7UMD34wDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCUdvBFtFlgslH0H3%2BCZ8G4DWjb14TU0xGIkN%2B%2B9am8dQIhAJyjRijWTj7WlVUelYvUiAXyLI7YI%2BiqwIHMuDdveXvU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC