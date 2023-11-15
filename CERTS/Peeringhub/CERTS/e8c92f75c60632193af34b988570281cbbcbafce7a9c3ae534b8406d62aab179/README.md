# STIR/SHAKEN CA Ecosystem Compliance

## Certificate VOCALTRANSIT SHAKEN 783J

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 205 day(s)\
Remaining Validity Period: 189 day(s)\
Subject: CN=VOCALTRANSIT SHAKEN 783J, OU=VOCALTRANSIT, O=IT Vocal LLC, ST=NV, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDJDCCAsqgAwIBAgIRAOOo7x%2B%2FCZOtxzEnsNEMiOYwCgYIKoZIzj0EAwIwfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwHhcNMjMxMDMwMDAwMDAwWhcNMjQwNTIyMDAwMDAwWjBrMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTlYxFTATBgNVBAoMDElUIFZvY2FsIExMQzEVMBMGA1UECwwMVk9DQUxUUkFOU0lUMSEwHwYDVQQDDBhWT0NBTFRSQU5TSVQgU0hBS0VOIDc4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATTQWOOmk8j1AIIYU82s7oRmBbl4EcKa1bK145vG%2FVoVprXLAiEvWT9AvCRoE0kJll14d9sW0PtgVbf7QkxBNDJo4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFMwTdYzavf7km%2FEpH9%2FOV3%2FBGAIlMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzgzSjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSAAwRQIhALg2vcj%2BG1RUoIL5R5Kk8bE3dF25tyvF3Fd5snZ9RAC8AiBnz5%2BB3P4ejjO8VKbp7zyaovsqBiO5LsxPIyzwbX2SRg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC