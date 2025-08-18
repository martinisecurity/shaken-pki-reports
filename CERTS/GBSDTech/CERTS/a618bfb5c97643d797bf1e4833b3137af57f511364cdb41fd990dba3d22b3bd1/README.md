# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Allure Telecom Inc.

Tested At: 18 Aug 25 21:09 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 242 day(s)\
Subject: CN=Allure Telecom Inc., OU=VOIP, O=alluretelecom.com, ST=TX, C=US\
Issuer: CN=SHAKEN 1RouteGroup Intermediate CA, O=1RouteGroup, ST=Texas, C=US\
Link: https://stir.alluretelecom.com/static/846K-cert-2025.pem

[View certificate details](https://x509.io/?cert=MIIC0DCCAnWgAwIBAgICETUwCgYIKoZIzj0EAwIwYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAsxUm91dGVHcm91cDErMCkGA1UEAwwiU0hBS0VOIDFSb3V0ZUdyb3VwIEludGVybWVkaWF0ZSBDQTAeFw0yNTA0MTgwMzQyNDJaFw0yNjA0MTcwMzQyNDJaMGMxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJUWDEaMBgGA1UECgwRYWxsdXJldGVsZWNvbS5jb20xDTALBgNVBAsMBFZPSVAxHDAaBgNVBAMME0FsbHVyZSBUZWxlY29tIEluYy4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQkoPmgaDR6GWRKtjkoWk%2FUbl5uGtWJngHXHLBwRV%2FnOtRSZUrebbyY4N89lZ%2Bj41P74IlP%2Fs9lIw9YPbAIKrx6o4IBGjCCARYwFgYIKwYBBQUHARoECjAIoAYWBDg0NkswDAYDVR0TAQH%2FBAIwADAdBgNVHQ4EFgQUvHOLBgRw4qQc1qPQqs509tQB5dwwHwYDVR0jBBgwFoAUp2xICZOgv0HgE2Fx5gPtNYcd7oEwDgYDVR0PAQH%2FBAQDAgeAMIGEBgNVHR8EfTB7MHmgPqA8hjpodHRwczovL2F1dGhlbnRpY2F0ZS1leHQtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsojekNTAzMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMRMwEQYDVQQDDApTVEktUEEgQ1JMMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBBDAKBggqhkjOPQQDAgNJADBGAiEA8MMKVisGhDXmmhrnJP7ulxa%2BymRw4T4rXdTaIdmnOmECIQDlhYKsqwCvmq7%2F5OGQtokLUEFetNZYjc4tRvQcKM%2F5KQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 13 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Allure Telecom Inc.' does not contain 'SHAKEN' |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 846K', but common name is 'Allure Telecom Inc.' |


Generated: 18 Aug 25 21:13 UTC