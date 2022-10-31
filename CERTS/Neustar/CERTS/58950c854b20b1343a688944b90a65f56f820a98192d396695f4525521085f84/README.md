# STIR/SHAKEN CA Ecosystem Compliance

## Certificate peerlessnetwork.com

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 69 day(s)\
Subject: C=US, ST=IL, L=Chicago, O=support, OU=Peerless Network, CN=peerlessnetwork.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11184.10174.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIEAjCCAuqgAwIBAgIUcIf6QEhg69Xwg2B0nNtyD%2B3TIzcwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAxMDcxODQ4MzNaFw0yMzAxMDcxODQ4MzNaMHcxHDAaBgNVBAMME3BlZXJsZXNzbmV0d29yay5jb20xGTAXBgNVBAsMEFBlZXJsZXNzIE5ldHdvcmsxEDAOBgNVBAoMB3N1cHBvcnQxEDAOBgNVBAcMB0NoaWNhZ28xCzAJBgNVBAgMAklMMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABO4ice4Tf1c9OX3k0ESdSCoi%2BbHAPoIXqb%2FyD9T4oldyyPsSaMd%2BdUC1Gr4JlOlFp5YjS6C1wlHgTUwECfOgTcmjggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFJporv6CzOFv0%2BKuNFLh3PN6UJu5MA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYEMDYzRTANBgkqhkiG9w0BAQsFAAOCAQEAY1vClbOUlsU%2FcBzUsIM5KsR36FcMCg2sZeqxUhDyaAsIvP163pTH0jAC159gT6cKq2KN7oH2VFZ6PX5qaAMp10o8iYZh4YcK8XvROCBOn%2BcZkDOFkNIZ8Zr3RZ8eKgXQCBVSTECLQogClbf7IXcpjkeq2CiXvTPe7zGEy55ahNiZ0bTCi0Y0i1PC89s6fkuOVSFkt4mP%2BcJiZC5s7cLIW2CpsB8eP19FsnGE8S%2BTrHEtFkG7YndtclaghKDqa55YWpY4AKZGuiNFXJRAivV%2B0pKkZi3bNboT4ITzSA6eD5vrMNPd9t0MSBRZHQupo6q5QQFdO%2B9G5xG7M2GeydwRXg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_authority_key_identifier
- e_sti_basic_constraints
- e_sti_certificate_policies
- e_sti_crl_distribution
- e_sti_crl_distribution_not_reachable
- e_sti_extension_unknown
- e_sti_issuer_dn
- e_sti_key_usage
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject
- e_sti_subject_cn
- e_sti_subject_key_identifier
- e_sti_subject_public_key
- e_sti_tn_auth_list
- e_sti_version
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22