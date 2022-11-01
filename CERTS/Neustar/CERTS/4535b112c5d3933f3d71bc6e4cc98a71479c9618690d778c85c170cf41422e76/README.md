# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 854J

Tested At: 01 Nov 22 20:30 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 182 day(s)\
Subject: CN=SHAKEN 854J, O=snetconnect, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11541.10167

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FDCCAqKgAwIBAgIURh5iwp7Dzfmnaj%2BXReIgGPphKNUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUwMjEzMTA0NFoXDTIzMDUwMjEzMTA0NFowOTELMAkGA1UEBhMCVVMxFDASBgNVBAoMC3NuZXRjb25uZWN0MRQwEgYDVQQDDAtTSEFLRU4gODU0SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKBdW6GiVzoadzW%2F61FO%2BpvMwxKXQuyVqjsngDzt3ePItbSjqyaWFx0MTiRWg5HKKpVP7BPeynhHPa3MV9NtP4ijggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEODU0SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFKk5W%2BMniW7yYKHc2H%2F1nwW8sNvaMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAzXlYGFT0oqFzL4ztP1DQ27bIvdMhcCYnBedhSDSltZUCIEtSnqbST%2BM3V4qPrbsP6CwoH2h0XH5fjQ3FK8N2t7L%2B)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 20:31:14