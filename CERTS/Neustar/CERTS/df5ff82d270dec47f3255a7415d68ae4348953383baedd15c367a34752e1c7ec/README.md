# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 899J

Tested At: 21 Aug 23 20:16 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -215 day(s)\
Subject: CN=SHAKEN 899J, O=Telco Connection, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://ssc.getsipnav.com/certs/a80448b68598a8d87c75bf6df015f9f2a30ae551

[View certificate details](https://understandingwebpki.com/?cert=MIIDATCCAqegAwIBAgIUYOcNsD2%2FNwTs2ZRtDNExASx%2BE1QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDExODE4MDYwOVoXDTIzMDExODE4MDYwOVowPjELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFRlbGNvIENvbm5lY3Rpb24xFDASBgNVBAMMC1NIQUtFTiA4OTlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbXfkFZOcyJ0CVEMJ7TIFoaKe6fW9osdZjAcPAOzpKZBMstpCNPg3x1%2BLGkqKBkVCaF%2FckO%2BsyUrR5m2KhXyVZKOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ4OTlKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUbshvnZX%2FTs4l10lHVUBqFoSyYTwwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB%2BGzgamdCdACLhABuW6WyeAF%2F3p%2Ff3hWGCwMaT%2F255MAiEAoSfNP3yhJk7mtY6nBsnHoIJZCXPbAtLLx61vkGhkcWo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC