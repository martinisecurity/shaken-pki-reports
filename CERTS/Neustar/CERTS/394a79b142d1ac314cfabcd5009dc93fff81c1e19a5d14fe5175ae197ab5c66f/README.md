# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Firstcomm5917

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 510 day(s)\
Subject: C=US, ST=IL, L=Oakbrook, O=Firstcomm, OU=Firstcomm.com, CN=Firstcomm5917\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11414.10183.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2FDCCAuSgAwIBAgIUMAcRlds%2BL3YCpQyVEuWjcD67tacwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMjQxNDAyMDVaFw0yNDAzMjQxNDAyMDVaMHExFjAUBgNVBAMMDUZpcnN0Y29tbTU5MTcxFjAUBgNVBAsMDUZpcnN0Y29tbS5jb20xEjAQBgNVBAoMCUZpcnN0Y29tbTERMA8GA1UEBwwIT2FrYnJvb2sxCzAJBgNVBAgMAklMMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABItmwV6bFtQt0Wxb3k8PZLwfF%2BZBfnLpb7goQ3qwynOfNcLgGLeoIIeKut1%2BnWxkZiW5yJhMya3GVM%2F6Mi%2BHrGajggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFMA%2FUzexPNsgdZtjWcpXOBQUELYmMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENTkxNzANBgkqhkiG9w0BAQsFAAOCAQEAZPf60FbZ2FnGtmXq1ZSKlURo0%2FP4fF5hArft6TmuUTBlugk0a0cTFKK1UJb%2F3fnNtnGVRz3WprkNlvsSfIdGGphxPgyOeuWNLp4EOMz0erK8wcYVtEz%2B3C60yghZ8Lg2A%2BrtGlXwosLwMPKFApDjQzYLvlRjzbYejjNJLUXcZBQagD36yDHbeYxx3Oy4KADDpbaeK8zHbKeKDqcQg%2BFIQaLQ23qonZLXv5fObR3vRUWE3CT2MvdMiyFcqg7A5sM%2FmzKxB65tMV2GPUiAK4%2FKo1%2FgdC1%2Bj7W%2B%2FGGzalZSTD83SC0sVMs3SYgrQzojHhL9JDLWx4l1T%2F6k3fGHpysDjA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_crl_distribution_not_reachable](../../ISSUES/e_sti_crl_distribution_not_reachable/README.md) | error | ATIS&#x2011;1000080 | Unable to retrieve CRL specified in CRLdp from allow listed IP address |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22