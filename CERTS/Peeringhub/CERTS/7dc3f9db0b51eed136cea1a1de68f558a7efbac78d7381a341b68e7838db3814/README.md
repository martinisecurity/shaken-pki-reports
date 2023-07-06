# STIR/SHAKEN CA Ecosystem Compliance

## Certificate VOCALTRANSIT SHAKEN 783J v3

Tested At: 06 Jul 23 14:06 UTC\
Initial Validity Period: 167 day(s)\
Remaining Validity Period: 132 day(s)\
Subject: CN=VOCALTRANSIT SHAKEN 783J v3, OU=VOCALTRANSIT, O=IT Vocal LLC, ST=NV, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://shaken.vocaltransit.com/783J/vocaltransit-STIRSHAKEN.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDJjCCAs2gAwIBAgIRAJDzg%2BQH4Z9l6xDdqUAZddIwCgYIKoZIzj0EAwIwfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwHhcNMjMwNjAxMDAwMDAwWhcNMjMxMTE1MDAwMDAwWjBuMQswCQYDVQQGEwJVUzELMAkGA1UECAwCTlYxFTATBgNVBAoMDElUIFZvY2FsIExMQzEVMBMGA1UECwwMVk9DQUxUUkFOU0lUMSQwIgYDVQQDDBtWT0NBTFRSQU5TSVQgU0hBS0VOIDc4M0ogdjMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATTQWOOmk8j1AIIYU82s7oRmBbl4EcKa1bK145vG%2FVoVprXLAiEvWT9AvCRoE0kJll14d9sW0PtgVbf7QkxBNDJo4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFMwTdYzavf7km%2FEpH9%2FOV3%2FBGAIlMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENzgzSjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDRwAwRAIgSNmKOV0Qqa%2FUo4zay7tAh3q7zlGGM%2FTkKqh8P1w5Ke0CIDxkMlcGCX%2BH%2FLLgkcUaqMgV6tYr95oE4KvpkPPe2lNA)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC