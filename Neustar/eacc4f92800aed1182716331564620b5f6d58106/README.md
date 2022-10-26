# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate eacc4f92800aed1182716331564620b5f6d58106
Tested At: 2022-10-26 20:32:11 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 309 day(s)\
Subject: CN=SHAKEN 500J, O=GoTo Communications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://pstn-cdn.live.gtc.goto.com/certs/stirshaken/goto-2022-09

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCDCCAq6gAwIBAgIUT1Yxbwi4CyxuFBnXZ2A1P6OD%2Fm8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgzMTE0NDc0NFoXDTIzMDgzMTE0NDc0NFowRTELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF0dvVG8gQ29tbXVuaWNhdGlvbnMgSW5jMRQwEgYDVQQDDAtTSEFLRU4gNTAwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM89YM02tGXm826qqkEQgq1l6oQQsOpmdEK44LYjnpryF%2B9Ovv%2FilPUeAZN0igyFzCjByfhpr1fi%2Bs8UPAeFjVWjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENTAwSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFBuc9QU5VZAbJlFPNsMryN%2FaOUnrMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBBPX5UAi7ya2F78XqMBe%2B3dw2VLkNYb%2BnwxtMV9645%2FgIhAOAjRSfhuaEmYaru%2FJB1R5fuGp3xXGqTNBBBlhy8UBOV)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |

Generated: 26/10/2022 at 20:32:17