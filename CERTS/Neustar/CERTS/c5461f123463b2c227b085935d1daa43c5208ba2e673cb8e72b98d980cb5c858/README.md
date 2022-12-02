# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 669B

Tested At: 02 Dec 22 07:45 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 184 day(s)\
Subject: CN=SHAKEN 669B, O=BTC Broadband, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11408.10160.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FzCCAqSgAwIBAgIUKzoArI3vbO4anlh4G4TUs78kLwMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYwMzIwMDMwOVoXDTIzMDYwMzIwMDMwOVowOzELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDUJUQyBCcm9hZGJhbmQxFDASBgNVBAMMC1NIQUtFTiA2NjlCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9kR%2FjSMOBRZdABYdmHe3fqDlH9YhxjBIoKohRaJv8%2FvQkQoOHRTYgek%2BXgG89Gmhy3mSKBgOm0YejDRBIYBreqOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ2NjlCMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUzOUGCfe7iuv5b%2BjzRUKj5hfRdf4wDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDtBEasSiNJtFMm5JA8fadPjvawy5pRvkWK6kizsgTCMgIhAMa7SJpu%2F7%2F4ZvS9LDNO4%2FSvp7essxChZkJmdl8RBG1r)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Dec 22 07:46 UTC