# STIR/SHAKEN CA Ecosystem Compliance

## Certificate cox.com

Tested At: 09 Mar 23 23:16 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 17 day(s)\
Subject: C=US, ST=GA, L=Atlanta, O=Communications, OU=Cox Communications, CN=cox.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11126.10191

[View certificate details](https://understandingwebpki.com/?cert=MIID%2FzCCAuegAwIBAgIUKKoQnojqJjuV0GIgL%2BXFsXcsmF8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAzMjUxNzI5MTdaFw0yMzAzMjYxNzI5MTdaMHQxEDAOBgNVBAMMB2NveC5jb20xGzAZBgNVBAsMEkNveCBDb21tdW5pY2F0aW9uczEXMBUGA1UECgwOQ29tbXVuaWNhdGlvbnMxEDAOBgNVBAcMB0F0bGFudGExCzAJBgNVBAgMAkdBMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEUeqnRCPKZwBnDeqECK%2Fdy1dwljc92%2By%2BL9MDvrxe73Y%2FmG2K68ovePEp1jWMOh4KGaCPvBsCBEtLq5SPvLDwajggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFG8opCW0%2F5uCiJ9j1jdkUFLhUkPGMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYENzY2MTANBgkqhkiG9w0BAQsFAAOCAQEAY4aP21H9sLT4WUh6L%2FCjPi%2B8FMGxtnsiwAAZb0z1iAKibsPJMP%2FYvK74PIpg0aZFiU1f5pFR0IWDa75KwJem4R47gZ2%2Bk3sde9unAF3609spDlxOFoZdf%2B6SbQawter6d9FbAx6yeRJEvx5ssQDl%2FVq5jU%2BTI0tch6X1RL7CAvmsmqLY26%2Bi7vQxLaKSgDX1iDMeEWEejLr2%2FiGKSXsbnaiF2qm1tTfoRCdQ00Cmw2xXj80%2BoGr6gNDurpE2BwZMquDBM3rHZH1T3%2BrhN1qrVB%2Bh0s3iO4RQpq8oD4RfDxn%2FW0r1sX8YtybDIvlSdmbb4cBqs6z5hTmnpjh%2F2eZV1Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
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
- e_us_cp_ambiguous_identifier
- e_us_cp_subject_sn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 10 Mar 23 02:25 UTC