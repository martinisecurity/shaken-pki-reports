# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 197D

Tested At: 17 Dec 22 17:06 UTC\
Initial Validity Period: 398 day(s)\
Remaining Validity Period: 39 day(s)\
Subject: CN=SHAKEN 197D, OU=Vonage Eng, O=Vonage, L=Holmdel, ST=NJ, C=US\
Issuer: CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA\
Link: https://sticr-cstga.ccid.neustar/api/v1/certificate/3f7be0c667b2f3c24546cf26629eb906.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDMTCCAtegAwIBAgIUaocVZNC6XofjlHGuqVZELngYmfUwCgYIKoZIzj0EAwIwgY8xCzAJBgNVBAYTAkNBMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEcMBoGA1UECwwTd3d3LmNhLmNjaWQubmV1c3RhcjE3MDUGA1UEAwwuTmV1c3RhciBDYW5hZGEgQ2VydGlmaWVkIENhbGxlciBJRCBTSEFLRU4gQ0EtMTAeFw0yMTEyMjIxODU3MjVaFw0yMzAxMjQxODU3MjVaMGgxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJOSjEQMA4GA1UEBwwHSG9sbWRlbDEPMA0GA1UECgwGVm9uYWdlMRMwEQYDVQQLDApWb25hZ2UgRW5nMRQwEgYDVQQDDAtTSEFLRU4gMTk3RDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE%2F7LBgwvnfzRI5s6wMc5odTj19J4cfnwCtTOLTZwVxsbEOyAchQHxMuQauR1j4x64u267WLXK4R5SmboQo4RzCjggE1MIIBMTAWBggrBgEFBQcBGgQKMAigBhYEMTk3RDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFFPuVX0JqdTNaA6xmavSADV%2FCEq4MGAGCCsGAQUFBwEBBFQwUjBQBggrBgEFBQcwAoZEaHR0cDovL2NhY2VydHMtY2EuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENhbmFkYUNBMS5jcnQwFQYDVR0gBA4wDDAKBggrBgEEAYO3HzBABgNVHR8EOTA3MDWgM6Axhi9odHRwczovL3N0aXBhLWNybC1jc3RnYS5jY2lkLm5ldXN0YXIvYXBpL3YxL2NybDAdBgNVHQ4EFgQUlxOKecNMVUIy0It0s1j%2BZ7JNJGYwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIEf9V0jPWdCp2S%2Fmiaz1Xjxhw1yZtAaLpA9vr2vDxkdDAiEAqrc0DxBE7vsuVkakKdhtpnQpJJ%2B7wMBMy%2B986XPJPOA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Dec 22 17:07 UTC