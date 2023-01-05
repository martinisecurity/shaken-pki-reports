# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5493

Tested At: 05 Jan 23 18:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 216 day(s)\
Subject: CN=SHAKEN 5493, O=Netcarrier Telecom Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/95.234

[View certificate details](https://understandingwebpki.com/?cert=MIIDBjCCAq2gAwIBAgIUapYg%2FuTtmb%2FInr2VsEWoq7Fd%2BycwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgwODE4MzgyOVoXDTIzMDgwODE4MzgyOVowRDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFk5ldGNhcnJpZXIgVGVsZWNvbSBJbmMxFDASBgNVBAMMC1NIQUtFTiA1NDkzMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWV%2BuePpZIsXPzU6kOBQAIAqswtW7qUaYyhj2ml5N2dLjgyLYyFFn7aXLpwIbxFY0SPEJ6ldzTpSjggw9AlgBJaOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ1NDkzMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUaVSkUm0b2BNbUMK6%2FGR4ImMNz5AwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIAHlkFD3XufBivoinhMcNUmA471%2FTRV5ZzXHvyD8QRVGAiAc1lgUrfNBe6t%2FpHhSfoLnpmtQmmm7OAx36ZbIrJ9lPQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 05 Jan 23 18:35 UTC