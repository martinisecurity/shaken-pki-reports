# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 772J

Tested At: 01 Mar 23 18:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 82 day(s)\
Subject: CN=SHAKEN 772J, O=VOIP ESSENTIAL INC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11540.10157

[View certificate details](https://understandingwebpki.com/?cert=MIIDAjCCAqmgAwIBAgIUejaugFMwXt866%2Fh%2FVJ5RENtbnDUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyMjE1MzQyN1oXDTIzMDUyMjE1MzQyN1owQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMElZPSVAgRVNTRU5USUFMIElOQzEUMBIGA1UEAwwLU0hBS0VOIDc3MkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAREwyXCvXhMwDghjPzD%2FjJ6Bz66xppjX2R4Xu704zofFmXICXMDA0PPtrQbrcBPwmmahkUgkpk4fJwBiMlPMu4fo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDc3MkowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTidp0mkEZEFcUq5oxro83GMKTFazAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgZ0oDga7j4U%2BKn177MhJPvJ36llxfEeZcjB4Hzf2qPvkCIDjtj2cKQQ8e9ODmY1JuAM2WP4U93idJjM8gl%2B%2FpwoVY)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01 Mar 23 18:22 UTC