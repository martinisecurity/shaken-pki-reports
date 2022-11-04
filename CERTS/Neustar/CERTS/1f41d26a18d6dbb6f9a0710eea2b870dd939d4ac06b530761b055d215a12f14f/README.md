# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 393J

Tested At: 04 Nov 22 01:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 133 day(s)\
Subject: CN=SHAKEN 393J, O=Cloudli Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11508.10138

[View certificate details](https://understandingwebpki.com/?cert=MIIDBzCCAq2gAwIBAgIUSCmus0zEL1M9RWrZTHfOKXuwhr0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDMxNjEzMTAzN1oXDTIzMDMxNjEzMTAzN1owRDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkNsb3VkbGkgQ29tbXVuaWNhdGlvbnMxFDASBgNVBAMMC1NIQUtFTiAzOTNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBJxxPSfDFYuVqovhtWH1AEKY32REj7TZgaSJEwbKjNxTixuwjgf46dqkKx1ya37gulIEyPCMVndNDoz4wESheqOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQzOTNKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUPgvN01kZo3qN2AzxzDB0vn1H%2B2gwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIDzb9LLobIB9WcrGKyAi0k5UX1MTrxyR2sYNwz0gw1cfAiEAml2kI1L89ZzXLz3O8B4QFGFMmWlcO6voQ1ZbtU%2Fki0M%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 04 Nov 22 01:11 UTC