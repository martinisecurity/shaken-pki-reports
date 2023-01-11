# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID SHAKEN CA-1

Tested At: 11 Jan 23 23:18 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3143 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

[View certificate details](https://understandingwebpki.com/?cert=MIIDOjCCAt%2BgAwIBAgIUTe1uXF8VSGPfy0E4TIUy0Kk5KcowCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIxMDgxOTAzMjUwMloXDTMxMDgyMDAzMjUwMlowgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfbaQ7SzC9ijiFCviXk8%2BjN2lleKFSQsKWz%2BsvR5MYhlkg88KNM5Lf2Rlm0mSpeqIK%2BQHuQHCR5otmPXrGJBeQKOCASYwggEiMA8GA1UdEwEB%2FwQFMAMBAf8wHwYDVR0jBBgwFoAUFNWx4g9D2wES8o4AOlMKXL1z19AwXQYIKwYBBQUHAQEEUTBPME0GCCsGAQUFBzAChkFodHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkUm9vdENBLmNydDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSQAwRgIhAM%2B2sC%2FLU6AmmeZVeQanPH2cY75OpxZ9HEV8YFrEjiplAiEA%2FvF4l4gWmQrHfy4K7aWJIVKseiKDEgDKjSmdjnXikaI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | US_SHAKEN_CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |

### Not Effective

- e_atis_ca_certificate_policies
- e_atis_ca_crl_distribution
- e_atis_ca_extension_unknown
- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- n_atis_ca_certificate_policy_critical

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 11 Jan 23 23:18 UTC