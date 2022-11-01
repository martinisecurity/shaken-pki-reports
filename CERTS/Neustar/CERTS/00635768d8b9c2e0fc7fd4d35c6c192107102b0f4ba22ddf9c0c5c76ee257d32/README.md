# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 506J

Tested At: 01 Nov 22 16:28 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 72 day(s)\
Subject: CN=SHAKEN 506J, O=Twilio International, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11078.10126

View: [Click to view](https://understandingwebpki.com/?cert=MIIDBTCCAqugAwIBAgIUQrclyQ%2FvA%2FmnteKbMxSELqbAvNQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDExMTIxMzE0NFoXDTIzMDExMTIxMzE0NFowQjELMAkGA1UEBhMCVVMxHTAbBgNVBAoMFFR3aWxpbyBJbnRlcm5hdGlvbmFsMRQwEgYDVQQDDAtTSEFLRU4gNTA2SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA32hWF656C8sOMlkYV5R6aCKLTYMjShLIA3dpkzZnXWIDMGDjVVS56o%2BUMz3r1BpsZlAQN3DoZP58H1wfO8pnajggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENTA2SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFLCOB861kwYqaH7U1jQ4JFG3zhlWMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiB6vjpTmqioxGesvixucPlkoaZOQOxJFt2nl1oqeWojewIhAKU%2FsklN4beSJHt1jA0T13lamBos8Vcj7qISf5PwiiSK)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 16:30:07