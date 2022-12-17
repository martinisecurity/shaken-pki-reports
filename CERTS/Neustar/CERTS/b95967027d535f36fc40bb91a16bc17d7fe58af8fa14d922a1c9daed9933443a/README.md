# STIR/SHAKEN CA Ecosystem Compliance

## Certificate digitalipvoice.com

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 205 day(s)\
Subject: C=US, ST=FL, L=Tampa, O=Digital IP Voice, OU=Service Division, CN=digitalipvoice.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://az.tax/1

[View certificate details](https://understandingwebpki.com/?cert=MIIECDCCAvCgAwIBAgIUK%2BNEp8Dd1pjyaeipxeJtqZnWW%2BIwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA3MDkxNTI5MzhaFw0yMzA3MTAxNTI5MzhaMH0xGzAZBgNVBAMMEmRpZ2l0YWxpcHZvaWNlLmNvbTEZMBcGA1UECwwQU2VydmljZSBEaXZpc2lvbjEZMBcGA1UECgwQRGlnaXRhbCBJUCBWb2ljZTEOMAwGA1UEBwwFVGFtcGExCzAJBgNVBAgMAkZMMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJApmqRN8oGRcuWQwIxUhoea6e%2BulhGdvmnWb5T3lZOZCc3p4QyUTn4Ye2iCohv56SvFZ2BNFjn9jo8iVOixPTmjggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFA%2BRMcVlit8M7AxIrYdlV0HhDTaBMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENTIwRjANBgkqhkiG9w0BAQsFAAOCAQEAgdLS5S2N%2F2kZwSZ5lJrX8YaPTuv0p8kstvw4S5%2FQGzquNn%2BL7jHS8ptxIHOlChY1aiU1Je4f3c7wfz4ibbmrCE3yFEMNwGcFyTaCCdKExBZz9djOLg4nArQxmYGJ9IM9MHDEeHLbrgIwfn%2BQa4OfICeI7H8pJgjGP2eOtlQMQVgA2XCscMdjowMrYfAR%2BJQzejKOg9GiKRPwzIg8Ol7MzVViGjoBbWjHmZkJ9YF%2FzcLD0RBTaPir0seXYGm%2FS5VbOpncD1fnR16ukRfFkFPMH3Uos5zJtAZtT%2BzcVvGjAEHy4eUSjTUsZ29buNUaYeTNIhB6uDO3fHtr2xVA4%2FzOiw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_authority_key_identifier
- e_atis_basic_constraints
- e_atis_certificate_policies
- e_atis_crl_distribution
- e_atis_extension_unknown
- e_atis_issuer_dn
- e_atis_key_usage
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject
- e_atis_subject_cn
- e_atis_subject_key_identifier
- e_atis_subject_public_key
- e_atis_tn_auth_list
- e_atis_version

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 17 Dec 22 17:07 UTC