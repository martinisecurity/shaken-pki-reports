# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Drop Inc 258K

Tested At: 12 Feb 24 16:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 289 day(s)\
Subject: CN=SHAKEN Drop Inc 258K, OU=Drop, O=Drop Inc, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: http://sti.comsapi.com/258k/ca.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkh6cwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEyNzE3NTgxOFoXDTI0MTEyNjE3NTgxOFowYTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQKDAhEcm9wIEluYzENMAsGA1UECwwERHJvcDEdMBsGA1UEAwwUU0hBS0VOIERyb3AgSW5jIDI1OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOhFkDpXxZUkcO1A7tlPq22zAs8qdt3ksBa096kRLfKB0tr5bmzJLlU6etipasInle1oSQ%2F39yJYnT2kWR2rjpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyNThLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUWxCrTIltzBjbumn9usIpJtYOgvEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCXBC1DAtuSMiLeFu2jJhrybPt6xWJ2BjxWbIylS7y2ygIgIPd5GU7etQBkdhbFHo7%2FqkTMBs%2FvZ8607GMwAS9oIlc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 258K', but common name is 'SHAKEN Drop Inc 258K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 12 Feb 24 17:02 UTC