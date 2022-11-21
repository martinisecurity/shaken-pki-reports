# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 704J

Tested At: 21 Nov 22 23:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 191 day(s)\
Subject: CN=SHAKEN 704J, O=Clearly IP, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://r.stir.tel/704J/230531154218Z/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqGgAwIBAgIUPDrZEW4Wpo7tqy7mOBspDIuEdHEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUzMTE1NDIxOFoXDTIzMDUzMTE1NDIxOFowODELMAkGA1UEBhMCVVMxEzARBgNVBAoMCkNsZWFybHkgSVAxFDASBgNVBAMMC1NIQUtFTiA3MDRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErLommlM69oWbFpGEXYVvdKRUMYsVeYmarifamf4QudKkLjwazhp2E0xaWBz7K69IzvKgXwsmQYuTOljzyhBUQ6OCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ3MDRKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUxVyd5lL0DgYzBRZYLBTtO5XIIHcwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBDfMQUqTs1uxSkIxBhXl46fralbIgKur4f%2FzW1Q1Q1yAiAfHHmMK5TEik4zDGkzyAEXFsAjoWX0DvtQv0TaEziSrA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Nov 22 23:36 UTC