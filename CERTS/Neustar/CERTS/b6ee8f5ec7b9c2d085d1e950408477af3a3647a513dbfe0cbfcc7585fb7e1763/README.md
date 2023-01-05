# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 804J

Tested At: 05 Jan 23 18:25 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 72 day(s)\
Subject: CN=SHAKEN 804J, O=QuestBlue Systems Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://customer.questblue.com/assets/questblue_shaken.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIDBzCCAqygAwIBAgIUIycrA5UtthnuJ9WODqKyMNsWdsUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDMxODE4MTg0NFoXDTIzMDMxODE4MTg0NFowQzELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFVF1ZXN0Qmx1ZSBTeXN0ZW1zIEluYzEUMBIGA1UEAwwLU0hBS0VOIDgwNEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQcvKIf0dgo3qoIoJXvwbHGAzHsXOqjdOa%2F0eNL4LlZR9zikKrwep8PT108ZaQ2pBflzZusG5br%2B95bw5PzPM6ko4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDgwNEowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBQ5oTgChUPn5VURvDEtbKr3NY4cijAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJeq4RRef%2FF%2FSyqtfH9qaZCLtmJkeaXvRtSyIE2hMmW3AiEA92BSJkxmj4lqdUDT3ojFlNVHqad%2Bwt5Qe5yO6PxpGII%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 05 Jan 23 18:35 UTC