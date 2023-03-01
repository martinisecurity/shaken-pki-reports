# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Gardonville

Tested At: 01 Mar 23 18:14 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 427 day(s)\
Subject: C=US, ST=MN, L=Brandon, O=Gardenville Telephone, OU=Gardenville Telephone, CN=Gardonville\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/45.207

[View certificate details](https://understandingwebpki.com/?cert=MIIEDjCCAvagAwIBAgIUI1JOMxbqREKPsNfYwgOJFk%2BfVUwwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA0MzAyMjQ1NTBaFw0yNDA0MzAyMjQ1NTBaMIGCMRQwEgYDVQQDDAtHYXJkb252aWxsZTEeMBwGA1UECwwVR2FyZGVudmlsbGUgVGVsZXBob25lMR4wHAYDVQQKDBVHYXJkZW52aWxsZSBUZWxlcGhvbmUxEDAOBgNVBAcMB0JyYW5kb24xCzAJBgNVBAgMAk1OMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCi41mDUVufuNPuOTri4SVieKx9G3RIw%2FivbL2eQSvFzIFhyZvWWMTd2J9SFFYYUohD%2BjRMl8ktxBOmq4hF1K5qjggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFK0EPKvemkRJ3plRhW2IeOAMoGuqMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYEMTM5NjANBgkqhkiG9w0BAQsFAAOCAQEAW%2FAExAiHDJnhOAoXP50NPYuEysaoy1duACDg6zuQ8LaOPkAtphPiJRRCfEQDU1lrQJyQKGn5l3KgX7SHxeBuZfNZWib7tY%2BiDtxBA96p%2B%2FMijLU%2BL1vQyRw13v9%2FyTcECybgIxd1HXPWPx30sl2xPHk10cokJASjMFS5muP5yhKmvum%2FWAXHiAirCWzTruZTchz62xmwYjL0YtqmPzfjhUyAAQLblunhs5muNoubWl%2BJ%2BMYa0Tyy%2FvGQareasjYwHekbpxJ4fBi4wSTrpBsmpacoax64st7xtHwCohO20NxQl2vYW2xURJ2V5wH5pm%2F0ieioXALaFZclN8SntmM63Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01 Mar 23 18:22 UTC