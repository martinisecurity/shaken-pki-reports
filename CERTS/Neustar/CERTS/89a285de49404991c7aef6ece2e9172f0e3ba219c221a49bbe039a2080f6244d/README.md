# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 963J

Tested At: 02 Nov 22 15:13 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 155 day(s)\
Subject: CN=SHAKEN 963J, O=Freevoice, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: http://sip.phx.dlr.freevoicepbx.com/Freevoice_963J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIUEiVRu2gUPNmWiCu8l%2BfixE2zrsswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDQwNjEyMDYzNloXDTIzMDQwNjEyMDYzNlowNzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCUZyZWV2b2ljZTEUMBIGA1UEAwwLU0hBS0VOIDk2M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARLcL8TExlVPj7%2BFRrdZ8Q754VSJeAoXK7JKvt%2BSbikFPrUmCpBy33cgDsLYLeVIkM2THQSUInOTKHYG5XFprxRo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDk2M0owDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTxvTbktYDecV4VEeG%2BbPxUoF8rTTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJjkJ9Bte32omuApuPbBpUN3CvIvImelAyxjnbtCzhDpAiEAut1xWpq0K%2F0qd6vtjwCKB%2BysNfSye%2B80me5hgtvU0wo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Nov 22 15:15 UTC