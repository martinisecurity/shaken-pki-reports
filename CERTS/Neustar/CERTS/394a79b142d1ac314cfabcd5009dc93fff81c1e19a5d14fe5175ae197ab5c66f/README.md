# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Firstcomm5917

Tested At: 17 Dec 22 00:03 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 464 day(s)\
Subject: C=US, ST=IL, L=Oakbrook, O=Firstcomm, OU=Firstcomm.com, CN=Firstcomm5917\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11166.10183.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID%2FDCCAuSgAwIBAgIUMAcRlds%2BL3YCpQyVEuWjcD67tacwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMjQxNDAyMDVaFw0yNDAzMjQxNDAyMDVaMHExFjAUBgNVBAMMDUZpcnN0Y29tbTU5MTcxFjAUBgNVBAsMDUZpcnN0Y29tbS5jb20xEjAQBgNVBAoMCUZpcnN0Y29tbTERMA8GA1UEBwwIT2FrYnJvb2sxCzAJBgNVBAgMAklMMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABItmwV6bFtQt0Wxb3k8PZLwfF%2BZBfnLpb7goQ3qwynOfNcLgGLeoIIeKut1%2BnWxkZiW5yJhMya3GVM%2F6Mi%2BHrGajggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFMA%2FUzexPNsgdZtjWcpXOBQUELYmMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENTkxNzANBgkqhkiG9w0BAQsFAAOCAQEAZPf60FbZ2FnGtmXq1ZSKlURo0%2FP4fF5hArft6TmuUTBlugk0a0cTFKK1UJb%2F3fnNtnGVRz3WprkNlvsSfIdGGphxPgyOeuWNLp4EOMz0erK8wcYVtEz%2B3C60yghZ8Lg2A%2BrtGlXwosLwMPKFApDjQzYLvlRjzbYejjNJLUXcZBQagD36yDHbeYxx3Oy4KADDpbaeK8zHbKeKDqcQg%2BFIQaLQ23qonZLXv5fObR3vRUWE3CT2MvdMiyFcqg7A5sM%2FmzKxB65tMV2GPUiAK4%2FKo1%2FgdC1%2Bj7W%2B%2FGGzalZSTD83SC0sVMs3SYgrQzojHhL9JDLWx4l1T%2F6k3fGHpysDjA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Dec 22 00:12 UTC