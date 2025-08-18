# STIR/SHAKEN CA Ecosystem Compliance

## Certificate allaccesstelecom.com

Tested At: 18 Aug 25 21:05 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=allaccesstelecom.com, OU=VOIP, O=All Access Telecom Inc., ST=TX, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://q2c.allaccesstelecom.com/static/855K-cert.pem

[View certificate details](https://x509.io/?cert=MIIC1jCCAnygAwIBAgICERkwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNDA4MjYyMDE0MDNaFw0yNTA4MjUyMDE0MDNaMGoxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJUWDEgMB4GA1UECgwXQWxsIEFjY2VzcyBUZWxlY29tIEluYy4xDTALBgNVBAsMBFZPSVAxHTAbBgNVBAMMFGFsbGFjY2Vzc3RlbGVjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEcoYeLX0bTXmbmzCH97TnBv%2Bf%2BBSQU%2BNryT2nYprOePBNZz64nJwoEuXD6vErGs%2FWlcWpcC4BN2lN%2FuGGSh9aqqOCARowggEWMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTVLMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFB2GPfFWcgIrT1SIfWbyfn0gAfwnMB8GA1UdIwQYMBaAFKdsSAmToL9B4BNhceYD7TWHHe6BMA4GA1UdDwEB%2FwQEAwIHgDCBhAYDVR0fBH0wezB5oD6gPIY6aHR0cHM6Ly9hdXRoZW50aWNhdGUtZXh0LWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKI3pDUwMzELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTETMBEGA1UEAwwKU1RJLVBBIENSTDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwCgYIKoZIzj0EAwIDSAAwRQIhAMPmR52LyZnPLx2wjaCIJs%2FP2b%2BhG7ibQ7vRU6L2fDOIAiAJ44KYJVC%2FhWOXsOecN6N%2FuCmaTynzc0yS%2BFilfudiBg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'allaccesstelecom.com' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 855K', but common name is 'allaccesstelecom.com' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |


Generated: 18 Aug 25 21:13 UTC