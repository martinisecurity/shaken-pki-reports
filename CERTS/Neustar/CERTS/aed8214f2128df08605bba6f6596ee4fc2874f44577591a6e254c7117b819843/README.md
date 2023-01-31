# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 067K

Tested At: 31 Jan 23 17:51 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 145 day(s)\
Subject: CN=SHAKEN 067K, O=Junction Networks Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://stir-shaken.jnctn.net/onsip-stir-shaken.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBjCCAqygAwIBAgIUfZ2VsrXjN1U0l4cInmJw6canPDYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyNDE5NTgyM1oXDTIzMDYyNDE5NTgyM1owQzELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFUp1bmN0aW9uIE5ldHdvcmtzIEluYzEUMBIGA1UEAwwLU0hBS0VOIDA2N0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATutQg8Lvz1o1JyTFN3i3Kkh2xlp4Tz6UF1mwZAhCKuQVUJyzuDwZa1ySxHsCnEtxR2M8DmhnpMivCtrPSWNguyo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDA2N0swDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRDbdvo4GTfaf7mPd7eWaEubNsfYDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJntdavsd78iQnuO9Bpxu8OqTanSND08z58ZTYeFFpzqAiA0sf9jEKSyda5Tq1PK0sBut237Zdp0XRDDiVmuja3uCw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 31 Jan 23 17:51 UTC