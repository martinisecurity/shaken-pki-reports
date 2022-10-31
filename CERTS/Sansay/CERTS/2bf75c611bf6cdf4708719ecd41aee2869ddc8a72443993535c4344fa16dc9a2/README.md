# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Drop Inc 258K

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 12 day(s)\
Subject: CN=SHAKEN Drop Inc 258K, OU=Drop, O=Drop Inc, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/258K_CERT

View: [Click to view](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT6owCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMjE3MzUwMFoXDTIyMTExMTE3MzUwMFowYTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQKDAhEcm9wIEluYzENMAsGA1UECwwERHJvcDEdMBsGA1UEAwwUU0hBS0VOIERyb3AgSW5jIDI1OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOhFkDpXxZUkcO1A7tlPq22zAs8qdt3ksBa096kRLfKB0tr5bmzJLlU6etipasInle1oSQ%2F39yJYnT2kWR2rjpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyNThLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUWxCrTIltzBjbumn9usIpJtYOgvEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBJCOL6ATBDNtTQ82B3rCH4dEPkw7F67XhWRzrKjJA%2F6AiEAtLY8ujgI0JdZaa4DpmwfnBIfpw0ejRyKK7q3gBSG25g%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 258K' |


Generated: 31/10/2022 at 16:43:22