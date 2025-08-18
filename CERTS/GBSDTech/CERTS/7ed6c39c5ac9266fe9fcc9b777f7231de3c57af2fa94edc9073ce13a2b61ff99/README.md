# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ccctelecom.com

Tested At: 18 Aug 25 21:10 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 176 day(s)\
Subject: CN=ccctelecom.com, OU=VOIP, O=Carrier Connect Corp, ST=TX, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://www.ccctelecom.com/816K.pem

[View certificate details](https://x509.io/?cert=MIICzTCCAnOgAwIBAgICER8wCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTAyMTExODQyMTZaFw0yNjAyMTAxODQyMTZaMGExCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJUWDEdMBsGA1UECgwUQ2FycmllciBDb25uZWN0IENvcnAxDTALBgNVBAsMBFZPSVAxFzAVBgNVBAMMDmNjY3RlbGVjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEh%2BYzQz4x8MiNCSV7CyDM1I6WC6Y%2FoioJLzNEN1gH90wtJxLBzlLLNvA%2F9sqmfCW6fyNEwFEkpI53Hi2LTBvW3KOCARowggEWMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTZLMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFGJajwU1yIdB9rAMrv1NE3InTuhDMB8GA1UdIwQYMBaAFKdsSAmToL9B4BNhceYD7TWHHe6BMA4GA1UdDwEB%2FwQEAwIHgDCBhAYDVR0fBH0wezB5oD6gPIY6aHR0cHM6Ly9hdXRoZW50aWNhdGUtZXh0LWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKI3pDUwMzELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTETMBEGA1UEAwwKU1RJLVBBIENSTDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwCgYIKoZIzj0EAwIDSAAwRQIgBEcDp35zmZVGagPr6DiJ78Gp2xEfocOiNP7pgHUbt3kCIQCIIVuS9WzlM78ryMNvTJA5YIWmofcH4Wvw4SC%2FtM2BLg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'ccctelecom.com' does not contain 'SHAKEN' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 816K', but common name is 'ccctelecom.com' |


Generated: 18 Aug 25 21:13 UTC