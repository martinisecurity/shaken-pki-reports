# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 750J

Tested At: 15 Nov 23 17:58 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -299 day(s)\
Subject: CN=SHAKEN 750J, O=Microtalk USA\\, Inc., C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBDCCAqqgAwIBAgIUGImhFqIf5Hygm343foSVjnif%2BoYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDEyMDE3MzQzNVoXDTIzMDEyMDE3MzQzNVowQTELMAkGA1UEBhMCVVMxHDAaBgNVBAoME01pY3JvdGFsayBVU0EsIEluYy4xFDASBgNVBAMMC1NIQUtFTiA3NTBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKGsCorXZmvMvzLMJsIaqu1iQuh175yJHI9djkT%2F39bZmyhvvsQglKtuJfLfWK7S33JhPJYsvCYWN3lVtKYxMtKOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ3NTBKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUWKpDvoG%2FDdw3hsvQndoYjfmBqYQwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDlNUrWIoSMmehRlUz%2B1IokU0rE8msdcafKw0bYo1A2TgIgddh8JJPwokEktzLIEolmdQVg%2BMG%2FMz7%2FLs50eDGienA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 18:10 UTC