# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Drop Inc 258K

Tested At: 30 Jan 23 22:59 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 314 day(s)\
Subject: CN=SHAKEN Drop Inc 258K, OU=Drop, O=Drop Inc, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: http://sti.comsapi.com/258k/ca.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkV%2F0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIxMDAyMTEyN1oXDTIzMTIxMDAyMTEyN1owYTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQKDAhEcm9wIEluYzENMAsGA1UECwwERHJvcDEdMBsGA1UEAwwUU0hBS0VOIERyb3AgSW5jIDI1OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOhFkDpXxZUkcO1A7tlPq22zAs8qdt3ksBa096kRLfKB0tr5bmzJLlU6etipasInle1oSQ%2F39yJYnT2kWR2rjpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyNThLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUWxCrTIltzBjbumn9usIpJtYOgvEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIDA5%2BJoD%2Fr0kwtz%2FAvvl6%2Bruw0X9vPGVqRL254%2FWqHwKAiEA2VSh45sMLCvrlUyER6waslSq3%2FZeKynFDhLtoXiAOt0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 258K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 30 Jan 23 23:10 UTC