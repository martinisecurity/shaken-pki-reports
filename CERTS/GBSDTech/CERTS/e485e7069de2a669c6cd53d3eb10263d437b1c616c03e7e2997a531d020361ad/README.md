# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ccctelecom.com

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 156 day(s)\
Subject: CN=ccctelecom.com, OU=VOIP, O=Carrier Connect Corp, ST=TX, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://www.ccctelecom.com/816K.pem

[View certificate details](https://x509.io/?cert=MIICrDCCAlOgAwIBAgICEQswCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNDAxMzAxOTE3MDVaFw0yNTAxMjgxOTE3MDVaMGExCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJUWDEdMBsGA1UECgwUQ2FycmllciBDb25uZWN0IENvcnAxDTALBgNVBAsMBFZPSVAxFzAVBgNVBAMMDmNjY3RlbGVjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEh%2BYzQz4x8MiNCSV7CyDM1I6WC6Y%2FoioJLzNEN1gH90wtJxLBzlLLNvA%2F9sqmfCW6fyNEwFEkpI53Hi2LTBvW3KOB%2BzCB%2BDAWBggrBgEFBQcBGgQKMAigBhYEODE2SzAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRiWo8FNciHQfawDK79TRNyJ07oQzAfBgNVHSMEGDAWgBR2xOfJOx4nEedS1E2tePaQBz6blDAOBgNVHQ8BAf8EBAMCB4AwZwYDVR0fBGAwXjBcoCGgH4YdaHR0cHM6Ly9zdGktcGEuY29tL3NoYWtlbi9jcmyiN6Q1MDMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEExEzARBgNVBAMMClNUSS1QQSBDUkwwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMAoGCCqGSM49BAMCA0cAMEQCIDirXr%2BhaxTIIRQbt7tlJJdBAjSGopbob1frMnIgMG9fAiAzl4bO7YsHfcs6idINxswdwvdnGj9R5cqemoF%2FExJzYw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'ccctelecom.com' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 816K', but common name is 'ccctelecom.com' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |


Generated: 26 Aug 24 18:03 UTC