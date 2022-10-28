# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate b6f33eebd6fa1f397a22fe4d6300df28960f3061
Tested At: 2022-10-28 10:32:57 +0000 UTC\
Initial Validity Period: 3653 day(s)\
Remaining Validity Period: 3218 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: http://5.161.95.22/191c4c42dd7fa6115e84100637e42c99.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIIDOjCCAt%2BgAwIBAgIUTe1uXF8VSGPfy0E4TIUy0Kk5KcowCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIxMDgxOTAzMjUwMloXDTMxMDgyMDAzMjUwMlowgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfbaQ7SzC9ijiFCviXk8%2BjN2lleKFSQsKWz%2BsvR5MYhlkg88KNM5Lf2Rlm0mSpeqIK%2BQHuQHCR5otmPXrGJBeQKOCASYwggEiMA8GA1UdEwEB%2FwQFMAMBAf8wHwYDVR0jBBgwFoAUFNWx4g9D2wES8o4AOlMKXL1z19AwXQYIKwYBBQUHAQEEUTBPME0GCCsGAQUFBzAChkFodHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkUm9vdENBLmNydDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjAOBgNVHQ8BAf8EBAMCAYYwCgYIKoZIzj0EAwIDSQAwRgIhAM%2B2sC%2FLU6AmmeZVeQanPH2cY75OpxZ9HEV8YFrEjiplAiEA%2FvF4l4gWmQrHfy4K7aWJIVKseiKDEgDKjSmdjnXikaI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| n_pki_ca_key_usage | notice | SHAKEN PKI Best Practice | For CA certificates, the Key Usage extension should contain a single key usage value of keyCertSign |
| e_cp1_3_ca_key_usage_crl_sign | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| w_pki_ca_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- n_sti_ca_certificate_policy_critical
- e_sti_ca_extension_unknown
- e_sti_ca_certificate_policies
- e_sti_ca_serial_number
- e_sti_ca_crl_distribution
- e_sti_ca_subject_cn

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 10:33:25