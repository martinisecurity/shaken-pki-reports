# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 534J

Tested At: 04 Nov 22 01:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 217 day(s)\
Subject: CN=SHAKEN 534J, O=UComTel, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/145.146

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIUUZwuIZc1m%2BWiNuhqHDWbz7CzbjwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYwODE1MTk1NFoXDTIzMDYwODE1MTk1NFowNTELMAkGA1UEBhMCVVMxEDAOBgNVBAoMB1VDb21UZWwxFDASBgNVBAMMC1NIQUtFTiA1MzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEuoNiyf7BuBDpCrImu84vvVPmFoLiPwMya%2Bi9YtiDh04Z4pO7apDhmn9wxAv9gcqFt2qsdnwiK%2Be6Cl3VjWblyaOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ1MzRKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUFWo3eW%2Bo3r9UkfxxOrNzKQs53F4wDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIHmLOBulmhXBrb5hJvUKZkjXM%2F2XNHa2DLtGkp7CbVCjAiEAmCEpn9D6gRp%2FT8SjyB02LXSbWpAtvfi%2BYBI51ocis9k%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 04 Nov 22 01:11 UTC