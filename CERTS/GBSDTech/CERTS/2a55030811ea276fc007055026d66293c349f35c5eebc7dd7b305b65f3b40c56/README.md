# STIR/SHAKEN CA Ecosystem Compliance

## Certificate NovoLink SHAKEN cert

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 232 day(s)\
Subject: CN=NovoLink SHAKEN cert, O=NovoLink, ST=Texas, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: http://sti.novolink.net/crypt/sti-shaken.crt

[View certificate details](https://x509.io/?cert=MIICmzCCAkGgAwIBAgICERAwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNDA0MTUyMjI2MjFaFw0yNTA0MTQyMjI2MjFaME8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczERMA8GA1UECgwITm92b0xpbmsxHTAbBgNVBAMMFE5vdm9MaW5rIFNIQUtFTiBjZXJ0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYmHykqnZSXQCbNNoACrCRiBp8XuDOiow9VF3V1EU4jwknz507j2y%2FN4B1KQn5EoD70WFuEaRM1tcowmozZznGqOB%2BzCB%2BDAWBggrBgEFBQcBGgQKMAigBhYEMzMyRzAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQCOJhq9UJT5KxYBuwCcgbHmUE7BzAfBgNVHSMEGDAWgBR2xOfJOx4nEedS1E2tePaQBz6blDAOBgNVHQ8BAf8EBAMCB4AwZwYDVR0fBGAwXjBcoCGgH4YdaHR0cHM6Ly9zdGktcGEuY29tL3NoYWtlbi9jcmyiN6Q1MDMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMAoGCCqGSM49BAMCA0gAMEUCIFBoiO1jc%2BXav3hZVpeXVSfbNxBI5NH3xKgHr0a%2FkqFSAiEA1QM8FXQjDZzh4luwbsdNKIMpAM%2BIEnXtb%2Bg6qpFm6%2F0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 332G', but common name is 'NovoLink SHAKEN cert' |


Generated: 26 Aug 24 18:03 UTC