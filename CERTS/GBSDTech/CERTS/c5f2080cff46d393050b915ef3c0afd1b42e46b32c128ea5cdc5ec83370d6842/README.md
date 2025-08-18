# STIR/SHAKEN CA Ecosystem Compliance

## Certificate INTERBEL TELEPHONE SHAKEN

Tested At: 18 Aug 25 20:16 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 296 day(s)\
Subject: CN=INTERBEL TELEPHONE SHAKEN, O=Interbel Telephone, ST=MT, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://certs.iverify-aninetworks.net/interbeltelephone_20260610.crt

[View certificate details](https://x509.io/?cert=MIICyDCCAm2gAwIBAgICEXswCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTA2MTExNzExMjFaFw0yNjA2MTAxNzExMjFaMFsxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJNVDEbMBkGA1UECgwSSW50ZXJiZWwgVGVsZXBob25lMSIwIAYDVQQDDBlJTlRFUkJFTCBURUxFUEhPTkUgU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnLvnLpYVrqhm%2FkvnLz96Z%2BATjIBeDoVeCSc8eg3jBKcZwte1G7bBFNbD28DtBHOo7hjTWI6IGTNSkb037mkIYqOCARowggEWMBYGCCsGAQUFBwEaBAowCKAGFgQyMjQyMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFJD65PDzDJyUMf0sNa2smviZyzDkMB8GA1UdIwQYMBaAFKdsSAmToL9B4BNhceYD7TWHHe6BMA4GA1UdDwEB%2FwQEAwIHgDCBhAYDVR0fBH0wezB5oD6gPIY6aHR0cHM6Ly9hdXRoZW50aWNhdGUtZXh0LWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKI3pDUwMzELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTETMBEGA1UEAwwKU1RJLVBBIENSTDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwCgYIKoZIzj0EAwIDSQAwRgIhAOv2QfyyaRantMMMcfzOgWcAbTupcAKmLWy9aFgvGhhAAiEA4SyvJrEZamOw5OhLIIDUGAvSfDF134CEfa6EF1WB2D8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 2242', but common name is 'INTERBEL TELEPHONE SHAKEN' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |


Generated: 18 Aug 25 21:13 UTC