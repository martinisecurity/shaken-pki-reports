# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Orange

Tested At: 09 Mar 23 23:16 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 396 day(s)\
Subject: C=US, ST=VA, L=Courbevoie, O=Orange, OU=Orange Business Services, CN=Orange\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11180.10184.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID%2FzCCAuegAwIBAgIUC16nA1ub1c5sd5K84QVLtSufwCMwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA0MDgwOTQ3MDdaFw0yNDA0MDgwOTQ3MDdaMHQxDzANBgNVBAMMBk9yYW5nZTEhMB8GA1UECwwYT3JhbmdlIEJ1c2luZXNzIFNlcnZpY2VzMQ8wDQYDVQQKDAZPcmFuZ2UxEzARBgNVBAcMCkNvdXJiZXZvaWUxCzAJBgNVBAgMAlZBMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCpgk6ApznudOJKB%2FHdwCPc73%2Fw06hjEkqDZmUmdDvF%2Bk4gjcVSVDUjxG15VOYa59GOuQmnIFYdbfTvEyZ9ogl6jggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFGqYvasPMVQB6Seb%2BJlEnTdf1X%2FxMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENzY2QzANBgkqhkiG9w0BAQsFAAOCAQEACfNrFIjiPlDUvxOCmKQrlp3Lv%2FBDx7ACC%2Frw%2FVK83F%2F8AK0PHs89d6O9iUoRXuMkLGFaVgtYO0xWy9dWhX%2FA6JU71D89gj5pXLj8u75%2BE73KJEp2krnnTMWSUrA7tK0zeWESBzk4ieN5t6pfZ9eDkmdR4heXP5khhaxszrHIOZ9E%2FUYV44CNw%2B8PWZE5gx3ZIxYe1RzVuve2kVqli8qZ8W630Ie2PeozwvN8FKeRotpkvk8uzQKRiIk80Q717X94LaO6lTZiNcuhzBI%2F8ssrC0JP99R2nGdutCDJ7Fq4gVMRpHT0Covl2mBxwJ4hVWNz9Q1KBYDy3rglpqdl9BpEPA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 10 Mar 23 02:25 UTC