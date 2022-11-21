# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672J

Tested At: 21 Nov 22 23:24 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 185 day(s)\
Subject: CN=SHAKEN 672J, O=Zoom Voice Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/134.135

[View certificate details](https://understandingwebpki.com/?cert=MIIDCjCCArCgAwIBAgIUL1tyIs9%2BCrGb5VwRLzI%2Bh8AmE%2FYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyNTE2NDE1NloXDTIzMDUyNTE2NDE1NlowRzELMAkGA1UEBhMCVVMxIjAgBgNVBAoMGVpvb20gVm9pY2UgQ29tbXVuaWNhdGlvbnMxFDASBgNVBAMMC1NIQUtFTiA2NzJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfhT6jxmxskytcnz3M83NQUiZHAcY26%2FSANFGZsqNOr3dKMgexBdV6Lx3wX4Mnbls6FP%2FmJDsemFscsonWnfUQqOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUbLIPzm%2Fav9cOp89xzZxYFAfYJIEwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCwP8ex2dl9QXjDXKPikIL%2FfZgFRtMODWB9EORjllWsKQIgZ%2F4cvJ6LZ5Ot1J6IY8EqsWpn0U6r4bk5BHh44QqcLwo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Nov 22 23:27 UTC