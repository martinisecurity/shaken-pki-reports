# STIR/SHAKEN CA Ecosystem Compliance

## Certificate alluretelecom.com

Tested At: 22 Aug 24 16:03 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 243 day(s)\
Subject: CN=alluretelecom.com, OU=VOIP, O=Allure Telecom Inc., ST=TX, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://stir.alluretelecom.com/static/846K-cert.pem

[View certificate details](https://x509.io/?cert=MIICrzCCAlWgAwIBAgICERIwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNDA0MjIyMDQzNDFaFw0yNTA0MjEyMDQzNDFaMGMxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJUWDEcMBoGA1UECgwTQWxsdXJlIFRlbGVjb20gSW5jLjENMAsGA1UECwwEVk9JUDEaMBgGA1UEAwwRYWxsdXJldGVsZWNvbS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQkoPmgaDR6GWRKtjkoWk%2FUbl5uGtWJngHXHLBwRV%2FnOtRSZUrebbyY4N89lZ%2Bj41P74IlP%2Fs9lIw9YPbAIKrx6o4H7MIH4MBYGCCsGAQUFBwEaBAowCKAGFgQ4NDZLMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFLxziwYEcOKkHNaj0KrOdPbUAeXcMB8GA1UdIwQYMBaAFHbE58k7HicR51LUTa149pAHPpuUMA4GA1UdDwEB%2FwQEAwIHgDBnBgNVHR8EYDBeMFygIaAfhh1odHRwczovL3N0aS1wYS5jb20vc2hha2VuL2NybKI3pDUwMzELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTETMBEGA1UEAwwKU1RJLVBBIENSTDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQQwCgYIKoZIzj0EAwIDSAAwRQIgE0hBZlKxYPN9QJM9WY8YwZM%2B8I0Nqx%2B1VijRQ%2BFNItUCIQDC8gS0Wt8T57D4BS3r%2FAmAhmd6Nqpagem86BkAXhxf5w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 846K', but common name is 'alluretelecom.com' |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'alluretelecom.com' does not contain 'SHAKEN' |


Generated: 22 Aug 24 16:06 UTC