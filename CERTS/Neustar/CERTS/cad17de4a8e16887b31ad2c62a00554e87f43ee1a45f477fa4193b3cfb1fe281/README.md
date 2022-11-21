# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 235K

Tested At: 21 Nov 22 20:44 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 221 day(s)\
Subject: CN=SHAKEN 235K, O=TeligentIP Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11544.10197

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqWgAwIBAgIUS08%2FYtluzyhu6Tp1%2BVqbv%2BLGqaowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYzMDAxNTEzNloXDTIzMDYzMDAxNTEzNlowPDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlRlbGlnZW50SVAgSW5jMRQwEgYDVQQDDAtTSEFLRU4gMjM1SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABB827nKD2olR3c4eKwY1Bgwm11uCHbUttQ0zXqpSeySjgCytGa5Bghib5Pg7EwZJbHu1U%2BHscyu6sMdmd9AQFxyjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMjM1SzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFKvqVxTpRtyYmpfL3ktBAca3tUgfMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEApzPwKVf1CBO%2BkOoLQqdtH9630QW3IqcHBNBJprBU764CIQDbOCFXQgcJ34rvknCjt%2FoRrlEpaVuw7Vb3NIHnJJKbIw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 21 Nov 22 20:55 UTC