# STIR/SHAKEN CA Ecosystem Compliance

## Certificate peerlessnetwork.com

Tested At: 21 Nov 22 20:21 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 47 day(s)\
Subject: C=US, ST=IL, L=Chicago, O=support, OU=Peerless Network, CN=peerlessnetwork.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11034.10174

[View certificate details](https://understandingwebpki.com/?cert=MIIEAjCCAuqgAwIBAgIUcIf6QEhg69Xwg2B0nNtyD%2B3TIzcwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAxMDcxODQ4MzNaFw0yMzAxMDcxODQ4MzNaMHcxHDAaBgNVBAMME3BlZXJsZXNzbmV0d29yay5jb20xGTAXBgNVBAsMEFBlZXJsZXNzIE5ldHdvcmsxEDAOBgNVBAoMB3N1cHBvcnQxEDAOBgNVBAcMB0NoaWNhZ28xCzAJBgNVBAgMAklMMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABO4ice4Tf1c9OX3k0ESdSCoi%2BbHAPoIXqb%2FyD9T4oldyyPsSaMd%2BdUC1Gr4JlOlFp5YjS6C1wlHgTUwECfOgTcmjggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFJporv6CzOFv0%2BKuNFLh3PN6UJu5MA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYEMDYzRTANBgkqhkiG9w0BAQsFAAOCAQEAY1vClbOUlsU%2FcBzUsIM5KsR36FcMCg2sZeqxUhDyaAsIvP163pTH0jAC159gT6cKq2KN7oH2VFZ6PX5qaAMp10o8iYZh4YcK8XvROCBOn%2BcZkDOFkNIZ8Zr3RZ8eKgXQCBVSTECLQogClbf7IXcpjkeq2CiXvTPe7zGEy55ahNiZ0bTCi0Y0i1PC89s6fkuOVSFkt4mP%2BcJiZC5s7cLIW2CpsB8eP19FsnGE8S%2BTrHEtFkG7YndtclaghKDqa55YWpY4AKZGuiNFXJRAivV%2B0pKkZi3bNboT4ITzSA6eD5vrMNPd9t0MSBRZHQupo6q5QQFdO%2B9G5xG7M2GeydwRXg%3D%3D)

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


Generated: 21 Nov 22 20:33 UTC