# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate ed757497d6919742972670344cc56fde4ff35f24
Tested At: 2022-10-28 19:22:08 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 76 day(s)\
Subject: CN=SHAKEN 506J, O=Twilio International, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11078.10126

View: [Click to view](https://understandingwebpki.com/?cert=MIIDBTCCAqugAwIBAgIUQrclyQ%2FvA%2FmnteKbMxSELqbAvNQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDExMTIxMzE0NFoXDTIzMDExMTIxMzE0NFowQjELMAkGA1UEBhMCVVMxHTAbBgNVBAoMFFR3aWxpbyBJbnRlcm5hdGlvbmFsMRQwEgYDVQQDDAtTSEFLRU4gNTA2SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA32hWF656C8sOMlkYV5R6aCKLTYMjShLIA3dpkzZnXWIDMGDjVVS56o%2BUMz3r1BpsZlAQN3DoZP58H1wfO8pnajggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENTA2SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFLCOB861kwYqaH7U1jQ4JFG3zhlWMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiB6vjpTmqioxGesvixucPlkoaZOQOxJFt2nl1oqeWojewIhAKU%2FsklN4beSJHt1jA0T13lamBos8Vcj7qISf5PwiiSK)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_sti_subject_cn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

Generated: 28/10/2022 at 19:22:10