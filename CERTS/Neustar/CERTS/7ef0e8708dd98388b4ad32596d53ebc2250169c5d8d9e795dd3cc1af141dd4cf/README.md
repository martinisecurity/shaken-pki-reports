# STIR/SHAKEN CA Ecosystem Compliance

## Certificate prod SHAKEN 811J

Tested At: 30 Nov 22 15:51 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 121 day(s)\
Subject: CN=prod SHAKEN 811J, OU=Alianza, O=CaptionCall, L=Salt Lake City, ST=Utah, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://api.alianza.com/v2/stir-shaken/certs/2dd72c30-5e06-49a5-bbec-fff3cdf9444b/key.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDOzCCAuGgAwIBAgIUCiWr83SJjBWHPTZSyOf6we35ntUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDMzMDIyNTMwOVoXDTIzMDMzMDIyNTMwOVoweDELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxFzAVBgNVBAcMDlNhbHQgTGFrZSBDaXR5MRQwEgYDVQQKDAtDYXB0aW9uQ2FsbDEQMA4GA1UECwwHQWxpYW56YTEZMBcGA1UEAwwQcHJvZCBTSEFLRU4gODExSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKiSBxMGSmBV0kVlim5ASMUMC8ww3VGrM1T4hzbLezbd2f8JsMxdLNj60LgvNFLez6Y2QnZIiroBuiXS1Gku9g6jggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEODExSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFNK8qKfNXPjfgBaTD19iyhv1o2vMMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBTYU62JxhvxyItA%2BWDhG%2FkzTuUlUPEqkWMouNjiCisRgIhAJpFxRyME7NF6On34JrZsPFVrXx5AzT1wHqPcJogntai)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 30 Nov 22 16:07 UTC