# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 3130

Tested At: 01 Nov 22 16:29 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 204 day(s)\
Subject: CN=SHAKEN 3130, O=Bullseye Telecom, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11416.10156.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDAjCCAqegAwIBAgIUWS8aagggPgva5OUnmehys%2FYvj3EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyMzE2NDMzNVoXDTIzMDUyMzE2NDMzNVowPjELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEEJ1bGxzZXllIFRlbGVjb20xFDASBgNVBAMMC1NIQUtFTiAzMTMwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEL%2B0TTKJuhFC%2Fr7Fzn0wKq5byD2pPmf%2BTyzVAKf3VxZjzQLPmhQJfQutH1uzcOkRF%2BVW%2FKlXpHy1pocYqapFPd6OCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQzMTMwMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQU2Khr2w3sKMlBPHNikG3niX3JGCgwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDhF7nyiDWlnJJwCRofSHJB08u%2Fuk0FmCaI%2BXHkI2WZ%2FAIhAIMQwE4XXGvCut6URk8JGf5ChNp2t2m%2B46jj155VJmKF)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 16:30:07