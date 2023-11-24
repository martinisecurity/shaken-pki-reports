# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-8468

Tested At: 24 Nov 23 11:07 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 165 day(s)\
Subject: C=US, ST=NY, L=Albany, O=Firstlight, OU=VoIP, CN=SHAKEN-8468\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11260.10189.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID8DCCAtigAwIBAgIUP3TRzaCNW3sjU1Nsf8lBUPyyqPEwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTA1MDYxNTUwNTlaFw0yNDA1MDYxNTUwNTlaMGUxFDASBgNVBAMMC1NIQUtFTi04NDY4MQ0wCwYDVQQLDARWb0lQMRMwEQYDVQQKDApGaXJzdGxpZ2h0MQ8wDQYDVQQHDAZBbGJhbnkxCzAJBgNVBAgMAk5ZMQswCQYDVQQGEwJVUzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAM8TR2ww2m%2Fh8OK8ivCY99lbRSmJ69rEbQiZmJ4H%2Fte4rL%2FANpCYvHfeHA9p31kGk95zfA%2B9Pf1Vv67mwB1iSWjggFIMIIBRDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFDu5XMsxF5xIz9TaZuuNtqjV4Qn3MIGBBggrBgEFBQcBAQR1MHMwRwYIKwYBBQUHMAKGO2h0dHA6Ly9jYWNlcnRzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MCgGCCsGAQUFBzABhhxodHRwOi8vb2NzcC1jYTEuY2NpZC5uZXVzdGFyMEgGA1UdHwRBMD8wPaA7oDmGN2h0dHA6Ly9jcmwuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcmwwHQYDVR0OBBYEFHY7G4aASW8hAhxf7pVBmiLW1H5IMA4GA1UdDwEB%2FwQEAwIHgDAWBggrBgEFBQcBGgQKMAigBhYEODQ2ODANBgkqhkiG9w0BAQsFAAOCAQEAIN%2FQRzAIvD3ndLrBRCIO6yGDSUQ9rMD%2FSpuiv9ZX8ielqy0eIvJCdxprEmcrXY3tf7Hct82lCvPKUdSRcyOcFKhvtGlaroYHzZ7FWQtF5i0VGYLcdXF7tIto%2Ffnn2Gz2uy1GN%2FFRAA2gki%2BOLzIp4c9ffT71MKMjHHt3GaMLeYJxmqL2%2B0oy3pLTU84NDuoW82VL1S6luiKnj%2FNMJcaN9vDkvQGzYEJmA5aziYRnInNMSHqqsGycD7dAmxhKPo9ycoeQrwUo4zBav4PNwgh5%2B94f5nCU5T9FqsILc5IbRrgkzZPtNrs98BVu%2FFmNqHdPoKqQVDN4%2Fjg4NmcPSkhcig%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |

### Not Effective

- e_atis_ext_crl_distribution_struct
- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC